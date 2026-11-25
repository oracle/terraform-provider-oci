// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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
// The Core Services API provides both IPv4-only and dual-stack (IPv4 and IPv6) endpoints. By default, OCI-published SDKs and the CLI use IPv4-only endpoints.
// Use a dual-stack endpoint to connect over IPv6 while retaining IPv4 compatibility.
// Dual-stack endpoints use the following format: https://{service}.{region}.ds.oci.{secondLevelDomain}.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstancePoolLaunchBeforeTerminatingDetails Default launch-before-termination settings for instance pool refreshes.
type InstancePoolLaunchBeforeTerminatingDetails struct {

	// The maximum percentage of instances that can be in the pool during a refresh.
	// The service uses this value to calculate launch-before-terminate surge capacity.
	MaxHealthyPercent *int `mandatory:"true" json:"maxHealthyPercent"`

	// Whether to refresh eligible instances that already use the target instance configuration.
	IsFullRefresh *bool `mandatory:"true" json:"isFullRefresh"`
}

func (m InstancePoolLaunchBeforeTerminatingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePoolLaunchBeforeTerminatingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
