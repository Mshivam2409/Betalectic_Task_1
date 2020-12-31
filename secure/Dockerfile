FROM golang
ENV GO111MODULE=on
RUN mkdir /secure
WORKDIR /secure
EXPOSE 5080 27017
COPY ./ ./
RUN echo "replace betalectic_task_1_secure => /secure" >> go.mod
RUN go mod download
RUN go build main.go
CMD [ "./main" ]
