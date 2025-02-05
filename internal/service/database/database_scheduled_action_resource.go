// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseScheduledActionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseScheduledAction,
		Read:     readDatabaseScheduledAction,
		Update:   updateDatabaseScheduledAction,
		Delete:   deleteDatabaseScheduledAction,
		Schema: map[string]*schema.Schema{
			// Required
			"action_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scheduling_plan_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scheduling_window_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"action_members": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"member_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"member_order": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"estimated_time_in_mins": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"action_params": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"action_order": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"estimated_time_in_mins": {
				Type:     schema.TypeInt,
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

func createDatabaseScheduledAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseScheduledActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseScheduledAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseScheduledActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseScheduledAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseScheduledActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseScheduledAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseScheduledActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseScheduledActionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ScheduledAction
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseScheduledActionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseScheduledActionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ScheduledActionLifecycleStateCreating),
	}
}

func (s *DatabaseScheduledActionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ScheduledActionLifecycleStateNeedsAttention),
		string(oci_database.ScheduledActionLifecycleStateAvailable),
	}
}

func (s *DatabaseScheduledActionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ScheduledActionLifecycleStateDeleting),
	}
}

func (s *DatabaseScheduledActionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ScheduledActionLifecycleStateDeleted),
	}
}

func (s *DatabaseScheduledActionResourceCrud) Create() error {
	request := oci_database.CreateScheduledActionRequest{}

	if actionMembers, ok := s.D.GetOkExists("action_members"); ok {
		interfaces := actionMembers.([]interface{})
		tmp := make([]oci_database.ActionMember, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_members", stateDataIndex)
			converted, err := s.mapToActionMember(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("action_members") {
			request.ActionMembers = tmp
		}
	}

	if actionParams, ok := s.D.GetOkExists("action_params"); ok {
		request.ActionParams = tfresource.ObjectMapToStringMap(actionParams.(map[string]interface{}))
	}

	if actionType, ok := s.D.GetOkExists("action_type"); ok {
		request.ActionType = oci_database.CreateScheduledActionDetailsActionTypeEnum(actionType.(string))
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if schedulingPlanId, ok := s.D.GetOkExists("scheduling_plan_id"); ok {
		tmp := schedulingPlanId.(string)
		request.SchedulingPlanId = &tmp
	}

	if schedulingWindowId, ok := s.D.GetOkExists("scheduling_window_id"); ok {
		tmp := schedulingWindowId.(string)
		request.SchedulingWindowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateScheduledAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledAction
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	//has to wait some time, otherwise subsequent querying will fail
	time.Sleep(time.Second * 5)
	return s.Get()
}

func (s *DatabaseScheduledActionResourceCrud) Get() error {
	request := oci_database.GetScheduledActionRequest{}

	tmp := s.D.Id()
	request.ScheduledActionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetScheduledAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledAction
	return nil
}

func (s *DatabaseScheduledActionResourceCrud) Update() error {
	request := oci_database.UpdateScheduledActionRequest{}

	if actionMembers, ok := s.D.GetOkExists("action_members"); ok {
		interfaces := actionMembers.([]interface{})
		tmp := make([]oci_database.ActionMember, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_members", stateDataIndex)
			converted, err := s.mapToActionMember(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("action_members") {
			request.ActionMembers = tmp
		}
	}

	if actionParams, ok := s.D.GetOkExists("action_params"); ok {
		request.ActionParams = tfresource.ObjectMapToStringMap(actionParams.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ScheduledActionId = &tmp

	if schedulingWindowId, ok := s.D.GetOkExists("scheduling_window_id"); ok {
		tmp := schedulingWindowId.(string)
		request.SchedulingWindowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateScheduledAction(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "scheduledaction", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseScheduledActionResourceCrud) Delete() error {
	request := oci_database.DeleteScheduledActionRequest{}

	tmp := s.D.Id()
	request.ScheduledActionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteScheduledAction(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "scheduledaction", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseScheduledActionResourceCrud) SetData() error {
	actionMembers := []interface{}{}
	for _, item := range s.Res.ActionMembers {
		actionMembers = append(actionMembers, ActionMemberToMap(item))
	}
	s.D.Set("action_members", actionMembers)

	if s.Res.ActionOrder != nil {
		s.D.Set("action_order", *s.Res.ActionOrder)
	}

	s.D.Set("action_params", s.Res.ActionParams)

	s.D.Set("action_type", s.Res.ActionType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EstimatedTimeInMins != nil {
		s.D.Set("estimated_time_in_mins", *s.Res.EstimatedTimeInMins)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SchedulingPlanId != nil {
		s.D.Set("scheduling_plan_id", *s.Res.SchedulingPlanId)
	}

	if s.Res.SchedulingWindowId != nil {
		s.D.Set("scheduling_window_id", *s.Res.SchedulingWindowId)
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

func (s *DatabaseScheduledActionResourceCrud) mapToActionMember(fieldKeyFormat string) (oci_database.ActionMember, error) {
	result := oci_database.ActionMember{}

	if estimatedTimeInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "estimated_time_in_mins")); ok {
		tmp := estimatedTimeInMins.(int)
		result.EstimatedTimeInMins = &tmp
	}

	if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
		tmp := memberId.(string)
		result.MemberId = &tmp
	}

	if memberOrder, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_order")); ok {
		tmp := memberOrder.(int)
		result.MemberOrder = &tmp
	}

	return result, nil
}

func ActionMemberToMap(obj oci_database.ActionMember) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EstimatedTimeInMins != nil {
		result["estimated_time_in_mins"] = int(*obj.EstimatedTimeInMins)
	}

	if obj.MemberId != nil {
		result["member_id"] = string(*obj.MemberId)
	}

	if obj.MemberOrder != nil {
		result["member_order"] = int(*obj.MemberOrder)
	}

	return result
}

func ScheduledActionSummaryToMap(obj oci_database.ScheduledActionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	actionMembers := []interface{}{}
	for _, item := range obj.ActionMembers {
		actionMembers = append(actionMembers, ActionMemberToMap(item))
	}
	result["action_members"] = actionMembers

	if obj.ActionOrder != nil {
		result["action_order"] = int(*obj.ActionOrder)
	}

	result["action_params"] = obj.ActionParams

	result["action_type"] = string(obj.ActionType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EstimatedTimeInMins != nil {
		result["estimated_time_in_mins"] = int(*obj.EstimatedTimeInMins)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SchedulingPlanId != nil {
		result["scheduling_plan_id"] = string(*obj.SchedulingPlanId)
	}

	if obj.SchedulingWindowId != nil {
		result["scheduling_window_id"] = string(*obj.SchedulingWindowId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
