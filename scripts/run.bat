@ECHO OFF
SETLOCAL
chcp 866>nul
CLS

SET BUILD_PATH=.\my-app

SET GO111MODULE=auto
SET GOMODCACHE=%CD%\packages

REM download all packages if project require
go mod tidy

ECHO Building service...

rem build with flags
go build -o %BUILD_PATH%.exe

rem start in dev mode (default)
%BUILD_PATH%.exe

rem start with flags - port, create tax events
@REM %BUILD_PATH%.exe -port 7000 -create-tax-events