// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApigatewayUsagePlanDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["usage_plan_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApigatewayUsagePlanResource(), fieldMap, readSingularApigatewayUsagePlan)
}

func readSingularApigatewayUsagePlan(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayUsagePlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsagePlansClient()

	return tfresource.ReadResource(sync)
}

type ApigatewayUsagePlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.UsagePlansClient
	Res    *oci_apigateway.GetUsagePlanResponse
}

func (s *ApigatewayUsagePlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayUsagePlanDataSourceCrud) Get() error {
	request := oci_apigateway.GetUsagePlanRequest{}

	if usagePlanId, ok := s.D.GetOkExists("usage_plan_id"); ok {
		tmp := usagePlanId.(string)
		request.UsagePlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.GetUsagePlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayUsagePlanDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	entitlements := []interface{}{}
	for _, item := range s.Res.Entitlements {
		entitlements = append(entitlements, EntitlementToMap(item))
	}
	s.D.Set("entitlements", entitlements)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
