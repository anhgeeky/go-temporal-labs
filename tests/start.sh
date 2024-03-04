#!/bin/bash
sh -c 'go run ./temporal/main.go & go run ./micro/main.go & wait'