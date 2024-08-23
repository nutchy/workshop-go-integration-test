# Use the official Go image as the base image
FROM golang:1.22.3

# Set the working directory inside the container
WORKDIR /app

# Install go-junit-report
RUN go install github.com/jstemmer/go-junit-report/v2@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Run tests and generate JUnit report
CMD go test -v ./... 2>&1 | go-junit-report -set-exit-code > test-report.xml
# CMD ["go", "test", "-v" ,"./..."]
