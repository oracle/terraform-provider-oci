// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//
// **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
// If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequestErrorCollection Results of a workRequestError search. Contains both WorkRequestError items and other information, such as metadata.
type WorkRequestErrorCollection struct {

	// List of workRequestError objects.
	Items []WorkRequestError `mandatory:"true" json:"items"`
}

func (m WorkRequestErrorCollection) String() string {
	return common.PointerString(m)
}
