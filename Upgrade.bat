@echo off
git pull && powershell -command "Stop-service -Force -name "RestorerRobot" -ErrorAction SilentlyContinue; go mod tidy; go build; Start-service -name "RestorerRobot""
:: Hail Hydra