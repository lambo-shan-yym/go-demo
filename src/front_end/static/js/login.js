$(function () {
    $("#loginBtn").click(function () {
        $("#userInfoForm").validate({
            rules: {
                username: {
                    required: true,
                    minlength: 5
                },
                password: {
                    required: true,
                    minlength: 6,
                    maxlength: 20
                },

            },
            errorElement: "em",
            messages: {
                username: {
                    required: "请输入用户名",
                    minlength: "用户名不能小于5位"

                },
                password: {
                    required: "请输入密码",
                    minlength: "密码长度不能小于6位",
                    maxlength: "密码长度不能大于20位"
                },
            },
            submitHandler: function (form) {
                username = $("#username").val()
                password = $("#password").val()
                // 25d55ad283aa400af464c76d713c07ad
                data1 = {username: username, password: $.md5(password)};
                $.ajax({
                    async: true,
                    type: 'post',
                    url: "/login",
                    contentType: 'application/json',
                    data: JSON.stringify(data1),
                    success: function (data) {
                        //后台用JsonResponse返回数据
                        //data 就会被转成字典
                        console.log(data)
                        //JSON.parse(data) 把字符串类型转成字典
                        if (data.code == 0) {
                            //成功,往localStorage中存储toekn，然后跳转到指定页面
                            localStorage.setItem('Authorization', data.result.token);
                            //alert(localStorage.getItem("Authorization"))
                            window.location.href = "/user_info_page";
                        } else {
                            myAlert('系统提示','用户名或者密码错误！',function(){
                                //要回调的方法
                                //window.location.href="http://www.baidu.com"
                            });

                            // layer.alert('用户名或密码错误', {
                            //     skin: 'layui-layer-lan'
                            //     , closeBtn: 0
                            //     , anim: 4 //动画类型
                            // });
                        }
                    },
                    error: function (data) {
                        layer.alert('系统异常，请稍后再试', {
                            skin: 'layui-layer-lan'
                            , closeBtn: 0
                            , anim: 4 //动画类型
                        });
                    }
                });
            }
        });

    })
})