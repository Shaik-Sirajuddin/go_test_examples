# Go Test Examples

This project contains simple Go examples with unit tests to demonstrate how to run tests in Go.

## Project Structure

- `math/`: Simple mathematical functions.
    - `Add(a, b int) int`: Returns the sum of two integers.
    - `Subtract(a, b int) int`: Returns the difference between two integers.
- `stringutils/`: String utility functions.
    - `Reverse(s string) string`: Returns the reversed string.
    - `IsEmpty(s string) bool`: Returns true if the string is empty.

## How to Run Tests

To run all tests in the project, use the following command:

```bash
go test ./...
```

For verbose output:

```bash
go test -v ./...
```

## CI/CD

The project includes a `Jenkinsfile` for CI/CD integration and a helper script `.jenkins_actions/go-test.sh` to execute tests in a Jenkins pipeline.
