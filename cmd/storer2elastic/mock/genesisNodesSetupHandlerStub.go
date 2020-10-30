package mock

import (
	"time"

	"github.com/ElrondNetwork/elrond-go/sharding"
)

// GenesisNodesSetupHandlerStub -
type GenesisNodesSetupHandlerStub struct {
	AllInitialNodesCalled                func() []sharding.GenesisNodeInfoHandler
	InitialNodesInfoForShardCalled       func(shardId uint32) ([]sharding.GenesisNodeInfoHandler, []sharding.GenesisNodeInfoHandler, error)
	InitialNodesInfoCalled               func() (map[uint32][]sharding.GenesisNodeInfoHandler, map[uint32][]sharding.GenesisNodeInfoHandler)
	GetStartTimeCalled                   func() int64
	GetRoundDurationCalled               func() uint64
	GetChainIdCalled                     func() string
	GetMinTransactionVersionCalled       func() uint32
	GetShardConsensusGroupSizeCalled     func() uint32
	GetMetaConsensusGroupSizeCalled      func() uint32
	MinNumberOfShardNodesCalled          func() uint32
	MinNumberOfMetaNodesCalled           func() uint32
	GetHysteresisCalled                  func() float32
	GetAdaptivityCalled                  func() bool
	NumberOfShardsCalled                 func() uint32
	MinNumberOfNodesCalled               func() uint32
	MinNumberOfNodesWithHysteresisCalled func() uint32
}

// InitialNodesPubKeys -
func (g *GenesisNodesSetupHandlerStub) InitialNodesPubKeys() map[uint32][]string {
	panic("implement me")
}

// GetShardIDForPubKey -
func (g *GenesisNodesSetupHandlerStub) GetShardIDForPubKey(pubkey []byte) (uint32, error) {
	panic("implement me")
}

// InitialEligibleNodesPubKeysForShard -
func (g *GenesisNodesSetupHandlerStub) InitialEligibleNodesPubKeysForShard(shardId uint32) ([]string, error) {
	panic("implement me")
}

// InitialNodesInfoForShard -
func (g *GenesisNodesSetupHandlerStub) InitialNodesInfoForShard(shardId uint32) ([]sharding.GenesisNodeInfoHandler, []sharding.GenesisNodeInfoHandler, error) {
	if g.InitialNodesInfoForShardCalled != nil {
		return g.InitialNodesInfoForShardCalled(shardId)
	}

	return nil, nil, nil
}

// InitialNodesInfo -
func (g *GenesisNodesSetupHandlerStub) InitialNodesInfo() (map[uint32][]sharding.GenesisNodeInfoHandler, map[uint32][]sharding.GenesisNodeInfoHandler) {
	if g.InitialNodesInfoCalled != nil {
		return g.InitialNodesInfoCalled()
	}

	return nil, nil
}

// GetStartTime -
func (g *GenesisNodesSetupHandlerStub) GetStartTime() int64 {
	if g.GetStartTimeCalled != nil {
		return g.GetStartTimeCalled()
	}

	return time.Now().Unix()
}

// GetRoundDuration -
func (g *GenesisNodesSetupHandlerStub) GetRoundDuration() uint64 {
	if g.GetRoundDurationCalled != nil {
		return g.GetRoundDurationCalled()
	}

	return 4500
}

// GetChainId -
func (g *GenesisNodesSetupHandlerStub) GetChainId() string {
	if g.GetChainIdCalled != nil {
		return g.GetChainIdCalled()
	}

	return "chainID"
}

// GetMinTransactionVersion -
func (g *GenesisNodesSetupHandlerStub) GetMinTransactionVersion() uint32 {
	if g.GetMinTransactionVersionCalled != nil {
		return g.GetMinTransactionVersionCalled()
	}

	return 1
}

// GetShardConsensusGroupSize -
func (g *GenesisNodesSetupHandlerStub) GetShardConsensusGroupSize() uint32 {
	if g.GetShardConsensusGroupSizeCalled != nil {
		return g.GetShardConsensusGroupSizeCalled()
	}

	return 1
}

// GetMetaConsensusGroupSize -
func (g *GenesisNodesSetupHandlerStub) GetMetaConsensusGroupSize() uint32 {
	if g.GetMetaConsensusGroupSizeCalled != nil {
		return g.GetMetaConsensusGroupSizeCalled()
	}

	return 1
}

// MinNumberOfShardNodes -
func (g *GenesisNodesSetupHandlerStub) MinNumberOfShardNodes() uint32 {
	if g.MinNumberOfShardNodesCalled != nil {
		return g.MinNumberOfShardNodesCalled()
	}

	return 1
}

// MinNumberOfMetaNodes -
func (g *GenesisNodesSetupHandlerStub) MinNumberOfMetaNodes() uint32 {
	if g.MinNumberOfMetaNodesCalled != nil {
		return g.MinNumberOfMetaNodesCalled()
	}

	return 1
}

// GetHysteresis -
func (g *GenesisNodesSetupHandlerStub) GetHysteresis() float32 {
	if g.GetHysteresisCalled != nil {
		return g.GetHysteresisCalled()
	}

	return 0
}

// GetAdaptivity -
func (g *GenesisNodesSetupHandlerStub) GetAdaptivity() bool {
	if g.GetAdaptivityCalled != nil {
		return g.GetAdaptivityCalled()
	}

	return false
}

// NumberOfShards -
func (g *GenesisNodesSetupHandlerStub) NumberOfShards() uint32 {
	if g.NumberOfShardsCalled != nil {
		return g.NumberOfShardsCalled()
	}

	return 2
}

// MinNumberOfNodes -
func (g *GenesisNodesSetupHandlerStub) MinNumberOfNodes() uint32 {
	if g.MinNumberOfNodesCalled != nil {
		return g.MinNumberOfNodesCalled()
	}

	return 1
}

// MinNumberOfNodesWithHysteresis -
func (g *GenesisNodesSetupHandlerStub) MinNumberOfNodesWithHysteresis() uint32 {
	if g.MinNumberOfNodesWithHysteresisCalled != nil {
		return g.MinNumberOfNodesWithHysteresisCalled()
	}
	return g.MinNumberOfNodes()
}

// AllInitialNodes -
func (g *GenesisNodesSetupHandlerStub) AllInitialNodes() []sharding.GenesisNodeInfoHandler {
	if g.AllInitialNodesCalled != nil {
		return g.AllInitialNodesCalled()
	}
	return nil
}

// IsInterfaceNil -
func (g *GenesisNodesSetupHandlerStub) IsInterfaceNil() bool {
	return g == nil
}