$(function () {
    $.ajax({
        async: true,
        type: 'get',
        url: "/user/get_user_info",
        contentType: 'application/json',
        headers: {'Accept': 'application/json', 'Authorization': localStorage.getItem("Authorization")},
        success: function (data) {
            if (data.code == 0) {
                //成功,往localStorage中存储toekn，然后跳转到指定页面
                $("#username").val(data.result.username)
                $("#nickname").val(data.result.nickname)
                $("#img0").attr("src", data.result["profile_picture"])
            } else {
                isTokenExpired(data)
            }
        },
        error: function (data) {

        }
    });


    $("#pic").change(function () {
        if ($.support.msie) {
            $("#img0").attr("src", $(this).val())
        } else {
            var objUrl = getObjectURL(this.files[0]);
            console.log("objUrl=" + objUrl);
            if (objUrl) {
                $("#img0").attr("src", objUrl);
            }
        }
    })

    //建立一個可存取到該file的url
    function getObjectURL(file) {
        var url = null;
        if (window.createObjectURL != undefined) {
            url = window.createObjectURL(file);
        } else if (window.URL != undefined) {
            url = window.URL.createObjectURL(file);
        } else if (window.webkitURL != undefined) {
            url = window.webkitURL.createObjectURL(file);
        }
        return url;
    }

    $("#userInfoEditBtn").click(function () {
        $('#nickname').removeAttr("disabled");
        $('#userInfoEditSubmitBtn').removeAttr("disabled");
    });
    $("#userInfoEditSubmitBtn").click(function () {
        $('#nickname').attr("disabled", true);
        $('#userInfoEditSubmitBtn').attr("disabled", true);
    });

    $("#logoutBtn").click(function () {

        myConfirm('系统确认框', '您确定要退出登录吗?', function (r) {
            if (r) {
                //...点确定之后执行的内容
                $.ajax({
                    async: true,
                    type: 'post',
                    url: "/user/logout",
                    contentType: 'application/json',
                    headers: {'Accept': 'application/json', 'Authorization': localStorage.getItem("Authorization")},
                    success: function (data) {
                        if (data.code == 0) {
                            localStorage.removeItem('Authorization');
                            window.location.href = "/index";
                        } else {

                        }
                    },
                    error: function (data) {
                    }
                });
            }
        });
    })

    $("#userInfoEditSubmitBtn").click(function () {
        var nickNameTextVal = $("#nickname").val();
        if (nickNameTextVal == null || nickNameTextVal == "") {
            alert("昵称不能为空")
        }
        data1 = {nickname: nickNameTextVal}
        $.ajax({
            async: true,
            type: 'put',
            url: "/user/update_user_info",
            contentType: 'application/json',
            headers: {'Accept': 'application/json', 'Authorization': localStorage.getItem("Authorization")},
            data: JSON.stringify(data1),
            success: function (data) {

                if (data.code == 0) {
                    myAlert('系统提示', '个人信息修改成功', function () {
                        //要回调的方法
                        location.reload()
                    });
                } else {

                }
            },
            error: function (data) {
            }
        });
    })

    $("#profileEditBtn").click(function () {
        $('#pic').click();
        $('#profileUpdateBtn').attr("disabled", false);

    })


    $("#profileUpdateBtn").click(function () {
        var fileInput = $('#pic').get(0).files[0]
        if (!fileInput) {
            return
        }
        var formData = new FormData();
        formData.append('file', $('#pic')[0].files[0]);  //添加图片信息的参数
        $.ajax({
            url: "/user/update_user_profile_picture",
            type: "put",
            dataType: "json",
            cache: false,
            data: formData,
            headers: {'Authorization': localStorage.getItem("Authorization")},
            processData: false,// 不处理数据
            contentType: false, // 不设置内容类型
            success: function (data) {
                if (data.code == 0) {
                    myAlert('系统提示', '个人头像修改成功', function () {
                        //要回调的方法
                        location.reload()
                    });
                } else {

                }
            }
        })
    })
})