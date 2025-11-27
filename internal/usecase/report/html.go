package reportusecase

import (
	"fmt"
	"strings"

	"github.com/bishal05das/blog_app_1/internal/domain"
)

type HtmlFormatterUseCase struct{}

func NewHtmlFormatterUseCase() *HtmlFormatterUseCase {
	return &HtmlFormatterUseCase{}
}

func (uc *HtmlFormatterUseCase) Format(data any) string {
	var sb strings.Builder
	sb.WriteString("<html><body><ul>")
	switch v := data.(type) {
	case *domain.Post:
		sb.WriteString(fmt.Sprintf("<li><b>ID:</b> %d</li>", v.ID))
		sb.WriteString(fmt.Sprintf("<li><b>Title:</b> %s</li>", v.Title))
		sb.WriteString(fmt.Sprintf("<li><b>UserID:</b> %d</li>", v.UserID))
		sb.WriteString(fmt.Sprintf("<li><b>Score:</b> %d</li>", v.Score))
	case *domain.User:
		sb.WriteString(fmt.Sprintf("<li><b>ID:</b> %d</li>", v.ID))
		sb.WriteString(fmt.Sprintf("<li><b>Name:</b> %s</li>", v.Name))
		sb.WriteString(fmt.Sprintf("<li><b>IsAdmin:</b> %t</li>", v.IsAdmin))
		sb.WriteString(fmt.Sprintf("<li><b>Reputation:</b> %d</li>", v.Reputation))
	default:
		return "<html><body>Unsupported data type</body></html>"
	}
	sb.WriteString("</ul></body></html>")
	return sb.String()
}
