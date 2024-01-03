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

// InternalByoipv6RangeAllocations Details containing the ipv6 prefixes allocation under an InternalByoipv6Range.
type InternalByoipv6RangeAllocations struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `InternalByoipv6Range` resource to which the ipv6 CIDR block belongs.
	Byoipv6RangeId *string `mandatory:"false" json:"byoipv6RangeId"`

	// The ipv6 CIDR block to be used in the vcn under a ByoipRange. It could be all of the CIDR block identified in `byoipv6RangeId`, or a subrange.
	// Example: `2001:0db8:0123:45::/56`
	Byoipv6AllocatedCidr *string `mandatory:"false" json:"byoipv6AllocatedCidr"`

	// The OCID of the ByoipRange's VCN.
	VcnId *string `mandatory:"false" json:"vcnId"`
}

func (m InternalByoipv6RangeAllocations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalByoipv6RangeAllocations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
