<a href='https://gitee.com/yhm_my/go-iris'><img src='https://gitee.com/yhm_my/go-iris/widgets/widget_1.svg' alt='go iris web'></img></a>
# go iris web实战（响应式web）

## 目前的界面效果
![登录](https://images.gitee.com/uploads/images/2019/0108/173619_918bd02f_1537471.png "屏幕截图.png")

![首页1](https://images.gitee.com/uploads/images/2019/0108/173654_4cfd4836_1537471.png "屏幕截图.png")

![首页2](https://images.gitee.com/uploads/images/2019/0108/173718_83b02d34_1537471.png "屏幕截图.png")

### 背景介绍
`Golang+Iris（web框架）+Casbin（权限）+JWT+Vue（渐进式js）`的web server框架，可前后端分离。<br />
Iris的教程比较少，比较零散基础，且框架集合的完整实战案例极少，几乎没有。<br />
源于开源，馈与社区。<br />
称着还有精力在这方面。
***QQ交流群：955576223***

> #### 软件架构
>    * 目前支持单web架构，如果部署成前后端分离，可用nginx中间件代理。
>    * 前端项目持续续更新中...，目前在front-vue分支

### 使用教程
1. 每次修改`/conf/app.yml`或`/conf/db.yml`的配置数据后，都需要在项目下执行命令：`go-bindata -pkg parse -o inits/parse/conf-data.go conf/`会生成`/inits/parse/conf-data.go`数据文件（执行成功后没任何提示）；
2. 部署时如有变动也需要再执行一遍上述命令，如此才能使配置修改生效；
3. `go-bindata`的安装和使用教程在项目下的`/doc/go-bindata-usage`文件中说明；
4. 除了首页、登录、注册接口其他都需要在Header中设置token信息；
5. HTTP Header <key:value> 设置：
    * key   -> Authorization
    * value -> bearer xxx
6. 如果不使用前端，可以使用server根目录下已经打包好的`/resources/*`前端文件；
7. 如果要使用前端：
    * clone或下载`front-vue`分支代码
    * 推荐安装`vue >= 2.x`和`node.js >= v8.9.3(LTS)`环境。IDE推荐安装webstone
    * `npm install`安装本地前端环境
    * `npm run dev`启动本地前端环境
    * `npm run build`打包前端文件
    * 可以将打包的dist目录下的文件拷贝到server目录的`/resources/`目录下

#### 部署（不使用nginx情况下），这里以windows 64bit环境下为例。步骤如下：
1. 编译server端项目。在项目下**使用命令行**执行下面的命令(根据你的需要选择)：
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
```
2. ***启动项目：***将打包后的文件 和 `/resources/*`目录文件 放在同一级目录中，执行go打包后的可执行文件，启动。如下图：
![部署时包的结构](https://images.gitee.com/uploads/images/2019/0108/214456_90a778b1_1537471.png "屏幕截图.png")

### 安装环境
1. golang >= 1.9
2. nginx 不必须
如果不使用前端环境，直接使用项目下的`/resource/*的文件`，则可以不需要下面的环境：
3. vue >= 2.x
4. node.js >= v8.9.3（LTS）

> 待需优化项，如：
>    > 相同密码没随机加密
>    > 同一用户生成的token，生成两次前一次没失效
>    >数据库连接池等等....


#### 参与贡献
1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request