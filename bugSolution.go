func calculate(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("division by zero")
  }
  return a / b, nil
}

func main() {
  result, err := calculate(10, 0)
  if err != nil {
    fmt.Println("Error:", err) // Handle the error explicitly
    return
  }
  fmt.Println("Result:", result)
} 