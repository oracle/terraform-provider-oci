// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiChargebackPlanDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["chargebackplan_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(OpsiChargebackPlanResource(), fieldMap, readSingularOpsiChargebackPlanWithContext)
}

func readSingularOpsiChargebackPlanWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OpsiChargebackPlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OpsiChargebackPlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetChargebackPlanResponse
}

func (s *OpsiChargebackPlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiChargebackPlanDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_opsi.GetChargebackPlanRequest{}

	if chargebackplanId, ok := s.D.GetOkExists("chargebackplan_id"); ok {
		tmp := chargebackplanId.(string)
		request.ChargebackplanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetChargebackPlan(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiChargebackPlanDataSourceCrud) SetData() error {
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

	s.D.Set("entity_source", s.Res.EntitySource)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCustomizable != nil {
		s.D.Set("is_customizable", *s.Res.IsCustomizable)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("plan_category", s.Res.PlanCategory)

	planCustomItems := []interface{}{}
	for _, item := range s.Res.PlanCustomItems {
		planCustomItems = append(planCustomItems, CreatePlanCustomItemDetailsToMap(item))
	}
	s.D.Set("plan_custom_items", planCustomItems)

	if s.Res.PlanDescription != nil {
		s.D.Set("plan_description", *s.Res.PlanDescription)
	}

	if s.Res.PlanName != nil {
		s.D.Set("plan_name", *s.Res.PlanName)
	}

	if s.Res.PlanType != nil {
		s.D.Set("plan_type", *s.Res.PlanType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
