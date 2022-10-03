$(function () {
    loginApp.init();
})
var loginApp = {
    init: function () {
        this.getCaptcha()
        this.captchaImgChage()
    },
    getCaptcha: function () {
        $.post("/base/captcha", function (response) {
            // $("#captchaId").val(response.captchaId)
            $("#captchaId").val(response.data.captchaId)
            $("#captchaImg").attr("src", response.data.captchaImage)
        })
    },
    captchaImgChage: function () {
        var that = this;
        $("#captchaImg").click(function () {
            that.getCaptcha()
        })
    }
}

