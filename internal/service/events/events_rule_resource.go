// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package events

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_events "github.com/oracle/oci-go-sdk/v56/events"
)

func EventsRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createEventsRule,
		Read:     readEventsRule,
		Update:   updateEventsRule,
		Delete:   deleteEventsRule,
		Schema: map[string]*schema.Schema{
			// Required
			"actions": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"actions": {
							Type:     schema.TypeSet,
							Required: true,
							Set:      actionsHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"FAAS",
											"ONS",
											"OSS",
										}, true),
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"function_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"stream_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"topic_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
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
			"lifecycle_message": {
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

func createEventsRule(d *schema.ResourceData, m interface{}) error {
	sync := &EventsRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EventsClient()

	return tfresource.CreateResource(d, sync)
}

func readEventsRule(d *schema.ResourceData, m interface{}) error {
	sync := &EventsRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EventsClient()

	return tfresource.ReadResource(sync)
}

func updateEventsRule(d *schema.ResourceData, m interface{}) error {
	sync := &EventsRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EventsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteEventsRule(d *schema.ResourceData, m interface{}) error {
	sync := &EventsRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EventsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type EventsRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_events.EventsClient
	Res                    *oci_events.Rule
	DisableNotFoundRetries bool
}

func (s *EventsRuleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *EventsRuleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_events.RuleLifecycleStateCreating),
	}
}

func (s *EventsRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_events.RuleLifecycleStateActive),
		string(oci_events.RuleLifecycleStateInactive),
	}
}

func (s *EventsRuleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_events.RuleLifecycleStateDeleting),
	}
}

func (s *EventsRuleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_events.RuleLifecycleStateDeleted),
	}
}

func (s *EventsRuleResourceCrud) Create() error {
	request := oci_events.CreateRuleRequest{}

	if actions, ok := s.D.GetOkExists("actions"); ok {
		if tmpList := actions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", 0)
			tmp, err := s.mapToActionDetailsList(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Actions = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if condition, ok := s.D.GetOkExists("condition"); ok {
		tmp := condition.(string)
		request.Condition = &tmp
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

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "events")

	response, err := s.Client.CreateRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Rule
	return nil
}

func (s *EventsRuleResourceCrud) Get() error {
	request := oci_events.GetRuleRequest{}

	tmp := s.D.Id()
	request.RuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "events")

	response, err := s.Client.GetRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Rule
	return nil
}

func (s *EventsRuleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_events.UpdateRuleRequest{}

	if actions, ok := s.D.GetOkExists("actions"); ok {
		if tmpList := actions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", 0)
			tmp, err := s.mapToActionDetailsList(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Actions = &tmp
		}
	}

	if condition, ok := s.D.GetOkExists("condition"); ok {
		tmp := condition.(string)
		request.Condition = &tmp
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

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	tmp := s.D.Id()
	request.RuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "events")

	response, err := s.Client.UpdateRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Rule
	return nil
}

func (s *EventsRuleResourceCrud) Delete() error {
	request := oci_events.DeleteRuleRequest{}

	tmp := s.D.Id()
	request.RuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "events")

	_, err := s.Client.DeleteRule(context.Background(), request)
	return err
}

func (s *EventsRuleResourceCrud) SetData() error {
	if s.Res.Actions != nil {
		s.D.Set("actions", []interface{}{ActionListToMap(s.Res.Actions, false)})
	} else {
		s.D.Set("actions", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Condition != nil {
		s.D.Set("condition", *s.Res.Condition)
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

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecycleMessage != nil {
		s.D.Set("lifecycle_message", *s.Res.LifecycleMessage)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *EventsRuleResourceCrud) mapToActionDetails(fieldKeyFormat string) (oci_events.ActionDetails, error) {
	var baseObject oci_events.ActionDetails
	//discriminator
	actionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_type"))
	var actionType string
	if ok {
		actionType = actionTypeRaw.(string)
	} else {
		actionType = "" // default value
	}
	switch strings.ToLower(actionType) {
	case strings.ToLower("FAAS"):
		details := oci_events.CreateFaaSActionDetails{}
		if functionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "function_id")); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("ONS"):
		details := oci_events.CreateNotificationServiceActionDetails{}
		if topicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "topic_id")); ok {
			tmp := topicId.(string)
			details.TopicId = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("OSS"):
		details := oci_events.CreateStreamingServiceActionDetails{}
		if streamId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_id")); ok {
			tmp := streamId.(string)
			details.StreamId = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown action_type '%v' was specified", actionType)
	}
	return baseObject, nil
}

func EventsActionToMap(obj oci_events.Action) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_events.FaaSAction:
		result["action_type"] = "FAAS"

		if v.FunctionId != nil {
			result["function_id"] = string(*v.FunctionId)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}

		if v.LifecycleMessage != nil {
			result["lifecycle_message"] = string(*v.LifecycleMessage)
		}

		result["state"] = string(v.LifecycleState)
	case oci_events.NotificationServiceAction:
		result["action_type"] = "ONS"

		if v.TopicId != nil {
			result["topic_id"] = string(*v.TopicId)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}

		if v.LifecycleMessage != nil {
			result["lifecycle_message"] = string(*v.LifecycleMessage)
		}

		result["state"] = string(v.LifecycleState)
	case oci_events.StreamingServiceAction:
		result["action_type"] = "OSS"

		if v.StreamId != nil {
			result["stream_id"] = string(*v.StreamId)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}

		if v.LifecycleMessage != nil {
			result["lifecycle_message"] = string(*v.LifecycleMessage)
		}

		result["state"] = string(v.LifecycleState)
	default:
		log.Printf("[WARN] Received 'action_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *EventsRuleResourceCrud) mapToActionDetailsList(fieldKeyFormat string) (oci_events.ActionDetailsList, error) {
	result := oci_events.ActionDetailsList{}

	if actions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "actions")); ok {
		set := actions.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_events.ActionDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := actionsHashCodeForSets(interfaces[i])
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "actions"), stateDataIndex)
			converted, err := s.mapToActionDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "actions")) {
			result.Actions = tmp
		}
	}

	return result, nil
}

func ActionListToMap(obj *oci_events.ActionList, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	actions := []interface{}{}
	for _, item := range obj.Actions {
		if item.GetLifecycleState() != oci_events.ActionLifecycleStateDeleted {
			actions = append(actions, EventsActionToMap(item))
		}
	}
	if datasource {
		result["actions"] = actions
	} else {
		result["actions"] = schema.NewSet(actionsHashCodeForSets, actions)
	}

	return result
}

func actionsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if actionType, ok := m["action_type"]; ok && actionType != "" {
		buf.WriteString(fmt.Sprintf("%v-", actionType))
	}
	if description, ok := m["description"]; ok && description != "" {
		buf.WriteString(fmt.Sprintf("%v-", description))
	}
	if functionId, ok := m["function_id"]; ok && functionId != "" {
		buf.WriteString(fmt.Sprintf("%v-", functionId))
	}
	if isEnabled, ok := m["is_enabled"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", isEnabled))
	}
	if streamId, ok := m["stream_id"]; ok && streamId != "" {
		buf.WriteString(fmt.Sprintf("%v-", streamId))
	}
	if topicId, ok := m["topic_id"]; ok && topicId != "" {
		buf.WriteString(fmt.Sprintf("%v-", topicId))
	}
	return hashcode.String(buf.String())
}
func (s *EventsRuleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_events.ChangeRuleCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RuleId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "events")

	_, err := s.Client.ChangeRuleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
