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

func FleetSoftwareUpdateFsuCollectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetSoftwareUpdateFsuCollections,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsu_collection_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetSoftwareUpdateFsuCollectionResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetSoftwareUpdateFsuCollections(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCollectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.ReadResource(sync)
}

type FleetSoftwareUpdateFsuCollectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res    *oci_fleet_software_update.ListFsuCollectionsResponse
}

func (s *FleetSoftwareUpdateFsuCollectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetSoftwareUpdateFsuCollectionsDataSourceCrud) Get() error {
	request := oci_fleet_software_update.ListFsuCollectionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_software_update.ListFsuCollectionsLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_fleet_software_update.ListFsuCollectionsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_software_update")

	response, err := s.Client.ListFsuCollections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFsuCollections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetSoftwareUpdateFsuCollectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetSoftwareUpdateFsuCollectionsDataSource-", FleetSoftwareUpdateFsuCollectionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fsuCollection := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FsuCollectionSummaryToMap(item))
	}
	fsuCollection["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetSoftwareUpdateFsuCollectionsDataSource().Schema["fsu_collection_summary_collection"].Elem.(*schema.Resource).Schema)
		fsuCollection["items"] = items
	}

	resources = append(resources, fsuCollection)
	if err := s.D.Set("fsu_collection_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
