// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ServiceCatalog The model for an Oracle Cloud Infrastructure Service Catalog.
type ServiceCatalog struct {

	// The unique identifier for the Service catalog.
	Id *string `mandatory:"true" json:"id"`

	// The Compartment id where the service catalog exists
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the service catalog.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The lifecycle state of the service catalog.
	LifecycleState ServiceCatalogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the service catalog was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-05-26T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the service catalog was last modified, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-12-10T05:10:29.721Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m ServiceCatalog) String() string {
	return common.PointerString(m)
}

// ServiceCatalogLifecycleStateEnum Enum with underlying type: string
type ServiceCatalogLifecycleStateEnum string

// Set of constants representing the allowable values for ServiceCatalogLifecycleStateEnum
const (
	ServiceCatalogLifecycleStateActive  ServiceCatalogLifecycleStateEnum = "ACTIVE"
	ServiceCatalogLifecycleStateDeleted ServiceCatalogLifecycleStateEnum = "DELETED"
)

var mappingServiceCatalogLifecycleState = map[string]ServiceCatalogLifecycleStateEnum{
	"ACTIVE":  ServiceCatalogLifecycleStateActive,
	"DELETED": ServiceCatalogLifecycleStateDeleted,
}

// GetServiceCatalogLifecycleStateEnumValues Enumerates the set of values for ServiceCatalogLifecycleStateEnum
func GetServiceCatalogLifecycleStateEnumValues() []ServiceCatalogLifecycleStateEnum {
	values := make([]ServiceCatalogLifecycleStateEnum, 0)
	for _, v := range mappingServiceCatalogLifecycleState {
		values = append(values, v)
	}
	return values
}
