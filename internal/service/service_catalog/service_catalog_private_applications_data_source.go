// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v58/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ServiceCatalogPrivateApplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceCatalogPrivateApplications,
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
			"private_application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_application_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ServiceCatalogPrivateApplicationResource()),
						},
					},
				},
			},
		},
	}
}

func readServiceCatalogPrivateApplications(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.ReadResource(sync)
}

type ServiceCatalogPrivateApplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.ListPrivateApplicationsResponse
}

func (s *ServiceCatalogPrivateApplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogPrivateApplicationsDataSourceCrud) Get() error {
	request := oci_service_catalog.ListPrivateApplicationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if privateApplicationId, ok := s.D.GetOkExists("id"); ok {
		tmp := privateApplicationId.(string)
		request.PrivateApplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_catalog")

	response, err := s.Client.ListPrivateApplications(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivateApplications(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceCatalogPrivateApplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceCatalogPrivateApplicationsDataSource-", ServiceCatalogPrivateApplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	privateApplication := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PrivateApplicationSummaryToMap(item))
	}
	privateApplication["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ServiceCatalogPrivateApplicationsDataSource().Schema["private_application_collection"].Elem.(*schema.Resource).Schema)
		privateApplication["items"] = items
	}

	resources = append(resources, privateApplication)
	if err := s.D.Set("private_application_collection", resources); err != nil {
		return err
	}

	return nil
}
