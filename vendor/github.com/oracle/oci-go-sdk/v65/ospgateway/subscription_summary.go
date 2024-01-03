// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscriptionSummary Subscription object which contains the common subscription data.
type SubscriptionSummary struct {

	// Subscription plan number.
	SubscriptionPlanNumber *string `mandatory:"true" json:"subscriptionPlanNumber"`

	// Subscription id identifier (OCID).
	Id *string `mandatory:"false" json:"id"`

	// Subscription plan type.
	PlanType SubscriptionSummaryPlanTypeEnum `mandatory:"false" json:"planType,omitempty"`

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
	UpgradeState SubscriptionSummaryUpgradeStateEnum `mandatory:"false" json:"upgradeState,omitempty"`

	// This field is used to describe the Upgrade State in case of error (E.g. Upgrade failure caused by interfacing Tax details- TaxError)
	UpgradeStateDetails SubscriptionSummaryUpgradeStateDetailsEnum `mandatory:"false" json:"upgradeStateDetails,omitempty"`

	// Account type.
	AccountType SubscriptionSummaryAccountTypeEnum `mandatory:"false" json:"accountType,omitempty"`

	TaxInfo *TaxInfo `mandatory:"false" json:"taxInfo"`

	// Payment option list of a subscription.
	PaymentOptions []PaymentOption `mandatory:"false" json:"paymentOptions"`

	PaymentGateway *PaymentGateway `mandatory:"false" json:"paymentGateway"`

	BillingAddress *Address `mandatory:"false" json:"billingAddress"`

	// Date of upgrade/conversion when planType changed from FREE_TIER to PAYG
	TimePlanUpgrade *common.SDKTime `mandatory:"false" json:"timePlanUpgrade"`

	// Date of upgrade/conversion when account type changed from PERSONAL to CORPORATE
	TimePersonalToCorporateConv *common.SDKTime `mandatory:"false" json:"timePersonalToCorporateConv"`
}

