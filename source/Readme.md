# Go Kafka Producer 

## Setup
```bash
## Requirements
# Download and setup the kafka broke
https://kafka.apache.org/downloads
```
## Starting a The go kafka producer
```
cd <repo>

# build the project
go build -o source .

# Make sure kafka borker is running
# Run the project
./source
#OR
go run .

# Testing
cd test
go test -v .

# api_test FAILED AND SUCCESS depends on the user requirements

```
