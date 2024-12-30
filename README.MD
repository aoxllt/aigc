# aigc-hunter-api

### 介绍
aigc-hunter 后台

### 软件架构
#### py版:
flask+oss
#### go版:
goframe+mysql+redis

### 安装教程
#### py版（在分支里）：
1. `pip install -r requirements.txt`安装所需的库
2. 修改utility/config.py下的oss数据和检测系统url
3. 启动app.py即可
#### go版：
1. 修改internal/consts/consts.go的检测系统的url
2. 修改manifest/config/config.yaml和hack/config.yaml下的mysql和redis的配置
3. 具体方法 [goframe配置模板](https://goframe.org/docs/web/server-config-file-template)
