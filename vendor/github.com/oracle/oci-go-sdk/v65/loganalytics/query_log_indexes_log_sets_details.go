// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// QueryLogIndexesLogSetsDetails Input model for QueryLogIndexesLogSets.  The only mandatory field is
// `dataType`.  `logIndex` and `logSet` are optional filters.
type QueryLogIndexesLogSetsDetails struct {

	// The data type to query (e.g. LOG, APM …)
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`

	// Optional log‑index number to filter on.  If omitted the mapping for all
	// log indexes is returned.
	LogIndex *int `mandatory:"false" json:"logIndex"`

	// Optional log‑set name to filter on.  If omitted the mapping for all
	// log sets is returned.
	LogSet *string `mandatory:"false" json:"logSet"`
}

func (m QueryLogIndexesLogSetsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryLogIndexesLogSetsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStorageDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetStorageDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
