#!/bin/bash

mkdir -p cmd
mkdir -p internal/config
mkdir -p internal/handler
mkdir -p internal/model
mkdir -p internal/repository
mkdir -p internal/service
mkdir -p pkg

touch cmd/main.go

echo "module myapps" > go.mod
echo "Project MyApps initialized."
