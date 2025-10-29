// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutomaticDrConfigurationMemberAutonomousDatabase Properties for an Autonomous Database Serverless member in an Automatic DR configuration.
type AutomaticDrConfigurationMemberAutonomousDatabase struct {

	// The unique id of the member. Must not be modified by user.
	// Example: `ocid1.database.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// A flag indicating if the automatic switchover should be enabled for the Autonomous Database Serverless member in the Automatic DR configuration.
	// Example: `false`
	IsAutoSwitchoverEnabled *bool `mandatory:"false" json:"isAutoSwitchoverEnabled"`

	// A flag indicating if the automatic failover should be enabled for the Autonomous Database Serverless member in the Automatic DR configuration.
	// Example: `false`
	IsAutoFailoverEnabled *bool `mandatory:"false" json:"isAutoFailoverEnabled"`
}

// GetMemberId returns MemberId
func (m AutomaticDrConfigurationMemberAutonomousDatabase) GetMemberId() *string {
	return m.MemberId
}

func (m AutomaticDrConfigurationMemberAutonomousDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutomaticDrConfigurationMemberAutonomousDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutomaticDrConfigurationMemberAutonomousDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutomaticDrConfigurationMemberAutonomousDatabase AutomaticDrConfigurationMemberAutonomousDatabase
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeAutomaticDrConfigurationMemberAutonomousDatabase
	}{
		"AUTONOMOUS_DATABASE",
		(MarshalTypeAutomaticDrConfigurationMemberAutonomousDatabase)(m),
	}

	return json.Marshal(&s)
}
