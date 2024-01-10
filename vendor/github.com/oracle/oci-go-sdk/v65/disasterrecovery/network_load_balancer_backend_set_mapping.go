// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkLoadBalancerBackendSetMapping A backend set mapping between source and destination network load balancer.
type NetworkLoadBalancerBackendSetMapping struct {

	// This flag specifies if this backend set is used for traffic for non-movable compute instances.
	// Backend sets that point to non-movable instances are only enabled or disabled during DR. For non-movable instances this flag should be set to 'true'.
	// Backend sets that point to movable instances are emptied and their contents are transferred to the destination region network load balancer.  For movable instances this flag should be set to 'false'.
	// Example: `true`
	IsBackendSetForNonMovable *bool `mandatory:"true" json:"isBackendSetForNonMovable"`

	// The name of the source backend set.
	// Example: `example_backend_set`
	SourceBackendSetName *string `mandatory:"true" json:"sourceBackendSetName"`

	// The name of the destination backend set.
	// Example: `example_backend_set`
	DestinationBackendSetName *string `mandatory:"true" json:"destinationBackendSetName"`
}

func (m NetworkLoadBalancerBackendSetMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkLoadBalancerBackendSetMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
