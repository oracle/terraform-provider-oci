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

// ExadataVmClusterSummary The VM Clusters of the Exadata infrastructure.
type ExadataVmClusterSummary struct {

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

	// The Oracle home directory.
	HomeDirectory *string `mandatory:"false" json:"homeDirectory"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel ExadataVmClusterSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The infrastructure deployment type.
	DeploymentType CloudExadataInfrastructureDeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`
}

// GetId returns Id
func (m ExadataVmClusterSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ExadataVmClusterSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExadataVmClusterSummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExadataVmClusterSummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExadataVmClusterSummary) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m ExadataVmClusterSummary) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m ExadataVmClusterSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ExadataVmClusterSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m ExadataVmClusterSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m ExadataVmClusterSummary) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m ExadataVmClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataVmClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadataVmClusterSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExadataVmClusterSummaryLicenseModelEnumStringValues(), ",")))
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
func (m ExadataVmClusterSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadataVmClusterSummary ExadataVmClusterSummary
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeExadataVmClusterSummary
	}{
		"VM_CLUSTER_SUMMARY",
		(MarshalTypeExadataVmClusterSummary)(m),
	}

	return json.Marshal(&s)
}

// ExadataVmClusterSummaryLicenseModelEnum Enum with underlying type: string
type ExadataVmClusterSummaryLicenseModelEnum string

// Set of constants representing the allowable values for ExadataVmClusterSummaryLicenseModelEnum
const (
	ExadataVmClusterSummaryLicenseModelLicenseIncluded     ExadataVmClusterSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	ExadataVmClusterSummaryLicenseModelBringYourOwnLicense ExadataVmClusterSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExadataVmClusterSummaryLicenseModelEnum = map[string]ExadataVmClusterSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       ExadataVmClusterSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExadataVmClusterSummaryLicenseModelBringYourOwnLicense,
}

var mappingExadataVmClusterSummaryLicenseModelEnumLowerCase = map[string]ExadataVmClusterSummaryLicenseModelEnum{
	"license_included":       ExadataVmClusterSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": ExadataVmClusterSummaryLicenseModelBringYourOwnLicense,
}

// GetExadataVmClusterSummaryLicenseModelEnumValues Enumerates the set of values for ExadataVmClusterSummaryLicenseModelEnum
func GetExadataVmClusterSummaryLicenseModelEnumValues() []ExadataVmClusterSummaryLicenseModelEnum {
	values := make([]ExadataVmClusterSummaryLicenseModelEnum, 0)
	for _, v := range mappingExadataVmClusterSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataVmClusterSummaryLicenseModelEnumStringValues Enumerates the set of values in String for ExadataVmClusterSummaryLicenseModelEnum
func GetExadataVmClusterSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExadataVmClusterSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataVmClusterSummaryLicenseModelEnum(val string) (ExadataVmClusterSummaryLicenseModelEnum, bool) {
	enum, ok := mappingExadataVmClusterSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
