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

func RedisRedisClusterGetOciCacheUserResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisRedisClusterGetOciCacheUser,
		Read:     readRedisRedisClusterGetOciCacheUser,
		Delete:   deleteRedisRedisClusterGetOciCacheUser,
		Schema: map[string]*schema.Schema{
			// Required
			"redis_cluster_id": {
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
			"oci_cache_users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oci_cache_user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createRedisRedisClusterGetOciCacheUser(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterGetOciCacheUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.CreateResource(d, sync)
}

func readRedisRedisClusterGetOciCacheUser(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteRedisRedisClusterGetOciCacheUser(d *schema.ResourceData, m interface{}) error {
	return nil
}

type RedisRedisClusterGetOciCacheUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.RedisClusterClient
	Res                    []oci_redis.AttachedOciCacheUser
	DisableNotFoundRetries bool
}

func (s *RedisRedisClusterGetOciCacheUserResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("RedisRedisClusterGetOciCacheUserResource-", RedisRedisClusterGetOciCacheUserResource(), s.D)
}

func (s *RedisRedisClusterGetOciCacheUserResourceCrud) Create() error {
	request := oci_redis.ListAttachedOciCacheUsersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
		tmp := redisClusterId.(string)
		request.RedisClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.ListAttachedOciCacheUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = response.Items
	return nil
}

func (s *RedisRedisClusterGetOciCacheUserResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	var users []map[string]interface{}
	for _, user := range s.Res {
		users = append(users, map[string]interface{}{
			"oci_cache_user_id": user.OciCacheUserId,
		})
	}
	if err := s.D.Set("oci_cache_users", users); err != nil {
		return err
	}

	s.D.SetId(s.ID())
	return nil
}
