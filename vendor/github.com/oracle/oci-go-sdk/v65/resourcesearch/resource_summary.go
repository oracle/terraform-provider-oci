// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Search Service API
//
// Search for resources in your cloud network.
//

package resourcesearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceSummary A resource that exists in the cloud network that you're querying.
type ResourceSummary struct {

	// The resource type name.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The unique identifier for this particular resource, usually an OCID.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The OCID of the compartment that contains this resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time that this resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The display name (or name) of this resource, if one exists.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The availability domain where this resource exists, if applicable.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The lifecycle state of this resource, if applicable.
	LifecycleState *string `mandatory:"false" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags associated with this resource, if any. System tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	SearchContext *SearchContext `mandatory:"false" json:"searchContext"`

	// Additional identifiers to use together in a "Get" request for a specified resource, only required for resource types
	// that explicitly cannot be retrieved by using a single identifier, such as the resource's OCID.
	IdentityContext map[string]interface{} `mandatory:"false" json:"identityContext"`

	// Additional resource attribute fields of this resource that match queries with a return clause, if any.
	// For example, if you ran a query to find the private IP addresses, public IP addresses, and isPrimary field of
	// the VNIC attachment on instance resources, that field would be included in the ResourceSummary object as:
	// {"additionalDetails": {"attachedVnic": [{"publicIP" : "172.110.110.110","privateIP" : "10.10.10.10","isPrimary" : true},
	// {"publicIP" : "172.110.110.111","privateIP" : "10.10.10.11","isPrimary" : false}]}.
	// The structure of the additional details attribute fields depends on the matching resource.
	AdditionalDetails map[string]interface{} `mandatory:"false" json:"additionalDetails"`
}

func (m ResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
