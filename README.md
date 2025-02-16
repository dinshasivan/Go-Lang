# Go-Lang
Go is a programming language created by Google. Go is a fast statically typed, compiled language. It is a general purpose language. Go has built-in testing support. Go is an object-oriented language(in its own way).

## Variable Declaration

```var var_name type```

**Explicit Declaration** 

```var name string```

**Type Inference**

```var name = "value"```

**short variable declaration**

```
    name := "dinsha"
    age := "23"
```

The ```:=``` Operator declares and initialize the varial in one step.

## Format Specifiers 

1. Basic Format Specifiers
    ``` %v``` Defualt format for a value
    ``` %T``` Type of value
    ``` %%``` print a literal %
2. String and Character Format Specifiers
    ``` %s``` String
    ``` %q``` Quoted String
    ``` %c``` Character  
    ``` %U``` Unicode format  
3. Pointer Format Specifier
    ``` %p``` Pointer address

## Array

An array in Go is a fixed-size collection of elements of the same data type. The size of an array is defined at declaration and cannot be changed.

**Declare Array**

```arr := [...]int{10,20,30,40}```

Find array length using **len()** method.

**Slice**

A slice is a dynamic-sized, more flexible alternative to an array.

**Declare Slice and initialize values**

```slice := []string{"Go", "Lang", "Slice"}```

The **make()** function create a slice with a specified length and capacity.

```
    arr := [5]int{1, 2, 3, 4, 5}

    slice := arr[1:4] 
```
The **append()** function in Go is used to add elements to a slice. It dynamically increases the slice's capacity if needed.

## Standard Library 📚
Go's standard library provides a rich set of built-in packages for various functionalities like I/O, string manipulation, cryptography, networking, and more.

**Important packages**
1. **strings:** Manipulating UFT-8 encoded string
2. **strconv:** Implements conversions to and from string       representations of basics data type.
3. **sort:** Provides primitive for sorting slices and user-defined collections.
4. **math:** Provides basic constants and mathematical functions
5. **net/http:**  Provides HTTP client and server implementations.
6. **syn:** Provides basic synchronization primitives.

## Loops
Go has only one looping construct: the for loop. However, it can be used in different ways to achieve while and do-while loop behavior.

**Basic ```for``` loop**
```
    for i := 0; i < 5; i++{
        //statements
    }
```

**```for``` loop as ```while``` loop**
```
    x := 0
	for x < 5{
	    /statement
	    // increment x 
	}
```

## Conditionals & Booleans
Go provides standard conditional statements like ```if```, ```else if```, ```else```, and ```switch```. Boolean values (```true``` / ```false```) are essential in controlling program flow.

## Functions

Functions in Go help in modularizing code and improving reusability. They are defined using the func keyword.

* Use parameteers for input and return values for output
* Go supports multiple retun values.
* Functions can be assigned to varables and passed as arguments

## Package Scope 📦

Go organizes code into packages to manage scope and reusability.
Scope determines where a variable, function, or constant can be accessed within a program.

### Type of Scope
1. Local (Block) Scope - Inside a functions or block {}
2. Package Scope - Accessible across the package
3. Global Scope - If exporteed, accessible from other packages

**Package Scope**

. Variables, constants, and functions declared at the package level are accessible within the entire package.
 
. Use uppercase names to export them for use in other packages

## Map

A map in Go is a built-in data type that stores key-value pairs, similar to dictionaries in Python or objects in JavaScript. Maps are unordered.

Use ```make()``` or map literals to create maps.

Access elements with ```map[key]```. 


Use ```delete(map, key)``` to remove a key. 

Check if a key exists with ```value, exists := map[key]```.

Iterate with ```for range```.

## Pass By Value
In Go, function arguments are passed by value by default. This means that when you pass a variable to a function, a copy of the variable is created, and modifications inside the function do not affect the original variable.
For structs, this means an entire copy is passed.

### Pointers
A pointer in Go is a variable that stores the memory address of another variable. Instead of holding an actual value, it "points" to where the value is stored in memory.

Use pointers (```*T```) to modify the original variable.

### Struct
A struct is a collection of fields. It groups related values under a single type.


### Receiver Functions
A receiver function is a function with a receiver parameter, which is a special parameter before the function name. It defines which type the function belongs to.

```
    func (receiver Type) function_name(parameters) retun_type{
        //function body
    }
```
. Receiver functions define methods for structs and custom types.

. Value receivers do not modify the original struct (pass by value).

. Pointer receivers modify struct fields (pass by reference).

. Use receiver functions to encapsulate behavior inside structs.

