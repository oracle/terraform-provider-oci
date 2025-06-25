// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccmDemandSignalItemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccmDemandSignalItems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"demand_signal_namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occm_demand_signal_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occm_demand_signal_item_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CapacityManagementOccmDemandSignalItemResource()),
						},
					},
				},
			},
		},
	}
}

func readCapacityManagementOccmDemandSignalItems(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalItemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccmDemandSignalItemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.DemandSignalClient
	Res    *oci_capacity_management.ListOccmDemandSignalItemsResponse
}

func (s *CapacityManagementOccmDemandSignalItemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccmDemandSignalItemsDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccmDemandSignalItemsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if demandSignalNamespace, ok := s.D.GetOkExists("demand_signal_namespace"); ok {
		request.DemandSignalNamespace = oci_capacity_management.ListOccmDemandSignalItemsDemandSignalNamespaceEnum(demandSignalNamespace.(string))
	}

	if occmDemandSignalId, ok := s.D.GetOkExists("occm_demand_signal_id"); ok {
		tmp := occmDemandSignalId.(string)
		request.OccmDemandSignalId = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccmDemandSignalItems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccmDemandSignalItems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccmDemandSignalItemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccmDemandSignalItemsDataSource-", CapacityManagementOccmDemandSignalItemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	occmDemandSignalItem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		mappedItem := OccmDemandSignalItemSummaryToMap(item)
		if mappedItem != nil {
			items = append(items, mappedItem)
		}
	}
	occmDemandSignalItem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		log.Printf("[DEBUG] Applying filters: %+v", f)
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccmDemandSignalItemsDataSource().Schema["occm_demand_signal_item_collection"].Elem.(*schema.Resource).Schema)
		occmDemandSignalItem["items"] = items
	}

	resources = append(resources, occmDemandSignalItem)
	if err := s.D.Set("occm_demand_signal_item_collection", resources); err != nil {
		return err
	}

	return nil
}
