## Variables In Go
Notes:
1. In Go a variable belongs and created at Runtime: only when the program runs can a variable created in memory
2. At compile time it only knows the type of the variable but doesn't assign any value to it
3. In Go a declared variable must be used else Go will return an error due to its high memory management
4. There are two ways to declare a variable in Go using `var` keyword or `:=`
5. variable declaration syntatic sugar `var <name> <type> = <value>`
6. When the variable value is explicit, you can omit it because Go can infer the variable type
7. When you want to declare a variable that you will not use in Go, make sure to use the blank identifier known as `_` else you will have a compile time error
8. The short declaration operator works only in block scopes `s := "Best teacher in the world"`
9. Cannot use `:=` for already defined variables, if you want to change the value of a named variable you use only the `=` operator
10. Also note that `var (
    salary float64
    firstName string
    gender bool
    )` is not the same as `var (salary float64 firstName string gender bool)`
11. In summary use short declaration `:=` when we know the initial value and normal declaration `var` otherwise or any time
12. Use normal declaration for multiple declaration for clarity


## GPT modified
Let's evaluate each statement:

1. **True:** In Go, a variable is created at runtime. Memory for the variable is allocated when the program runs.
2. **True:** At compile time, Go only knows the type of the variable, but it doesn't assign any value to it.
3. **True:** If a declared variable is not used, Go returns an error during compilation.
4. **True:** There are two ways to declare a variable in Go: using the `var` keyword or using the `:=` short declaration operator.
5. **False:** The correct syntax for variable declaration in Go is `var <name> <type> = <value>`.
6. **True:** When the variable value is explicit, you can omit the type because Go can infer it.
7. **False:** You can declare a variable and not use it without causing a compile-time error.
8. **True:** The short declaration operator (`:=`) works only in block scopes.
9. **True:** You cannot use `:=` for already defined variables; you use only the `=` operator to change the value of a named variable.
10. **False:** `var` with parentheses or without them is the same in Go. Both are valid ways of declaring multiple variables.
11. **True:** Use short declaration (`:=`) when you know the initial value, and use normal declaration (`var`) otherwise or when declaring multiple variables for clarity.
12. **True:** Use normal declaration for multiple declarations for clarity.


## Correction

Sure, let's correct the statements that are false:

5. **True:** The correct syntax for variable declaration in Go is `var <name> <type>`, and you can optionally assign a value. The `= <value>` part is not required for declaration, but it can be used to initialize the variable.

7. **True:** You can declare a variable and not use it without causing a compile-time error in Go.

Let me know if you have any more questions or if you'd like further clarification on any statement!