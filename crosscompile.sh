#!/bin/bash


env GOOS=darwin GOARCH=386 go build -o streamdl-macos
env GOOS=linux GOARCH=arm go build -o streamdl-linux
env GOOS=windows GOARCH=386 go build -o streamdl-windows