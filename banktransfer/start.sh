#!/bin/bash
sh -c 'go run ./worker/main.go & go run ./api/main.go & wait'