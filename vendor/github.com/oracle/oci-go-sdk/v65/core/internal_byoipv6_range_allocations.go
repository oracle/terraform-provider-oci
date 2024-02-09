// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalByoipv6RangeAllocations Details containing the ipv6 prefixes allocation under an InternalByoipv6Range.
type InternalByoipv6RangeAllocations struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `InternalByoipv6Range` resource to which the IPv6 prefix belongs.
	Byoipv6RangeId *string `mandatory:"false" json:"byoipv6RangeId"`

	// The IPv6 prefix to be used in the VCN under a ByoipRange. It could be all of the prefix identified in `byoipv6RangeId`, or a subrange.
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
