// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateManagedInstanceDetails Information to update a managed instance
type UpdateManagedInstanceDetails struct {

	// OCID of the ONS topic used to send notification to users
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	// True if user allow data collection for this instance
	IsDataCollectionAuthorized *bool `mandatory:"false" json:"isDataCollectionAuthorized"`
}

func (m UpdateManagedInstanceDetails) String() string {
	return common.PointerString(m)
}
