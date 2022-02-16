// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Subscription Subscription details object which extends the SubscriptionSummary
type Subscription struct {

	// Subscription plan number.
	SubscriptionPlanNumber *string `mandatory:"true" json:"subscriptionPlanNumber"`

	// Subscription id identifier (OCID).
	Id *string `mandatory:"false" json:"id"`

	// Subscription plan type.
	PlanType SubscriptionPlanTypeEnum `mandatory:"false" json:"planType,omitempty"`

	// Start date of the subscription.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Ship to customer account site address id.
	ShipToCustAcctSiteId *string `mandatory:"false" json:"shipToCustAcctSiteId"`

	// Ship to customer account role.
	ShipToCustAcctRoleId *string `mandatory:"false" json:"shipToCustAcctRoleId"`

	// Bill to customer Account id.
	BillToCustAccountId *string `mandatory:"false" json:"billToCustAccountId"`

	// Payment intension.
	IsIntentToPay *bool `mandatory:"false" json:"isIntentToPay"`

	// Currency code
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// GSI Subscription external code.
	GsiOrgCode *string `mandatory:"false" json:"gsiOrgCode"`

	// Language short code (en, de, hu, etc)
	LanguageCode *string `mandatory:"false" json:"languageCode"`

	// GSI organization external identifier.
	OrganizationId *string `mandatory:"false" json:"organizationId"`

	// Status of the upgrade.
	UpgradeState SubscriptionUpgradeStateEnum `mandatory:"false" json:"upgradeState,omitempty"`

	// This field is used to describe the Upgrade State in case of error (E.g. Upgrade failure caused by interfacing Tax details- TaxError)
	UpgradeStateDetails SubscriptionUpgradeStateDetailsEnum `mandatory:"false" json:"upgradeStateDetails,omitempty"`

	TaxInfo *TaxInfo `mandatory:"false" json:"taxInfo"`

	// Payment option list of a subscription.
	PaymentOptions []PaymentOption `mandatory:"false" json:"paymentOptions"`

	PaymentGateway *PaymentGateway `mandatory:"false" json:"paymentGateway"`

	BillingAddress *BillingAddress `mandatory:"false" json:"billingAddress"`

	// Date of upgrade/conversion when planType changed from FREE_TIER to PAYG
	TimePlanUpgrade *common.SDKTime `mandatory:"false" json:"timePlanUpgrade"`
}

