// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpointSummary PrivateEndpoint resource summary.
type PrivateEndpointSummary struct {

	// PrivateEndpoint identifier
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment where private endpoint exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Identifier of the customer subnet against which private endpoint is created.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Identifier of the VCN in which subnet exists.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Private endpoint display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the the private endpoint was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the private endpoint was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Lifecycle state of private endpoint.
	LifecycleState PrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// PrivateEndpoint description.
	Description *string `mandatory:"false" json:"description"`

	// The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Detailed message for the lifecycle state.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PrivateEndpointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
