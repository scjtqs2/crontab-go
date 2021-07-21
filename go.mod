module github.com/scjtqs/crontab-go

go 1.13

require (
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/robfig/cron/v3 v3.0.1
	github.com/scjtqs/utils v0.0.0-20210721044355-610ada9b6ca7
	github.com/sirupsen/logrus v1.8.1
	github.com/t-tomalak/logrus-easy-formatter v0.0.0-20190827215021-c074f06c5816
	go.uber.org/dig v1.11.0
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

//replace github.com/scjtqs/utils => /Users/apple/Workspace/git/utils-github
