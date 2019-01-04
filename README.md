# go iris web

#### 介绍
go+iris+casbin+jwt+vue的web框架，可前后分离。<br />
由于目前框架集合的完整案例极少，几乎没有。<br />
源于开源，馈与社区。<br />
称着还有精力在这方面。
***QQ交流群：955576223***

#### 软件架构
目前支持单web架构，如果部署成前后端分离，可用nginx中间件代理。
* 前端项目后续更新...

待需优化项，如：
1. 相同密码没随机加密
2. 同一用户生成的token，生成两次前一次没失效
3. 数据库连接池等等....

#### 安装教程

##### 安装环境
1. golang >= 1.9
2. nginx 不必须
3. vue >= 2.x
4. node.js >= v8.9.3（LTS）

#### 使用说明
##### HTTP Header <key:value> 设置：
- key-> Authorization<br />
- value-> bearer xxxxxxxx<br />
##### 启动环境
1. `npm run install`安装本地前端环境
2. `npm run dev`启动本地前端环境
3. xxxx

#### *部署*（不使用nginx情况下）

1. 打包app.yml和db.yml配置文件数据。在项目下执行命令：`go-bindata -pkg parse -o conf/parse/conf-data.go conf/`
2. 拷贝配置文件和前端静态文件。再打包`go install`**和可执行文件放在同级**
- 由于casbin不支持数据打包，所以需要拷贝conf文件及该目录下的rbac_model.conf
- 拷贝resources目录及下面的所有文件



#### 参与贡献

1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request


#### 码云特技

1. 使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2. 码云官方博客 [blog.gitee.com](https://blog.gitee.com)
3. 你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解码云上的优秀开源项目
4. [GVP](https://gitee.com/gvp) 全称是码云最有价值开源项目，是码云综合评定出的优秀开源项目
5. 码云官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6. 码云封面人物是一档用来展示码云会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)