// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateDatabaseToolsPrivateEndpointDetails The information about new DatabaseToolsPrivateEndpoint.
type CreateDatabaseToolsPrivateEndpointDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the containing Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsEndpointService.
	EndpointServiceId *string `mandatory:"true" json:"endpointServiceId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A description of the DatabaseToolsPrivateEndpoint.
	Description *string `mandatory:"false" json:"description"`

	// The private IP address that represents the access point for the associated endpoint service.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups
	// that the private endpoint's VNIC belongs to.  For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m CreateDatabaseToolsPrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsPrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
