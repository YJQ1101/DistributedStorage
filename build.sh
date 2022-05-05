for i in `seq 1 6`; do mkdir -p /home/yjq1101/tmp/$i/objects; done
export RABBITMQ_SERVER=amqp://test:test@127.0.0.1:5672
LISTEN_ADDRESS=192.168.0.1:12345 STORAGE_ROOT=/home/yjq1101/tmp/1 go run dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.0.2:12345 STORAGE_ROOT=/home/yjq1101/tmp/2 go run dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.0.3:12345 STORAGE_ROOT=/home/yjq1101/tmp/3 go run dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.0.4:12345 STORAGE_ROOT=/home/yjq1101/tmp/4 go run dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.0.5:12345 STORAGE_ROOT=/home/yjq1101/tmp/5 go run dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.0.6:12345 STORAGE_ROOT=/home/yjq1101/tmp/6 go run dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.1.1:12345 go run apiServer/apiServer.go &
LISTEN_ADDRESS=192.168.1.2:12345 go run apiServer/apiServer.go &
