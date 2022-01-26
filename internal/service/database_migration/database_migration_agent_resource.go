// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database_migration "github.com/oracle/oci-go-sdk/v56/databasemigration"
)

func DatabaseMigrationAgentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMigrationAgent,
		Read:     readDatabaseMigrationAgent,
		Update:   updateDatabaseMigrationAgent,
		Delete:   deleteDatabaseMigrationAgent,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseMigrationAgent(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMigrationAgent(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseMigrationAgent(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseMigrationAgent(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseMigrationAgentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_migration.DatabaseMigrationClient
	Res                    *oci_database_migration.Agent
	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationAgentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseMigrationAgentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesCreating),
	}
}

func (s *DatabaseMigrationAgentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesActive),
	}
}

func (s *DatabaseMigrationAgentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesDeleting),
	}
}

func (s *DatabaseMigrationAgentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesDeleted),
	}
}

func (s *DatabaseMigrationAgentResourceCrud) Create() error {
	request := oci_database_migration.UpdateAgentRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	/*if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}*/

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		tmp := publicKey.(string)
		request.PublicKey = &tmp
	}

	if streamId, ok := s.D.GetOkExists("stream_id"); ok {
		tmp := streamId.(string)
		request.StreamId = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.UpdateAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Agent
	return nil
}

func (s *DatabaseMigrationAgentResourceCrud) Get() error {
	request := oci_database_migration.GetAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.GetAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Agent
	return nil
}

func (s *DatabaseMigrationAgentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_migration.UpdateAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

	/*if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}*/

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		tmp := publicKey.(string)
		request.PublicKey = &tmp
	}

	if streamId, ok := s.D.GetOkExists("stream_id"); ok {
		tmp := streamId.(string)
		request.StreamId = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.UpdateAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Agent
	return nil
}

func (s *DatabaseMigrationAgentResourceCrud) Delete() error {
	request := oci_database_migration.DeleteAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.DeleteAgent(context.Background(), request)
	return err
}

func (s *DatabaseMigrationAgentResourceCrud) SetData() error {
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

	if s.Res.PublicKey != nil {
		s.D.Set("public_key", *s.Res.PublicKey)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StreamId != nil {
		s.D.Set("stream_id", *s.Res.StreamId)
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

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func AgentSummaryToMap(obj oci_database_migration.AgentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.StreamId != nil {
		result["stream_id"] = string(*obj.StreamId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *DatabaseMigrationAgentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_migration.ChangeAgentCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AgentId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.ChangeAgentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
