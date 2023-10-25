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

// HdfsConnection Represents the metadata of a Hadoop Distributed File System Connection.
type HdfsConnection struct {

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

	// The Hadoop Distributed File System technology type.
	TechnologyType HdfsConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m HdfsConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m HdfsConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m HdfsConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m HdfsConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m HdfsConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m HdfsConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m HdfsConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m HdfsConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m HdfsConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m HdfsConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m HdfsConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m HdfsConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m HdfsConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m HdfsConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m HdfsConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m HdfsConnection) GetSubnetId() *string {
	return m.SubnetId
}

func (m HdfsConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HdfsConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHdfsConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetHdfsConnectionTechnologyTypeEnumStringValues(), ",")))
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
func (m HdfsConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHdfsConnection HdfsConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeHdfsConnection
	}{
		"HDFS",
		(MarshalTypeHdfsConnection)(m),
	}

	return json.Marshal(&s)
}

// HdfsConnectionTechnologyTypeEnum Enum with underlying type: string
type HdfsConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for HdfsConnectionTechnologyTypeEnum
const (
	HdfsConnectionTechnologyTypeHdfs HdfsConnectionTechnologyTypeEnum = "HDFS"
)

var mappingHdfsConnectionTechnologyTypeEnum = map[string]HdfsConnectionTechnologyTypeEnum{
	"HDFS": HdfsConnectionTechnologyTypeHdfs,
}

var mappingHdfsConnectionTechnologyTypeEnumLowerCase = map[string]HdfsConnectionTechnologyTypeEnum{
	"hdfs": HdfsConnectionTechnologyTypeHdfs,
}

// GetHdfsConnectionTechnologyTypeEnumValues Enumerates the set of values for HdfsConnectionTechnologyTypeEnum
func GetHdfsConnectionTechnologyTypeEnumValues() []HdfsConnectionTechnologyTypeEnum {
	values := make([]HdfsConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingHdfsConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHdfsConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for HdfsConnectionTechnologyTypeEnum
func GetHdfsConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"HDFS",
	}
}

// GetMappingHdfsConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHdfsConnectionTechnologyTypeEnum(val string) (HdfsConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingHdfsConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
