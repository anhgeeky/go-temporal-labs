#!/bin/bash
sh -c 'go run ./serivces/mcs-account/main.go & go run ./serivces/mcs-money-transfer/main.go & go run ./serivces/mcs-notification/main.go & go run ./serivces/mcs-payment/main.go & wait'