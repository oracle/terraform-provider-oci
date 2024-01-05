// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeConsoleHistoryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDbNodeConsoleHistory,
		Read:     readDatabaseDbNodeConsoleHistory,
		Update:   updateDatabaseDbNodeConsoleHistory,
		Delete:   deleteDatabaseDbNodeConsoleHistory,
		Schema: map[string]*schema.Schema{
			// Required
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDbNodeConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDbNodeConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseDbNodeConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseDbNodeConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDbNodeConsoleHistoryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ConsoleHistory
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) ID() string {
	return GetDbNodeConsoleHistoryCompositeId(s.D.Get("db_node_id").(string), *s.Res.Id)
	//return GetDbNodeConsoleHistoryCompositeId(s.D.Get("console_history_id").(string), s.D.Get("db_node_id").(string))
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ConsoleHistoryLifecycleStateRequested),
	}
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ConsoleHistoryLifecycleStateSucceeded),
	}
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ConsoleHistoryLifecycleStateDeleting),
	}
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ConsoleHistoryLifecycleStateDeleted),
	}
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) Create() error {
	request := oci_database.CreateConsoleHistoryRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) Get() error {
	request := oci_database.GetConsoleHistoryRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	tmp := s.D.Id()
	request.ConsoleHistoryId = &tmp

	dbNodeId, consoleHistoryId, err := parseDbNodeConsoleHistoryCompositeId(s.D.Id())
	if err == nil {
		request.DbNodeId = &dbNodeId
		request.ConsoleHistoryId = &consoleHistoryId
		log.Printf("request during parse %s", request)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory

	return nil
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) Update() error {
	request := oci_database.UpdateConsoleHistoryRequest{}

	tmp := s.D.Id()
	request.ConsoleHistoryId = &tmp

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "node", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) Delete() error {
	request := oci_database.DeleteConsoleHistoryRequest{}

	tmp := s.D.Id()
	request.ConsoleHistoryId = &tmp

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteConsoleHistory(context.Background(), request)
	return err
}

func (s *DatabaseDbNodeConsoleHistoryResourceCrud) SetData() error {

	dbNodeId, consoleHistoryId, err := parseDbNodeConsoleHistoryCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(consoleHistoryId)
		s.D.Set("db_node_id", &dbNodeId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbNodeId != nil {
		s.D.Set("db_node_id", *s.Res.DbNodeId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("freeform_tags", s.Res.FreeformTags)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func GetDbNodeConsoleHistoryCompositeId(dbNodeId string, consoleHistoryId string) string {
	dbNodeId = url.PathEscape(dbNodeId)
	consoleHistoryId = url.PathEscape(consoleHistoryId)
	compositeId := "dbNodes/" + dbNodeId + "/consoleHistories/" + consoleHistoryId
	return compositeId
}

func parseDbNodeConsoleHistoryCompositeId(compositeId string) (dbNodeId string, consoleHistoryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dbNodes/.*/consoleHistories/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dbNodeId, _ = url.PathUnescape(parts[1])
	consoleHistoryId, _ = url.PathUnescape(parts[3])

	return
}

func ConsoleHistorySummaryToMap(obj oci_database.ConsoleHistorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DbNodeId != nil {
		result["db_node_id"] = string(*obj.DbNodeId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
