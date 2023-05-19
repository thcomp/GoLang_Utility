@echo off
setlocal enabledelayedexpansion

set processName=Builder.exe
go mod tidy
set GOOS=windows
rem set GOARCH=386
go build -o !processName! main.go
