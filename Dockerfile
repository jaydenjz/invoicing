# Reference: https://github.com/GoogleCloudPlatform/golang-samples/blob/main/run/helloworld/Dockerfile

# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.18 as builder

# Create and change to the app directory.
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy local code to the container image.
COPY . .

# Build the binary.
RUN cd cmd/invoicing && \
    go build -v -o /bin/invoicing

FROM scratch

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/config /config
COPY --from=builder /bin/invoicing /bin/invoicing

# Run the web service on container startup.
CMD ["/bin/invoicing"]

# docker build --tag docker-invoicing .
# docker run --name invoicing -d docker-invoicing