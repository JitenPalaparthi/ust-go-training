## basics

- Find 1-n prime numbers
- find out the biggest number among a,b,c through a function

## range loop

- reverse the string using range loop
- count number of spaces in a string
- find number of words in a string

## Arrays

- find sum of elements in an array
- reverse an array
- take input of an array and return another array with only prime numbers
- take three arrays ids, names, salaries and write tax computation
  - sal > 12000 then tax is 10%
  - sal > 20000 then tax is 12%
  - sal > 50000 then tax is 20%
  - print emply id, name , sal and his tax

## Slice

- create and instantiate a slice and reverse a slice
- create and instantiate a slice find the repeated elements in the slice
- take three slices ids, names, salaries and write tax computation
  - sal > 12000 then tax is 10%
  - sal > 20000 then tax is 12%
  - sal > 50000 then tax is 20%
  - print emply id, name , sal and his tax
- create two 2D slices , assign values and perform matrix addition and substraction using functions

## Variadic

- Create a sumOf function. Use ...interface{} as a parameter, perform sumOf based on input types as int8,int16,int32,int64,int,uint8,uint16,uint32,uint64,float32,float64 and string. Make sure all values in variadic argument are same. Fill the stub in 23-variadic program.
  
  func sumOfAny(vals ...interface{})interface{}{
      // implement here
    return nil
  }

## map

- pass two slices of same type and same length as parameters
  - 1- consider elements of first slice is key and second slice as value store them in a map.
  - 2- The function returns map and an error
  - if length of slice1 is not equal to length of slice2 then return error
  - if slice one has duplicate elements return error becase duplicate elements cannot be given as keys.
- Write a function to convert a map to a string of json form and also to byte array

## pointers

- Extend sunOf(v ...interface{})(interface{}) to use pointers. That means even pointers can be passed as arguments and then the function must work.

## structs

- Create a package named shapes.
  - Create sub packages rect,square,triangle and circle
  - Create a func in each package called New()
  - New must be like a constructor with required parameters and returns Rect,Square,Circle and Triangle with respective to the package. That means in the package Rect , the new function takes Length and Width and returns a &Rect, similarly for all as per their parameters into their respective packages. Make sure New is not a method it is a function in a package that works similar to Constructor
  - create methods as AreaOf and PerimeterOf for each and every sub package. That means for rect, square, triangle and circle.
  - Call them all in main

## user defined types

- Implement a type from a type called string
  - create a method called Greet
  - implement another type from the existing type
  - again create a method calle Greet
  - lets myString1 and myString2 have two methods Greet and Greet
  - call them both using type casting

- Create a user defined map
  - add functions Insert(Key,value), Update(key,value), Delete(key), GetKeys()[]interface{},GetValues()[]interfaces{}

## interface

- Create an interface Perimeter()float32 in the existing example 36-interfaces
  - Implement IPerimeter for Rect and Square types
  - Call it in the main function like the way Area is called

- Add Circle and Triangle to the shapes package in example 36-interfaces
  - Implement IArea and IPerimeter
  - Create objects for Square and Triangle
  - Call them in Area and Perimeter Dependency functions.

## anonymous func

- create anonymous functions to perform add, sub, multiply and divide operatios.
- store those functions in variables
- call them using variables. Used in line no 18 in example 37

## goroutines

- run all anonymous functions from previous task as go routines

## channels

- write a function called increment.
  - call the function 100 times in a loop as a goroutine
  - the final value after calling increment func for 100 times should be 100
  - that means increment starts from 1 to 100 and finally it must print 100

- implement greet function in program no 40 without using time.Sleep instead use channels

## channels and waitGroup

- re-implement example 46-mutex with channel and WaitGroup rather than mutex and WaitGroup

## run postgress in docker

  - Step-0:install docker from docker.com
  - https://docs.docker.com/engine/install/ubuntu/
  - step-1 ```docker pull postgres```
  - step-2 ```docker run -d --name pg -e POSTGRES_PASSWORD=admin123 -e POSTGRES_USER=admin -e POSTGRES_DB=contactsdb -p 5432:5432 postgres```
  - step-3: ensure postgres container is running ```docker ps```

## Create Calculator with Generics

- Addition, Substraction, Multiplication using generics

## Delete operation

- Currently in contacts demo, the delete operation always returns 1. Modify it to 0 if no delete happens.

## getAll operation in contacts microservice

- implement getAllContacts in contacts microservice
  
## implement a JWT token for echo and gin

- jwt token to be used. In order to do that
-  create user
-  upon successful user login, it returns jwt token
-  for all subsequent requests must use that token to perform create, update, get and delete contact operations

## Install grpc tools
- Install BloomRPC --> Its a client tool similar to Postman but for grpc
- Follow the documentation steps: https://community.chocolatey.org/packages/bloomrpc
- Install protoc compiler from https://github.com/protocolbuffers/protobuf/releases/tag/v21.3
- Install protoc gen for go. Run below two commands in the terminal or powershell
- ```go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28```
- ```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2```