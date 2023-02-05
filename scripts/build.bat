@ECHO OFF
SETLOCAL
chcp 866>nul

CD ..

rm -rf build/

SET BUILD_PATH=.\build\my-app

SET GO111MODULE=auto
SET GOMODCACHE=%CD%\packages
SET FLAGS="-w -s"

REM download all packages if project require
go mod tidy


ECHO Building release...

REM Linux Build settings
SET GOOS=linux
SET GOARCH=amd64
SET CGO_ENABLED=0

go build -ldflags %FLAGS% -o %BUILD_PATH%
ECHO Linux app: created successfuly

@REM cp .env.prod release/.env