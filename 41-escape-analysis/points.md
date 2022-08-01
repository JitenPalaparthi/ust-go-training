### Escape analysis

- As a go developer, how do you know whether variable is stored in stack or heap?
- This can be done through escape analysis

### To perform ecsape analysis

```go tool compile -m main.go```
or
```go build -gcflags="-m" main.go```

### Main points in escape analysis

1. escapes to heap: particular variable has to be shared between main and other function or method;hence,at compile time itself it is decided to store in heap.
2. does not escape to heap: variable is not stored in heap memory;hence, less burden on GC.
3. moved to heap: variable is stored in heap memory.