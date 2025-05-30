// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCaptureFilterDetails A capture filter contains a set of rules governing what traffic a VTAP mirrors.
type CreateCaptureFilterDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capture filter.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Indicates which service will use this capture filter
	FilterType CreateCaptureFilterDetailsFilterTypeEnum `mandatory:"true" json:"filterType"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The set of rules governing what traffic a VTAP mirrors.
	VtapCaptureFilterRules []VtapCaptureFilterRuleDetails `mandatory:"false" json:"vtapCaptureFilterRules"`

	// The set of rules governing what traffic the Flow Log collects when creating a flow log capture filter.
	FlowLogCaptureFilterRules []FlowLogCaptureFilterRuleDetails `mandatory:"false" json:"flowLogCaptureFilterRules"`
}

func (m CreateCaptureFilterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCaptureFilterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateCaptureFilterDetailsFilterTypeEnum(string(m.FilterType)); !ok && m.FilterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterType: %s. Supported values are: %s.", m.FilterType, strings.Join(GetCreateCaptureFilterDetailsFilterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateCaptureFilterDetailsFilterTypeEnum Enum with underlying type: string
type CreateCaptureFilterDetailsFilterTypeEnum string

// Set of constants representing the allowable values for CreateCaptureFilterDetailsFilterTypeEnum
const (
	CreateCaptureFilterDetailsFilterTypeVtap    CreateCaptureFilterDetailsFilterTypeEnum = "VTAP"
	CreateCaptureFilterDetailsFilterTypeFlowlog CreateCaptureFilterDetailsFilterTypeEnum = "FLOWLOG"
)

var mappingCreateCaptureFilterDetailsFilterTypeEnum = map[string]CreateCaptureFilterDetailsFilterTypeEnum{
	"VTAP":    CreateCaptureFilterDetailsFilterTypeVtap,
	"FLOWLOG": CreateCaptureFilterDetailsFilterTypeFlowlog,
}

var mappingCreateCaptureFilterDetailsFilterTypeEnumLowerCase = map[string]CreateCaptureFilterDetailsFilterTypeEnum{
	"vtap":    CreateCaptureFilterDetailsFilterTypeVtap,
	"flowlog": CreateCaptureFilterDetailsFilterTypeFlowlog,
}

// GetCreateCaptureFilterDetailsFilterTypeEnumValues Enumerates the set of values for CreateCaptureFilterDetailsFilterTypeEnum
func GetCreateCaptureFilterDetailsFilterTypeEnumValues() []CreateCaptureFilterDetailsFilterTypeEnum {
	values := make([]CreateCaptureFilterDetailsFilterTypeEnum, 0)
	for _, v := range mappingCreateCaptureFilterDetailsFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCaptureFilterDetailsFilterTypeEnumStringValues Enumerates the set of values in String for CreateCaptureFilterDetailsFilterTypeEnum
func GetCreateCaptureFilterDetailsFilterTypeEnumStringValues() []string {
	return []string{
		"VTAP",
		"FLOWLOG",
	}
}

// GetMappingCreateCaptureFilterDetailsFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCaptureFilterDetailsFilterTypeEnum(val string) (CreateCaptureFilterDetailsFilterTypeEnum, bool) {
	enum, ok := mappingCreateCaptureFilterDetailsFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
