// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v65/dataconnectivity"

	"terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/tfresource"
)

func DataConnectivityRegistriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataConnectivityRegistries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_deep_lookup": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"registry_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataConnectivityRegistryResource()),
						},
					},
				},
			},
		},
	}
}

func readDataConnectivityRegistries(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.ListRegistriesResponse
}

func (s *DataConnectivityRegistriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistriesDataSourceCrud) Get() error {
	request := oci_data_connectivity.ListRegistriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isDeepLookup, ok := s.D.GetOkExists("is_deep_lookup"); ok {
		tmp := isDeepLookup.(bool)
		request.IsDeepLookup = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	// if state, ok := s.D.GetOkExists("state"); ok {
	// 	request.LifecycleState = oci_data_connectivity.RegistryLifecycleStateEnum(state.(string))
	// }

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.ListRegistries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRegistries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataConnectivityRegistriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataConnectivityRegistriesDataSource-", DataConnectivityRegistriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	registry := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RegistrySummaryToMap(item))
	}
	registry["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataConnectivityRegistriesDataSource().Schema["registry_summary_collection"].Elem.(*schema.Resource).Schema)
		registry["items"] = items
	}

	resources = append(resources, registry)
	if err := s.D.Set("registry_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
