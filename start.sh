#!/bin/bash

# Export environment variable
export APP_ENV=local

# Debugging statement to print the value of APP_ENV
echo "APP_ENV is set to: $APP_ENV"

# Run the Go application
go run ./app/main.go
