#!/bin/bash
choco install zip -y
go build .
rm wordlists/*
zip -r brutus-hash-hunter-0.9.1.zip config wordlists brutus-hash-hunter.exe 