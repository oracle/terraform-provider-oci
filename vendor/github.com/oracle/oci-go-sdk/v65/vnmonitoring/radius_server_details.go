// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// RadiusServerDetails RADIUS server is used for ClientVPN authentication only.
type RadiusServerDetails struct {

	// The IP address of the RADIUS server.
	ServerIp *string `mandatory:"false" json:"serverIp"`

	// The password is used for RADIUS authentication.
	SharedSecret *string `mandatory:"false" json:"sharedSecret"`

	// The port for the authentication of RADIUS service. Default is 1812.
	AuthenticationPort *int `mandatory:"false" json:"authenticationPort"`

	// The accounting port is used to access RADIUS service. Moreover, the Accounting Port is only required when RADIUS Accounting is enabled.
	AccountingPort *int `mandatory:"false" json:"accountingPort"`
}

func (m RadiusServerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RadiusServerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
