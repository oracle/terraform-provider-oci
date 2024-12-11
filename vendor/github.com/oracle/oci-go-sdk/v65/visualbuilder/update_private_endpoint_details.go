// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdatePrivateEndpointDetails Private endpoint update configuration details.
type UpdatePrivateEndpointDetails struct {

	// The subnet OCID for the private endpoint. If provided then a new Private Endpoint will be created and a new Private Endpoint IP may be generated.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Network Security Group OCIDs for the Private Endpoint.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`
}

func (m UpdatePrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePrivateEndpointDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePrivateEndpointDetails UpdatePrivateEndpointDetails
	s := struct {
		DiscriminatorParam string `json:"networkEndpointType"`
		MarshalTypeUpdatePrivateEndpointDetails
	}{
		"PRIVATE",
		(MarshalTypeUpdatePrivateEndpointDetails)(m),
	}

	return json.Marshal(&s)
}
