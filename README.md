# go-order-processing_queue
Order processing system with golang use go-nsq 

# Usage
`
go test --cmd sender --nsqd_address 127.0.0.1
go test --cmd consumer --nsqd_address 127.0.0.1
`

# NSQD
docker-compose with nsqd
change hostname & broadcast address to use
