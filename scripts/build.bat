@ECHO OFF
SETLOCAL
chcp 866>nul

CD ../src/

rm -rf build/

SET BUILD_PATH=..\build
SET APP_PATH=%BUILD_PATH%\app
SET FLAGS="-w -s"

REM download all packages if project require
go mod tidy


ECHO Building release...

REM Linux Build settings
SET GOOS=linux
SET GOARCH=amd64
SET CGO_ENABLED=0


go build -ldflags %FLAGS% -o %APP_PATH%
ECHO Linux app: created successfuly - (%APP_PATH%)

cp .env %BUILD_PATH%\.env