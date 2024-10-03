@ECHO OFF

:: %~dp0 is a special variable in batch files that expands to the directory of the currently executing script.
CD "%~dp0..\src"

rm -rf build/

SET BUILD_PATH=..\build
SET APP_PATH=%BUILD_PATH%\app
SET FLAGS="-w -s"

:: download all packages if project require
go mod tidy


echo Building Go project for Linux...

:: Linux Build settings
SET GOOS=linux
SET GOARCH=amd64
SET CGO_ENABLED=0

:: Build the Go binary
:: -ldflags "-s -w" removes debug info and reduces binary size
:: -trimpath removes local file system paths
go build -ldflags %FLAGS% -o %APP_PATH%

:: Check if the build was successful
if %ERRORLEVEL% neq 0 (
    echo Build failed!
    exit /b %ERRORLEVEL%
) else (
    echo Build created successfully - (%APP_PATH%)
    cp .env %BUILD_PATH%\.env
)
