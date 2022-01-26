// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PublisherSummary Summary details about the publisher of the resource.
type PublisherSummary struct {

	// The unique identifier for the publisher.
	Id *string `mandatory:"true" json:"id"`

	// The name of the publisher.
	DisplayName *string `mandatory:"true" json:"displayName"`
}

func (m PublisherSummary) String() string {
	return common.PointerString(m)
}
