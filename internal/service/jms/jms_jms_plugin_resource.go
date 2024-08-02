// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJmsPluginResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsJmsPlugin,
		Read:     readJmsJmsPlugin,
		Update:   updateJmsJmsPlugin,
		Delete:   deleteJmsJmsPlugin,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"fleet_id": {
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

			// Computed
			"agent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_distribution": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plugin_version": {
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
			"time_last_seen": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_registered": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createJmsJmsPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJmsPluginResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsJmsPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJmsPluginResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

func updateJmsJmsPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJmsPluginResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsJmsPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJmsPluginResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsJmsPluginResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms.JavaManagementServiceClient
	Res                    *oci_jms.JmsPlugin
	DisableNotFoundRetries bool
}

func (s *JmsJmsPluginResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *JmsJmsPluginResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *JmsJmsPluginResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_jms.JmsPluginLifecycleStateActive),
		string(oci_jms.JmsPluginLifecycleStateNeedsAttention),
	}
}

func (s *JmsJmsPluginResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *JmsJmsPluginResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_jms.JmsPluginLifecycleStateDeleted),
	}
}

func (s *JmsJmsPluginResourceCrud) Create() error {
	request := oci_jms.CreateJmsPluginRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.CreateJmsPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JmsPlugin
	return nil
}

func (s *JmsJmsPluginResourceCrud) Get() error {
	request := oci_jms.GetJmsPluginRequest{}

	tmp := s.D.Id()
	request.JmsPluginId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.GetJmsPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JmsPlugin
	return nil
}

func (s *JmsJmsPluginResourceCrud) Update() error {
	request := oci_jms.UpdateJmsPluginRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.JmsPluginId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.UpdateJmsPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JmsPlugin
	return nil
}

func (s *JmsJmsPluginResourceCrud) Delete() error {
	request := oci_jms.DeleteJmsPluginRequest{}

	tmp := s.D.Id()
	request.JmsPluginId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	_, err := s.Client.DeleteJmsPlugin(context.Background(), request)
	return err
}

func (s *JmsJmsPluginResourceCrud) SetData() error {
	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	s.D.Set("agent_type", s.Res.AgentType)

	s.D.Set("availability_status", s.Res.AvailabilityStatus)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.FleetId != nil {
		s.D.Set("fleet_id", *s.Res.FleetId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.OsArchitecture != nil {
		s.D.Set("os_architecture", *s.Res.OsArchitecture)
	}

	if s.Res.OsDistribution != nil {
		s.D.Set("os_distribution", *s.Res.OsDistribution)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.PluginVersion != nil {
		s.D.Set("plugin_version", *s.Res.PluginVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeLastSeen != nil {
		s.D.Set("time_last_seen", s.Res.TimeLastSeen.String())
	}

	if s.Res.TimeRegistered != nil {
		s.D.Set("time_registered", s.Res.TimeRegistered.String())
	}

	return nil
}

func JmsPluginSummaryToMap(obj oci_jms.JmsPluginSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	result["agent_type"] = string(obj.AgentType)

	result["availability_status"] = string(obj.AvailabilityStatus)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.FleetId != nil {
		result["fleet_id"] = string(*obj.FleetId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.OsArchitecture != nil {
		result["os_architecture"] = string(*obj.OsArchitecture)
	}

	if obj.OsDistribution != nil {
		result["os_distribution"] = string(*obj.OsDistribution)
	}

	result["os_family"] = string(obj.OsFamily)

	if obj.PluginVersion != nil {
		result["plugin_version"] = string(*obj.PluginVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeLastSeen != nil {
		result["time_last_seen"] = obj.TimeLastSeen.String()
	}

	if obj.TimeRegistered != nil {
		result["time_registered"] = obj.TimeRegistered.String()
	}

	return result
}
