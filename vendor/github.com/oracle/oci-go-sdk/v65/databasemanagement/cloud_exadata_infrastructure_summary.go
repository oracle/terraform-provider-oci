// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudExadataInfrastructureSummary The Exadata infrastructure.
type CloudExadataInfrastructureSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
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
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The rack size of the Exadata infrastructure.
	RackSize CloudExadataInfrastructureSummaryRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel CloudExadataInfrastructureSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The type of the Exadata infrastructure.
	InfrastructureType CloudExadataInfrastructureDeploymentTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`
}

// GetId returns Id
func (m CloudExadataInfrastructureSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m CloudExadataInfrastructureSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m CloudExadataInfrastructureSummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m CloudExadataInfrastructureSummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m CloudExadataInfrastructureSummary) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m CloudExadataInfrastructureSummary) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m CloudExadataInfrastructureSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m CloudExadataInfrastructureSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m CloudExadataInfrastructureSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m CloudExadataInfrastructureSummary) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m CloudExadataInfrastructureSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudExadataInfrastructureSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudExadataInfrastructureSummaryRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetCloudExadataInfrastructureSummaryRackSizeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudExadataInfrastructureSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCloudExadataInfrastructureSummaryLicenseModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDbmResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbmResourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudExadataInfrastructureDeploymentTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetCloudExadataInfrastructureDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudExadataInfrastructureSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudExadataInfrastructureSummary CloudExadataInfrastructureSummary
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeCloudExadataInfrastructureSummary
	}{
		"CLOUD_INFRASTRUCTURE_SUMMARY",
		(MarshalTypeCloudExadataInfrastructureSummary)(m),
	}

	return json.Marshal(&s)
}

// CloudExadataInfrastructureSummaryRackSizeEnum Enum with underlying type: string
type CloudExadataInfrastructureSummaryRackSizeEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureSummaryRackSizeEnum
const (
	CloudExadataInfrastructureSummaryRackSizeFull    CloudExadataInfrastructureSummaryRackSizeEnum = "FULL"
	CloudExadataInfrastructureSummaryRackSizeHalf    CloudExadataInfrastructureSummaryRackSizeEnum = "HALF"
	CloudExadataInfrastructureSummaryRackSizeQuarter CloudExadataInfrastructureSummaryRackSizeEnum = "QUARTER"
	CloudExadataInfrastructureSummaryRackSizeEighth  CloudExadataInfrastructureSummaryRackSizeEnum = "EIGHTH"
)

var mappingCloudExadataInfrastructureSummaryRackSizeEnum = map[string]CloudExadataInfrastructureSummaryRackSizeEnum{
	"FULL":    CloudExadataInfrastructureSummaryRackSizeFull,
	"HALF":    CloudExadataInfrastructureSummaryRackSizeHalf,
	"QUARTER": CloudExadataInfrastructureSummaryRackSizeQuarter,
	"EIGHTH":  CloudExadataInfrastructureSummaryRackSizeEighth,
}

var mappingCloudExadataInfrastructureSummaryRackSizeEnumLowerCase = map[string]CloudExadataInfrastructureSummaryRackSizeEnum{
	"full":    CloudExadataInfrastructureSummaryRackSizeFull,
	"half":    CloudExadataInfrastructureSummaryRackSizeHalf,
	"quarter": CloudExadataInfrastructureSummaryRackSizeQuarter,
	"eighth":  CloudExadataInfrastructureSummaryRackSizeEighth,
}

// GetCloudExadataInfrastructureSummaryRackSizeEnumValues Enumerates the set of values for CloudExadataInfrastructureSummaryRackSizeEnum
func GetCloudExadataInfrastructureSummaryRackSizeEnumValues() []CloudExadataInfrastructureSummaryRackSizeEnum {
	values := make([]CloudExadataInfrastructureSummaryRackSizeEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureSummaryRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureSummaryRackSizeEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureSummaryRackSizeEnum
func GetCloudExadataInfrastructureSummaryRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
	}
}

// GetMappingCloudExadataInfrastructureSummaryRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureSummaryRackSizeEnum(val string) (CloudExadataInfrastructureSummaryRackSizeEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureSummaryRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudExadataInfrastructureSummaryLicenseModelEnum Enum with underlying type: string
type CloudExadataInfrastructureSummaryLicenseModelEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureSummaryLicenseModelEnum
const (
	CloudExadataInfrastructureSummaryLicenseModelLicenseIncluded     CloudExadataInfrastructureSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	CloudExadataInfrastructureSummaryLicenseModelBringYourOwnLicense CloudExadataInfrastructureSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCloudExadataInfrastructureSummaryLicenseModelEnum = map[string]CloudExadataInfrastructureSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       CloudExadataInfrastructureSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CloudExadataInfrastructureSummaryLicenseModelBringYourOwnLicense,
}

var mappingCloudExadataInfrastructureSummaryLicenseModelEnumLowerCase = map[string]CloudExadataInfrastructureSummaryLicenseModelEnum{
	"license_included":       CloudExadataInfrastructureSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": CloudExadataInfrastructureSummaryLicenseModelBringYourOwnLicense,
}

// GetCloudExadataInfrastructureSummaryLicenseModelEnumValues Enumerates the set of values for CloudExadataInfrastructureSummaryLicenseModelEnum
func GetCloudExadataInfrastructureSummaryLicenseModelEnumValues() []CloudExadataInfrastructureSummaryLicenseModelEnum {
	values := make([]CloudExadataInfrastructureSummaryLicenseModelEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureSummaryLicenseModelEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureSummaryLicenseModelEnum
func GetCloudExadataInfrastructureSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCloudExadataInfrastructureSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureSummaryLicenseModelEnum(val string) (CloudExadataInfrastructureSummaryLicenseModelEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
