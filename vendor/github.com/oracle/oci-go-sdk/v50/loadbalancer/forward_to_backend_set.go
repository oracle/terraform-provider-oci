// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ForwardToBackendSet Action to forward requests to a given backend set.
type ForwardToBackendSet struct {

	// Name of the backend set the listener will forward the traffic to.
	// Example: `backendSetForImages`
	BackendSetName *string `mandatory:"false" json:"backendSetName"`
}

func (m ForwardToBackendSet) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ForwardToBackendSet) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeForwardToBackendSet ForwardToBackendSet
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeForwardToBackendSet
	}{
		"FORWARD_TO_BACKENDSET",
		(MarshalTypeForwardToBackendSet)(m),
	}

	return json.Marshal(&s)
}
