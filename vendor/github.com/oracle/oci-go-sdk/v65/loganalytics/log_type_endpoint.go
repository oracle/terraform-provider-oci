// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogTypeEndpoint The LOG type endpoint configuration. Logs are fetched from the specified endpoint.
// For time based incremental collection, specify the START_TIME macro with the desired time format,
// example: {START_TIME:yyMMddHHmmssZ}.
// For offset based incremental collection, specify the START_OFFSET macro with offset identifier in the API response,
// example: {START_OFFSET:$.offset}
type LogTypeEndpoint struct {
	LogEndpoint *LogEndpoint `mandatory:"true" json:"logEndpoint"`
}

func (m LogTypeEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogTypeEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LogTypeEndpoint) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLogTypeEndpoint LogTypeEndpoint
	s := struct {
		DiscriminatorParam string `json:"endpointType"`
		MarshalTypeLogTypeEndpoint
	}{
		"LOG",
		(MarshalTypeLogTypeEndpoint)(m),
	}

	return json.Marshal(&s)
}
