1. `dictionary := map[string]string{"test" : "this is a test"}`
2. same as declaring an array except requires two types
3. keys must be a comparable type
4. value type can be anything, even a map
5. `return d[key]` returns value
6. it is good to wrap a map in a type such as `type Dictionary map[string]string`, this makes it more consise to create
7. maps feel like a reference type, but it is wrong, passing a map copies the pointer
8. empty maps are teh same as nil maps, should never be setup, never do `var m map[string]string` instead do `var m = map[string]string{}` or `var m = make(map[string]string)`
9. writing to a nil or empty map will cause runtime panic
10. multiple vars can be initialized like so
    var (
        ErrNotFound = errors.New("error not found")
        ErrWordExists = errors.New("error word already exists")
    )
11. `delete(mapName, word)` will remove a k:v from the map