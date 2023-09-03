FROM repo.rdvalidasi.com/golang/docker:1-19-6 AS builder

WORKDIR /opt/Cc

COPY . .

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags "-w -s" -o telegram_robot

FROM repo.rdvalidasi.com/library/alpine:3.5

RUN apk add --no-cache tzdata
ENV TZ Asia/Bangkok

WORKDIR /opt/Cc
RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY --from=builder /opt/Cc/telegram_robot .
CMD ["./telegram_robot"]
