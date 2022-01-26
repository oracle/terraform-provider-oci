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

// Storage This is the storage configuration and status of a tenancy in Logan Analytics application
type Storage struct {

	// This indicates if old data can be archived for a tenancy
	IsArchivingEnabled *bool `mandatory:"true" json:"isArchivingEnabled"`

	ArchivingConfiguration *ArchivingConfiguration `mandatory:"true" json:"archivingConfiguration"`
}

func (m Storage) String() string {
	return common.PointerString(m)
}
