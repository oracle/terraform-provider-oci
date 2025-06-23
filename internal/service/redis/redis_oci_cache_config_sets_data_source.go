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

func RedisOciCacheConfigSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRedisOciCacheConfigSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"oci_cache_config_set_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(RedisOciCacheConfigSetResource()),
						},
					},
				},
			},
		},
	}
}

func readRedisOciCacheConfigSets(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheConfigSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheConfigSetClient()

	return tfresource.ReadResource(sync)
}

type RedisOciCacheConfigSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheConfigSetClient
	Res    *oci_redis.ListOciCacheConfigSetsResponse
}

func (s *RedisOciCacheConfigSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheConfigSetsDataSourceCrud) Get() error {
	request := oci_redis.ListOciCacheConfigSetsRequest{}

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
		request.LifecycleState = oci_redis.OciCacheConfigSetLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.ListOciCacheConfigSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOciCacheConfigSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RedisOciCacheConfigSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RedisOciCacheConfigSetsDataSource-", RedisOciCacheConfigSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	ociCacheConfigSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OciCacheConfigSetSummaryToMap(item))
	}
	ociCacheConfigSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RedisOciCacheConfigSetsDataSource().Schema["oci_cache_config_set_collection"].Elem.(*schema.Resource).Schema)
		ociCacheConfigSet["items"] = items
	}

	resources = append(resources, ociCacheConfigSet)
	if err := s.D.Set("oci_cache_config_set_collection", resources); err != nil {
		return err
	}

	return nil
}
