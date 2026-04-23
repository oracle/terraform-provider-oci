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

func CostadCostAlertSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createCostadCostAlertSubscriptionWithContext,
		ReadContext:   readCostadCostAlertSubscriptionWithContext,
		UpdateContext: updateCostadCostAlertSubscriptionWithContext,
		DeleteContext: deleteCostadCostAlertSubscriptionWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"channels": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"cost_anomaly_monitors": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCostadCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readCostadCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateCostadCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteCostadCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type CostadCostAlertSubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_costad.CostAdClient
	Res                    *oci_costad.CostAlertSubscription
	DisableNotFoundRetries bool
}

func (s *CostadCostAlertSubscriptionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CostadCostAlertSubscriptionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *CostadCostAlertSubscriptionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_costad.CostAlertSubscriptionLifecycleStateActive),
	}
}

func (s *CostadCostAlertSubscriptionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CostadCostAlertSubscriptionResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *CostadCostAlertSubscriptionResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_costad.CreateCostAlertSubscriptionRequest{}

	if channels, ok := s.D.GetOkExists("channels"); ok {
		tmp := channels.(string)
		request.Channels = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.CreateCostAlertSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAlertSubscription
	return nil
}

func (s *CostadCostAlertSubscriptionResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_costad.GetCostAlertSubscriptionRequest{}

	tmp := s.D.Id()
	request.CostAlertSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.GetCostAlertSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAlertSubscription
	return nil
}

func (s *CostadCostAlertSubscriptionResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_costad.UpdateCostAlertSubscriptionRequest{}

	if channels, ok := s.D.GetOkExists("channels"); ok {
		tmp := channels.(string)
		request.Channels = &tmp
	}

	tmp := s.D.Id()
	request.CostAlertSubscriptionId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.UpdateCostAlertSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAlertSubscription
	return nil
}

func (s *CostadCostAlertSubscriptionResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_costad.DeleteCostAlertSubscriptionRequest{}

	tmp := s.D.Id()
	request.CostAlertSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	_, err := s.Client.DeleteCostAlertSubscription(ctx, request)
	return err
}

func (s *CostadCostAlertSubscriptionResourceCrud) SetData() error {
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

func CostAlertSubscriptionSummaryToMap(obj oci_costad.CostAlertSubscriptionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChannelTypes != nil {
		result["channels"] = string(*obj.ChannelTypes)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
