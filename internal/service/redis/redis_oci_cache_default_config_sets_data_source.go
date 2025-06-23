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

func RedisOciCacheDefaultConfigSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRedisOciCacheDefaultConfigSets,
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
			"software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_cache_default_config_set_collection": {
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

									// Optional

									// Computed
									"default_configuration_details": {
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
															"allowed_values": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"config_key": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"default_config_value": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"description": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_modifiable": {
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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

func readRedisOciCacheDefaultConfigSets(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheDefaultConfigSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheDefaultConfigSetClient()

	return tfresource.ReadResource(sync)
}

type RedisOciCacheDefaultConfigSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheDefaultConfigSetClient
	Res    *oci_redis.ListOciCacheDefaultConfigSetsResponse
}

func (s *RedisOciCacheDefaultConfigSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheDefaultConfigSetsDataSourceCrud) Get() error {
	request := oci_redis.ListOciCacheDefaultConfigSetsRequest{}

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

	if softwareVersion, ok := s.D.GetOkExists("software_version"); ok {
		request.SoftwareVersion = oci_redis.OciCacheConfigSetSoftwareVersionEnum(softwareVersion.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_redis.OciCacheDefaultConfigSetLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.ListOciCacheDefaultConfigSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOciCacheDefaultConfigSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RedisOciCacheDefaultConfigSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RedisOciCacheDefaultConfigSetsDataSource-", RedisOciCacheDefaultConfigSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	ociCacheDefaultConfigSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OciCacheDefaultConfigSetSummaryToMap(item))
	}
	ociCacheDefaultConfigSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RedisOciCacheDefaultConfigSetsDataSource().Schema["oci_cache_default_config_set_collection"].Elem.(*schema.Resource).Schema)
		ociCacheDefaultConfigSet["items"] = items
	}

	resources = append(resources, ociCacheDefaultConfigSet)
	if err := s.D.Set("oci_cache_default_config_set_collection", resources); err != nil {
		return err
	}

	return nil
}

func DefaultConfigurationDetailsToMap(obj *oci_redis.DefaultConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DefaultConfigurationInfoToMap(item))
	}
	result["items"] = items

	return result
}

func DefaultConfigurationInfoToMap(obj oci_redis.DefaultConfigurationInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowedValues != nil {
		result["allowed_values"] = string(*obj.AllowedValues)
	}

	if obj.ConfigKey != nil {
		result["config_key"] = string(*obj.ConfigKey)
	}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.DefaultConfigValue != nil {
		result["default_config_value"] = string(*obj.DefaultConfigValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsModifiable != nil {
		result["is_modifiable"] = bool(*obj.IsModifiable)
	}

	return result
}

func OciCacheDefaultConfigSetSummaryToMap(obj oci_redis.OciCacheDefaultConfigSetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["software_version"] = string(obj.SoftwareVersion)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
