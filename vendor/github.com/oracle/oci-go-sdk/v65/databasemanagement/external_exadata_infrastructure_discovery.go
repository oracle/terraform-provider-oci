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

// ExternalExadataInfrastructureDiscovery The result of the Exadata infrastructure discovery.
type ExternalExadataInfrastructureDiscovery struct {

	// The name of the entity.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique key of the discovery request.
	DiscoveryKey *string `mandatory:"true" json:"discoveryKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the entity discovered.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the agent used for monitoring.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated connector.
	ConnectorId *string `mandatory:"false" json:"connectorId"`

	// The version of the entity.
	Version *string `mandatory:"false" json:"version"`

	// The internal identifier of the entity.
	InternalId *string `mandatory:"false" json:"internalId"`

	// The status of the entity.
	Status *string `mandatory:"false" json:"status"`

	// The error code of the discovery.
	DiscoverErrorCode *string `mandatory:"false" json:"discoverErrorCode"`

	// The error message of the discovery.
	DiscoverErrorMsg *string `mandatory:"false" json:"discoverErrorMsg"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The Oracle home path of the Exadata infrastructure.
	GridHomePath *string `mandatory:"false" json:"gridHomePath"`

	// The list of DB systems in the Exadata infrastructure.
	DbSystems []ExternalDatabaseSystemDiscoverySummary `mandatory:"false" json:"dbSystems"`

	StorageGrid *ExternalStorageGridDiscoverySummary `mandatory:"false" json:"storageGrid"`

	// The list of storage servers in the Exadata infrastructure.
	StorageServers []ExternalStorageServerDiscoverySummary `mandatory:"false" json:"storageServers"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel ExternalExadataInfrastructureDiscoveryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The size of the Exadata infrastructure.
	RackSize ExternalExadataInfrastructureDiscoveryRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The status of the entity discovery.
	DiscoverStatus EntityDiscoveredDiscoverStatusEnum `mandatory:"false" json:"discoverStatus,omitempty"`
}

// GetId returns Id
func (m ExternalExadataInfrastructureDiscovery) GetId() *string {
	return m.Id
}

// GetAgentId returns AgentId
func (m ExternalExadataInfrastructureDiscovery) GetAgentId() *string {
	return m.AgentId
}

// GetConnectorId returns ConnectorId
func (m ExternalExadataInfrastructureDiscovery) GetConnectorId() *string {
	return m.ConnectorId
}

// GetDisplayName returns DisplayName
func (m ExternalExadataInfrastructureDiscovery) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExternalExadataInfrastructureDiscovery) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExternalExadataInfrastructureDiscovery) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExternalExadataInfrastructureDiscovery) GetStatus() *string {
	return m.Status
}

// GetDiscoverStatus returns DiscoverStatus
func (m ExternalExadataInfrastructureDiscovery) GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum {
	return m.DiscoverStatus
}

// GetDiscoverErrorCode returns DiscoverErrorCode
func (m ExternalExadataInfrastructureDiscovery) GetDiscoverErrorCode() *string {
	return m.DiscoverErrorCode
}

// GetDiscoverErrorMsg returns DiscoverErrorMsg
func (m ExternalExadataInfrastructureDiscovery) GetDiscoverErrorMsg() *string {
	return m.DiscoverErrorMsg
}

func (m ExternalExadataInfrastructureDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataInfrastructureDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalExadataInfrastructureDiscoveryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalExadataInfrastructureDiscoveryLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalExadataInfrastructureDiscoveryRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetExternalExadataInfrastructureDiscoveryRackSizeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEntityDiscoveredDiscoverStatusEnum(string(m.DiscoverStatus)); !ok && m.DiscoverStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoverStatus: %s. Supported values are: %s.", m.DiscoverStatus, strings.Join(GetEntityDiscoveredDiscoverStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalExadataInfrastructureDiscovery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataInfrastructureDiscovery ExternalExadataInfrastructureDiscovery
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeExternalExadataInfrastructureDiscovery
	}{
		"INFRASTRUCTURE_DISCOVER",
		(MarshalTypeExternalExadataInfrastructureDiscovery)(m),
	}

	return json.Marshal(&s)
}

// ExternalExadataInfrastructureDiscoveryLicenseModelEnum Enum with underlying type: string
type ExternalExadataInfrastructureDiscoveryLicenseModelEnum string

// Set of constants representing the allowable values for ExternalExadataInfrastructureDiscoveryLicenseModelEnum
const (
	ExternalExadataInfrastructureDiscoveryLicenseModelLicenseIncluded     ExternalExadataInfrastructureDiscoveryLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalExadataInfrastructureDiscoveryLicenseModelBringYourOwnLicense ExternalExadataInfrastructureDiscoveryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalExadataInfrastructureDiscoveryLicenseModelEnum = map[string]ExternalExadataInfrastructureDiscoveryLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalExadataInfrastructureDiscoveryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalExadataInfrastructureDiscoveryLicenseModelBringYourOwnLicense,
}

var mappingExternalExadataInfrastructureDiscoveryLicenseModelEnumLowerCase = map[string]ExternalExadataInfrastructureDiscoveryLicenseModelEnum{
	"license_included":       ExternalExadataInfrastructureDiscoveryLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalExadataInfrastructureDiscoveryLicenseModelBringYourOwnLicense,
}

// GetExternalExadataInfrastructureDiscoveryLicenseModelEnumValues Enumerates the set of values for ExternalExadataInfrastructureDiscoveryLicenseModelEnum
func GetExternalExadataInfrastructureDiscoveryLicenseModelEnumValues() []ExternalExadataInfrastructureDiscoveryLicenseModelEnum {
	values := make([]ExternalExadataInfrastructureDiscoveryLicenseModelEnum, 0)
	for _, v := range mappingExternalExadataInfrastructureDiscoveryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataInfrastructureDiscoveryLicenseModelEnumStringValues Enumerates the set of values in String for ExternalExadataInfrastructureDiscoveryLicenseModelEnum
func GetExternalExadataInfrastructureDiscoveryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalExadataInfrastructureDiscoveryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataInfrastructureDiscoveryLicenseModelEnum(val string) (ExternalExadataInfrastructureDiscoveryLicenseModelEnum, bool) {
	enum, ok := mappingExternalExadataInfrastructureDiscoveryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalExadataInfrastructureDiscoveryRackSizeEnum Enum with underlying type: string
type ExternalExadataInfrastructureDiscoveryRackSizeEnum string

// Set of constants representing the allowable values for ExternalExadataInfrastructureDiscoveryRackSizeEnum
const (
	ExternalExadataInfrastructureDiscoveryRackSizeFull    ExternalExadataInfrastructureDiscoveryRackSizeEnum = "FULL"
	ExternalExadataInfrastructureDiscoveryRackSizeHalf    ExternalExadataInfrastructureDiscoveryRackSizeEnum = "HALF"
	ExternalExadataInfrastructureDiscoveryRackSizeQuarter ExternalExadataInfrastructureDiscoveryRackSizeEnum = "QUARTER"
	ExternalExadataInfrastructureDiscoveryRackSizeEighth  ExternalExadataInfrastructureDiscoveryRackSizeEnum = "EIGHTH"
	ExternalExadataInfrastructureDiscoveryRackSizeUnknown ExternalExadataInfrastructureDiscoveryRackSizeEnum = "UNKNOWN"
)

var mappingExternalExadataInfrastructureDiscoveryRackSizeEnum = map[string]ExternalExadataInfrastructureDiscoveryRackSizeEnum{
	"FULL":    ExternalExadataInfrastructureDiscoveryRackSizeFull,
	"HALF":    ExternalExadataInfrastructureDiscoveryRackSizeHalf,
	"QUARTER": ExternalExadataInfrastructureDiscoveryRackSizeQuarter,
	"EIGHTH":  ExternalExadataInfrastructureDiscoveryRackSizeEighth,
	"UNKNOWN": ExternalExadataInfrastructureDiscoveryRackSizeUnknown,
}

var mappingExternalExadataInfrastructureDiscoveryRackSizeEnumLowerCase = map[string]ExternalExadataInfrastructureDiscoveryRackSizeEnum{
	"full":    ExternalExadataInfrastructureDiscoveryRackSizeFull,
	"half":    ExternalExadataInfrastructureDiscoveryRackSizeHalf,
	"quarter": ExternalExadataInfrastructureDiscoveryRackSizeQuarter,
	"eighth":  ExternalExadataInfrastructureDiscoveryRackSizeEighth,
	"unknown": ExternalExadataInfrastructureDiscoveryRackSizeUnknown,
}

// GetExternalExadataInfrastructureDiscoveryRackSizeEnumValues Enumerates the set of values for ExternalExadataInfrastructureDiscoveryRackSizeEnum
func GetExternalExadataInfrastructureDiscoveryRackSizeEnumValues() []ExternalExadataInfrastructureDiscoveryRackSizeEnum {
	values := make([]ExternalExadataInfrastructureDiscoveryRackSizeEnum, 0)
	for _, v := range mappingExternalExadataInfrastructureDiscoveryRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataInfrastructureDiscoveryRackSizeEnumStringValues Enumerates the set of values in String for ExternalExadataInfrastructureDiscoveryRackSizeEnum
func GetExternalExadataInfrastructureDiscoveryRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
		"UNKNOWN",
	}
}

// GetMappingExternalExadataInfrastructureDiscoveryRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataInfrastructureDiscoveryRackSizeEnum(val string) (ExternalExadataInfrastructureDiscoveryRackSizeEnum, bool) {
	enum, ok := mappingExternalExadataInfrastructureDiscoveryRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
