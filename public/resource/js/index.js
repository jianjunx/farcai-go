$(function () {
  loginLogic();
  pagination();
});

// 登录部分逻辑
function loginLogic() {
  var TOKEN_KEY = "AUTH_TOKEN";
  var USER_KEY = "USER_INFO";

  if (localStorage.getItem(TOKEN_KEY)) {
    $(".login-action").hide();
    $(".login-end").show();
    var userInfo = JSON.parse(localStorage.getItem(USER_KEY)) || {};
    $(".login-username").text(userInfo.userName);
  }
  // 登录
  $(".login-submint").click(function () {
    var tipEle = $(".login-tip");
    var name = $(".login-name").val();
    var passwd = $(".login-passwd").val();
    if (!name) return tipEle.show().text("请输入用户名");
    if (!passwd) return tipEle.show().text("请输入密码");
    // 加盐
    var salt = "_farcai_salt";
    // md5加密
    var MD5Passwd = new Hashes.MD5().hex(passwd + salt);
    $.ajax({
      url: "/api/v1/login",
      data: JSON.stringify({ username: name, passwd: MD5Passwd }),
      contentType: "application/json",
      type: "POST",
      success: function (res) {
        if (res.code == !200) return tipEle.show().text(res.error);
        var data = res.data || {};
        localStorage.setItem(TOKEN_KEY, data.token);
        localStorage.setItem(USER_KEY, JSON.stringify(data.userInfo));
        location.href = "/";
      },
      error: function (err) {
        console.log("err", err);
        tipEle.show().text("登录错误，请重试");
      },
    });
  });
  // 退出登录
  $(".login-out").click(function () {
    localStorage.removeItem(USER_KEY);
    localStorage.removeItem(TOKEN_KEY);
    $(".login-action").show();
    $(".login-end").hide();
  });
}

// 翻页逻辑
function pagination() {
  var query = new URLSearchParams(location.search);
  var page = query.get("page") || 1;
  $(".pagination-next").click(function () {
    page++;
    location.search = "?page=" + page;
  });
  $(".pagination-prev").click(function () {
    page--;
    if (page == 1) return (location.href = "/");
    location.search = "?page=" + page;
  });
  $(".pagination-btn").click(function (event) {
    var val = $(event.target).attr("value");
    if (val == 1) return (location.href = "/");
    location.search = "?page=" + val;
  });
}
