package lib

type InputInterface interface {
	GetNextLine() *string
	Close()
}

func GetNextLine(i InputInterface) *string {
	return i.GetNextLine()
}

func Close(i InputInterface) {
	i.Close()
}
