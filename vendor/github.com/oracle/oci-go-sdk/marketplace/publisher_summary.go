// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PublisherSummary Summary details about the publisher of the listing.
type PublisherSummary struct {

	// Unique identifier for the publisher.
	Id *string `mandatory:"false" json:"id"`

	// The name of the publisher.
	Name *string `mandatory:"false" json:"name"`

	// A description of the publisher.
	Description *string `mandatory:"false" json:"description"`
}

func (m PublisherSummary) String() string {
	return common.PointerString(m)
}
