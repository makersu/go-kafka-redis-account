# go-kafka-redis-account

## start zookeeper, kafka and redis

## build and run producer
```
> go build && ./go-kafka-redis-account
Welcome to Account service: producer

-> create###Mark Su
Message: {Event:{AccId:58b94498-3f2f-4cb0-9469-9e4ccfc0c83c Type:CreateEvent} AccName:Mark Su}
Message is stored in partition 0, offset 0
```

## build and run consumer
```
> go build && ./go-kafka-redis-account

Welcome to Account service: consumer

Press [enter] to exit consumer

Processing CreateEvent:
{"AccId":"58b94498-3f2f-4cb0-9469-9e4ccfc0c83c","Type":"CreateEvent","AccName":"Mark Su"}
Redis: {Id:58b94498-3f2f-4cb0-9469-9e4ccfc0c83c Name:Mark Su Balance:0}
```

## debug kafka
```
> kafka-console-consumer --zookeeper localhost:2181 --topic go-kafka-redis-account  --from-beginning
{"AccId":"58b94498-3f2f-4cb0-9469-9e4ccfc0c83c","Type":"CreateEvent","AccName":"Mark Su"}
```

## debug redis
```
redis-cli
127.0.0.1:6379> hgetall 58b94498-3f2f-4cb0-9469-9e4ccfc0c83c
```