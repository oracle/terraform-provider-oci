// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StorageUsageTrendAggregation Usage data per filesystem.
type StorageUsageTrendAggregation struct {

	// Name of filesystem.
	FileSystemName *string `mandatory:"true" json:"fileSystemName"`

	// Mount points are specialized NTFS filesystem objects.
	MountPoint *string `mandatory:"true" json:"mountPoint"`

	// Size of filesystem.
	FileSystemSizeInGBs *float64 `mandatory:"true" json:"fileSystemSizeInGBs"`

	// List of usage data samples for a filesystem.
	UsageData []StorageUsageTrend `mandatory:"true" json:"usageData"`
}

func (m StorageUsageTrendAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StorageUsageTrendAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
