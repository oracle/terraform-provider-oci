// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisRedisClusterCreateIdentityTokenResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisRedisClusterCreateIdentityToken,
		Read:     readRedisRedisClusterCreateIdentityToken,
		Delete:   deleteRedisRedisClusterCreateIdentityToken,
		Schema: map[string]*schema.Schema{
			// Required
			"public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"redis_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"redis_user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"identity_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createRedisRedisClusterCreateIdentityToken(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterCreateIdentityTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisIdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readRedisRedisClusterCreateIdentityToken(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteRedisRedisClusterCreateIdentityToken(d *schema.ResourceData, m interface{}) error {
	return nil
}

type RedisRedisClusterCreateIdentityTokenResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.RedisIdentityClient
	Res                    *oci_redis.IdentityTokenDetailsResponse
	DisableNotFoundRetries bool
}

func (s *RedisRedisClusterCreateIdentityTokenResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("RedisRedisClusterCreateIdentityTokenResource-", RedisRedisClusterCreateIdentityTokenResource(), s.D)
}

func (s *RedisRedisClusterCreateIdentityTokenResourceCrud) Create() error {
	request := oci_redis.CreateIdentityTokenRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		tmp := publicKey.(string)
		request.PublicKey = &tmp
	}

	if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
		tmp := redisClusterId.(string)
		request.RedisClusterId = &tmp
	}

	if redisUser, ok := s.D.GetOkExists("redis_user"); ok {
		tmp := redisUser.(string)
		request.RedisUser = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")
	log.Printf("[DEBUG] *** 00100 request ***: %v", request)

	response, err := s.Client.CreateIdentityToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityTokenDetailsResponse
	return nil
}

func (s *RedisRedisClusterCreateIdentityTokenResourceCrud) SetData() error {
	if s.Res.IdentityToken != nil {
		s.D.Set("identity_token", *s.Res.IdentityToken)
	}

	if s.Res.RedisUser != nil {
		s.D.Set("redis_user", *s.Res.RedisUser)
	}

	s.D.SetId(s.ID())
	return nil
}
