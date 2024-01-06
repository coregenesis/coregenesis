package server

import (
	"github.com/coregenesis/coregenesis/consensus"
	consensusDev "github.com/coregenesis/coregenesis/consensus/dev"
	consensusDummy "github.com/coregenesis/coregenesis/consensus/dummy"
	consensusIBFT "github.com/coregenesis/coregenesis/consensus/ibft"
	"github.com/coregenesis/coregenesis/secrets"
	"github.com/coregenesis/coregenesis/secrets/awsssm"
	"github.com/coregenesis/coregenesis/secrets/gcpssm"
	"github.com/coregenesis/coregenesis/secrets/hashicorpvault"
	"github.com/coregenesis/coregenesis/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
