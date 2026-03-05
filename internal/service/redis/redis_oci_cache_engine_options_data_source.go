// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisOciCacheEngineOptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRedisOciCacheEngineOptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"oci_cache_engine_options_collection": {
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
									"engine_versions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version": {
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
				},
			},
		},
	}
}

func readRedisOciCacheEngineOptions(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheEngineOptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheEngineOptionsClient()

	return tfresource.ReadResource(sync)
}

type RedisOciCacheEngineOptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheEngineOptionsClient
	Res    *oci_redis.ListOciCacheEngineOptionsResponse
}

func (s *RedisOciCacheEngineOptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheEngineOptionsDataSourceCrud) Get() error {
	request := oci_redis.ListOciCacheEngineOptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.ListOciCacheEngineOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOciCacheEngineOptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RedisOciCacheEngineOptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RedisOciCacheEngineOptionsDataSource-", RedisOciCacheEngineOptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	ociCacheEngineOption := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OciCacheEngineOptionSummaryToMap(item))
	}
	ociCacheEngineOption["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RedisOciCacheEngineOptionsDataSource().Schema["oci_cache_engine_options_collection"].Elem.(*schema.Resource).Schema)
		ociCacheEngineOption["items"] = items
	}

	resources = append(resources, ociCacheEngineOption)
	if err := s.D.Set("oci_cache_engine_options_collection", resources); err != nil {
		return err
	}

	return nil
}

func EngineVersionToMap(obj oci_redis.EngineVersion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["version"] = string(obj.Version)

	return result
}

func OciCacheEngineOptionSummaryToMap(obj oci_redis.OciCacheEngineOptionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	engineVersions := []interface{}{}
	for _, item := range obj.EngineVersions {
		engineVersions = append(engineVersions, EngineVersionToMap(item))
	}
	result["engine_versions"] = engineVersions

	return result
}
