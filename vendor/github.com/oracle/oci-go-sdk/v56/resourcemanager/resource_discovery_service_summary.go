// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
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
