FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./ReceptionService /app/ReceptionService
COPY ./DepartmentMicroservice /app/DepartmentMicroservice
COPY ./DiseaseService /app/DiseaseService
WORKDIR /app/DiseaseService
RUN go mod download
RUN go build -o diseaseserver ./serve && chmod +x ./serve
ENTRYPOINT ["./diseaseserver"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/DiseaseService .
RUN chmod +x ./diseaseserver
CMD ["./diseaseserver"]
EXPOSE 4003