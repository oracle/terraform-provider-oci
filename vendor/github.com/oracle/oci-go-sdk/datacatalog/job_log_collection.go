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

// JobLogCollection Results of a job logs Listing. A job log is an audit log record inserted during the lifecycle of a job execution instance.
type JobLogCollection struct {

	// Collection of Job logs.
	Items []JobLogSummary `mandatory:"true" json:"items"`
}

func (m JobLogCollection) String() string {
	return common.PointerString(m)
}
