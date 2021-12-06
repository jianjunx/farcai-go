var TOKEN_KEY = "AUTH_TOKEN";
var USER_KEY = "USER_INFO";
var SALT = "_farcai_salt"; // 加盐

$(function () {
  loginLogic();
  pagination();
  tocInit();
  initEditLogic();
  headerActive();
});
function setAjaxToken(xhr) {
  xhr.setRequestHeader("Authorization", localStorage.getItem("AUTH_TOKEN"));
}
function headerActive() {
  var nav = $('a[href="' + location.pathname + '"]');
  if (nav.length == 0) {
    // $('a[href="/"]').addClass("active");
    return;
  }
  nav.addClass("active");
}
function initEditLogic() {
  var edit = $(".detail-edit");
  if (localStorage.getItem(TOKEN_KEY) && edit.length > 0) {
    edit.show();
    var delEle = $(".detail-delete");
    // 绑定删除事件
    delEle.click(function () {
      deleteDetail(delEle.attr("pid"));
    });
  }
}
// 登录部分逻辑
function loginLogic() {
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

    // md5加密
    var MD5Passwd = new Hashes.MD5().hex(passwd + SALT);
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
    if (page == 1) return (location.search = "");
    location.search = "?page=" + page;
  });
  // $(".pagination-btn").click(function (event) {
  //   var val = $(event.target).attr("value");
  //   if (val == 1) return (location.href = "/");
  //   location.search = "?page=" + val;
  // });
}
function deleteDetail(id) {
  var r = confirm("是否确认删除？");
  if (!r) return;
  $.ajax({
    url: "/api/v1/post/" + id,
    type: "DELETE",
    contentType: "application/json",
    success: function (res) {
      if (res.code != 200) alert(res.error);
      location.href = "/";
    },
    beforeSend: setAjaxToken,
  });
}
function tocInit() {
  var tocBox = $("#toc-box");
  if (tocBox.length == 0) return;
  var titles = $(
    "#view-content h1,#view-content h2,#view-content h3,#view-content h4,#view-content h5,#view-content h6"
  );
  var len = titles.length;
  if (len === 0) return;
  for (var i = 0; i < len; i++) {
    var ele = titles[i];
    const item = $(ele).clone();
    item.addClass("toc-label");
    item.attr("id", ele.id + "_sub");
    tocBox.append(item);
  }
  tocScrollTo(tocBox);
}
function tocScrollTo(tocBox) {
  tocBox.click(function (event) {
    $(event.target).addClass("active").siblings().removeClass("active");
    var id = event.target.id.replace("_sub", "");
    var top = document.getElementById(id).offsetTop;
    window.scrollTo(0, top - 80);
  });
}
