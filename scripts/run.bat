@ECHO OFF

:: %~dp0 is a special variable in batch files that expands to the directory of the currently executing script.
CD "%~dp0..\src"

SET BUILD_PATH=.\my-app

:: download all packages if project require
go mod tidy

ECHO Building service...
go build -o %BUILD_PATH%.exe

:: start in dev mode (default)
%BUILD_PATH%.exe
:: start with specified "port" if use flags
@REM %BUILD_PATH%.exe --port 7000
