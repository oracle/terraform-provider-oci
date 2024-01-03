// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PerSignalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPerSignalDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPerSignalDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PerSignalDetailsStatusEnum Enum with underlying type: string
type PerSignalDetailsStatusEnum string

// Set of constants representing the allowable values for PerSignalDetailsStatusEnum
const (
	PerSignalDetailsStatusAccepted PerSignalDetailsStatusEnum = "ACCEPTED"
	PerSignalDetailsStatusDropped  PerSignalDetailsStatusEnum = "DROPPED"
	PerSignalDetailsStatusOther    PerSignalDetailsStatusEnum = "OTHER"
)

var mappingPerSignalDetailsStatusEnum = map[string]PerSignalDetailsStatusEnum{
	"ACCEPTED": PerSignalDetailsStatusAccepted,
	"DROPPED":  PerSignalDetailsStatusDropped,
	"OTHER":    PerSignalDetailsStatusOther,
}

var mappingPerSignalDetailsStatusEnumLowerCase = map[string]PerSignalDetailsStatusEnum{
	"accepted": PerSignalDetailsStatusAccepted,
	"dropped":  PerSignalDetailsStatusDropped,
	"other":    PerSignalDetailsStatusOther,
}

// GetPerSignalDetailsStatusEnumValues Enumerates the set of values for PerSignalDetailsStatusEnum
func GetPerSignalDetailsStatusEnumValues() []PerSignalDetailsStatusEnum {
	values := make([]PerSignalDetailsStatusEnum, 0)
	for _, v := range mappingPerSignalDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPerSignalDetailsStatusEnumStringValues Enumerates the set of values in String for PerSignalDetailsStatusEnum
func GetPerSignalDetailsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"DROPPED",
		"OTHER",
	}
}

// GetMappingPerSignalDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPerSignalDetailsStatusEnum(val string) (PerSignalDetailsStatusEnum, bool) {
	enum, ok := mappingPerSignalDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
