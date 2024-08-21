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

func RedisRedisClusterNodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRedisRedisClusterNodes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redis_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"redis_node_collection": {
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
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_endpoint_fqdn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_endpoint_ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"redis_cluster_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"shard_number": {
										Type:     schema.TypeInt,
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

func readRedisRedisClusterNodes(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterNodesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.ReadResource(sync)
}

type RedisRedisClusterNodesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.RedisClusterClient
	Res    *oci_redis.ListRedisClusterNodesResponse
}

func (s *RedisRedisClusterNodesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisRedisClusterNodesDataSourceCrud) Get() error {
	request := oci_redis.ListRedisClusterNodesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
		tmp := redisClusterId.(string)
		request.RedisClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.ListRedisClusterNodes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRedisClusterNodes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RedisRedisClusterNodesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RedisRedisClusterNodesDataSource-", RedisRedisClusterNodesDataSource(), s.D))
	resources := []map[string]interface{}{}
	redisClusterNode := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NodeSummaryToMap(item))
	}
	redisClusterNode["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RedisRedisClusterNodesDataSource().Schema["redis_node_collection"].Elem.(*schema.Resource).Schema)
		redisClusterNode["items"] = items
	}

	resources = append(resources, redisClusterNode)
	if err := s.D.Set("redis_node_collection", resources); err != nil {
		return err
	}

	return nil
}

func NodeSummaryToMap(obj oci_redis.NodeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.PrivateEndpointFqdn != nil {
		result["private_endpoint_fqdn"] = string(*obj.PrivateEndpointFqdn)
	}

	if obj.PrivateEndpointIpAddress != nil {
		result["private_endpoint_ip_address"] = string(*obj.PrivateEndpointIpAddress)
	}

	if obj.RedisClusterId != nil {
		result["redis_cluster_id"] = string(*obj.RedisClusterId)
	}

	if obj.ShardNumber != nil {
		result["shard_number"] = int(*obj.ShardNumber)
	}

	return result
}
