1. Very large section - a lot to learn, going to redo
2. `numArray := [5]int{1,2,3,4,5}`
3. `sliceArray := []int{1,2,3,4,5,6,7`
4. `%v` displays default for value in printing with Errorf(), good for arrays/slices
5. arrays/slices indexed like normal `arr[i]`
6. `_` blank identifier for loops
7. `range` good for iterating over arrays/slices ex. `for _, val := range numbs {...}`
8. size of array is encoded in type so [5]int is not same type as [7]int. compile error
9. size of slices `DO NOT` get encoded as type
10. arrays and slices are not compatable
11. each test has a cost. test enough to have confidence that a function works as intended but not much further
12. `go test -cover` helps find edge cases in your tests
13. slices cannot be compared with equality operators in go. using `reflect.DeepEqual` *may* be helpful to cheat this. tests to see if any of the variables are the same
14. `reflect.DeepEqual` is not type safe
15. `len(varName)` returns length of variable
16. `make([]int, sizeVar)` creates a slice of starting capacity of sizeVar
17. slices can be sliced using [x:y]. ex. `tail := sliceName[1:]`, tail is equal to sliceName, minus the value and the 0th index
18. `...T` takes as many parameters as you send it of type T. Example func `func addAll(nums ...int)int {..}`