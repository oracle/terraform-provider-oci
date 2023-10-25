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

// GenericConnectionSummary Summary of the Generic Connection.
type GenericConnectionSummary struct {

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

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Generic technology type.
	TechnologyType GenericConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`
}

// GetId returns Id
func (m GenericConnectionSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m GenericConnectionSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m GenericConnectionSummary) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m GenericConnectionSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m GenericConnectionSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m GenericConnectionSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m GenericConnectionSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m GenericConnectionSummary) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m GenericConnectionSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m GenericConnectionSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m GenericConnectionSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m GenericConnectionSummary) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m GenericConnectionSummary) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m GenericConnectionSummary) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m GenericConnectionSummary) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m GenericConnectionSummary) GetSubnetId() *string {
	return m.SubnetId
}

func (m GenericConnectionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericConnectionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGenericConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetGenericConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GenericConnectionSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericConnectionSummary GenericConnectionSummary
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeGenericConnectionSummary
	}{
		"GENERIC",
		(MarshalTypeGenericConnectionSummary)(m),
	}

	return json.Marshal(&s)
}
