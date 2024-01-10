// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpoint A private endpoint allowing Resource Manager to access nonpublic cloud resources. For more information about private endpoints, see Private Endpoint Management (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/private-endpoints.htm).
type PrivateEndpoint struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this private endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the private endpoint. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The source IP addresses that Resource Manager uses to connect to your network. Automatically assigned by Resource Manager.
	SourceIps []string `mandatory:"false" json:"sourceIps"`

	// The OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of
	// network security groups (NSGs) (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm)
	// for the private endpoint.
	// Order does not matter.
	NsgIdList []string `mandatory:"false" json:"nsgIdList"`

	// When `true`, allows the private endpoint to be used with a configuration source provider.
	IsUsedWithConfigurationSourceProvider *bool `mandatory:"false" json:"isUsedWithConfigurationSourceProvider"`

	// DNS zones to use for accessing private Git servers.
	// For private Git server instructions, see
	// Private Git Server (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/private-endpoints.htm#private-git).
	// Specify DNS fully qualified domain names (FQDNs); DNS Proxy forwards related DNS FQDN queries to the consumer DNS resolver.
	// For DNS FQDNs not specified, queries go to service provider VCN resolver.
	// Example: `abc.oraclevcn.com`
	DnsZones []string `mandatory:"false" json:"dnsZones"`

	// The date and time at which the private endpoint was created.
	// Format is defined by RFC3339.
	// Example: `2020-11-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current lifecycle state of the private endpoint.
	LifecycleState PrivateEndpointLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m PrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateEndpointLifecycleStateEnum Enum with underlying type: string
type PrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateEndpointLifecycleStateEnum
const (
	PrivateEndpointLifecycleStateActive   PrivateEndpointLifecycleStateEnum = "ACTIVE"
	PrivateEndpointLifecycleStateCreating PrivateEndpointLifecycleStateEnum = "CREATING"
	PrivateEndpointLifecycleStateDeleting PrivateEndpointLifecycleStateEnum = "DELETING"
	PrivateEndpointLifecycleStateDeleted  PrivateEndpointLifecycleStateEnum = "DELETED"
	PrivateEndpointLifecycleStateFailed   PrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingPrivateEndpointLifecycleStateEnum = map[string]PrivateEndpointLifecycleStateEnum{
	"ACTIVE":   PrivateEndpointLifecycleStateActive,
	"CREATING": PrivateEndpointLifecycleStateCreating,
	"DELETING": PrivateEndpointLifecycleStateDeleting,
	"DELETED":  PrivateEndpointLifecycleStateDeleted,
	"FAILED":   PrivateEndpointLifecycleStateFailed,
}

var mappingPrivateEndpointLifecycleStateEnumLowerCase = map[string]PrivateEndpointLifecycleStateEnum{
	"active":   PrivateEndpointLifecycleStateActive,
	"creating": PrivateEndpointLifecycleStateCreating,
	"deleting": PrivateEndpointLifecycleStateDeleting,
	"deleted":  PrivateEndpointLifecycleStateDeleted,
	"failed":   PrivateEndpointLifecycleStateFailed,
}

// GetPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumValues() []PrivateEndpointLifecycleStateEnum {
	values := make([]PrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateEndpointLifecycleStateEnum(val string) (PrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
