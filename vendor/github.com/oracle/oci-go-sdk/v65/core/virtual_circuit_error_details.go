// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VirtualCircuitErrorDetails Details for an error on a Virtual Circuit.
type VirtualCircuitErrorDetails struct {

	// Describes the severity level of the error.
	SeverityLevel VirtualCircuitErrorDetailsSeverityLevelEnum `mandatory:"true" json:"severityLevel"`

	// Unique code describes the error type.
	ErrorCode *string `mandatory:"true" json:"errorCode"`

	// A detailed description of the error.
	ErrorDescription *string `mandatory:"true" json:"errorDescription"`

	// Resolution for the error.
	Solution *string `mandatory:"true" json:"solution"`

	// Link to more Oracle resources or relevant documentation.
	OciResourcesLink *string `mandatory:"true" json:"ociResourcesLink"`
}

func (m VirtualCircuitErrorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualCircuitErrorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVirtualCircuitErrorDetailsSeverityLevelEnum(string(m.SeverityLevel)); !ok && m.SeverityLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SeverityLevel: %s. Supported values are: %s.", m.SeverityLevel, strings.Join(GetVirtualCircuitErrorDetailsSeverityLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VirtualCircuitErrorDetailsSeverityLevelEnum Enum with underlying type: string
type VirtualCircuitErrorDetailsSeverityLevelEnum string

// Set of constants representing the allowable values for VirtualCircuitErrorDetailsSeverityLevelEnum
const (
	VirtualCircuitErrorDetailsSeverityLevelError   VirtualCircuitErrorDetailsSeverityLevelEnum = "ERROR"
	VirtualCircuitErrorDetailsSeverityLevelWarning VirtualCircuitErrorDetailsSeverityLevelEnum = "WARNING"
)

var mappingVirtualCircuitErrorDetailsSeverityLevelEnum = map[string]VirtualCircuitErrorDetailsSeverityLevelEnum{
	"ERROR":   VirtualCircuitErrorDetailsSeverityLevelError,
	"WARNING": VirtualCircuitErrorDetailsSeverityLevelWarning,
}

var mappingVirtualCircuitErrorDetailsSeverityLevelEnumLowerCase = map[string]VirtualCircuitErrorDetailsSeverityLevelEnum{
	"error":   VirtualCircuitErrorDetailsSeverityLevelError,
	"warning": VirtualCircuitErrorDetailsSeverityLevelWarning,
}

// GetVirtualCircuitErrorDetailsSeverityLevelEnumValues Enumerates the set of values for VirtualCircuitErrorDetailsSeverityLevelEnum
func GetVirtualCircuitErrorDetailsSeverityLevelEnumValues() []VirtualCircuitErrorDetailsSeverityLevelEnum {
	values := make([]VirtualCircuitErrorDetailsSeverityLevelEnum, 0)
	for _, v := range mappingVirtualCircuitErrorDetailsSeverityLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualCircuitErrorDetailsSeverityLevelEnumStringValues Enumerates the set of values in String for VirtualCircuitErrorDetailsSeverityLevelEnum
func GetVirtualCircuitErrorDetailsSeverityLevelEnumStringValues() []string {
	return []string{
		"ERROR",
		"WARNING",
	}
}

// GetMappingVirtualCircuitErrorDetailsSeverityLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualCircuitErrorDetailsSeverityLevelEnum(val string) (VirtualCircuitErrorDetailsSeverityLevelEnum, bool) {
	enum, ok := mappingVirtualCircuitErrorDetailsSeverityLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
