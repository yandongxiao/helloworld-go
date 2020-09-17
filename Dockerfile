FROM golang:1.15.2
COPY ./main.go /main.go
CMD ["go", "run", "/main.go"]
