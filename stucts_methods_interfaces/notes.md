1. structs are as follows
    type Rectangle struct {
        Width float64
        Height float64
    }
2. make a struct like `rectangle := Rectangle{10.0, 12.0}` or `rectangle := Rectangle{Width: 10.0, Height: 12.0}`
3. access a field like `rectangle.Width`
4. methods for structs are as follows
    func (r Rectangle) Area() float64 {
        return Width * Height 
    }
5. func (receiverName RecieverType) MethodName(args)
6. interfaces, any type that satisfies the interface can be used. 
7. interfaces are as follows
    type Shape interface {
        Area() float64
    }
8. table driven tests
    areaTests := []struct {
        shape Shape
        want float64
    }{
        {shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
        {shape: Circle{Radius: 10}, want: 314.1592653589793},
    }
9. above is an anonymous structs for the slice of structs
10. very easy to add to the test table if near identical tests are being ran for each test
11. adding a name field to the test is good for ease of use.
12. `"%#v got %g want %g"` the %#v will print the test struct
13. interfaces hide complexity to other parts of the program