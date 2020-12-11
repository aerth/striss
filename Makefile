bin/striss: *.go
	go build -o $@
clean:
	rm -f bin/stiss
