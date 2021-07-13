/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package view

import (
	view2 "github.com/hyperledger-labs/fabric-smart-client/platform/view"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"

	"github.com/hyperledger-labs/fabric-token-sdk/token/driver"
)

type SigService interface {
	GetVerifier(id view.Identity) (view2.Verifier, error)
	GetSigner(id view.Identity) (view2.Signer, error)
	RegisterSigner(identity view.Identity, signer view2.Signer, verifier view2.Verifier) error
}

type SigServiceWrapper struct {
	s SigService
}

func NewSigServiceWrapper(s SigService) *SigServiceWrapper {
	return &SigServiceWrapper{s: s}
}

func (s *SigServiceWrapper) GetVerifier(id view.Identity) (driver.Verifier, error) {
	return s.s.GetVerifier(id)
}

func (s *SigServiceWrapper) GetSigner(id view.Identity) (driver.Signer, error) {
	return s.s.GetSigner(id)
}

func (s *SigServiceWrapper) RegisterSigner(identity view.Identity, signer driver.Signer, verifier driver.Verifier) error {
	return s.s.RegisterSigner(identity, signer, verifier)
}
