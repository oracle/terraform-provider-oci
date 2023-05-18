// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpointSummary To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type PrivateEndpointSummary struct {

	// The name given to the Private Endpoint. Avoid entering confidential information.
	// Example: my-new-pe1
	Name *string `mandatory:"true" json:"name"`

	// The Object Storage namespace with which the Private Endpoint is associated.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The compartment ID in which the Private Endpoint is authorized.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the user who created the Private Endpoint.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The date and time the Private Endpoint was created, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Private Endpoint was updated, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`

	// A prefix to use for the private endpoint. The customer VCN's DNS records are
	// updated with this prefix. The prefix input from the customer will be the first sub-domain in the endpointFqdn.
	// Example: If the prefix chosen is "abc", then the endpointFqdn will be 'abc.private.objectstorage.<region>.oraclecloud.com'
	Prefix *string `mandatory:"true" json:"prefix"`

	// The entity tag for the Private Endpoint.
	ETag *string `mandatory:"true" json:"eTag"`

	// The summaries of Private Endpoints' lifecycle state.
	LifecycleState PrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags *interface{} `mandatory:"false" json:"definedTags"`
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
