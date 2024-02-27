// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OdaInstanceAttachmentSummary Description of an ODA instance attachment.
type OdaInstanceAttachmentSummary struct {

	// Unique immutable identifier that was assigned when the ODA instance attachment was created.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the ODA instance to which the attachment applies.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The OCID of the target instance (which could be any other OCI PaaS/SaaS resource), to which the ODA instance is or is being attached.
	AttachToId *string `mandatory:"true" json:"attachToId"`

	// The type of attachment defined as an enum.
	AttachmentType OdaInstanceAttachmentSummaryAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`

	// The current state of the attachment.
	LifecycleState OdaInstanceAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Attachment-specific metadata, defined by the target service.
	AttachmentMetadata *string `mandatory:"false" json:"attachmentMetadata"`

	// List of operation names that are restricted while this ODA instance is attached.
	RestrictedOperations []string `mandatory:"false" json:"restrictedOperations"`

	Owner *OdaInstanceOwner `mandatory:"false" json:"owner"`

	// The time the attachment was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the attachment was last modified. An RFC3339 formatted datetime string
	TimeLastUpdate *common.SDKTime `mandatory:"false" json:"timeLastUpdate"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OdaInstanceAttachmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OdaInstanceAttachmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaInstanceAttachmentSummaryAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetOdaInstanceAttachmentSummaryAttachmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOdaInstanceAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOdaInstanceAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdaInstanceAttachmentSummaryAttachmentTypeEnum Enum with underlying type: string
type OdaInstanceAttachmentSummaryAttachmentTypeEnum string

// Set of constants representing the allowable values for OdaInstanceAttachmentSummaryAttachmentTypeEnum
const (
	OdaInstanceAttachmentSummaryAttachmentTypeFusion OdaInstanceAttachmentSummaryAttachmentTypeEnum = "FUSION"
	OdaInstanceAttachmentSummaryAttachmentTypeMax    OdaInstanceAttachmentSummaryAttachmentTypeEnum = "MAX"
)

var mappingOdaInstanceAttachmentSummaryAttachmentTypeEnum = map[string]OdaInstanceAttachmentSummaryAttachmentTypeEnum{
	"FUSION": OdaInstanceAttachmentSummaryAttachmentTypeFusion,
	"MAX":    OdaInstanceAttachmentSummaryAttachmentTypeMax,
}

var mappingOdaInstanceAttachmentSummaryAttachmentTypeEnumLowerCase = map[string]OdaInstanceAttachmentSummaryAttachmentTypeEnum{
	"fusion": OdaInstanceAttachmentSummaryAttachmentTypeFusion,
	"max":    OdaInstanceAttachmentSummaryAttachmentTypeMax,
}

// GetOdaInstanceAttachmentSummaryAttachmentTypeEnumValues Enumerates the set of values for OdaInstanceAttachmentSummaryAttachmentTypeEnum
func GetOdaInstanceAttachmentSummaryAttachmentTypeEnumValues() []OdaInstanceAttachmentSummaryAttachmentTypeEnum {
	values := make([]OdaInstanceAttachmentSummaryAttachmentTypeEnum, 0)
	for _, v := range mappingOdaInstanceAttachmentSummaryAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaInstanceAttachmentSummaryAttachmentTypeEnumStringValues Enumerates the set of values in String for OdaInstanceAttachmentSummaryAttachmentTypeEnum
func GetOdaInstanceAttachmentSummaryAttachmentTypeEnumStringValues() []string {
	return []string{
		"FUSION",
		"MAX",
	}
}

// GetMappingOdaInstanceAttachmentSummaryAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaInstanceAttachmentSummaryAttachmentTypeEnum(val string) (OdaInstanceAttachmentSummaryAttachmentTypeEnum, bool) {
	enum, ok := mappingOdaInstanceAttachmentSummaryAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
