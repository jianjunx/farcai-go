var CONTENT_KEY = "CACHE_CONTENT"; // 编辑器内容缓存key
var TITLE_KEY = "CACHE_TITLE"; // 标题缓存key
var AUTO_SAVE_TIME = 5000; // 自动保存时间
var cos = null;
var MdEditor = null;
var currentCategory;
var headInput = null;
var ArticleItem = {};

function setAjaxToken(xhr) {
  xhr.setRequestHeader("Authorization", localStorage.getItem("AUTH_TOKEN"));
}
function initEditor() {
  // 取默认标题
  headInput.val(ArticleItem.title);
  // 初始化编辑器
  MdEditor = editormd("editormd", {
    width: "99.5%",
    height: window.innerHeight - 78,
    syncScrolling: "single",
    path: CNDURL + "/lib/",
    placeholder: "",
    appendMarkdown: ArticleItem.markdown,
    saveHTMLToTextarea: true,
    tocm: true,
    imageUpload: true,
    emoji: true,
    imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
    // imageUploadURL: "/api/v1/uploadfile",
    imageUploadCalback: function (files, cb) {
      uploadImage(files[0], cb);
    },
  });
}
function uploadImage(file, cb) {
  cos.putObject(
    {
      Bucket: COS_BUCKET /* 必须 */,
      Region: COS_REGION /* 存储桶所在地域，必须字段 */,
      Key: COS_PATH + "/" + Date.now() + "_" + file.name /* 必须 */,
      StorageClass: "STANDARD",
      Body: file, // 上传文件对象
      onProgress: function (progressData) {
        console.log(JSON.stringify(progressData));
      },
    },
    function (err, data) {
      if (!err && data.statusCode == 200) cb("//" + data.Location);
    }
  );
}

function initCOS() {
  cos = new COS({
    getAuthorization: function (options, callback) {
      // 异步获取临时密钥
      $.ajax({
        url: "/api/v1/credentials/cos",
        type: "GET",
        contentType: "application/json",
        success: function (res) {
          if (res.code == !200) return alert(res.error);
          var data = res.data || {};
          var credentials = data.Credentials || {};
          console.log("data", data);
          var params = {
            TmpSecretId: credentials.TmpSecretId,
            TmpSecretKey: credentials.TmpSecretKey,
            SecurityToken: credentials.Token,
            StartTime: data.StartTime, // 时间戳，单位秒，如：1580000000
            ExpiredTime: data.ExpiredTime, // 时间戳，单位秒，如：1580000000
            ScopeLimit: true, // 细粒度控制权限需要设为 true，会限制密钥只在相同请求时重复使用
          };
          callback(params);
        },
        beforeSend: setAjaxToken,
      });
    },
  });
}

function getArticleItem(id) {
  $.ajax({
    url: "/api/v1/article/" + id,
    type: "GET",
    contentType: "application/json",
    success: function (res) {
      if (res.code == !200) return alert(res.error);
      ArticleItem = res.data || {};
      initEditor();
    },
    beforeSend: setAjaxToken,
  });
}

function initCache() {
  headInput = $(".header-input");
  var query = new URLSearchParams(location.search);
  var _id = query.get("id");
  if (_id) return getArticleItem(_id);
  // 取本地缓存
  ArticleItem.title = window.localStorage.getItem(TITLE_KEY);
  ArticleItem.markdown = window.localStorage.getItem(CONTENT_KEY) || "";
  // initEditor
  initEditor();
  // 自动保存
  setInterval(() => {}, AUTO_SAVE_TIME);
}

function saveHandler() {
  window.localStorage.setItem(TITLE_KEY, headInput.val());
  window.localStorage.setItem(CONTENT_KEY, MdEditor.getMarkdown());
}

// 发布
function publishHandler() {
  if (!currentCategory) return $(".publish-tip").text("请选择分类");
  ArticleItem.title = headInput.val();
  ArticleItem.markdown = MdEditor.getMarkdown();
  ArticleItem.content = MdEditor.getHTML();
  ArticleItem.categoryId = currentCategory;

  $.ajax({
    url: "/api/v1/article",
    type: ArticleItem.articleId ? "PUT" : "POST",
    contentType: "application/json",
    data: JSON.stringify(ArticleItem),
    success: function (res) {
      if (res.code == !200) return alert(res.error);
      ArticleItem = res.data || {};
    },
    beforeSend: setAjaxToken,
  });
}

$(function () {
  initCOS();
  // 初始化缓存
  initCache();
  // 保存
  $(".save-btn").click(saveHandler);
  var drop = $(".publish-drop");
  // 显示
  $(".publish-show").click(function () {
    drop.show();
  });
  // 隐藏
  $(".publish-close").click(function () {
    drop.hide();
  });
  // 发布逻辑
  $(".publish-btn").click(publishHandler);
  // 选择分类
  $(".category").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    currentCategory = target.attr("value");
    $(".publish-tip").text("");
  });
});
