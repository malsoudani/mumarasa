* Basic Types in Go: 
    - `bool`
    - `string`
    - `int, int8, int16, int32, int64`
    - `uint, uint8, uint16, uint32, uint64, uintptr`.  `uintptr` is a big enough integer to hold a memory address.
    - `byte (uint8)`
    - `rune (int32)`
    - `float32, float64`
    - `complex64, complex128`

* Other Types:
    - Array: like other programming languages
    - Slice: It's more like a vector or a list in other programming languages, in the since that its an array that can grow in size.
    - Struct: much like other programming languages this is just a collection of fields. 
    - Pointer: any type can be a pointer by holding the memory address of that given type.
    - Function: a function in Go can be used as a type too!, it can be stored, passed, .. etc. Kinda like JavaScript.
    - Interface: these aren't like Java's interfaces for example, you don't have to explictly implement them, although they still define function like other programming languages interfaces.
    - Map: Much like other programming languages, basically a key, value store.
    - Channels: can be used as communication between Goroutines, More on that later.

* User Defined Types
    - In Go there is no inheretence and other oop techniques used.
    - to compensate for that we have user defined types that are declared as follows : `type MySpecialType redifenition-of`
    example: `type Greeting string`

* Function in Go
    - Go functions can have multiple return types
    - A function in go can be used like any other type.
    - Function literals, lets us define function within a function where the inner function can remember the context that the function was declated at, much like clousures in other languages.

* passing functions in go: 
    - functions can be treated as variables to the extent of passing it and assigning it to a variable function that takes a given parameter.
    - a function can be used as a custom type example: `type MySpecialFunction func(string) ()`
    - closures can be utilized to define a function to be called within a parent function, and to place a different customization level on the function.