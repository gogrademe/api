FROM alpine:edge

COPY . /app

RUN apk --no-cache add curl go \
  && cd /app \
  && curl https://glide.sh/get | sh \
  && glide install \
  && go install \
  && apk --no-cache del go
