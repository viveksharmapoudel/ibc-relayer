package icon

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/hyperledger-labs/yui-relayer/core"
)

type Prover struct {
	chain *Chain
	//config *ProverConfig
}

// var _ core.ProverI = (*Prover)(nil)

func NewProver(chain *Chain) *Prover {
	return &Prover{chain: chain}
}

// ****************implement proverI interface****************

// initalize prover
func (pr *Prover) Init(homePath string, timeout time.Duration, codec codec.ProtoCodecMarshaler, debug bool) error {
	return nil
}

// SetRelayInfo sets source's path and counterparty's info to the chain
func (pr *Prover) SetRelayInfo(_ *core.PathEnd, _ *core.ProvableChain, _ *core.PathEnd) error {
	return nil // prover uses chain's path instead
}

func (pr *Prover) SetupForRelay(ctx context.Context) error {
	return nil
}

// GetChainID returns the chain ID
func (p *Prover) GetChainID() string {
	return ""
}

// QueryHeader returns the header corresponding to the height
func (p *Prover) QueryHeader(height int64) (out core.HeaderI, err error) {
	return nil, nil
}

// QueryLatestHeader returns the latest header from the chain
func (p *Prover) QueryLatestHeader() (out core.HeaderI, err error) {
	return nil, nil
}

// GetLatestLightHeight returns the latest height on the light client
func (p *Prover) GetLatestLightHeight() (int64, error) {
	return 0, nil
}

// CreateMsgCreateClient creates a CreateClientMsg to this chain
func (p *Prover) CreateMsgCreateClient(clientID string, dstHeader core.HeaderI, signer sdk.AccAddress) (*clienttypes.MsgCreateClient, error) {
	return nil, nil
}

// SetupHeader creates a new header based on a given header
func (p *Prover) SetupHeader(dst core.LightClientIBCQueryierI, baseSrcHeader core.HeaderI) (core.HeaderI, error) {
	return nil, nil
}

// UpdateLightWithHeader updates a header on the light client and returns the header and height corresponding to the chain
func (p *Prover) UpdateLightWithHeader() (header core.HeaderI, provableHeight int64, queryableHeight int64, err error) {
	return nil, 0, 0, nil
}

// QueryClientConsensusState returns the ClientConsensusState and its proof
func (p *Prover) QueryClientConsensusStateWithProof(height int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	return nil, nil
}

// QueryClientStateWithProof returns the ClientState and its proof
func (p *Prover) QueryClientStateWithProof(height int64) (*clienttypes.QueryClientStateResponse, error) {
	return nil, nil
}

// QueryConnectionWithProof returns the Connection and its proof
func (p *Prover) QueryConnectionWithProof(height int64) (*conntypes.QueryConnectionResponse, error) {
	return nil, nil
}

// QueryChannelWithProof returns the Channel and its proof
func (p *Prover) QueryChannelWithProof(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	return nil, nil
}

// QueryPacketCommitmentWithProof returns the packet commitment and its proof
func (p *Prover) QueryPacketCommitmentWithProof(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	return nil, nil
}

// QueryPacketAcknowledgementCommitmentWithProof returns the packet acknowledgement commitment and its proof
func (p *Prover) QueryPacketAcknowledgementCommitmentWithProof(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	return nil, nil
}
