// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBlockchainPlatformDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBlockchainPlatformPlatformRoleEnum(string(m.PlatformRole)); !ok && m.PlatformRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformRole: %s. Supported values are: %s.", m.PlatformRole, strings.Join(GetBlockchainPlatformPlatformRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBlockchainPlatformComputeShapeEnum(string(m.ComputeShape)); !ok && m.ComputeShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeShape: %s. Supported values are: %s.", m.ComputeShape, strings.Join(GetBlockchainPlatformComputeShapeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
