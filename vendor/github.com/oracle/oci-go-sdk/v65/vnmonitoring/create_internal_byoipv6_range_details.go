// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateInternalByoipv6RangeDetails Details to create internal byoip range
type CreateInternalByoipv6RangeDetails struct {

	// VCNIP-supplied unique identifier of the ByoipRange.
	Byoipv6RangeId *string `mandatory:"true" json:"byoipv6RangeId"`

	// The CIDR IPv6 address block of the Imported ByoipRange. The CIDR length is always /64.
	Byoipv6RangeCidrBlock *string `mandatory:"true" json:"byoipv6RangeCidrBlock"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Byoipv6Range.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m CreateInternalByoipv6RangeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalByoipv6RangeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
