// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOccmDemandSignalDetails Details about different fields that can be used to update the demand signal.
type UpdateOccmDemandSignalDetails struct {

	// Use this field to update the display name of the demand signal
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Use this field to update the description of the demand signal.
	Description *string `mandatory:"false" json:"description"`

	// The subset of demand signal states available for updating the demand signal.
	LifecycleDetails UpdateOccmDemandSignalDetailsLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m UpdateOccmDemandSignalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOccmDemandSignalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateOccmDemandSignalDetailsLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetUpdateOccmDemandSignalDetailsLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateOccmDemandSignalDetailsLifecycleDetailsEnum Enum with underlying type: string
type UpdateOccmDemandSignalDetailsLifecycleDetailsEnum string

// Set of constants representing the allowable values for UpdateOccmDemandSignalDetailsLifecycleDetailsEnum
const (
	UpdateOccmDemandSignalDetailsLifecycleDetailsSubmitted UpdateOccmDemandSignalDetailsLifecycleDetailsEnum = "SUBMITTED"
)

var mappingUpdateOccmDemandSignalDetailsLifecycleDetailsEnum = map[string]UpdateOccmDemandSignalDetailsLifecycleDetailsEnum{
	"SUBMITTED": UpdateOccmDemandSignalDetailsLifecycleDetailsSubmitted,
}

var mappingUpdateOccmDemandSignalDetailsLifecycleDetailsEnumLowerCase = map[string]UpdateOccmDemandSignalDetailsLifecycleDetailsEnum{
	"submitted": UpdateOccmDemandSignalDetailsLifecycleDetailsSubmitted,
}

// GetUpdateOccmDemandSignalDetailsLifecycleDetailsEnumValues Enumerates the set of values for UpdateOccmDemandSignalDetailsLifecycleDetailsEnum
func GetUpdateOccmDemandSignalDetailsLifecycleDetailsEnumValues() []UpdateOccmDemandSignalDetailsLifecycleDetailsEnum {
	values := make([]UpdateOccmDemandSignalDetailsLifecycleDetailsEnum, 0)
	for _, v := range mappingUpdateOccmDemandSignalDetailsLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOccmDemandSignalDetailsLifecycleDetailsEnumStringValues Enumerates the set of values in String for UpdateOccmDemandSignalDetailsLifecycleDetailsEnum
func GetUpdateOccmDemandSignalDetailsLifecycleDetailsEnumStringValues() []string {
	return []string{
		"SUBMITTED",
	}
}

// GetMappingUpdateOccmDemandSignalDetailsLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOccmDemandSignalDetailsLifecycleDetailsEnum(val string) (UpdateOccmDemandSignalDetailsLifecycleDetailsEnum, bool) {
	enum, ok := mappingUpdateOccmDemandSignalDetailsLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
