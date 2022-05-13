from golang:1.18-alpine as builder

run mkdir /app

workdir /app

copy . .

run go build -o main cmd/main.go


from alpine:3.15
run mkdir /app
workdir /app
copy --from=builder /app/main .
copy ./pkg/config/envs ./pkg/config/envs
EXPOSE 5432
cmd ["/app/main"]