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

// UpdateOciCacheUserDetails Details required to update an existing OCI cache user.
type UpdateOciCacheUserDetails struct {

	// Description of OCI cache user.
	Description *string `mandatory:"false" json:"description"`

	AuthenticationMode AuthenticationMode `mandatory:"false" json:"authenticationMode"`

	// ACL string of OCI cache user.
	AclString *string `mandatory:"false" json:"aclString"`

	// OCI cache user status. ON enables and OFF disables the OCI cache user to login to the associated clusters.
	Status OciCacheUserStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOciCacheUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOciCacheUserDetails) ValidateEnumValue() (bool, error) {
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
func (m *UpdateOciCacheUserDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description        *string                           `json:"description"`
		AuthenticationMode authenticationmode                `json:"authenticationMode"`
		AclString          *string                           `json:"aclString"`
		Status             OciCacheUserStatusEnum            `json:"status"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

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

	m.Status = model.Status

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
