package valueobject

import "sync"

type Player struct {
	id         uint64
	components sync.Map
}
