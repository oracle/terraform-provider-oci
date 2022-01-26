// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v56/datacatalog"
)

func DatacatalogCatalogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatacatalogCatalogs,
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
			"catalogs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatacatalogCatalogResource()),
			},
		},
	}
}

func readDatacatalogCatalogs(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

type DatacatalogCatalogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.ListCatalogsResponse
}

func (s *DatacatalogCatalogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogCatalogsDataSourceCrud) Get() error {
	request := oci_datacatalog.ListCatalogsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datacatalog.ListCatalogsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacatalog")

	response, err := s.Client.ListCatalogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCatalogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatacatalogCatalogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatacatalogCatalogsDataSource-", DatacatalogCatalogsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		catalog := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		catalog["attached_catalog_private_endpoints"] = r.AttachedCatalogPrivateEndpoints

		if r.DefinedTags != nil {
			catalog["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			catalog["display_name"] = *r.DisplayName
		}

		catalog["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			catalog["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			catalog["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.NumberOfObjects != nil {
			catalog["number_of_objects"] = *r.NumberOfObjects
		}

		catalog["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			catalog["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			catalog["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, catalog)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatacatalogCatalogsDataSource().Schema["catalogs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("catalogs", resources); err != nil {
		return err
	}

	return nil
}
