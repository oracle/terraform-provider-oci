// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PredefinedExpressionThresholdScalingConfiguration The scaling configuration for the predefined metric expression rule.
type PredefinedExpressionThresholdScalingConfiguration struct {

	// A metric value at which the scaling operation will be triggered.
	Threshold *int `mandatory:"true" json:"threshold"`

	// The period of time that the condition defined in the alarm must persist before the alarm state
	// changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the
	// alarm must persist in breaching the condition for five minutes before the alarm updates its
	// state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five
	// minutes before the alarm updates its state to "OK."
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M.
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`

	// The value is used for adjusting the count of instances by.
	InstanceCountAdjustment *int `mandatory:"false" json:"instanceCountAdjustment"`
}

// GetPendingDuration returns PendingDuration
func (m PredefinedExpressionThresholdScalingConfiguration) GetPendingDuration() *string {
	return m.PendingDuration
}

// GetInstanceCountAdjustment returns InstanceCountAdjustment
func (m PredefinedExpressionThresholdScalingConfiguration) GetInstanceCountAdjustment() *int {
	return m.InstanceCountAdjustment
}

func (m PredefinedExpressionThresholdScalingConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PredefinedExpressionThresholdScalingConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PredefinedExpressionThresholdScalingConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePredefinedExpressionThresholdScalingConfiguration PredefinedExpressionThresholdScalingConfiguration
	s := struct {
		DiscriminatorParam string `json:"scalingConfigurationType"`
		MarshalTypePredefinedExpressionThresholdScalingConfiguration
	}{
		"THRESHOLD",
		(MarshalTypePredefinedExpressionThresholdScalingConfiguration)(m),
	}

	return json.Marshal(&s)
}
