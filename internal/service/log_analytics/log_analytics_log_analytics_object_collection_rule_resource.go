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

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsLogAnalyticsObjectCollectionRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsObjectCollectionRule,
		Read:     readLogAnalyticsLogAnalyticsObjectCollectionRule,
		Update:   updateLogAnalyticsLogAnalyticsObjectCollectionRule,
		Delete:   deleteLogAnalyticsLogAnalyticsObjectCollectionRule,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_source_name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"os_bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"os_namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"char_encoding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collection_type": {
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_id": {
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
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_force_historic_collection": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"log_set": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_set_ext_regex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_set_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object_name_filters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"overrides": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"match_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"property_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"property_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"poll_since": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"poll_till": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timezone": {
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

func createLogAnalyticsLogAnalyticsObjectCollectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsObjectCollectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsLogAnalyticsObjectCollectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsLogAnalyticsObjectCollectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.LogAnalyticsObjectCollectionRule
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) ID() string {
	return GetLogAnalyticsObjectCollectionRuleCompositeId(*s.Res.Id, s.D.Get("namespace").(string))
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_log_analytics.ObjectCollectionRuleLifecycleStatesActive),
	}
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_log_analytics.ObjectCollectionRuleLifecycleStatesDeleted),
	}
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) Create() error {
	request := oci_log_analytics.CreateLogAnalyticsObjectCollectionRuleRequest{}

	if charEncoding, ok := s.D.GetOkExists("char_encoding"); ok {
		tmp := charEncoding.(string)
		request.CharEncoding = &tmp
	}

	if collectionType, ok := s.D.GetOkExists("collection_type"); ok {
		request.CollectionType = oci_log_analytics.ObjectCollectionRuleCollectionTypesEnum(collectionType.(string))
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if isForceHistoricCollection, ok := s.D.GetOkExists("is_force_historic_collection"); ok {
		tmp := isForceHistoricCollection.(bool)
		request.IsForceHistoricCollection = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	if logSet, ok := s.D.GetOkExists("log_set"); ok {
		tmp := logSet.(string)
		request.LogSet = &tmp
	}

	if logSetExtRegex, ok := s.D.GetOkExists("log_set_ext_regex"); ok {
		tmp := logSetExtRegex.(string)
		request.LogSetExtRegex = &tmp
	}

	if logSetKey, ok := s.D.GetOkExists("log_set_key"); ok {
		request.LogSetKey = oci_log_analytics.LogSetKeyTypesEnum(logSetKey.(string))
	}

	if logSourceName, ok := s.D.GetOkExists("log_source_name"); ok {
		tmp := logSourceName.(string)
		request.LogSourceName = &tmp
	}

	if logType, ok := s.D.GetOkExists("log_type"); ok {
		request.LogType = oci_log_analytics.LogTypesEnum(logType.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if objectNameFilters, ok := s.D.GetOkExists("object_name_filters"); ok {
		interfaces := objectNameFilters.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object_name_filters") {
			request.ObjectNameFilters = tmp
		}
	}

	if osBucketName, ok := s.D.GetOkExists("os_bucket_name"); ok {
		tmp := osBucketName.(string)
		request.OsBucketName = &tmp
	}

	if osNamespace, ok := s.D.GetOkExists("os_namespace"); ok {
		tmp := osNamespace.(string)
		request.OsNamespace = &tmp
	}

	if overrides, ok := s.D.GetOkExists("overrides"); ok {
		interfaces := overrides.([]interface{})
		tmp := make([]oci_log_analytics.PropertyOverride, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", "overrides", stateDataIndex)
			converted, err := s.mapToPropertyOverrides(fieldKeyFormatNextLevel)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("overrides") {
			request.Overrides = map[string][]oci_log_analytics.PropertyOverride{"items": tmp}
		}
	}

	if pollSince, ok := s.D.GetOkExists("poll_since"); ok {
		tmp := pollSince.(string)
		request.PollSince = &tmp
	}

	if pollTill, ok := s.D.GetOkExists("poll_till"); ok {
		tmp := pollTill.(string)
		request.PollTill = &tmp
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.CreateLogAnalyticsObjectCollectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsObjectCollectionRule
	return nil
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsObjectCollectionRuleRequest{}

	tmp := s.D.Id()
	request.LogAnalyticsObjectCollectionRuleId = &tmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}
	if logAnalyticsObjectCollectionRuleId, ok := s.D.GetOkExists("log_analytics_object_collection_rule_id"); ok {
		tmp := logAnalyticsObjectCollectionRuleId.(string)
		request.LogAnalyticsObjectCollectionRuleId = &tmp
	}

	logAnalyticsObjectCollectionRuleId, namespace, err := parseLogAnalyticsObjectCollectionRuleCompositeId(s.D.Id())
	if err == nil {
		request.LogAnalyticsObjectCollectionRuleId = &logAnalyticsObjectCollectionRuleId
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetLogAnalyticsObjectCollectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsObjectCollectionRule
	return nil
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_log_analytics.UpdateLogAnalyticsObjectCollectionRuleRequest{}

	if charEncoding, ok := s.D.GetOkExists("char_encoding"); ok {
		tmp := charEncoding.(string)
		request.CharEncoding = &tmp
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

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	tmp := s.D.Id()
	request.LogAnalyticsObjectCollectionRuleId = &tmp

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	if logSet, ok := s.D.GetOkExists("log_set"); ok {
		tmp := logSet.(string)
		request.LogSet = &tmp
	}

	if logSetExtRegex, ok := s.D.GetOkExists("log_set_ext_regex"); ok {
		tmp := logSetExtRegex.(string)
		request.LogSetExtRegex = &tmp
	}

	if logSetKey, ok := s.D.GetOkExists("log_set_key"); ok {
		request.LogSetKey = oci_log_analytics.LogSetKeyTypesEnum(logSetKey.(string))
	}

	if logSourceName, ok := s.D.GetOkExists("log_source_name"); ok {
		tmp := logSourceName.(string)
		request.LogSourceName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if objectNameFilters, ok := s.D.GetOkExists("object_name_filters"); ok {
		interfaces := objectNameFilters.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object_name_filters") {
			request.ObjectNameFilters = tmp
		}
	}

	if overrides, ok := s.D.GetOkExists("overrides"); ok {
		interfaces := overrides.([]interface{})
		tmp := make([]oci_log_analytics.PropertyOverride, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", "overrides", stateDataIndex)
			converted, err := s.mapToPropertyOverrides(fieldKeyFormatNextLevel)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("overrides") {
			request.Overrides = map[string][]oci_log_analytics.PropertyOverride{"items": tmp}
		}
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateLogAnalyticsObjectCollectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsObjectCollectionRule
	return nil
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) Delete() error {
	request := oci_log_analytics.DeleteLogAnalyticsObjectCollectionRuleRequest{}

	tmp := s.D.Id()
	request.LogAnalyticsObjectCollectionRuleId = &tmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.DeleteLogAnalyticsObjectCollectionRule(context.Background(), request)
	return err
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) SetData() error {

	logAnalyticsObjectCollectionRuleId, namespace, err := parseLogAnalyticsObjectCollectionRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(logAnalyticsObjectCollectionRuleId)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CharEncoding != nil {
		s.D.Set("char_encoding", *s.Res.CharEncoding)
	}

	s.D.Set("collection_type", s.Res.CollectionType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EntityId != nil {
		s.D.Set("entity_id", *s.Res.EntityId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.IsForceHistoricCollection != nil {
		s.D.Set("is_force_historic_collection", *s.Res.IsForceHistoricCollection)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogGroupId != nil {
		s.D.Set("log_group_id", *s.Res.LogGroupId)
	}

	if s.Res.LogSet != nil {
		s.D.Set("log_set", *s.Res.LogSet)
	}

	if s.Res.LogSetExtRegex != nil {
		s.D.Set("log_set_ext_regex", *s.Res.LogSetExtRegex)
	}

	s.D.Set("log_set_key", s.Res.LogSetKey)

	if s.Res.LogSourceName != nil {
		s.D.Set("log_source_name", *s.Res.LogSourceName)
	}

	s.D.Set("log_type", s.Res.LogType)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("object_name_filters", s.Res.ObjectNameFilters)

	if s.Res.OsBucketName != nil {
		s.D.Set("os_bucket_name", *s.Res.OsBucketName)
	}

	if s.Res.OsNamespace != nil {
		s.D.Set("os_namespace", *s.Res.OsNamespace)
	}

	if s.Res.Overrides != nil {
		s.D.Set("overrides", propertyOverridesToMap(s.Res.Overrides))
	} else {
		s.D.Set("overrides", nil)
	}

	if s.Res.PollSince != nil {
		s.D.Set("poll_since", *s.Res.PollSince)
	}

	if s.Res.PollTill != nil {
		s.D.Set("poll_till", *s.Res.PollTill)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	return nil
}

func GetLogAnalyticsObjectCollectionRuleCompositeId(logAnalyticsObjectCollectionRuleId string, namespace string) string {
	logAnalyticsObjectCollectionRuleId = url.PathEscape(logAnalyticsObjectCollectionRuleId)
	namespace = url.PathEscape(namespace)
	compositeId := "namespaces/" + namespace + "/logAnalyticsObjectCollectionRules/" + logAnalyticsObjectCollectionRuleId
	return compositeId
}

func parseLogAnalyticsObjectCollectionRuleCompositeId(compositeId string) (logAnalyticsObjectCollectionRuleId string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/logAnalyticsObjectCollectionRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	logAnalyticsObjectCollectionRuleId, _ = url.PathUnescape(parts[3])

	return
}

func LogAnalyticsObjectCollectionRuleSummaryToMap(obj oci_log_analytics.LogAnalyticsObjectCollectionRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["collection_type"] = string(obj.CollectionType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["log_type"] = string(obj.LogType)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["object_name_filters"] = obj.ObjectNameFilters

	if obj.OsBucketName != nil {
		result["os_bucket_name"] = string(*obj.OsBucketName)
	}

	if obj.OsNamespace != nil {
		result["os_namespace"] = string(*obj.OsNamespace)
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

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_log_analytics.ChangeLogAnalyticsObjectCollectionRuleCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LogAnalyticsObjectCollectionRuleId = &idTmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		changeCompartmentRequest.NamespaceName = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.ChangeLogAnalyticsObjectCollectionRuleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleResourceCrud) mapToPropertyOverrides(fieldKeyFormat string) (oci_log_analytics.PropertyOverride, error) {
	result := oci_log_analytics.PropertyOverride{}

	if matchType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "match_type")); ok {
		tmp := matchType.(string)
		result.MatchType = &tmp
	}
	if matchValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "match_value")); ok {
		tmp := matchValue.(string)
		result.MatchValue = &tmp
	}
	if propertyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "property_name")); ok {
		tmp := propertyName.(string)
		result.PropertyName = &tmp
	}
	if propertyValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "property_value")); ok {
		tmp := propertyValue.(string)
		result.PropertyValue = &tmp
	}

	return result, nil
}

func propertyOverridesToMap(obj map[string][]oci_log_analytics.PropertyOverride) []map[string]string {
	overrides := make([]map[string]string, len(obj["items"]))
	i := 0
	for _, item := range obj["items"] {
		overrides[i] = propertyOverrideToMap(item)
		i++
	}

	return overrides
}

func propertyOverrideToMap(obj oci_log_analytics.PropertyOverride) map[string]string {
	result := make(map[string]string)

	result["match_type"] = *obj.MatchType
	result["match_value"] = *obj.MatchValue
	result["property_name"] = *obj.PropertyName
	result["property_value"] = *obj.PropertyValue

	return result
}
