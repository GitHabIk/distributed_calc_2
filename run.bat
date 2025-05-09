@echo off
set GOARCH=amd64
set GOOS=windows
set CGO_ENABLED=1

go run ./cmd/calc_service/main.go

pause
