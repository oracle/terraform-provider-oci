// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalExadataInfrastructureSummary The Exadata infrastructure.
type ExternalExadataInfrastructureSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The version of the Exadata resource.
	Version *string `mandatory:"false" json:"version"`

	// The internal ID of the Exadata resource.
	InternalId *string `mandatory:"false" json:"internalId"`

	// The status of the Exadata resource.
	Status *string `mandatory:"false" json:"status"`

	// The timestamp of the creation of the Exadata resource.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The timestamp of the last update of the Exadata resource.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The details of the lifecycle state of the Exadata resource.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The additional details of the resource defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// The Oracle grid home path.
	GridHomePath *string `mandatory:"false" json:"gridHomePath"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The rack size of the Exadata infrastructure.
	RackSize ExternalExadataInfrastructureSummaryRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel ExternalExadataInfrastructureSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ExternalExadataInfrastructureSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ExternalExadataInfrastructureSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExternalExadataInfrastructureSummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExternalExadataInfrastructureSummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExternalExadataInfrastructureSummary) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m ExternalExadataInfrastructureSummary) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m ExternalExadataInfrastructureSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ExternalExadataInfrastructureSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m ExternalExadataInfrastructureSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m ExternalExadataInfrastructureSummary) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m ExternalExadataInfrastructureSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataInfrastructureSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalExadataInfrastructureSummaryRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetExternalExadataInfrastructureSummaryRackSizeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalExadataInfrastructureSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalExadataInfrastructureSummaryLicenseModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDbmResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbmResourceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalExadataInfrastructureSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataInfrastructureSummary ExternalExadataInfrastructureSummary
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeExternalExadataInfrastructureSummary
	}{
		"INFRASTRUCTURE_SUMMARY",
		(MarshalTypeExternalExadataInfrastructureSummary)(m),
	}

	return json.Marshal(&s)
}

// ExternalExadataInfrastructureSummaryRackSizeEnum Enum with underlying type: string
type ExternalExadataInfrastructureSummaryRackSizeEnum string

// Set of constants representing the allowable values for ExternalExadataInfrastructureSummaryRackSizeEnum
const (
	ExternalExadataInfrastructureSummaryRackSizeFull    ExternalExadataInfrastructureSummaryRackSizeEnum = "FULL"
	ExternalExadataInfrastructureSummaryRackSizeHalf    ExternalExadataInfrastructureSummaryRackSizeEnum = "HALF"
	ExternalExadataInfrastructureSummaryRackSizeQuarter ExternalExadataInfrastructureSummaryRackSizeEnum = "QUARTER"
	ExternalExadataInfrastructureSummaryRackSizeEighth  ExternalExadataInfrastructureSummaryRackSizeEnum = "EIGHTH"
)

var mappingExternalExadataInfrastructureSummaryRackSizeEnum = map[string]ExternalExadataInfrastructureSummaryRackSizeEnum{
	"FULL":    ExternalExadataInfrastructureSummaryRackSizeFull,
	"HALF":    ExternalExadataInfrastructureSummaryRackSizeHalf,
	"QUARTER": ExternalExadataInfrastructureSummaryRackSizeQuarter,
	"EIGHTH":  ExternalExadataInfrastructureSummaryRackSizeEighth,
}

var mappingExternalExadataInfrastructureSummaryRackSizeEnumLowerCase = map[string]ExternalExadataInfrastructureSummaryRackSizeEnum{
	"full":    ExternalExadataInfrastructureSummaryRackSizeFull,
	"half":    ExternalExadataInfrastructureSummaryRackSizeHalf,
	"quarter": ExternalExadataInfrastructureSummaryRackSizeQuarter,
	"eighth":  ExternalExadataInfrastructureSummaryRackSizeEighth,
}

// GetExternalExadataInfrastructureSummaryRackSizeEnumValues Enumerates the set of values for ExternalExadataInfrastructureSummaryRackSizeEnum
func GetExternalExadataInfrastructureSummaryRackSizeEnumValues() []ExternalExadataInfrastructureSummaryRackSizeEnum {
	values := make([]ExternalExadataInfrastructureSummaryRackSizeEnum, 0)
	for _, v := range mappingExternalExadataInfrastructureSummaryRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataInfrastructureSummaryRackSizeEnumStringValues Enumerates the set of values in String for ExternalExadataInfrastructureSummaryRackSizeEnum
func GetExternalExadataInfrastructureSummaryRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
	}
}

// GetMappingExternalExadataInfrastructureSummaryRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataInfrastructureSummaryRackSizeEnum(val string) (ExternalExadataInfrastructureSummaryRackSizeEnum, bool) {
	enum, ok := mappingExternalExadataInfrastructureSummaryRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalExadataInfrastructureSummaryLicenseModelEnum Enum with underlying type: string
type ExternalExadataInfrastructureSummaryLicenseModelEnum string

// Set of constants representing the allowable values for ExternalExadataInfrastructureSummaryLicenseModelEnum
const (
	ExternalExadataInfrastructureSummaryLicenseModelLicenseIncluded     ExternalExadataInfrastructureSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalExadataInfrastructureSummaryLicenseModelBringYourOwnLicense ExternalExadataInfrastructureSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalExadataInfrastructureSummaryLicenseModelEnum = map[string]ExternalExadataInfrastructureSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalExadataInfrastructureSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalExadataInfrastructureSummaryLicenseModelBringYourOwnLicense,
}

var mappingExternalExadataInfrastructureSummaryLicenseModelEnumLowerCase = map[string]ExternalExadataInfrastructureSummaryLicenseModelEnum{
	"license_included":       ExternalExadataInfrastructureSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalExadataInfrastructureSummaryLicenseModelBringYourOwnLicense,
}

// GetExternalExadataInfrastructureSummaryLicenseModelEnumValues Enumerates the set of values for ExternalExadataInfrastructureSummaryLicenseModelEnum
func GetExternalExadataInfrastructureSummaryLicenseModelEnumValues() []ExternalExadataInfrastructureSummaryLicenseModelEnum {
	values := make([]ExternalExadataInfrastructureSummaryLicenseModelEnum, 0)
	for _, v := range mappingExternalExadataInfrastructureSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataInfrastructureSummaryLicenseModelEnumStringValues Enumerates the set of values in String for ExternalExadataInfrastructureSummaryLicenseModelEnum
func GetExternalExadataInfrastructureSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalExadataInfrastructureSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataInfrastructureSummaryLicenseModelEnum(val string) (ExternalExadataInfrastructureSummaryLicenseModelEnum, bool) {
	enum, ok := mappingExternalExadataInfrastructureSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
