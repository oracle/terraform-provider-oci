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

// MonitoredResourceAssociation Association details between two monitored resources.
type MonitoredResourceAssociation struct {

	// Association Type.
	AssociationType *string `mandatory:"true" json:"associationType"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Tenancy Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	TenantId *string `mandatory:"true" json:"tenantId"`

	// Source Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SourceResourceId *string `mandatory:"true" json:"sourceResourceId"`

	// Destination Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DestinationResourceId *string `mandatory:"true" json:"destinationResourceId"`

	SourceResourceDetails *AssociationResourceDetails `mandatory:"false" json:"sourceResourceDetails"`

	DestinationResourceDetails *AssociationResourceDetails `mandatory:"false" json:"destinationResourceDetails"`

	// The time when the association was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Association category. Possible values are:
	// - System created (SYSTEM),
	// - User created using API (USER_API)
	// - User created using tags (USER_TAG_ASSOC).
	Category MonitoredResourceAssociationCategoryEnum `mandatory:"false" json:"category,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MonitoredResourceAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredResourceAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMonitoredResourceAssociationCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetMonitoredResourceAssociationCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitoredResourceAssociationCategoryEnum Enum with underlying type: string
type MonitoredResourceAssociationCategoryEnum string

// Set of constants representing the allowable values for MonitoredResourceAssociationCategoryEnum
const (
	MonitoredResourceAssociationCategorySystem       MonitoredResourceAssociationCategoryEnum = "SYSTEM"
	MonitoredResourceAssociationCategoryUserApi      MonitoredResourceAssociationCategoryEnum = "USER_API"
	MonitoredResourceAssociationCategoryUserTagAssoc MonitoredResourceAssociationCategoryEnum = "USER_TAG_ASSOC"
)

var mappingMonitoredResourceAssociationCategoryEnum = map[string]MonitoredResourceAssociationCategoryEnum{
	"SYSTEM":         MonitoredResourceAssociationCategorySystem,
	"USER_API":       MonitoredResourceAssociationCategoryUserApi,
	"USER_TAG_ASSOC": MonitoredResourceAssociationCategoryUserTagAssoc,
}

var mappingMonitoredResourceAssociationCategoryEnumLowerCase = map[string]MonitoredResourceAssociationCategoryEnum{
	"system":         MonitoredResourceAssociationCategorySystem,
	"user_api":       MonitoredResourceAssociationCategoryUserApi,
	"user_tag_assoc": MonitoredResourceAssociationCategoryUserTagAssoc,
}

// GetMonitoredResourceAssociationCategoryEnumValues Enumerates the set of values for MonitoredResourceAssociationCategoryEnum
func GetMonitoredResourceAssociationCategoryEnumValues() []MonitoredResourceAssociationCategoryEnum {
	values := make([]MonitoredResourceAssociationCategoryEnum, 0)
	for _, v := range mappingMonitoredResourceAssociationCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredResourceAssociationCategoryEnumStringValues Enumerates the set of values in String for MonitoredResourceAssociationCategoryEnum
func GetMonitoredResourceAssociationCategoryEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"USER_API",
		"USER_TAG_ASSOC",
	}
}

// GetMappingMonitoredResourceAssociationCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredResourceAssociationCategoryEnum(val string) (MonitoredResourceAssociationCategoryEnum, bool) {
	enum, ok := mappingMonitoredResourceAssociationCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
