package reportusecase

type Report interface {
	Generate(id int,formatter Formatter) string
}