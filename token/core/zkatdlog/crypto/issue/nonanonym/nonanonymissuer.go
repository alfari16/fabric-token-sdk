/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package nonanonym

import (
	api2 "github.com/hyperledger-labs/fabric-smart-client/platform/view/api"
	"github.com/hyperledger-labs/fabric-token-sdk/token/core/zkatdlog/crypto"
	"github.com/hyperledger-labs/fabric-token-sdk/token/core/zkatdlog/crypto/common"
	issue2 "github.com/hyperledger-labs/fabric-token-sdk/token/core/zkatdlog/crypto/issue"
	"github.com/hyperledger-labs/fabric-token-sdk/token/core/zkatdlog/crypto/token"
	"github.com/pkg/errors"
)

//go:generate counterfeiter -o mock/signing_identity.go -fake-name SigningIdentity . SigningIdentity

// signing identity
type SigningIdentity interface {
	api2.SigningIdentity
}

//go:generate counterfeiter -o mock/identity.go -fake-name Identity . Identity

// identity
type Identity interface {
	api2.Identity
}

type Issuer struct {
	Signer       SigningIdentity
	PublicParams *crypto.PublicParams
	Type         string
}

func (i *Issuer) New(ttype string, signer common.SigningIdentity, pp *crypto.PublicParams) {
	i.Signer = signer
	i.Type = ttype
	i.PublicParams = pp
}

func (i *Issuer) GenerateZKIssue(values []uint64, owners [][]byte) (*issue2.IssueAction, []*token.TokenInformation, error) {
	tokens, tw, err := token.GetTokensWithWitness(values, i.Type, i.PublicParams.ZKATPedParams)
	if err != nil {
		return nil, nil, err
	}

	prover := issue2.NewProver(tw, tokens, false, i.PublicParams)
	proof, err := prover.Prove()
	if err != nil {
		return nil, nil, errors.Errorf("failed to generate zero knwoledge proof for issue")
	}

	issue, err := issue2.NewIssue(i.Signer.GetPublicVersion(), tokens, owners, proof, false)
	if err != nil {
		return nil, nil, err
	}

	signerRaw, err := i.Signer.GetPublicVersion().Serialize()
	if err != nil {
		return nil, nil, err
	}

	inf := make([]*token.TokenInformation, len(values))
	for j := 0; j < len(inf); j++ {
		inf[j] = &token.TokenInformation{
			Type:           i.Type,
			Value:          tw[j].Value,
			BlindingFactor: tw[j].BlindingFactor,
			Owner:          owners[j],
			Issuer:         signerRaw,
		}
	}

	return issue, inf, nil
}

func (i *Issuer) SignTokenActions(raw []byte, txID string) ([]byte, error) {
	return i.Signer.Sign(append(raw, []byte(txID)...))
}