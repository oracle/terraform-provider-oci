// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
