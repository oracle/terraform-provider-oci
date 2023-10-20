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

// GenericConnection Represents the metadata of a Generic Connection.
type GenericConnection struct {

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

	// Host and port separated by colon.
	// Example: `"server.example.com:1234"`
	// For multiple hosts, provide a comma separated list.
	// Example: `"server1.example.com:1000,server1.example.com:2000"`
	Host *string `mandatory:"true" json:"host"`

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

	// The Generic technology type.
	TechnologyType GenericConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m GenericConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m GenericConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m GenericConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m GenericConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m GenericConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m GenericConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m GenericConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m GenericConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m GenericConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m GenericConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m GenericConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m GenericConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m GenericConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m GenericConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m GenericConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m GenericConnection) GetSubnetId() *string {
	return m.SubnetId
}

func (m GenericConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenericConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetGenericConnectionTechnologyTypeEnumStringValues(), ",")))
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
func (m GenericConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericConnection GenericConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeGenericConnection
	}{
		"GENERIC",
		(MarshalTypeGenericConnection)(m),
	}

	return json.Marshal(&s)
}

// GenericConnectionTechnologyTypeEnum Enum with underlying type: string
type GenericConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for GenericConnectionTechnologyTypeEnum
const (
	GenericConnectionTechnologyTypeGeneric GenericConnectionTechnologyTypeEnum = "GENERIC"
)

var mappingGenericConnectionTechnologyTypeEnum = map[string]GenericConnectionTechnologyTypeEnum{
	"GENERIC": GenericConnectionTechnologyTypeGeneric,
}

var mappingGenericConnectionTechnologyTypeEnumLowerCase = map[string]GenericConnectionTechnologyTypeEnum{
	"generic": GenericConnectionTechnologyTypeGeneric,
}

// GetGenericConnectionTechnologyTypeEnumValues Enumerates the set of values for GenericConnectionTechnologyTypeEnum
func GetGenericConnectionTechnologyTypeEnumValues() []GenericConnectionTechnologyTypeEnum {
	values := make([]GenericConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingGenericConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for GenericConnectionTechnologyTypeEnum
func GetGenericConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"GENERIC",
	}
}

// GetMappingGenericConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericConnectionTechnologyTypeEnum(val string) (GenericConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingGenericConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
