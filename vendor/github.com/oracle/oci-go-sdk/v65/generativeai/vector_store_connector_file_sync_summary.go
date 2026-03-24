// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VectorStoreConnectorFileSyncSummary Summary information for a VectorStoreConnectorFileSync
type VectorStoreConnectorFileSyncSummary struct {

	// An OCID that uniquely identifies a VectorStoreConnectorFileSync operation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// An OCID that identifies the VectorStoreConnector under which this FileSync operation is created.
	VectorStoreConnectorId *string `mandatory:"true" json:"vectorStoreConnectorId"`

	// Owning compartment OCID for a VectorStoreConnectorFileSync.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Owning tenant OCID for a VectorStoreConnector
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The date and time that the FileSync operation was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the VectorStoreConnectorFileSync operation.
	// Accepted Values are:
	//   - ACCEPTED
	//   - IN_PROGRESS
	//   - FAILED
	//   - SUCCEEDED
	//   - CANCELING
	//   - CANCELED
	LifecycleState VectorStoreConnectorFileSyncLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of the FileSync operation based on how it is triggered. The type can be either MANUAL or SCHEDULED
	TriggerType VectorStoreConnectorFileSyncSummaryTriggerTypeEnum `mandatory:"true" json:"triggerType"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the FileSync operation has started in the format of an RFC3339 datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time when the FileSync operation has ended in the format of an RFC3339 datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The duration (in seconds) taken for the VectorStoreConnectorFileSync operation to complete.
	DurationInSeconds *int `mandatory:"false" json:"durationInSeconds"`

	Stats *VectorStoreConnectorStats `mandatory:"false" json:"stats"`
}

func (m VectorStoreConnectorFileSyncSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VectorStoreConnectorFileSyncSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVectorStoreConnectorFileSyncLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVectorStoreConnectorFileSyncLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVectorStoreConnectorFileSyncSummaryTriggerTypeEnum(string(m.TriggerType)); !ok && m.TriggerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggerType: %s. Supported values are: %s.", m.TriggerType, strings.Join(GetVectorStoreConnectorFileSyncSummaryTriggerTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VectorStoreConnectorFileSyncSummaryTriggerTypeEnum Enum with underlying type: string
type VectorStoreConnectorFileSyncSummaryTriggerTypeEnum string

// Set of constants representing the allowable values for VectorStoreConnectorFileSyncSummaryTriggerTypeEnum
const (
	VectorStoreConnectorFileSyncSummaryTriggerTypeManual    VectorStoreConnectorFileSyncSummaryTriggerTypeEnum = "MANUAL"
	VectorStoreConnectorFileSyncSummaryTriggerTypeScheduled VectorStoreConnectorFileSyncSummaryTriggerTypeEnum = "SCHEDULED"
)

var mappingVectorStoreConnectorFileSyncSummaryTriggerTypeEnum = map[string]VectorStoreConnectorFileSyncSummaryTriggerTypeEnum{
	"MANUAL":    VectorStoreConnectorFileSyncSummaryTriggerTypeManual,
	"SCHEDULED": VectorStoreConnectorFileSyncSummaryTriggerTypeScheduled,
}

var mappingVectorStoreConnectorFileSyncSummaryTriggerTypeEnumLowerCase = map[string]VectorStoreConnectorFileSyncSummaryTriggerTypeEnum{
	"manual":    VectorStoreConnectorFileSyncSummaryTriggerTypeManual,
	"scheduled": VectorStoreConnectorFileSyncSummaryTriggerTypeScheduled,
}

// GetVectorStoreConnectorFileSyncSummaryTriggerTypeEnumValues Enumerates the set of values for VectorStoreConnectorFileSyncSummaryTriggerTypeEnum
func GetVectorStoreConnectorFileSyncSummaryTriggerTypeEnumValues() []VectorStoreConnectorFileSyncSummaryTriggerTypeEnum {
	values := make([]VectorStoreConnectorFileSyncSummaryTriggerTypeEnum, 0)
	for _, v := range mappingVectorStoreConnectorFileSyncSummaryTriggerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVectorStoreConnectorFileSyncSummaryTriggerTypeEnumStringValues Enumerates the set of values in String for VectorStoreConnectorFileSyncSummaryTriggerTypeEnum
func GetVectorStoreConnectorFileSyncSummaryTriggerTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"SCHEDULED",
	}
}

// GetMappingVectorStoreConnectorFileSyncSummaryTriggerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVectorStoreConnectorFileSyncSummaryTriggerTypeEnum(val string) (VectorStoreConnectorFileSyncSummaryTriggerTypeEnum, bool) {
	enum, ok := mappingVectorStoreConnectorFileSyncSummaryTriggerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
