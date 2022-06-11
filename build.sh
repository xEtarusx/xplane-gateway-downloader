#!/bin/bash

# Build .exe for windows
GOOS=windows GOARCH=amd64 go build -o bin/xplane-gateway-downloader-windows.exe .

# Build for mac 32-bit
#GOOS=darwin GOARCH=386 go build -o bin/xplane-gateway-downloader-386-darwin .

# Build for mac 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/xplane-gateway-downloader-amd64-darwin .

# Build for linux 32-bit
GOOS=linux GOARCH=386 go build -o bin/xplane-gateway-downloader-386-linux .

# Build for linux 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/xplane-gateway-downloader-amd64-linux .
