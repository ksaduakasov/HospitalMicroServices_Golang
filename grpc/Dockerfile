FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./ReceptionService /app/ReceptionService
COPY ./DepartmentMicroservice /app/DepartmentMicroservice
WORKDIR /app/grpc
RUN go mod download
RUN go build -o grpcserver ./server && chmod +x ./server
ENTRYPOINT ["./grpcserver"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/grpc .
CMD ["./grpcserver"]
EXPOSE 4002