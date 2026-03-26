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

// ExadataInfrastructureDiscoverySummary The summary of the Exadata system infrastructure discovery.
type ExadataInfrastructureDiscoverySummary struct {

	// The name of the entity.
	DisplayName *string `mandatory:"true" json:"displayName"`

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

	// The size of the Exadata infrastructure.
	RackSize ExadataInfrastructureDiscoverySummaryRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The status of the entity discovery.
	DiscoverStatus EntityDiscoveredDiscoverStatusEnum `mandatory:"false" json:"discoverStatus,omitempty"`
}

// GetId returns Id
func (m ExadataInfrastructureDiscoverySummary) GetId() *string {
	return m.Id
}

// GetAgentId returns AgentId
func (m ExadataInfrastructureDiscoverySummary) GetAgentId() *string {
	return m.AgentId
}

// GetConnectorId returns ConnectorId
func (m ExadataInfrastructureDiscoverySummary) GetConnectorId() *string {
	return m.ConnectorId
}

// GetDisplayName returns DisplayName
func (m ExadataInfrastructureDiscoverySummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExadataInfrastructureDiscoverySummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExadataInfrastructureDiscoverySummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExadataInfrastructureDiscoverySummary) GetStatus() *string {
	return m.Status
}

// GetDiscoverStatus returns DiscoverStatus
func (m ExadataInfrastructureDiscoverySummary) GetDiscoverStatus() EntityDiscoveredDiscoverStatusEnum {
	return m.DiscoverStatus
}

// GetDiscoverErrorCode returns DiscoverErrorCode
func (m ExadataInfrastructureDiscoverySummary) GetDiscoverErrorCode() *string {
	return m.DiscoverErrorCode
}

// GetDiscoverErrorMsg returns DiscoverErrorMsg
func (m ExadataInfrastructureDiscoverySummary) GetDiscoverErrorMsg() *string {
	return m.DiscoverErrorMsg
}

func (m ExadataInfrastructureDiscoverySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructureDiscoverySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadataInfrastructureDiscoverySummaryRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetExadataInfrastructureDiscoverySummaryRackSizeEnumStringValues(), ",")))
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
func (m ExadataInfrastructureDiscoverySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadataInfrastructureDiscoverySummary ExadataInfrastructureDiscoverySummary
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeExadataInfrastructureDiscoverySummary
	}{
		"MANAGED_INFRASTRUCTURE_DISCOVER_SUMMARY",
		(MarshalTypeExadataInfrastructureDiscoverySummary)(m),
	}

	return json.Marshal(&s)
}

// ExadataInfrastructureDiscoverySummaryRackSizeEnum Enum with underlying type: string
type ExadataInfrastructureDiscoverySummaryRackSizeEnum string

// Set of constants representing the allowable values for ExadataInfrastructureDiscoverySummaryRackSizeEnum
const (
	ExadataInfrastructureDiscoverySummaryRackSizeFull    ExadataInfrastructureDiscoverySummaryRackSizeEnum = "FULL"
	ExadataInfrastructureDiscoverySummaryRackSizeHalf    ExadataInfrastructureDiscoverySummaryRackSizeEnum = "HALF"
	ExadataInfrastructureDiscoverySummaryRackSizeQuarter ExadataInfrastructureDiscoverySummaryRackSizeEnum = "QUARTER"
	ExadataInfrastructureDiscoverySummaryRackSizeEighth  ExadataInfrastructureDiscoverySummaryRackSizeEnum = "EIGHTH"
)

var mappingExadataInfrastructureDiscoverySummaryRackSizeEnum = map[string]ExadataInfrastructureDiscoverySummaryRackSizeEnum{
	"FULL":    ExadataInfrastructureDiscoverySummaryRackSizeFull,
	"HALF":    ExadataInfrastructureDiscoverySummaryRackSizeHalf,
	"QUARTER": ExadataInfrastructureDiscoverySummaryRackSizeQuarter,
	"EIGHTH":  ExadataInfrastructureDiscoverySummaryRackSizeEighth,
}

var mappingExadataInfrastructureDiscoverySummaryRackSizeEnumLowerCase = map[string]ExadataInfrastructureDiscoverySummaryRackSizeEnum{
	"full":    ExadataInfrastructureDiscoverySummaryRackSizeFull,
	"half":    ExadataInfrastructureDiscoverySummaryRackSizeHalf,
	"quarter": ExadataInfrastructureDiscoverySummaryRackSizeQuarter,
	"eighth":  ExadataInfrastructureDiscoverySummaryRackSizeEighth,
}

// GetExadataInfrastructureDiscoverySummaryRackSizeEnumValues Enumerates the set of values for ExadataInfrastructureDiscoverySummaryRackSizeEnum
func GetExadataInfrastructureDiscoverySummaryRackSizeEnumValues() []ExadataInfrastructureDiscoverySummaryRackSizeEnum {
	values := make([]ExadataInfrastructureDiscoverySummaryRackSizeEnum, 0)
	for _, v := range mappingExadataInfrastructureDiscoverySummaryRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureDiscoverySummaryRackSizeEnumStringValues Enumerates the set of values in String for ExadataInfrastructureDiscoverySummaryRackSizeEnum
func GetExadataInfrastructureDiscoverySummaryRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
	}
}

// GetMappingExadataInfrastructureDiscoverySummaryRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureDiscoverySummaryRackSizeEnum(val string) (ExadataInfrastructureDiscoverySummaryRackSizeEnum, bool) {
	enum, ok := mappingExadataInfrastructureDiscoverySummaryRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
