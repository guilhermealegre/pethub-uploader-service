#FROM 176300518568.dkr.ecr.eu-west-1.amazonaws.com/baseimages:golang1.20.4-alpine3.17 as builder
FROM 176300518568.dkr.ecr.eu-west-1.amazonaws.com/baseimages:golang1.22.2-alpine3.19 as builder

COPY . /go/src/be-uploader-service
WORKDIR /go/src/be-uploader-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o be_uploader cmd/server/main.go

RUN chown root:root be_uploader
RUN chown 755 be_uploader

#FROM 176300518568.dkr.ecr.eu-west-1.amazonaws.com/baseimages:alpine3.14
FROM 176300518568.dkr.ecr.eu-west-1.amazonaws.com/baseimages:alpine3.19.1

COPY --from=builder --chown=root:root /go/src/be-uploader-service/be_uploader .

RUN apk --no-cache add ca-certificates \
    curl \
    bash

RUN apk add aws-cli

COPY internal/swagger/docs/v1 internal/swagger/docs/v1
COPY build/bash-multi.entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh


ENTRYPOINT ["/entrypoint.sh"]
CMD ["./be_uploader"]
