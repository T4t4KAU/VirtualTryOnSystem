docker stop $(docker ps -a | grep venuns | awk '{print $1}')
kill $(lsof -i :7077 -t)
kill $(lsof -i :8088 -t)
kill $(lsof -i :9099 -t)