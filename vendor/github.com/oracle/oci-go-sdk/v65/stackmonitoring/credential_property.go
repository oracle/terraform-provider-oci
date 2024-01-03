// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CredentialProperty Monitored resource credential property.
type CredentialProperty struct {

	// The name of the credential property, should confirm with names of properties of this credential's type.
	// Example: For JMXCreds type, credential property name for weblogic user is 'Username'.
	Name *string `mandatory:"true" json:"name"`

	// The value of the credential property name.
	// Example: For JMXCreds type, credential property value for 'Username' property is 'weblogic'.
	Value *string `mandatory:"true" json:"value"`
}

func (m CredentialProperty) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CredentialProperty) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
