## Constants

1. Constants are created at compile time and cannot be changed when the program is running
2. We can declare a const in Go without actually using the constant. The go-compiler would not return an error
this is because constants are created at compile time and value assigned also at compiled time and memory has already
been allocated to it, and it cannot be changed for any reason. Unlike variables because at compile time, Go knows
about the variable and its type but doesn't assign a value to it instead reserves a memory for the variable type
through inferring or if the variable type is explicitly stated and this variable can be changed because it is given
at runtime and can be changed again during runtime
3. Constants must be initialized when declared
4. In Go constants belong to compile-time, so errors are discovered early while variables belong to runtime so
   errors are discovered quite late example 
   `
     x, y := 5, 0
     fmt.Println(x/y) // Supposed to get an error but would not get an error at compile time because this is a 
                      // variable problem and these values would be assign to them at runtime hence cannot be
                      // detected early enough during compilation
     const x int = 5
     y := 0
     fmt.Println(x/y) // Error would be detected early due to compile time handling for constants
   `
5. In a grouped constants, a constant repeats the previous one 
    ` const (
       min1 = -500
       min2
       min3
      )
      fmt.Println(min1, min2, min3) => -500 -500 -500
    `