package domain

import (
	"container/ring"
	"github.com/phuhao00/gddd/chat/internal/domain/valueobject"
	"sync"
)

type Chat struct {
	group2msg     map[string]*ring.Ring
	group2players map[string]sync.Map
}

func (c *Chat) AddGroupPlayer(groupKey string, player valueobject.Player) {

}

func (c *Chat) DelGroupPlayer(groupKey, playerUID string) {

}

func (c *Chat) BroadCastMsg(groupKey string, message *valueobject.Message) {

}
