FROM golang:alpine as builder
LABEL stage=builder
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 go build -o /bin/ascii-art-web-dockerize

FROM scratch
LABEL name="ascii-art-web-dockerize"
LABEL description="Ascii Art Web Dockerize"
COPY --from=builder /bin/ascii-art-web-dockerize /
COPY --from=builder /src/font /font
COPY --from=builder /src/front /front
EXPOSE 8080
CMD ["/ascii-art-web-dockerize"]