// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataInfrastructureFleetStatusByCategory The number of Exadata infrastructures in the fleet, grouped by deployment type and rack-size.
type ExadataInfrastructureFleetStatusByCategory struct {

	// The infrastructure deployment type.
	DeploymentType ExadataInfrastructureDeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// The size of the Exadata infrastructure.
	RackSize ExadataInfrastructureFleetStatusByCategoryRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The number of Exadata infrastructures in the fleet.
	InventoryCount *int `mandatory:"false" json:"inventoryCount"`
}

func (m ExadataInfrastructureFleetStatusByCategory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructureFleetStatusByCategory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataInfrastructureDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetExadataInfrastructureDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInfrastructureFleetStatusByCategoryRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetExadataInfrastructureFleetStatusByCategoryRackSizeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataInfrastructureFleetStatusByCategoryRackSizeEnum Enum with underlying type: string
type ExadataInfrastructureFleetStatusByCategoryRackSizeEnum string

// Set of constants representing the allowable values for ExadataInfrastructureFleetStatusByCategoryRackSizeEnum
const (
	ExadataInfrastructureFleetStatusByCategoryRackSizeFull    ExadataInfrastructureFleetStatusByCategoryRackSizeEnum = "FULL"
	ExadataInfrastructureFleetStatusByCategoryRackSizeHalf    ExadataInfrastructureFleetStatusByCategoryRackSizeEnum = "HALF"
	ExadataInfrastructureFleetStatusByCategoryRackSizeQuarter ExadataInfrastructureFleetStatusByCategoryRackSizeEnum = "QUARTER"
	ExadataInfrastructureFleetStatusByCategoryRackSizeEighth  ExadataInfrastructureFleetStatusByCategoryRackSizeEnum = "EIGHTH"
	ExadataInfrastructureFleetStatusByCategoryRackSizeOther   ExadataInfrastructureFleetStatusByCategoryRackSizeEnum = "OTHER"
)

var mappingExadataInfrastructureFleetStatusByCategoryRackSizeEnum = map[string]ExadataInfrastructureFleetStatusByCategoryRackSizeEnum{
	"FULL":    ExadataInfrastructureFleetStatusByCategoryRackSizeFull,
	"HALF":    ExadataInfrastructureFleetStatusByCategoryRackSizeHalf,
	"QUARTER": ExadataInfrastructureFleetStatusByCategoryRackSizeQuarter,
	"EIGHTH":  ExadataInfrastructureFleetStatusByCategoryRackSizeEighth,
	"OTHER":   ExadataInfrastructureFleetStatusByCategoryRackSizeOther,
}

var mappingExadataInfrastructureFleetStatusByCategoryRackSizeEnumLowerCase = map[string]ExadataInfrastructureFleetStatusByCategoryRackSizeEnum{
	"full":    ExadataInfrastructureFleetStatusByCategoryRackSizeFull,
	"half":    ExadataInfrastructureFleetStatusByCategoryRackSizeHalf,
	"quarter": ExadataInfrastructureFleetStatusByCategoryRackSizeQuarter,
	"eighth":  ExadataInfrastructureFleetStatusByCategoryRackSizeEighth,
	"other":   ExadataInfrastructureFleetStatusByCategoryRackSizeOther,
}

// GetExadataInfrastructureFleetStatusByCategoryRackSizeEnumValues Enumerates the set of values for ExadataInfrastructureFleetStatusByCategoryRackSizeEnum
func GetExadataInfrastructureFleetStatusByCategoryRackSizeEnumValues() []ExadataInfrastructureFleetStatusByCategoryRackSizeEnum {
	values := make([]ExadataInfrastructureFleetStatusByCategoryRackSizeEnum, 0)
	for _, v := range mappingExadataInfrastructureFleetStatusByCategoryRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureFleetStatusByCategoryRackSizeEnumStringValues Enumerates the set of values in String for ExadataInfrastructureFleetStatusByCategoryRackSizeEnum
func GetExadataInfrastructureFleetStatusByCategoryRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
		"OTHER",
	}
}

// GetMappingExadataInfrastructureFleetStatusByCategoryRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureFleetStatusByCategoryRackSizeEnum(val string) (ExadataInfrastructureFleetStatusByCategoryRackSizeEnum, bool) {
	enum, ok := mappingExadataInfrastructureFleetStatusByCategoryRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
