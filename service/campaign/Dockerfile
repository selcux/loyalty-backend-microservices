FROM golang:alpine as build

ENV PROJECT_PATH="/go/src/gitlab.com/adesso-turkey/loyalty-backend-microservices"

WORKDIR $PROJECT_PATH

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

COPY go.mod .
#COPY go.sum .

RUN go mod download

COPY . $PROJECT_PATH

ENV SERVICE_NAME="campaign"

RUN go build -ldflags="-s -w" -o /app ./service/$SERVICE_NAME/cmd
#RUN go build -o app ./cmd/$SERVICE_NAME

#WORKDIR /dist

#RUN mv $PROJECT_PATH/app /

#RUN ls -all

#RUN echo $PROJECT_PATH

#ENTRYPOINT ["/dist/app"]

FROM scratch
COPY --from=build /app /app
ENTRYPOINT ["/app"]