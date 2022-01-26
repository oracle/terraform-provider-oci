// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ValidatePatternResult Details regarding the validation of a pattern resource.
type ValidatePatternResult struct {

	// The status returned from the pattern validation.
	Status *string `mandatory:"true" json:"status"`

	// The message from the pattern validation.
	Message *string `mandatory:"false" json:"message"`

	// The expression used in the pattern validation.
	Expression *string `mandatory:"false" json:"expression"`

	// The prefix used in the pattern validation.
	FilePathPrefix *string `mandatory:"false" json:"filePathPrefix"`

	// Collection of logical entities derived from the pattern, as applied to a list of file paths.
	DerivedLogicalEntities []DerivedLogicalEntities `mandatory:"false" json:"derivedLogicalEntities"`
}

func (m ValidatePatternResult) String() string {
	return common.PointerString(m)
}
