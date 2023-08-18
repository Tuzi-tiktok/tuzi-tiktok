FROM ubuntu:22.10
LABEL authors="Aliothmoon"
WORKDIR /app
COPY bin .

ENV TUZI_RFG="TUZI_RFG=100.93.41.99:8848?n=c8bf818f-7660-4590-b763-157cffcd209b&d=bootstrap.yml"

ENTRYPOINT ["top","-b"]