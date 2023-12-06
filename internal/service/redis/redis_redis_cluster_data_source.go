// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisRedisClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["redis_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(RedisRedisClusterResource(), fieldMap, readSingularRedisRedisCluster)
}

func readSingularRedisRedisCluster(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.ReadResource(sync)
}

type RedisRedisClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.RedisClusterClient
	Res    *oci_redis.GetRedisClusterResponse
}

func (s *RedisRedisClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisRedisClusterDataSourceCrud) Get() error {
	request := oci_redis.GetRedisClusterRequest{}

	if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
		tmp := redisClusterId.(string)
		request.RedisClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.GetRedisCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RedisRedisClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NodeCollection != nil {
		s.D.Set("node_collection", []interface{}{NodeCollectionToMap(s.Res.NodeCollection)})
	} else {
		s.D.Set("node_collection", nil)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	if s.Res.NodeMemoryInGBs != nil {
		s.D.Set("node_memory_in_gbs", *s.Res.NodeMemoryInGBs)
	}

	if s.Res.PrimaryEndpointIpAddress != nil {
		s.D.Set("primary_endpoint_ip_address", *s.Res.PrimaryEndpointIpAddress)
	}

	if s.Res.PrimaryFqdn != nil {
		s.D.Set("primary_fqdn", *s.Res.PrimaryFqdn)
	}

	if s.Res.ReplicasEndpointIpAddress != nil {
		s.D.Set("replicas_endpoint_ip_address", *s.Res.ReplicasEndpointIpAddress)
	}

	if s.Res.ReplicasFqdn != nil {
		s.D.Set("replicas_fqdn", *s.Res.ReplicasFqdn)
	}

	s.D.Set("software_version", s.Res.SoftwareVersion)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

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
