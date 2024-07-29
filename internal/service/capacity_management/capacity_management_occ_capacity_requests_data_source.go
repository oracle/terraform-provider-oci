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

func CapacityManagementOccCapacityRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccCapacityRequests,
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
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_availability_catalog_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_capacity_request_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CapacityManagementOccCapacityRequestResource()),
						},
					},
				},
			},
		},
	}
}

func readCapacityManagementOccCapacityRequests(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCapacityRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccCapacityRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.ListOccCapacityRequestsResponse
}

func (s *CapacityManagementOccCapacityRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccCapacityRequestsDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccCapacityRequestsRequest{}

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
		request.Namespace = oci_capacity_management.ListOccCapacityRequestsNamespaceEnum(namespace.(string))
	}

	if occAvailabilityCatalogId, ok := s.D.GetOkExists("occ_availability_catalog_id"); ok {
		tmp := occAvailabilityCatalogId.(string)
		request.OccAvailabilityCatalogId = &tmp
	}

	if requestType, ok := s.D.GetOkExists("request_type"); ok {
		request.RequestType = oci_capacity_management.OccCapacityRequestRequestTypeEnum(requestType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccCapacityRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccCapacityRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccCapacityRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccCapacityRequestsDataSource-", CapacityManagementOccCapacityRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	occCapacityRequest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccCapacityRequestSummaryToMap(item))
	}
	occCapacityRequest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccCapacityRequestsDataSource().Schema["occ_capacity_request_collection"].Elem.(*schema.Resource).Schema)
		occCapacityRequest["items"] = items
	}

	resources = append(resources, occCapacityRequest)
	if err := s.D.Set("occ_capacity_request_collection", resources); err != nil {
		return err
	}

	return nil
}
