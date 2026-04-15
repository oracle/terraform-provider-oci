// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdditionalEgressRule Additional egress rule.
type AdditionalEgressRule struct {

	// Rule description.
	Description *string `mandatory:"true" json:"description"`

	// Specifies the destination CIDR block the port should be opened for. Must be IPv4 only, and cannot be part of any private range from RFC 1918 (https://datatracker.ietf.org/doc/html/rfc1918).
	DestinationCidr *string `mandatory:"true" json:"destinationCidr"`

	// The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value.
	MaxDestinationPort *int `mandatory:"true" json:"maxDestinationPort"`

	// The minimum port number, which must not be greater than the maximum port number.
	MinDestinationPort *int `mandatory:"true" json:"minDestinationPort"`
}

func (m AdditionalEgressRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdditionalEgressRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
