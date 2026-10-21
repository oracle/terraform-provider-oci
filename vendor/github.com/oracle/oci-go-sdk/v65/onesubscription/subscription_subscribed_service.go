// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SubscriptionSubscribedService Subscribed service summary.
type SubscriptionSubscribedService struct {

	// SPM (Subscription Plan Management) internal subscribed service ID.
	Id *string `mandatory:"true" json:"id"`

	Product *SubscriptionProduct `mandatory:"false" json:"product"`

	// Subscribed service quantity.
	Quantity *string `mandatory:"false" json:"quantity"`

	// Subscribed service status.
	Status *string `mandatory:"false" json:"status"`

	// Subscribed service operation type.
	OperationType *string `mandatory:"false" json:"operationType"`

	// Subscribed service net unit price.
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Subscribed service used amount.
	UsedAmount *string `mandatory:"false" json:"usedAmount"`

	// Subscribed sercice available or remaining amount.
	AvailableAmount *string `mandatory:"false" json:"availableAmount"`

	// Funded allocation line value.
	// Example: 12000.00
	FundedAllocationValue *string `mandatory:"false" json:"fundedAllocationValue"`

	// The name of the partner to which the subscription belongs, however depending on which subscription, the invoicing can differ.
	PartnerTransactionType *string `mandatory:"false" json:"partnerTransactionType"`

	// Term value in months.
	TermValue *int64 `mandatory:"false" json:"termValue"`

	// Term value UOM (Unit of Measure).
	TermValueUom *string `mandatory:"false" json:"termValueUom"`

	// Booking opportunity number of subscribed service.
	BookingOptyNumber *string `mandatory:"false" json:"bookingOptyNumber"`

	// Subscribed service total value.
	TotalValue *string `mandatory:"false" json:"totalValue"`

	// Subscribed service promotion amount.
	OriginalPromoAmount *string `mandatory:"false" json:"originalPromoAmount"`

	// Sales order number associated with the subscribed service.
	OrderNumber *int64 `mandatory:"false" json:"orderNumber"`

	// Subscribed service data center region.
	DataCenterRegion *string `mandatory:"false" json:"dataCenterRegion"`

	// Subscribed service pricing model.
	PricingModel *string `mandatory:"false" json:"pricingModel"`

	// Subscribed service program type.
	ProgramType *string `mandatory:"false" json:"programType"`

	// Subscribed service promotion type.
	PromoType *string `mandatory:"false" json:"promoType"`

	// Subscribed service CSI (Customer Service Identifier) number.
	Csi *int64 `mandatory:"false" json:"csi"`

	// Subscribed service intent to pay parameter.
	IsIntentToPay *bool `mandatory:"false" json:"isIntentToPay"`

	// Subscribed service start date.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Subscribed service end date.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// List of commitment services of a line.
	CommitmentServices []CommitmentService `mandatory:"false" json:"commitmentServices"`

	// List of Associated Subscriptions
	AssociatedSubscriptions []AssociatedSubscriptions `mandatory:"false" json:"associatedSubscriptions"`
}

func (m SubscriptionSubscribedService) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionSubscribedService) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
