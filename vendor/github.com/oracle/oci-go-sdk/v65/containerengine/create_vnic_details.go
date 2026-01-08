// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVnicDetails The properties of the secondary vnics
type CreateVnicDetails struct {

	// the ocid of the subnet to create the vnic in
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Display name for secondary vnic
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet
	AssignIpv6Ip *bool `mandatory:"false" json:"assignIpv6Ip"`

	// Whether the VNIC should be assigned a public IP address
	AssignPublicIp *bool `mandatory:"false" json:"assignPublicIp"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The number of ip addresses to attach to secondary vnic
	IpCount *int `mandatory:"false" json:"ipCount"`

	// The application resource that corresponds to this secondary vnic. Used to map pods to this specific vnic for scheduling
	ApplicationResources []string `mandatory:"false" json:"applicationResources"`

	// A list of IPv6 prefixes from which the VNIC should be assigned an IPv6 address. You can provide only the prefix
	// and OCI selects an available address from the range. You can optionally choose to leave the prefix range empty
	// and instead provide the specific IPv6 address that should be used from within that range.
	Ipv6AddressIpv6SubnetCidrPairDetails []Ipv6AddressIpv6SubnetCidrPairDetails `mandatory:"false" json:"ipv6AddressIpv6SubnetCidrPairDetails"`

	// A list of the OCIDs of the network security groups (NSGs) to add the VNIC to
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Whether the source/destination check is disabled on the VNIC
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`
}

func (m CreateVnicDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVnicDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
