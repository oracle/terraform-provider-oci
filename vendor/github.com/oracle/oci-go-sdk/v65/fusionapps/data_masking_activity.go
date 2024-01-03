// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataMaskingActivity Details of data masking activity.
type DataMaskingActivity struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Fusion Environment Identifier.
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// The current state of the DataMaskingActivity.
	LifecycleState DataMaskingActivityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the data masking activity started. An RFC3339 formatted datetime string.
	TimeMaskingStart *common.SDKTime `mandatory:"true" json:"timeMaskingStart"`

	// The time the data masking activity ended. An RFC3339 formatted datetime string.
	TimeMaskingFinish *common.SDKTime `mandatory:"true" json:"timeMaskingFinish"`
}

func (m DataMaskingActivity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataMaskingActivity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataMaskingActivityLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataMaskingActivityLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataMaskingActivityLifecycleStateEnum Enum with underlying type: string
type DataMaskingActivityLifecycleStateEnum string

// Set of constants representing the allowable values for DataMaskingActivityLifecycleStateEnum
const (
	DataMaskingActivityLifecycleStateAccepted   DataMaskingActivityLifecycleStateEnum = "ACCEPTED"
	DataMaskingActivityLifecycleStateInProgress DataMaskingActivityLifecycleStateEnum = "IN_PROGRESS"
	DataMaskingActivityLifecycleStateFailed     DataMaskingActivityLifecycleStateEnum = "FAILED"
	DataMaskingActivityLifecycleStateSucceeded  DataMaskingActivityLifecycleStateEnum = "SUCCEEDED"
	DataMaskingActivityLifecycleStateCanceled   DataMaskingActivityLifecycleStateEnum = "CANCELED"
)

var mappingDataMaskingActivityLifecycleStateEnum = map[string]DataMaskingActivityLifecycleStateEnum{
	"ACCEPTED":    DataMaskingActivityLifecycleStateAccepted,
	"IN_PROGRESS": DataMaskingActivityLifecycleStateInProgress,
	"FAILED":      DataMaskingActivityLifecycleStateFailed,
	"SUCCEEDED":   DataMaskingActivityLifecycleStateSucceeded,
	"CANCELED":    DataMaskingActivityLifecycleStateCanceled,
}

var mappingDataMaskingActivityLifecycleStateEnumLowerCase = map[string]DataMaskingActivityLifecycleStateEnum{
	"accepted":    DataMaskingActivityLifecycleStateAccepted,
	"in_progress": DataMaskingActivityLifecycleStateInProgress,
	"failed":      DataMaskingActivityLifecycleStateFailed,
	"succeeded":   DataMaskingActivityLifecycleStateSucceeded,
	"canceled":    DataMaskingActivityLifecycleStateCanceled,
}

// GetDataMaskingActivityLifecycleStateEnumValues Enumerates the set of values for DataMaskingActivityLifecycleStateEnum
func GetDataMaskingActivityLifecycleStateEnumValues() []DataMaskingActivityLifecycleStateEnum {
	values := make([]DataMaskingActivityLifecycleStateEnum, 0)
	for _, v := range mappingDataMaskingActivityLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataMaskingActivityLifecycleStateEnumStringValues Enumerates the set of values in String for DataMaskingActivityLifecycleStateEnum
func GetDataMaskingActivityLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
	}
}

// GetMappingDataMaskingActivityLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataMaskingActivityLifecycleStateEnum(val string) (DataMaskingActivityLifecycleStateEnum, bool) {
	enum, ok := mappingDataMaskingActivityLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
