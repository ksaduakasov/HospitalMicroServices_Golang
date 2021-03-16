FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./ReceptionService /app/ReceptionService
COPY ./DepartmentMicroservice /app/DepartmentMicroservice
WORKDIR /app/ReceptionService
RUN go mod download
RUN go build -o receptionserver ./api && chmod +x ./api
ENTRYPOINT ["./receptionserver"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/ReceptionService .
CMD ["./receptionserver"]
EXPOSE 4000