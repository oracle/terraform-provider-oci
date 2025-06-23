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

func RedisOciCacheConfigSetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oci_cache_config_set_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(RedisOciCacheConfigSetResource(), fieldMap, readSingularRedisOciCacheConfigSet)
}

func readSingularRedisOciCacheConfigSet(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheConfigSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheConfigSetClient()

	return tfresource.ReadResource(sync)
}

type RedisOciCacheConfigSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheConfigSetClient
	Res    *oci_redis.GetOciCacheConfigSetResponse
}

func (s *RedisOciCacheConfigSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheConfigSetDataSourceCrud) Get() error {
	request := oci_redis.GetOciCacheConfigSetRequest{}

	if ociCacheConfigSetId, ok := s.D.GetOkExists("oci_cache_config_set_id"); ok {
		tmp := ociCacheConfigSetId.(string)
		request.OciCacheConfigSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.GetOciCacheConfigSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RedisOciCacheConfigSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationDetails != nil {
		s.D.Set("configuration_details", []interface{}{ConfigurationDetailsToMap(s.Res.ConfigurationDetails)})
	} else {
		s.D.Set("configuration_details", nil)
	}

	if s.Res.DefaultConfigSetId != nil {
		s.D.Set("default_config_set_id", *s.Res.DefaultConfigSetId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("software_version", s.Res.SoftwareVersion)

	s.D.Set("state", s.Res.LifecycleState)

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
