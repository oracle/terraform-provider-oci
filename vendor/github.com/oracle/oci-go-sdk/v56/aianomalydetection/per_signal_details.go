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

// PerSignalDetails Detailed information like statistics, metrics and status for a signal
type PerSignalDetails struct {

	// The name of a signal.
	SignalName *string `mandatory:"true" json:"signalName"`

	// Min value within a signal.
	Min *float64 `mandatory:"true" json:"min"`

	// Max value within a signal.
	Max *float64 `mandatory:"true" json:"max"`

	// Standard deviation of values within a signal.
	Std *float64 `mandatory:"true" json:"std"`

	// Status of the signal:
	//  * ACCEPTED - the signal is used for training the model
	//  * DROPPED - the signal does not meet requirement, and is dropped before training the model.
	//  * OTHER - placeholder for other status
	Status PerSignalDetailsStatusEnum `mandatory:"true" json:"status"`

	// The ratio of missing values in a signal filled/imputed by the IDP algorithm.
	MviRatio *float64 `mandatory:"false" json:"mviRatio"`

	// A boolean value to indicate if a signal is quantized or not.
	IsQuantized *bool `mandatory:"false" json:"isQuantized"`

	// Accuracy metric for a signal.
	Fap *float32 `mandatory:"false" json:"fap"`

	// detailed information for a signal.
	Details *string `mandatory:"false" json:"details"`
}

func (m PerSignalDetails) String() string {
	return common.PointerString(m)
}

// PerSignalDetailsStatusEnum Enum with underlying type: string
type PerSignalDetailsStatusEnum string

// Set of constants representing the allowable values for PerSignalDetailsStatusEnum
const (
	PerSignalDetailsStatusAccepted PerSignalDetailsStatusEnum = "ACCEPTED"
	PerSignalDetailsStatusDropped  PerSignalDetailsStatusEnum = "DROPPED"
	PerSignalDetailsStatusOther    PerSignalDetailsStatusEnum = "OTHER"
)

var mappingPerSignalDetailsStatus = map[string]PerSignalDetailsStatusEnum{
	"ACCEPTED": PerSignalDetailsStatusAccepted,
	"DROPPED":  PerSignalDetailsStatusDropped,
	"OTHER":    PerSignalDetailsStatusOther,
}

// GetPerSignalDetailsStatusEnumValues Enumerates the set of values for PerSignalDetailsStatusEnum
func GetPerSignalDetailsStatusEnumValues() []PerSignalDetailsStatusEnum {
	values := make([]PerSignalDetailsStatusEnum, 0)
	for _, v := range mappingPerSignalDetailsStatus {
		values = append(values, v)
	}
	return values
}
