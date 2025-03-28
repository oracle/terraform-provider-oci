// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Namespaces (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSecurityAttributeDetails Details of the security attribute to be updated for a specific security attribute namespace.
type UpdateSecurityAttributeDetails struct {

	// The description of the security attribute during creation.
	Description *string `mandatory:"false" json:"description"`

	// Whether the security attribute is retired.
	// See Managing Security Attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm).
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	Validator BaseSecurityAttributeValidator `mandatory:"false" json:"validator"`
}

func (m UpdateSecurityAttributeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSecurityAttributeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSecurityAttributeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description *string                        `json:"description"`
		IsRetired   *bool                          `json:"isRetired"`
		Validator   basesecurityattributevalidator `json:"validator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.IsRetired = model.IsRetired

	nn, e = model.Validator.UnmarshalPolymorphicJSON(model.Validator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Validator = nn.(BaseSecurityAttributeValidator)
	} else {
		m.Validator = nil
	}

	return
}
