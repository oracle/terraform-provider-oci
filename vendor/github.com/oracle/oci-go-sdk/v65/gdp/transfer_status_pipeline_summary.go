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

// TransferStatusPipelineSummary Summary information for a transfer status listed for a specific pipeline.
type TransferStatusPipelineSummary struct {

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
	TransferPhase TransferStatusPipelineSummaryTransferPhaseEnum `mandatory:"true" json:"transferPhase"`

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

func (m TransferStatusPipelineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferStatusPipelineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTransferStatusLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferStatusLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTransferStatusPipelineSummaryTransferPhaseEnum(string(m.TransferPhase)); !ok && m.TransferPhase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransferPhase: %s. Supported values are: %s.", m.TransferPhase, strings.Join(GetTransferStatusPipelineSummaryTransferPhaseEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferStatusPipelineSummaryTransferPhaseEnum Enum with underlying type: string
type TransferStatusPipelineSummaryTransferPhaseEnum string

// Set of constants representing the allowable values for TransferStatusPipelineSummaryTransferPhaseEnum
const (
	TransferStatusPipelineSummaryTransferPhaseQueued         TransferStatusPipelineSummaryTransferPhaseEnum = "QUEUED"
	TransferStatusPipelineSummaryTransferPhaseTransferring   TransferStatusPipelineSummaryTransferPhaseEnum = "TRANSFERRING"
	TransferStatusPipelineSummaryTransferPhaseCompleted      TransferStatusPipelineSummaryTransferPhaseEnum = "COMPLETED"
	TransferStatusPipelineSummaryTransferPhaseFailed         TransferStatusPipelineSummaryTransferPhaseEnum = "FAILED"
	TransferStatusPipelineSummaryTransferPhaseNeedsAttention TransferStatusPipelineSummaryTransferPhaseEnum = "NEEDS_ATTENTION"
)

var mappingTransferStatusPipelineSummaryTransferPhaseEnum = map[string]TransferStatusPipelineSummaryTransferPhaseEnum{
	"QUEUED":          TransferStatusPipelineSummaryTransferPhaseQueued,
	"TRANSFERRING":    TransferStatusPipelineSummaryTransferPhaseTransferring,
	"COMPLETED":       TransferStatusPipelineSummaryTransferPhaseCompleted,
	"FAILED":          TransferStatusPipelineSummaryTransferPhaseFailed,
	"NEEDS_ATTENTION": TransferStatusPipelineSummaryTransferPhaseNeedsAttention,
}

var mappingTransferStatusPipelineSummaryTransferPhaseEnumLowerCase = map[string]TransferStatusPipelineSummaryTransferPhaseEnum{
	"queued":          TransferStatusPipelineSummaryTransferPhaseQueued,
	"transferring":    TransferStatusPipelineSummaryTransferPhaseTransferring,
	"completed":       TransferStatusPipelineSummaryTransferPhaseCompleted,
	"failed":          TransferStatusPipelineSummaryTransferPhaseFailed,
	"needs_attention": TransferStatusPipelineSummaryTransferPhaseNeedsAttention,
}

// GetTransferStatusPipelineSummaryTransferPhaseEnumValues Enumerates the set of values for TransferStatusPipelineSummaryTransferPhaseEnum
func GetTransferStatusPipelineSummaryTransferPhaseEnumValues() []TransferStatusPipelineSummaryTransferPhaseEnum {
	values := make([]TransferStatusPipelineSummaryTransferPhaseEnum, 0)
	for _, v := range mappingTransferStatusPipelineSummaryTransferPhaseEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferStatusPipelineSummaryTransferPhaseEnumStringValues Enumerates the set of values in String for TransferStatusPipelineSummaryTransferPhaseEnum
func GetTransferStatusPipelineSummaryTransferPhaseEnumStringValues() []string {
	return []string{
		"QUEUED",
		"TRANSFERRING",
		"COMPLETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingTransferStatusPipelineSummaryTransferPhaseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferStatusPipelineSummaryTransferPhaseEnum(val string) (TransferStatusPipelineSummaryTransferPhaseEnum, bool) {
	enum, ok := mappingTransferStatusPipelineSummaryTransferPhaseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
