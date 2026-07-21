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

// ArchivingConfigurationsLogIndexOverride This is the configuration for archiving the given logIndex data.
type ArchivingConfigurationsLogIndexOverride struct {

	// This is the dataType for the given logIndex level archiving configuration.
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`

	// This is the logIndex for the given archiving configuration.
	LogIndex *string `mandatory:"true" json:"logIndex"`

	// This is the duration in active storage before the given logIndex data is archived, as described in
	// https://en.wikipedia.org/wiki/ISO_8601#Durations.
	// The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W).
	ActiveStorageDuration *string `mandatory:"true" json:"activeStorageDuration"`

	// This is the duration before the given archived logIndex data is deleted from object storage, as described in
	// https://en.wikipedia.org/wiki/ISO_8601#Durations
	// The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W).
	ArchivalStorageDuration *string `mandatory:"true" json:"archivalStorageDuration"`
}

func (m ArchivingConfigurationsLogIndexOverride) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArchivingConfigurationsLogIndexOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStorageDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetStorageDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
