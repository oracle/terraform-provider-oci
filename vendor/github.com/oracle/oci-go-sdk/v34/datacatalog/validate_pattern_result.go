// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v34/common"
)

// ValidatePatternResult Details regarding the validation of a pattern resource.
type ValidatePatternResult struct {

	// The status returned from the pattern validation.
	Status *string `mandatory:"true" json:"status"`

	// The message from the pattern validation.
	Message *string `mandatory:"false" json:"message"`

	// The expression used in the pattern validation.
	Expression *string `mandatory:"false" json:"expression"`

	// Collection of logical entities derived from the expression applied to a list of file paths.
	DerivedLogicalEntities []DerivedLogicalEntities `mandatory:"false" json:"derivedLogicalEntities"`
}

func (m ValidatePatternResult) String() string {
	return common.PointerString(m)
}
