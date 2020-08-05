// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResourceDiscoveryServiceSummary A service supported for use with Resource Discovery.
type ResourceDiscoveryServiceSummary struct {

	// A supported service. Example: `core`
	// For reference on service names, see the Terraform provider documentation (https://www.terraform.io/docs/providers/oci/guides/resource_discovery.html#services).
	Name *string `mandatory:"false" json:"name"`

	// The scope of the service as used with Resource Discovery.
	// This property determines the type of compartment OCID required: root compartment (`TENANCY`) or not (`COMPARTMENT`).
	// For example, `identity` is at the root compartment scope while `database` is at the compartment scope.
	DiscoveryScope ResourceDiscoveryServiceSummaryDiscoveryScopeEnum `mandatory:"false" json:"discoveryScope,omitempty"`
}

func (m ResourceDiscoveryServiceSummary) String() string {
	return common.PointerString(m)
}

// ResourceDiscoveryServiceSummaryDiscoveryScopeEnum Enum with underlying type: string
type ResourceDiscoveryServiceSummaryDiscoveryScopeEnum string

// Set of constants representing the allowable values for ResourceDiscoveryServiceSummaryDiscoveryScopeEnum
const (
	ResourceDiscoveryServiceSummaryDiscoveryScopeTenancy     ResourceDiscoveryServiceSummaryDiscoveryScopeEnum = "TENANCY"
	ResourceDiscoveryServiceSummaryDiscoveryScopeCompartment ResourceDiscoveryServiceSummaryDiscoveryScopeEnum = "COMPARTMENT"
)

var mappingResourceDiscoveryServiceSummaryDiscoveryScope = map[string]ResourceDiscoveryServiceSummaryDiscoveryScopeEnum{
	"TENANCY":     ResourceDiscoveryServiceSummaryDiscoveryScopeTenancy,
	"COMPARTMENT": ResourceDiscoveryServiceSummaryDiscoveryScopeCompartment,
}

// GetResourceDiscoveryServiceSummaryDiscoveryScopeEnumValues Enumerates the set of values for ResourceDiscoveryServiceSummaryDiscoveryScopeEnum
func GetResourceDiscoveryServiceSummaryDiscoveryScopeEnumValues() []ResourceDiscoveryServiceSummaryDiscoveryScopeEnum {
	values := make([]ResourceDiscoveryServiceSummaryDiscoveryScopeEnum, 0)
	for _, v := range mappingResourceDiscoveryServiceSummaryDiscoveryScope {
		values = append(values, v)
	}
	return values
}
