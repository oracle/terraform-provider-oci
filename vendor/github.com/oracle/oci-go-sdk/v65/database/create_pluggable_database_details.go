// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePluggableDatabaseDetails Parameters for creating a pluggable database in a specified container database (CDB).
// Additional option `pdbCreationTypeDetails` can be used for creating Pluggable Database using different operations, e.g. LocalClone, Remote Clone, Relocate.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreatePluggableDatabaseDetails struct {

	// The name for the pluggable database (PDB). The name is unique in the context of a Database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	PdbName *string `mandatory:"true" json:"pdbName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the CDB
	ContainerDatabaseId *string `mandatory:"true" json:"containerDatabaseId"`

	// A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword *string `mandatory:"false" json:"pdbAdminPassword"`

	// The existing TDE wallet password of the CDB.
	TdeWalletPassword *string `mandatory:"false" json:"tdeWalletPassword"`

	// The locked mode of the pluggable database admin account. If false, the user needs to provide the PDB Admin Password to connect to it.
	// If true, the pluggable database will be locked and user cannot login to it.
	ShouldPdbAdminAccountBeLocked *bool `mandatory:"false" json:"shouldPdbAdminAccountBeLocked"`

	// The DB system administrator password of the Container Database.
	ContainerDatabaseAdminPassword *string `mandatory:"false" json:"containerDatabaseAdminPassword"`

	// Indicates whether to take Pluggable Database Backup after the operation.
	ShouldCreatePdbBackup *bool `mandatory:"false" json:"shouldCreatePdbBackup"`

	PdbCreationTypeDetails CreatePluggableDatabaseCreationTypeDetails `mandatory:"false" json:"pdbCreationTypeDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreatePluggableDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePluggableDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreatePluggableDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PdbAdminPassword               *string                                    `json:"pdbAdminPassword"`
		TdeWalletPassword              *string                                    `json:"tdeWalletPassword"`
		ShouldPdbAdminAccountBeLocked  *bool                                      `json:"shouldPdbAdminAccountBeLocked"`
		ContainerDatabaseAdminPassword *string                                    `json:"containerDatabaseAdminPassword"`
		ShouldCreatePdbBackup          *bool                                      `json:"shouldCreatePdbBackup"`
		PdbCreationTypeDetails         createpluggabledatabasecreationtypedetails `json:"pdbCreationTypeDetails"`
		FreeformTags                   map[string]string                          `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{}          `json:"definedTags"`
		PdbName                        *string                                    `json:"pdbName"`
		ContainerDatabaseId            *string                                    `json:"containerDatabaseId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.PdbAdminPassword = model.PdbAdminPassword

	m.TdeWalletPassword = model.TdeWalletPassword

	m.ShouldPdbAdminAccountBeLocked = model.ShouldPdbAdminAccountBeLocked

	m.ContainerDatabaseAdminPassword = model.ContainerDatabaseAdminPassword

	m.ShouldCreatePdbBackup = model.ShouldCreatePdbBackup

	nn, e = model.PdbCreationTypeDetails.UnmarshalPolymorphicJSON(model.PdbCreationTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PdbCreationTypeDetails = nn.(CreatePluggableDatabaseCreationTypeDetails)
	} else {
		m.PdbCreationTypeDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.PdbName = model.PdbName

	m.ContainerDatabaseId = model.ContainerDatabaseId

	return
}
