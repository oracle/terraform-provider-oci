// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package self

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_self "github.com/oracle/oci-go-sdk/v65/self"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func SelfSubscriptionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subscription_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(SelfSubscriptionResource(), fieldMap, readSingularSelfSubscriptionWithContext)
}

func readSingularSelfSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &SelfSubscriptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SelfSubscriptionClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type SelfSubscriptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_self.SubscriptionClient
	Res    *oci_self.GetSubscriptionResponse
}

func (s *SelfSubscriptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SelfSubscriptionDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_self.GetSubscriptionRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "self")

	response, err := s.Client.GetSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SelfSubscriptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	additionalDetails := []interface{}{}
	for _, item := range s.Res.AdditionalDetails {
		additionalDetails = append(additionalDetails, ExtendedMetadataToMap(item))
	}
	s.D.Set("additional_details", additionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.ProductId != nil {
		s.D.Set("product_id", *s.Res.ProductId)
	}

	if s.Res.Realm != nil {
		s.D.Set("realm", *s.Res.Realm)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.SellerId != nil {
		s.D.Set("seller_id", *s.Res.SellerId)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubscriptionDetails != nil {
		s.D.Set("subscription_details", []interface{}{SubscriptionDetailsToMap(s.Res.SubscriptionDetails)})
	} else {
		s.D.Set("subscription_details", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
