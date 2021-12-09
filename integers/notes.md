1. if a function takes mulitple arguments of teh same type, the type only has to be given on the last one
    func Add(a, b int)
is the same as
    func Add(a int, b int)

2. adding examples to tests can be good. Examplexxx()

3. `go test -v` is a more verbose way to show tests that are ran

4. on examples a comment with `// Output: xxxxx` is required for it to be ran

5. forgot to add last time `godoc -http=:8080` runs the godocs as a local server