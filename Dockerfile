FROM golang:1.21-alpine as builder
WORKDIR /

COPY . ./
RUN go mod download


RUN go build -o /service-barebones

FROM alpine
COPY --from=builder /service-barebones .

EXPOSE 80
CMD [ "/service-barebones" ]`