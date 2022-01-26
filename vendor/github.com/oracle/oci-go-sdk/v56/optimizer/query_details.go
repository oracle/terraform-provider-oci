// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// QueryDetails The request object for querying the resource action details.
type QueryDetails struct {

	// The query describing which resources to search for.
	// For more information, see Query Language Syntax (https://docs.cloud.oracle.com/iaas/Content/CloudAdvisor/Reference/query-syntax.htm).
	Query *string `mandatory:"false" json:"query"`
}

func (m QueryDetails) String() string {
	return common.PointerString(m)
}
