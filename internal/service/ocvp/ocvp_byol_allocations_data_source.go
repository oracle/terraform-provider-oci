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

func OcvpByolAllocationsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readOcvpByolAllocationsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"available_units_greater_than_or_equal_to": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"byol_allocation_id": {
				Type:     schema.TypeString,
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
			"byol_allocation_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OcvpByolAllocationResource()),
						},
					},
				},
			},
		},
	}
}

func readOcvpByolAllocationsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OcvpByolAllocationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ByolAllocationClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OcvpByolAllocationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.ByolAllocationClient
	Res    *oci_ocvp.ListByolAllocationsResponse
}

func (s *OcvpByolAllocationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpByolAllocationsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ocvp.ListByolAllocationsRequest{}

	if availableUnitsGreaterThanOrEqualTo, ok := s.D.GetOkExists("available_units_greater_than_or_equal_to"); ok {
		tmp := float32(availableUnitsGreaterThanOrEqualTo.(float64))
		request.AvailableUnitsGreaterThanOrEqualTo = &tmp
	}

	if byolAllocationId, ok := s.D.GetOkExists("id"); ok {
		tmp := byolAllocationId.(string)
		request.ByolAllocationId = &tmp
	}

	if byolId, ok := s.D.GetOkExists("byol_id"); ok {
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
		request.SoftwareType = oci_ocvp.ByolAllocationSoftwareTypeEnum(softwareType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ocvp.ByolAllocationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListByolAllocations(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListByolAllocations(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OcvpByolAllocationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpByolAllocationsDataSource-", OcvpByolAllocationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	byolAllocation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ByolAllocationSummaryToMap(item))
	}
	byolAllocation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OcvpByolAllocationsDataSource().Schema["byol_allocation_collection"].Elem.(*schema.Resource).Schema)
		byolAllocation["items"] = items
	}

	resources = append(resources, byolAllocation)
	if err := s.D.Set("byol_allocation_collection", resources); err != nil {
		return err
	}

	return nil
}
