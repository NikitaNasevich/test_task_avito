FROM golang:alpine as builder

LABEL maintainer="Mikita Nasevich"

RUN apk update && apk add --no-cache git

RUN mkdir /avito_0
WORKDIR /avito_0
COPY . .

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /avito_0/main .
COPY --from=builder /avito_0/.env .

#ENV PORT=3000
#ENV DATABASE=bestuser:bestuser@(127.0.0.1:3308)/test_avito

EXPOSE 3000

ENTRYPOINT ./main