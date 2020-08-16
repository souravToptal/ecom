# Build the Go Binary.

FROM golang:1.12.1 as build
ENV CGO_ENABLED 0
ARG VCS_REF
ARG PACKAGE_NAME
ARG PACKAGE_PREFIX
RUN mkdir -p /go/src/github.com/souravToptal/ecom
COPY . /go/src/github.com/souravToptal/ecom
WORKDIR /go/src/github.com/souravToptal/ecom/cmd/${PACKAGE_PREFIX}${PACKAGE_NAME}
RUN go build -ldflags "-s -w -X main.build=${VCS_REF}" -a -tags netgo

# Run the Go Binary in Alpine.

FROM alpine:3.7
RUN apk add --no-cache ca-certificates openssl
ARG BUILD_DATE
ARG VCS_REF
ARG PACKAGE_NAME
ARG PACKAGE_PREFIX
COPY --from=build /go/src/github.com/souravToptal/ecom/cmd/${PACKAGE_PREFIX}${PACKAGE_NAME}/${PACKAGE_NAME} /app/main
WORKDIR /app
CMD /app/main
