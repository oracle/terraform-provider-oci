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

// Publisher The model for a publisher.
type Publisher struct {

	// Unique identifier for the publisher.
	Id *string `mandatory:"false" json:"id"`

	// The name of the publisher.
	Name *string `mandatory:"false" json:"name"`

	// A description of the publisher.
	Description *string `mandatory:"false" json:"description"`

	// The year the publisher's company or organization was founded.
	YearFounded *int64 `mandatory:"false" json:"yearFounded"`

	// The publisher's website.
	WebsiteUrl *string `mandatory:"false" json:"websiteUrl"`

	// The email address of the publisher.
	ContactEmail *string `mandatory:"false" json:"contactEmail"`

	// The phone number of the publisher.
	ContactPhone *string `mandatory:"false" json:"contactPhone"`

	// The address of the publisher's headquarters.
	HqAddress *string `mandatory:"false" json:"hqAddress"`

	Logo *UploadData `mandatory:"false" json:"logo"`

	// Reference links.
	Links []Link `mandatory:"false" json:"links"`
}

func (m Publisher) String() string {
	return common.PointerString(m)
}
