// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
