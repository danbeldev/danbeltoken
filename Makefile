danbeltoken:
	make refresh
	docker run --rm -v ${PWD}/contract:/sources ethereum/solc:0.8.0 --abi --bin /sources/danbeltoken.sol -o /sources --overwrite
	docker run --rm -v ${PWD}/contract:/sources ethereum/client-go:alltools-v1.11.2 abigen --abi /sources/DanBelToken.abi --bin /sources/DanBelToken.bin --pkg contract --out /sources/DanBelToken.go
	sudo chmod 777 contract/*