// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAgentPluginResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudBridgeAgentPlugin,
		Read:     readCloudBridgeAgentPlugin,
		Update:   updateCloudBridgeAgentPlugin,
		Delete:   deleteCloudBridgeAgentPlugin,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"plugin_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"desired_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
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

func createCloudBridgeAgentPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentPluginResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudBridgeAgentPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentPluginResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.ReadResource(sync)
}

func updateCloudBridgeAgentPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentPluginResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudBridgeAgentPlugin(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CloudBridgeAgentPluginResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_bridge.OcbAgentSvcClient
	Res                    *oci_cloud_bridge.Plugin
	DisableNotFoundRetries bool
}

func (s *CloudBridgeAgentPluginResourceCrud) ID() string {
	return GetAgentPluginCompositeId(s.D.Get("agent_id").(string), s.D.Get("plugin_name").(string))
}

func (s *CloudBridgeAgentPluginResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *CloudBridgeAgentPluginResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_bridge.PluginLifecycleStateActive),
		string(oci_cloud_bridge.PluginLifecycleStateNeedsAttention),
	}
}

func (s *CloudBridgeAgentPluginResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CloudBridgeAgentPluginResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_bridge.PluginLifecycleStateDeleted),
	}
}

func (s *CloudBridgeAgentPluginResourceCrud) Create() error {
	request := oci_cloud_bridge.UpdatePluginRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if desiredState, ok := s.D.GetOkExists("desired_state"); ok {
		request.DesiredState = oci_cloud_bridge.PluginDesiredStateEnum(desiredState.(string))
	}

	if pluginName, ok := s.D.GetOkExists("plugin_name"); ok {
		tmp := pluginName.(string)
		request.PluginName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.UpdatePlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Plugin
	return nil
}

func (s *CloudBridgeAgentPluginResourceCrud) Get() error {
	request := oci_cloud_bridge.GetPluginRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if pluginName, ok := s.D.GetOkExists("plugin_name"); ok {
		tmp := pluginName.(string)
		request.PluginName = &tmp
	}

	agentId, pluginName, err := parseAgentPluginCompositeId(s.D.Id())
	if err == nil {
		request.AgentId = &agentId
		request.PluginName = &pluginName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.GetPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Plugin
	return nil
}

func (s *CloudBridgeAgentPluginResourceCrud) Update() error {
	request := oci_cloud_bridge.UpdatePluginRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if desiredState, ok := s.D.GetOkExists("desired_state"); ok {
		request.DesiredState = oci_cloud_bridge.PluginDesiredStateEnum(desiredState.(string))
	}

	if pluginName, ok := s.D.GetOkExists("plugin_name"); ok {
		tmp := pluginName.(string)
		request.PluginName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.UpdatePlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Plugin
	return nil
}

func (s *CloudBridgeAgentPluginResourceCrud) SetData() error {

	agentId, pluginName, err := parseAgentPluginCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("agent_id", &agentId)
		s.D.Set("plugin_name", &pluginName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("desired_state", s.Res.DesiredState)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PluginVersion != nil {
		s.D.Set("plugin_version", *s.Res.PluginVersion)
	}

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

func GetAgentPluginCompositeId(agentId string, pluginName string) string {
	agentId = url.PathEscape(agentId)
	pluginName = url.PathEscape(pluginName)
	compositeId := "agents/" + agentId + "/plugins/" + pluginName
	return compositeId
}

func parseAgentPluginCompositeId(compositeId string) (agentId string, pluginName string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("agents/.*/plugins/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	agentId, _ = url.PathUnescape(parts[1])
	pluginName, _ = url.PathUnescape(parts[3])

	return
}
