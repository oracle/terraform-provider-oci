// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// EventConfig Describes an event configuration, for a given object type and property. Primarily, whether a property change will result in an event being emitted.
type EventConfig struct {

	// Unique type key identifier.
	TypeId *string `mandatory:"false" json:"typeId"`

	// Name of the type.
	TypeName *string `mandatory:"false" json:"typeName"`

	// Unique property key identifier.
	PropertyId *string `mandatory:"false" json:"propertyId"`

	// Name of the property.
	PropertyName *string `mandatory:"false" json:"propertyName"`

	// Status of the configuration.
	EventConfigStatus EventConfigStatusEnum `mandatory:"false" json:"eventConfigStatus,omitempty"`

	// The date and time the event was configured, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the configuration. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created the configuration.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who last modified the configuration.
	UpdatedById *string `mandatory:"false" json:"updatedById"`
}

func (m EventConfig) String() string {
	return common.PointerString(m)
}
