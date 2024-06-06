# Use an official OpenJDK runtime as a parent image
FROM openjdk:17-jdk-slim

# Set the working directory in the container
WORKDIR /app

# Copy the application's JAR file into the container at /app
COPY  ./target/IngestorJob-1.0-SNAPSHOT.jar /app/app.jar

# Command to run the Flink job
CMD ["flink", "run", "-c", "consumer.App", "/app/app.jar"]

