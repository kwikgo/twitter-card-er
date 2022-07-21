
BIN := ./twitter-card-er

all: $(BIN)

$(BIN): *.go
	go build

run: all
	env TCE_BASEURL=http://localhost:9911 $(BIN)

install:
	go install

clean:
	rm -f $(BIN)

