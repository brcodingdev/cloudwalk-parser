## Description
Cloudwalk parser challenge
[Requirements](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1)

## Stack
* Go 1.20
* Cobra `For command line`
* Testify `Unit tests`
* Redis `Store parsed matches`
* make

### Prerequisites
* Go environment
* docker
* make
* docker-compose

## Installation
- Open Makefile, and change the HOST_REDIS to your localhost
  HOST_REDIS = localhost:6379

```bash
# start redis 
$ docker-compose up -d
```

## Parse file
```bash
# parse file with default file ./assets/log_1.log
$ make parse

# if you want parsing other file via command line
# load = action to take
# -f path of file
$ HOST_REDIS=localhost:6379 go run ./cmd/cli load -f ./assets/log_2.log
```

## Report
```bash
# report last file parsed
$ make report
```

## Test
```bash
# unit tests
$ make test
```

## Lint
```bash
# lint the application
$ make lint
```