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

// CloudExadataInfrastructure The details of the Exadata infrastructure.
type CloudExadataInfrastructure struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	DisplayName *string `mandatory:"true" json:"displayName"`

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

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	StorageGrid *CloudExadataStorageGridSummary `mandatory:"false" json:"storageGrid"`

	// The list of VM Clusters in the Exadata infrastructure.
	VmClusters []ExadataVmClusterSummary `mandatory:"false" json:"vmClusters"`

	// The list of OCIDs  (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartments.
	DatabaseCompartments []string `mandatory:"false" json:"databaseCompartments"`

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
	RackSize CloudExadataInfrastructureRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel CloudExadataInfrastructureLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The infrastructure deployment type.
	DeploymentType CloudExadataInfrastructureDeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`
}

// GetId returns Id
func (m CloudExadataInfrastructure) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m CloudExadataInfrastructure) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m CloudExadataInfrastructure) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m CloudExadataInfrastructure) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m CloudExadataInfrastructure) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m CloudExadataInfrastructure) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m CloudExadataInfrastructure) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m CloudExadataInfrastructure) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m CloudExadataInfrastructure) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m CloudExadataInfrastructure) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m CloudExadataInfrastructure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudExadataInfrastructure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudExadataInfrastructureRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetCloudExadataInfrastructureRackSizeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudExadataInfrastructureLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCloudExadataInfrastructureLicenseModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDbmResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbmResourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudExadataInfrastructureDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetCloudExadataInfrastructureDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudExadataInfrastructure) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudExadataInfrastructure CloudExadataInfrastructure
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeCloudExadataInfrastructure
	}{
		"CLOUD_INFRASTRUCTURE",
		(MarshalTypeCloudExadataInfrastructure)(m),
	}

	return json.Marshal(&s)
}

// CloudExadataInfrastructureRackSizeEnum Enum with underlying type: string
type CloudExadataInfrastructureRackSizeEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureRackSizeEnum
const (
	CloudExadataInfrastructureRackSizeFull    CloudExadataInfrastructureRackSizeEnum = "FULL"
	CloudExadataInfrastructureRackSizeHalf    CloudExadataInfrastructureRackSizeEnum = "HALF"
	CloudExadataInfrastructureRackSizeQuarter CloudExadataInfrastructureRackSizeEnum = "QUARTER"
	CloudExadataInfrastructureRackSizeEighth  CloudExadataInfrastructureRackSizeEnum = "EIGHTH"
)

var mappingCloudExadataInfrastructureRackSizeEnum = map[string]CloudExadataInfrastructureRackSizeEnum{
	"FULL":    CloudExadataInfrastructureRackSizeFull,
	"HALF":    CloudExadataInfrastructureRackSizeHalf,
	"QUARTER": CloudExadataInfrastructureRackSizeQuarter,
	"EIGHTH":  CloudExadataInfrastructureRackSizeEighth,
}

var mappingCloudExadataInfrastructureRackSizeEnumLowerCase = map[string]CloudExadataInfrastructureRackSizeEnum{
	"full":    CloudExadataInfrastructureRackSizeFull,
	"half":    CloudExadataInfrastructureRackSizeHalf,
	"quarter": CloudExadataInfrastructureRackSizeQuarter,
	"eighth":  CloudExadataInfrastructureRackSizeEighth,
}

// GetCloudExadataInfrastructureRackSizeEnumValues Enumerates the set of values for CloudExadataInfrastructureRackSizeEnum
func GetCloudExadataInfrastructureRackSizeEnumValues() []CloudExadataInfrastructureRackSizeEnum {
	values := make([]CloudExadataInfrastructureRackSizeEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureRackSizeEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureRackSizeEnum
func GetCloudExadataInfrastructureRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
	}
}

// GetMappingCloudExadataInfrastructureRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureRackSizeEnum(val string) (CloudExadataInfrastructureRackSizeEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudExadataInfrastructureLicenseModelEnum Enum with underlying type: string
type CloudExadataInfrastructureLicenseModelEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureLicenseModelEnum
const (
	CloudExadataInfrastructureLicenseModelLicenseIncluded     CloudExadataInfrastructureLicenseModelEnum = "LICENSE_INCLUDED"
	CloudExadataInfrastructureLicenseModelBringYourOwnLicense CloudExadataInfrastructureLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCloudExadataInfrastructureLicenseModelEnum = map[string]CloudExadataInfrastructureLicenseModelEnum{
	"LICENSE_INCLUDED":       CloudExadataInfrastructureLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CloudExadataInfrastructureLicenseModelBringYourOwnLicense,
}

var mappingCloudExadataInfrastructureLicenseModelEnumLowerCase = map[string]CloudExadataInfrastructureLicenseModelEnum{
	"license_included":       CloudExadataInfrastructureLicenseModelLicenseIncluded,
	"bring_your_own_license": CloudExadataInfrastructureLicenseModelBringYourOwnLicense,
}

// GetCloudExadataInfrastructureLicenseModelEnumValues Enumerates the set of values for CloudExadataInfrastructureLicenseModelEnum
func GetCloudExadataInfrastructureLicenseModelEnumValues() []CloudExadataInfrastructureLicenseModelEnum {
	values := make([]CloudExadataInfrastructureLicenseModelEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureLicenseModelEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureLicenseModelEnum
func GetCloudExadataInfrastructureLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCloudExadataInfrastructureLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureLicenseModelEnum(val string) (CloudExadataInfrastructureLicenseModelEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
