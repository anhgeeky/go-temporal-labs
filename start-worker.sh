#!/bin/bash
sh -c 'go run ./pkg/workflows/banktransfer/cmd/worker/main.go & wait'