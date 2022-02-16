// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v58/loganalytics"
)

func LogAnalyticsLogAnalyticsEntityResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsEntity,
		Read:     readLogAnalyticsLogAnalyticsEntity,
		Update:   updateLogAnalyticsLogAnalyticsEntity,
		Delete:   deleteLogAnalyticsLogAnalyticsEntity,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_type_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"cloud_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"source_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timezone_region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"are_logs_collected": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_type_internal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_agent_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_agent_display_name": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsLogAnalyticsEntity(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsEntity(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsLogAnalyticsEntity(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsLogAnalyticsEntity(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsLogAnalyticsEntityResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.LogAnalyticsEntity
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) ID() string {
	return getLogAnalyticsEntityCompositeId(*s.Res.Id, s.D.Get("namespace").(string))
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_log_analytics.EntityLifecycleStatesActive),
	}
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_log_analytics.EntityLifecycleStatesDeleted),
	}
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) Create() error {
	request := oci_log_analytics.CreateLogAnalyticsEntityRequest{}

	if cloudResourceId, ok := s.D.GetOkExists("cloud_resource_id"); ok {
		tmp := cloudResourceId.(string)
		request.CloudResourceId = &tmp
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

	if entityTypeName, ok := s.D.GetOkExists("entity_type_name"); ok {
		tmp := entityTypeName.(string)
		request.EntityTypeName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		request.Properties = utils.ObjectMapToStringMap(properties.(map[string]interface{}))
	}

	if sourceId, ok := s.D.GetOkExists("source_id"); ok {
		tmp := sourceId.(string)
		request.SourceId = &tmp
	}

	if timezoneRegion, ok := s.D.GetOkExists("timezone_region"); ok {
		tmp := timezoneRegion.(string)
		request.TimezoneRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.CreateLogAnalyticsEntity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsEntity
	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsEntityRequest{}

	tmp := s.D.Id()
	request.LogAnalyticsEntityId = &tmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	logAnalyticsEntityId, namespace, err := parseLogAnalyticsEntityCompositeId(s.D.Id())
	if err == nil {
		request.LogAnalyticsEntityId = &logAnalyticsEntityId
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetLogAnalyticsEntity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsEntity
	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_log_analytics.UpdateLogAnalyticsEntityRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	tmp := s.D.Id()
	request.LogAnalyticsEntityId = &tmp

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok && s.D.HasChange("name") {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		request.Properties = utils.ObjectMapToStringMap(properties.(map[string]interface{}))
	}

	if timezoneRegion, ok := s.D.GetOkExists("timezone_region"); ok {
		tmp := timezoneRegion.(string)
		request.TimezoneRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateLogAnalyticsEntity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsEntity
	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) Delete() error {
	request := oci_log_analytics.DeleteLogAnalyticsEntityRequest{}

	tmp := s.D.Id()
	request.LogAnalyticsEntityId = &tmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.DeleteLogAnalyticsEntity(context.Background(), request)
	return err
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) SetData() error {

	logAnalyticsEntityId, namespace, err := parseLogAnalyticsEntityCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(logAnalyticsEntityId)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AreLogsCollected != nil {
		s.D.Set("are_logs_collected", *s.Res.AreLogsCollected)
	}

	if s.Res.CloudResourceId != nil {
		s.D.Set("cloud_resource_id", *s.Res.CloudResourceId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.EntityTypeInternalName != nil {
		s.D.Set("entity_type_internal_name", *s.Res.EntityTypeInternalName)
	}

	if s.Res.EntityTypeName != nil {
		s.D.Set("entity_type_name", *s.Res.EntityTypeName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ManagementAgentCompartmentId != nil {
		s.D.Set("management_agent_compartment_id", *s.Res.ManagementAgentCompartmentId)
	}

	if s.Res.ManagementAgentDisplayName != nil {
		s.D.Set("management_agent_display_name", *s.Res.ManagementAgentDisplayName)
	}

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("properties", s.Res.Properties)

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimezoneRegion != nil {
		s.D.Set("timezone_region", *s.Res.TimezoneRegion)
	}

	return nil
}

func getLogAnalyticsEntityCompositeId(logAnalyticsEntityId string, namespace string) string {
	logAnalyticsEntityId = url.PathEscape(logAnalyticsEntityId)
	namespace = url.PathEscape(namespace)
	compositeId := "namespaces/" + namespace + "/logAnalyticsEntities/" + logAnalyticsEntityId
	return compositeId
}

func parseLogAnalyticsEntityCompositeId(compositeId string) (logAnalyticsEntityId string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/logAnalyticsEntities/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	logAnalyticsEntityId, _ = url.PathUnescape(parts[3])

	return
}

func LogAnalyticsEntitySummaryToMap(obj oci_log_analytics.LogAnalyticsEntitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreLogsCollected != nil {
		result["are_logs_collected"] = bool(*obj.AreLogsCollected)
	}

	if obj.CloudResourceId != nil {
		result["cloud_resource_id"] = string(*obj.CloudResourceId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.EntityTypeInternalName != nil {
		result["entity_type_internal_name"] = string(*obj.EntityTypeInternalName)
	}

	if obj.EntityTypeName != nil {
		result["entity_type_name"] = string(*obj.EntityTypeName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ManagementAgentId != nil {
		result["management_agent_id"] = string(*obj.ManagementAgentId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.SourceId != nil {
		result["source_id"] = string(*obj.SourceId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TimezoneRegion != nil {
		result["timezone_region"] = string(*obj.TimezoneRegion)
	}

	return result
}

func (s *LogAnalyticsLogAnalyticsEntityResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_log_analytics.ChangeLogAnalyticsEntityCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LogAnalyticsEntityId = &idTmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		changeCompartmentRequest.NamespaceName = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.ChangeLogAnalyticsEntityCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
