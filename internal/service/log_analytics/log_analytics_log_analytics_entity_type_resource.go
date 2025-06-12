// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

type sdkLogger interface {
	//LogLevel returns the log level of sdkLogger
	LogLevel() int

	//Log logs v with the provided format if the current log level is loglevel
	Log(logLevel int, format string, v ...interface{}) error
}

const debugLogging = 2

var defaultLogger sdkLogger

func LogAnalyticsLogAnalyticsEntityTypeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsEntityType,
		Read:     readLogAnalyticsLogAnalyticsEntityType,
		Delete:   deleteLogAnalyticsLogAnalyticsEntityType,
		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_agent_eligibility_status": {
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

			"properties": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
		},
	}
}

func createLogAnalyticsLogAnalyticsEntityType(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsEntityType(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func deleteLogAnalyticsLogAnalyticsEntityType(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LogAnalyticsLogAnalyticsEntityTypeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.LogAnalyticsEntityType
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) ID() string {
	return GetLogAnalyticsEntityTypeCompositeId(s.D.Get("namespace").(string), s.D.Get("name").(string))
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_log_analytics.EntityLifecycleStatesActive),
	}
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_log_analytics.EntityLifecycleStatesDeleted),
	}
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) Create() error {
	request := oci_log_analytics.CreateLogAnalyticsEntityTypeRequest{}

	if category, ok := s.D.GetOkExists("category"); ok {
		tmp := category.(string)
		request.Category = &tmp
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
		interfaces := properties.([]interface{})
		tmp := make([]oci_log_analytics.EntityTypeProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "properties", stateDataIndex)
			converted, err := s.mapToEntityTypeProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("properties") {
			request.Properties = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")
	_, err := s.Client.CreateLogAnalyticsEntityType(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsEntityTypeRequest{}

	if entityTypeName, ok := s.D.GetOkExists("name"); ok {
		tmp := entityTypeName.(string)
		request.EntityTypeName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	namespace, entityTypeName, err := parseLogAnalyticsEntityTypeCompositeId(s.D.Id())
	if err == nil {
		request.NamespaceName = &namespace
		request.EntityTypeName = &entityTypeName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetLogAnalyticsEntityType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsEntityType

	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) SetData() error {

	namespace, entityTypeName, err := parseLogAnalyticsEntityTypeCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &entityTypeName)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Category != nil {
		s.D.Set("category", *s.Res.Category)
	}

	s.D.Set("cloud_type", *&s.Res.CloudType)

	if s.Res.InternalName != nil {
		s.D.Set("internal_name", *&s.Res.InternalName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("management_agent_eligibility_status", s.Res.ManagementAgentEligibilityStatus)

	propetiesMap := []interface{}{}
	for _, item := range s.Res.Properties {
		propetiesMap = append(propetiesMap, EntityPropertiesToMap(item))
	}
	s.D.Set("properties", propetiesMap)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}
	return nil
}

func GetLogAnalyticsEntityTypeCompositeId(namespace string, entityTypeName string) string {
	namespace = url.PathEscape(namespace)
	entityTypeName = url.PathEscape(entityTypeName)
	compositeId := "namespaces/" + namespace + "/logAnalyticsEntityTypes/" + entityTypeName
	return compositeId
}

func parseLogAnalyticsEntityTypeCompositeId(compositeId string) (namespace string, entityTypeName string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/logAnalyticsEntityTypes/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	entityTypeName, _ = url.PathUnescape(parts[3])

	return
}

func (s *LogAnalyticsLogAnalyticsEntityTypeResourceCrud) mapToEntityTypeProperty(fieldKeyFormat string) (oci_log_analytics.EntityTypeProperty, error) {
	result := oci_log_analytics.EntityTypeProperty{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func EntityTypePropertyToMap(obj oci_log_analytics.EntityTypeProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func LogAnalyticsEntityTypeSummaryToMap(obj oci_log_analytics.LogAnalyticsEntityTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Category != nil {
		result["category"] = string(*obj.Category)
	}

	result["cloud_type"] = string(obj.CloudType)

	if obj.InternalName != nil {
		result["internal_name"] = string(*obj.InternalName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func EntityPropertiesToMap(obj oci_log_analytics.EntityTypeProperty) map[string]interface{} {
	result := map[string]interface{}{}
	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}
	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	return result
}
