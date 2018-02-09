@echo off

setlocal
echo %~d0:%~p0
set GOPATH=%GOPATH%;%~d0%~p0

go test .
endlocal
pause
