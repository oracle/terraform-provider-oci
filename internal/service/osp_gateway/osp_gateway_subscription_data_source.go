// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v65/ospgateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OspGatewaySubscriptionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["osp_home_region"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["subscription_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OspGatewaySubscriptionResource(), fieldMap, readSingularOspGatewaySubscription)
}

func readSingularOspGatewaySubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewaySubscriptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscriptionServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewaySubscriptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.SubscriptionServiceClient
	Res    *oci_osp_gateway.GetSubscriptionResponse
}

func (s *OspGatewaySubscriptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewaySubscriptionDataSourceCrud) Get() error {
	request := oci_osp_gateway.GetSubscriptionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osp_gateway")

	response, err := s.Client.GetSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OspGatewaySubscriptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("account_type", s.Res.AccountType)

	if s.Res.BillToCustAccountId != nil {
		s.D.Set("bill_to_cust_account_id", *s.Res.BillToCustAccountId)
	}

	if s.Res.BillingAddress != nil {
		s.D.Set("billing_address", []interface{}{AddressToMap(s.Res.BillingAddress)})
	} else {
		s.D.Set("billing_address", nil)
	}

	if s.Res.CurrencyCode != nil {
		s.D.Set("currency_code", *s.Res.CurrencyCode)
	}

	if s.Res.GsiOrgCode != nil {
		s.D.Set("gsi_org_code", *s.Res.GsiOrgCode)
	}

	if s.Res.IsIntentToPay != nil {
		s.D.Set("is_intent_to_pay", *s.Res.IsIntentToPay)
	}

	if s.Res.LanguageCode != nil {
		s.D.Set("language_code", *s.Res.LanguageCode)
	}

	if s.Res.OrganizationId != nil {
		s.D.Set("organization_id", *s.Res.OrganizationId)
	}

	if s.Res.PaymentGateway != nil {
		s.D.Set("payment_gateway", []interface{}{PaymentGatewayToMap(s.Res.PaymentGateway)})
	} else {
		s.D.Set("payment_gateway", nil)
	}

	paymentOptions := []interface{}{}
	for _, item := range s.Res.PaymentOptions {
		paymentOptions = append(paymentOptions, PaymentOptionToMap(item))
	}
	s.D.Set("payment_options", paymentOptions)

	s.D.Set("plan_type", s.Res.PlanType)

	if s.Res.ShipToCustAcctRoleId != nil {
		s.D.Set("ship_to_cust_acct_role_id", *s.Res.ShipToCustAcctRoleId)
	}

	if s.Res.ShipToCustAcctSiteId != nil {
		s.D.Set("ship_to_cust_acct_site_id", *s.Res.ShipToCustAcctSiteId)
	}

	if s.Res.SubscriptionPlanNumber != nil {
		s.D.Set("subscription_plan_number", *s.Res.SubscriptionPlanNumber)
	}

	if s.Res.TaxInfo != nil {
		s.D.Set("tax_info", []interface{}{TaxInfoToMap(s.Res.TaxInfo)})
	} else {
		s.D.Set("tax_info", nil)
	}

	if s.Res.TimePersonalToCorporateConv != nil {
		s.D.Set("time_personal_to_corporate_conv", s.Res.TimePersonalToCorporateConv.String())
	}

	if s.Res.TimePlanUpgrade != nil {
		s.D.Set("time_plan_upgrade", s.Res.TimePlanUpgrade.String())
	}

	if s.Res.TimeStart != nil {
		s.D.Set("time_start", s.Res.TimeStart.String())
	}

	s.D.Set("upgrade_state", s.Res.UpgradeState)

	s.D.Set("upgrade_state_details", s.Res.UpgradeStateDetails)

	return nil
}
