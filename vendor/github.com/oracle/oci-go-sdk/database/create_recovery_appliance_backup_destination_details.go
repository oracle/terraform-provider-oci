// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateRecoveryApplianceBackupDestinationDetails Used for creating Recovery Appliance backup destinations.
type CreateRecoveryApplianceBackupDestinationDetails struct {

	// The user-provided name of the backup destination.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The connection string for connecting to the Recovery Appliance.
	ConnectionString *string `mandatory:"true" json:"connectionString"`

	// The Virtual Private Catalog (VPC) users that are used to access the Recovery Appliance.
	VpcUsers []string `mandatory:"true" json:"vpcUsers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m CreateRecoveryApplianceBackupDestinationDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateRecoveryApplianceBackupDestinationDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m CreateRecoveryApplianceBackupDestinationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateRecoveryApplianceBackupDestinationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateRecoveryApplianceBackupDestinationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateRecoveryApplianceBackupDestinationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateRecoveryApplianceBackupDestinationDetails CreateRecoveryApplianceBackupDestinationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateRecoveryApplianceBackupDestinationDetails
	}{
		"RECOVERY_APPLIANCE",
		(MarshalTypeCreateRecoveryApplianceBackupDestinationDetails)(m),
	}

	return json.Marshal(&s)
}
