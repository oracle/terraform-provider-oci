// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// EstimatePurgeArchivalDataSizeDetails This is the input used to estimate the size of archival data that might be purged
type EstimatePurgeArchivalDataSizeDetails struct {

	// This is the time before which data will be purged
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// This is the type of data to be purged
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`

	// This is the list of logSets associated with this purge archival data estimate
	LogSets *string `mandatory:"false" json:"logSets"`
}

func (m EstimatePurgeArchivalDataSizeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EstimatePurgeArchivalDataSizeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStorageDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetStorageDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
