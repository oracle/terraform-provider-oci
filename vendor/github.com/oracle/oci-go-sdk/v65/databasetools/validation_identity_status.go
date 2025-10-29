// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidationIdentityStatus Status details for an identity type
type ValidationIdentityStatus struct {

	// The Database Tools identity type.
	Type IdentityTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The status of the identity. AVAILABLE if the identity type can be used with this connection. UNAVAILABLE if not.
	Status DatabaseToolsIdentityStatusEnum `mandatory:"false" json:"status,omitempty"`

	// If the status is UNAVAILABLE this displays the cause.
	StatusDetails *string `mandatory:"false" json:"statusDetails"`
}

func (m ValidationIdentityStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidationIdentityStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIdentityTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdentityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsIdentityStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDatabaseToolsIdentityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
