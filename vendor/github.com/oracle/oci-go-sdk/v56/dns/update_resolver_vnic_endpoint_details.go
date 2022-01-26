// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateResolverVnicEndpointDetails The body for updating an existing resolver VNIC endpoint.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateResolverVnicEndpointDetails struct {

	// An array of network security group OCIDs for the resolver endpoint. These must be part of the VCN that the
	// resolver endpoint is a part of.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m UpdateResolverVnicEndpointDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateResolverVnicEndpointDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateResolverVnicEndpointDetails UpdateResolverVnicEndpointDetails
	s := struct {
		DiscriminatorParam string `json:"endpointType"`
		MarshalTypeUpdateResolverVnicEndpointDetails
	}{
		"VNIC",
		(MarshalTypeUpdateResolverVnicEndpointDetails)(m),
	}

	return json.Marshal(&s)
}
