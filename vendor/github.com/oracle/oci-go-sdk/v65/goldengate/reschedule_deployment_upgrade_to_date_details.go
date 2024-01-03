// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RescheduleDeploymentUpgradeToDateDetails Definition of the additional attributes for default deployment upgrade cancel.
type RescheduleDeploymentUpgradeToDateDetails struct {

	// The time of upgrade schedule. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeSchedule *common.SDKTime `mandatory:"true" json:"timeSchedule"`
}

func (m RescheduleDeploymentUpgradeToDateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RescheduleDeploymentUpgradeToDateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RescheduleDeploymentUpgradeToDateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRescheduleDeploymentUpgradeToDateDetails RescheduleDeploymentUpgradeToDateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRescheduleDeploymentUpgradeToDateDetails
	}{
		"RESCHEDULE_TO_DATE",
		(MarshalTypeRescheduleDeploymentUpgradeToDateDetails)(m),
	}

	return json.Marshal(&s)
}
