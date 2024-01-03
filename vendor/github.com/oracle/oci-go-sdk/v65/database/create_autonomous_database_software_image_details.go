// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAutonomousDatabaseSoftwareImageDetails Parameters for creating a Autonomous Database Software Image
type CreateAutonomousDatabaseSoftwareImageDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous Database Software Image. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The source Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) from which to create Autonomous Database Software Image.
	SourceCdbId *string `mandatory:"true" json:"sourceCdbId"`

	// To what shape the image is meant for.
	ImageShapeFamily CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum `mandatory:"true" json:"imageShapeFamily"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAutonomousDatabaseSoftwareImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousDatabaseSoftwareImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum(string(m.ImageShapeFamily)); !ok && m.ImageShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageShapeFamily: %s. Supported values are: %s.", m.ImageShapeFamily, strings.Join(GetCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum Enum with underlying type: string
type CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum
const (
	CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum = "EXADATA_SHAPE"
	CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyExaccShape   CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum = "EXACC_SHAPE"
)

var mappingCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum = map[string]CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum{
	"EXADATA_SHAPE": CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape,
	"EXACC_SHAPE":   CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyExaccShape,
}

var mappingCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnumLowerCase = map[string]CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum{
	"exadata_shape": CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape,
	"exacc_shape":   CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyExaccShape,
}

// GetCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnumValues Enumerates the set of values for CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum
func GetCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnumValues() []CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum {
	values := make([]CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnumStringValues Enumerates the set of values in String for CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum
func GetCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnumStringValues() []string {
	return []string{
		"EXADATA_SHAPE",
		"EXACC_SHAPE",
	}
}

// GetMappingCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum(val string) (CreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnum, bool) {
	enum, ok := mappingCreateAutonomousDatabaseSoftwareImageDetailsImageShapeFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
