PROGRAMS := $(patsubst %.go,%,$(wildcard *.go))

.PHONY: all clean

all: $(PROGRAMS)

clean:
	rm -f $(PROGRAMS)

%: %.go
	GOPATH=`pwd` go build $<
