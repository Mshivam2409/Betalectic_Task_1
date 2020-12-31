FROM golang
ENV GO111MODULE=on
RUN mkdir /auth
WORKDIR /auth
EXPOSE 5050 27017
COPY ./ ./
RUN echo "replace betalectic_task_1 => /auth" >> go.mod
RUN go mod download
RUN go build main.go
CMD [ "./main" ]
