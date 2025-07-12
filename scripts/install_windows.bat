@echo off
setlocal

REM Change this path if you want a different install location
set TARGET=%USERPROFILE%\bin
if not exist "%TARGET%" mkdir "%TARGET%"
copy gitzy.exe "%TARGET%\gitzy.exe"

set "NEWPATH=%PATH%;%TARGET%"
setx PATH "%NEWPATH%"

echo Installed gitzy.exe to %TARGET%
echo Please restart your terminal or log out and back in to update your PATH.
pause
