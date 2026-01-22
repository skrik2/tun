.PHONY: test clean

test:
	- go run ./tuncmd/cmd/main.go
	
clean:
	- rm -rf build/*