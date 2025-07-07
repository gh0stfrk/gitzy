@echo off
setlocal
REM Change this path if you want a different install location
set TARGET=%USERPROFILE%\bin
if not exist "%TARGET%" mkdir "%TARGET%"
copy gitzy.exe "%TARGET%\gitzy.exe"
setx PATH "%PATH%;%TARGET%"
echo Installed gitzy.exe to %TARGET%
pause
