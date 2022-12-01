#!/bin/bash
echo === EXAMPLE ===
go run $1.go $1.example.txt
echo
echo
echo === INPUT ===
go run $1.go $1.txt
