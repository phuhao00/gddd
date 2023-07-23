package service

import (
	"github.com/phuhao00/gddd/chat/internal/domain/valueobject"
)

type IChat interface {
	Private(message valueobject.Message)
	WorldChat(message valueobject.Message)
}
