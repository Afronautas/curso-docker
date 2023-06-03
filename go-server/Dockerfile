FROM golang as builder
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o app main.go

FROM scratch
WORKDIR /app/src
ENV PORT 8082
EXPOSE 8082
COPY --from=builder /usr/src/app/app .
ENTRYPOINT [ "./app" ]