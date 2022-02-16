// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package appmgmt_control

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_appmgmt_control "github.com/oracle/oci-go-sdk/v58/appmgmtcontrol"
)

func AppmgmtControlMonitoredInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAppmgmtControlMonitoredInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitored_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"management_agent_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"monitoring_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
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
							},
						},
					},
				},
			},
		},
	}
}

func readAppmgmtControlMonitoredInstances(d *schema.ResourceData, m interface{}) error {
	sync := &AppmgmtControlMonitoredInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AppmgmtControlClient()

	return tfresource.ReadResource(sync)
}

type AppmgmtControlMonitoredInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_appmgmt_control.AppmgmtControlClient
	Res    *oci_appmgmt_control.ListMonitoredInstancesResponse
}

func (s *AppmgmtControlMonitoredInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AppmgmtControlMonitoredInstancesDataSourceCrud) Get() error {
	request := oci_appmgmt_control.ListMonitoredInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "appmgmt_control")

	response, err := s.Client.ListMonitoredInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMonitoredInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AppmgmtControlMonitoredInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AppmgmtControlMonitoredInstancesDataSource-", AppmgmtControlMonitoredInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitoredInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredInstanceSummaryToMap(item))
	}
	monitoredInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AppmgmtControlMonitoredInstancesDataSource().Schema["monitored_instance_collection"].Elem.(*schema.Resource).Schema)
		monitoredInstance["items"] = items
	}

	resources = append(resources, monitoredInstance)
	if err := s.D.Set("monitored_instance_collection", resources); err != nil {
		return err
	}

	return nil
}

func MonitoredInstanceSummaryToMap(obj oci_appmgmt_control.MonitoredInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.ManagementAgentId != nil {
		result["management_agent_id"] = string(*obj.ManagementAgentId)
	}

	result["monitoring_state"] = string(obj.MonitoringState)

	result["state"] = string(obj.LifecycleState)

	return result
}
