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

// TargetDatabase The details of the Data Safe target database.
type TargetDatabase struct {

	// The OCID of the compartment which contains the Data Safe target database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Data Safe target database.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the target database in Data Safe.
	DisplayName *string `mandatory:"true" json:"displayName"`

	DatabaseDetails DatabaseDetails `mandatory:"true" json:"databaseDetails"`

	// The current state of the target database in Data Safe.
	LifecycleState TargetDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time of the target database registration and creation in Data Safe.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The description of the target database in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	Credentials *Credentials `mandatory:"false" json:"credentials"`

	TlsConfig *TlsConfig `mandatory:"false" json:"tlsConfig"`

	ConnectionOption ConnectionOption `mandatory:"false" json:"connectionOption"`

	// The OCIDs of associated resources like database, Data Safe private endpoint etc.
	AssociatedResourceIds []string `mandatory:"false" json:"associatedResourceIds"`

	// Details about the current state of the target database in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time of the target database update in Data Safe.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TargetDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TargetDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string                           `json:"description"`
		Credentials           *Credentials                      `json:"credentials"`
		TlsConfig             *TlsConfig                        `json:"tlsConfig"`
		ConnectionOption      connectionoption                  `json:"connectionOption"`
		AssociatedResourceIds []string                          `json:"associatedResourceIds"`
		LifecycleDetails      *string                           `json:"lifecycleDetails"`
		TimeUpdated           *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		SystemTags            map[string]map[string]interface{} `json:"systemTags"`
		CompartmentId         *string                           `json:"compartmentId"`
		Id                    *string                           `json:"id"`
		DisplayName           *string                           `json:"displayName"`
		DatabaseDetails       databasedetails                   `json:"databaseDetails"`
		LifecycleState        TargetDatabaseLifecycleStateEnum  `json:"lifecycleState"`
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
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

	m.AssociatedResourceIds = make([]string, len(model.AssociatedResourceIds))
	copy(m.AssociatedResourceIds, model.AssociatedResourceIds)
	m.LifecycleDetails = model.LifecycleDetails

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.CompartmentId = model.CompartmentId

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	nn, e = model.DatabaseDetails.UnmarshalPolymorphicJSON(model.DatabaseDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseDetails = nn.(DatabaseDetails)
	} else {
		m.DatabaseDetails = nil
	}

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	return
}
