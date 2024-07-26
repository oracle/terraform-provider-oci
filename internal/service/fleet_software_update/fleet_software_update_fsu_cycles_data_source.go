// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuCyclesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetSoftwareUpdateFsuCycles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"collection_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsu_collection_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsu_cycle_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetSoftwareUpdateFsuCycleResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetSoftwareUpdateFsuCycles(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCyclesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.ReadResource(sync)
}

type FleetSoftwareUpdateFsuCyclesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res    *oci_fleet_software_update.ListFsuCyclesResponse
}

func (s *FleetSoftwareUpdateFsuCyclesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetSoftwareUpdateFsuCyclesDataSourceCrud) Get() error {
	request := oci_fleet_software_update.ListFsuCyclesRequest{}

	if collectionType, ok := s.D.GetOkExists("collection_type"); ok {
		request.CollectionType = oci_fleet_software_update.ListFsuCyclesCollectionTypeEnum(collectionType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fsuCollectionId, ok := s.D.GetOkExists("fsu_collection_id"); ok {
		tmp := fsuCollectionId.(string)
		request.FsuCollectionId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_software_update.ListFsuCyclesLifecycleStateEnum(state.(string))
	}

	if targetVersion, ok := s.D.GetOkExists("target_version"); ok {
		tmp := targetVersion.(string)
		request.TargetVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_software_update")

	response, err := s.Client.ListFsuCycles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFsuCycles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetSoftwareUpdateFsuCyclesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetSoftwareUpdateFsuCyclesDataSource-", FleetSoftwareUpdateFsuCyclesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fsuCycle := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FsuCycleSummaryToMap(item))
	}
	fsuCycle["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetSoftwareUpdateFsuCyclesDataSource().Schema["fsu_cycle_summary_collection"].Elem.(*schema.Resource).Schema)
		fsuCycle["items"] = items
	}

	resources = append(resources, fsuCycle)
	if err := s.D.Set("fsu_cycle_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
