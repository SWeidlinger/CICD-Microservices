FROM golang:1.22-alpine

# Set maintainer label: maintainer=[YOUR-EMAIL]
LABEL maintainer='s.weidlinger99@gmx.at'

# Set working directory: `/src`
WORKDIR /src

# Copy local file `main.go` to the working directory
COPY *.go .
COPY *.mod .
COPY *.sum .

# Build the GO app as myapp binary and move it to /usr/
RUN go build -o /usr/myapp ./...

#Expose port 8888
EXPOSE 8010

# Run the service myapp when a container of this image is launched
CMD ["/usr/myapp"]