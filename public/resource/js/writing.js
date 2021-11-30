var CONTENT_KEY = "CACHE_CONTENT"; // 编辑器内容缓存key
var TITLE_KEY = "CACHE_TITLE"; // 标题缓存key
var AUTO_SAVE_TIME = 5000; // 自动保存时间
var MdEditor = null;

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
  setInterval(() => {
    window.localStorage.setItem(TITLE_KEY, headInput.val());
    window.localStorage.setItem(CONTENT_KEY, MdEditor.getMarkdown());
  }, AUTO_SAVE_TIME);
}
$(function () {
  // initEditor
  initEditor();
  // 初始化缓存
  initCache();
});
