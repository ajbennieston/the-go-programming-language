PROGRAMS := $(patsubst %.go,%,$(wildcard *.go))

.PHONY: all clean

all: $(PROGRAMS) 

clean:
	rm -f $(PROGRAMS)

%: %.go ../src/golang.org/x/net/html
	GOPATH=`pwd`/.. go build $<

../src/golang.org/x/net/html:
	-GOPATH=`pwd`/.. go get
