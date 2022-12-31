.PHONY: protos

protos: # execute from building-microservices-with-go/currency
	protoc -I protos/ \
			--go_out=protos/currency \
			--go_opt=paths=source_relative \
			--go-grpc_out=protos/currency \
			--go-grpc_opt=paths=source_relative \
			protos/currency.proto

testget:
	grpcurl --plaintext -d '{"Base": "GBP", "Destination": "USD"}' localhost:9092 protos.Currency.GetRate


testsubscriber:
	grpcurl --plaintext --msg-template -d @ localhost:9092 protos.Currency/SubscribeRates