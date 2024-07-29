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

func CapacityManagementInternalOccAvailabilityCatalogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementInternalOccAvailabilityCatalogs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"catalog_state": {
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_customer_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"occ_availability_catalog_collection": {
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
									"catalog_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"metadata_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"format_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"occ_customer_group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readCapacityManagementInternalOccAvailabilityCatalogs(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccAvailabilityCatalogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementInternalOccAvailabilityCatalogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.ListOccAvailabilityCatalogsInternalResponse
}

func (s *CapacityManagementInternalOccAvailabilityCatalogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementInternalOccAvailabilityCatalogsDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccAvailabilityCatalogsInternalRequest{}

	if catalogState, ok := s.D.GetOkExists("catalog_state"); ok {
		request.CatalogState = oci_capacity_management.OccAvailabilityCatalogCatalogStateEnum(catalogState.(string))
	}

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

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		request.Namespace = oci_capacity_management.ListOccAvailabilityCatalogsInternalNamespaceEnum(namespace.(string))
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccAvailabilityCatalogsInternal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccAvailabilityCatalogsInternal(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementInternalOccAvailabilityCatalogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementInternalOccAvailabilityCatalogsDataSource-", CapacityManagementInternalOccAvailabilityCatalogsDataSource(), s.D))
	resources := []map[string]interface{}{}
	internalOccAvailabilityCatalog := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccAvailabilityCatalogSummaryToMap(item))
	}
	internalOccAvailabilityCatalog["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementInternalOccAvailabilityCatalogsDataSource().Schema["occ_availability_catalog_collection"].Elem.(*schema.Resource).Schema)
		internalOccAvailabilityCatalog["items"] = items
	}

	resources = append(resources, internalOccAvailabilityCatalog)
	if err := s.D.Set("occ_availability_catalog_collection", resources); err != nil {
		return err
	}

	return nil
}
