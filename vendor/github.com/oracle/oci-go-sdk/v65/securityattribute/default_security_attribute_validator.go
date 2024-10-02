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

// DefaultSecurityAttributeValidator Use this validator to clear any existing validator on the security attribute with the UpdateSecurityAttribute
// operation. Using this `validatorType` is the same as not setting any value on the validator field.
// The resultant value for `validatorType` returned in the response body is `null`.
type DefaultSecurityAttributeValidator struct {
}

func (m DefaultSecurityAttributeValidator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultSecurityAttributeValidator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DefaultSecurityAttributeValidator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultSecurityAttributeValidator DefaultSecurityAttributeValidator
	s := struct {
		DiscriminatorParam string `json:"validatorType"`
		MarshalTypeDefaultSecurityAttributeValidator
	}{
		"DEFAULT",
		(MarshalTypeDefaultSecurityAttributeValidator)(m),
	}

	return json.Marshal(&s)
}