func (m Subscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Subscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubscriptionPlanTypeEnum(string(m.PlanType)); !ok && m.PlanType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanType: %s. Supported values are: %s.", m.PlanType, strings.Join(GetSubscriptionPlanTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubscriptionUpgradeStateEnum(string(m.UpgradeState)); !ok && m.UpgradeState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeState: %s. Supported values are: %s.", m.UpgradeState, strings.Join(GetSubscriptionUpgradeStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubscriptionUpgradeStateDetailsEnum(string(m.UpgradeStateDetails)); !ok && m.UpgradeStateDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeStateDetails: %s. Supported values are: %s.", m.UpgradeStateDetails, strings.Join(GetSubscriptionUpgradeStateDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Subscription) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                     *string                             `json:"id"`
		PlanType               SubscriptionPlanTypeEnum            `json:"planType"`
		TimeStart              *common.SDKTime                     `json:"timeStart"`
		ShipToCustAcctSiteId   *string                             `json:"shipToCustAcctSiteId"`
		ShipToCustAcctRoleId   *string                             `json:"shipToCustAcctRoleId"`
		BillToCustAccountId    *string                             `json:"billToCustAccountId"`
		IsIntentToPay          *bool                               `json:"isIntentToPay"`
		CurrencyCode           *string                             `json:"currencyCode"`
		GsiOrgCode             *string                             `json:"gsiOrgCode"`
		LanguageCode           *string                             `json:"languageCode"`
		OrganizationId         *string                             `json:"organizationId"`
		UpgradeState           SubscriptionUpgradeStateEnum        `json:"upgradeState"`
		UpgradeStateDetails    SubscriptionUpgradeStateDetailsEnum `json:"upgradeStateDetails"`
		TaxInfo                *TaxInfo                            `json:"taxInfo"`
		PaymentOptions         []paymentoption                     `json:"paymentOptions"`
		PaymentGateway         *PaymentGateway                     `json:"paymentGateway"`
		BillingAddress         *BillingAddress                     `json:"billingAddress"`
		TimePlanUpgrade        *common.SDKTime                     `json:"timePlanUpgrade"`
		SubscriptionPlanNumber *string                             `json:"subscriptionPlanNumber"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.PlanType = model.PlanType

	m.TimeStart = model.TimeStart

	m.ShipToCustAcctSiteId = model.ShipToCustAcctSiteId

	m.ShipToCustAcctRoleId = model.ShipToCustAcctRoleId

	m.BillToCustAccountId = model.BillToCustAccountId

	m.IsIntentToPay = model.IsIntentToPay

	m.CurrencyCode = model.CurrencyCode

	m.GsiOrgCode = model.GsiOrgCode

	m.LanguageCode = model.LanguageCode

	m.OrganizationId = model.OrganizationId

	m.UpgradeState = model.UpgradeState

	m.UpgradeStateDetails = model.UpgradeStateDetails

	m.TaxInfo = model.TaxInfo

	m.PaymentOptions = make([]PaymentOption, len(model.PaymentOptions))
	for i, n := range model.PaymentOptions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.PaymentOptions[i] = nn.(PaymentOption)
		} else {
			m.PaymentOptions[i] = nil
		}
	}

	m.PaymentGateway = model.PaymentGateway

	m.BillingAddress = model.BillingAddress

	m.TimePlanUpgrade = model.TimePlanUpgrade

	m.SubscriptionPlanNumber = model.SubscriptionPlanNumber

	return
}

// SubscriptionPlanTypeEnum Enum with underlying type: string
type SubscriptionPlanTypeEnum string

// Set of constants representing the allowable values for SubscriptionPlanTypeEnum
const (
	SubscriptionPlanTypeFreeTier SubscriptionPlanTypeEnum = "FREE_TIER"
	SubscriptionPlanTypePayg     SubscriptionPlanTypeEnum = "PAYG"
)

var mappingSubscriptionPlanTypeEnum = map[string]SubscriptionPlanTypeEnum{
	"FREE_TIER": SubscriptionPlanTypeFreeTier,
	"PAYG":      SubscriptionPlanTypePayg,
}

// GetSubscriptionPlanTypeEnumValues Enumerates the set of values for SubscriptionPlanTypeEnum
func GetSubscriptionPlanTypeEnumValues() []SubscriptionPlanTypeEnum {
	values := make([]SubscriptionPlanTypeEnum, 0)
	for _, v := range mappingSubscriptionPlanTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionPlanTypeEnumStringValues Enumerates the set of values in String for SubscriptionPlanTypeEnum
func GetSubscriptionPlanTypeEnumStringValues() []string {
	return []string{
		"FREE_TIER",
		"PAYG",
	}
}

// GetMappingSubscriptionPlanTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionPlanTypeEnum(val string) (SubscriptionPlanTypeEnum, bool) {
	mappingSubscriptionPlanTypeEnumIgnoreCase := make(map[string]SubscriptionPlanTypeEnum)
	for k, v := range mappingSubscriptionPlanTypeEnum {
		mappingSubscriptionPlanTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSubscriptionPlanTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SubscriptionUpgradeStateEnum Enum with underlying type: string
type SubscriptionUpgradeStateEnum string

// Set of constants representing the allowable values for SubscriptionUpgradeStateEnum
const (
	SubscriptionUpgradeStatePromo     SubscriptionUpgradeStateEnum = "PROMO"
	SubscriptionUpgradeStateSubmitted SubscriptionUpgradeStateEnum = "SUBMITTED"
	SubscriptionUpgradeStateError     SubscriptionUpgradeStateEnum = "ERROR"
	SubscriptionUpgradeStateUpgraded  SubscriptionUpgradeStateEnum = "UPGRADED"
)

var mappingSubscriptionUpgradeStateEnum = map[string]SubscriptionUpgradeStateEnum{
	"PROMO":     SubscriptionUpgradeStatePromo,
	"SUBMITTED": SubscriptionUpgradeStateSubmitted,
	"ERROR":     SubscriptionUpgradeStateError,
	"UPGRADED":  SubscriptionUpgradeStateUpgraded,
}

// GetSubscriptionUpgradeStateEnumValues Enumerates the set of values for SubscriptionUpgradeStateEnum
func GetSubscriptionUpgradeStateEnumValues() []SubscriptionUpgradeStateEnum {
	values := make([]SubscriptionUpgradeStateEnum, 0)
	for _, v := range mappingSubscriptionUpgradeStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionUpgradeStateEnumStringValues Enumerates the set of values in String for SubscriptionUpgradeStateEnum
func GetSubscriptionUpgradeStateEnumStringValues() []string {
	return []string{
		"PROMO",
		"SUBMITTED",
		"ERROR",
		"UPGRADED",
	}
}

// GetMappingSubscriptionUpgradeStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionUpgradeStateEnum(val string) (SubscriptionUpgradeStateEnum, bool) {
	mappingSubscriptionUpgradeStateEnumIgnoreCase := make(map[string]SubscriptionUpgradeStateEnum)
	for k, v := range mappingSubscriptionUpgradeStateEnum {
		mappingSubscriptionUpgradeStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSubscriptionUpgradeStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SubscriptionUpgradeStateDetailsEnum Enum with underlying type: string
type SubscriptionUpgradeStateDetailsEnum string

// Set of constants representing the allowable values for SubscriptionUpgradeStateDetailsEnum
const (
	SubscriptionUpgradeStateDetailsTaxError     SubscriptionUpgradeStateDetailsEnum = "TAX_ERROR"
	SubscriptionUpgradeStateDetailsUpgradeError SubscriptionUpgradeStateDetailsEnum = "UPGRADE_ERROR"
)

var mappingSubscriptionUpgradeStateDetailsEnum = map[string]SubscriptionUpgradeStateDetailsEnum{
	"TAX_ERROR":     SubscriptionUpgradeStateDetailsTaxError,
	"UPGRADE_ERROR": SubscriptionUpgradeStateDetailsUpgradeError,
}

// GetSubscriptionUpgradeStateDetailsEnumValues Enumerates the set of values for SubscriptionUpgradeStateDetailsEnum
func GetSubscriptionUpgradeStateDetailsEnumValues() []SubscriptionUpgradeStateDetailsEnum {
	values := make([]SubscriptionUpgradeStateDetailsEnum, 0)
	for _, v := range mappingSubscriptionUpgradeStateDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionUpgradeStateDetailsEnumStringValues Enumerates the set of values in String for SubscriptionUpgradeStateDetailsEnum
func GetSubscriptionUpgradeStateDetailsEnumStringValues() []string {
	return []string{
		"TAX_ERROR",
		"UPGRADE_ERROR",
	}
}

// GetMappingSubscriptionUpgradeStateDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionUpgradeStateDetailsEnum(val string) (SubscriptionUpgradeStateDetailsEnum, bool) {
	mappingSubscriptionUpgradeStateDetailsEnumIgnoreCase := make(map[string]SubscriptionUpgradeStateDetailsEnum)
	for k, v := range mappingSubscriptionUpgradeStateDetailsEnum {
		mappingSubscriptionUpgradeStateDetailsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSubscriptionUpgradeStateDetailsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
