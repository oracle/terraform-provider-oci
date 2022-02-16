// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseDbNodeConsoleConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDbNodeConsoleConnection,
		Read:     readDatabaseDbNodeConsoleConnection,
		Delete:   deleteDatabaseDbNodeConsoleConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDbNodeConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDbNodeConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func deleteDatabaseDbNodeConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDbNodeConsoleConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ConsoleConnection
	DisableNotFoundRetries bool
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) ID() string {
	return GetDbNodeConsoleConnectionCompositeId(s.D.Get("db_node_id").(string), *s.Res.Id)
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ConsoleConnectionLifecycleStateCreating),
	}
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ConsoleConnectionLifecycleStateActive),
	}
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ConsoleConnectionLifecycleStateDeleting),
	}
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ConsoleConnectionLifecycleStateDeleted),
	}
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) Create() error {
	request := oci_database.CreateConsoleConnectionRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		tmp := publicKey.(string)
		request.PublicKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleConnection

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) Get() error {
	request := oci_database.GetConsoleConnectionRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	tmp := s.D.Id()
	request.ConsoleConnectionId = &tmp

	dbNodeId, id, err := ParseDbNodeConsoleConnectionCompositeId(s.D.Id())
	if err == nil {
		request.DbNodeId = &dbNodeId
		request.ConsoleConnectionId = &id
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleConnection
	return nil
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) Delete() error {
	request := oci_database.DeleteConsoleConnectionRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	dbNodeId, id, error := ParseDbNodeConsoleConnectionCompositeId(s.D.Id())
	if error == nil {
		request.DbNodeId = &dbNodeId
		request.ConsoleConnectionId = &id
	} else {
		log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteConsoleConnection(context.Background(), request)
	return err
}

func (s *DatabaseDbNodeConsoleConnectionResourceCrud) SetData() error {

	dbNodeId, id, err := ParseDbNodeConsoleConnectionCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("db_node_id", &dbNodeId)
		s.D.Set("id", &id)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.DbNodeId != nil {
		s.D.Set("db_node_id", *s.Res.DbNodeId)
	}

	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}

func GetDbNodeConsoleConnectionCompositeId(dbNodeId string, id string) string {
	dbNodeId = url.PathEscape(dbNodeId)
	id = url.PathEscape(id)
	compositeId := "dbNodes/" + dbNodeId + "/consoleConnections/" + id
	return compositeId
}

func ParseDbNodeConsoleConnectionCompositeId(compositeId string) (dbNodeId string, id string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dbNodes/.*/consoleConnections/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dbNodeId, _ = url.PathUnescape(parts[1])
	id, _ = url.PathUnescape(parts[3])

	return
}
