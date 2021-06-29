// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v43/common"
)

// ImportDataAssetDetails Specifies the file contents to be imported.
type ImportDataAssetDetails struct {

	// The file contents to be imported. File size not to exceed 10 MB.
	ImportFileContents []byte `mandatory:"true" json:"importFileContents"`
}

func (m ImportDataAssetDetails) String() string {
	return common.PointerString(m)
}
