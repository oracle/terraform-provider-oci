// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateConnectionDetails Properties used in connection update operations.
type UpdateConnectionDetails struct {

	// A description of the connection.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The list of customized properties along with the values for this object
	CustomPropertyMembers []CustomPropertySetUsage `mandatory:"false" json:"customPropertyMembers"`

	// A map of maps that contains the properties which are specific to the connection type. Each connection type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// connections have required properties within the "default" category. To determine the set of optional and
	// required properties for a connection type, a query can be done on '/types?type=connection' that returns a
	// collection of all connection types. The appropriate connection type, which will include definitions of all
	// of it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "username": "user1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`

	// A map of maps that contains the encrypted values for sensitive properties which are specific to the
	// connection type. Each connection type definition defines it's set of required and optional properties.
	// The map keys are category names and the values are maps of property name to property value. Every property is
	// contained inside of a category. Most connections have required properties within the "default" category.
	// To determine the set of optional and required properties for a connection type, a query can be done
	// on '/types?type=connection' that returns a collection of all connection types. The appropriate connection
	// type, which will include definitions of all of it's properties, can be identified from this collection.
	// Example: `{"encProperties": { "default": { "password": "example-password"}}}`
	EncProperties map[string]map[string]string `mandatory:"false" json:"encProperties"`

	// Indicates whether this connection is the default connection.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m UpdateConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
