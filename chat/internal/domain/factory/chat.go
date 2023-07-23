package factory

import (
	"github.com/phuhao00/gddd/chat/internal/domain"
)

type Chat struct{}

func (c *Chat) Generate() domain.Chat {
	return domain.Chat{}
}
