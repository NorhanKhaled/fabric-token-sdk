/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package driver

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type IdentityUsage int

const (
	IssuerRole = iota
	AuditorRole
	OwnerRole
	CertifierRole
)

type GetIdentityFunc func() (view.Identity, error)

type IdentityInfo struct {
	ID           string
	EnrollmentID string
	GetIdentity  GetIdentityFunc
}

// IdentityProvider handles the long-term identities on top of which wallets are defined.
type IdentityProvider interface {
	LookupIdentifier(usage IdentityUsage, v interface{}) (view.Identity, string)

	// GetIdentityInfo returns the long-term identity info associated to the passed id, nil if not found.
	GetIdentityInfo(usage IdentityUsage, id string) *IdentityInfo

	// GetAuditInfo returns the audit information associated to the passed identity, nil otherwise
	GetAuditInfo(identity view.Identity) ([]byte, error)

	// GetSigner returns a Signer for passed identity.
	GetSigner(identity view.Identity) (Signer, error)

	GetEnrollmentID(auditInfo []byte) (string, error)

	GetIdentityMetadata(identity view.Identity) ([]byte, error)

	// Bind binds id to the passed identity long term identity. The same signer, verifier, and audit of the long term
	// identity is associated to id.
	Bind(id view.Identity, longTerm view.Identity) error
}
