# Tahap build
FROM golang:1.23-alpine AS build

# Direktori kerja
WORKDIR /goapp

# Copy semua file ke dalam image build
COPY . .

# Unduh dependencies dan lakukan build
RUN go mod download
RUN go build -v -o /goapp/gokredit ./cmd/main.go


# Tahap final untuk runtime
FROM alpine:3.14

# Direktori kerja pada image runtime
WORKDIR /app

# Salin file binary hasil build ke image runtime
COPY --from=build /goapp /app/

# Menambahkan binary ke PATH
ENV PATH="/app:${PATH}"

# Membuka port 8082
EXPOSE 8080

# Command yang dijalankan saat container start
ENTRYPOINT [ "gokredit" ]

# Instruksi untuk build image
# docker build -t zikrigusli/kreditplus:1 .