// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// StepMessage Contents of a step message.
type StepMessage struct {

	// The status message of the steps in a recipe during pipeline initialization.
	// https://docs.oracle.com/en/middleware/goldengate/core/23/oggra/rest-endpoints.html
	Message *string `mandatory:"true" json:"message"`

	// The code returned when GoldenGate reports an error while running a step during pipeline initialization.
	// https://docs.oracle.com/en/middleware/goldengate/core/23/error-messages/ogg-00001-ogg-40000.html#GUID-97FF7AA7-7A5C-4AA7-B29F-3CC8D26761F2
	Code *string `mandatory:"true" json:"code"`

	// Date and time of a message issued by steps in a recipe during pipeline initialization.
	// The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-07-25T21:10:29.600Z`.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// The severity returned when calling GoldenGate API messages for a step in a recipe during pipeline initialization.
	// https://docs.oracle.com/en/middleware/goldengate/core/23/oggra/rest-endpoints.html
	Severity SeverityTypeEnum `mandatory:"true" json:"severity"`
}

func (m StepMessage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StepMessage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSeverityTypeEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetSeverityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
