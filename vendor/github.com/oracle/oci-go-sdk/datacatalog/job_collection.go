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

// JobCollection Results of a jobs listing. Jobs are scheduled instances of a job definition.
type JobCollection struct {

	// Collection of jobs.
	Items []JobSummary `mandatory:"true" json:"items"`
}

func (m JobCollection) String() string {
	return common.PointerString(m)
}
