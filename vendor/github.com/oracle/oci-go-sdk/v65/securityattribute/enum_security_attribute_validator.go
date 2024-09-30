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

// EnumSecurityAttributeValidator Used to validate the value set for a security attribute and contains the list of allowable `values`.
// You must specify at least one valid value in the `values` array. You can't have blank or
// or empty strings (`""`). Duplicate values are not allowed.
type EnumSecurityAttributeValidator struct {

	// The list of allowed values for a security attribute value.
	Values []string `mandatory:"false" json:"values"`
}

func (m EnumSecurityAttributeValidator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnumSecurityAttributeValidator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EnumSecurityAttributeValidator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEnumSecurityAttributeValidator EnumSecurityAttributeValidator
	s := struct {
		DiscriminatorParam string `json:"validatorType"`
		MarshalTypeEnumSecurityAttributeValidator
	}{
		"ENUM",
		(MarshalTypeEnumSecurityAttributeValidator)(m),
	}

	return json.Marshal(&s)
}
