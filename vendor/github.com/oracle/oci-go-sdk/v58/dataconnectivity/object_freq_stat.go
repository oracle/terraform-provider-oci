// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ObjectFreqStat One specific element. Its meaning changes in the context i.e. For ValueFrequency, the value represents a column value. For Patterns the value represents a pattern. For DataType the value represents a data type. For DataType formats (pertaining to date time) the value represent a format.
type ObjectFreqStat struct {

	// Value of the confidence of the profile result
	Value *string `mandatory:"false" json:"value"`

	// Placeholder for now, in future we will return the confidence of the profile result (because we are using sampled data and not whole data)
	Confidence *int `mandatory:"false" json:"confidence"`

	// How many times that value occurred.
	Freq *int64 `mandatory:"false" json:"freq"`

	// Frequency percentage across the sampled row counts (excluding nulls).
	FreqPercentage *float64 `mandatory:"false" json:"freqPercentage"`
}

func (m ObjectFreqStat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectFreqStat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
