// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeDiscoverySchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudBridgeDiscoverySchedules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"discovery_schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_schedule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudBridgeDiscoveryScheduleResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudBridgeDiscoverySchedules(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeDiscoverySchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeDiscoverySchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.DiscoveryClient
	Res    *oci_cloud_bridge.ListDiscoverySchedulesResponse
}

func (s *CloudBridgeDiscoverySchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeDiscoverySchedulesDataSourceCrud) Get() error {
	request := oci_cloud_bridge.ListDiscoverySchedulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if discoveryScheduleId, ok := s.D.GetOkExists("id"); ok {
		tmp := discoveryScheduleId.(string)
		request.DiscoveryScheduleId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_bridge.ListDiscoverySchedulesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.ListDiscoverySchedules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDiscoverySchedules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudBridgeDiscoverySchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeDiscoverySchedulesDataSource-", CloudBridgeDiscoverySchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	discoverySchedule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiscoveryScheduleSummaryToMap(item))
	}
	discoverySchedule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudBridgeDiscoverySchedulesDataSource().Schema["discovery_schedule_collection"].Elem.(*schema.Resource).Schema)
		discoverySchedule["items"] = items
	}

	resources = append(resources, discoverySchedule)
	if err := s.D.Set("discovery_schedule_collection", resources); err != nil {
		return err
	}

	return nil
}
