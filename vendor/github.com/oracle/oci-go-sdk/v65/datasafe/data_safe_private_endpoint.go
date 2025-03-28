// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSafePrivateEndpoint A Data Safe private endpoint that allows Data Safe to connect to databases in a customer's virtual cloud network (VCN).
type DataSafePrivateEndpoint struct {

	// The OCID of the Data Safe private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the private endpoint.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID of the underlying private endpoint.
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`

	// The private IP address of the private endpoint.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The three-label fully qualified domain name (FQDN) of the private endpoint. The customer VCN's DNS records are updated with this FQDN.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// The description of the private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the private endpoint was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the private endpoint.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DataSafePrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSafePrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
