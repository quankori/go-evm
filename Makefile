events:
	- go run cmd/events/main.go
transfer:
	- go run cmd/transfer/main.go
block:
	- go run cmd/block/main.go
gen-token:
	- abigen --abi internal/abi/token_abi.json --pkg token --out internal/solc/token.go