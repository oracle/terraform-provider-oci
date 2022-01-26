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

func DatacatalogCatalogPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatacatalogCatalogPrivateEndpoints,
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
			"catalog_private_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatacatalogCatalogPrivateEndpointResource()),
			},
		},
	}
}

func readDatacatalogCatalogPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

type DatacatalogCatalogPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.ListCatalogPrivateEndpointsResponse
}

func (s *DatacatalogCatalogPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogCatalogPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_datacatalog.ListCatalogPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datacatalog.ListCatalogPrivateEndpointsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacatalog")

	response, err := s.Client.ListCatalogPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCatalogPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatacatalogCatalogPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatacatalogCatalogPrivateEndpointsDataSource-", DatacatalogCatalogPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		catalogPrivateEndpoint := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		catalogPrivateEndpoint["attached_catalogs"] = r.AttachedCatalogs

		if r.DefinedTags != nil {
			catalogPrivateEndpoint["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			catalogPrivateEndpoint["display_name"] = *r.DisplayName
		}

		catalogPrivateEndpoint["dns_zones"] = r.DnsZones

		catalogPrivateEndpoint["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			catalogPrivateEndpoint["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			catalogPrivateEndpoint["lifecycle_details"] = *r.LifecycleDetails
		}

		catalogPrivateEndpoint["state"] = r.LifecycleState

		if r.SubnetId != nil {
			catalogPrivateEndpoint["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			catalogPrivateEndpoint["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			catalogPrivateEndpoint["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, catalogPrivateEndpoint)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatacatalogCatalogPrivateEndpointsDataSource().Schema["catalog_private_endpoints"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("catalog_private_endpoints", resources); err != nil {
		return err
	}

	return nil
}
