for i in `seq 1 6`; do mkdir -p /home/yjq1101/tmp/$i/objects; done
export RABBITMQ_SERVER=amqp://test:test@127.0.0.1:5672
LISTEN_ADDRESS=localhost:10001 STORAGE_ROOT=/home/yjq1101/Documents/DistributedStorage/tmp/1 go run dataServer/dataServer.go &
LISTEN_ADDRESS=localhost:10002 STORAGE_ROOT=/home/yjq1101/Documents/DistributedStorage/tmp/2 go run dataServer/dataServer.go &
LISTEN_ADDRESS=localhost:10003 STORAGE_ROOT=/home/yjq1101/Documents/DistributedStorage/tmp/3 go run dataServer/dataServer.go &
LISTEN_ADDRESS=localhost:10004 STORAGE_ROOT=/home/yjq1101/Documents/DistributedStorage/tmp/4 go run dataServer/dataServer.go &
LISTEN_ADDRESS=localhost:10005 STORAGE_ROOT=/home/yjq1101/Documents/DistributedStorage/tmp/5 go run dataServer/dataServer.go &
LISTEN_ADDRESS=localhost:10006 STORAGE_ROOT=/home/yjq1101/Documents/DistributedStorage/tmp/6 go run dataServer/dataServer.go &
LISTEN_ADDRESS=localhost:10007 go run apiServer/apiServer.go &
LISTEN_ADDRESS=localhost:10008 go run apiServer/apiServer.go &
