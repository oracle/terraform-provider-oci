// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateBlockchainPlatformDetails Blockchain Platform details for creating a new service.
type CreateBlockchainPlatformDetails struct {

	// Platform Instance Display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Role of platform - founder or participant
	PlatformRole BlockchainPlatformPlatformRoleEnum `mandatory:"true" json:"platformRole"`

	// Compute shape - STANDARD or ENTERPRISE_SMALL or ENTERPRISE_MEDIUM or ENTERPRISE_LARGE or ENTERPRISE_EXTRA_LARGE
	ComputeShape BlockchainPlatformComputeShapeEnum `mandatory:"true" json:"computeShape"`

	// IDCS access token with Identity Domain Administrator role
	IdcsAccessToken *string `mandatory:"true" json:"idcsAccessToken"`

	// Platform Instance Description
	Description *string `mandatory:"false" json:"description"`

	// Bring your own license
	IsByol *bool `mandatory:"false" json:"isByol"`

	// Platform version
	PlatformVersion *string `mandatory:"false" json:"platformVersion"`

	// Identifier for a federated user
	FederatedUserId *string `mandatory:"false" json:"federatedUserId"`

	// Base64 encoded text in ASCII character set of a Thirdparty CA Certificates archive file.
	// The Archive file is a zip file containing third part CA Certificates,
	// the ca key and certificate files used when issuing enrollment certificates
	// (ECerts) and transaction certificates (TCerts). The chainfile (if it exists)
	// contains the certificate chain which should be trusted for this CA, where
	// the 1st in the chain is always the root CA certificate.
	// File list in zip file [ca-cert.pem,ca-key.pem,ca-chain.pem(optional)].
	CaCertArchiveText *string `mandatory:"false" json:"caCertArchiveText"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBlockchainPlatformDetails) String() string {
	return common.PointerString(m)
}
