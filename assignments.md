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