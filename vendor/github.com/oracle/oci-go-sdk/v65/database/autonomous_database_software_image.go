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

// AutonomousDatabaseSoftwareImage Autonomous Database Software Images created from Autonomous Container Database
type AutonomousDatabaseSoftwareImage struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database Software Image.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database version with which the Autonomous Database Software Image is to be built.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// The user-friendly name for the Autonomous Database Software Image. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Autonomous Database Software Image.
	LifecycleState AutonomousDatabaseSoftwareImageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Autonomous Database Software Image was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The Release Updates.
	ReleaseUpdate *string `mandatory:"true" json:"releaseUpdate"`

	// To what shape the image is meant for.
	ImageShapeFamily AutonomousDatabaseSoftwareImageImageShapeFamilyEnum `mandatory:"true" json:"imageShapeFamily"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// One-off patches included in the Autonomous Database Software Image
	AutonomousDsiOneOffPatches []string `mandatory:"false" json:"autonomousDsiOneOffPatches"`
}

func (m AutonomousDatabaseSoftwareImage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseSoftwareImage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseSoftwareImageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseSoftwareImageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnum(string(m.ImageShapeFamily)); !ok && m.ImageShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageShapeFamily: %s. Supported values are: %s.", m.ImageShapeFamily, strings.Join(GetAutonomousDatabaseSoftwareImageImageShapeFamilyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseSoftwareImageLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseSoftwareImageLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSoftwareImageLifecycleStateEnum
const (
	AutonomousDatabaseSoftwareImageLifecycleStateAvailable    AutonomousDatabaseSoftwareImageLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseSoftwareImageLifecycleStateFailed       AutonomousDatabaseSoftwareImageLifecycleStateEnum = "FAILED"
	AutonomousDatabaseSoftwareImageLifecycleStateProvisioning AutonomousDatabaseSoftwareImageLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseSoftwareImageLifecycleStateExpired      AutonomousDatabaseSoftwareImageLifecycleStateEnum = "EXPIRED"
	AutonomousDatabaseSoftwareImageLifecycleStateTerminated   AutonomousDatabaseSoftwareImageLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseSoftwareImageLifecycleStateTerminating  AutonomousDatabaseSoftwareImageLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseSoftwareImageLifecycleStateUpdating     AutonomousDatabaseSoftwareImageLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDatabaseSoftwareImageLifecycleStateEnum = map[string]AutonomousDatabaseSoftwareImageLifecycleStateEnum{
	"AVAILABLE":    AutonomousDatabaseSoftwareImageLifecycleStateAvailable,
	"FAILED":       AutonomousDatabaseSoftwareImageLifecycleStateFailed,
	"PROVISIONING": AutonomousDatabaseSoftwareImageLifecycleStateProvisioning,
	"EXPIRED":      AutonomousDatabaseSoftwareImageLifecycleStateExpired,
	"TERMINATED":   AutonomousDatabaseSoftwareImageLifecycleStateTerminated,
	"TERMINATING":  AutonomousDatabaseSoftwareImageLifecycleStateTerminating,
	"UPDATING":     AutonomousDatabaseSoftwareImageLifecycleStateUpdating,
}

var mappingAutonomousDatabaseSoftwareImageLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseSoftwareImageLifecycleStateEnum{
	"available":    AutonomousDatabaseSoftwareImageLifecycleStateAvailable,
	"failed":       AutonomousDatabaseSoftwareImageLifecycleStateFailed,
	"provisioning": AutonomousDatabaseSoftwareImageLifecycleStateProvisioning,
	"expired":      AutonomousDatabaseSoftwareImageLifecycleStateExpired,
	"terminated":   AutonomousDatabaseSoftwareImageLifecycleStateTerminated,
	"terminating":  AutonomousDatabaseSoftwareImageLifecycleStateTerminating,
	"updating":     AutonomousDatabaseSoftwareImageLifecycleStateUpdating,
}

// GetAutonomousDatabaseSoftwareImageLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseSoftwareImageLifecycleStateEnum
func GetAutonomousDatabaseSoftwareImageLifecycleStateEnumValues() []AutonomousDatabaseSoftwareImageLifecycleStateEnum {
	values := make([]AutonomousDatabaseSoftwareImageLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseSoftwareImageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSoftwareImageLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSoftwareImageLifecycleStateEnum
func GetAutonomousDatabaseSoftwareImageLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"FAILED",
		"PROVISIONING",
		"EXPIRED",
		"TERMINATED",
		"TERMINATING",
		"UPDATING",
	}
}

// GetMappingAutonomousDatabaseSoftwareImageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSoftwareImageLifecycleStateEnum(val string) (AutonomousDatabaseSoftwareImageLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSoftwareImageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSoftwareImageImageShapeFamilyEnum Enum with underlying type: string
type AutonomousDatabaseSoftwareImageImageShapeFamilyEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSoftwareImageImageShapeFamilyEnum
const (
	AutonomousDatabaseSoftwareImageImageShapeFamilyExaccShape   AutonomousDatabaseSoftwareImageImageShapeFamilyEnum = "EXACC_SHAPE"
	AutonomousDatabaseSoftwareImageImageShapeFamilyExadataShape AutonomousDatabaseSoftwareImageImageShapeFamilyEnum = "EXADATA_SHAPE"
)

var mappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnum = map[string]AutonomousDatabaseSoftwareImageImageShapeFamilyEnum{
	"EXACC_SHAPE":   AutonomousDatabaseSoftwareImageImageShapeFamilyExaccShape,
	"EXADATA_SHAPE": AutonomousDatabaseSoftwareImageImageShapeFamilyExadataShape,
}

var mappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnumLowerCase = map[string]AutonomousDatabaseSoftwareImageImageShapeFamilyEnum{
	"exacc_shape":   AutonomousDatabaseSoftwareImageImageShapeFamilyExaccShape,
	"exadata_shape": AutonomousDatabaseSoftwareImageImageShapeFamilyExadataShape,
}

// GetAutonomousDatabaseSoftwareImageImageShapeFamilyEnumValues Enumerates the set of values for AutonomousDatabaseSoftwareImageImageShapeFamilyEnum
func GetAutonomousDatabaseSoftwareImageImageShapeFamilyEnumValues() []AutonomousDatabaseSoftwareImageImageShapeFamilyEnum {
	values := make([]AutonomousDatabaseSoftwareImageImageShapeFamilyEnum, 0)
	for _, v := range mappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSoftwareImageImageShapeFamilyEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSoftwareImageImageShapeFamilyEnum
func GetAutonomousDatabaseSoftwareImageImageShapeFamilyEnumStringValues() []string {
	return []string{
		"EXACC_SHAPE",
		"EXADATA_SHAPE",
	}
}

// GetMappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnum(val string) (AutonomousDatabaseSoftwareImageImageShapeFamilyEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
