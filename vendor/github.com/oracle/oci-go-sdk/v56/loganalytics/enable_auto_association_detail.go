// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// EnableAutoAssociationDetail The information required to enable log source auto-association.
type EnableAutoAssociationDetail struct {

	// The unique identifier of the log group to use when auto-associting the log source to
	// eligible entities.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`
}

func (m EnableAutoAssociationDetail) String() string {
	return common.PointerString(m)
}
