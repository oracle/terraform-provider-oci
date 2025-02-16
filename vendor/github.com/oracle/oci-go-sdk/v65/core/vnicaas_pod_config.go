// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VnicaasPodConfig Configuration of the VNICaaS POD. Presence of this config indicates that Service VNIC needs to be attached in a
// platform enabled VNICaaS POD. Anycast IP of VNICaaS POD will be substrateIp field of VnicAttachment request.
type VnicaasPodConfig struct {

	// VNICaas fleet name.
	FleetName *string `mandatory:"false" json:"fleetName"`

	// VNICaaS POD identifier. FleetName and podId are used as part of the topic that Service VNICs mappings are published to.
	PodId *string `mandatory:"false" json:"podId"`

	// IP address used as key for POD level allocations like slotId, to ensure that they are unique for all
	// service VNICs attached to a POD. It cannot conflict with IPs configured for other PODs.
	IpForPodLevelAllocations *string `mandatory:"false" json:"ipForPodLevelAllocations"`
}

func (m VnicaasPodConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VnicaasPodConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
