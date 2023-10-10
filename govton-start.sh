#!/bin/bash
cd core
go build govton.go
nohup ./govton -f etc/govton-api-1.yaml > /root/govton/logs/govton-1.log 2>&1 &
nohup ./govton -f etc/govton-api-2.yaml > /root/govton/logs/govton-2.log 2>&1 &
nohup ./govton -f etc/govton-api-3.yaml > /root/govton/logs/govton-3.log 2>&1 &