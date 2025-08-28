// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VulnerabilitiesSummary Summary of vulnerabilities found in registered resources grouped by severity.
type VulnerabilitiesSummary struct {

	// Total number of vulnerabilities.
	Total *int `mandatory:"true" json:"total"`

	// Cumulative number of resources that have critical level vulnerabilities.
	Critical *int `mandatory:"true" json:"critical"`

	// Cumulative number of resources that have high level vulnerabilities.
	High *int `mandatory:"true" json:"high"`

	// Cumulative number of resources that have medium level vulnerabilities.
	Medium *int `mandatory:"true" json:"medium"`

	// Cumulative number of resources that have info level vulnerabilities.
	Info *int `mandatory:"true" json:"info"`

	// Cumulative number of resources that have low level vulnerabilities.
	Low *int `mandatory:"true" json:"low"`
}

func (m VulnerabilitiesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VulnerabilitiesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
