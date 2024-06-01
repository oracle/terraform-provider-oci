// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// Db2Connection Represents the metadata of a DB2 Connection.
type Db2Connection struct {

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

	// The username Oracle GoldenGate uses to connect to the DB2 database.
	// This username must already exist and be available by the DB2 to be connected to.
	Username *string `mandatory:"true" json:"username"`

	// The name or address of a host.
	Host *string `mandatory:"true" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"true" json:"port"`

	// The name of the database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

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

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of name-value pair attribute entries.
	// Used as additional parameters in connection string.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// The DB2 technology type.
	TechnologyType Db2ConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Protocol for the DB2 database.
	SecurityProtocol Db2ConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m Db2Connection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m Db2Connection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m Db2Connection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m Db2Connection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m Db2Connection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m Db2Connection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m Db2Connection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m Db2Connection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m Db2Connection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m Db2Connection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m Db2Connection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m Db2Connection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m Db2Connection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m Db2Connection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m Db2Connection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m Db2Connection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m Db2Connection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m Db2Connection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m Db2Connection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Db2Connection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDb2ConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetDb2ConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDb2ConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetDb2ConnectionSecurityProtocolEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Db2Connection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDb2Connection Db2Connection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeDb2Connection
	}{
		"DB2",
		(MarshalTypeDb2Connection)(m),
	}

	return json.Marshal(&s)
}

// Db2ConnectionTechnologyTypeEnum Enum with underlying type: string
type Db2ConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for Db2ConnectionTechnologyTypeEnum
const (
	Db2ConnectionTechnologyTypeDb2Zos Db2ConnectionTechnologyTypeEnum = "DB2_ZOS"
)

var mappingDb2ConnectionTechnologyTypeEnum = map[string]Db2ConnectionTechnologyTypeEnum{
	"DB2_ZOS": Db2ConnectionTechnologyTypeDb2Zos,
}

var mappingDb2ConnectionTechnologyTypeEnumLowerCase = map[string]Db2ConnectionTechnologyTypeEnum{
	"db2_zos": Db2ConnectionTechnologyTypeDb2Zos,
}

// GetDb2ConnectionTechnologyTypeEnumValues Enumerates the set of values for Db2ConnectionTechnologyTypeEnum
func GetDb2ConnectionTechnologyTypeEnumValues() []Db2ConnectionTechnologyTypeEnum {
	values := make([]Db2ConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingDb2ConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDb2ConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for Db2ConnectionTechnologyTypeEnum
func GetDb2ConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"DB2_ZOS",
	}
}

// GetMappingDb2ConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDb2ConnectionTechnologyTypeEnum(val string) (Db2ConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingDb2ConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// Db2ConnectionSecurityProtocolEnum Enum with underlying type: string
type Db2ConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for Db2ConnectionSecurityProtocolEnum
const (
	Db2ConnectionSecurityProtocolPlain Db2ConnectionSecurityProtocolEnum = "PLAIN"
	Db2ConnectionSecurityProtocolTls   Db2ConnectionSecurityProtocolEnum = "TLS"
)

var mappingDb2ConnectionSecurityProtocolEnum = map[string]Db2ConnectionSecurityProtocolEnum{
	"PLAIN": Db2ConnectionSecurityProtocolPlain,
	"TLS":   Db2ConnectionSecurityProtocolTls,
}

var mappingDb2ConnectionSecurityProtocolEnumLowerCase = map[string]Db2ConnectionSecurityProtocolEnum{
	"plain": Db2ConnectionSecurityProtocolPlain,
	"tls":   Db2ConnectionSecurityProtocolTls,
}

// GetDb2ConnectionSecurityProtocolEnumValues Enumerates the set of values for Db2ConnectionSecurityProtocolEnum
func GetDb2ConnectionSecurityProtocolEnumValues() []Db2ConnectionSecurityProtocolEnum {
	values := make([]Db2ConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingDb2ConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetDb2ConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for Db2ConnectionSecurityProtocolEnum
func GetDb2ConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
	}
}

// GetMappingDb2ConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDb2ConnectionSecurityProtocolEnum(val string) (Db2ConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingDb2ConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
