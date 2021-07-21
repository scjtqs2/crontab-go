BUILD=`date +%FT%T%z`
COMMIT_SHA1=`git rev-parse HEAD`
VER='$(call current_branch)'

LDFLAGS=-ldflags " -s -X main.Build=${BUILD} -X main.Version=${Version}.${COMMIT_SHA1}"

build :
	rm -rf dist
	mkdir dist
	CGO_ENABLED=0  go build  ${LDFLAGS} ./dist/crontab .
	chmod -R +x ./dist

clean:
	rm -rf dist

define current_branch
	   local folder="$(pwd)"
        [ -n "$1" ] && folder="$1"
        git -C "$folder" rev-parse --abbrev-ref HEAD | grep -v HEAD || \
        git -C "$folder" describe --tags HEAD || \
        git -C "$folder" rev-parse HEAD
endef
