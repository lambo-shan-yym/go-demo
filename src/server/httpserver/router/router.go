package router

import (
	"entry_task/src/server/httpserver/api"
	"entry_task/src/server/httpserver/code"
	"entry_task/src/tools"
	log "github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.NoRoute(NoRouterHandler)
	r.NoMethod(NoMethodHandler)
	r.Use(ErrorHandler())
	r.LoadHTMLGlob(
		"src/front_end/static/*.html") // html模板
	r.GET("/index", api.LoginHtml)
	r.GET("/user_info_page", api.UserInfoPageHtml)
	r.POST("/login", api.Login)

	authorized := r.Group("/user", JWTAuth())
	authorized.POST("/logout", api.Logout)
	authorized.GET("/get_user_info", api.GetUserInfo)
	authorized.PUT("/update_user_info", api.UpdateUserInfo)
	authorized.PUT("/update_user_profile_picture", api.UpdateUserProfilePicture)

	r.Static("/front_end/upload", "./src/front_end/upload/")
	r.Static("/src/front_end/upload", "./src/front_end/upload/")
	r.Static("/front_end/static", "./src/front_end/static/")
	return r
}

func NoRouterHandler(ctx *gin.Context) {
	e := code.NotFound
	ctx.JSON(e.StatusCode, e)
	return
}

func NoMethodHandler(ctx *gin.Context) {
	e := code.NoMethod
	ctx.JSON(e.StatusCode, e)
	return
}

// 统一异常处理
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var exception *code.Exception
				if e, ok := err.(*code.Exception); ok {
					exception = e
				} else if e, ok := err.(error); ok {
					exception = code.OtherException(e.Error())
				} else {
					exception = code.ServerException
				}
				ctx.JSON(exception.StatusCode, exception)
				return
			}
		}()
		ctx.Next()
	}
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Abort()
			panic(code.Unauthorized)
			return
		}

		//log.Print("get token: ", token)

		j := tools.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			log.Errorf("【http服务】解析token发生异常，异常信息：%s", err.Error())
			if err == tools.TokenExpired {
				c.Abort()
				panic(code.TokenExpired)
				return
			}
			panic(code.FormatException(code.SystemException, err.Error()))
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Next()
	}
}

// 测试使用的token，token为username
func JWTAuthForTestEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Abort()
			panic(code.Unauthorized)
			return
		}
		claims := &tools.CustomClaims{
			Username:       token,
			StandardClaims: jwt.StandardClaims{},
		}
		c.Set("claims", claims)
	}
}
