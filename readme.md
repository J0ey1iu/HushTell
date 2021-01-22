# 项目说明
## 文件夹说明
1. dist 打包后的文件目录
2. node_modules 打包所需的包
3. src 源文件目录
4. package.json 包目录
5. webpack.config.js webpack目录

## 运行项目
1. 安装npm
2. 在该文件夹中cmd 运行npm install以安装所有包
3. cmd中输入 npm start 热启动项目
4. cmd中输入 .\node_modules\.bin\webpack 生成打包文件到dist

## API
1. POST请求头格式：'Content-Type': 'multipart/form-data'
2. 传递数据格式：
    note:
    {
        'mytext', 文本文件
        'options', {
            "emailTip": 布尔值
            "readTip": 布尔值
            "encryption": 布尔值
            "encrytionPwd": 字符串
            "readTime": 字符串数字，单位是hour
            "saveTime": 字符串数字，单位是day
        })
    }

    file：
    {
        'myfile', 文件
        'options', {
            "emailTip": 布尔值
            "readTip": 布尔值
            "encryption": 布尔值
            "encrytionPwd": 字符串
            "readTime": 字符串数字，单位是hour
            "saveTime": 字符串数字，单位是day
        })        
    }
