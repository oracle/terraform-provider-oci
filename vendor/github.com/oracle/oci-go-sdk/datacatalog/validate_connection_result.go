// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
