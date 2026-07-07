// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TransferStatus Detailed transfer status record for CDS pipeline file tracking.
type TransferStatus struct {

	// ID of the transfer job emitted by the data plane.
	TransferId *string `mandatory:"true" json:"transferId"`

	// Name of the transferred object.
	Filename *string `mandatory:"true" json:"filename"`

	// Guarded Data Pipeline OCID that initiated the transfer.
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// Compartment OCID that owns the transfer status.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Source region of the transfer status.
	SourceRegion *string `mandatory:"true" json:"sourceRegion"`

	// Destination region where the transfer completed.
	DestinationRegion *string `mandatory:"true" json:"destinationRegion"`

	// Current lifecycle state of the transfer.
	LifecycleState TransferStatusLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Domain-specific transfer phase reported by the data plane.
	TransferPhase TransferStatusTransferPhaseEnum `mandatory:"true" json:"transferPhase"`

	// Indicates whether the transfer completed successfully.
	IsTransferComplete *bool `mandatory:"true" json:"isTransferComplete"`

	// The time the transfer was created in the data plane.
	TimeUploaded *common.SDKTime `mandatory:"true" json:"timeUploaded"`

	// The time the transfer status was last updated.
	TimeLastUpdated *common.SDKTime `mandatory:"true" json:"timeLastUpdated"`

	// Timeline of state changes keyed by destination region.
	StateTransitions map[string][]StateDetailSummary `mandatory:"true" json:"stateTransitions"`

	// Hex-encoded hash of the transferred content.
	ContentHash *string `mandatory:"false" json:"contentHash"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TransferStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTransferStatusLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferStatusLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTransferStatusTransferPhaseEnum(string(m.TransferPhase)); !ok && m.TransferPhase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransferPhase: %s. Supported values are: %s.", m.TransferPhase, strings.Join(GetTransferStatusTransferPhaseEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferStatusLifecycleStateEnum Enum with underlying type: string
type TransferStatusLifecycleStateEnum string

// Set of constants representing the allowable values for TransferStatusLifecycleStateEnum
const (
	TransferStatusLifecycleStateCreating       TransferStatusLifecycleStateEnum = "CREATING"
	TransferStatusLifecycleStateUpdating       TransferStatusLifecycleStateEnum = "UPDATING"
	TransferStatusLifecycleStateActive         TransferStatusLifecycleStateEnum = "ACTIVE"
	TransferStatusLifecycleStateInactive       TransferStatusLifecycleStateEnum = "INACTIVE"
	TransferStatusLifecycleStateDeleting       TransferStatusLifecycleStateEnum = "DELETING"
	TransferStatusLifecycleStateDeleted        TransferStatusLifecycleStateEnum = "DELETED"
	TransferStatusLifecycleStateFailed         TransferStatusLifecycleStateEnum = "FAILED"
	TransferStatusLifecycleStateNeedsAttention TransferStatusLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingTransferStatusLifecycleStateEnum = map[string]TransferStatusLifecycleStateEnum{
	"CREATING":        TransferStatusLifecycleStateCreating,
	"UPDATING":        TransferStatusLifecycleStateUpdating,
	"ACTIVE":          TransferStatusLifecycleStateActive,
	"INACTIVE":        TransferStatusLifecycleStateInactive,
	"DELETING":        TransferStatusLifecycleStateDeleting,
	"DELETED":         TransferStatusLifecycleStateDeleted,
	"FAILED":          TransferStatusLifecycleStateFailed,
	"NEEDS_ATTENTION": TransferStatusLifecycleStateNeedsAttention,
}

var mappingTransferStatusLifecycleStateEnumLowerCase = map[string]TransferStatusLifecycleStateEnum{
	"creating":        TransferStatusLifecycleStateCreating,
	"updating":        TransferStatusLifecycleStateUpdating,
	"active":          TransferStatusLifecycleStateActive,
	"inactive":        TransferStatusLifecycleStateInactive,
	"deleting":        TransferStatusLifecycleStateDeleting,
	"deleted":         TransferStatusLifecycleStateDeleted,
	"failed":          TransferStatusLifecycleStateFailed,
	"needs_attention": TransferStatusLifecycleStateNeedsAttention,
}

// GetTransferStatusLifecycleStateEnumValues Enumerates the set of values for TransferStatusLifecycleStateEnum
func GetTransferStatusLifecycleStateEnumValues() []TransferStatusLifecycleStateEnum {
	values := make([]TransferStatusLifecycleStateEnum, 0)
	for _, v := range mappingTransferStatusLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferStatusLifecycleStateEnumStringValues Enumerates the set of values in String for TransferStatusLifecycleStateEnum
func GetTransferStatusLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingTransferStatusLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferStatusLifecycleStateEnum(val string) (TransferStatusLifecycleStateEnum, bool) {
	enum, ok := mappingTransferStatusLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TransferStatusTransferPhaseEnum Enum with underlying type: string
type TransferStatusTransferPhaseEnum string

// Set of constants representing the allowable values for TransferStatusTransferPhaseEnum
const (
	TransferStatusTransferPhaseQueued         TransferStatusTransferPhaseEnum = "QUEUED"
	TransferStatusTransferPhaseTransferring   TransferStatusTransferPhaseEnum = "TRANSFERRING"
	TransferStatusTransferPhaseCompleted      TransferStatusTransferPhaseEnum = "COMPLETED"
	TransferStatusTransferPhaseFailed         TransferStatusTransferPhaseEnum = "FAILED"
	TransferStatusTransferPhaseNeedsAttention TransferStatusTransferPhaseEnum = "NEEDS_ATTENTION"
)

var mappingTransferStatusTransferPhaseEnum = map[string]TransferStatusTransferPhaseEnum{
	"QUEUED":          TransferStatusTransferPhaseQueued,
	"TRANSFERRING":    TransferStatusTransferPhaseTransferring,
	"COMPLETED":       TransferStatusTransferPhaseCompleted,
	"FAILED":          TransferStatusTransferPhaseFailed,
	"NEEDS_ATTENTION": TransferStatusTransferPhaseNeedsAttention,
}

var mappingTransferStatusTransferPhaseEnumLowerCase = map[string]TransferStatusTransferPhaseEnum{
	"queued":          TransferStatusTransferPhaseQueued,
	"transferring":    TransferStatusTransferPhaseTransferring,
	"completed":       TransferStatusTransferPhaseCompleted,
	"failed":          TransferStatusTransferPhaseFailed,
	"needs_attention": TransferStatusTransferPhaseNeedsAttention,
}

// GetTransferStatusTransferPhaseEnumValues Enumerates the set of values for TransferStatusTransferPhaseEnum
func GetTransferStatusTransferPhaseEnumValues() []TransferStatusTransferPhaseEnum {
	values := make([]TransferStatusTransferPhaseEnum, 0)
	for _, v := range mappingTransferStatusTransferPhaseEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferStatusTransferPhaseEnumStringValues Enumerates the set of values in String for TransferStatusTransferPhaseEnum
func GetTransferStatusTransferPhaseEnumStringValues() []string {
	return []string{
		"QUEUED",
		"TRANSFERRING",
		"COMPLETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingTransferStatusTransferPhaseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferStatusTransferPhaseEnum(val string) (TransferStatusTransferPhaseEnum, bool) {
	enum, ok := mappingTransferStatusTransferPhaseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
