package ports

type ArithmeticPort interface {
	Addition(a, b int32) (int32, error)
	Substraction(a, b int32) (int32, error)
	Multipication(a, b int32) (int32, error)
	Division(a, b int32) (int32, error)
}
