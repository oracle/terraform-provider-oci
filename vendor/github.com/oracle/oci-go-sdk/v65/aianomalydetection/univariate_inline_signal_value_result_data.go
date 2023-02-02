// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnivariateInlineSignalValueResultData The signal value of the detection result.
type UnivariateInlineSignalValueResultData struct {

	// The timestamp of the time series.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The index of the time series.
	RowIndex *int `mandatory:"false" json:"rowIndex"`

	// The score represents the severity of the anomaly.
	AnomalyScore *float64 `mandatory:"false" json:"anomalyScore"`

	// The estimated value if there is an anomaly.
	EstimatedValue *float64 `mandatory:"false" json:"estimatedValue"`

	// The derived value after running preprocessing.
	ImputedValue *float64 `mandatory:"false" json:"imputedValue"`

	// The value provided by the user.
	ActualValue *float64 `mandatory:"false" json:"actualValue"`

	// This flag indicates if this data point is anomalous or not.
	IsAnomaly *bool `mandatory:"false" json:"isAnomaly"`
}

func (m UnivariateInlineSignalValueResultData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnivariateInlineSignalValueResultData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
