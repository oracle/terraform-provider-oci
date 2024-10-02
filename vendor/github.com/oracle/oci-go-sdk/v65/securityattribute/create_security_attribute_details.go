// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSecurityAttributeDetails Details of the security attribute to be created for a specific security attribute namespace.
type CreateSecurityAttributeDetails struct {

	// The name you assign to the security attribute during creation. This is the security attribute key.
	// The name must be unique within the namespace and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the security attribute during creation.
	Description *string `mandatory:"true" json:"description"`

	Validator BaseSecurityAttributeValidator `mandatory:"false" json:"validator"`
}

func (m CreateSecurityAttributeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSecurityAttributeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateSecurityAttributeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Validator   basesecurityattributevalidator `json:"validator"`
		Name        *string                        `json:"name"`
		Description *string                        `json:"description"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Validator.UnmarshalPolymorphicJSON(model.Validator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Validator = nn.(BaseSecurityAttributeValidator)
	} else {
		m.Validator = nil
	}

	m.Name = model.Name

	m.Description = model.Description

	return
}
