// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InfrastructureContact Contact details for Database Infrastructure.
type InfrastructureContact struct {

	// The name of the Database Infrastructure contact.
	Name *string `mandatory:"true" json:"name"`

	// The email for the Database Infrastructure contact.
	Email *string `mandatory:"true" json:"email"`

	// If `true`, this Database Infrastructure contact is a primary contact.
	// If `false`, this Database Infrastructure is a secondary contact.
	IsPrimary *bool `mandatory:"true" json:"isPrimary"`

	// The phone number for the Database Infrastructure contact.
	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	// If `true`, this Database Infrastructure contact is a valid My Oracle Support (MOS) contact.
	// If `false`, this Database Infrastructure contact is not a valid MOS contact.
	IsContactMosValidated *bool `mandatory:"false" json:"isContactMosValidated"`
}

func (m InfrastructureContact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InfrastructureContact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
