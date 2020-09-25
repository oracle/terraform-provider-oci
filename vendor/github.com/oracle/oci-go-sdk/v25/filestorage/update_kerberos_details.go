// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// UpdateKerberosDetails Kerberos details needed to update configuration.
type UpdateKerberosDetails struct {

	// The realm of the kerberos server Mount Target interacts with.
	KerberosRealm *string `mandatory:"false" json:"kerberosRealm"`

	// Describes how long to keep keytab entries(in seconds) after they have been rotated.
	KerberosKeyLifeSeconds *int `mandatory:"false" json:"kerberosKeyLifeSeconds"`
}

func (m UpdateKerberosDetails) String() string {
	return common.PointerString(m)
}
