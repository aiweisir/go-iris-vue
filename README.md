<a href='https://gitee.com/yhm_my/go-iris'><img src='https://gitee.com/yhm_my/go-iris/widgets/widget_1.svg' alt='go iris web'></img></a>
# go iris web（响应式移动端）

## 目前的界面效果
![![输入图片说明](https://images.gitee.com/uploads/images/2019/0108/173445_b6936399_1537471.png "屏幕截图.png")](https://images.gitee.com/uploads/images/2019/0108/173445_f85990af_1537471.png "屏幕截图.png")

![输入图片说明](https://images.gitee.com/uploads/images/2019/0108/173510_e83e8a36_1537471.png "屏幕截图.png")

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
1. `npm install`安装本地前端环境
2. `npm run dev`启动本地前端环境
3. xxxx

#### 部署（不使用nginx情况下）

1. 打包app.yml和db.yml配置文件数据。在项目下执行命令：`go-bindata -pkg parse -o inits/parse/conf-data.go conf/`
2. 拷贝配置文件和前端静态文件。再打包项目`go install`，**和可执行文件放在同级**
- 拷贝resources目录及下面的所有文件



#### 参与贡献

1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request