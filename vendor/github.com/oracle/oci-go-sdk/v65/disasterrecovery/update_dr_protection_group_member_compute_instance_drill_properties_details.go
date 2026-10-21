// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDrProtectionGroupMemberComputeInstanceDrillPropertiesDetails Optional update to DR drill-specific settings for a movable compute instance. If this object is omitted,
// the existing drill VNIC mappings are preserved. If provided, vnicMappings is required and must specify the
// complete VNIC mapping configuration for DR drills. Partial lists do not merge with existing mappings. An
// empty list clears the drill-specific mappings. A non-empty list is validated and replaces all existing
// drill-specific mappings. When no drill mappings are configured, DR drills use the regular recovery vnicMappings.
type UpdateDrProtectionGroupMemberComputeInstanceDrillPropertiesDetails struct {

	// A list of compute instance VNIC mappings to use for DR drills.
	VnicMappings []ComputeInstanceMovableVnicMapping `mandatory:"true" json:"vnicMappings"`
}

func (m UpdateDrProtectionGroupMemberComputeInstanceDrillPropertiesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDrProtectionGroupMemberComputeInstanceDrillPropertiesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
