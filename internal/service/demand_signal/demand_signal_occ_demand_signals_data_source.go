// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package demand_signal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DemandSignalOccDemandSignalsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDemandSignalOccDemandSignals,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_demand_signal_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DemandSignalOccDemandSignalResource()),
						},
					},
				},
			},
		},
	}
}

func readDemandSignalOccDemandSignals(d *schema.ResourceData, m interface{}) error {
	sync := &DemandSignalOccDemandSignalsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccDemandSignalClient()

	return tfresource.ReadResource(sync)
}

type DemandSignalOccDemandSignalsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_demand_signal.OccDemandSignalClient
	Res    *oci_demand_signal.ListOccDemandSignalsResponse
}

func (s *DemandSignalOccDemandSignalsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DemandSignalOccDemandSignalsDataSourceCrud) Get() error {
	request := oci_demand_signal.ListOccDemandSignalsRequest{}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_demand_signal.OccDemandSignalLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "demand_signal")

	response, err := s.Client.ListOccDemandSignals(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccDemandSignals(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DemandSignalOccDemandSignalsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DemandSignalOccDemandSignalsDataSource-", DemandSignalOccDemandSignalsDataSource(), s.D))
	resources := []map[string]interface{}{}
	occDemandSignal := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccDemandSignalSummaryToMap(item))
	}
	occDemandSignal["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DemandSignalOccDemandSignalsDataSource().Schema["occ_demand_signal_collection"].Elem.(*schema.Resource).Schema)
		occDemandSignal["items"] = items
	}

	resources = append(resources, occDemandSignal)
	if err := s.D.Set("occ_demand_signal_collection", resources); err != nil {
		return err
	}

	return nil
}
