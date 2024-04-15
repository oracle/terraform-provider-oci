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

// AutonomousDatabaseSoftwareImageSummary The Database service supports the creation of Autonomous Database Software Images for use in creating Autonomous Container Database.
type AutonomousDatabaseSoftwareImageSummary struct {

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

func (m AutonomousDatabaseSoftwareImageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseSoftwareImageSummary) ValidateEnumValue() (bool, error) {
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
