#!/bin/bash

########## Windows ##########

# Build .exe for windows
GOOS=windows GOARCH=amd64 go build -o bin/xplane-gateway-downloader.exe .

# go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# zip windows application
tar -cvf xplane-gateway-downloader-windows.zip xplane-gateway-downloader.exe

# get back out for next go build cmd
cd ../


########## Mac OS ##########

# Build for mac 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/xplane-gateway-downloader .

# go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# zip mac 64-bit application
tar -cvf xplane-gateway-downloader-mac_os.zip xplane-gateway-downloader

# rename file
mv xplane-gateway-downloader xplane-gateway-downloader-amd64-darwin

# get back out for next go build cmd
cd ../


########## Linux 32-bit ##########

# Build for linux 32-bit
GOOS=linux GOARCH=386 go build -o bin/xplane-gateway-downloader .

# go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# zip mac 64-bit application
tar -cvf xplane-gateway-downloader-linux-32bit.zip xplane-gateway-downloader

# rename file
mv xplane-gateway-downloader xplane-gateway-downloader-386-linux

# get back out for next go build cmd
cd ../


########## Linux 64-bit ##########

# Build for linux 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/xplane-gateway-downloader .

# go into bin folder to avoid bin folder appearing in .zip file
cd bin/

# zip mac 64-bit application
tar -cvf xplane-gateway-downloader-linux-64bit.zip xplane-gateway-downloader

# rename file
mv xplane-gateway-downloader xplane-gateway-downloader-amd64-linux

# get back out for next go build cmd
cd ../
