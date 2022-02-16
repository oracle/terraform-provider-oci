// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Screenshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
