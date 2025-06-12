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

func CapacityManagementOccmDemandSignalsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccmDemandSignals,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occm_demand_signal_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CapacityManagementOccmDemandSignalResource()),
						},
					},
				},
			},
		},
	}
}

func readCapacityManagementOccmDemandSignals(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccmDemandSignalsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.DemandSignalClient
	Res    *oci_capacity_management.ListOccmDemandSignalsResponse
}

func (s *CapacityManagementOccmDemandSignalsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccmDemandSignalsDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccmDemandSignalsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		tmp := lifecycleDetails.(string)
		request.LifecycleDetails = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccmDemandSignals(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccmDemandSignals(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccmDemandSignalsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccmDemandSignalsDataSource-", CapacityManagementOccmDemandSignalsDataSource(), s.D))
	resources := []map[string]interface{}{}
	occmDemandSignal := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccmDemandSignalSummaryToMap(item))
	}
	occmDemandSignal["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccmDemandSignalsDataSource().Schema["occm_demand_signal_collection"].Elem.(*schema.Resource).Schema)
		occmDemandSignal["items"] = items
	}

	resources = append(resources, occmDemandSignal)
	if err := s.D.Set("occm_demand_signal_collection", resources); err != nil {
		return err
	}

	return nil
}
