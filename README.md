# HushTell后端测试branch
这个branch用于后端的开发与测试，docker中使用的是这个branch

## 功能说明
1. 从前端页面接收表单并存储为匿名文件，目录在`temp`
2. 给每个文件设定动态url：`f/{IPHash}/{filename}`
3. 有两个timer自动删除文件
	1. globalTime，一天后自动销毁文件，无论有没有被访问过, 3秒刷新一次
	2. accessTimer，被访问3秒后自动销毁文件，1秒刷新一次
	3. 如果一个文件夹中没有任何文件，该文件夹也会被销毁

## 如果前后端配合正常
前端上传文件后应当跳转到`/upload`页面，该页面会提示上传成功与否，同时会给出指向文件的动态链接，点击链接会再次跳转，显示文件的查找结果

## TODO
目前的项目结构比较初级，所有的数据处理和路由都在后端，先实现一个这样的版本，跑起来之后再考虑其他的架构

## 文件说明
- `/config`：全局变量或设置的存储位置
- `/model`：项目中用到的类或实体
- `/templates`：html渲染模板
- `/util`：工具函数
- `main.go`：主程序，使用`go run main.go`运行，不足的库会自动补充
- `docker`：dockerfile的存储地点，基本没用

## docker使用说明
首先安装好docker，重启电脑，运行docker之后，在命令行中输入：

`docker run -it -p 8000:8000 joeyliu086/hushtell`

首次使用会涉及到下载，下载可能会有点慢，一次下载之后就不用再次下载了

docker跑起来之后可以通过http://localhost:8000 进入测试页