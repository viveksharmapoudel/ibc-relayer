package polygon

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyperledger-labs/yui-relayer/core"
)

type Chain struct {

	// specs related to chain
}

// ****************implement chainI interface****************

// ChainID returns ID of the chain
func (c *Chain) ChainID() string {
	return ""
}

// GetLatestHeight gets the chain for the latest height and returns it
func (c *Chain) GetLatestHeight() (int64, error) {
	return 0, nil
}

// GetAddress returns the address of relayer
func (c *Chain) GetAddress() (sdk.AccAddress, error) {
	return nil, nil
}

// Codec returns the codec
func (c *Chain) Codec() codec.ProtoCodecMarshaler {
	return nil
}

// SetRelayInfo sets source's path and counterparty's info to the chain
func (c *Chain) SetRelayInfo(path *core.PathEnd, counterparty *core.ProvableChain, counterpartyPath *core.PathEnd) error {
	return nil
}

// Path returns the path
func (c *Chain) Path() *core.PathEnd {
	return nil
}

// SendMsgs sends msgs to the chain
func (c *Chain) SendMsgs(msgs []sdk.Msg) ([]byte, error) {
	return nil, nil
}

// Send sends msgs to the chain and logging a result of it
// It returns a boolean value whether the result is success
func (c *Chain) Send(msgs []sdk.Msg) bool {
	return false
}

// Init initializes the chain
func (c *Chain) Init(homePath string, timeout time.Duration, codec codec.ProtoCodecMarshaler, debug bool) error {
	return nil
}

// SetupForRelay performs chain-specific setup before starting the relay
func (c *Chain) SetupForRelay(ctx context.Context) error {
	return nil
}

// RegisterMsgEventListener registers a given EventListener to the chain
func (c *Chain) RegisterMsgEventListener(core.MsgEventListener) {
	return
}
