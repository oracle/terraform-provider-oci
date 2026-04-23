// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package costad

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_costad "github.com/oracle/oci-go-sdk/v65/costad"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CostadCostAlertSubscriptionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cost_alert_subscription_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(CostadCostAlertSubscriptionResource(), fieldMap, readSingularCostadCostAlertSubscriptionWithContext)
}

func readSingularCostadCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAlertSubscriptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type CostadCostAlertSubscriptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_costad.CostAdClient
	Res    *oci_costad.GetCostAlertSubscriptionResponse
}

func (s *CostadCostAlertSubscriptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CostadCostAlertSubscriptionDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_costad.GetCostAlertSubscriptionRequest{}

	if costAlertSubscriptionId, ok := s.D.GetOkExists("cost_alert_subscription_id"); ok {
		tmp := costAlertSubscriptionId.(string)
		request.CostAlertSubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "costad")

	response, err := s.Client.GetCostAlertSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CostadCostAlertSubscriptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Channels != nil {
		s.D.Set("channels", *s.Res.Channels)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CostAnomalyMonitors != nil {
		costAnomalyMonitorsJSON, err := json.Marshal(s.Res.CostAnomalyMonitors)
		if err != nil {
			return err
		}
		s.D.Set("cost_anomaly_monitors", string(costAnomalyMonitorsJSON))
	} else {
		s.D.Set("cost_anomaly_monitors", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
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
