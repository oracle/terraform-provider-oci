// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateLoadBalancerDetails Configuration details to update a load balancer.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateLoadBalancerDetails struct {

	// The user-friendly display name for the load balancer. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether or not the load balancer has delete protection enabled.
	// If "true", the loadbalancer will be protected against deletion if configured to accept traffic.
	// If "false", the loadbalancer will not be protected against deletion.
	// If null or unset, the value for delete protection will not be changed.
	// Example: `true`
	IsDeleteProtectionEnabled *bool `mandatory:"false" json:"isDeleteProtectionEnabled"`

	// Whether or not the load balancer has the Request Id feature enabled for HTTP listeners.
	// If "true", the load balancer will attach a unique request id header to every request
	// passed through from the load balancer to load balancer backends. This same request id
	// header also will be added to the response the lb received from the backend handling
	// the request before the load balancer returns the response to the requestor. The name
	// of the unique request id header is set the by value of requestIdHeader.
	// If "false", the loadbalancer not add this unique request id header to either the request
	// passed through to the load balancer backends nor to the reponse returned to the user.
	// New load balancers have the Request Id feature enabled unless isRequestIdEnabled is set to False.
	// Example: `true`
	IsRequestIdEnabled *bool `mandatory:"false" json:"isRequestIdEnabled"`

	// If isRequestIdEnabled is true then this field contains the name of the header field
	// that contains the unique request id that is attached to every request from
	// the load balancer to the load balancer backends and to every response from the load
	// balancer.
	// If a request to the load balancer already contains a header with same name as specified
	// in requestIdHeader then the load balancer will not change the value of that field.
	// If isRequestIdEnabled is false then this field is ignored.
	// **Notes:**
	// * Unless the header name is "" it must start with "X-" prefix.
	// * Setting the header name to "" will set it to the default: X-Request-Id.
	RequestIdHeader *string `mandatory:"false" json:"requestIdHeader"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLoadBalancerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
