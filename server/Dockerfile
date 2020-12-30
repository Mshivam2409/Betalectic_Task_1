FROM golang
ENV GO111MODULE=on
RUN mkdir /server
WORKDIR /server
EXPOSE 5050 27017
COPY ./ ./
RUN echo "replace betalectic_task_1 => /server" >> go.mod
RUN go mod download
RUN go build main.go
CMD [ "./main" ]
