#!/bin/bash
sh -c 'go run ./workflows/banktransfer/cmd/worker/main.go & go run ./services/mcs-account/main.go & go run ./services/mcs-money-transfer/main.go & wait'