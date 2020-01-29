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

// ValidateConnectionDetails Validate connection from the connection metadata or oracle wallet file.
type ValidateConnectionDetails struct {
	ConnectionDetail *CreateConnectionDetails `mandatory:"false" json:"connectionDetail"`

	// The information used to validate the connection.
	ConnectionPayload []byte `mandatory:"false" json:"connectionPayload"`
}

func (m ValidateConnectionDetails) String() string {
	return common.PointerString(m)
}
