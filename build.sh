#!/bin/bash

# Build .exe for windows
GOOS=windows GOARCH=amd64 go build -o bin/xplane-gateway-downloader.exe .

# Build for mac 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/xplane-gateway-downloader-darwin-amd64 .

# Build for linux 32-bit
GOOS=linux GOARCH=386 go build -o bin/xplane-gateway-downloader-linux-386 .

# Build for linux 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/xplane-gateway-downloader-linux-amd64 .
