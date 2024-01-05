// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeConsoleHistoryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["console_history_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["db_node_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseDbNodeConsoleHistoryResource(), fieldMap, readSingularDatabaseDbNodeConsoleHistory)
}

func readSingularDatabaseDbNodeConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleHistoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodeConsoleHistoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetConsoleHistoryResponse
}

func (s *DatabaseDbNodeConsoleHistoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeConsoleHistoryDataSourceCrud) Get() error {
	request := oci_database.GetConsoleHistoryRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if consoleHistoryId, ok := s.D.GetOkExists("console_history_id"); ok {
		tmp := consoleHistoryId.(string)
		request.ConsoleHistoryId = &tmp
	}

	dbNodeId, consoleHistoryId, err := parseDbNodeConsoleHistoryCompositeId(s.D.Id())
	if err == nil {
		request.DbNodeId = &dbNodeId
		request.ConsoleHistoryId = &consoleHistoryId
		log.Printf("Request during parse %s", request)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	log.Printf("s.Res %s", s.Res)
	return nil
}

func (s *DatabaseDbNodeConsoleHistoryDataSourceCrud) SetData() error {
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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
