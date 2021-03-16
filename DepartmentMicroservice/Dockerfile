FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./DepartmentMicroservice /app/DepartmentMicroservice
COPY ./ReceptionService /app/ReceptionService
WORKDIR /app/DepartmentMicroservice
RUN go mod download
RUN go build -o departmentserver ./serve && chmod +x ./serve
ENTRYPOINT ["./departmentserver"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/DepartmentMicroservice .
RUN chmod +x ./departmentserver
CMD ["./departmentserver"]
EXPOSE 4001