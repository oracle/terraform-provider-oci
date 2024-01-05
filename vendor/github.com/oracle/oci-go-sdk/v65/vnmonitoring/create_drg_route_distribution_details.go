// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDrgRouteDistributionDetails Details used to create a route distribution.
type CreateDrgRouteDistributionDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG route table belongs to.
	DrgId *string `mandatory:"true" json:"drgId"`

	// Whether this distribution defines how routes get imported into route tables or exported through DRG Attachments
	DistributionType CreateDrgRouteDistributionDetailsDistributionTypeEnum `mandatory:"true" json:"distributionType"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateDrgRouteDistributionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDrgRouteDistributionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDrgRouteDistributionDetailsDistributionTypeEnum(string(m.DistributionType)); !ok && m.DistributionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DistributionType: %s. Supported values are: %s.", m.DistributionType, strings.Join(GetCreateDrgRouteDistributionDetailsDistributionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDrgRouteDistributionDetailsDistributionTypeEnum Enum with underlying type: string
type CreateDrgRouteDistributionDetailsDistributionTypeEnum string

// Set of constants representing the allowable values for CreateDrgRouteDistributionDetailsDistributionTypeEnum
const (
	CreateDrgRouteDistributionDetailsDistributionTypeImport CreateDrgRouteDistributionDetailsDistributionTypeEnum = "IMPORT"
)

var mappingCreateDrgRouteDistributionDetailsDistributionTypeEnum = map[string]CreateDrgRouteDistributionDetailsDistributionTypeEnum{
	"IMPORT": CreateDrgRouteDistributionDetailsDistributionTypeImport,
}

var mappingCreateDrgRouteDistributionDetailsDistributionTypeEnumLowerCase = map[string]CreateDrgRouteDistributionDetailsDistributionTypeEnum{
	"import": CreateDrgRouteDistributionDetailsDistributionTypeImport,
}

// GetCreateDrgRouteDistributionDetailsDistributionTypeEnumValues Enumerates the set of values for CreateDrgRouteDistributionDetailsDistributionTypeEnum
func GetCreateDrgRouteDistributionDetailsDistributionTypeEnumValues() []CreateDrgRouteDistributionDetailsDistributionTypeEnum {
	values := make([]CreateDrgRouteDistributionDetailsDistributionTypeEnum, 0)
	for _, v := range mappingCreateDrgRouteDistributionDetailsDistributionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDrgRouteDistributionDetailsDistributionTypeEnumStringValues Enumerates the set of values in String for CreateDrgRouteDistributionDetailsDistributionTypeEnum
func GetCreateDrgRouteDistributionDetailsDistributionTypeEnumStringValues() []string {
	return []string{
		"IMPORT",
	}
}

// GetMappingCreateDrgRouteDistributionDetailsDistributionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDrgRouteDistributionDetailsDistributionTypeEnum(val string) (CreateDrgRouteDistributionDetailsDistributionTypeEnum, bool) {
	enum, ok := mappingCreateDrgRouteDistributionDetailsDistributionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
