docker stop $(docker ps -a | grep venuns | awk '{print $1}') > /dev/null
kill $(lsof -i :7077 -t)
kill $(lsof -i :8088 -t)
kill $(lsof -i :9099 -t)

cd /root/govton/core
go build govton.go
nohup ./govton -f etc/govton-api-1.yaml > /root/govton/logs/govton-1.log 2>&1 &
nohup ./govton -f etc/govton-api-2.yaml > /root/govton/logs/govton-2.log 2>&1 &
nohup ./govton -f etc/govton-api-3.yaml > /root/govton/logs/govton-3.log 2>&1 &