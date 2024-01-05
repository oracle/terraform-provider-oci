// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateTargetDatabaseDetails The details used to register the database in Data Safe and to create the Data Safe target database.
type CreateTargetDatabaseDetails struct {

	// The OCID of the compartment in which to create the Data Safe target database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DatabaseDetails DatabaseDetails `mandatory:"true" json:"databaseDetails"`

	// The display name of the target database in Data Safe. The name is modifiable and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the target database in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	Credentials *Credentials `mandatory:"false" json:"credentials"`

	TlsConfig *TlsConfig `mandatory:"false" json:"tlsConfig"`

	ConnectionOption ConnectionOption `mandatory:"false" json:"connectionOption"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateTargetDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTargetDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateTargetDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		Credentials      *Credentials                      `json:"credentials"`
		TlsConfig        *TlsConfig                        `json:"tlsConfig"`
		ConnectionOption connectionoption                  `json:"connectionOption"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId    *string                           `json:"compartmentId"`
		DatabaseDetails  databasedetails                   `json:"databaseDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.Credentials = model.Credentials

	m.TlsConfig = model.TlsConfig

	nn, e = model.ConnectionOption.UnmarshalPolymorphicJSON(model.ConnectionOption.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionOption = nn.(ConnectionOption)
	} else {
		m.ConnectionOption = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	nn, e = model.DatabaseDetails.UnmarshalPolymorphicJSON(model.DatabaseDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseDetails = nn.(DatabaseDetails)
	} else {
		m.DatabaseDetails = nil
	}

	return
}
