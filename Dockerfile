FROM golang:1.20
LABEL authors="Aliothmoon"
WORKDIR /app
COPY . /app
# Build all services
ENV GOPROXY=https://goproxy.cn TUZI_CFG=/app/shared/config TUZI_SEC=/app/internal/service/auth/config
RUN ["go","work","sync"]
#ENTRYPOINT ["sh", "build-all.sh"]


#Bak
FROM golang:1.20 as bin
LABEL authors="Aliothmoon"
WORKDIR /app
COPY . /app
# Build all services
ENV GOPROXY=https://goproxy.cn TUZI_CFG=/app/shared/config TUZI_SEC=/app/internal/service/auth/config
# Add Layer To Cache
RUN ["sh", "/app/internal/service/auth/build.sh"]

FROM ubuntu:22.10
WORKDIR /app
COPY --from=bin  /app/output/bin/auth-api ./
COPY --from=bin /app/shared/config/*.yaml ./
COPY --from=bin /app/internal/service/auth/config/secret.yaml ./
CMD  ["./auth-api"]