// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccAvailabilityCatalogOccAvailabilities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"date_expected_capacity_handover": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_availability_catalog_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workload_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_availability_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"available_quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"catalog_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"date_expected_capacity_handover": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"date_final_customer_order": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"demanded_quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"total_available_quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"workload_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCapacityManagementOccAvailabilityCatalogOccAvailabilities(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.ListOccAvailabilitiesResponse
}

func (s *CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccAvailabilitiesRequest{}

	if dateExpectedCapacityHandover, ok := s.D.GetOkExists("date_expected_capacity_handover"); ok {
		tmp := dateExpectedCapacityHandover.(string)
		request.DateExpectedCapacityHandover = &tmp
	}

	if occAvailabilityCatalogId, ok := s.D.GetOkExists("occ_availability_catalog_id"); ok {
		tmp := occAvailabilityCatalogId.(string)
		request.OccAvailabilityCatalogId = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if workloadType, ok := s.D.GetOkExists("workload_type"); ok {
		tmp := workloadType.(string)
		request.WorkloadType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccAvailabilities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccAvailabilities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSource-", CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	occAvailabilityCatalogOccAvailability := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccAvailabilitySummaryToMap(item))
	}
	occAvailabilityCatalogOccAvailability["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSource().Schema["occ_availability_collection"].Elem.(*schema.Resource).Schema)
		occAvailabilityCatalogOccAvailability["items"] = items
	}

	resources = append(resources, occAvailabilityCatalogOccAvailability)
	if err := s.D.Set("occ_availability_collection", resources); err != nil {
		return err
	}

	return nil
}
