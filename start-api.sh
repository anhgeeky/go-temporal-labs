#!/bin/bash
sh -c 'go run ./services/mcs-account/main.go & go run ./services/mcs-money-transfer/main.go & go run ./services/mcs-notification/main.go & wait'