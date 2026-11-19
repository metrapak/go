package methods_interfaces

// A method is a function with a special receiver argument.
//
//	The receiver appears in its own argument list between the func keyword and the method name.
//	In this example, the Abs method has a receiver of type Vertex named v.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	abs := v.Abs()
	println(abs)
}
