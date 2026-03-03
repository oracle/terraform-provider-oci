// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpByolsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readOcvpByolsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"available_units_greater_than_or_equal_to": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"byol_id": {
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
			"software_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"byol_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OcvpByolResource()),
						},
					},
				},
			},
		},
	}
}

func readOcvpByolsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OcvpByolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ByolClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OcvpByolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.ByolClient
	Res    *oci_ocvp.ListByolsResponse
}

func (s *OcvpByolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpByolsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ocvp.ListByolsRequest{}

	if availableUnitsGreaterThanOrEqualTo, ok := s.D.GetOkExists("available_units_greater_than_or_equal_to"); ok {
		tmp := float32(availableUnitsGreaterThanOrEqualTo.(float64))
		request.AvailableUnitsGreaterThanOrEqualTo = &tmp
	}

	if byolId, ok := s.D.GetOkExists("id"); ok {
		tmp := byolId.(string)
		request.ByolId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if softwareType, ok := s.D.GetOkExists("software_type"); ok {
		request.SoftwareType = oci_ocvp.ByolSoftwareTypeEnum(softwareType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ocvp.ByolLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListByols(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListByols(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OcvpByolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpByolsDataSource-", OcvpByolsDataSource(), s.D))
	resources := []map[string]interface{}{}
	byol := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ByolSummaryToMap(item))
	}
	byol["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OcvpByolsDataSource().Schema["byol_collection"].Elem.(*schema.Resource).Schema)
		byol["items"] = items
	}

	resources = append(resources, byol)
	if err := s.D.Set("byol_collection", resources); err != nil {
		return err
	}

	return nil
}
