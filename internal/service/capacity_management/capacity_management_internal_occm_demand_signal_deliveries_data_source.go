// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementInternalOccmDemandSignalDeliveriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementInternalOccmDemandSignalDeliveries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_customer_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"occm_demand_signal_item_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_occm_demand_signal_delivery_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CapacityManagementInternalOccmDemandSignalDeliveryResource()),
						},
					},
				},
			},
		},
	}
}

func readCapacityManagementInternalOccmDemandSignalDeliveries(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalDeliveriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementInternalOccmDemandSignalDeliveriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.InternalDemandSignalClient
	Res    *oci_capacity_management.ListInternalOccmDemandSignalDeliveriesResponse
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveriesDataSourceCrud) Get() error {
	request := oci_capacity_management.ListInternalOccmDemandSignalDeliveriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	if occmDemandSignalItemId, ok := s.D.GetOkExists("occm_demand_signal_item_id"); ok {
		tmp := occmDemandSignalItemId.(string)
		request.OccmDemandSignalItemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListInternalOccmDemandSignalDeliveries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInternalOccmDemandSignalDeliveries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementInternalOccmDemandSignalDeliveriesDataSource-", CapacityManagementInternalOccmDemandSignalDeliveriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	internalOccmDemandSignalDelivery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InternalOccmDemandSignalDeliverySummaryToMap(item))
	}
	internalOccmDemandSignalDelivery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementInternalOccmDemandSignalDeliveriesDataSource().Schema["internal_occm_demand_signal_delivery_collection"].Elem.(*schema.Resource).Schema)
		internalOccmDemandSignalDelivery["items"] = items
	}

	resources = append(resources, internalOccmDemandSignalDelivery)
	if err := s.D.Set("internal_occm_demand_signal_delivery_collection", resources); err != nil {
		return err
	}

	return nil
}
