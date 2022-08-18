FROM amd64/golang:1.18 as builder

# version of the image
# must not include a leading v(1.2.3). example: 1.2.3
ARG VERSION

WORKDIR /go/src/terraform-providers/terraform-provider-ikea
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download
COPY ./ ./

RUN env GOOS=linux GOARCH=amd64 go build -o terraform-provider-ikea_v$VERSION .

FROM hashicorp/terraform:1.2.2

# see above
ARG VERSION

# install requirements for pipeline execution
RUN apk add --update git bash curl unzip zip openssl make gcc musl-dev linux-headers libffi-dev libc6-compat coreutils jq

# Create a group and user
RUN addgroup -S terraform && adduser -S terraform -G terraform
USER terraform

# copy binary to home directory
COPY --from=builder /go/src/terraform-providers/terraform-provider-ikea/terraform-provider-ikea_v$VERSION /home/terraform/.terraform.d/plugins/qaware.com/terraform/ikea/$VERSION/linux_amd64/terraform-provider-ikea_v$VERSION

ENTRYPOINT ["terraform"]
