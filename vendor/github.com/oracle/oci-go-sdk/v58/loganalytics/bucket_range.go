// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BucketRange Represents querylanguage bucket command input arguments in parse output.
type BucketRange struct {

	// Lower bound of the bucket range specified in the querystring for the numeric field referenced in tbe bucket command.
	Lower *float32 `mandatory:"false" json:"lower"`

	// Upper bound of the bucket range specified in the querystring for the numeric field referenced in tbe bucket command.
	Upper *float32 `mandatory:"false" json:"upper"`

	// Optional alias of the bucket range if specified in the querystring.
	Alias *string `mandatory:"false" json:"alias"`
}

func (m BucketRange) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BucketRange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
