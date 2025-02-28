# Use a lightweight Python image
FROM python:3.11-slim


# Set the working directory
WORKDIR /app


# Install dependencies
COPY reqs.txt /app/
RUN pip install  --prefer-binary -r reqs.txt

COPY . .

# Expose the gRPC server port
EXPOSE 50051

# Command to run the server
CMD ["python", "server.py"]
