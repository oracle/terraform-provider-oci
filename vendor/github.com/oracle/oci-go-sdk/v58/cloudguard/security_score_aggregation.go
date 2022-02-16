// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SecurityScoreAggregation Provides the dimensions and their corresponding count value.
type SecurityScoreAggregation struct {

	// The key-value pairs of dimensions and their names.
	DimensionsMap map[string]string `mandatory:"true" json:"dimensionsMap"`

	// The security rating with given dimension/s
	SecurityRating SecurityRatingEnum `mandatory:"true" json:"securityRating"`

	// The security score with given dimension/s
	SecurityScore *int `mandatory:"true" json:"securityScore"`
}

func (m SecurityScoreAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityScoreAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityRatingEnum(string(m.SecurityRating)); !ok && m.SecurityRating != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityRating: %s. Supported values are: %s.", m.SecurityRating, strings.Join(GetSecurityRatingEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
