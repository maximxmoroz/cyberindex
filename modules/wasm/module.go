package wasm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cybercongress/cyberindex/database"
	"github.com/forbole/juno/v2/modules"
	"github.com/forbole/juno/v2/modules/messages"
	"github.com/forbole/juno/v2/types"
)

var (
	_ modules.Module 		= &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
	messagesParser  messages.MessageAddressesParser
	cdc 		   	codec.Codec
	db              *database.CyberDb
}

func NewModule(
	messagesParser messages.MessageAddressesParser,
	cdc 		   codec.Codec,
	db *database.CyberDb,
) *Module {
	return &Module{
		messagesParser:  messagesParser,
		cdc: 			 cdc,
		db:              db,
	}
}

func (m *Module) Name() string {
	return "wasm"
}

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *types.Tx) error {
	return HandleMsg(tx, index, msg, m.db)
}
