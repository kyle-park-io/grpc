# build stage
FROM golang:alpine as builder

ARG PROGRAM_VER=dev-docker

COPY . /build
WORKDIR /build

RUN go build -ldflags "-X main.programVer=${PROGRAM_VER}" -o app

# ---
FROM alpine:3.16

COPY --from=builder /build/app /app/app

EXPOSE 50051

WORKDIR /app

ENTRYPOINT ["./app"]
