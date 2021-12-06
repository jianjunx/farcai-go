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
# 运行镜像 加载本地local.env环境变量 映射3200端口 将logs文件夹与本地home/log映射
docker run --env-file ./local.env -p 3200:3100 --name farcai-server -d -v /home/logs:/logs farcai:镜像名和版本号
```
