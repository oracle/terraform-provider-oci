// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdditionalDetails Additional details about the provisioned services
type AdditionalDetails struct {

	// connections assigned to Golden Gate deployment
	AssignedConnections []AssignedConnectionDetails `mandatory:"false" json:"assignedConnections"`

	// OCID of model
	ModelId *string `mandatory:"false" json:"modelId"`

	// version of model
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// region of cluster
	OciRegion *string `mandatory:"false" json:"ociRegion"`

	// details of all endpoints assigned to cluster
	EndpointDetails []EndpointAdditional `mandatory:"false" json:"endpointDetails"`

	// OCID of model
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`
}

func (m AdditionalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdditionalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
