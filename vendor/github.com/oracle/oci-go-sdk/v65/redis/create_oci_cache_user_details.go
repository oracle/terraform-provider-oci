// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOciCacheUserDetails Details required to create a new OCI cache user.
type CreateOciCacheUserDetails struct {

	// OCI cache user name is required to connect to an OCI cache cluster.
	Name *string `mandatory:"true" json:"name"`

	// Description of OCI cache user.
	Description *string `mandatory:"true" json:"description"`

	// OCI cache user compartment ID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	AuthenticationMode AuthenticationMode `mandatory:"true" json:"authenticationMode"`

	// ACL string of OCI cache user.
	AclString *string `mandatory:"true" json:"aclString"`

	// OCI cache user status. ON enables and OFF disables the OCI cache user to login to the associated clusters. Default value is ON.
	Status OciCacheUserStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOciCacheUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciCacheUserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOciCacheUserStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOciCacheUserStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateOciCacheUserDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Status             OciCacheUserStatusEnum            `json:"status"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		Name               *string                           `json:"name"`
		Description        *string                           `json:"description"`
		CompartmentId      *string                           `json:"compartmentId"`
		AuthenticationMode authenticationmode                `json:"authenticationMode"`
		AclString          *string                           `json:"aclString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Status = model.Status

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.Description = model.Description

	m.CompartmentId = model.CompartmentId

	nn, e = model.AuthenticationMode.UnmarshalPolymorphicJSON(model.AuthenticationMode.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthenticationMode = nn.(AuthenticationMode)
	} else {
		m.AuthenticationMode = nil
	}

	m.AclString = model.AclString

	return
}
