// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestCryptoAnalysesDetails Details of the request to start a JFR crypto event analysis.
// When the targets aren't specified, then all managed instances currently in the fleet are selected.
type RequestCryptoAnalysesDetails struct {

	// The attachment targets to start JFR.
	Targets []JfrAttachmentTarget `mandatory:"false" json:"targets"`

	// Duration of the JFR recording in minutes.
	RecordingDurationInMinutes *int `mandatory:"false" json:"recordingDurationInMinutes"`

	// Period to looking for JVMs. In addition to attach to running JVMs when given the command,
	// JVM started within the waiting period will also be attached for JFR. The value should be
	// larger than the agent polling interval setting for the fleet to ensure agent can get the
	// instructions. If not specified, the agent polling interval for the fleet is used.
	WaitingPeriodInMinutes *int `mandatory:"false" json:"waitingPeriodInMinutes"`
}

func (m RequestCryptoAnalysesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestCryptoAnalysesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
