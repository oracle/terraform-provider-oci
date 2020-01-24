// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ValidateConnectionResult Details regarding the validation of a connection resource.
type ValidateConnectionResult struct {

	// The status returned from the connection validation.
	Status ConnectionResultEnum `mandatory:"true" json:"status"`

	// The message from the connection validation.
	Message *string `mandatory:"false" json:"message"`
}

func (m ValidateConnectionResult) String() string {
	return common.PointerString(m)
}
