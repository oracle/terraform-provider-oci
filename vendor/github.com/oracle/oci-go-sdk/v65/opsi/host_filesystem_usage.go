// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostFilesystemUsage Filesystem Usage metric for the host.
type HostFilesystemUsage struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Mount points are specialized NTFS filesystem objects
	MountPoint *string `mandatory:"false" json:"mountPoint"`

	FileSystemUsageInGB *float64 `mandatory:"false" json:"fileSystemUsageInGB"`

	FileSystemAvailInPercent *float64 `mandatory:"false" json:"fileSystemAvailInPercent"`
}

// GetTimeCollected returns TimeCollected
func (m HostFilesystemUsage) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostFilesystemUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostFilesystemUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostFilesystemUsage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostFilesystemUsage HostFilesystemUsage
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostFilesystemUsage
	}{
		"HOST_FILESYSTEM_USAGE",
		(MarshalTypeHostFilesystemUsage)(m),
	}

	return json.Marshal(&s)
}
