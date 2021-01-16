############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# install git to fetch dependencies and gcc for runtime/cgo
RUN apk update && apk add --no-cache git build-base

WORKDIR /src
COPY . .

# Build
RUN CGO_ENABLED=0 go build -o /bin/wap

############################
# STEP 2 build a small image
############################
FROM scratch

# default env vars
ENV LOGLEVEL INFO

COPY --from=builder /bin/wap /wap

# Run app
ENTRYPOINT ["/wap"]
