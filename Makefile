BUILD=`date +%FT%T%z`
COMMIT_SHA1=`git rev-parse HEAD`
VER=`sh git_branch.sh`

LDFLAGS=-ldflags " -s -X main.Build=${BUILD} -X main.Version=${VER}.${COMMIT_SHA1}"

build :
	rm -rf dist
	mkdir dist
	CGO_ENABLED=0  go build  ${LDFLAGS} -o ./dist/crontab .
	chmod -R +x ./dist

clean:
	rm -rf dist

