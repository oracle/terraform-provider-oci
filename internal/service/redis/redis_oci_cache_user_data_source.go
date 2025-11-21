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

func RedisOciCacheUserDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oci_cache_user_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(RedisOciCacheUserResource(), fieldMap, readSingularRedisOciCacheUser)
}

func readSingularRedisOciCacheUser(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheUserClient()

	return tfresource.ReadResource(sync)
}

type RedisOciCacheUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheUserClient
	Res    *oci_redis.GetOciCacheUserResponse
}

func (s *RedisOciCacheUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheUserDataSourceCrud) Get() error {
	request := oci_redis.GetOciCacheUserRequest{}

	if ociCacheUserId, ok := s.D.GetOkExists("oci_cache_user_id"); ok {
		tmp := ociCacheUserId.(string)
		request.OciCacheUserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.GetOciCacheUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RedisOciCacheUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AclString != nil {
		s.D.Set("acl_string", *s.Res.AclString)
	}

	if s.Res.AuthenticationMode != nil {
		authenticationModeArray := []interface{}{}
		if authenticationModeMap := AuthenticationModeToMap(&s.Res.AuthenticationMode); authenticationModeMap != nil {
			authenticationModeArray = append(authenticationModeArray, authenticationModeMap)
		}
		s.D.Set("authentication_mode", authenticationModeArray)
	} else {
		s.D.Set("authentication_mode", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
