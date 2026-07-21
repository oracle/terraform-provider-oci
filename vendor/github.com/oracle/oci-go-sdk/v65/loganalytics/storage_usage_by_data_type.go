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

// StorageUsageByDataType Storage usage information for one data type
type StorageUsageByDataType struct {

	// The data type (e.g. LOG, APM, …) for which the mapping is requested.
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`

	// This is the number of bytes of active data for this data type
	ActiveDataSizeInBytes *int64 `mandatory:"true" json:"activeDataSizeInBytes"`

	// This is the number of bytes of archived data for this data type.
	ArchivedDataSizeInBytes *int64 `mandatory:"true" json:"archivedDataSizeInBytes"`

	// This is the number of bytes of recalled data from archived for this data type.
	RecalledArchivedDataSizeInBytes *int64 `mandatory:"true" json:"recalledArchivedDataSizeInBytes"`
}

func (m StorageUsageByDataType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StorageUsageByDataType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStorageDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetStorageDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
