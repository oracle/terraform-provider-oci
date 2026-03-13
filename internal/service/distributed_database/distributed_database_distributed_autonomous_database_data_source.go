// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedAutonomousDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["distributed_autonomous_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	//fieldMap["metadata"] = &schema.Schema{
	//	Type: schema.TypeString,
	// WORKAROUND / FIX REQUIRED FOR GENERATED CODE ISSUE:
	//
	// Terraform provider internal validation fails with:
	//   "metadata: Elem must be set for lists"
	//   "metadata: One of optional, required, or computed must be set"
	//
	// Root cause:
	// The code generator emitted an invalid schema definition for the `metadata`
	// field in multiple Distributed Database data sources:
	//
	//   - metadata is defined as TypeList but missing Elem
	//   - metadata has no Optional / Required / Computed flag set
	//
	// Terraform schema rules require:
	//   - TypeList / TypeSet MUST define Elem
	//   - Every schema field MUST specify exactly one of:
	//       Optional, Required, or Computed
	//
	// Because of this, Terraform fails during InternalValidate *before*
	// any user configuration is evaluated, making this a provider-side bug.
	//
	// Correct schema shape must be:
	//
	//   "metadata": {
	//       Type:     schema.TypeList,
	//       Computed: true,
	//       Elem: &schema.Resource{
	//           Schema: <metadata schema>
	//       },
	//   }
	//
	// Affected data sources:
	//   - oci_distributed_database_distributed_autonomous_databases
	//   - oci_distributed_database_distributed_autonomous_database
	//   - oci_distributed_database_distributed_databases
	//   - oci_distributed_database_distributed_database
	//
	// This must be fixed in the code generator to avoid recurring regressions.
	//
	// See JIRA: TOP-9438
	//Required: false,
	//Optional: true,
	//}

	fieldMap["metadata_query"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Optional query parameter forwarded to the GET /distributedAutonomousDatabases API as `metadata`.",
	}

	fieldMap["metadata"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"map": {
					Type:     schema.TypeMap,
					Computed: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}

	return tfresource.GetSingularDataSourceItemSchemaWithContext(DistributedDatabaseDistributedAutonomousDatabaseResource(), fieldMap, readSingularDistributedDatabaseDistributedAutonomousDatabaseWithContext)
}

func readSingularDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedAutonomousDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedAutonomousDbServiceClient
	Res    *oci_distributed_database.GetDistributedAutonomousDatabaseResponse
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedAutonomousDatabaseRequest{}

	if distributedAutonomousDatabaseId, ok := s.D.GetOkExists("distributed_autonomous_database_id"); ok {
		tmp := distributedAutonomousDatabaseId.(string)
		request.DistributedAutonomousDatabaseId = &tmp
	}

	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The data source treats 'metadata' as a list and calls an undefined helper (mapTostring),
	// but GetDistributedAutonomousDatabaseRequest.Metadata is a *string in the OCI Go SDK.
	// Fix: assign the schema string directly to request.Metadata; do not index into a list
	// or call a non-existent mapper.
	// See JIRA: TOP-9424

	/*if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapTostring(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = &tmp
		}
	}*/

	if v, ok := s.D.GetOkExists("metadata_query"); ok {
		tmp := v.(string)
		request.Metadata = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.GetDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	catalogDetails := []interface{}{}
	for _, item := range s.Res.CatalogDetails {
		catalogDetails = append(catalogDetails, DistributedAutonomousDatabaseCatalogToMap(item))
	}
	s.D.Set("catalog_details", catalogDetails)

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.Chunks != nil {
		s.D.Set("chunks", *s.Res.Chunks)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{DistributedAutonomousDatabaseConnectionStringToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbBackupConfig != nil {
		s.D.Set("db_backup_config", []interface{}{DistributedAutonomousDbBackupConfigToMap(s.Res.DbBackupConfig)})
	} else {
		s.D.Set("db_backup_config", nil)
	}

	s.D.Set("db_deployment_type", s.Res.DbDeploymentType)

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	gsmDetails := []interface{}{}
	for _, item := range s.Res.GsmDetails {
		gsmDetails = append(gsmDetails, DistributedAutonomousDatabaseGsmToMap(item))
	}
	s.D.Set("gsm_details", gsmDetails)

	if s.Res.LatestGsmImage != nil {
		s.D.Set("latest_gsm_image", []interface{}{DistributedAutonomousDatabaseGsmImageToMap(s.Res.LatestGsmImage)})
	} else {
		s.D.Set("latest_gsm_image", nil)
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
		s.D.Set("metadata", []interface{}{DistributedAutonomousDbMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

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

	if s.Res.ReplicationFactor != nil {
		s.D.Set("replication_factor", *s.Res.ReplicationFactor)
	}

	s.D.Set("replication_method", s.Res.ReplicationMethod)

	if s.Res.ReplicationUnit != nil {
		s.D.Set("replication_unit", *s.Res.ReplicationUnit)
	}

	shardDetails := []interface{}{}
	for _, item := range s.Res.ShardDetails {
		shardDetails = append(shardDetails, DistributedAutonomousDatabaseShardToMap(item))
	}
	s.D.Set("shard_details", shardDetails)

	s.D.Set("sharding_method", s.Res.ShardingMethod)

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
