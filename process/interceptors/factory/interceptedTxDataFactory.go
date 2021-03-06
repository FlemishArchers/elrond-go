package factory

import (
	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/atomic"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/core/versioning"
	"github.com/ElrondNetwork/elrond-go/crypto"
	"github.com/ElrondNetwork/elrond-go/hashing"
	"github.com/ElrondNetwork/elrond-go/marshal"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/process/transaction"
	"github.com/ElrondNetwork/elrond-go/sharding"
)

var _ process.InterceptedDataFactory = (*interceptedTxDataFactory)(nil)

var log = logger.GetOrCreate("interceptors/factory")

type interceptedTxDataFactory struct {
	protoMarshalizer            marshal.Marshalizer
	signMarshalizer             marshal.Marshalizer
	hasher                      hashing.Hasher
	keyGen                      crypto.KeyGenerator
	singleSigner                crypto.SingleSigner
	pubkeyConverter             core.PubkeyConverter
	shardCoordinator            sharding.Coordinator
	feeHandler                  process.FeeHandler
	whiteListerVerifiedTxs      process.WhiteListHandler
	argsParser                  process.ArgumentsParser
	chainID                     []byte
	minTransactionVersion       uint32
	enableSignedTxWithHashEpoch uint32
	epochStartTrigger           process.EpochStartTriggerHandler
	txSignHasher                hashing.Hasher
	txVersionChecker            process.TxVersionCheckerHandler
	flagEnableSignedTxWithHash  atomic.Flag
}

// NewInterceptedTxDataFactory creates an instance of interceptedTxDataFactory
func NewInterceptedTxDataFactory(argument *ArgInterceptedDataFactory) (*interceptedTxDataFactory, error) {
	if argument == nil {
		return nil, process.ErrNilArgumentStruct
	}
	if check.IfNil(argument.ProtoMarshalizer) {
		return nil, process.ErrNilMarshalizer
	}
	if check.IfNil(argument.TxSignMarshalizer) {
		return nil, process.ErrNilMarshalizer
	}
	if check.IfNil(argument.Hasher) {
		return nil, process.ErrNilHasher
	}
	if check.IfNil(argument.KeyGen) {
		return nil, process.ErrNilKeyGen
	}
	if check.IfNil(argument.Signer) {
		return nil, process.ErrNilSingleSigner
	}
	if check.IfNil(argument.AddressPubkeyConv) {
		return nil, process.ErrNilPubkeyConverter
	}
	if check.IfNil(argument.ShardCoordinator) {
		return nil, process.ErrNilShardCoordinator
	}
	if check.IfNil(argument.FeeHandler) {
		return nil, process.ErrNilEconomicsFeeHandler
	}
	if check.IfNil(argument.WhiteListerVerifiedTxs) {
		return nil, process.ErrNilWhiteListHandler
	}
	if check.IfNil(argument.ArgsParser) {
		return nil, process.ErrNilArgumentParser
	}
	if len(argument.ChainID) == 0 {
		return nil, process.ErrInvalidChainID
	}
	if argument.MinTransactionVersion == 0 {
		return nil, process.ErrInvalidTransactionVersion
	}
	if check.IfNil(argument.EpochStartTrigger) {
		return nil, process.ErrNilEpochStartTrigger
	}
	if check.IfNil(argument.TxSignHasher) {
		return nil, process.ErrNilHasher
	}
	if check.IfNil(argument.EpochNotifier) {
		return nil, process.ErrNilEpochNotifier
	}

	itdf := &interceptedTxDataFactory{
		protoMarshalizer:            argument.ProtoMarshalizer,
		signMarshalizer:             argument.TxSignMarshalizer,
		hasher:                      argument.Hasher,
		keyGen:                      argument.KeyGen,
		singleSigner:                argument.Signer,
		pubkeyConverter:             argument.AddressPubkeyConv,
		shardCoordinator:            argument.ShardCoordinator,
		feeHandler:                  argument.FeeHandler,
		whiteListerVerifiedTxs:      argument.WhiteListerVerifiedTxs,
		argsParser:                  argument.ArgsParser,
		chainID:                     argument.ChainID,
		minTransactionVersion:       argument.MinTransactionVersion,
		epochStartTrigger:           argument.EpochStartTrigger,
		enableSignedTxWithHashEpoch: argument.EnableSignTxWithHashEpoch,
		txSignHasher:                argument.TxSignHasher,
		txVersionChecker:            versioning.NewTxVersionChecker(argument.MinTransactionVersion),
	}

	argument.EpochNotifier.RegisterNotifyHandler(itdf)

	return itdf, nil
}

// Create creates instances of InterceptedData by unmarshalling provided buffer
func (itdf *interceptedTxDataFactory) Create(buff []byte) (process.InterceptedData, error) {
	return transaction.NewInterceptedTransaction(
		buff,
		itdf.protoMarshalizer,
		itdf.signMarshalizer,
		itdf.hasher,
		itdf.keyGen,
		itdf.singleSigner,
		itdf.pubkeyConverter,
		itdf.shardCoordinator,
		itdf.feeHandler,
		itdf.whiteListerVerifiedTxs,
		itdf.argsParser,
		itdf.chainID,
		itdf.flagEnableSignedTxWithHash.IsSet(),
		itdf.txSignHasher,
		itdf.txVersionChecker,
	)
}

// IsInterfaceNil returns true if there is no value under the interface
func (itdf *interceptedTxDataFactory) IsInterfaceNil() bool {
	return itdf == nil
}

// EpochConfirmed is called whenever a new epoch is confirmed
func (itdf *interceptedTxDataFactory) EpochConfirmed(epoch uint32) {
	itdf.flagEnableSignedTxWithHash.Toggle(epoch >= itdf.enableSignedTxWithHashEpoch)
	log.Debug("interceptors: transaction signed with hash", "enabled", itdf.flagEnableSignedTxWithHash.IsSet())
}
