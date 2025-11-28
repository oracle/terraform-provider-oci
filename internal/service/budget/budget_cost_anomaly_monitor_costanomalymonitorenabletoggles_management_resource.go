// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
		Read:     readBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
		Update:   updateBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
		Delete:   deleteBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement,
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

			// Optional

			// Computed
			//"compartment_id": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"cost_alert_subscription_map": {
			//	Type:     schema.TypeList,
			//	Computed: true,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			// Required
			//			"enable_costanomalymonitorenabletoggle": {
			//				Type:     schema.TypeBool,
			//				Required: true,
			//			},
			//
			//			// Optional
			//
			//			// Computed
			//			"cost_alert_subscription_id": {
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"operator": {
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"threshold_absolute_value": {
			//				Type:     schema.TypeInt,
			//				Computed: true,
			//			},
			//			"threshold_relative_percent": {
			//				Type:     schema.TypeInt,
			//				Computed: true,
			//			},
			//		},
			//	},
			//},
			//"defined_tags": {
			//	Type:     schema.TypeMap,
			//	Computed: true,
			//	Elem:     schema.TypeString,
			//},
			//"description": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"freeform_tags": {
			//	Type:     schema.TypeMap,
			//	Computed: true,
			//	Elem:     schema.TypeString,
			//},
			//"lifecycle_details": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"name": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"state": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"system_tags": {
			//	Type:     schema.TypeMap,
			//	Computed: true,
			//	Elem:     schema.TypeString,
			//},
			//"target_resource_filter": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"time_created": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"time_updated": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			//"type": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
		},
	}
}

func createBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()
	sync.Res = &BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(d *schema.ResourceData, m interface{}) error {
	if d.Id() == "" {
		d.SetId(tfresource.GenerateDataSourceHashID("BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource-", BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource(), d))
	}
	return nil
	//return nil
}

func updateBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()
	sync.Res = &BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()
	sync.Res = &BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse struct {
	enableResponse  *oci_budget.EnableCostAnomalyMonitorResponse
	disableResponse *oci_budget.DisableCostAnomalyMonitorResponse
}

type BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_budget.CostAdClient
	Res                    *BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResponse
	DisableNotFoundRetries bool
}

func (s *BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource-", BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource(), s.D)
}

func (s *BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_costanomalymonitorenabletoggle"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_budget.EnableCostAnomalyMonitorRequest{}

		if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
			tmp := costAnomalyMonitorId.(string)
			request.CostAnomalyMonitorId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

		response, err := s.Client.EnableCostAnomalyMonitor(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_budget.DisableCostAnomalyMonitorRequest{}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.DisableCostAnomalyMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_costanomalymonitorenabletoggle"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_budget.EnableCostAnomalyMonitorRequest{}

		if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
			tmp := costAnomalyMonitorId.(string)
			request.CostAnomalyMonitorId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

		response, err := s.Client.EnableCostAnomalyMonitor(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_budget.DisableCostAnomalyMonitorRequest{}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.DisableCostAnomalyMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_costanomalymonitorenabletoggle"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_budget.DisableCostAnomalyMonitorRequest{}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.DisableCostAnomalyMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceCrud) SetData() error {
	return nil
}
