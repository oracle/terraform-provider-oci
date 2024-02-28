// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogRuleTaskDetails The log filter task.
// For configuration instructions, see
// Creating a Connector (https://docs.cloud.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type LogRuleTaskDetails struct {

	// A filter or mask to limit the source used in the flow defined by the connector.
	Condition *string `mandatory:"true" json:"condition"`
}

func (m LogRuleTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogRuleTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LogRuleTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLogRuleTaskDetails LogRuleTaskDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeLogRuleTaskDetails
	}{
		"logRule",
		(MarshalTypeLogRuleTaskDetails)(m),
	}

	return json.Marshal(&s)
}
