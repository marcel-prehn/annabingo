FROM node:lts-buster-slim as frontend
RUN mkdir /build
COPY ui /build/ui
WORKDIR /build/ui
RUN npm install
RUN npm run build

FROM golang:alpine as backend
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=frontend /build/ui/build /app/ui/build
COPY --from=backend /build/main /app/
WORKDIR /app
CMD ["./main"]