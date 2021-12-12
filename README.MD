# 简介

用 Go 写的博客系统练手项目，基于 goframe 框架和 Mariadb,功能简单够用。

# 快速开始

### 修改配置

将`/config/config.toml`中的配置修改成自己的

### 配置环境变量

Markdown 编辑器的图片上传功能使用了腾讯云 COS，所以启动是在本地配置下 COS 的 `SECRETID` 和 `SECRETKEY` 环境变量

```bash
$ export COS_SECRETID="xxx"
$ export COS_SECRETKEY="xxx"
```

### 运行

```bash
$ go mod tidy
$ go run main.go
```

#### 热更新

安装[Air](https://github.com/cosmtrek/air)然后启动预览

```bash
$ make run
```

# Docker 部署

执行以下命令构建镜像

```bash
$ make image
```

构建完镜像后
修改项目根目录下的`local.env`配置

```bash
# 运行镜像 加载本地.env环境变量文件 映射3000端口 将logs文件夹与本地home/log映射
docker run --env-file env/blog.env -p 3000:3100 --name blog-v1.0.0 -d -v /home/logs/blog:/tmp/logs farcai:镜像名和版本号
```

# 功能说明

项目首次部署后需要对部分数据进行初始化，可通过路由查看对应的接口，再通过 Postman 进行添加。

### 注册

地址：/api/v1/signup

```json
{
  "name": "username",
  "passwd": "md5加密后的数据",
  "repasswd": "同上"
}
```

目前没有提供注册页面，账号需要自己从数据库添加，注意：密码在前端登录时进行了 MD5 加密，所以在数据库中添加密码也必须是 MD5 加密后的 [MD5](https://www.cmd5.com/hash.aspx?s=123456) 在这里加密一下，其中：

```
Pass # 明文密码
Salt # 加盐 默认是：_farcai_salt，也可以在index.js里修改成自己的
```

注册完在页面登录一下

### 添加分类

地址：/api/v1/category POST
请求头：Authorization=打开浏览器调试从页面 localstore 里面取

```json
{
  "name": "Golang"
}
```

### 添加自定也页面

添加自定义页面与添加文章相同，在发布弹窗下选择自定义页，并填写自定义链接。
访问时地址就是`host/自定义链接`，网页标题就是添加的文章标题。

### TOC

要开启 TOC 需要在文章底部添加`[TOC]`，这是会自动生成 TOC

### 评论

在config.toml文件中开启```Valine = true```并配置相应的id和key，参考[Valine](https://valine.js.org/quickstart.html)

## License
[MulanPSL-2.0](https://github.com/jianjunx/farcai-go/blob/master/LICENSE)