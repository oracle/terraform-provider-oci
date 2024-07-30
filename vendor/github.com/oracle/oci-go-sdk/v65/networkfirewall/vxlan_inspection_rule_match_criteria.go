// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VxlanInspectionRuleMatchCriteria Criteria to evaluate against incoming network traffic.
// A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic.
type VxlanInspectionRuleMatchCriteria struct {

	// An array of address list names to be evaluated against the traffic source address.
	SourceAddress []string `mandatory:"false" json:"sourceAddress"`

	// An array of address list names to be evaluated against the traffic destination address.
	DestinationAddress []string `mandatory:"false" json:"destinationAddress"`
}

func (m VxlanInspectionRuleMatchCriteria) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VxlanInspectionRuleMatchCriteria) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
