FROM alpine:latest

COPY elix.exe /app/elix.exe

WORKDIR /app

CMD ["./elix.exe"]