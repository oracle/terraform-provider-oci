// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package compute_cloud_at_customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ComputeCloudAtCustomerCccUpgradeSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readComputeCloudAtCustomerCccUpgradeSchedules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ccc_upgrade_schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ccc_upgrade_schedule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ComputeCloudAtCustomerCccUpgradeScheduleResource()),
						},
					},
				},
			},
		},
	}
}

func readComputeCloudAtCustomerCccUpgradeSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccUpgradeSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.ReadResource(sync)
}

type ComputeCloudAtCustomerCccUpgradeSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient
	Res    *oci_compute_cloud_at_customer.ListCccUpgradeSchedulesResponse
}

func (s *ComputeCloudAtCustomerCccUpgradeSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ComputeCloudAtCustomerCccUpgradeSchedulesDataSourceCrud) Get() error {
	request := oci_compute_cloud_at_customer.ListCccUpgradeSchedulesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_compute_cloud_at_customer.ListCccUpgradeSchedulesAccessLevelEnum(accessLevel.(string))
	}

	if cccUpgradeScheduleId, ok := s.D.GetOkExists("id"); ok {
		tmp := cccUpgradeScheduleId.(string)
		request.CccUpgradeScheduleId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "compute_cloud_at_customer")

	response, err := s.Client.ListCccUpgradeSchedules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCccUpgradeSchedules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ComputeCloudAtCustomerCccUpgradeSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ComputeCloudAtCustomerCccUpgradeSchedulesDataSource-", ComputeCloudAtCustomerCccUpgradeSchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	cccUpgradeSchedule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CccUpgradeScheduleSummaryToMap(item))
	}
	cccUpgradeSchedule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ComputeCloudAtCustomerCccUpgradeSchedulesDataSource().Schema["ccc_upgrade_schedule_collection"].Elem.(*schema.Resource).Schema)
		cccUpgradeSchedule["items"] = items
	}

	resources = append(resources, cccUpgradeSchedule)
	if err := s.D.Set("ccc_upgrade_schedule_collection", resources); err != nil {
		return err
	}

	return nil
}
