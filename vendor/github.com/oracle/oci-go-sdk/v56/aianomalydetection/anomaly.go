// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Anomaly An object to hold value information for each anomaly point
type Anomaly struct {

	// Name of a signal where current anomaly point belongs to
	SignalName *string `mandatory:"true" json:"signalName"`

	// The actual value for the anomaly point at given signal and timestamp/row
	ActualValue *float64 `mandatory:"true" json:"actualValue"`

	// The estimated value for the anomaly point at given signal and timestamp/row
	EstimatedValue *float64 `mandatory:"true" json:"estimatedValue"`

	// A significant score ranged from 0 to 1 to each anomaly point
	AnomalyScore *float64 `mandatory:"true" json:"anomalyScore"`

	// The value imputed by one of IDP step for missing values in origin data
	ImputedValue *float64 `mandatory:"false" json:"imputedValue"`
}

func (m Anomaly) String() string {
	return common.PointerString(m)
}
