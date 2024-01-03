// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateContainerVnicDetails Information to create a virtual network interface card (VNIC) which gives
// the containers on this container instance access to a virtual client network (VCN).
// You use this object when creating the primary VNIC during container instance launch or when creating a secondary VNIC.
// This VNIC is created in the same compartment as the specified subnet on
// behalf of the customer.
// The VNIC created by this call contains both the tags specified
// in this object as well as any tags specified in the parent container instance.
type CreateContainerVnicDetails struct {

	// The OCID of the subnet to create the VNIC in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A user-friendly name for the VNIC. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The hostname for the VNIC's primary private IP. Used for DNS.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// Whether the VNIC should be assigned a public IP address.
	IsPublicIpAssigned *bool `mandatory:"false" json:"isPublicIpAssigned"`

	// Whether the source/destination check is disabled on the VNIC.
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`

	// A list of the OCIDs of the network security groups (NSGs) to add the VNIC to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A private IP address of your choice to assign to the VNIC. Must be an
	// available IP address within the subnet's CIDR.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateContainerVnicDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerVnicDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
