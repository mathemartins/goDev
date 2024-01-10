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
10. 