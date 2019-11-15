// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Screenshot The model for a listing's screenshot.
type Screenshot struct {

	// The name of the screenshot.
	Name *string `mandatory:"false" json:"name"`

	// A description of the screenshot.
	Description *string `mandatory:"false" json:"description"`

	// The content URL of the screenshot.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the screenshot.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The file extension of the screenshot.
	FileExtension *string `mandatory:"false" json:"fileExtension"`
}

func (m Screenshot) String() string {
	return common.PointerString(m)
}
