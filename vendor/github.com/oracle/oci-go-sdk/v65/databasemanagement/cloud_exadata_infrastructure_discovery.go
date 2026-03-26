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

// CloudExadataInfrastructureDiscovery The result of the Exadata infrastructure discovery.
type CloudExadataInfrastructureDiscovery struct {

	// The name of the entity.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique key of the discovery request.
	DiscoveryKey *string `mandatory:"true" json:"discoveryKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the entity discovered.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent used for monitoring.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated connector.
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

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The Oracle home path of the Exadata infrastructure.
	GridHomePath *string `mandatory:"false" json:"gridHomePath"`

	// The list of VM Clusters in the Exadata infrastructure.
	VmClusters []VmClusterDiscoverySummary `mandatory:"false" json:"vmClusters"`

	StorageGrid *StorageGridDiscoverySummary `mandatory:"false" json:"storageGrid"`

	// The list of storage servers in the Exadata infrastructure.
	StorageServers []StorageServerDiscoverySummary `mandatory:"false" json:"storageServers"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel CloudExadataInfrastructureDiscoveryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The size of the Exadata infrastructure.
	RackSize CloudExadataInfrastructureDiscoveryRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The status of the entity discovery.
	DiscoverStatus EntityDiscoveredDiscoverStatusEnum `mandatory:"false" json:"discoverStatus,omitempty"`
}

// GetId returns Id
func (m CloudExadataInfrastructureDiscovery) GetId() *string {
	return m.Id
}

// GetAgentId returns AgentId
func (m CloudExadataInfrastructureDiscovery) GetAgentId() *string {
	return m.AgentId
}

// GetConnectorId returns ConnectorId
func (m CloudExadataInfrastructureDiscovery) GetConnectorId() *string {
	return m.ConnectorId
}

// GetDisplayName returns DisplayName
func (m CloudExadataInfrastructureDiscovery) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m CloudExadataInfrastructureDiscovery) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m CloudExadataInfrastructureDiscovery) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m CloudExadataInfrastructureDiscovery) GetStatus() *string {
	return m.Status
}

// GetDiscoverStatus returns DiscoverStatus
func (m CloudExadataInfrastructureDiscovery) GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum {
	return m.DiscoverStatus
}

// GetDiscoverErrorCode returns DiscoverErrorCode
func (m CloudExadataInfrastructureDiscovery) GetDiscoverErrorCode() *string {
	return m.DiscoverErrorCode
}

// GetDiscoverErrorMsg returns DiscoverErrorMsg
func (m CloudExadataInfrastructureDiscovery) GetDiscoverErrorMsg() *string {
	return m.DiscoverErrorMsg
}

func (m CloudExadataInfrastructureDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudExadataInfrastructureDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudExadataInfrastructureDiscoveryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCloudExadataInfrastructureDiscoveryLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudExadataInfrastructureDiscoveryRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetCloudExadataInfrastructureDiscoveryRackSizeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEntityDiscoveredDiscoverStatusEnum(string(m.DiscoverStatus)); !ok && m.DiscoverStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoverStatus: %s. Supported values are: %s.", m.DiscoverStatus, strings.Join(GetEntityDiscoveredDiscoverStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudExadataInfrastructureDiscovery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudExadataInfrastructureDiscovery CloudExadataInfrastructureDiscovery
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeCloudExadataInfrastructureDiscovery
	}{
		"CLOUD_INFRASTRUCTURE_DISCOVER",
		(MarshalTypeCloudExadataInfrastructureDiscovery)(m),
	}

	return json.Marshal(&s)
}

