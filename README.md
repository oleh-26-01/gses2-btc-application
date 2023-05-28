# BTC Web API Repository

This repository contains the source code and Dockerfile for the BTC Application, a Go web application built with the Iris framework. If you want to run the project, follow the instructions below.

## Getting Started

To run the BTC Application, you can either build and run the Docker image or execute the project directly using Go.

### Running with Docker

1. Make sure you have Docker installed on your machine.

2. Clone or download this repository to your local machine.

3. Open a terminal or command prompt and navigate to the project directory.

4. Build the Docker image using the following command:

   ```
   docker build -t btc-application .
   ```

   This command will build the Docker image using the provided Dockerfile.

5. Once the image is built, you can run a container based on the image with the following command:

   ```
   docker run -p 8080:8080 btc-application
   ```

   This will start the BTC Application container, and you can access it by navigating to [http://localhost:8080](http://localhost:8080) in your web browser.

   **Note:** If you prefer to download the pre-built Docker image, you can pull it from Docker Hub using the following command:

   ```
   docker pull yarosholeg04/btc-application
   ```

   After pulling the image, you can run a container based on the image with the following command:

   ```
   docker run -p 8080:8080 yarosholeg04/btc-application
   ```

   This will start the BTC Application container, and you can access it by navigating to [http://localhost:8080](http://localhost:8080) in your web browser.

### Running without Docker

1. Make sure you have Go installed on your machine.

2. Clone or download this repository to your local machine.

3. Open a terminal or command prompt and navigate to the project directory.

4. Run the following command to download the required dependencies:

   ```
   go mod download
   ```

5. Once the dependencies are downloaded, you can start the application by running the `main.go` file:

   ```
   go run main.go
   ```

   The application should now be running, and you can access it by navigating to [http://localhost:8080](http://localhost:8080) in your web browser.

## Contributing

If you would like to contribute to this project, feel free to fork the repository and submit a pull request with your changes. Your contributions are greatly appreciated!
