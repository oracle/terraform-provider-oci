// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Subscription, Commitment and and Rate Card Details
//
// Set of APIs that return the Subscription Details, Commitment and Effective Rate Card Details
//

package osubsubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscribedServiceSummary Subscribed Service summary
type SubscribedServiceSummary struct {

	// SPM internal Subscribed Service ID
	Id *string `mandatory:"true" json:"id"`

	Product *SubscriptionProduct `mandatory:"false" json:"product"`

	// Subscribed service quantity
	Quantity *string `mandatory:"false" json:"quantity"`

	// Subscribed service status
	Status *string `mandatory:"false" json:"status"`

	// Subscribed service operation type
	OperationType *string `mandatory:"false" json:"operationType"`

	// Subscribed service net unit price
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Funded Allocation line value
	// example: 12000.00
	FundedAllocationValue *string `mandatory:"false" json:"fundedAllocationValue"`

	// This field contains the name of the partner to which the subscription belongs - depending on which the invoicing may differ
	PartnerTransactionType *string `mandatory:"false" json:"partnerTransactionType"`

	// Term value in Months
	TermValue *int64 `mandatory:"false" json:"termValue"`

	// Term value UOM
	TermValueUOM *string `mandatory:"false" json:"termValueUOM"`

	// Booking Opportunity Number of Subscribed Service
	BookingOptyNumber *string `mandatory:"false" json:"bookingOptyNumber"`

	// Subscribed service total value
	TotalValue *string `mandatory:"false" json:"totalValue"`

	// Sales Order Number associated to the subscribed service
	OrderNumber *int64 `mandatory:"false" json:"orderNumber"`

	// Subscribed service data center region
	DataCenterRegion *string `mandatory:"false" json:"dataCenterRegion"`

	// Subscribed service pricing model
	PricingModel *string `mandatory:"false" json:"pricingModel"`

	// Subscribed service program type
	ProgramType *string `mandatory:"false" json:"programType"`

	// Subscribed service promotion type
	PromoType *string `mandatory:"false" json:"promoType"`

	// Subscribed service CSI number
	Csi *int64 `mandatory:"false" json:"csi"`

	// Subscribed service intent to pay flag
	IsIntentToPay *bool `mandatory:"false" json:"isIntentToPay"`

	// Subscribed service start date
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Subscribed service end date
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// List of Commitment services of a line
	CommitmentServices []Commitment `mandatory:"false" json:"commitmentServices"`
}

func (m SubscribedServiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscribedServiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
