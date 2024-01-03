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

// SubscribedService Subscribed service contract details
type SubscribedService struct {

	// SPM internal Subscribed Service ID
	Id *string `mandatory:"false" json:"id"`

	// Subscribed Service line type
	Type *string `mandatory:"false" json:"type"`

	// Subscribed service line number
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// Subscription ID associated to the subscribed service
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	Product *RateCardProduct `mandatory:"false" json:"product"`

	// Subscribed service start date
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Subscribed service end date
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Subscribed service quantity
	Quantity *string `mandatory:"false" json:"quantity"`

	// Subscribed service status
	Status *string `mandatory:"false" json:"status"`

	// Subscribed service operation type
	OperationType *string `mandatory:"false" json:"operationType"`

	// Subscribed service net unit price
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Indicates the period for which the commitment amount can be utilised exceeding which the amount lapses. Also used in calculation of total contract line value
	PricePeriod *string `mandatory:"false" json:"pricePeriod"`

	// Subscribed service line net amount
	LineNetAmount *string `mandatory:"false" json:"lineNetAmount"`

	// Indicates if the commitment lines can have different quantities
	IsVariableCommitment *bool `mandatory:"false" json:"isVariableCommitment"`

	// Indicates if a service can recieve usages and consequently have available amounts computed
	IsAllowance *bool `mandatory:"false" json:"isAllowance"`

	// Subscribed service used amount
	UsedAmount *string `mandatory:"false" json:"usedAmount"`

	// Subscribed sercice available or remaining amount
	AvailableAmount *string `mandatory:"false" json:"availableAmount"`

	// Funded Allocation line value
	// example: 12000.00
	FundedAllocationValue *string `mandatory:"false" json:"fundedAllocationValue"`

	// Indicator on whether or not there has been usage for the subscribed service
	IsHavingUsage *bool `mandatory:"false" json:"isHavingUsage"`

	// If true compares rate between ratecard and the active pricelist and minimum rate would be fetched
	IsCapToPriceList *bool `mandatory:"false" json:"isCapToPriceList"`

	// Subscribed service credit percentage
	CreditPercentage *string `mandatory:"false" json:"creditPercentage"`

	// This field contains the name of the partner to which the subscription belongs - depending on which the invoicing may differ
	PartnerTransactionType *string `mandatory:"false" json:"partnerTransactionType"`

	// Used in context of service credit lines
	IsCreditEnabled *bool `mandatory:"false" json:"isCreditEnabled"`

	// Overage Policy of Subscribed Service
	OveragePolicy *string `mandatory:"false" json:"overagePolicy"`

	// Overage Bill To of Subscribed Service
	OverageBillTo *string `mandatory:"false" json:"overageBillTo"`

	// Pay As You Go policy of Subscribed Service (Can be null - indicating no payg policy)
	PaygPolicy *string `mandatory:"false" json:"paygPolicy"`

	// Not null if this service has an associated promotion line in SPM. Contains the line identifier from Order Management of
	// the associated promo line.
	PromoOrderLineId *int64 `mandatory:"false" json:"promoOrderLineId"`

	// Promotion Pricing Type of Subscribed Service (Can be null - indicating no promotion pricing)
	PromotionPricingType *string `mandatory:"false" json:"promotionPricingType"`

	// Subscribed service Rate Card Discount Percentage
	RateCardDiscountPercentage *string `mandatory:"false" json:"rateCardDiscountPercentage"`

	// Subscribed service Overage Discount Percentage
	OverageDiscountPercentage *string `mandatory:"false" json:"overageDiscountPercentage"`

	BillToCustomer *SubscribedServiceBusinessPartner `mandatory:"false" json:"billToCustomer"`

	BillToContact *SubscribedServiceUser `mandatory:"false" json:"billToContact"`

	BillToAddress *SubscribedServiceAddress `mandatory:"false" json:"billToAddress"`

	// Payment Number of Subscribed Service
	PaymentNumber *string `mandatory:"false" json:"paymentNumber"`

	// Subscribed service payment expiry date
	TimePaymentExpiry *common.SDKTime `mandatory:"false" json:"timePaymentExpiry"`

	PaymentTerm *SubscribedServicePaymentTerm `mandatory:"false" json:"paymentTerm"`

	// Payment Method of Subscribed Service
	PaymentMethod *string `mandatory:"false" json:"paymentMethod"`

	// Subscribed service Transaction Extension Id
	TransactionExtensionId *int64 `mandatory:"false" json:"transactionExtensionId"`

	// Sales Channel of Subscribed Service
	SalesChannel *string `mandatory:"false" json:"salesChannel"`

	// Subscribed service eligible to renew field
	EligibleToRenew *string `mandatory:"false" json:"eligibleToRenew"`

	// SPM renewed Subscription ID
	RenewedSubscribedServiceId *string `mandatory:"false" json:"renewedSubscribedServiceId"`

	// Term value in Months
	TermValue *int64 `mandatory:"false" json:"termValue"`

	// Term value UOM
	TermValueUom *string `mandatory:"false" json:"termValueUom"`

	// Subscribed service Opportunity Id
	RenewalOptyId *int64 `mandatory:"false" json:"renewalOptyId"`

	// Renewal Opportunity Number of Subscribed Service
	RenewalOptyNumber *string `mandatory:"false" json:"renewalOptyNumber"`

	// Renewal Opportunity Type of Subscribed Service
	RenewalOptyType *string `mandatory:"false" json:"renewalOptyType"`

	// Booking Opportunity Number of Subscribed Service
	BookingOptyNumber *string `mandatory:"false" json:"bookingOptyNumber"`

	// Subscribed service Revenue Line Id
	RevenueLineId *int64 `mandatory:"false" json:"revenueLineId"`

	// Revenue Line NUmber of Subscribed Service
	RevenueLineNumber *string `mandatory:"false" json:"revenueLineNumber"`

	// Subscribed service Major Set
	MajorSet *int64 `mandatory:"false" json:"majorSet"`

	// Subscribed service Major Set Start date
	TimeMajorsetStart *common.SDKTime `mandatory:"false" json:"timeMajorsetStart"`

	// Subscribed service Major Set End date
	TimeMajorsetEnd *common.SDKTime `mandatory:"false" json:"timeMajorsetEnd"`

	// Subscribed service System ARR
	SystemArrInLc *string `mandatory:"false" json:"systemArrInLc"`

	// Subscribed service System ARR in Standard Currency
	SystemArrInSc *string `mandatory:"false" json:"systemArrInSc"`

	// Subscribed service System ATR-ARR
	SystemAtrArrInLc *string `mandatory:"false" json:"systemAtrArrInLc"`

	// Subscribed service System ATR-ARR in Standard Currency
	SystemAtrArrInSc *string `mandatory:"false" json:"systemAtrArrInSc"`

	// Subscribed service Revised ARR
	RevisedArrInLc *string `mandatory:"false" json:"revisedArrInLc"`

	// Subscribed service Revised ARR in Standard Currency
	RevisedArrInSc *string `mandatory:"false" json:"revisedArrInSc"`

	// Subscribed service total value
	TotalValue *string `mandatory:"false" json:"totalValue"`

	// Subscribed service Promotion Amount
	OriginalPromoAmount *string `mandatory:"false" json:"originalPromoAmount"`

	// Sales Order Header associated to the subscribed service
	OrderHeaderId *int64 `mandatory:"false" json:"orderHeaderId"`

	// Sales Order Number associated to the subscribed service
	OrderNumber *int64 `mandatory:"false" json:"orderNumber"`

	// Order Type of Subscribed Service
	OrderType *string `mandatory:"false" json:"orderType"`

	// Sales Order Line Id associated to the subscribed service
	OrderLineId *int64 `mandatory:"false" json:"orderLineId"`

	// Sales Order Line Number associated to the subscribed service
	OrderLineNumber *int `mandatory:"false" json:"orderLineNumber"`

	// Subscribed service commitment schedule Id
	CommitmentScheduleId *string `mandatory:"false" json:"commitmentScheduleId"`

	// Subscribed service sales account party id
	SalesAccountPartyId *int64 `mandatory:"false" json:"salesAccountPartyId"`

	// Subscribed service data center
	DataCenter *string `mandatory:"false" json:"dataCenter"`

	// Subscribed service data center region
	DataCenterRegion *string `mandatory:"false" json:"dataCenterRegion"`

	// Subscribed service admin email id
	AdminEmail *string `mandatory:"false" json:"adminEmail"`

	// Subscribed service buyer email id
	BuyerEmail *string `mandatory:"false" json:"buyerEmail"`

	// Subscribed service source
	SubscriptionSource *string `mandatory:"false" json:"subscriptionSource"`

	// Subscribed service provisioning source
	ProvisioningSource *string `mandatory:"false" json:"provisioningSource"`

	// Subscribed service fulfillment set
	FulfillmentSet *string `mandatory:"false" json:"fulfillmentSet"`

	// Subscribed service intent to pay flag
	IsIntentToPay *bool `mandatory:"false" json:"isIntentToPay"`

	// Subscribed service payg flag
	IsPayg *bool `mandatory:"false" json:"isPayg"`

	// Subscribed service pricing model
	PricingModel *string `mandatory:"false" json:"pricingModel"`

	// Subscribed service program type
	ProgramType *string `mandatory:"false" json:"programType"`

	// Subscribed service start date type
	StartDateType *string `mandatory:"false" json:"startDateType"`

	// Subscribed service provisioning date
	TimeProvisioned *common.SDKTime `mandatory:"false" json:"timeProvisioned"`

	// Subscribed service promotion type
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

	// Subscribed service CSI number
	Csi *int64 `mandatory:"false" json:"csi"`

	// Identifier for a customer's transactions for purchase of ay oracle services
	CustomerTransactionReference *string `mandatory:"false" json:"customerTransactionReference"`

	// Subscribed service partner credit amount
	PartnerCreditAmount *string `mandatory:"false" json:"partnerCreditAmount"`

	// Indicates if the Subscribed service has a single ratecard
	IsSingleRateCard *bool `mandatory:"false" json:"isSingleRateCard"`

	// Subscribed service agreement ID
	AgreementId *int64 `mandatory:"false" json:"agreementId"`

	// Subscribed service agrrement name
	AgreementName *string `mandatory:"false" json:"agreementName"`

	// Subscribed service agrrement type
	AgreementType *string `mandatory:"false" json:"agreementType"`

	// Subscribed service invoice frequency
	BillingFrequency *string `mandatory:"false" json:"billingFrequency"`

	// Subscribed service welcome email sent date
	TimeWelcomeEmailSent *common.SDKTime `mandatory:"false" json:"timeWelcomeEmailSent"`

	// Subscribed service service configuration email sent date
	TimeServiceConfigurationEmailSent *common.SDKTime `mandatory:"false" json:"timeServiceConfigurationEmailSent"`

	// Subscribed service customer config date
	TimeCustomerConfig *common.SDKTime `mandatory:"false" json:"timeCustomerConfig"`

	// Subscribed service agrrement end date
	TimeAgreementEnd *common.SDKTime `mandatory:"false" json:"timeAgreementEnd"`

	// List of Commitment services of a line
	CommitmentServices []CommitmentService `mandatory:"false" json:"commitmentServices"`

	// List of Rate Cards of a Subscribed Service
	RateCards []RateCardSummary `mandatory:"false" json:"rateCards"`

	// Subscribed service creation date
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// User that created the subscribed service
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Subscribed service last update date
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// User that updated the subscribed service
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// SPM Ratecard Type
	RatecardType *string `mandatory:"false" json:"ratecardType"`
}

func (m SubscribedService) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscribedService) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
