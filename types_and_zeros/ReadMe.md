## Types and Zeros In Go
Notes:
1. Go is a static type programming language and that means it does type checking at compile time
2. If you don't provide the type, it should be able to automatically infer the type at compile time, if not provided explicitly
3. We cannot assign a float to an int in Go
4. The type of variable cannot be changed once declared, because Go registers the variable type at compile time and would stay the same throughout the code build lifecycle
5. Type conversion in Go can be done, and this helps the executable to act during compile time and prepare a memory for the allocated value to be changed `var a = "5"; b := 3; b = int(a)`
   `var a = "5"
      b := 3
      b = int(a) // This line will result in a compilation error. You should use strconv.Atoi(a) for string to int conversion.
   `
6. In Go there is no such thing as uninitialized variable, every variable is initialized at runtime and knows its type and only assigns value to it at runtime, but knows about the existence of a variable at compile time and knows the type
7. Zero values or default values in Go `numeric types = 0; bool type = false; string type = ""; pointer type = nil`
8. 