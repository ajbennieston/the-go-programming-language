PROGRAMS := $(patsubst %.go,%,$(wildcard *.go))

.PHONY: all clean

all: $(PROGRAMS)

clean:
	rm -f $(PROGRAMS)

%: %.go
	go build $^

test-tempconv: test-tempconv.go src/gopl.io/ch2/tempconv/*.go
	GOPATH=`pwd` go build $<
