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

// SubscribedServiceSummary Subscribed service contract details.
type SubscribedServiceSummary struct {

	// SPM (Subscription Plan Management) internal subscribed service ID.
	Id *string `mandatory:"true" json:"id"`

	// Subscribed service line type.
	Type *string `mandatory:"false" json:"type"`

	// Subscribed service line number.
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// Subscription ID associated with the subscribed service.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	Product *RateCardProduct `mandatory:"false" json:"product"`

	// Subscribed service start date.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Subscribed service end date.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Subscribed service quantity.
	Quantity *string `mandatory:"false" json:"quantity"`

	// Subscribed service status.
	Status *string `mandatory:"false" json:"status"`

	// Subscribed service operation type.
	OperationType *string `mandatory:"false" json:"operationType"`

	// Subscribed service net unit price.
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Indicates the period for which the commitment amount can be utilized, exceeding which, the amount lapses. Also used in calculation of total contract line value.
	PricePeriod *string `mandatory:"false" json:"pricePeriod"`

	// Subscribed service line net amount.
	LineNetAmount *string `mandatory:"false" json:"lineNetAmount"`

	// Indicates if the commitment lines can have different quantities.
	IsVariableCommitment *bool `mandatory:"false" json:"isVariableCommitment"`

	// Indicates if a service can recieve usages and consequently have available amounts computed.
	IsAllowance *bool `mandatory:"false" json:"isAllowance"`

	// Subscribed service used amount.
	UsedAmount *string `mandatory:"false" json:"usedAmount"`

	// Subscribed sercice available or remaining amount.
	AvailableAmount *string `mandatory:"false" json:"availableAmount"`

	// Funded allocation line value.
	// Example: 12000.00
	FundedAllocationValue *string `mandatory:"false" json:"fundedAllocationValue"`

	// Indicates whether or not there has been usage for the subscribed service.
	IsHavingUsage *bool `mandatory:"false" json:"isHavingUsage"`

	// If true compares the rate between the rate card and the active price list, and the minimum rate would be fetched.
	IsCapToPriceList *bool `mandatory:"false" json:"isCapToPriceList"`

	// Subscribed service credit percentage.
	CreditPercentage *string `mandatory:"false" json:"creditPercentage"`

	// Contains the name of the partner to which the subscription belongs, however depending on which subscription, the invoicing can differ.
	PartnerTransactionType *string `mandatory:"false" json:"partnerTransactionType"`

	// Used in context of service credit lines.
	IsCreditEnabled *bool `mandatory:"false" json:"isCreditEnabled"`

	// Overage policy of subscribed service.
	OveragePolicy *string `mandatory:"false" json:"overagePolicy"`

	// Overage bill to of subscribed service.
	OverageBillTo *string `mandatory:"false" json:"overageBillTo"`

	// Pay As You Go policy of subscribed service. Can be null, indicating no PAYG policy.
	PaygPolicy *string `mandatory:"false" json:"paygPolicy"`

	// Not null if this service has an associated promotion line in SPM (Subscription Plan Management). Contains the line identifier from order management of
	// the associated promo line.
	PromoOrderLineId *int64 `mandatory:"false" json:"promoOrderLineId"`

	// Promotion pricing type of subscribed service. Can be null, indicating no promotional pricing.
	PromotionPricingType *string `mandatory:"false" json:"promotionPricingType"`

	// Subscribed service rate card ciscount percentage.
	RateCardDiscountPercentage *string `mandatory:"false" json:"rateCardDiscountPercentage"`

	// Subscribed service overage discount percentage.
	OverageDiscountPercentage *string `mandatory:"false" json:"overageDiscountPercentage"`

	BillToCustomer *SubscribedServiceBusinessPartner `mandatory:"false" json:"billToCustomer"`

	BillToContact *SubscribedServiceUser `mandatory:"false" json:"billToContact"`

	BillToAddress *SubscribedServiceAddress `mandatory:"false" json:"billToAddress"`

	// Payment number of subscribed service.
	PaymentNumber *string `mandatory:"false" json:"paymentNumber"`

	// Subscribed service payment expiration date.
	TimePaymentExpiry *common.SDKTime `mandatory:"false" json:"timePaymentExpiry"`

	PaymentTerm *SubscribedServicePaymentTerm `mandatory:"false" json:"paymentTerm"`

	// Payment method of subscribed service.
	PaymentMethod *string `mandatory:"false" json:"paymentMethod"`

	// Subscribed service transaction extension ID.
	TransactionExtensionId *int64 `mandatory:"false" json:"transactionExtensionId"`

	// Sales channel of subscribed service.
	SalesChannel *string `mandatory:"false" json:"salesChannel"`

	// Subscribed service eligible to renew field.
	EligibleToRenew *string `mandatory:"false" json:"eligibleToRenew"`

	// SPM (Subscription Plan Management) renewed subscription ID.
	RenewedSubscribedServiceId *string `mandatory:"false" json:"renewedSubscribedServiceId"`

	// Term value in months.
	TermValue *int64 `mandatory:"false" json:"termValue"`

	// Term value UOM (Unit of Measure).
	TermValueUom *string `mandatory:"false" json:"termValueUom"`

	// Subscribed service opportunity ID.
	RenewalOptyId *int64 `mandatory:"false" json:"renewalOptyId"`

	// Renewal opportunity number of subscribed service.
	RenewalOptyNumber *string `mandatory:"false" json:"renewalOptyNumber"`

	// Renewal opportunity type of subscribed service.
	RenewalOptyType *string `mandatory:"false" json:"renewalOptyType"`

	// Booking opportunity number of subscribed service.
	BookingOptyNumber *string `mandatory:"false" json:"bookingOptyNumber"`

	// Subscribed service revenue line ID.
	RevenueLineId *int64 `mandatory:"false" json:"revenueLineId"`

	// Revenue line number of subscribed service.
	RevenueLineNumber *string `mandatory:"false" json:"revenueLineNumber"`

	// Subscribed service major set.
	MajorSet *int64 `mandatory:"false" json:"majorSet"`

	// Subscribed service major set start date.
	TimeMajorsetStart *common.SDKTime `mandatory:"false" json:"timeMajorsetStart"`

	// Subscribed service major set end date.
	TimeMajorsetEnd *common.SDKTime `mandatory:"false" json:"timeMajorsetEnd"`

	// Subscribed service system ARR (Annualized Recurring Revenue).
	SystemArrInLc *string `mandatory:"false" json:"systemArrInLc"`

	// Subscribed service system ARR in standard currency.
	SystemArrInSc *string `mandatory:"false" json:"systemArrInSc"`

	// Subscribed service system ATR-ARR (Annualized Recurring Revenue - Available to Renew).
	SystemAtrArrInLc *string `mandatory:"false" json:"systemAtrArrInLc"`

	// Subscribed service system ATR-ARR in standard currency.
	SystemAtrArrInSc *string `mandatory:"false" json:"systemAtrArrInSc"`

	// Subscribed service revised ARR.
	RevisedArrInLc *string `mandatory:"false" json:"revisedArrInLc"`

	// Subscribed service revised ARR in standard currency.
	RevisedArrInSc *string `mandatory:"false" json:"revisedArrInSc"`

	// Subscribed service total value.
	TotalValue *string `mandatory:"false" json:"totalValue"`

	// Subscribed service promotion amount.
	OriginalPromoAmount *string `mandatory:"false" json:"originalPromoAmount"`

	// Sales order header associated with the subscribed service.
	OrderHeaderId *int64 `mandatory:"false" json:"orderHeaderId"`

	// Sales order number associated with the subscribed service.
	OrderNumber *int64 `mandatory:"false" json:"orderNumber"`

	// Order type of subscribed service.
	OrderType *string `mandatory:"false" json:"orderType"`

	// Sales order line ID associated with the subscribed service.
	OrderLineId *int64 `mandatory:"false" json:"orderLineId"`

	// Sales order line number associated with the subscribed service.
	OrderLineNumber *int `mandatory:"false" json:"orderLineNumber"`

	// Subscribed service commitment schedule ID.
	CommitmentScheduleId *string `mandatory:"false" json:"commitmentScheduleId"`

	// Subscribed service sales account party ID.
	SalesAccountPartyId *int64 `mandatory:"false" json:"salesAccountPartyId"`

	// Subscribed service data center.
	DataCenter *string `mandatory:"false" json:"dataCenter"`

	// Subscribed service data center region.
	DataCenterRegion *string `mandatory:"false" json:"dataCenterRegion"`

	// Subscribed service admin email ID.
	AdminEmail *string `mandatory:"false" json:"adminEmail"`

	// Subscribed service buyer email ID.
	BuyerEmail *string `mandatory:"false" json:"buyerEmail"`

	// Subscribed service source.
	SubscriptionSource *string `mandatory:"false" json:"subscriptionSource"`

	// Subscribed service provisioning source.
	ProvisioningSource *string `mandatory:"false" json:"provisioningSource"`

	// Subscribed service fulfillment set.
	FulfillmentSet *string `mandatory:"false" json:"fulfillmentSet"`

	// Subscribed service intent to pay parameter.
	IsIntentToPay *bool `mandatory:"false" json:"isIntentToPay"`

	// Subscribed service PAYG parameter.
	IsPayg *bool `mandatory:"false" json:"isPayg"`

	// Subscribed service pricing model.
	PricingModel *string `mandatory:"false" json:"pricingModel"`

	// Subscribed service program type.
	ProgramType *string `mandatory:"false" json:"programType"`

	// Subscribed service start date type.
	StartDateType *string `mandatory:"false" json:"startDateType"`

	// Subscribed service provisioning date.
	TimeProvisioned *common.SDKTime `mandatory:"false" json:"timeProvisioned"`

	// Subscribed service promotion type.
	PromoType *string `mandatory:"false" json:"promoType"`

	ServiceToCustomer *SubscribedServiceBusinessPartner `mandatory:"false" json:"serviceToCustomer"`

	ServiceToContact *SubscribedServiceUser `mandatory:"false" json:"serviceToContact"`

	ServiceToAddress *SubscribedServiceAddress `mandatory:"false" json:"serviceToAddress"`

	SoldToCustomer *SubscribedServiceBusinessPartner `mandatory:"false" json:"soldToCustomer"`

	SoldToContact *SubscribedServiceUser `mandatory:"false" json:"soldToContact"`

	EndUserCustomer *SubscribedServiceBusinessPartner `mandatory:"false" json:"endUserCustomer"`

	EndUserContact *SubscribedServiceUser `mandatory:"false" json:"endUserContact"`

	EndUserAddress *SubscribedServiceAddress `mandatory:"false" json:"endUserAddress"`

	ResellerCustomer *SubscribedServiceBusinessPartner `mandatory:"false" json:"resellerCustomer"`

	ResellerContact *SubscribedServiceUser `mandatory:"false" json:"resellerContact"`

	ResellerAddress *SubscribedServiceAddress `mandatory:"false" json:"resellerAddress"`

	// Subscribed service CSI number.
	Csi *int64 `mandatory:"false" json:"csi"`

	// Identifier for a customer's transactions for purchase of any Oracle services.
	CustomerTransactionReference *string `mandatory:"false" json:"customerTransactionReference"`

	// Subscribed service partner credit amount.
	PartnerCreditAmount *string `mandatory:"false" json:"partnerCreditAmount"`

	// Indicates if the subscribed service has a single rate card.
	IsSingleRateCard *bool `mandatory:"false" json:"isSingleRateCard"`

	// Subscribed service agreement ID.
	AgreementId *int64 `mandatory:"false" json:"agreementId"`

	// Subscribed service agreement name.
	AgreementName *string `mandatory:"false" json:"agreementName"`

	// Subscribed service agreement type.
	AgreementType *string `mandatory:"false" json:"agreementType"`

	// Subscribed service invoice frequency.
	BillingFrequency *string `mandatory:"false" json:"billingFrequency"`

	// Subscribed service welcome email sent date.
	TimeWelcomeEmailSent *common.SDKTime `mandatory:"false" json:"timeWelcomeEmailSent"`

	// Subscribed service configuration email sent date.
	TimeServiceConfigurationEmailSent *common.SDKTime `mandatory:"false" json:"timeServiceConfigurationEmailSent"`

	// Subscribed service customer configuration date.
	TimeCustomerConfig *common.SDKTime `mandatory:"false" json:"timeCustomerConfig"`

	// Subscribed service agreement end date.
	TimeAgreementEnd *common.SDKTime `mandatory:"false" json:"timeAgreementEnd"`

	// Subscribed service creation date.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// User that created the subscribed service.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Subscribed service last update date.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// User that updated the subscribed service.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// SPM (Subscription Plan Management) rate card type.
	RatecardType *string `mandatory:"false" json:"ratecardType"`
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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
