FROM golang:1.19-alpine AS build

RUN apk --no-cache add tzdata

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/main.go

FROM alpine:latest AS release

WORKDIR /

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /app /app

RUN mkdir /dump

ENV TZ=Asia/Jakarta

ENTRYPOINT ["/app", "--dns", "8.8.8.8"]