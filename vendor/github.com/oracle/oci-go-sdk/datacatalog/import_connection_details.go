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

// ImportConnectionDetails Import connection from the connection metadata and oracle wallet file.
type ImportConnectionDetails struct {

	// The information used to import the connection.
	ConnectionPayload []byte `mandatory:"true" json:"connectionPayload"`

	ConnectionDetail *CreateConnectionDetails `mandatory:"false" json:"connectionDetail"`
}

func (m ImportConnectionDetails) String() string {
	return common.PointerString(m)
}
