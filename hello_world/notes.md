1. TDD and writing tests before the function forces me to think about how a function will work before I actually implement it, which is good

2. i can commit many times without pushing to master, all commits will be pushed when pushed

3. functions can be given a return variable in setup and simply type return to get return it,
example

    func retString(a string) (temp string) {
        temp = a
        return
    }

this will return temp automatically, if a is empty, an empty string is returned

4. xxx_test.go should be the name of a test

5. tests should `import "testing"`

6. tests can have subtests using `t.Run("some explanation", func (t *testing.T) {..test here..})`

7. `t *testing.T` is required in tests

8. `t.Helper()` in helper functions should be added inside of tests

9. capitalizing the first letter of things in golang export them automatically, only do it, if intended to be used outside of this package

10. `go mod init` or `go mod init <package name>` initializes a module?