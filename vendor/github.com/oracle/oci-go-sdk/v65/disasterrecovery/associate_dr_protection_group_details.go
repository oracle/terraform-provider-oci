// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AssociateDrProtectionGroupDetails The details for associating a DR protection group with a peer DR protection group.
type AssociateDrProtectionGroupDetails struct {

	// The role of the DR protection group.
	// Example: `STANDBY`
	Role DrProtectionGroupRoleEnum `mandatory:"true" json:"role"`

	// The OCID of the peer DR protection group.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	PeerId *string `mandatory:"false" json:"peerId"`

	// The region of the peer DR protection group.
	// Example: `us-ashburn-1`
	PeerRegion *string `mandatory:"false" json:"peerRegion"`
}

func (m AssociateDrProtectionGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociateDrProtectionGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrProtectionGroupRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDrProtectionGroupRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
