

# 案例
* 通过k8s资源创建naval文件
```
naval-cli create k8s -i="./example/k8s/task1.yml" -o="./example/k8s"  --id=task1
```

* 通过docker-compose创建naval文件
```
naval-cli  create compose -i="./example/compose/task1.yml" -o="./example/compose"  --id=task2
```

* 通过naval文件请求naval来创建任务
```
naval-cli add -i="./example/k8s/task1-20230728144138.yml"
```

* 通过naval文件请求naval来修改任务
```
naval-cli update -i="./example/k8s/task1-20230728144138.yml"
```

* 通过任务id请求naval删除任务
```
naval-cli delete -i="task1"  
```
