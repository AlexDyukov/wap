############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

WORKDIR /usr/share/zoneinfo
RUN apk update && apk --no-cache add tzdata zip
# No compression needed because go's
# tz loader doesn't handle compressed data.
RUN zip -q -r -0 /zoneinfo.zip .

WORKDIR /go/src/wap
COPY . .
# Static build required so that we can safely copy the binary over.
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

############################
# STEP 2 build small image
############################
FROM scratch

# timezone data:
ENV ZONEINFO /zoneinfo.zip
COPY --from=builder /zoneinfo.zip /

# app binary
ENV LOGLEVEL INFO
COPY --from=builder /go/bin/wap /wap
ENTRYPOINT ["/wap"]
