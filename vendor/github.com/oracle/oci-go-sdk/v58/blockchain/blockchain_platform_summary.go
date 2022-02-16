// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BlockchainPlatformSummary Blockchain Platform Instance Summary.
type BlockchainPlatformSummary struct {

	// unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Platform Instance Display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Compute shape - STANDARD or ENTERPRISE_SMALL or ENTERPRISE_MEDIUM or ENTERPRISE_LARGE or ENTERPRISE_EXTRA_LARGE or ENTERPRISE_CUSTOM
	ComputeShape BlockchainPlatformComputeShapeEnum `mandatory:"true" json:"computeShape"`

	// Platform Instance Description
	Description *string `mandatory:"false" json:"description"`

	// The time the the Platform Instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Platform Instance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Role of platform - founder or participant
	PlatformRole BlockchainPlatformPlatformRoleEnum `mandatory:"false" json:"platformRole,omitempty"`

	// Service endpoint URL, valid post-provisioning
	ServiceEndpoint *string `mandatory:"false" json:"serviceEndpoint"`

	// The current state of the Platform Instance.
	LifecycleState BlockchainPlatformLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BlockchainPlatformSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BlockchainPlatformSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBlockchainPlatformComputeShapeEnum(string(m.ComputeShape)); !ok && m.ComputeShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeShape: %s. Supported values are: %s.", m.ComputeShape, strings.Join(GetBlockchainPlatformComputeShapeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBlockchainPlatformPlatformRoleEnum(string(m.PlatformRole)); !ok && m.PlatformRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformRole: %s. Supported values are: %s.", m.PlatformRole, strings.Join(GetBlockchainPlatformPlatformRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBlockchainPlatformLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBlockchainPlatformLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
