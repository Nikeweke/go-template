@echo off
:: Enable command echoing only for debugging, otherwise keep silent
setlocal enabledelayedexpansion

:: Print info
echo Building Go project for Linux...

:: Change to the directory containing the source code
cd "%~dp0..\src"

:: Set environment variables for Linux build (amd64 architecture)
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

:: Optional: clean the Go cache (comment out if not needed)
:: go clean -cache

:: Build the Go binary
:: -ldflags "-s -w" removes debug info and reduces binary size
:: -trimpath removes local file system paths
go build -ldflags="-s -w" -trimpath -o ../build/myprogram-linux-amd64

:: Check if the build was successful
if %ERRORLEVEL% neq 0 (
    echo Build failed!
    exit /b %ERRORLEVEL%
) else (
    echo Build succeeded!
)

:: Optional: Print the size of the compiled binary
echo Binary size:
dir ..\build\myprogram-linux-amd64

pause
