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

