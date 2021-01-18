# API设计开发
接下来按照页面和其对应api进行说明
## 页面一：初始页面
页面的主要元素：
1. nav
2. 两个按钮，让用户选择是输入信息，还是上传文件
3. 版权的标志，注册信息等（暂时不管）
对应的api：
- 无，由前端自行跳转到接下来的对应页面
- 如果用户选择了输入：跳转到`/text`
- 如果用户选择了文件：跳转到`/upload`

## 页面二：/text
页面的主要元素：
1. nav
2. 说明
3. 输入框
4. 提交按钮
5. 选项
	1. 设置允许阅读时间（打开链接后多长时间销毁信息）
	2. 设置保存时间（服务器保存信息的最长时间）
	3. 设置是否加密信息（信息加密）
	4. 设置密码（可以为空）
	5. 设置是否告知对方此信息会自动销毁
	6. 设置是否需要邮件提醒
	7. ~~设置IP限制等等~~

对应的api：
- `/api/upload-text`：
	- 在用户提交时调用
	- api返回对应的动态链接，如果收到true则展示链接给用户，false则提示用户重新尝试

## 页面三：/upload
（注意提示用户可以上传的文件类型、文件大小、文件时长的限制等）
页面的主要元素：
1. nav
2. 说明
3. 输入框
4. 提交按钮
5. 选项
	1. 设置允许阅读时间（打开链接后多长时间销毁信息）
	2. 设置保存时间（服务器保存信息的最长时间）
	3. 设置是否加密信息（信息加密）
	4. 设置密码（可以为空）
	5. 设置是否告知对方此信息会自动销毁
	6. 设置是否需要邮件提醒
	7. ~~设置IP限制等等~~

对应的api：
- `/api/upload-file`：
	- 在用户提交时调用
	- api返回true/false，如果收到true则展示链接给用户，false则提示用户重新尝试

## 页面四：动态文件地址（展示页面）
TODO