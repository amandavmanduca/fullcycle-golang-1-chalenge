package interfaces

type FileHandlerInterface interface {
	Write(text string) error
}
