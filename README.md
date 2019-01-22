<a href='https://gitee.com/wx85278161/go-iris'><img src='https://gitee.com/wx85278161/go-iris/widgets/widget_1.svg' alt='go iris web'></img></a>
# go iris web实战（响应式web）

## 目前的界面效果
![登录](https://images.gitee.com/uploads/images/2019/0108/173619_918bd02f_1537471.png "屏幕截图.png")
![用户管理](https://images.gitee.com/uploads/images/2019/0122/172900_4bb52b8f_1537471.png "屏幕截图.png")
![角色管理](https://images.gitee.com/uploads/images/2019/0122/172936_6bacbb35_1537471.png "屏幕截图.png")
![菜单管理](https://images.gitee.com/uploads/images/2019/0122/172953_bc31acf3_1537471.png "屏幕截图.png")

### 背景介绍
`Golang + Iris(web框架) + Casbin(权限) + JWT + Vue(渐进式js)`的web server框架，可前后端分离。<br />
Iris的教程较少、零散、基础，且框架集合的完整实战案例极少(毕竟多数是用于工作，商业项目)，几乎没有。后期可以直接使用。<br />
源于开源，馈与社区。<br />
称着还有精力在这方面。
***QQ交流群：955576223***

> #### 软件架构
> 目前支持单web架构，如果部署成前后端分离，可用nginx中间件代理(已添加跨域访问设置)。
>    * 采用了Casbin做Restful的rbac权限控制；
>    * 采用jwt做用户认证、回话控制；
>    * 采用Mysql+xorm做持久层；
>    * Vue前端项目持续更新中...，目前在front-vue分支；

***
#### 项目目录结构
```
go-iris
  +-- a 该目录放的是临时的测试方法
  +-- conf 所有的配置文件目录
  +-- doc 说明文档（含go-bindata和mysql文件）
  +-- exec_packahe 可执行的打包文件（目前只有win 64bit的打包）
  +-- inits 所有需初始化的目录
  |       +-- parse 所有配置文件的初始化目录
  |       +-- init.go 用于初始化系统root用户，并注入所有service
  +-- middleware 包含的中间件目录
  |       +-- casbins 用于rbac权限的中间件的目录
  |       +-- jwts jwt中间件目录
  +-- resources 打包的前端静态资源文件
  +-- utils 工具包目录
  +-- web
  |       +-- db 数据库dao层目录
  |       +-- models 模型文件目录
  |       +-- routes 所有分发出来的路由的目录
  |       +-- supports 提供辅助方法的目录
  +-- main.go 入口
```

### 使用教程
1. 每次修改`/conf/app.yml`或`/conf/db.yml`的配置后，都需要在项目下执行命令打包配置数据：`go-bindata -pkg parse -o inits/parse/conf-data.go conf/`会生成`/inits/parse/conf-data.go`数据文件（执行成功后不会有任何提示，则反之）；
2. **部署时如有上述配置文件修改也需要再执行一遍上述命令，如此才能使配置修改生效**；
3. `go-bindata`的安装和使用教程在项目下的`/doc/go-bindata-usage`文件中说明；
4. 如果不使用前端，可以使用server端根目录下已经打包好的`/resources/*`前端文件；
5. 如果要使用前端：
    * clone或下载`front-vue`分支代码
    * 推荐安装`vue >= 2.x`和`node.js >= v8.9.3(LTS)`环境。IDE推荐安装webstone
    * `npm install`安装本地前端环境
    * `npm run dev`启动本地前端环境
    * `npm run build`打包前端文件
    * 可以将打包的dist目录下的文件拷贝到server端目录的`/resources/`目录下

***
#### 部署（不使用nginx情况下），这里在windows 64bit环境下操作为例。依如下步骤操作：
1. **编译server端项目**。在项目下**使用命令行**执行下面的命令(根据你的需要选择目标OS)：
```
[[编译成当前环境]]
go install
[[编译成Linux 64bit]]
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
go install
[[编译成Mac]]
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=darwin
go install
编译后的可执行文件在你本地go环境的GOPATH/bin/下找到。
```

2. **启动项目**。将server端打包后的可执行文件 和 `/resources/*`前端目录文件 放在同一级目录中，执行go打包后的可执行文件，启动。如下图：
![部署时包的结构](https://images.gitee.com/uploads/images/2019/0108/214456_90a778b1_1537471.png "屏幕截图.png")

> * 启动的本地服务地址：localhost:8088<br/>
> * 超级用户登录：
>    > 初始账号：root<br />
>    > 初始密码: 123456
> * 一般用户登录：
>    > 账号：yhm1<br />
>    > 密码：123456


***
> 安装环境
> * golang >= 1.9
> * nginx 不必须
>    > 如果不使用前端环境，直接使用项目下的`/resource/*`的文件，则可以不需要下面的环境：
>    > * vue >= 2.x
>    > * node.js >= v8.9.3（LTS）

> 待需优化项，如：
> * 前端静态文件数据打包
> * 相同密码没随机加密
> * 同一用户生成的token，生成两次前一次没失效
> * 数据库连接池等等....


#### 参与贡献
1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request