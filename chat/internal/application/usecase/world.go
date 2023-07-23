package usecase

import (
	"github.com/phuhao00/gddd/chat/internal/application/service"
	"github.com/phuhao00/gddd/chat/internal/domain/valueobject"
)

func World(chat service.IChat, message valueobject.Message) {
	chat.WorldChat(message)
}
