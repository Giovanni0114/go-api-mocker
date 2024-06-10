# Go API Mocker

This project is a Go-based application that generates an API with random values based on a JSON schema. 
It is designed to help developers quickly set up mock APIs for testing and development purposes.

## Getting Started

### Running the Application

1. Ensure you have a valid JSON schema in the `schemas/schema.json` file. Here is an example schema:
    ```json
    {
        "endpoints": [
            {
                "path": "/random/string",
                "type": "string"
            },
            {
                "path": "/random/number",
                "type": "number"
            },
            {
                "path": "/random/bool",
                "type": "bool"
            }
        ]
    }
    ```

2. Run the application:
   ```sh
   go run cmd/go-api-mocker/main.go
   ```

3. The server will start on `http://localhost:8080`. You can access the endpoints defined in the schema, for example:
   - `http://localhost:8080/random/string`
   - `http://localhost:8080/random/number`
   - `http://localhost:8080/random/bool`


## Future Iterations

### Iteration 1: Enhanced Schema Validation and Response Structuring

- Validate the JSON schema more rigorously.
- Support additional data types (e.g., arrays, objects).

### Iteration 2: User-Defined Data Constraints and Endpoint Customization

- Allow users to define constraints for generated data (e.g., string length, number range).
- Enable endpoint customization with query parameters.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


## Disclaimer

In this project code snippets and README.md file was created with help of ChatGPT (version 4o).
