// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciControlCenterCp API
//
// A description of the OciControlCenterCp API
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccAvailabilityCatalog Details of the availability catalog resource.
type OccAvailabilityCatalog struct {

	// The OCID of the availability catalog.
	Id *string `mandatory:"true" json:"id"`

	// The name of the OCI service in consideration. For example, Compute, Exadata, and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The OCID of the tenancy where the availability catalog resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the availability catalog.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The customer group OCID to which the availability catalog belongs.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// The different states associated with the availability catalog.
	CatalogState OccAvailabilityCatalogCatalogStateEnum `mandatory:"true" json:"catalogState"`

	MetadataDetails *MetadataDetails `mandatory:"true" json:"metadataDetails"`

	// The time when the availability catalog was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the availability catalog was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the resource.
	LifecycleState OccAvailabilityCatalogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Text information about the availability catalog.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Details about capacity available for  different resources in catalog.
	Details []OccAvailabilitySummary `mandatory:"false" json:"details"`
}

func (m OccAvailabilityCatalog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccAvailabilityCatalog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilityCatalogCatalogStateEnum(string(m.CatalogState)); !ok && m.CatalogState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CatalogState: %s. Supported values are: %s.", m.CatalogState, strings.Join(GetOccAvailabilityCatalogCatalogStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilityCatalogLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccAvailabilityCatalogLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccAvailabilityCatalogCatalogStateEnum Enum with underlying type: string
type OccAvailabilityCatalogCatalogStateEnum string

// Set of constants representing the allowable values for OccAvailabilityCatalogCatalogStateEnum
const (
	OccAvailabilityCatalogCatalogStateNotUploaded  OccAvailabilityCatalogCatalogStateEnum = "NOT_UPLOADED"
	OccAvailabilityCatalogCatalogStateUploadFailed OccAvailabilityCatalogCatalogStateEnum = "UPLOAD_FAILED"
	OccAvailabilityCatalogCatalogStateStaged       OccAvailabilityCatalogCatalogStateEnum = "STAGED"
	OccAvailabilityCatalogCatalogStatePublished    OccAvailabilityCatalogCatalogStateEnum = "PUBLISHED"
	OccAvailabilityCatalogCatalogStateOutdated     OccAvailabilityCatalogCatalogStateEnum = "OUTDATED"
	OccAvailabilityCatalogCatalogStateDeleted      OccAvailabilityCatalogCatalogStateEnum = "DELETED"
)

var mappingOccAvailabilityCatalogCatalogStateEnum = map[string]OccAvailabilityCatalogCatalogStateEnum{
	"NOT_UPLOADED":  OccAvailabilityCatalogCatalogStateNotUploaded,
	"UPLOAD_FAILED": OccAvailabilityCatalogCatalogStateUploadFailed,
	"STAGED":        OccAvailabilityCatalogCatalogStateStaged,
	"PUBLISHED":     OccAvailabilityCatalogCatalogStatePublished,
	"OUTDATED":      OccAvailabilityCatalogCatalogStateOutdated,
	"DELETED":       OccAvailabilityCatalogCatalogStateDeleted,
}

var mappingOccAvailabilityCatalogCatalogStateEnumLowerCase = map[string]OccAvailabilityCatalogCatalogStateEnum{
	"not_uploaded":  OccAvailabilityCatalogCatalogStateNotUploaded,
	"upload_failed": OccAvailabilityCatalogCatalogStateUploadFailed,
	"staged":        OccAvailabilityCatalogCatalogStateStaged,
	"published":     OccAvailabilityCatalogCatalogStatePublished,
	"outdated":      OccAvailabilityCatalogCatalogStateOutdated,
	"deleted":       OccAvailabilityCatalogCatalogStateDeleted,
}

// GetOccAvailabilityCatalogCatalogStateEnumValues Enumerates the set of values for OccAvailabilityCatalogCatalogStateEnum
func GetOccAvailabilityCatalogCatalogStateEnumValues() []OccAvailabilityCatalogCatalogStateEnum {
	values := make([]OccAvailabilityCatalogCatalogStateEnum, 0)
	for _, v := range mappingOccAvailabilityCatalogCatalogStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccAvailabilityCatalogCatalogStateEnumStringValues Enumerates the set of values in String for OccAvailabilityCatalogCatalogStateEnum
func GetOccAvailabilityCatalogCatalogStateEnumStringValues() []string {
	return []string{
		"NOT_UPLOADED",
		"UPLOAD_FAILED",
		"STAGED",
		"PUBLISHED",
		"OUTDATED",
		"DELETED",
	}
}

// GetMappingOccAvailabilityCatalogCatalogStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccAvailabilityCatalogCatalogStateEnum(val string) (OccAvailabilityCatalogCatalogStateEnum, bool) {
	enum, ok := mappingOccAvailabilityCatalogCatalogStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccAvailabilityCatalogLifecycleStateEnum Enum with underlying type: string
type OccAvailabilityCatalogLifecycleStateEnum string

// Set of constants representing the allowable values for OccAvailabilityCatalogLifecycleStateEnum
const (
	OccAvailabilityCatalogLifecycleStateCreating OccAvailabilityCatalogLifecycleStateEnum = "CREATING"
	OccAvailabilityCatalogLifecycleStateUpdating OccAvailabilityCatalogLifecycleStateEnum = "UPDATING"
	OccAvailabilityCatalogLifecycleStateActive   OccAvailabilityCatalogLifecycleStateEnum = "ACTIVE"
	OccAvailabilityCatalogLifecycleStateDeleting OccAvailabilityCatalogLifecycleStateEnum = "DELETING"
	OccAvailabilityCatalogLifecycleStateDeleted  OccAvailabilityCatalogLifecycleStateEnum = "DELETED"
	OccAvailabilityCatalogLifecycleStateFailed   OccAvailabilityCatalogLifecycleStateEnum = "FAILED"
)

var mappingOccAvailabilityCatalogLifecycleStateEnum = map[string]OccAvailabilityCatalogLifecycleStateEnum{
	"CREATING": OccAvailabilityCatalogLifecycleStateCreating,
	"UPDATING": OccAvailabilityCatalogLifecycleStateUpdating,
	"ACTIVE":   OccAvailabilityCatalogLifecycleStateActive,
	"DELETING": OccAvailabilityCatalogLifecycleStateDeleting,
	"DELETED":  OccAvailabilityCatalogLifecycleStateDeleted,
	"FAILED":   OccAvailabilityCatalogLifecycleStateFailed,
}

var mappingOccAvailabilityCatalogLifecycleStateEnumLowerCase = map[string]OccAvailabilityCatalogLifecycleStateEnum{
	"creating": OccAvailabilityCatalogLifecycleStateCreating,
	"updating": OccAvailabilityCatalogLifecycleStateUpdating,
	"active":   OccAvailabilityCatalogLifecycleStateActive,
	"deleting": OccAvailabilityCatalogLifecycleStateDeleting,
	"deleted":  OccAvailabilityCatalogLifecycleStateDeleted,
	"failed":   OccAvailabilityCatalogLifecycleStateFailed,
}

// GetOccAvailabilityCatalogLifecycleStateEnumValues Enumerates the set of values for OccAvailabilityCatalogLifecycleStateEnum
func GetOccAvailabilityCatalogLifecycleStateEnumValues() []OccAvailabilityCatalogLifecycleStateEnum {
	values := make([]OccAvailabilityCatalogLifecycleStateEnum, 0)
	for _, v := range mappingOccAvailabilityCatalogLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccAvailabilityCatalogLifecycleStateEnumStringValues Enumerates the set of values in String for OccAvailabilityCatalogLifecycleStateEnum
func GetOccAvailabilityCatalogLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOccAvailabilityCatalogLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccAvailabilityCatalogLifecycleStateEnum(val string) (OccAvailabilityCatalogLifecycleStateEnum, bool) {
	enum, ok := mappingOccAvailabilityCatalogLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
