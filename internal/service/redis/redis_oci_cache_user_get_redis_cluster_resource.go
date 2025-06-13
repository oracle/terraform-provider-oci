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

func RedisOciCacheUserGetRedisClusterResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisOciCacheUserGetRedisCluster,
		Read:     readRedisOciCacheUserGetRedisCluster,
		Delete:   deleteRedisOciCacheUserGetRedisCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"oci_cache_user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"oci_cache_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oci_cache_cluster_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createRedisOciCacheUserGetRedisCluster(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheUserGetRedisClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheUserClient()

	return tfresource.CreateResource(d, sync)
}

func readRedisOciCacheUserGetRedisCluster(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteRedisOciCacheUserGetRedisCluster(d *schema.ResourceData, m interface{}) error {
	return nil
}

type RedisOciCacheUserGetRedisClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.OciCacheUserClient
	Res                    []oci_redis.AttachedOciCacheCluster
	DisableNotFoundRetries bool
}

func (s *RedisOciCacheUserGetRedisClusterResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("RedisOciCacheUserGetRedisClusterResource-", RedisOciCacheUserGetRedisClusterResource(), s.D)
}

func (s *RedisOciCacheUserGetRedisClusterResourceCrud) Create() error {
	request := oci_redis.ListAttachedRedisClustersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if ociCacheUserId, ok := s.D.GetOkExists("oci_cache_user_id"); ok {
		tmp := ociCacheUserId.(string)
		request.OciCacheUserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.ListAttachedRedisClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = response.Items
	return nil
}

func (s *RedisOciCacheUserGetRedisClusterResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	var clusters []map[string]interface{}
	for _, cluster := range s.Res {
		clusters = append(clusters, map[string]interface{}{
			"oci_cache_cluster_id": cluster.OciCacheClusterId,
		})
	}
	if err := s.D.Set("oci_cache_clusters", clusters); err != nil {
		return err
	}

	s.D.SetId(s.ID())
	return nil
}
