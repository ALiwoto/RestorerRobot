@echo off
TITLE Build latest version
echo.
echo Building from branch:
git branch
echo.
echo Pulling latest version....
git pull
echo.
echo.
echo Building latest binary....
go build
echo.
echo.
timeout 4
exit
:: Hail Hydra
