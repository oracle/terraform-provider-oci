// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v56/datacatalog"
)

func DatacatalogCatalogTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatacatalogCatalogTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"catalog_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_type_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fields": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_approved": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_internal": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_tag": {
				Type:     schema.TypeString,
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
			"type_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"catalog_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
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

func readDatacatalogCatalogTypes(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

type DatacatalogCatalogTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.ListTypesResponse
}

func (s *DatacatalogCatalogTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogCatalogTypesDataSourceCrud) Get() error {
	request := oci_datacatalog.ListTypesRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if externalTypeName, ok := s.D.GetOkExists("external_type_name"); ok {
		tmp := externalTypeName.(string)
		request.ExternalTypeName = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		set := fields.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_datacatalog.ListTypesFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacatalog.ListTypesFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if isApproved, ok := s.D.GetOkExists("is_approved"); ok {
		tmp := isApproved.(string)
		request.IsApproved = &tmp
	}

	if isInternal, ok := s.D.GetOkExists("is_internal"); ok {
		tmp := isInternal.(string)
		request.IsInternal = &tmp
	}

	if isTag, ok := s.D.GetOkExists("is_tag"); ok {
		tmp := isTag.(string)
		request.IsTag = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datacatalog.ListTypesLifecycleStateEnum(state.(string))
	}

	if typeCategory, ok := s.D.GetOkExists("type_category"); ok {
		tmp := typeCategory.(string)
		request.TypeCategory = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacatalog")

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

func (s *DatacatalogCatalogTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatacatalogCatalogTypesDataSource-", DatacatalogCatalogTypesDataSource(), s.D))

	resources := []map[string]interface{}{}
	catalogType := map[string]interface{}{}

	items := []interface{}{}

	for _, item := range s.Res.Items {
		items = append(items, TypeSummaryToMap(item))
	}
	catalogType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatacatalogCatalogTypesDataSource().Schema["type_collection"].Elem.(*schema.Resource).Schema)
		catalogType["items"] = items
	}

	if s.Res.Count != nil {
		catalogType["count"] = *s.Res.Count
	}

	resources = append(resources, catalogType)
	if err := s.D.Set("type_collection", resources); err != nil {
		return err
	}

	return nil
}

func TypeSummaryToMap(obj oci_datacatalog.TypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CatalogId != nil {
		result["catalog_id"] = string(*obj.CatalogId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TypeCategory != nil {
		result["type_category"] = string(*obj.TypeCategory)
	}

	if obj.Uri != nil {
		result["uri"] = string(*obj.Uri)
	}

	return result
}
