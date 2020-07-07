package api

import (
	pb "entry_task/src/proto"
	"entry_task/src/server/httpserver/client"
	"entry_task/src/server/httpserver/code"
	"entry_task/src/server/httpserver/dto"
	"entry_task/src/tools"
	log "github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"time"
)

func Login(ctx *gin.Context) {

	// 参数绑定
	var userLoginInfo dto.UserLoginArg
	ctx.BindJSON(&userLoginInfo)
	log.Debugf("【http服务】收到用户登录HTTP API请求，用户名：%s,密码:%s", userLoginInfo.Username, userLoginInfo.Password)
	// 参数校验
	if userLoginInfo.Username == "" {
		panic(code.FormatException(code.RequiredParam, "username"))
	}

	if userLoginInfo.Password == "" {
		panic(code.FormatException(code.RequiredParam, "password"))
	}

	if len(userLoginInfo.Password) != 32 {
		panic(code.FormatException(code.ParamInvalid, "password"))
	}

	// 通过rpc api调用登录逻辑
	conn, err := client.GrpcPool.Get(ctx)
	defer conn.Close()
	resp, err := pb.NewUserServiceClient(conn).Login(ctx, &pb.UserLoginRequest{Username: userLoginInfo.Username, Password: userLoginInfo.Password})
	if err != nil {
		log.Error("【http服务】调用rpc接口出现错误，错误信息：", err.Error())
		panic(code.ServerException)
		return
	}

	log.Debugf("【http服务】调用rpc登录接口返回的结果：%+v", *resp)
	if resp.Code != code.CodeSuccess {
		tools.FailResponse(ctx, int(resp.Code), resp.Msg, nil)
		return
	}

	generateToken(ctx, userLoginInfo.Username)

}

func Logout(ctx *gin.Context) {
	tools.SuccessResponse(ctx, nil)
}

func GetUserInfo(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*tools.CustomClaims)
	conn, err := client.GrpcPool.Get(ctx)
	defer conn.Close()
	resp, err := pb.NewUserServiceClient(conn).GetUserInfo(ctx, &pb.CommonRequest{Username: claims.Username})
	if err != nil {
		log.Error("【http服务】调用rpc退出登录接口出现错误，错误信息：", err.Error())
		panic(code.ServerException)
		return
	}
	log.Debugf("【http服务】调用rpc获取用户信息接口返回的结果：%v", resp)
	if resp.Code != code.CodeSuccess {
		tools.FailResponse(ctx, int(resp.Code), resp.Msg, nil)
		return
	}
	tools.SuccessResponse(ctx,
		map[string]string{"username": resp.Username, "nickname": resp.Nickname, "profile_picture": resp.ProfilePicture})

}

func UpdateUserProfilePicture(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*tools.CustomClaims)
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Error("【http服务】获取头像文件参数file出现错误，错误信息：", err.Error())
	}
	imageName := tools.GenerateImgName(file.Filename)
	dst := path.Join(tools.GetHttpConfig().Image.SavePath, imageName)
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		log.Error("【http服务】上传图谱出现错误，错误信息：", err.Error())
	}
	imageURL := tools.GetHttpConfig().Image.PrefixUrl + "/" + dst
	resp, err := client.GetGRClient().EditUserInfo(ctx, &pb.EditUserInfoRequest{Username: claims.Username, ProfilePicture: imageURL})

	if err != nil {
		log.Error("【http服务】调用rpc修改用户头像接口出现错误，错误信息：", err.Error())
		panic(code.ServerException)
		return
	}
	log.Debugf("【http服务】调用rpc修改用户头像接口返回的结果：%v", resp)
	if resp.Code != code.CodeSuccess {
		tools.FailResponse(ctx, int(resp.Code), resp.Msg, nil)
		return
	}
	tools.SuccessResponse(ctx, nil)

}

func UpdateUserInfo(ctx *gin.Context) {
	// 参数绑定
	var userUpdateInfoArg dto.UserUpdateInfoArg
	ctx.BindJSON(&userUpdateInfoArg)
	// 参数校验
	if len(userUpdateInfoArg.Nickname) > 32 {
		panic(code.FormatException(code.ParamInvalid, "nickname"))
	}
	claims := ctx.MustGet("claims").(*tools.CustomClaims)
	resp, err := client.GetGRClient().EditUserInfo(ctx, &pb.EditUserInfoRequest{Username: claims.Username, Nickname: userUpdateInfoArg.Nickname})
	if err != nil {
		log.Error("【http服务】调用rpc修改用户信息接口出现错误，错误信息：", err.Error())
		panic(code.ServerException)
		return
	}
	log.Debugf("【http服务】调用rpc修改用户信息接口返回的结果：%v", resp)
	if resp.Code != code.CodeSuccess {
		tools.FailResponse(ctx, int(resp.Code), resp.Msg, nil)
		return
	}
	tools.SuccessResponse(ctx, nil)

}

func LoginHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
	})
}

func UserInfoPageHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user_info.html", gin.H{
	})
}

// 生成令牌
func generateToken(c *gin.Context, username string) {
	j := &tools.JWT{
		[]byte("newtrekYang"),
	}
	claims := tools.CustomClaims{
		username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekYang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		log.Errorf("【http】服务创建token发生异常，异常信息：%s", err.Error())
		panic(code.CreateTokenFail)
	}

	result := make(map[string]string)
	result["token"] = token
	tools.SuccessResponse(c, result)
	return
}
