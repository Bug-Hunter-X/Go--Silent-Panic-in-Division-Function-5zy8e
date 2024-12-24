# Go: Silent Panic in Division Function

This example demonstrates a common error in Go: silently ignoring errors in a function that returns multiple values (including an error). When the `calculate` function encounters a division-by-zero error, it returns an error object.  However, if the calling code doesn't explicitly check for and handle this error, the program will panic.

The `bug.go` file demonstrates the issue; `bugSolution.go` provides the fix.

## How to Reproduce
1. Save `bug.go`.
2. Run it using `go run bug.go`.
3. Observe the panic.

## Solution
The solution is to always check the error returned by a function that signals errors through multiple return values.  The `bugSolution.go` file shows the correct way to handle the error.