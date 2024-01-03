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

// LogListTypeEndpoint The LOG_LIST type endpoint configuration. The list of logs is first fetched using the listEndpoint configuration,
// and then the logs are subsequently fetched using the logEndpoints, which reference the list endpoint response.
// For time based incremental collection, specify the START_TIME macro with the desired time format,
// example: {START_TIME:yyMMddHHmmssZ}.
// For offset based incremental collection, specify the START_OFFSET macro with offset identifier in the API response,
// example: {START_OFFSET:$.offset}
type LogListTypeEndpoint struct {
	ListEndpoint *LogListEndpoint `mandatory:"true" json:"listEndpoint"`

	// Log endpoints, which reference the listEndpoint response, to fetch log data.
	LogEndpoints []LogEndpoint `mandatory:"true" json:"logEndpoints"`
}

func (m LogListTypeEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogListTypeEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LogListTypeEndpoint) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLogListTypeEndpoint LogListTypeEndpoint
	s := struct {
		DiscriminatorParam string `json:"endpointType"`
		MarshalTypeLogListTypeEndpoint
	}{
		"LOG_LIST",
		(MarshalTypeLogListTypeEndpoint)(m),
	}

	return json.Marshal(&s)
}
