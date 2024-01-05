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

// DeniedSecurityActionDetails Defines details for the security action taken on denied traffic.
type DeniedSecurityActionDetails struct {

	// If true, the evaluated security list and network security group ID details are incomplete.
	IsRestrictedOrPartial *bool `mandatory:"true" json:"isRestrictedOrPartial"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of evaluated security lists associcated
	// with the OCI resource's subnet.
	EvaluatedSecurityListIds []string `mandatory:"false" json:"evaluatedSecurityListIds"`

	// List of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of evaluated network security groups
	// associated with the OCI resource's VNIC.
	EvaluatedNsgIds []string `mandatory:"false" json:"evaluatedNsgIds"`
}

func (m DeniedSecurityActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeniedSecurityActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
