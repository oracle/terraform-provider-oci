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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceAssociationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsNamespaceAssociation,
		Read:     readLogAnalyticsNamespaceAssociation,
		Update:   updateLogAnalyticsNamespaceAssociation,
		Delete:   deleteLogAnalyticsNamespaceAssociation,
		Schema: map[string]*schema.Schema{
			// Required
			"entity_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"log_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"association_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"patterns": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									// Computed
									"effective_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"is_from_republish": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"agent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"agent_entity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_type_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_type_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"failure_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_group_compartment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"retry_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"source_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_type_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_attempted": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsNamespaceAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsNamespaceAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsNamespaceAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsNamespaceAssociationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.LogAnalyticsAssociation
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) ID() string {
	return GetNamespaceAssociationCompositeId(s.D.Get("namespace").(string), s.D.Get("compartment_id").(string), s.D.Get("entity_id").(string), s.D.Get("source_name").(string))
}

func GetNamespaceAssociationCompositeId(namespace string, compartmentId string, entityId string, sourceName string) string {
	namespace = url.PathEscape(namespace)
	entityId = url.PathEscape(entityId)
	sourceName = url.PathEscape(sourceName)
	compartmentId = url.PathEscape(compartmentId)
	compositeId := fmt.Sprintf("namespaces/%s/associations/%s/%s/%s", namespace, compartmentId, entityId, sourceName)
	return compositeId
}

func parseNamespaceAssociationCompositeId(compositeId string) (namespace string, compartmentId string, entityId string, sourceName string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/associations/.*/.*/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	compartmentId, _ = url.PathUnescape(parts[3])
	entityId, _ = url.PathUnescape(parts[4])
	sourceName, _ = url.PathUnescape(parts[5])

	return
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) Create() error {
	return s.Update()
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) Get() error {
	request := oci_log_analytics.ListSourceAssociationsRequest{}
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	namespace, compartmentId, entityId, sourceName, err := parseNamespaceAssociationCompositeId(s.D.Id())
	if err == nil {
		request.NamespaceName = &namespace
		request.CompartmentId = &compartmentId
		request.EntityId = &entityId
		request.SourceName = &sourceName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		return err
	}

	response, err := s.Client.ListSourceAssociations(context.Background(), request)

	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsAssociationCollection.Items[0]
	return nil
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) Update() error {
	request := oci_log_analytics.UpsertAssociationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isFromRepublish, ok := s.D.GetOkExists("is_from_republish"); ok {
		tmp := isFromRepublish.(bool)
		request.IsFromRepublish = &tmp
	}

	var namespaceName string
	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		request.NamespaceName = &namespaceName
	}

	upsertAssocItem, err := s.makeUpsertLogAnalyticsAssociation()
	if err != nil {
		return err
	}
	request.Items = []oci_log_analytics.UpsertLogAnalyticsAssociation{upsertAssocItem}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpsertAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNamespaceAssociationFromWorkRequest(&namespaceName, workId, false, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeCreateAssociations, s.D.Timeout(schema.TimeoutCreate))
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) Delete() error {
	request := oci_log_analytics.DeleteAssociationsRequest{}

	var namespaceName string
	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		request.NamespaceName = &namespaceName
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	deleteAssocItem, err := s.makeDeleteLogAnalyticsAssociation()
	if err != nil {
		return err
	}
	request.Items = []oci_log_analytics.DeleteLogAnalyticsAssociation{deleteAssocItem}

	response, err := s.Client.DeleteAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNamespaceAssociationFromWorkRequest(&namespaceName, workId, true, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeDeleteAssociations, s.D.Timeout(schema.TimeoutCreate))
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) getNamespaceAssociationFromWorkRequest(namespaceName *string, workId *string, isDelete bool, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := NamespaceAssociationWaitForWorkRequest(namespaceName, workId, "namespace",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	if isDelete {
		return nil
	}

	var compartmentId *string
	if compartmentIdfromD, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentIdfromD.(string)
		compartmentId = &tmp
	}

	var entityId *string
	if entityIdfromD, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityIdfromD.(string)
		entityId = &tmp
	}

	var sourceName *string
	if sourceNamefromD, ok := s.D.GetOkExists("source_name"); ok {
		tmp := sourceNamefromD.(string)
		sourceName = &tmp
	}

	s.D.SetId(GetNamespaceAssociationCompositeId(*namespaceName, *compartmentId, *entityId, *sourceName))
	return s.Get()
}

func NamespaceAssociationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "log_analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_log_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func NamespaceAssociationWaitForWorkRequest(namespaceName *string, wId *string, entityType string, action oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_log_analytics.LogAnalyticsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "log_analytics")
	retryPolicy.ShouldRetryOperation = NamespaceAssociationWorkRequestShouldRetryFunc(timeout)

	response := oci_log_analytics.GetConfigWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateInProgress),
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateAccepted),
		},
		Target: []string{
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateSucceeded),
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetConfigWorkRequest(context.Background(),
				oci_log_analytics.GetConfigWorkRequestRequest{
					NamespaceName: namespaceName,
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.LogAnalyticsConfigWorkRequest
			return wr, string(wr.LifecycleState), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	// The work request may have failed, check for errors if identifier is not found or work failed or got cancelled
	if response.LogAnalyticsConfigWorkRequest.LifecycleState == oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateFailed {
		return nil, getErrorFromLogAnalyticsNamespaceAssociationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return nil, nil
}

func getErrorFromLogAnalyticsNamespaceAssociationWorkRequest(client *oci_log_analytics.LogAnalyticsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum) error {
	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s", *workId, entityType, action)

	return workRequestErr
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) SetData() error {
	namespace, compartmentId, entityId, sourceName, err := parseNamespaceAssociationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("namespace", namespace)
		s.D.Set("compartment_id", compartmentId)
		s.D.Set("entity_id", entityId)
		s.D.Set("source_name", sourceName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}
	if s.Res.AssociationProperties != nil {
		var associationProperties []interface{}
		for _, associationProperty := range s.Res.AssociationProperties {
			associationProperties = append(associationProperties, AssociationPropertyToMap(associationProperty))
		}
		s.D.Set("association_properties", associationProperties)
	}
	if s.Res.EntityName != nil {
		s.D.Set("entity_name", *s.Res.EntityName)
	}
	if s.Res.EntityTypeName != nil {
		s.D.Set("entity_type_name", *s.Res.EntityTypeName)
	}
	if s.Res.Host != nil {
		s.D.Set("host", *s.Res.Host)
	}
	s.D.Set("lifecycle_state", s.Res.LifeCycleState)
	if s.Res.LogGroupId != nil {
		s.D.Set("log_group_id", *s.Res.LogGroupId)
	}
	if s.Res.LogGroupName != nil {
		s.D.Set("log_group_name", *s.Res.LogGroupName)
	}
	if s.Res.LogGroupCompartment != nil {
		s.D.Set("log_group_compartment", *s.Res.LogGroupCompartment)
	}
	if s.Res.RetryCount != nil {
		s.D.Set("retry_count", *s.Res.RetryCount)
	}
	if s.Res.SourceDisplayName != nil {
		s.D.Set("source_display_name", *s.Res.SourceDisplayName)
	}
	if s.Res.SourceTypeName != nil {
		s.D.Set("source_type_name", *s.Res.SourceTypeName)
	}
	if s.Res.TimeLastAttempted != nil {
		s.D.Set("time_last_attempted", s.Res.TimeLastAttempted.String())
	}
	if s.Res.AgentEntityName != nil {
		s.D.Set("agent_entity_name", *s.Res.AgentEntityName)
	}
	if s.Res.EntityTypeDisplayName != nil {
		s.D.Set("entity_type_display_name", *s.Res.EntityTypeDisplayName)
	}
	if s.Res.FailureMessage != nil {
		s.D.Set("failure_message", *s.Res.FailureMessage)
	}

	return nil
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) mapToAssociationProperty(fieldKeyFormat string) (oci_log_analytics.AssociationProperty, error) {
	result := oci_log_analytics.AssociationProperty{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if patterns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patterns")); ok {
		interfaces := patterns.([]interface{})
		tmp := make([]oci_log_analytics.PatternOverride, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patterns"), stateDataIndex)
			converted, err := s.mapToPatternOverride(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patterns")) {
			result.Patterns = tmp
		}
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AssociationPropertyToMap(obj oci_log_analytics.AssociationProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	patterns := []interface{}{}
	if obj.Patterns != nil {
		for _, item := range obj.Patterns {
			patterns = append(patterns, LogAnalyticsAssociationsPatternOverrideToMap(item))
		}
		result["patterns"] = patterns
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) mapToPatternOverride(fieldKeyFormat string) (oci_log_analytics.PatternOverride, error) {
	result := oci_log_analytics.PatternOverride{}

	if effectiveLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "effective_level")); ok {
		if effectiveLevel != "" {
			fmt.Println("Inside here. Checking.")
			tmp := effectiveLevel.(string)
			result.EffectiveLevel = &tmp
		}
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func LogAnalyticsAssociationsPatternOverrideToMap(obj oci_log_analytics.PatternOverride) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EffectiveLevel != nil {
		result["effective_level"] = string(*obj.EffectiveLevel)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) makeUpsertLogAnalyticsAssociation() (oci_log_analytics.UpsertLogAnalyticsAssociation, error) {
	result := oci_log_analytics.UpsertLogAnalyticsAssociation{}

	if associationProperties, ok := s.D.GetOkExists("association_properties"); ok {
		interfaces := associationProperties.([]interface{})
		tmp := make([]oci_log_analytics.AssociationProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "association_properties", stateDataIndex)
			converted, err := s.mapToAssociationProperty(fieldKeyFormat)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("association_properties") {
			result.AssociationProperties = tmp
		}
	}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		result.EntityId = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if sourceName, ok := s.D.GetOkExists("source_name"); ok {
		tmp := sourceName.(string)
		result.SourceName = &tmp
	}

	return result, nil
}

func (s *LogAnalyticsNamespaceAssociationResourceCrud) makeDeleteLogAnalyticsAssociation() (oci_log_analytics.DeleteLogAnalyticsAssociation, error) {
	result := oci_log_analytics.DeleteLogAnalyticsAssociation{}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		result.EntityId = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if sourceName, ok := s.D.GetOkExists("source_name"); ok {
		tmp := sourceName.(string)
		result.SourceName = &tmp
	}

	return result, nil
}
