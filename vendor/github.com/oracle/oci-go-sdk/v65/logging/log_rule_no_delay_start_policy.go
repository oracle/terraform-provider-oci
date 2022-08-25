// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, agent configurations, log data models,
// continuous queries, and managed continuous queries.
// For more information, see https://docs.oracle.com/en-us/iaas/Content/Logging/Concepts/loggingoverview.htm.
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogRuleNoDelayStartPolicy Log rule start policy that starts the log rule immediately.
type LogRuleNoDelayStartPolicy struct {
}

func (m LogRuleNoDelayStartPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogRuleNoDelayStartPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LogRuleNoDelayStartPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLogRuleNoDelayStartPolicy LogRuleNoDelayStartPolicy
	s := struct {
		DiscriminatorParam string `json:"startPolicyType"`
		MarshalTypeLogRuleNoDelayStartPolicy
	}{
		"NO_DELAY_START_POLICY",
		(MarshalTypeLogRuleNoDelayStartPolicy)(m),
	}

	return json.Marshal(&s)
}
