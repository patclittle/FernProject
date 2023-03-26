#!/bin/bash

# Build and run the Go program
go build -o main && ./main &

# Print a message to the console
echo "Go to 'http://localhost:8080' to try it out"

# Wait for the program to finish
wait