func (m SubscriptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubscriptionSummaryPlanTypeEnum(string(m.PlanType)); !ok && m.PlanType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanType: %s. Supported values are: %s.", m.PlanType, strings.Join(GetSubscriptionSummaryPlanTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubscriptionSummaryUpgradeStateEnum(string(m.UpgradeState)); !ok && m.UpgradeState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeState: %s. Supported values are: %s.", m.UpgradeState, strings.Join(GetSubscriptionSummaryUpgradeStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubscriptionSummaryUpgradeStateDetailsEnum(string(m.UpgradeStateDetails)); !ok && m.UpgradeStateDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeStateDetails: %s. Supported values are: %s.", m.UpgradeStateDetails, strings.Join(GetSubscriptionSummaryUpgradeStateDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubscriptionSummaryAccountTypeEnum(string(m.AccountType)); !ok && m.AccountType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccountType: %s. Supported values are: %s.", m.AccountType, strings.Join(GetSubscriptionSummaryAccountTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SubscriptionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                          *string                                    `json:"id"`
		PlanType                    SubscriptionSummaryPlanTypeEnum            `json:"planType"`
		TimeStart                   *common.SDKTime                            `json:"timeStart"`
		ShipToCustAcctSiteId        *string                                    `json:"shipToCustAcctSiteId"`
		ShipToCustAcctRoleId        *string                                    `json:"shipToCustAcctRoleId"`
		BillToCustAccountId         *string                                    `json:"billToCustAccountId"`
		IsIntentToPay               *bool                                      `json:"isIntentToPay"`
		CurrencyCode                *string                                    `json:"currencyCode"`
		GsiOrgCode                  *string                                    `json:"gsiOrgCode"`
		LanguageCode                *string                                    `json:"languageCode"`
		OrganizationId              *string                                    `json:"organizationId"`
		UpgradeState                SubscriptionSummaryUpgradeStateEnum        `json:"upgradeState"`
		UpgradeStateDetails         SubscriptionSummaryUpgradeStateDetailsEnum `json:"upgradeStateDetails"`
		AccountType                 SubscriptionSummaryAccountTypeEnum         `json:"accountType"`
		TaxInfo                     *TaxInfo                                   `json:"taxInfo"`
		PaymentOptions              []paymentoption                            `json:"paymentOptions"`
		PaymentGateway              *PaymentGateway                            `json:"paymentGateway"`
		BillingAddress              *Address                                   `json:"billingAddress"`
		TimePlanUpgrade             *common.SDKTime                            `json:"timePlanUpgrade"`
		TimePersonalToCorporateConv *common.SDKTime                            `json:"timePersonalToCorporateConv"`
		SubscriptionPlanNumber      *string                                    `json:"subscriptionPlanNumber"`
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

	m.AccountType = model.AccountType

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

	m.TimePersonalToCorporateConv = model.TimePersonalToCorporateConv

	m.SubscriptionPlanNumber = model.SubscriptionPlanNumber

	return
}

// SubscriptionSummaryPlanTypeEnum Enum with underlying type: string
type SubscriptionSummaryPlanTypeEnum string

// Set of constants representing the allowable values for SubscriptionSummaryPlanTypeEnum
const (
	SubscriptionSummaryPlanTypeFreeTier SubscriptionSummaryPlanTypeEnum = "FREE_TIER"
	SubscriptionSummaryPlanTypePayg     SubscriptionSummaryPlanTypeEnum = "PAYG"
)

var mappingSubscriptionSummaryPlanTypeEnum = map[string]SubscriptionSummaryPlanTypeEnum{
	"FREE_TIER": SubscriptionSummaryPlanTypeFreeTier,
	"PAYG":      SubscriptionSummaryPlanTypePayg,
}

var mappingSubscriptionSummaryPlanTypeEnumLowerCase = map[string]SubscriptionSummaryPlanTypeEnum{
	"free_tier": SubscriptionSummaryPlanTypeFreeTier,
	"payg":      SubscriptionSummaryPlanTypePayg,
}

// GetSubscriptionSummaryPlanTypeEnumValues Enumerates the set of values for SubscriptionSummaryPlanTypeEnum
func GetSubscriptionSummaryPlanTypeEnumValues() []SubscriptionSummaryPlanTypeEnum {
	values := make([]SubscriptionSummaryPlanTypeEnum, 0)
	for _, v := range mappingSubscriptionSummaryPlanTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionSummaryPlanTypeEnumStringValues Enumerates the set of values in String for SubscriptionSummaryPlanTypeEnum
func GetSubscriptionSummaryPlanTypeEnumStringValues() []string {
	return []string{
		"FREE_TIER",
		"PAYG",
	}
}

// GetMappingSubscriptionSummaryPlanTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionSummaryPlanTypeEnum(val string) (SubscriptionSummaryPlanTypeEnum, bool) {
	enum, ok := mappingSubscriptionSummaryPlanTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubscriptionSummaryUpgradeStateEnum Enum with underlying type: string
type SubscriptionSummaryUpgradeStateEnum string

// Set of constants representing the allowable values for SubscriptionSummaryUpgradeStateEnum
const (
	SubscriptionSummaryUpgradeStatePromo     SubscriptionSummaryUpgradeStateEnum = "PROMO"
	SubscriptionSummaryUpgradeStateSubmitted SubscriptionSummaryUpgradeStateEnum = "SUBMITTED"
	SubscriptionSummaryUpgradeStateError     SubscriptionSummaryUpgradeStateEnum = "ERROR"
	SubscriptionSummaryUpgradeStateUpgraded  SubscriptionSummaryUpgradeStateEnum = "UPGRADED"
)

var mappingSubscriptionSummaryUpgradeStateEnum = map[string]SubscriptionSummaryUpgradeStateEnum{
	"PROMO":     SubscriptionSummaryUpgradeStatePromo,
	"SUBMITTED": SubscriptionSummaryUpgradeStateSubmitted,
	"ERROR":     SubscriptionSummaryUpgradeStateError,
	"UPGRADED":  SubscriptionSummaryUpgradeStateUpgraded,
}

var mappingSubscriptionSummaryUpgradeStateEnumLowerCase = map[string]SubscriptionSummaryUpgradeStateEnum{
	"promo":     SubscriptionSummaryUpgradeStatePromo,
	"submitted": SubscriptionSummaryUpgradeStateSubmitted,
	"error":     SubscriptionSummaryUpgradeStateError,
	"upgraded":  SubscriptionSummaryUpgradeStateUpgraded,
}

// GetSubscriptionSummaryUpgradeStateEnumValues Enumerates the set of values for SubscriptionSummaryUpgradeStateEnum
func GetSubscriptionSummaryUpgradeStateEnumValues() []SubscriptionSummaryUpgradeStateEnum {
	values := make([]SubscriptionSummaryUpgradeStateEnum, 0)
	for _, v := range mappingSubscriptionSummaryUpgradeStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionSummaryUpgradeStateEnumStringValues Enumerates the set of values in String for SubscriptionSummaryUpgradeStateEnum
func GetSubscriptionSummaryUpgradeStateEnumStringValues() []string {
	return []string{
		"PROMO",
		"SUBMITTED",
		"ERROR",
		"UPGRADED",
	}
}

// GetMappingSubscriptionSummaryUpgradeStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionSummaryUpgradeStateEnum(val string) (SubscriptionSummaryUpgradeStateEnum, bool) {
	enum, ok := mappingSubscriptionSummaryUpgradeStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubscriptionSummaryUpgradeStateDetailsEnum Enum with underlying type: string
type SubscriptionSummaryUpgradeStateDetailsEnum string

// Set of constants representing the allowable values for SubscriptionSummaryUpgradeStateDetailsEnum
const (
	SubscriptionSummaryUpgradeStateDetailsTaxError     SubscriptionSummaryUpgradeStateDetailsEnum = "TAX_ERROR"
	SubscriptionSummaryUpgradeStateDetailsUpgradeError SubscriptionSummaryUpgradeStateDetailsEnum = "UPGRADE_ERROR"
)

var mappingSubscriptionSummaryUpgradeStateDetailsEnum = map[string]SubscriptionSummaryUpgradeStateDetailsEnum{
	"TAX_ERROR":     SubscriptionSummaryUpgradeStateDetailsTaxError,
	"UPGRADE_ERROR": SubscriptionSummaryUpgradeStateDetailsUpgradeError,
}

var mappingSubscriptionSummaryUpgradeStateDetailsEnumLowerCase = map[string]SubscriptionSummaryUpgradeStateDetailsEnum{
	"tax_error":     SubscriptionSummaryUpgradeStateDetailsTaxError,
	"upgrade_error": SubscriptionSummaryUpgradeStateDetailsUpgradeError,
}

// GetSubscriptionSummaryUpgradeStateDetailsEnumValues Enumerates the set of values for SubscriptionSummaryUpgradeStateDetailsEnum
func GetSubscriptionSummaryUpgradeStateDetailsEnumValues() []SubscriptionSummaryUpgradeStateDetailsEnum {
	values := make([]SubscriptionSummaryUpgradeStateDetailsEnum, 0)
	for _, v := range mappingSubscriptionSummaryUpgradeStateDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionSummaryUpgradeStateDetailsEnumStringValues Enumerates the set of values in String for SubscriptionSummaryUpgradeStateDetailsEnum
func GetSubscriptionSummaryUpgradeStateDetailsEnumStringValues() []string {
	return []string{
		"TAX_ERROR",
		"UPGRADE_ERROR",
	}
}

// GetMappingSubscriptionSummaryUpgradeStateDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionSummaryUpgradeStateDetailsEnum(val string) (SubscriptionSummaryUpgradeStateDetailsEnum, bool) {
	enum, ok := mappingSubscriptionSummaryUpgradeStateDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubscriptionSummaryAccountTypeEnum Enum with underlying type: string
type SubscriptionSummaryAccountTypeEnum string

// Set of constants representing the allowable values for SubscriptionSummaryAccountTypeEnum
const (
	SubscriptionSummaryAccountTypePersonal           SubscriptionSummaryAccountTypeEnum = "PERSONAL"
	SubscriptionSummaryAccountTypeCorporate          SubscriptionSummaryAccountTypeEnum = "CORPORATE"
	SubscriptionSummaryAccountTypeCorporateSubmitted SubscriptionSummaryAccountTypeEnum = "CORPORATE_SUBMITTED"
)

var mappingSubscriptionSummaryAccountTypeEnum = map[string]SubscriptionSummaryAccountTypeEnum{
	"PERSONAL":            SubscriptionSummaryAccountTypePersonal,
	"CORPORATE":           SubscriptionSummaryAccountTypeCorporate,
	"CORPORATE_SUBMITTED": SubscriptionSummaryAccountTypeCorporateSubmitted,
}

var mappingSubscriptionSummaryAccountTypeEnumLowerCase = map[string]SubscriptionSummaryAccountTypeEnum{
	"personal":            SubscriptionSummaryAccountTypePersonal,
	"corporate":           SubscriptionSummaryAccountTypeCorporate,
	"corporate_submitted": SubscriptionSummaryAccountTypeCorporateSubmitted,
}

// GetSubscriptionSummaryAccountTypeEnumValues Enumerates the set of values for SubscriptionSummaryAccountTypeEnum
func GetSubscriptionSummaryAccountTypeEnumValues() []SubscriptionSummaryAccountTypeEnum {
	values := make([]SubscriptionSummaryAccountTypeEnum, 0)
	for _, v := range mappingSubscriptionSummaryAccountTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionSummaryAccountTypeEnumStringValues Enumerates the set of values in String for SubscriptionSummaryAccountTypeEnum
func GetSubscriptionSummaryAccountTypeEnumStringValues() []string {
	return []string{
		"PERSONAL",
		"CORPORATE",
		"CORPORATE_SUBMITTED",
	}
}

// GetMappingSubscriptionSummaryAccountTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionSummaryAccountTypeEnum(val string) (SubscriptionSummaryAccountTypeEnum, bool) {
	enum, ok := mappingSubscriptionSummaryAccountTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
