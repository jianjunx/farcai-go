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
  var mdTocList = $(".markdown-toc-list");
  // 如果有TOC
  if (mdTocList.length > 0 && mdTocList.children().length > 0) {
    tocBox.append(mdTocList);
    tocScrollTo(tocBox);
  } else {
    $(".detail-left").css("width", "100%");
    $(".detail-right").hide();
  }
}
function tocScrollTo(tocBox) {
  // 组织默认事件
  var all = document.querySelectorAll("#toc-box a");
  for (var i = 0, len = all.length; i < len; i++) {
    all[i].href = "javascript:void(0)";
  }
  var prvEle = null;
  tocBox.on("click", "a", function (event) {
    event.stopPropagation();
    ele = $(event.target);
    ele.addClass("active");
    if (prvEle) prvEle.removeClass("active");
    prvEle = ele;
    var _href = $(event.target).text();
    var top = $("a[name='" + _href + "']").offset().top;
    window.scrollTo(0, top - 80);
  });
}
