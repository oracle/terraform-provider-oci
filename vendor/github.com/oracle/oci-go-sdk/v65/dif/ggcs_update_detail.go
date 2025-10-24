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

// GgcsUpdateDetail Details required to update the existing GGCS instance.
type GgcsUpdateDetail struct {

	// Instance id of the existing GGCS instance to be updated.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	Ocpu *int `mandatory:"false" json:"ocpu"`

	// The OCID of a public subnet in the customer tenancy. Can be provided only for public GGCS deployments.
	PublicSubnetId *string `mandatory:"false" json:"publicSubnetId"`

	// Connection details to be associated with the Goldengate deployment.
	Connections []GgcsConnectionDetails `mandatory:"false" json:"connections"`
}

func (m GgcsUpdateDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GgcsUpdateDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
