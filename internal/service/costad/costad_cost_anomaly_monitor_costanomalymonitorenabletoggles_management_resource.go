// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package costad

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_costad "github.com/oracle/oci-go-sdk/v65/costad"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
		ReadContext:   readCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
		UpdateContext: updateCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
		DeleteContext: deleteCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"cost_anomaly_monitor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_costanomalymonitorenabletoggle": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func createCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()
	sync.Res = &CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse{}

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.Id() == "" {
		d.SetId(tfresource.GenerateDataSourceHashID("CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource-", CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource(), d))
	}
	return nil
}

func updateCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()
	sync.Res = &CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse{}

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()
	sync.Res = &CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse struct {
	enableResponse  *oci_costad.EnableCostAnomalyMonitorResponse
	disableResponse *oci_costad.DisableCostAnomalyMonitorResponse
}

type CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_costad.CostAdClient
	Res                    *CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse
	DisableNotFoundRetries bool
}

func (s *CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource-", CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource(), s.D)
}

func (s *CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) CreateWithContext(ctx context.Context) error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_costanomalymonitorenabletoggle"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_costad.EnableCostAnomalyMonitorRequest{}

		if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
			tmp := costAnomalyMonitorId.(string)
			request.CostAnomalyMonitorId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

		response, err := s.Client.EnableCostAnomalyMonitor(ctx, request)
		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_costad.DisableCostAnomalyMonitorRequest{}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.DisableCostAnomalyMonitor(ctx, request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) UpdateWithContext(ctx context.Context) error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_costanomalymonitorenabletoggle"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_costad.EnableCostAnomalyMonitorRequest{}

		if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
			tmp := costAnomalyMonitorId.(string)
			request.CostAnomalyMonitorId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

		response, err := s.Client.EnableCostAnomalyMonitor(ctx, request)
		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_costad.DisableCostAnomalyMonitorRequest{}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.DisableCostAnomalyMonitor(ctx, request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) DeleteWithContext(ctx context.Context) error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_costanomalymonitorenabletoggle"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_costad.DisableCostAnomalyMonitorRequest{}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.DisableCostAnomalyMonitor(ctx, request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) SetData() error {
	return nil
}
