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

func LogAnalyticsLogAnalyticsLogGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsLogGroup,
		Read:     readLogAnalyticsLogAnalyticsLogGroup,
		Update:   updateLogAnalyticsLogAnalyticsLogGroup,
		Delete:   deleteLogAnalyticsLogAnalyticsLogGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
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
			"description": {
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

func createLogAnalyticsLogAnalyticsLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsLogAnalyticsLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsLogAnalyticsLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsLogAnalyticsLogGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.LogAnalyticsLogGroup
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsLogGroupResourceCrud) ID() string {
	return getLogAnalyticsLogGroupCompositeId(*s.Res.Id, s.D.Get("namespace").(string))
}

func (s *LogAnalyticsLogAnalyticsLogGroupResourceCrud) Create() error {
	request := oci_log_analytics.CreateLogAnalyticsLogGroupRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.CreateLogAnalyticsLogGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsLogGroup
	return nil
}

func (s *LogAnalyticsLogAnalyticsLogGroupResourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsLogGroupRequest{}

	tmp := s.D.Id()
	request.LogAnalyticsLogGroupId = &tmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	logAnalyticsLogGroupId, namespace, err := parseLogAnalyticsLogGroupCompositeId(s.D.Id())
	if err == nil {
		request.LogAnalyticsLogGroupId = &logAnalyticsLogGroupId
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetLogAnalyticsLogGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsLogGroup
	return nil
}

func (s *LogAnalyticsLogAnalyticsLogGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_log_analytics.UpdateLogAnalyticsLogGroupRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.LogAnalyticsLogGroupId = &tmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateLogAnalyticsLogGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsLogGroup
	return nil
}

func (s *LogAnalyticsLogAnalyticsLogGroupResourceCrud) Delete() error {
	request := oci_log_analytics.DeleteLogAnalyticsLogGroupRequest{}

	tmp := s.D.Id()
	request.LogAnalyticsLogGroupId = &tmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.DeleteLogAnalyticsLogGroup(context.Background(), request)
	return err
}

func (s *LogAnalyticsLogAnalyticsLogGroupResourceCrud) SetData() error {

	logAnalyticsLogGroupId, namespace, err := parseLogAnalyticsLogGroupCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(logAnalyticsLogGroupId)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func getLogAnalyticsLogGroupCompositeId(logAnalyticsLogGroupId string, namespace string) string {
	logAnalyticsLogGroupId = url.PathEscape(logAnalyticsLogGroupId)
	namespace = url.PathEscape(namespace)
	compositeId := "namespaces/" + namespace + "/logAnalyticsLogGroups/" + logAnalyticsLogGroupId
	return compositeId
}

func parseLogAnalyticsLogGroupCompositeId(compositeId string) (logAnalyticsLogGroupId string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/logAnalyticsLogGroups/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	logAnalyticsLogGroupId, _ = url.PathUnescape(parts[3])

	return
}

func LogAnalyticsLogGroupSummaryToMap(obj oci_log_analytics.LogAnalyticsLogGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *LogAnalyticsLogAnalyticsLogGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_log_analytics.ChangeLogAnalyticsLogGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LogAnalyticsLogGroupId = &idTmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		changeCompartmentRequest.NamespaceName = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.ChangeLogAnalyticsLogGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
