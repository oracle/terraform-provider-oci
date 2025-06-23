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

func RedisOciCacheConfigSetlistAssociatedOciCacheClusterResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisOciCacheConfigSetlistAssociatedOciCacheCluster,
		Read:     readRedisOciCacheConfigSetlistAssociatedOciCacheCluster,
		Delete:   deleteRedisOciCacheConfigSetlistAssociatedOciCacheCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"oci_cache_config_set_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createRedisOciCacheConfigSetlistAssociatedOciCacheCluster(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheConfigSetClient()

	return tfresource.CreateResource(d, sync)
}

func readRedisOciCacheConfigSetlistAssociatedOciCacheCluster(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteRedisOciCacheConfigSetlistAssociatedOciCacheCluster(d *schema.ResourceData, m interface{}) error {
	return nil
}

type RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.OciCacheConfigSetClient
	Res                    *oci_redis.AssociatedOciCacheClusterCollection
	DisableNotFoundRetries bool
}

func (s *RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("RedisOciCacheConfigSetlistAssociatedOciCacheClusterResource-", RedisOciCacheConfigSetlistAssociatedOciCacheClusterResource(), s.D)
}

func (s *RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceCrud) Create() error {
	request := oci_redis.ListAssociatedOciCacheClustersRequest{}

	if ociCacheConfigSetId, ok := s.D.GetOkExists("oci_cache_config_set_id"); ok {
		tmp := ociCacheConfigSetId.(string)
		request.OciCacheConfigSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.ListAssociatedOciCacheClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AssociatedOciCacheClusterCollection
	return nil
}

func (s *RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceCrud) SetData() error {
	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssociatedOciCacheClusterSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func AssociatedOciCacheClusterSummaryToMap(obj oci_redis.AssociatedOciCacheClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
