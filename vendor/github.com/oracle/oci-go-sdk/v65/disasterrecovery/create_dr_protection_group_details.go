// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDrProtectionGroupDetails The details for creating a DR protection group.
type CreateDrProtectionGroupDetails struct {

	// The OCID of the compartment in which to create the DR protection group.
	// Example: `ocid1.compartment.oc1..uniqueID`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the DR protection group.
	// Example: `EBS PHX Group`
	DisplayName *string `mandatory:"true" json:"displayName"`

	LogLocation *CreateObjectStorageLogLocationDetails `mandatory:"true" json:"logLocation"`

	Association *AssociateDrProtectionGroupDetails `mandatory:"false" json:"association"`

	// A list of DR protection group members.
	Members []CreateDrProtectionGroupMemberDetails `mandatory:"false" json:"members"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDrProtectionGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDrProtectionGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDrProtectionGroupDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Association   *AssociateDrProtectionGroupDetails     `json:"association"`
		Members       []createdrprotectiongroupmemberdetails `json:"members"`
		FreeformTags  map[string]string                      `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{}      `json:"definedTags"`
		CompartmentId *string                                `json:"compartmentId"`
		DisplayName   *string                                `json:"displayName"`
		LogLocation   *CreateObjectStorageLogLocationDetails `json:"logLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Association = model.Association

	m.Members = make([]CreateDrProtectionGroupMemberDetails, len(model.Members))
	for i, n := range model.Members {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Members[i] = nn.(CreateDrProtectionGroupMemberDetails)
		} else {
			m.Members[i] = nil
		}
	}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.LogLocation = model.LogLocation

	return
}
