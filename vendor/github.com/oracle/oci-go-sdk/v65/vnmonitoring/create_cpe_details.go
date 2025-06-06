// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCpeDetails The representation of CreateCpeDetails
type CreateCpeDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the CPE.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The public IP address of the on-premises router.
	// Example: `203.0.113.2`
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE device type. You can provide
	// a value if you want to later generate CPE device configuration content for IPSec connections
	// that use this CPE. You can also call UpdateCpe later to
	// provide a value. For a list of possible values, see
	// ListCpeDeviceShapes.
	// For more information about generating CPE device configuration content, see:
	//   * GetCpeDeviceConfigContent
	//   * GetIpsecCpeDeviceConfigContent
	//   * GetTunnelCpeDeviceConfigContent
	//   * GetTunnelCpeDeviceConfig
	CpeDeviceShapeId *string `mandatory:"false" json:"cpeDeviceShapeId"`

	// Indicates whether this CPE is of type `private` or not.
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`
}

func (m CreateCpeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCpeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
