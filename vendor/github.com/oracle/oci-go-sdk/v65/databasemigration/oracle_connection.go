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

// OracleConnection Represents the metadata of an Oracle Database Connection.
type OracleConnection struct {

	// The connection's OCID.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCI resource ID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The username.
	Username *string `mandatory:"true" json:"username"`

	// An object's Description.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information
	// for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// OCI resource ID.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// OCI resource ID.
	KeyId *string `mandatory:"false" json:"keyId"`

	// OCI resource ID.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// List of ingress IP addresses from where to connect to this connection's privateIp.
	IngressIps []IngressIpDetails `mandatory:"false" json:"ingressIps"`

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The password.
	Password *string `mandatory:"false" json:"password"`

	// The username.
	ReplicationUsername *string `mandatory:"false" json:"replicationUsername"`

	// The password.
	ReplicationPassword *string `mandatory:"false" json:"replicationPassword"`

	// OCI resource ID.
	SecretId *string `mandatory:"false" json:"secretId"`

	// OCI resource ID.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// Connect descriptor or Easy Connect Naming method used to connect to a database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The OCID of the database being referenced.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// Name of the host the SSH key is valid for.
	SshHost *string `mandatory:"false" json:"sshHost"`

	// Private SSH key string.
	SshKey *string `mandatory:"false" json:"sshKey"`

	// The username.
	SshUser *string `mandatory:"false" json:"sshUser"`

	// Sudo location
	SshSudoLocation *string `mandatory:"false" json:"sshSudoLocation"`

	// The Oracle technology type.
	TechnologyType OracleConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m OracleConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m OracleConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m OracleConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m OracleConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m OracleConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OracleConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OracleConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m OracleConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OracleConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m OracleConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OracleConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m OracleConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m OracleConnection) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m OracleConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetIngressIps returns IngressIps
func (m OracleConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m OracleConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetUsername returns Username
func (m OracleConnection) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m OracleConnection) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m OracleConnection) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m OracleConnection) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

// GetSecretId returns SecretId
func (m OracleConnection) GetSecretId() *string {
	return m.SecretId
}

// GetPrivateEndpointId returns PrivateEndpointId
func (m OracleConnection) GetPrivateEndpointId() *string {
	return m.PrivateEndpointId
}

func (m OracleConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetOracleConnectionTechnologyTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleConnection OracleConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeOracleConnection
	}{
		"ORACLE",
		(MarshalTypeOracleConnection)(m),
	}

	return json.Marshal(&s)
}

// OracleConnectionTechnologyTypeEnum Enum with underlying type: string
type OracleConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for OracleConnectionTechnologyTypeEnum
const (
	OracleConnectionTechnologyTypeAmazonRdsOracle       OracleConnectionTechnologyTypeEnum = "AMAZON_RDS_ORACLE"
	OracleConnectionTechnologyTypeOciAutonomousDatabase OracleConnectionTechnologyTypeEnum = "OCI_AUTONOMOUS_DATABASE"
	OracleConnectionTechnologyTypeOracleDatabase        OracleConnectionTechnologyTypeEnum = "ORACLE_DATABASE"
	OracleConnectionTechnologyTypeOracleExadata         OracleConnectionTechnologyTypeEnum = "ORACLE_EXADATA"
)

var mappingOracleConnectionTechnologyTypeEnum = map[string]OracleConnectionTechnologyTypeEnum{
	"AMAZON_RDS_ORACLE":       OracleConnectionTechnologyTypeAmazonRdsOracle,
	"OCI_AUTONOMOUS_DATABASE": OracleConnectionTechnologyTypeOciAutonomousDatabase,
	"ORACLE_DATABASE":         OracleConnectionTechnologyTypeOracleDatabase,
	"ORACLE_EXADATA":          OracleConnectionTechnologyTypeOracleExadata,
}

var mappingOracleConnectionTechnologyTypeEnumLowerCase = map[string]OracleConnectionTechnologyTypeEnum{
	"amazon_rds_oracle":       OracleConnectionTechnologyTypeAmazonRdsOracle,
	"oci_autonomous_database": OracleConnectionTechnologyTypeOciAutonomousDatabase,
	"oracle_database":         OracleConnectionTechnologyTypeOracleDatabase,
	"oracle_exadata":          OracleConnectionTechnologyTypeOracleExadata,
}

// GetOracleConnectionTechnologyTypeEnumValues Enumerates the set of values for OracleConnectionTechnologyTypeEnum
func GetOracleConnectionTechnologyTypeEnumValues() []OracleConnectionTechnologyTypeEnum {
	values := make([]OracleConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingOracleConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for OracleConnectionTechnologyTypeEnum
func GetOracleConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"AMAZON_RDS_ORACLE",
		"OCI_AUTONOMOUS_DATABASE",
		"ORACLE_DATABASE",
		"ORACLE_EXADATA",
	}
}

// GetMappingOracleConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleConnectionTechnologyTypeEnum(val string) (OracleConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingOracleConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
