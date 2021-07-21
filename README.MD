# 说明
> 这是一个用golang写的定制执行任务的工具,crontab的格式支持分钟级别和秒级.
> 
> 5个参数识别为分钟级。6个参数识别为秒级。
> 
> eg:每5秒执行一次 `*/5 * * * * *`

## 配置 config
> 使用了yaml格式的文件格式。
> 
> eg: config.yml
```yaml
crontab:
- cron: '*/10 * * * * *'
  command: echo 'hello world'
- cron: 0 * * * *
  command: echo 'hello !!!'
```

## how to build
> 一键式的编译
- install make、golang、git 安装 make、golang环境、git
- make  执行make命令
- you can find ./dist/crontab 

> 直接编译
- go build -o crontab 得到 crontab。

## how to use manually
> crontab-go -h 查看命令行用法。
> 
> crontab-go -c /home/yourpath/config.yaml 指定自定义配置文件地址
> 
> crontab-go -d 后台运行
> 
> eg： crontab-go -c /a/b/c/cf.yml -d
> 

