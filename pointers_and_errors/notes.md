1. function and method parameters are copied, changing the copy doesnt change the original
2. & address of operator gets the address like other langs
3. it is possible to dereference like `(*w).balance`, but go does it automatically so `w.balance` works just as well
4. nil is null in other langs
5. errors can return nil, and should if no error
6. trying to access a nil value results in `runtime panic`
7. `errors.New("this is an error")` creates a new error
8. `if got.Error() != want {}` is a way to test for errors
9. `t.Fatal("some statement")` will stop the test and return the statement.
10. `var ErrInsufficentFunds = errors.New("cannot withdraw more money than you have")` allows for a global, reusable error
11. the 3rd party `errcheck` module allows to find unchecked errors
12. pointers can be nil
13. the compiler will not help determine if pointers are nil
14. making new types from other types, or grouping useful types together with a useful name, is encouraged. 
15. the new type can be used with interfaces