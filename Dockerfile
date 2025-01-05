# Step 1: Use the official Go image as the base image
FROM golang:1.23

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go modules and dependencies
COPY go.mod go.sum ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy the application source code
COPY . .

# Step 6: Build the Go application
RUN go build -o main .

# Step 7: Expose the application port
EXPOSE 8000

# Step 8: Command to run the application
CMD ["./main"]