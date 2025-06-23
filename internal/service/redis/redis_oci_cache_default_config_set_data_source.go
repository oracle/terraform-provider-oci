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

func RedisOciCacheDefaultConfigSetDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularRedisOciCacheDefaultConfigSet,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"oci_cache_default_config_set_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularRedisOciCacheDefaultConfigSet(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheDefaultConfigSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheDefaultConfigSetClient()

	return tfresource.ReadResource(sync)
}

type RedisOciCacheDefaultConfigSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheDefaultConfigSetClient
	Res    *oci_redis.GetOciCacheDefaultConfigSetResponse
}

func (s *RedisOciCacheDefaultConfigSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheDefaultConfigSetDataSourceCrud) Get() error {
	request := oci_redis.GetOciCacheDefaultConfigSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if ociCacheDefaultConfigSetId, ok := s.D.GetOkExists("oci_cache_default_config_set_id"); ok {
		tmp := ociCacheDefaultConfigSetId.(string)
		request.OciCacheDefaultConfigSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.GetOciCacheDefaultConfigSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RedisOciCacheDefaultConfigSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefaultConfigurationDetails != nil {
		s.D.Set("default_configuration_details", []interface{}{DefaultConfigurationDetailsToMap(s.Res.DefaultConfigurationDetails)})
	} else {
		s.D.Set("default_configuration_details", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("software_version", s.Res.SoftwareVersion)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
