package wxserver

import (
	"github.com/AAllehyany/wx-simulator/wxcore"
	"github.com/google/uuid"
)

type GameServer struct {
	Games map[uuid.UUID]wxcore.GameState
}