// CloudExadataInfrastructureDiscoveryLicenseModelEnum Enum with underlying type: string
type CloudExadataInfrastructureDiscoveryLicenseModelEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureDiscoveryLicenseModelEnum
const (
	CloudExadataInfrastructureDiscoveryLicenseModelLicenseIncluded     CloudExadataInfrastructureDiscoveryLicenseModelEnum = "LICENSE_INCLUDED"
	CloudExadataInfrastructureDiscoveryLicenseModelBringYourOwnLicense CloudExadataInfrastructureDiscoveryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCloudExadataInfrastructureDiscoveryLicenseModelEnum = map[string]CloudExadataInfrastructureDiscoveryLicenseModelEnum{
	"LICENSE_INCLUDED":       CloudExadataInfrastructureDiscoveryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CloudExadataInfrastructureDiscoveryLicenseModelBringYourOwnLicense,
}

var mappingCloudExadataInfrastructureDiscoveryLicenseModelEnumLowerCase = map[string]CloudExadataInfrastructureDiscoveryLicenseModelEnum{
	"license_included":       CloudExadataInfrastructureDiscoveryLicenseModelLicenseIncluded,
	"bring_your_own_license": CloudExadataInfrastructureDiscoveryLicenseModelBringYourOwnLicense,
}

// GetCloudExadataInfrastructureDiscoveryLicenseModelEnumValues Enumerates the set of values for CloudExadataInfrastructureDiscoveryLicenseModelEnum
func GetCloudExadataInfrastructureDiscoveryLicenseModelEnumValues() []CloudExadataInfrastructureDiscoveryLicenseModelEnum {
	values := make([]CloudExadataInfrastructureDiscoveryLicenseModelEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureDiscoveryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureDiscoveryLicenseModelEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureDiscoveryLicenseModelEnum
func GetCloudExadataInfrastructureDiscoveryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCloudExadataInfrastructureDiscoveryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureDiscoveryLicenseModelEnum(val string) (CloudExadataInfrastructureDiscoveryLicenseModelEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureDiscoveryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudExadataInfrastructureDiscoveryRackSizeEnum Enum with underlying type: string
type CloudExadataInfrastructureDiscoveryRackSizeEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureDiscoveryRackSizeEnum
const (
	CloudExadataInfrastructureDiscoveryRackSizeFull    CloudExadataInfrastructureDiscoveryRackSizeEnum = "FULL"
	CloudExadataInfrastructureDiscoveryRackSizeHalf    CloudExadataInfrastructureDiscoveryRackSizeEnum = "HALF"
	CloudExadataInfrastructureDiscoveryRackSizeQuarter CloudExadataInfrastructureDiscoveryRackSizeEnum = "QUARTER"
	CloudExadataInfrastructureDiscoveryRackSizeEighth  CloudExadataInfrastructureDiscoveryRackSizeEnum = "EIGHTH"
	CloudExadataInfrastructureDiscoveryRackSizeUnknown CloudExadataInfrastructureDiscoveryRackSizeEnum = "UNKNOWN"
)

var mappingCloudExadataInfrastructureDiscoveryRackSizeEnum = map[string]CloudExadataInfrastructureDiscoveryRackSizeEnum{
	"FULL":    CloudExadataInfrastructureDiscoveryRackSizeFull,
	"HALF":    CloudExadataInfrastructureDiscoveryRackSizeHalf,
	"QUARTER": CloudExadataInfrastructureDiscoveryRackSizeQuarter,
	"EIGHTH":  CloudExadataInfrastructureDiscoveryRackSizeEighth,
	"UNKNOWN": CloudExadataInfrastructureDiscoveryRackSizeUnknown,
}

var mappingCloudExadataInfrastructureDiscoveryRackSizeEnumLowerCase = map[string]CloudExadataInfrastructureDiscoveryRackSizeEnum{
	"full":    CloudExadataInfrastructureDiscoveryRackSizeFull,
	"half":    CloudExadataInfrastructureDiscoveryRackSizeHalf,
	"quarter": CloudExadataInfrastructureDiscoveryRackSizeQuarter,
	"eighth":  CloudExadataInfrastructureDiscoveryRackSizeEighth,
	"unknown": CloudExadataInfrastructureDiscoveryRackSizeUnknown,
}

// GetCloudExadataInfrastructureDiscoveryRackSizeEnumValues Enumerates the set of values for CloudExadataInfrastructureDiscoveryRackSizeEnum
func GetCloudExadataInfrastructureDiscoveryRackSizeEnumValues() []CloudExadataInfrastructureDiscoveryRackSizeEnum {
	values := make([]CloudExadataInfrastructureDiscoveryRackSizeEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureDiscoveryRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureDiscoveryRackSizeEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureDiscoveryRackSizeEnum
func GetCloudExadataInfrastructureDiscoveryRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
		"UNKNOWN",
	}
}

// GetMappingCloudExadataInfrastructureDiscoveryRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureDiscoveryRackSizeEnum(val string) (CloudExadataInfrastructureDiscoveryRackSizeEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureDiscoveryRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
