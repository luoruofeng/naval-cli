# 概要
该项目是[Naval](https://github.com/luoruofeng/naval)平台的命令行工具。为[Naval](https://github.com/luoruofeng/naval)提供了*常用命令*和*生成Navel文件*的命令。  
  
其中*生成Navel文件*的命令是Naval API不提供的功能。这也是完成该项目的初衷。   
[Naval](https://github.com/luoruofeng/naval)最终将会提供三种访问方式：API、命令行工具、Dashboard。





# 安装
```shell
go install github.com/luoruofeng/naval-cli@latest
```
<br>



# 查看帮助文档
```shell
naval-cli --help
```

<br>

# 使用案例
* 通过k8s资源创建naval文件`(该工具最常用的功能,Naval API将不会提供该功能)`
```shell
naval-cli create k8s -i="./example/k8s/task1.yml" -o="./example/k8s"  --id=task1
```

<br>

* 通过docker-compose创建naval文件`(该工具最常用的功能,Naval API将不会提供该功能)`
```shell
naval-cli  create compose -i="./example/compose/task1.yml" -o="./example/compose"  --id=task2
```

<br>

* 通过naval文件请求naval来创建任务
```shell
naval-cli add -i="./example/k8s/task1-20230728144138.yml"
```

<br>


* 通过naval文件请求naval来修改任务
```shell
naval-cli update -i="./example/k8s/task1-20230728144138.yml"
```

<br>

* 通过任务id请求naval删除任务
```shell
naval-cli delete -i="task1"  
```

<br>

* 通过任务id请求naval执行任务
```shell
naval-cli exec -i="task1"  
```
<br>