// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseDbNodeConsoleConnectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["db_node_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseDbNodeConsoleConnectionResource(), fieldMap, readSingularDatabaseDbNodeConsoleConnection)
}

func readSingularDatabaseDbNodeConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleConnectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodeConsoleConnectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetConsoleConnectionResponse
}

func (s *DatabaseDbNodeConsoleConnectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeConsoleConnectionDataSourceCrud) Get() error {
	request := oci_database.GetConsoleConnectionRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.ConsoleConnectionId = &tmp
	}

	dbNodeId, id, error := ParseDbNodeConsoleConnectionCompositeId(s.D.Id())
	if error == nil {
		request.DbNodeId = &dbNodeId
		request.ConsoleConnectionId = &id
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbNodeConsoleConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
