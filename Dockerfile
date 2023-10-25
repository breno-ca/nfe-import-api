FROM golang:1.19.7 as builder
WORKDIR /desafio-tecnico-backend
COPY . .
COPY env.json .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/desafio

FROM scratch

COPY --from=builder ["/desafio-tecnico-backend/env.json", "./"]
COPY --from=builder /desafio-tecnico-backend/server /server
EXPOSE 8080
CMD ["/server"]