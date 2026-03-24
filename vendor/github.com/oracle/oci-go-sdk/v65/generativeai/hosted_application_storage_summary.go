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

// HostedApplicationStorageSummary Summary information about a hosted application storage.
type HostedApplicationStorageSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the hosted application storage.
	LifecycleState HostedApplicationStorageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An optional description of the hosted application.
	Description *string `mandatory:"false" json:"description"`

	// The compartment OCID to create the hosted application in.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// type like Cache, Postgresql and ADB.
	StorageType HostedApplicationStorageSummaryStorageTypeEnum `mandatory:"false" json:"storageType,omitempty"`

	// The date and time the hosted application was created, in the format defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the hosted application was updated, in the format defined by RFC 3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m HostedApplicationStorageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostedApplicationStorageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostedApplicationStorageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHostedApplicationStorageLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingHostedApplicationStorageSummaryStorageTypeEnum(string(m.StorageType)); !ok && m.StorageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageType: %s. Supported values are: %s.", m.StorageType, strings.Join(GetHostedApplicationStorageSummaryStorageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostedApplicationStorageSummaryStorageTypeEnum Enum with underlying type: string
type HostedApplicationStorageSummaryStorageTypeEnum string

// Set of constants representing the allowable values for HostedApplicationStorageSummaryStorageTypeEnum
const (
	HostedApplicationStorageSummaryStorageTypeCache      HostedApplicationStorageSummaryStorageTypeEnum = "CACHE"
	HostedApplicationStorageSummaryStorageTypePostgresql HostedApplicationStorageSummaryStorageTypeEnum = "POSTGRESQL"
	HostedApplicationStorageSummaryStorageTypeAdb        HostedApplicationStorageSummaryStorageTypeEnum = "ADB"
)

var mappingHostedApplicationStorageSummaryStorageTypeEnum = map[string]HostedApplicationStorageSummaryStorageTypeEnum{
	"CACHE":      HostedApplicationStorageSummaryStorageTypeCache,
	"POSTGRESQL": HostedApplicationStorageSummaryStorageTypePostgresql,
	"ADB":        HostedApplicationStorageSummaryStorageTypeAdb,
}

var mappingHostedApplicationStorageSummaryStorageTypeEnumLowerCase = map[string]HostedApplicationStorageSummaryStorageTypeEnum{
	"cache":      HostedApplicationStorageSummaryStorageTypeCache,
	"postgresql": HostedApplicationStorageSummaryStorageTypePostgresql,
	"adb":        HostedApplicationStorageSummaryStorageTypeAdb,
}

// GetHostedApplicationStorageSummaryStorageTypeEnumValues Enumerates the set of values for HostedApplicationStorageSummaryStorageTypeEnum
func GetHostedApplicationStorageSummaryStorageTypeEnumValues() []HostedApplicationStorageSummaryStorageTypeEnum {
	values := make([]HostedApplicationStorageSummaryStorageTypeEnum, 0)
	for _, v := range mappingHostedApplicationStorageSummaryStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHostedApplicationStorageSummaryStorageTypeEnumStringValues Enumerates the set of values in String for HostedApplicationStorageSummaryStorageTypeEnum
func GetHostedApplicationStorageSummaryStorageTypeEnumStringValues() []string {
	return []string{
		"CACHE",
		"POSTGRESQL",
		"ADB",
	}
}

// GetMappingHostedApplicationStorageSummaryStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostedApplicationStorageSummaryStorageTypeEnum(val string) (HostedApplicationStorageSummaryStorageTypeEnum, bool) {
	enum, ok := mappingHostedApplicationStorageSummaryStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
