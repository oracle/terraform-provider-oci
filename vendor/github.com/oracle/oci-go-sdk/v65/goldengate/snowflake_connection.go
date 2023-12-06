// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SnowflakeConnection Represents the metadata of a Snowflake Connection.
type SnowflakeConnection struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// JDBC connection URL.
	// e.g.: 'jdbc:snowflake://<account_name>.snowflakecomputing.com/?warehouse=<warehouse-name>&db=<db-name>'
	ConnectionUrl *string `mandatory:"true" json:"connectionUrl"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

	// List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.
	// Customers may optionally set up ingress security rules to restrict traffic from these IP addresses.
	IngressIps []IngressIpDetails `mandatory:"false" json:"ingressIps"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The username Oracle GoldenGate uses to connect to Snowflake.
	// This username must already exist and be available by Snowflake platform to be connected to.
	Username *string `mandatory:"false" json:"username"`

	// The Snowflake technology type.
	TechnologyType SnowflakeConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Used authentication mechanism to access Snowflake.
	AuthenticationType SnowflakeConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m SnowflakeConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m SnowflakeConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m SnowflakeConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m SnowflakeConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m SnowflakeConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m SnowflakeConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m SnowflakeConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m SnowflakeConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m SnowflakeConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m SnowflakeConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m SnowflakeConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m SnowflakeConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m SnowflakeConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m SnowflakeConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m SnowflakeConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m SnowflakeConnection) GetSubnetId() *string {
	return m.SubnetId
}

func (m SnowflakeConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SnowflakeConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSnowflakeConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetSnowflakeConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSnowflakeConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetSnowflakeConnectionAuthenticationTypeEnumStringValues(), ",")))
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
func (m SnowflakeConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSnowflakeConnection SnowflakeConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeSnowflakeConnection
	}{
		"SNOWFLAKE",
		(MarshalTypeSnowflakeConnection)(m),
	}

	return json.Marshal(&s)
}

// SnowflakeConnectionTechnologyTypeEnum Enum with underlying type: string
type SnowflakeConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for SnowflakeConnectionTechnologyTypeEnum
const (
	SnowflakeConnectionTechnologyTypeSnowflake SnowflakeConnectionTechnologyTypeEnum = "SNOWFLAKE"
)

var mappingSnowflakeConnectionTechnologyTypeEnum = map[string]SnowflakeConnectionTechnologyTypeEnum{
	"SNOWFLAKE": SnowflakeConnectionTechnologyTypeSnowflake,
}

var mappingSnowflakeConnectionTechnologyTypeEnumLowerCase = map[string]SnowflakeConnectionTechnologyTypeEnum{
	"snowflake": SnowflakeConnectionTechnologyTypeSnowflake,
}

// GetSnowflakeConnectionTechnologyTypeEnumValues Enumerates the set of values for SnowflakeConnectionTechnologyTypeEnum
func GetSnowflakeConnectionTechnologyTypeEnumValues() []SnowflakeConnectionTechnologyTypeEnum {
	values := make([]SnowflakeConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingSnowflakeConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSnowflakeConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for SnowflakeConnectionTechnologyTypeEnum
func GetSnowflakeConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"SNOWFLAKE",
	}
}

// GetMappingSnowflakeConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnowflakeConnectionTechnologyTypeEnum(val string) (SnowflakeConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingSnowflakeConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SnowflakeConnectionAuthenticationTypeEnum Enum with underlying type: string
type SnowflakeConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for SnowflakeConnectionAuthenticationTypeEnum
const (
	SnowflakeConnectionAuthenticationTypeBasic   SnowflakeConnectionAuthenticationTypeEnum = "BASIC"
	SnowflakeConnectionAuthenticationTypeKeyPair SnowflakeConnectionAuthenticationTypeEnum = "KEY_PAIR"
)

var mappingSnowflakeConnectionAuthenticationTypeEnum = map[string]SnowflakeConnectionAuthenticationTypeEnum{
	"BASIC":    SnowflakeConnectionAuthenticationTypeBasic,
	"KEY_PAIR": SnowflakeConnectionAuthenticationTypeKeyPair,
}

var mappingSnowflakeConnectionAuthenticationTypeEnumLowerCase = map[string]SnowflakeConnectionAuthenticationTypeEnum{
	"basic":    SnowflakeConnectionAuthenticationTypeBasic,
	"key_pair": SnowflakeConnectionAuthenticationTypeKeyPair,
}

// GetSnowflakeConnectionAuthenticationTypeEnumValues Enumerates the set of values for SnowflakeConnectionAuthenticationTypeEnum
func GetSnowflakeConnectionAuthenticationTypeEnumValues() []SnowflakeConnectionAuthenticationTypeEnum {
	values := make([]SnowflakeConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingSnowflakeConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSnowflakeConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for SnowflakeConnectionAuthenticationTypeEnum
func GetSnowflakeConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"KEY_PAIR",
	}
}

// GetMappingSnowflakeConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnowflakeConnectionAuthenticationTypeEnum(val string) (SnowflakeConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingSnowflakeConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
