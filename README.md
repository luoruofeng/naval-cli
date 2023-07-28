

# 案例
* 通过k8s资源创建naval文件
```
naval-cli create k8s -i="./example/k8s/task1.yml" -o="./example/k8s" --id=task1
```

* 通过docker-compose创建naval文件
```
go run . create compose -i="./example/compose/task1.yml" -o="./example/compose"  --id=task2
```

