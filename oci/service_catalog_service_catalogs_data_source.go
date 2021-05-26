// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v41/servicecatalog"
)

func init() {
	RegisterDatasource("oci_service_catalog_service_catalogs", ServiceCatalogServiceCatalogsDataSource())
}

func ServiceCatalogServiceCatalogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceCatalogServiceCatalogs,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_catalog_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_catalog_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(ServiceCatalogServiceCatalogResource()),
						},
					},
				},
			},
		},
	}
}

func readServiceCatalogServiceCatalogs(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceCatalogClient()

	return ReadResource(sync)
}

type ServiceCatalogServiceCatalogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.ListServiceCatalogsResponse
}

func (s *ServiceCatalogServiceCatalogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogServiceCatalogsDataSourceCrud) Get() error {
	request := oci_service_catalog.ListServiceCatalogsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if serviceCatalogId, ok := s.D.GetOkExists("id"); ok {
		tmp := serviceCatalogId.(string)
		request.ServiceCatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "service_catalog")

	response, err := s.Client.ListServiceCatalogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServiceCatalogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceCatalogServiceCatalogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ServiceCatalogServiceCatalogsDataSource-", ServiceCatalogServiceCatalogsDataSource(), s.D))
	resources := []map[string]interface{}{}
	serviceCatalog := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ServiceCatalogSummaryToMap(item))
	}
	serviceCatalog["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, ServiceCatalogServiceCatalogsDataSource().Schema["service_catalog_collection"].Elem.(*schema.Resource).Schema)
		serviceCatalog["items"] = items
	}

	resources = append(resources, serviceCatalog)
	if err := s.D.Set("service_catalog_collection", resources); err != nil {
		return err
	}

	return nil
}
