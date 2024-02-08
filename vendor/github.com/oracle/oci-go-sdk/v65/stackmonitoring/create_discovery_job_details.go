// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDiscoveryJobDetails The request of DiscoveryJob details.
type CreateDiscoveryJobDetails struct {

	// The OCID of Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DiscoveryDetails *DiscoveryDetails `mandatory:"true" json:"discoveryDetails"`

	// Add option submits new discovery Job. Add with retry option to re-submit failed discovery job. Refresh option refreshes the existing discovered resources.
	DiscoveryType CreateDiscoveryJobDetailsDiscoveryTypeEnum `mandatory:"false" json:"discoveryType,omitempty"`

	// Client who submits discovery job.
	DiscoveryClient *string `mandatory:"false" json:"discoveryClient"`

	// If this parameter set to true, the specified tags will be applied
	// to all resources discovered in the current request.
	// Default is true.
	ShouldPropagateTagsToDiscoveredResources *bool `mandatory:"false" json:"shouldPropagateTagsToDiscoveredResources"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDiscoveryJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDiscoveryJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDiscoveryJobDetailsDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetCreateDiscoveryJobDetailsDiscoveryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDiscoveryJobDetailsDiscoveryTypeEnum Enum with underlying type: string
type CreateDiscoveryJobDetailsDiscoveryTypeEnum string

// Set of constants representing the allowable values for CreateDiscoveryJobDetailsDiscoveryTypeEnum
const (
	CreateDiscoveryJobDetailsDiscoveryTypeAdd          CreateDiscoveryJobDetailsDiscoveryTypeEnum = "ADD"
	CreateDiscoveryJobDetailsDiscoveryTypeAddWithRetry CreateDiscoveryJobDetailsDiscoveryTypeEnum = "ADD_WITH_RETRY"
	CreateDiscoveryJobDetailsDiscoveryTypeRefresh      CreateDiscoveryJobDetailsDiscoveryTypeEnum = "REFRESH"
)

var mappingCreateDiscoveryJobDetailsDiscoveryTypeEnum = map[string]CreateDiscoveryJobDetailsDiscoveryTypeEnum{
	"ADD":            CreateDiscoveryJobDetailsDiscoveryTypeAdd,
	"ADD_WITH_RETRY": CreateDiscoveryJobDetailsDiscoveryTypeAddWithRetry,
	"REFRESH":        CreateDiscoveryJobDetailsDiscoveryTypeRefresh,
}

var mappingCreateDiscoveryJobDetailsDiscoveryTypeEnumLowerCase = map[string]CreateDiscoveryJobDetailsDiscoveryTypeEnum{
	"add":            CreateDiscoveryJobDetailsDiscoveryTypeAdd,
	"add_with_retry": CreateDiscoveryJobDetailsDiscoveryTypeAddWithRetry,
	"refresh":        CreateDiscoveryJobDetailsDiscoveryTypeRefresh,
}

// GetCreateDiscoveryJobDetailsDiscoveryTypeEnumValues Enumerates the set of values for CreateDiscoveryJobDetailsDiscoveryTypeEnum
func GetCreateDiscoveryJobDetailsDiscoveryTypeEnumValues() []CreateDiscoveryJobDetailsDiscoveryTypeEnum {
	values := make([]CreateDiscoveryJobDetailsDiscoveryTypeEnum, 0)
	for _, v := range mappingCreateDiscoveryJobDetailsDiscoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDiscoveryJobDetailsDiscoveryTypeEnumStringValues Enumerates the set of values in String for CreateDiscoveryJobDetailsDiscoveryTypeEnum
func GetCreateDiscoveryJobDetailsDiscoveryTypeEnumStringValues() []string {
	return []string{
		"ADD",
		"ADD_WITH_RETRY",
		"REFRESH",
	}
}

// GetMappingCreateDiscoveryJobDetailsDiscoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDiscoveryJobDetailsDiscoveryTypeEnum(val string) (CreateDiscoveryJobDetailsDiscoveryTypeEnum, bool) {
	enum, ok := mappingCreateDiscoveryJobDetailsDiscoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
