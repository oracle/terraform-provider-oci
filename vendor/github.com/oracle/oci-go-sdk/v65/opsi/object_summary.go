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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectSummary Summary resource object.
type ObjectSummary struct {

	// The name of the Awr Hub object.
	Name *string `mandatory:"false" json:"name"`

	// Size of the Awr Hub object in bytes.
	Size *int64 `mandatory:"false" json:"size"`

	// Base64-encoded MD5 hash of the Awr Hub object data.
	Md5 *string `mandatory:"false" json:"md5"`

	// The time at which the resource was first created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `mandatory:"false" json:"etag"`

	// The object's storage tier.
	StorageTier StorageTierEnum `mandatory:"false" json:"storageTier,omitempty"`

	// Archival state of an object for those in the archival tier.
	ArchivalState ArchivalStateEnum `mandatory:"false" json:"archivalState,omitempty"`

	// The date and time the Awr Hub object was modified
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`
}

func (m ObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStorageTierEnum(string(m.StorageTier)); !ok && m.StorageTier != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageTier: %s. Supported values are: %s.", m.StorageTier, strings.Join(GetStorageTierEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchivalStateEnum(string(m.ArchivalState)); !ok && m.ArchivalState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchivalState: %s. Supported values are: %s.", m.ArchivalState, strings.Join(GetArchivalStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
