// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectionSummary Represents the parameters common for all connections in list operation.
type ConnectionSummary interface {

	// The OCID of the connection being referenced.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the compartment.
	GetCompartmentId() *string

	// Lifecycle state for connection.
	GetLifecycleState() ConnectionLifecycleStateEnum

	// The time when this resource was created.
	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// The time when this resource was updated.
	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeUpdated() *common.SDKTime

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDescription() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// A message describing the current state in more detail. For example, can be used to provide actionable information
	// for a resource in Failed state.
	GetLifecycleDetails() *string

	// OCI resource ID.
	GetVaultId() *string

	// The OCID of the key used in cryptographic operations.
	GetKeyId() *string

	// OCI resource ID.
	GetSubnetId() *string

	// List of ingress IP addresses from where to connect to this connection's privateIp.
	GetIngressIps() []IngressIpDetails

	// An array of Network Security Group OCIDs used to define network access for Connections.
	GetNsgIds() []string
}

type connectionsummary struct {
	JsonData         []byte
	Description      *string                           `mandatory:"false" json:"description"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	VaultId          *string                           `mandatory:"false" json:"vaultId"`
	KeyId            *string                           `mandatory:"false" json:"keyId"`
	SubnetId         *string                           `mandatory:"false" json:"subnetId"`
	IngressIps       []IngressIpDetails                `mandatory:"false" json:"ingressIps"`
	NsgIds           []string                          `mandatory:"false" json:"nsgIds"`
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState   ConnectionLifecycleStateEnum      `mandatory:"true" json:"lifecycleState"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	ConnectionType   string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *connectionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectionsummary connectionsummary
	s := struct {
		Model Unmarshalerconnectionsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.VaultId = s.Model.VaultId
	m.KeyId = s.Model.KeyId
	m.SubnetId = s.Model.SubnetId
	m.IngressIps = s.Model.IngressIps
	m.NsgIds = s.Model.NsgIds
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "ORACLE":
		mm := OracleConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := MysqlConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConnectionSummary: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m connectionsummary) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m connectionsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m connectionsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m connectionsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleDetails returns LifecycleDetails
func (m connectionsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetVaultId returns VaultId
func (m connectionsummary) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m connectionsummary) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m connectionsummary) GetSubnetId() *string {
	return m.SubnetId
}

// GetIngressIps returns IngressIps
func (m connectionsummary) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m connectionsummary) GetNsgIds() []string {
	return m.NsgIds
}

// GetId returns Id
func (m connectionsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m connectionsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m connectionsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m connectionsummary) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m connectionsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m connectionsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m connectionsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectionsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
