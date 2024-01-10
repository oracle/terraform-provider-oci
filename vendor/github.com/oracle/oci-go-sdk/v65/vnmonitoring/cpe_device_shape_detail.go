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

// CpeDeviceShapeDetail The detailed information about a particular CPE device type. Compare with
// CpeDeviceShapeSummary.
type CpeDeviceShapeDetail struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE device shape.
	// This value uniquely identifies the type of CPE device.
	CpeDeviceShapeId *string `mandatory:"false" json:"cpeDeviceShapeId"`

	CpeDeviceInfo *CpeDeviceInfo `mandatory:"false" json:"cpeDeviceInfo"`

	// For certain CPE devices types, the customer can provide answers to
	// questions that are specific to the device type. This attribute contains
	// a list of those questions. The Networking service merges the answers with
	// other information and renders a set of CPE configuration content. To
	// provide the answers, use
	// UpdateTunnelCpeDeviceConfig.
	Parameters []CpeDeviceConfigQuestion `mandatory:"false" json:"parameters"`

	// A template of CPE device configuration information that will be merged with the customer's
	// answers to the questions to render the final CPE device configuration content. Also see:
	//   * GetCpeDeviceConfigContent
	//   * GetIpsecCpeDeviceConfigContent
	//   * GetTunnelCpeDeviceConfigContent
	Template *string `mandatory:"false" json:"template"`
}

func (m CpeDeviceShapeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CpeDeviceShapeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
