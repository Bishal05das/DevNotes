package reportusecase

type Formatter interface {
	Format(data any) string
}
