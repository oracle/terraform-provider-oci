// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UploadData The model for uploaded binary data, like logos and images.
type UploadData struct {

	// The name used to refer to the uploaded data.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The content URL of the uploaded data.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the uploaded data.
	MimeType *string `mandatory:"false" json:"mimeType"`
}

func (m UploadData) String() string {
	return common.PointerString(m)
}
