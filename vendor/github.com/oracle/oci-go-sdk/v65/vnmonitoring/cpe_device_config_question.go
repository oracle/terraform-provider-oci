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

// CpeDeviceConfigQuestion An individual question that the customer can answer about the CPE device.
// The customer provides answers to these questions in
// UpdateTunnelCpeDeviceConfig.
type CpeDeviceConfigQuestion struct {

	// A string that identifies the question.
	Key *string `mandatory:"false" json:"key"`

	// A descriptive label for the question (for example, to display in a form in a graphical interface).
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description or explanation of the question, to help the customer answer accurately.
	Explanation *string `mandatory:"false" json:"explanation"`
}

func (m CpeDeviceConfigQuestion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CpeDeviceConfigQuestion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
