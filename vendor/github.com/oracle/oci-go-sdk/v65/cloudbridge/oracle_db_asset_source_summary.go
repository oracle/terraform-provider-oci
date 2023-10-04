// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbAssetSourceSummary Summary of an Oracle database asset source provided in the list.
type OracleDbAssetSourceSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment for the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the environment.
	EnvironmentId *string `mandatory:"true" json:"environmentId"`

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The detailed state of the asset source.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the inventory that will contain created assets.
	InventoryId *string `mandatory:"true" json:"inventoryId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	AssetsCompartmentId *string `mandatory:"true" json:"assetsCompartmentId"`

	// Database server endpoint.
	Host *string `mandatory:"true" json:"host"`

	// Database port number.
	Port *int `mandatory:"true" json:"port"`

	// Database service name.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The time when the asset source was created in RFC3339 format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The point in time that the asset source was last updated in RFC3339 format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the asset source.
	LifecycleState AssetSourceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Database server connection protocol.
	Protocol DatabaseNetworkProtocolEnum `mandatory:"true" json:"protocol"`
}

//GetId returns Id
func (m OracleDbAssetSourceSummary) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m OracleDbAssetSourceSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetEnvironmentId returns EnvironmentId
func (m OracleDbAssetSourceSummary) GetEnvironmentId() *string {
	return m.EnvironmentId
}

//GetDisplayName returns DisplayName
func (m OracleDbAssetSourceSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetLifecycleState returns LifecycleState
func (m OracleDbAssetSourceSummary) GetLifecycleState() AssetSourceLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m OracleDbAssetSourceSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetInventoryId returns InventoryId
func (m OracleDbAssetSourceSummary) GetInventoryId() *string {
	return m.InventoryId
}

//GetAssetsCompartmentId returns AssetsCompartmentId
func (m OracleDbAssetSourceSummary) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

//GetTimeCreated returns TimeCreated
func (m OracleDbAssetSourceSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m OracleDbAssetSourceSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetFreeformTags returns FreeformTags
func (m OracleDbAssetSourceSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m OracleDbAssetSourceSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m OracleDbAssetSourceSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m OracleDbAssetSourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbAssetSourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAssetSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssetSourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseNetworkProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetDatabaseNetworkProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleDbAssetSourceSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleDbAssetSourceSummary OracleDbAssetSourceSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOracleDbAssetSourceSummary
	}{
		"ORACLE_DB",
		(MarshalTypeOracleDbAssetSourceSummary)(m),
	}

	return json.Marshal(&s)
}
