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

// ExternalDatabaseSystemDiscoverySummary The summary of the DB system discovery.
type ExternalDatabaseSystemDiscoverySummary struct {

	// The name of the entity.
	DisplayName *string `mandatory:"true" json:"displayName"`

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

	// The Oracle home path.
	OracleHome *string `mandatory:"false" json:"oracleHome"`

	// The display name of the ASM connector.
	AsmConnectorName *string `mandatory:"false" json:"asmConnectorName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The status of the entity discovery.
	DiscoverStatus EntityDiscoveredDiscoverStatusEnum `mandatory:"false" json:"discoverStatus,omitempty"`
}

// GetId returns Id
func (m ExternalDatabaseSystemDiscoverySummary) GetId() *string {
	return m.Id
}

// GetAgentId returns AgentId
func (m ExternalDatabaseSystemDiscoverySummary) GetAgentId() *string {
	return m.AgentId
}

// GetConnectorId returns ConnectorId
func (m ExternalDatabaseSystemDiscoverySummary) GetConnectorId() *string {
	return m.ConnectorId
}

// GetDisplayName returns DisplayName
func (m ExternalDatabaseSystemDiscoverySummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExternalDatabaseSystemDiscoverySummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExternalDatabaseSystemDiscoverySummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExternalDatabaseSystemDiscoverySummary) GetStatus() *string {
	return m.Status
}

// GetDiscoverStatus returns DiscoverStatus
func (m ExternalDatabaseSystemDiscoverySummary) GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum {
	return m.DiscoverStatus
}

// GetDiscoverErrorCode returns DiscoverErrorCode
func (m ExternalDatabaseSystemDiscoverySummary) GetDiscoverErrorCode() *string {
	return m.DiscoverErrorCode
}

// GetDiscoverErrorMsg returns DiscoverErrorMsg
func (m ExternalDatabaseSystemDiscoverySummary) GetDiscoverErrorMsg() *string {
	return m.DiscoverErrorMsg
}

func (m ExternalDatabaseSystemDiscoverySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseSystemDiscoverySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDatabaseSystemDiscoverySummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalDatabaseSystemDiscoverySummaryLicenseModelEnumStringValues(), ",")))
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
func (m ExternalDatabaseSystemDiscoverySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalDatabaseSystemDiscoverySummary ExternalDatabaseSystemDiscoverySummary
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeExternalDatabaseSystemDiscoverySummary
	}{
		"DATABASE_SYSTEM_DISCOVER_SUMMARY",
		(MarshalTypeExternalDatabaseSystemDiscoverySummary)(m),
	}

	return json.Marshal(&s)
}

// ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum Enum with underlying type: string
type ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum string

// Set of constants representing the allowable values for ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum
const (
	ExternalDatabaseSystemDiscoverySummaryLicenseModelLicenseIncluded     ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalDatabaseSystemDiscoverySummaryLicenseModelBringYourOwnLicense ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalDatabaseSystemDiscoverySummaryLicenseModelEnum = map[string]ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalDatabaseSystemDiscoverySummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalDatabaseSystemDiscoverySummaryLicenseModelBringYourOwnLicense,
}

var mappingExternalDatabaseSystemDiscoverySummaryLicenseModelEnumLowerCase = map[string]ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum{
	"license_included":       ExternalDatabaseSystemDiscoverySummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalDatabaseSystemDiscoverySummaryLicenseModelBringYourOwnLicense,
}

// GetExternalDatabaseSystemDiscoverySummaryLicenseModelEnumValues Enumerates the set of values for ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum
func GetExternalDatabaseSystemDiscoverySummaryLicenseModelEnumValues() []ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum {
	values := make([]ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum, 0)
	for _, v := range mappingExternalDatabaseSystemDiscoverySummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseSystemDiscoverySummaryLicenseModelEnumStringValues Enumerates the set of values in String for ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum
func GetExternalDatabaseSystemDiscoverySummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalDatabaseSystemDiscoverySummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseSystemDiscoverySummaryLicenseModelEnum(val string) (ExternalDatabaseSystemDiscoverySummaryLicenseModelEnum, bool) {
	enum, ok := mappingExternalDatabaseSystemDiscoverySummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
