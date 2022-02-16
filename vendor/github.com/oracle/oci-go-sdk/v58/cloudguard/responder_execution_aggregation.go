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

// ResponderExecutionAggregation Provides the dimensions and their corresponding count value.
type ResponderExecutionAggregation struct {

	// The key-value pairs of dimensions and their names. The key corresponds to the Analytic Dimension(s) chosen, and the value corresponds to the value of the dimension from the data. E.g. if the Analytic Dimension chosen is "RISK_LEVEL", then the value will be like "CRITICAL". If the Analytic Dimensions chosen are "RISK_LEVEL" and "RESOURCE_TYPE", then the map will have two key-value pairs of form {"RISK_LEVEL" &#58; "CRITICAL, "RESOURCE_TYPE" &#58; "LOAD_BALANCER"}
	DimensionsMap map[string]string `mandatory:"true" json:"dimensionsMap"`

	// The number of occurences with given dimension(s)
	Count *int `mandatory:"true" json:"count"`
}

func (m ResponderExecutionAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponderExecutionAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
