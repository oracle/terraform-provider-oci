// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseToolsPrivateEndpoint Description of DatabaseToolsPrivateEndpoint.
type DatabaseToolsPrivateEndpoint struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the containing Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsPrivateEndpoint.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsEndpointService.
	EndpointServiceId *string `mandatory:"true" json:"endpointServiceId"`

	// The time the DatabaseToolsPrivateEndpoint was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DatabaseToolsPrivateEndpoint was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The current state of the DatabaseToolsPrivateEndpoint.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of the DatabaseToolsPrivateEndpoint.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VCN that the private endpoint belongs to.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint's VNIC.
	PrivateEndpointVnicId *string `mandatory:"false" json:"privateEndpointVnicId"`

	// The private IP address that represents the access point for the associated endpoint service.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// Then FQDN to use for the private endpoint.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// A list of additional FQDNs that can be also be used for the private endpoint.
	AdditionalFqdns []string `mandatory:"false" json:"additionalFqdns"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups
	// that the private endpoint's VNIC belongs to.  For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	ReverseConnectionConfiguration *DatabaseToolsPrivateEndpointReverseConnectionConfiguration `mandatory:"false" json:"reverseConnectionConfiguration"`
}

func (m DatabaseToolsPrivateEndpoint) String() string {
	return common.PointerString(m)
}
