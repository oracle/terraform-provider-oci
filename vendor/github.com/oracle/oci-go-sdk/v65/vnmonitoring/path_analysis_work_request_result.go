// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PathAnalysisWorkRequestResult Defines the configuration of the path analysis result.
type PathAnalysisWorkRequestResult struct {

	// List of various paths from source node to destination node
	// for a given `PathAnalysisQuery`.
	Paths []Path `mandatory:"true" json:"paths"`

	// Time the `PathAnalysisResult` was generated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m PathAnalysisWorkRequestResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PathAnalysisWorkRequestResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PathAnalysisWorkRequestResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePathAnalysisWorkRequestResult PathAnalysisWorkRequestResult
	s := struct {
		DiscriminatorParam string `json:"resultType"`
		MarshalTypePathAnalysisWorkRequestResult
	}{
		"PATH_ANALYSIS",
		(MarshalTypePathAnalysisWorkRequestResult)(m),
	}

	return json.Marshal(&s)
}
