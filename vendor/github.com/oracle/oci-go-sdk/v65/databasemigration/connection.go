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

// Connection Represents the common details required for creating a new connection.
type Connection interface {

	// The OCID of the connection being referenced.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the compartment.
	GetCompartmentId() *string

	// The Connection's current lifecycle state.
	GetLifecycleState() ConnectionLifecycleStateEnum

	// The time when this resource was created.
	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// The time when this resource was updated.
	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeUpdated() *common.SDKTime

	// The username (credential) used when creating or updating this resource.
	GetUsername() *string

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

	// The message describing the current state of the connection's lifecycle in detail.
	// For example, can be used to provide actionable information for a connection in a Failed state.
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

	// The password (credential) used when creating or updating this resource.
	GetPassword() *string

	// The username (credential) used when creating or updating this resource.
	GetReplicationUsername() *string

	// The password (credential) used when creating or updating this resource.
	GetReplicationPassword() *string

	// The OCID of the resource being referenced.
	GetSecretId() *string

	// The OCID of the resource being referenced.
	GetPrivateEndpointId() *string
}

type connection struct {
	JsonData            []byte
	Description         *string                           `mandatory:"false" json:"description"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	LifecycleDetails    *string                           `mandatory:"false" json:"lifecycleDetails"`
	VaultId             *string                           `mandatory:"false" json:"vaultId"`
	KeyId               *string                           `mandatory:"false" json:"keyId"`
	SubnetId            *string                           `mandatory:"false" json:"subnetId"`
	IngressIps          []IngressIpDetails                `mandatory:"false" json:"ingressIps"`
	NsgIds              []string                          `mandatory:"false" json:"nsgIds"`
	Password            *string                           `mandatory:"false" json:"password"`
	ReplicationUsername *string                           `mandatory:"false" json:"replicationUsername"`
	ReplicationPassword *string                           `mandatory:"false" json:"replicationPassword"`
	SecretId            *string                           `mandatory:"false" json:"secretId"`
	PrivateEndpointId   *string                           `mandatory:"false" json:"privateEndpointId"`
	Id                  *string                           `mandatory:"true" json:"id"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState      ConnectionLifecycleStateEnum      `mandatory:"true" json:"lifecycleState"`
	TimeCreated         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated         *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	Username            *string                           `mandatory:"true" json:"username"`
	ConnectionType      string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *connection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnection connection
	s := struct {
		Model Unmarshalerconnection
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
	m.Username = s.Model.Username
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
	m.Password = s.Model.Password
	m.ReplicationUsername = s.Model.ReplicationUsername
	m.ReplicationPassword = s.Model.ReplicationPassword
	m.SecretId = s.Model.SecretId
	m.PrivateEndpointId = s.Model.PrivateEndpointId
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "MYSQL":
		mm := MysqlConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := OracleConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Connection: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m connection) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m connection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m connection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m connection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleDetails returns LifecycleDetails
func (m connection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetVaultId returns VaultId
func (m connection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m connection) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m connection) GetSubnetId() *string {
	return m.SubnetId
}

// GetIngressIps returns IngressIps
func (m connection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m connection) GetNsgIds() []string {
	return m.NsgIds
}

// GetPassword returns Password
func (m connection) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m connection) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m connection) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

// GetSecretId returns SecretId
func (m connection) GetSecretId() *string {
	return m.SecretId
}

// GetPrivateEndpointId returns PrivateEndpointId
func (m connection) GetPrivateEndpointId() *string {
	return m.PrivateEndpointId
}

// GetId returns Id
func (m connection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m connection) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m connection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m connection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m connection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m connection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetUsername returns Username
func (m connection) GetUsername() *string {
	return m.Username
}

func (m connection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionLifecycleStateEnum Enum with underlying type: string
type ConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for ConnectionLifecycleStateEnum
const (
	ConnectionLifecycleStateCreating ConnectionLifecycleStateEnum = "CREATING"
	ConnectionLifecycleStateUpdating ConnectionLifecycleStateEnum = "UPDATING"
	ConnectionLifecycleStateActive   ConnectionLifecycleStateEnum = "ACTIVE"
	ConnectionLifecycleStateInactive ConnectionLifecycleStateEnum = "INACTIVE"
	ConnectionLifecycleStateDeleting ConnectionLifecycleStateEnum = "DELETING"
	ConnectionLifecycleStateDeleted  ConnectionLifecycleStateEnum = "DELETED"
	ConnectionLifecycleStateFailed   ConnectionLifecycleStateEnum = "FAILED"
)

var mappingConnectionLifecycleStateEnum = map[string]ConnectionLifecycleStateEnum{
	"CREATING": ConnectionLifecycleStateCreating,
	"UPDATING": ConnectionLifecycleStateUpdating,
	"ACTIVE":   ConnectionLifecycleStateActive,
	"INACTIVE": ConnectionLifecycleStateInactive,
	"DELETING": ConnectionLifecycleStateDeleting,
	"DELETED":  ConnectionLifecycleStateDeleted,
	"FAILED":   ConnectionLifecycleStateFailed,
}

var mappingConnectionLifecycleStateEnumLowerCase = map[string]ConnectionLifecycleStateEnum{
	"creating": ConnectionLifecycleStateCreating,
	"updating": ConnectionLifecycleStateUpdating,
	"active":   ConnectionLifecycleStateActive,
	"inactive": ConnectionLifecycleStateInactive,
	"deleting": ConnectionLifecycleStateDeleting,
	"deleted":  ConnectionLifecycleStateDeleted,
	"failed":   ConnectionLifecycleStateFailed,
}

// GetConnectionLifecycleStateEnumValues Enumerates the set of values for ConnectionLifecycleStateEnum
func GetConnectionLifecycleStateEnumValues() []ConnectionLifecycleStateEnum {
	values := make([]ConnectionLifecycleStateEnum, 0)
	for _, v := range mappingConnectionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionLifecycleStateEnumStringValues Enumerates the set of values in String for ConnectionLifecycleStateEnum
func GetConnectionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingConnectionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionLifecycleStateEnum(val string) (ConnectionLifecycleStateEnum, bool) {
	enum, ok := mappingConnectionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
