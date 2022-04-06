events:
	- go run cmd/events/main.go
gen-token:
	- abigen --abi internal/abi/token_abi.json --pkg token --out internal/solc/token.go