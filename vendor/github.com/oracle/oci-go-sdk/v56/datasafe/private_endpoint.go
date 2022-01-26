// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PrivateEndpoint The details required to establish a connection to the database using a private endpoint.
type PrivateEndpoint struct {

	// The OCID of the Data Safe private endpoint.
	DatasafePrivateEndpointId *string `mandatory:"false" json:"datasafePrivateEndpointId"`
}

func (m PrivateEndpoint) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PrivateEndpoint) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrivateEndpoint PrivateEndpoint
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypePrivateEndpoint
	}{
		"PRIVATE_ENDPOINT",
		(MarshalTypePrivateEndpoint)(m),
	}

	return json.Marshal(&s)
}
