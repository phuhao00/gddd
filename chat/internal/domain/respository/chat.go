package respository

import (
	"github.com/phuhao00/gddd/chat/internal/domain/valueobject"
)

type IChatRepository interface {
	LoadGroupPlayers(groupKey string) []valueobject.Player
	LoadGroupMessages(groupKey string) []valueobject.Message
	SaveGroupMessages(groupKey string, messages []valueobject.Message)
	AddGroupMessages(groupKey string, message valueobject.Message)
}
