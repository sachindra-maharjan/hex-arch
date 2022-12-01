package ports

type DBPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
}



