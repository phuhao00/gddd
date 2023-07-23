package usecase

import (
	"github.com/phuhao00/gddd/chat/internal/application/service"
	"github.com/phuhao00/gddd/chat/internal/domain/valueobject"
)

func Private(chat service.IChat, message valueobject.Message) {
	chat.Private(message)
}
