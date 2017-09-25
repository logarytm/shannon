all: shannon

%: %.go
	go build $<
