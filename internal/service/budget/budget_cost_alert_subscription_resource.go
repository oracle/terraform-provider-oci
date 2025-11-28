// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BudgetCostAlertSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBudgetCostAlertSubscriptionWithContext,
		ReadContext:   readBudgetCostAlertSubscriptionWithContext,
		UpdateContext: updateBudgetCostAlertSubscriptionWithContext,
		DeleteContext: deleteBudgetCostAlertSubscriptionWithContext,
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

func mapToCostAnomalyMonitors(cost_anomaly_monitors string) (*interface{}, error) {
	var result interface{}
	var err error

	var obj interface{}
	err = json.Unmarshal([]byte(cost_anomaly_monitors), &obj)
	result = &obj

	return &result, err
}

func CostAnomalyMonitorsToMap(obj *interface{}) string {
	var result string

	if obj != nil {
		var bytes, _ = json.Marshal(obj)
		result = string(bytes)
	}

	return result
}

func createBudgetCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readBudgetCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBudgetCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteBudgetCostAlertSubscriptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAlertSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BudgetCostAlertSubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_budget.CostAdClient
	Res                    *oci_budget.CostAlertSubscription
	DisableNotFoundRetries bool
}

func (s *BudgetCostAlertSubscriptionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BudgetCostAlertSubscriptionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BudgetCostAlertSubscriptionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_budget.CostAlertSubscriptionLifecycleStateActive),
	}
}

func (s *BudgetCostAlertSubscriptionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BudgetCostAlertSubscriptionResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *BudgetCostAlertSubscriptionResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_budget.CreateCostAlertSubscriptionRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.CreateCostAlertSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAlertSubscription
	return nil
}

func (s *BudgetCostAlertSubscriptionResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_budget.GetCostAlertSubscriptionRequest{}

	tmp := s.D.Id()
	request.CostAlertSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.GetCostAlertSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAlertSubscription
	return nil
}

func (s *BudgetCostAlertSubscriptionResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_budget.UpdateCostAlertSubscriptionRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.UpdateCostAlertSubscription(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAlertSubscription
	return nil
}

func (s *BudgetCostAlertSubscriptionResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_budget.DeleteCostAlertSubscriptionRequest{}

	tmp := s.D.Id()
	request.CostAlertSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	_, err := s.Client.DeleteCostAlertSubscription(ctx, request)
	return err
}

func (s *BudgetCostAlertSubscriptionResourceCrud) SetData() error {
	if s.Res.Channels != nil {
		s.D.Set("channels", *s.Res.Channels)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CostAnomalyMonitors != nil {
		//s.D.Set("cost_anomaly_monitors", []interface{}{CostAnomalyMonitorsToMap(s.Res.CostAnomalyMonitors)})
		costAnomalyMonitorsJson, err := json.Marshal(s.Res.CostAnomalyMonitors)
		if err != nil {
			return err
		}
		s.D.Set("cost_anomaly_monitors", string(costAnomalyMonitorsJson))
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

func CostAlertSubscriptionSummaryToMap(obj oci_budget.CostAlertSubscriptionSummary) map[string]interface{} {
	result := map[string]interface{}{}
	fmt.Printf("--- START DEBUG: CostAlertSubscriptionSummaryToMap for ID: %s ---\n", *obj.Id)
	if obj.ChannelTypes != nil {
		result["channels"] = string(*obj.ChannelTypes)
	}
	//if obj.ChannelTypes != nil {
	//	// --- START OF FIX ---
	//	channelTypesString := string(*obj.ChannelTypes)
	//	fmt.Printf("[DEBUG] Raw ChannelTypes string: %s\n", channelTypesString)
	//	var channelTypesSlice []string
	//
	//	// Decode the JSON string (e.g., "[\"email\"]") into a Go slice ([]string)
	//	if err := json.Unmarshal([]byte(channelTypesString), &channelTypesSlice); err != nil {
	//		// PANIC AVOIDED: If unmarshalling fails, we panic here to surface the bug,
	//		// but normally we would log or handle this gracefully. In a helper function,
	//		// panicking or returning a map without the value is often standard.
	//		// For stability, it's safer to return the empty slice or panic with a clear message.
	//		panic(fmt.Sprintf("Failed to unmarshal channel_types JSON '%s': %v", channelTypesString, err))
	//	}
	//	fmt.Printf("[DEBUG] Unmarshaled ChannelTypes slice: %v (Type: %T)\n", channelTypesSlice, channelTypesSlice)
	//	// Set the corrected slice value (e.g., []string{"email"})
	//	result["channel_types"] = channelTypesSlice
	//	// --- END OF FIX ---
	//}

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
