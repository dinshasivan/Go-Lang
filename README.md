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
		//statement
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

