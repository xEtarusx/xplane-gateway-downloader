#!/bin/bash

########## Windows ##########

# Build .exe for windows
GOOS=windows GOARCH=amd64 go build -o bin/xplane-gateway-downloader.exe .

# Go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# Zip windows application
zip xplane-gateway-downloader-windows.zip xplane-gateway-downloader.exe

# Get back out for next go build cmd
cd ../


########## Mac OS ##########

# Build for mac 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/xplane-gateway-downloader .

# Go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# Zip mac 64-bit application
zip xplane-gateway-downloader-mac_os.zip xplane-gateway-downloader

# Rename file
mv xplane-gateway-downloader xplane-gateway-downloader-amd64-darwin

# Get back out for next go build cmd
cd ../


########## Linux 32-bit ##########

# Build for linux 32-bit
GOOS=linux GOARCH=386 go build -o bin/xplane-gateway-downloader .

# Go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# Zip linux 32-bit application
zip xplane-gateway-downloader-linux-32bit.zip xplane-gateway-downloader

# Rename file
mv xplane-gateway-downloader xplane-gateway-downloader-386-linux

# Get back out for next go build cmd
cd ../


########## Linux 64-bit ##########

# Build for linux 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/xplane-gateway-downloader .

# Go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# Zip linux 64-bit application
zip xplane-gateway-downloader-linux-64bit.zip xplane-gateway-downloader

# Rename file
mv xplane-gateway-downloader xplane-gateway-downloader-amd64-linux

# Get back out for next go build cmd
cd ../
