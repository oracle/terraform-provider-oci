// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PrivateEndpointDetails Private endpoint configuration details.
type PrivateEndpointDetails struct {

	// The VCN OCID for the private endpoint.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The subnet OCID for the private endpoint.
	SubnetId *string `mandatory:"true" json:"subnetId"`
}

func (m PrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PrivateEndpointDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrivateEndpointDetails PrivateEndpointDetails
	s := struct {
		DiscriminatorParam string `json:"networkEndpointType"`
		MarshalTypePrivateEndpointDetails
	}{
		"PRIVATE",
		(MarshalTypePrivateEndpointDetails)(m),
	}

	return json.Marshal(&s)
}
