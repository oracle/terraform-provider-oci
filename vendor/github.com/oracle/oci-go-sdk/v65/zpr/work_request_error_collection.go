// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Zero Trust Packet Routing Control Plane API
//
// Use the Zero Trust Packet Routing Control Plane API to manage ZPR configuration and policy. See the Zero Trust Packet Routing (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/home.htm) documentation for more information.
//

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestErrorCollection A list of work request errors. Can contain both errors and other information, such as metadata.
type WorkRequestErrorCollection struct {

	// A list of work request errors.
	Items []WorkRequestError `mandatory:"true" json:"items"`
}

func (m WorkRequestErrorCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestErrorCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
