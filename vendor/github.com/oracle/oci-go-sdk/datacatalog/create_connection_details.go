// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateConnectionDetails Properties used in connection create operations.
type CreateConnectionDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The key of the object type. Type key's can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"true" json:"typeKey"`

	// A map of maps that contains the properties which are specific to the connection type. Each connection type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// connections have required properties within the "default" category. To determine the set of optional and
	// required properties for a connection type, a query can be done on '/types?type=connection' that returns a
	// collection of all connection types. The appropriate connection type, which will include definitions of all
	// of it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "username": "user1"}}}`
	Properties map[string]map[string]string `mandatory:"true" json:"properties"`

	// A description of the connection.
	Description *string `mandatory:"false" json:"description"`

	// A map of maps that contains the encrypted values for sensitive properties which are specific to the
	// connection type. Each connection type definition defines it's set of required and optional properties.
	// The map keys are category names and the values are maps of property name to property value. Every property is
	// contained inside of a category. Most connections have required properties within the "default" category.
	// To determine the set of optional and required properties for a connection type, a query can be done
	// on '/types?type=connection' that returns a collection of all connection types. The appropriate connection
	// type, which will include definitions of all of it's properties, can be identified from this collection.
	// Example: `{"encProperties": { "default": { "password": "pwd"}}}`
	EncProperties map[string]map[string]string `mandatory:"false" json:"encProperties"`

	// Indicates whether this connection is the default connection. The first connection of a data asset defaults
	// to being the default, subsequent connections default to not being the default. If a default connection already
	// exists, then trying to create a connection as the default will fail. In this case the default connection would
	// need to be updated not to be the default and then the new connection can then be created as the default.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m CreateConnectionDetails) String() string {
	return common.PointerString(m)
}
