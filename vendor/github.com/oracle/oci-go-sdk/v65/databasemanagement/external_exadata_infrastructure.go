// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalExadataInfrastructure The details of the Exadata infrastructure.
type ExternalExadataInfrastructure struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata resource.
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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	StorageGrid *ExternalExadataStorageGridSummary `mandatory:"false" json:"storageGrid"`

	// A list of DB systems.
	DatabaseSystems []ExternalExadataDatabaseSystemSummary `mandatory:"false" json:"databaseSystems"`

	// The list of OCIDs  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartments.
	DatabaseCompartments []string `mandatory:"false" json:"databaseCompartments"`

	// The rack size of the Exadata infrastructure.
	RackSize ExternalExadataInfrastructureRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel ExternalExadataInfrastructureLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ExternalExadataInfrastructure) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ExternalExadataInfrastructure) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExternalExadataInfrastructure) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExternalExadataInfrastructure) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExternalExadataInfrastructure) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m ExternalExadataInfrastructure) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m ExternalExadataInfrastructure) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ExternalExadataInfrastructure) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m ExternalExadataInfrastructure) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m ExternalExadataInfrastructure) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m ExternalExadataInfrastructure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataInfrastructure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalExadataInfrastructureRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetExternalExadataInfrastructureRackSizeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalExadataInfrastructureLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalExadataInfrastructureLicenseModelEnumStringValues(), ",")))
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
func (m ExternalExadataInfrastructure) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataInfrastructure ExternalExadataInfrastructure
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeExternalExadataInfrastructure
	}{
		"INFRASTRUCTURE",
		(MarshalTypeExternalExadataInfrastructure)(m),
	}

	return json.Marshal(&s)
}

// ExternalExadataInfrastructureRackSizeEnum Enum with underlying type: string
type ExternalExadataInfrastructureRackSizeEnum string

// Set of constants representing the allowable values for ExternalExadataInfrastructureRackSizeEnum
const (
	ExternalExadataInfrastructureRackSizeFull    ExternalExadataInfrastructureRackSizeEnum = "FULL"
	ExternalExadataInfrastructureRackSizeHalf    ExternalExadataInfrastructureRackSizeEnum = "HALF"
	ExternalExadataInfrastructureRackSizeQuarter ExternalExadataInfrastructureRackSizeEnum = "QUARTER"
	ExternalExadataInfrastructureRackSizeEighth  ExternalExadataInfrastructureRackSizeEnum = "EIGHTH"
)

var mappingExternalExadataInfrastructureRackSizeEnum = map[string]ExternalExadataInfrastructureRackSizeEnum{
	"FULL":    ExternalExadataInfrastructureRackSizeFull,
	"HALF":    ExternalExadataInfrastructureRackSizeHalf,
	"QUARTER": ExternalExadataInfrastructureRackSizeQuarter,
	"EIGHTH":  ExternalExadataInfrastructureRackSizeEighth,
}

var mappingExternalExadataInfrastructureRackSizeEnumLowerCase = map[string]ExternalExadataInfrastructureRackSizeEnum{
	"full":    ExternalExadataInfrastructureRackSizeFull,
	"half":    ExternalExadataInfrastructureRackSizeHalf,
	"quarter": ExternalExadataInfrastructureRackSizeQuarter,
	"eighth":  ExternalExadataInfrastructureRackSizeEighth,
}

// GetExternalExadataInfrastructureRackSizeEnumValues Enumerates the set of values for ExternalExadataInfrastructureRackSizeEnum
func GetExternalExadataInfrastructureRackSizeEnumValues() []ExternalExadataInfrastructureRackSizeEnum {
	values := make([]ExternalExadataInfrastructureRackSizeEnum, 0)
	for _, v := range mappingExternalExadataInfrastructureRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataInfrastructureRackSizeEnumStringValues Enumerates the set of values in String for ExternalExadataInfrastructureRackSizeEnum
func GetExternalExadataInfrastructureRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
	}
}

// GetMappingExternalExadataInfrastructureRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataInfrastructureRackSizeEnum(val string) (ExternalExadataInfrastructureRackSizeEnum, bool) {
	enum, ok := mappingExternalExadataInfrastructureRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalExadataInfrastructureLicenseModelEnum Enum with underlying type: string
type ExternalExadataInfrastructureLicenseModelEnum string

// Set of constants representing the allowable values for ExternalExadataInfrastructureLicenseModelEnum
const (
	ExternalExadataInfrastructureLicenseModelLicenseIncluded     ExternalExadataInfrastructureLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalExadataInfrastructureLicenseModelBringYourOwnLicense ExternalExadataInfrastructureLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalExadataInfrastructureLicenseModelEnum = map[string]ExternalExadataInfrastructureLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalExadataInfrastructureLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalExadataInfrastructureLicenseModelBringYourOwnLicense,
}

var mappingExternalExadataInfrastructureLicenseModelEnumLowerCase = map[string]ExternalExadataInfrastructureLicenseModelEnum{
	"license_included":       ExternalExadataInfrastructureLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalExadataInfrastructureLicenseModelBringYourOwnLicense,
}

// GetExternalExadataInfrastructureLicenseModelEnumValues Enumerates the set of values for ExternalExadataInfrastructureLicenseModelEnum
func GetExternalExadataInfrastructureLicenseModelEnumValues() []ExternalExadataInfrastructureLicenseModelEnum {
	values := make([]ExternalExadataInfrastructureLicenseModelEnum, 0)
	for _, v := range mappingExternalExadataInfrastructureLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataInfrastructureLicenseModelEnumStringValues Enumerates the set of values in String for ExternalExadataInfrastructureLicenseModelEnum
func GetExternalExadataInfrastructureLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalExadataInfrastructureLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataInfrastructureLicenseModelEnum(val string) (ExternalExadataInfrastructureLicenseModelEnum, bool) {
	enum, ok := mappingExternalExadataInfrastructureLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
