package reportusecase

import (
	"encoding/json"

	"github.com/bishal05das/blog_app_1/internal/domain"
)

type JsonFormatterUseCase struct{}

func NewJsonFormatterUseCase() *JsonFormatterUseCase {
	return &JsonFormatterUseCase{}
}

func (uc *JsonFormatterUseCase) Format(data any) string {
	switch v := data.(type) {
    case *domain.Post:
        b, _ := json.MarshalIndent(v, "", "  ")
        return string(b)
    case *domain.User:
        b, _ := json.MarshalIndent(v, "", "  ")
        return string(b)
    default:
        return "{}"
    }
}


