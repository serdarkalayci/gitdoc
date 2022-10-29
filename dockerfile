############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/titanicAPI/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/gitdoc
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/gitdoc /go/bin/gitdoc
ADD static/version.txt /static/version.txt
ADD configuration/livesettings.json /configuration/livesettings.json
ENV BASE_URL :5550
# Run the hello binary.
ENTRYPOINT ["/go/bin/gitdoc"]