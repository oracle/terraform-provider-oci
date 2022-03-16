// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v62/dataconnectivity"
)

func DataConnectivityRegistryTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataConnectivityRegistryTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"registry_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"types_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									// Optional
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"connection_attributes": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"data_asset_attributes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"attribute_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_base64encoded": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_generated": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_mandatory": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_sensitive": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"valid_key_list": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
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

func readDataConnectivityRegistryTypes(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistryTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.ListTypesResponse
}

func (s *DataConnectivityRegistryTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistryTypesDataSourceCrud) Get() error {
	request := oci_data_connectivity.ListTypesRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.ListTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataConnectivityRegistryTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataConnectivityRegistryTypesDataSource-", DataConnectivityRegistryTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	registryType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataConnectivityTypeSummaryToMap(item))
	}
	registryType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataConnectivityRegistryTypesDataSource().Schema["types_summary_collection"].Elem.(*schema.Resource).Schema)
		registryType["items"] = items
	}

	resources = append(resources, registryType)
	if err := s.D.Set("types_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
