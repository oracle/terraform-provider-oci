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

// CreateOdaInstanceAttachmentDetails Properties required to create an ODA instance attachment.
type CreateOdaInstanceAttachmentDetails struct {

	// The OCID of the target instance (which could be any other OCI PaaS/SaaS resource), to which this ODA instance is being attached.
	AttachToId *string `mandatory:"true" json:"attachToId"`

	// The type of target instance which this ODA instance is being attached.
	AttachmentType CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`

	Owner *OdaInstanceAttachmentOwner `mandatory:"true" json:"owner"`

	// Attachment specific metadata. Defined by the target service.
	AttachmentMetadata *string `mandatory:"false" json:"attachmentMetadata"`

	// List of operations that are restricted while this instance is attached.
	RestrictedOperations []string `mandatory:"false" json:"restrictedOperations"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOdaInstanceAttachmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOdaInstanceAttachmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateOdaInstanceAttachmentDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetCreateOdaInstanceAttachmentDetailsAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum Enum with underlying type: string
type CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum
const (
	CreateOdaInstanceAttachmentDetailsAttachmentTypeFusion CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum = "FUSION"
)

var mappingCreateOdaInstanceAttachmentDetailsAttachmentTypeEnum = map[string]CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum{
	"FUSION": CreateOdaInstanceAttachmentDetailsAttachmentTypeFusion,
}

var mappingCreateOdaInstanceAttachmentDetailsAttachmentTypeEnumLowerCase = map[string]CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum{
	"fusion": CreateOdaInstanceAttachmentDetailsAttachmentTypeFusion,
}

// GetCreateOdaInstanceAttachmentDetailsAttachmentTypeEnumValues Enumerates the set of values for CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum
func GetCreateOdaInstanceAttachmentDetailsAttachmentTypeEnumValues() []CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum {
	values := make([]CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingCreateOdaInstanceAttachmentDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOdaInstanceAttachmentDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum
func GetCreateOdaInstanceAttachmentDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"FUSION",
	}
}

// GetMappingCreateOdaInstanceAttachmentDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOdaInstanceAttachmentDetailsAttachmentTypeEnum(val string) (CreateOdaInstanceAttachmentDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingCreateOdaInstanceAttachmentDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
