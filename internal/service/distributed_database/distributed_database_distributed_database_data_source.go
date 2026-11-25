// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["distributed_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["metadata_query"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Comma-separated argument names for which metadata should be retrieved.",
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DistributedDatabaseDistributedDatabaseResource(), fieldMap, readSingularDistributedDatabaseDistributedDatabaseWithContext)
}

func readSingularDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedDbServiceClient
	Res    *oci_distributed_database.GetDistributedDatabaseResponse
}

func (s *DistributedDatabaseDistributedDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedDatabaseDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedDatabaseRequest{Metadata: oci_common.String(operationalStateMetadataQuery)}

	if distributedDatabaseId, ok := s.D.GetOkExists("distributed_database_id"); ok {
		tmp := distributedDatabaseId.(string)
		request.DistributedDatabaseId = &tmp
	}

	if metadataQuery, ok := s.D.GetOkExists("metadata_query"); ok {
		tmp := metadataQuery.(string)
		request.Metadata = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.GetDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DistributedDatabaseDistributedDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutoResourceManagementConfig != nil {
		s.D.Set("auto_resource_management_config", []interface{}{AutoResourceManagementConfigurationDetailsToMap(s.Res.AutoResourceManagementConfig)})
	} else {
		s.D.Set("auto_resource_management_config", nil)
	}

	catalogDetails := []interface{}{}
	for _, item := range s.Res.CatalogDetails {
		catalogDetails = append(catalogDetails, DistributedDatabaseCatalogToMap(item))
	}
	s.D.Set("catalog_details", catalogDetails)

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	compositeDataGuardShardSpaces := []interface{}{}
	for _, item := range s.Res.CompositeDataGuardShardSpaces {
		compositeDataGuardShardSpaces = append(compositeDataGuardShardSpaces, CompositeDataGuardShardSpaceToMap(item))
	}
	s.D.Set("composite_data_guard_shard_spaces", compositeDataGuardShardSpaces)

	compositeRaftShardSpaces := []interface{}{}
	for _, item := range s.Res.CompositeRaftShardSpaces {
		compositeRaftShardSpaces = append(compositeRaftShardSpaces, CompositeRaftShardSpaceToMap(item))
	}
	s.D.Set("composite_raft_shard_spaces", compositeRaftShardSpaces)

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{DistributedDbConnectionStringToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbBackupConfig != nil {
		s.D.Set("db_backup_config", []interface{}{DistributedDbBackupConfigToMap(s.Res.DbBackupConfig)})
	} else {
		s.D.Set("db_backup_config", nil)
	}

	s.D.Set("db_deployment_type", s.Res.DbDeploymentType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	gdsControlNodeDetails := []interface{}{}
	for _, item := range s.Res.GdsControlNodeDetails {
		gdsControlNodeDetails = append(gdsControlNodeDetails, DistributedDatabaseGdsControlNodeToMap(item))
	}
	s.D.Set("gds_control_node_details", gdsControlNodeDetails)

	gsmDetails := []interface{}{}
	for _, item := range s.Res.GsmDetails {
		gsmDetails = append(gsmDetails, DistributedDatabaseGsmToMap(item))
	}
	s.D.Set("gsm_details", gsmDetails)

	if s.Res.GsmSshPublicKey != nil {
		s.D.Set("gsm_ssh_public_key", *s.Res.GsmSshPublicKey)
	}

	if s.Res.LatestGsmImageDetails != nil {
		s.D.Set("latest_gsm_image_details", []interface{}{DistributedDbGsmImageToMap(s.Res.LatestGsmImageDetails)})
	} else {
		s.D.Set("latest_gsm_image_details", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerPort != nil {
		s.D.Set("listener_port", *s.Res.ListenerPort)
	}

	if s.Res.ListenerPortTls != nil {
		s.D.Set("listener_port_tls", *s.Res.ListenerPortTls)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DistributedDbMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	s.D.Set("notification_topic_ids", s.Res.NotificationTopicIds)

	if s.Res.OnsPortLocal != nil {
		s.D.Set("ons_port_local", *s.Res.OnsPortLocal)
	}

	if s.Res.OnsPortRemote != nil {
		s.D.Set("ons_port_remote", *s.Res.OnsPortRemote)
	}

	if s.Res.Prefix != nil {
		s.D.Set("prefix", *s.Res.Prefix)
	}

	s.D.Set("private_endpoint_ids", s.Res.PrivateEndpointIds)

	s.D.Set("replication_method", s.Res.ReplicationMethod)

	if s.Res.ScanListenerPort != nil {
		s.D.Set("scan_listener_port", *s.Res.ScanListenerPort)
	}

	s.D.Set("sharding_method", s.Res.ShardingMethod)

	s.D.Set("state", s.Res.LifecycleState)
	if err := preserveOperationalState(s.D); err != nil {
		return err
	}

	if s.Res.SystemChunkCount != nil {
		s.D.Set("system_chunk_count", *s.Res.SystemChunkCount)
	}

	if s.Res.SystemDataGuardDatabases != nil {
		s.D.Set("system_data_guard_databases", []interface{}{SystemDataGuardDatabaseToMap(s.Res.SystemDataGuardDatabases)})
	} else {
		s.D.Set("system_data_guard_databases", nil)
	}

	systemRaftClusters := []interface{}{}
	for _, item := range s.Res.SystemRaftClusters {
		systemRaftClusters = append(systemRaftClusters, SystemRaftClusterToMap(item))
	}
	s.D.Set("system_raft_clusters", systemRaftClusters)

	if s.Res.SystemRaftReplicationUnitCount != nil {
		s.D.Set("system_raft_replication_unit_count", *s.Res.SystemRaftReplicationUnitCount)
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

	userShardSpaces := []interface{}{}
	for _, item := range s.Res.UserShardSpaces {
		userShardSpaces = append(userShardSpaces, UserShardSpaceToMap(item))
	}
	s.D.Set("user_shard_spaces", userShardSpaces)

	vcnNsgIds := []interface{}{}
	for _, item := range s.Res.VcnNsgIds {
		vcnNsgIds = append(vcnNsgIds, VcnNsgIdsDetailsToMap(item, true))
	}
	s.D.Set("vcn_nsg_ids", vcnNsgIds)

	return nil
}
