// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Publisher) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
