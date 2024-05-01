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

// ExternalExadataDatabaseSystemSummary The DB systems of the Exadata infrastructure.
type ExternalExadataDatabaseSystemSummary struct {

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

	// The Oracle license model that applies to the database management resources.
	LicenseModel ExternalExadataDatabaseSystemSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ExternalExadataDatabaseSystemSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ExternalExadataDatabaseSystemSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExternalExadataDatabaseSystemSummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExternalExadataDatabaseSystemSummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExternalExadataDatabaseSystemSummary) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m ExternalExadataDatabaseSystemSummary) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m ExternalExadataDatabaseSystemSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ExternalExadataDatabaseSystemSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m ExternalExadataDatabaseSystemSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m ExternalExadataDatabaseSystemSummary) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m ExternalExadataDatabaseSystemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataDatabaseSystemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalExadataDatabaseSystemSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalExadataDatabaseSystemSummaryLicenseModelEnumStringValues(), ",")))
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
func (m ExternalExadataDatabaseSystemSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataDatabaseSystemSummary ExternalExadataDatabaseSystemSummary
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeExternalExadataDatabaseSystemSummary
	}{
		"DATABASE_SYSTEM_SUMMARY",
		(MarshalTypeExternalExadataDatabaseSystemSummary)(m),
	}

	return json.Marshal(&s)
}

// ExternalExadataDatabaseSystemSummaryLicenseModelEnum Enum with underlying type: string
type ExternalExadataDatabaseSystemSummaryLicenseModelEnum string

// Set of constants representing the allowable values for ExternalExadataDatabaseSystemSummaryLicenseModelEnum
const (
	ExternalExadataDatabaseSystemSummaryLicenseModelLicenseIncluded     ExternalExadataDatabaseSystemSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalExadataDatabaseSystemSummaryLicenseModelBringYourOwnLicense ExternalExadataDatabaseSystemSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalExadataDatabaseSystemSummaryLicenseModelEnum = map[string]ExternalExadataDatabaseSystemSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalExadataDatabaseSystemSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalExadataDatabaseSystemSummaryLicenseModelBringYourOwnLicense,
}

var mappingExternalExadataDatabaseSystemSummaryLicenseModelEnumLowerCase = map[string]ExternalExadataDatabaseSystemSummaryLicenseModelEnum{
	"license_included":       ExternalExadataDatabaseSystemSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalExadataDatabaseSystemSummaryLicenseModelBringYourOwnLicense,
}

// GetExternalExadataDatabaseSystemSummaryLicenseModelEnumValues Enumerates the set of values for ExternalExadataDatabaseSystemSummaryLicenseModelEnum
func GetExternalExadataDatabaseSystemSummaryLicenseModelEnumValues() []ExternalExadataDatabaseSystemSummaryLicenseModelEnum {
	values := make([]ExternalExadataDatabaseSystemSummaryLicenseModelEnum, 0)
	for _, v := range mappingExternalExadataDatabaseSystemSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataDatabaseSystemSummaryLicenseModelEnumStringValues Enumerates the set of values in String for ExternalExadataDatabaseSystemSummaryLicenseModelEnum
func GetExternalExadataDatabaseSystemSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalExadataDatabaseSystemSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataDatabaseSystemSummaryLicenseModelEnum(val string) (ExternalExadataDatabaseSystemSummaryLicenseModelEnum, bool) {
	enum, ok := mappingExternalExadataDatabaseSystemSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
