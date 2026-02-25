// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuReadinessChecksDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readFleetSoftwareUpdateFsuReadinessChecksWithContext,
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
			"resource_id": {
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
			"fsu_readiness_check_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetSoftwareUpdateFsuReadinessCheckResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetSoftwareUpdateFsuReadinessChecksWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetSoftwareUpdateFsuReadinessChecksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetSoftwareUpdateFsuReadinessChecksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res    *oci_fleet_software_update.ListFsuReadinessChecksResponse
}

func (s *FleetSoftwareUpdateFsuReadinessChecksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetSoftwareUpdateFsuReadinessChecksDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_software_update.ListFsuReadinessChecksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_software_update.FsuReadinessCheckLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_fleet_software_update.FsuReadinessCheckTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_software_update")

	response, err := s.Client.ListFsuReadinessChecks(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFsuReadinessChecks(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetSoftwareUpdateFsuReadinessChecksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetSoftwareUpdateFsuReadinessChecksDataSource-", FleetSoftwareUpdateFsuReadinessChecksDataSource(), s.D))
	resources := []map[string]interface{}{}
	fsuReadinessCheck := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FsuReadinessCheckSummaryToMap(item))
	}
	fsuReadinessCheck["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetSoftwareUpdateFsuReadinessChecksDataSource().Schema["fsu_readiness_check_collection"].Elem.(*schema.Resource).Schema)
		fsuReadinessCheck["items"] = items
	}

	resources = append(resources, fsuReadinessCheck)
	if err := s.D.Set("fsu_readiness_check_collection", resources); err != nil {
		return err
	}

	return nil
}
