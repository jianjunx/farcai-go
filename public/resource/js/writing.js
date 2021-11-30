var CONTENT_KEY = "CACHE_CONTENT"; // 编辑器内容缓存key
var TITLE_KEY = "CACHE_TITLE"; // 标题缓存key
var AUTO_SAVE_TIME = 5000; // 自动保存时间
var MdEditor = null;
var currentCategory;
function initEditor() {
  MdEditor = editormd("editormd", {
    width: "99.5%",
    height: window.innerHeight - 78,
    syncScrolling: "single",
    path: CNDURL + "/editor-md/lib/",
    placeholder: "",
    appendMarkdown: window.localStorage.getItem(CONTENT_KEY) || "",
  });
}

function initCache() {
  var title = window.localStorage.getItem(TITLE_KEY);
  var headInput = $(".header-input");
  // 取默认标题
  headInput.val(title);
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
}

$(function () {
  // initEditor
  initEditor();
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
