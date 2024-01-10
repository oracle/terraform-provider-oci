// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InvoiceLineSummary Invoice Line
type InvoiceLineSummary struct {

	// SPM Invoice Line internal identifier
	Id *string `mandatory:"true" json:"id"`

	Product *InvoicingProduct `mandatory:"true" json:"product"`

	// Data Center Attribute.
	DataCenter *string `mandatory:"true" json:"dataCenter"`

	// Usage start time
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// Usage end time
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`

	// AR Invoice Number for Invoice Line
	ArInvoiceNumber *string `mandatory:"false" json:"arInvoiceNumber"`
}

func (m InvoiceLineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvoiceLineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
