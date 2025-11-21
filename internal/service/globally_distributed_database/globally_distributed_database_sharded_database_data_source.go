// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globally_distributed_database

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GloballyDistributedDatabaseShardedDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["metadata"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["sharded_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GloballyDistributedDatabaseShardedDatabaseResource(), fieldMap, readSingularGloballyDistributedDatabaseShardedDatabase)
}

func readSingularGloballyDistributedDatabaseShardedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabaseShardedDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.ReadResource(sync)
}

type GloballyDistributedDatabaseShardedDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_globally_distributed_database.ShardedDatabaseServiceClient
	Res    *oci_globally_distributed_database.GetShardedDatabaseResponse
}

func (s *GloballyDistributedDatabaseShardedDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GloballyDistributedDatabaseShardedDatabaseDataSourceCrud) Get() error {
	request := oci_globally_distributed_database.GetShardedDatabaseRequest{}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp := metadata.(string)
		request.Metadata = &tmp
	}

	if shardedDatabaseId, ok := s.D.GetOkExists("sharded_database_id"); ok {
		tmp := shardedDatabaseId.(string)
		request.ShardedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "globally_distributed_database")

	response, err := s.Client.GetShardedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.ShardedDatabase).(type) {
	case oci_globally_distributed_database.DedicatedShardedDatabase:
		s.D.Set("db_deployment_type", "DEDICATED")

		catalogDetails := []interface{}{}
		for _, item := range v.CatalogDetails {
			catalogDetails = append(catalogDetails, DedicatedCatalogDetailsToMap(item))
		}
		s.D.Set("catalog_details", catalogDetails)

		if v.CharacterSet != nil {
			s.D.Set("character_set", *v.CharacterSet)
		}

		if v.Chunks != nil {
			s.D.Set("chunks", *v.Chunks)
		}

		if v.ClusterCertificateCommonName != nil {
			s.D.Set("cluster_certificate_common_name", *v.ClusterCertificateCommonName)
		}

		if v.ConnectionStrings != nil {
			s.D.Set("connection_strings", []interface{}{ConnectionStringToMap(v.ConnectionStrings)})
		} else {
			s.D.Set("connection_strings", nil)
		}

		if v.DbVersion != nil {
			s.D.Set("db_version", *v.DbVersion)
		}

		s.D.Set("db_workload", v.DbWorkload)

		gsms := []interface{}{}
		for _, item := range v.Gsms {
			gsms = append(gsms, GsmDetailsToMap(item))
		}
		s.D.Set("gsms", gsms)

		if v.ListenerPort != nil {
			s.D.Set("listener_port", *v.ListenerPort)
		}

		if v.ListenerPortTls != nil {
			s.D.Set("listener_port_tls", *v.ListenerPortTls)
		}

		if v.NcharacterSet != nil {
			s.D.Set("ncharacter_set", *v.NcharacterSet)
		}

		if v.OnsPortLocal != nil {
			s.D.Set("ons_port_local", *v.OnsPortLocal)
		}

		if v.OnsPortRemote != nil {
			s.D.Set("ons_port_remote", *v.OnsPortRemote)
		}

		if v.Prefix != nil {
			s.D.Set("prefix", *v.Prefix)
		}

		if v.PrivateEndpoint != nil {
			s.D.Set("private_endpoint", *v.PrivateEndpoint)
		}

		if v.ReplicationFactor != nil {
			s.D.Set("replication_factor", *v.ReplicationFactor)
		}

		s.D.Set("replication_method", v.ReplicationMethod)

		if v.ReplicationUnit != nil {
			s.D.Set("replication_unit", *v.ReplicationUnit)
		}

		shardDetails := []interface{}{}
		for _, item := range v.ShardDetails {
			shardDetails = append(shardDetails, DedicatedShardDetailsToMap(item))
		}
		s.D.Set("shard_details", shardDetails)

		s.D.Set("sharding_method", v.ShardingMethod)

		if v.TimeZone != nil {
			s.D.Set("time_zone", *v.TimeZone)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleStateDetails != nil {
			s.D.Set("lifecycle_state_details", *v.LifecycleStateDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'db_deployment_type' of unknown type %v", s.Res.ShardedDatabase)
		return nil
	}

	return nil
}
