FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./ReceptionService /app/ReceptionService
COPY ./DepartmentMicroservice /app/DepartmentMicroservice
COPY ./DiseaseService /app/DiseaseService
COPY ./UserService /app/UserService
WORKDIR /app/UserService
RUN go mod download
RUN go build -o userserver ./serve && chmod +x ./serve
ENTRYPOINT ["./userserver"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/UserService .
RUN chmod +x ./userserver
CMD ["./userserver"]
EXPOSE 4004