// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDisasterRecoveryDetails Details for disaster recovery of the vb instance (update).
type UpdateDisasterRecoveryDetails struct {

	// Enable Disaster Recovery on primary Visual Builder.  If Disaster Recovery is enabled alredy, then it cannot be disabled.
	IsEnableDisasterRecovery *bool `mandatory:"false" json:"isEnableDisasterRecovery"`

	// Peer remote region to create the DR standby Vb Instance.
	PeerRegion *string `mandatory:"false" json:"peerRegion"`

	// Create disaster recovery VB instance using regional IDCS, if set to true.
	IsUsingRegionalIdcs *bool `mandatory:"false" json:"isUsingRegionalIdcs"`

	// Target state of the standby Visual Builder in remote region
	DisasterRecoveryState DisasterRecoveryDetailsDisasterRecoveryStateEnum `mandatory:"false" json:"disasterRecoveryState,omitempty"`
}

func (m UpdateDisasterRecoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDisasterRecoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDisasterRecoveryDetailsDisasterRecoveryStateEnum(string(m.DisasterRecoveryState)); !ok && m.DisasterRecoveryState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryState: %s. Supported values are: %s.", m.DisasterRecoveryState, strings.Join(GetDisasterRecoveryDetailsDisasterRecoveryStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
