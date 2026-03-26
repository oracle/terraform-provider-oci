// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisasterRecoveryPrecheckReport DR precheck result for standby peer in the specified placement (availabilityDomain and faultDomain).
type DisasterRecoveryPrecheckReport struct {

	// The timestamp when pre-check started. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-10-26T20:19:29.600Z`.
	TimePrecheckStarted *common.SDKTime `mandatory:"true" json:"timePrecheckStarted"`

	// The timestamp when pre-check operation finished. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-10-26T20:19:29.600Z`.
	TimePrecheckFinished *common.SDKTime `mandatory:"true" json:"timePrecheckFinished"`

	// Status of the DR precheck result.
	PrecheckStatus PrecheckStatusEnum `mandatory:"true" json:"precheckStatus"`

	// A list of precheck results.
	Checks []DisasterRecoveryPrecheckResult `mandatory:"true" json:"checks"`
}

func (m DisasterRecoveryPrecheckReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisasterRecoveryPrecheckReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrecheckStatusEnum(string(m.PrecheckStatus)); !ok && m.PrecheckStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrecheckStatus: %s. Supported values are: %s.", m.PrecheckStatus, strings.Join(GetPrecheckStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
