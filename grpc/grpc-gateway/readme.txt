
https://www.cnblogs.com/sunlong88/p/12128687.html



protoc -I. --go_out=plugins=grpc:. Prod.proto

protoc -I. --go_out=plugins=grpc:. Orders.proto
protoc -I. --go_out=plugins=grpc:. Models.proto



protoc  --grpc-gateway_out=logtostderr=true:. Prod.proto
protoc  --grpc-gateway_out=logtostderr=true:. Orders.proto




go run httpserver.go

go run server.go




curl --location --request POST 'http://localhost:8080/v1/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "order_no": "20200101",
    "user_id": "1001",
    "order_money": 95,
    "order_details": [
        {
            "order_no": "20200101",
            "prod_id": "1012"
        },
        {
            "order_no": "20200101",
            "prod_id": "1013"
        }
    ]
}'
