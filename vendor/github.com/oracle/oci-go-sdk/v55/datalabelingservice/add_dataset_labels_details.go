// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataLabelingService API
//
// A description of the DataLabelingService API
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// AddDatasetLabelsDetails Adds a subset of Labels to the Dataset's LabelSet.  This LabelSet will be merged with the current Dataset's LabelSet. Requests with duplicate Labels will be rejected.
type AddDatasetLabelsDetails struct {
	LabelSet *LabelSet `mandatory:"false" json:"labelSet"`
}

func (m AddDatasetLabelsDetails) String() string {
	return common.PointerString(m)
}
