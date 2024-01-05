// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecallArchivedDataDetails This is the input used to recall archived data
type RecallArchivedDataDetails struct {

	// This is the compartment OCID for permission checking
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This is the end of the time interval
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// This is the start of the time interval
	TimeDataStarted *common.SDKTime `mandatory:"true" json:"timeDataStarted"`

	// This is the type of the log data to be recalled
	DataType StorageDataTypeEnum `mandatory:"false" json:"dataType,omitempty"`

	// This is a list of comma-separated log sets that recalled data belongs to.
	LogSets *string `mandatory:"false" json:"logSets"`

	// This is the query that identifies the recalled data.
	Query *string `mandatory:"false" json:"query"`

	// This is the purpose of the recall
	Purpose *string `mandatory:"false" json:"purpose"`

	// This indicates if only new data has to be recalled in this recall request
	IsRecallNewDataOnly *bool `mandatory:"false" json:"isRecallNewDataOnly"`
}

func (m RecallArchivedDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecallArchivedDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStorageDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetStorageDataTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
