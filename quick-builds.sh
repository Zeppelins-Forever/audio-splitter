#!/bin/sh
echo "Run this in the same directory as main.go."
mkdir Releases
echo
# Builds
echo "Linux Builds"
GOOS=linux GOARCH=amd64 go build -o Releases/audio-splitter-Linux-x64
GOOS=linux GOARCH=arm64 go build -o Releases/audio-splitter-Linux-ARM64
GOOS=linux GOARCH=riscv64 go build -o Releases/audio-splitter-Linux-RISCV64
echo "Windows Builds"
GOOS=windows GOARCH=amd64 go build -o Releases/audio-splitter-Windows-x64.exe
GOOS=windows GOARCH=arm64 go build -o Releases/audio-splitter-Windows-ARM64.exe
echo "MacOS (Darwin) Builds"
GOOS=darwin GOARCH=amd64 go build -o Releases/audio-splitter-MacOS-x64
GOOS=darwin GOARCH=arm64 go build -o Releases/audio-splitter-MacOS-ARM64
echo "FreeBSD Builds"
GOOS=freebsd GOARCH=amd64 go build -o Releases/audio-splitter-FreeBSD-x64
GOOS=freebsd GOARCH=arm64 go build -o Releases/audio-splitter-FreeBSD-ARM64
echo "NetBSD Builds"
GOOS=netbsd GOARCH=amd64 go build -o Releases/audio-splitter-NetBSD-x64
GOOS=netbsd GOARCH=arm64 go build -o Releases/audio-splitter-NetBSD-ARM64
echo "OpenBSD Builds"
GOOS=openbsd GOARCH=amd64 go build -o Releases/audio-splitter-OpenBSD-x64
GOOS=openbsd GOARCH=arm64 go build -o Releases/audio-splitter-OpenBSD-ARM64
GOOS=openbsd GOARCH=riscv64 go build -o Releases/audio-splitter-OpenBSD-RISCV64
