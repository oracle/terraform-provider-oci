// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ArchivingConfiguration This is the configuration for data archiving in object storage
type ArchivingConfiguration struct {

	// This is the duration data in active storage before data is archived, as described in
	// https://en.wikipedia.org/wiki/ISO_8601#Durations.
	// The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W).
	ActiveStorageDuration *string `mandatory:"false" json:"activeStorageDuration"`

	// This is the duration before archived data is deleted from object storage, as described in
	// https://en.wikipedia.org/wiki/ISO_8601#Durations
	// The largest supported unit is D, e.g. P365D (not P1Y) or P14D (not P2W).
	ArchivalStorageDuration *string `mandatory:"false" json:"archivalStorageDuration"`
}

func (m ArchivingConfiguration) String() string {
	return common.PointerString(m)
}
