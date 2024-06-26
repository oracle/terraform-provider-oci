// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package database_migration

//
//import (
//	"context"
//
//	"github.com/oracle/terraform-provider-oci/internal/client"
//	"github.com/oracle/terraform-provider-oci/internal/tfresource"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
//)
//
//func DatabaseMigrationAgentDataSource() *schema.Resource {
//	fieldMap := make(map[string]*schema.Schema)
//	fieldMap["agent_id"] = &schema.Schema{
//		Type:     schema.TypeString,
//		Required: true,
//	}
//	return tfresource.GetSingularDataSourceItemSchema(DatabaseMigrationAgentResource(), fieldMap, readSingularDatabaseMigrationAgent)
//}
//
//func readSingularDatabaseMigrationAgent(d *schema.ResourceData, m interface{}) error {
//	sync := &DatabaseMigrationAgentDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type DatabaseMigrationAgentDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_database_migration.DatabaseMigrationClient
//	Res    *oci_database_migration.GetAgentResponse
//}
//
//func (s *DatabaseMigrationAgentDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *DatabaseMigrationAgentDataSourceCrud) Get() error {
//	request := oci_database_migration.GetAgentRequest{}
//
//	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
//		tmp := agentId.(string)
//		request.AgentId = &tmp
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")
//
//	response, err := s.Client.GetAgent(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *DatabaseMigrationAgentDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(*s.Res.Id)
//
//	if s.Res.CompartmentId != nil {
//		s.D.Set("compartment_id", *s.Res.CompartmentId)
//	}
//
//	if s.Res.DefinedTags != nil {
//		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
//	}
//
//	if s.Res.DisplayName != nil {
//		s.D.Set("display_name", *s.Res.DisplayName)
//	}
//
//	s.D.Set("freeform_tags", s.Res.FreeformTags)
//
//	if s.Res.LifecycleDetails != nil {
//		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
//	}
//
//	if s.Res.PublicKey != nil {
//		s.D.Set("public_key", *s.Res.PublicKey)
//	}
//
//	s.D.Set("state", s.Res.LifecycleState)
//
//	if s.Res.StreamId != nil {
//		s.D.Set("stream_id", *s.Res.StreamId)
//	}
//
//	if s.Res.SystemTags != nil {
//		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
//	}
//
//	if s.Res.TimeCreated != nil {
//		s.D.Set("time_created", s.Res.TimeCreated.String())
//	}
//
//	if s.Res.TimeUpdated != nil {
//		s.D.Set("time_updated", s.Res.TimeUpdated.String())
//	}
//
//	if s.Res.Version != nil {
//		s.D.Set("version", *s.Res.Version)
//	}
//
//	return nil
//}
