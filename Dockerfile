FROM golang:1.20
LABEL authors="Aliothmoon"
WORKDIR /app
COPY . /app
# Build all services
ENV GOPROXY=https://goproxy.cn TUZI_CFG=/app/shared/config TUZI_SEC=/app/internal/service/auth/config
RUN ["go","work","sync"]
#ENTRYPOINT ["sh", "build-all.sh"]
CMD ["top","-b"]