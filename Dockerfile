FROM pangpanglabs/golang:jan AS builder
WORKDIR /go/src/fruit-api/
COPY ./ /go/src/fruit-api/
# disable cgo 
ENV CGO_ENABLED=0
# build steps
RUN echo ">>> 1: go version" && go version \
    && echo ">>> 2: go get" && go-wrapper download \
    && echo ">>> 3: go install" && go-wrapper install

# make application docker image use alpine
FROM  alpine:3.6
RUN apk --no-cache add ca-certificates
WORKDIR /go/bin/
# copy config file to image (like config.json or config.staging.json)
#COPY --from=builder /go/src/fruit-api/config*.yml ./
# copy execute file to image
COPY --from=builder /go/bin/ ./
EXPOSE 5000
CMD ["./fruit-api"]