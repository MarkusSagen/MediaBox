

FROM golang:latest

LABEL maintainer="Markus Sagen <Markus.John.Sagen@gmail.com"

# Set working dir in container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.mod ./

# Download depependencies
# Dependencies are csched if go.mod and go.sum are unchanged
RUN go mod download

# Copy source to working dir
COPY . .

# Build go app
RUN go build -o server .

# Export port 
EXPOSE 8000

# Run commands on execution
CMD ["./server"]


