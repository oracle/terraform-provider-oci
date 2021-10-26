// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataLabelingService API
//
// A description of the DataLabelingService API
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// LabelSet An ordered collection of Labels that are unique by name.
type LabelSet struct {

	// An ordered collection of Labels that are unique by name.
	Items []Label `mandatory:"false" json:"items"`
}

func (m LabelSet) String() string {
	return common.PointerString(m)
}
