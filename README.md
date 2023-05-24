This is a complete Go application that builds a REST API with a single endpoint `/price?brandID={brandID}&applicationDate={applicationDate}&productID{productID}`. It accepts input parameters `brandID`, `applicationDate`, and `productID` and returns the corresponding price information from an in-memory database.

The database is initialized with example data, and the `GetPrice` function retrieves the price based on the input parameters. The response is formatted as JSON and returned to the client.


# Setup

1. Make sure you have Go installed on your system.
2. Clone the repository:

        git clone https://github.com/PaulPROST/APIREST_application.git

4. Navigate to the project directory:
5. Install the dependencies:

        go mod tidy


# Running the Service

To start the service, run the following command:

    go run main.go

By default, the service will be accessible at **'http://localhost:8080'**.


# Endpoints

You can then make requests to the endpoint using tools like cURL or Postman. For example, to test the first case, you can use the following command:

    curl -i "http://localhost:8080/price?brandID=1&date=2020-06-14-16.00.00&productID=35455"

The response should contain the price information for the given parameters.

# Testing

To run the tests, use the following command:

    go test ./internal/transport/http/