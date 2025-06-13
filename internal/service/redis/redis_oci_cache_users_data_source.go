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

func RedisOciCacheUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRedisOciCacheUsers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_cache_user_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OciCacheUserSummaryResource()),
						},
					},
				},
			},
		},
	}
}

func OciCacheUserSummaryResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             &schema.Schema{Type: schema.TypeString},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func readRedisOciCacheUsers(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheUserClient()

	return tfresource.ReadResource(sync)
}

type RedisOciCacheUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheUserClient
	Res    *oci_redis.ListOciCacheUsersResponse
}

func (s *RedisOciCacheUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheUsersDataSourceCrud) Get() error {
	request := oci_redis.ListOciCacheUsersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_redis.OciCacheUserLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.ListOciCacheUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOciCacheUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RedisOciCacheUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RedisOciCacheUsersDataSource-", RedisOciCacheUsersDataSource(), s.D))
	resources := []map[string]interface{}{}
	ociCacheUser := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OciCacheUserSummaryToMap(item))
	}
	ociCacheUser["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RedisOciCacheUsersDataSource().Schema["oci_cache_user_collection"].Elem.(*schema.Resource).Schema)
		ociCacheUser["items"] = items
	}

	resources = append(resources, ociCacheUser)
	if err := s.D.Set("oci_cache_user_collection", resources); err != nil {
		return err
	}

	return nil
}
