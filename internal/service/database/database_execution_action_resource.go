// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExecutionActionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExecutionAction,
		Read:     readDatabaseExecutionAction,
		Update:   updateDatabaseExecutionAction,
		Delete:   deleteDatabaseExecutionAction,
		Schema: map[string]*schema.Schema{
			// Required
			"action_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"execution_window_id": {
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
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"total_time_taken_in_mins": {
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
			"compartment_id": {
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

			// Computed
			"description": {
				Type:     schema.TypeString,
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
			"execution_action_order": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_substate": {
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
			"total_time_taken_in_mins": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDatabaseExecutionAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExecutionAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseExecutionAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExecutionAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExecutionActionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExecutionAction
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExecutionActionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExecutionActionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExecutionActionLifecycleStateScheduled),
		string(oci_database.ExecutionActionLifecycleStateInProgress),
	}
}

func (s *DatabaseExecutionActionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExecutionActionLifecycleStateSucceeded),
	}
}

func (s *DatabaseExecutionActionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseExecutionActionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExecutionActionLifecycleStateDeleted),
	}
}

func (s *DatabaseExecutionActionResourceCrud) Create() error {
	request := oci_database.CreateExecutionActionRequest{}

	if actionMembers, ok := s.D.GetOkExists("action_members"); ok {
		interfaces := actionMembers.([]interface{})
		tmp := make([]oci_database.ExecutionActionMember, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_members", stateDataIndex)
			converted, err := s.mapToExecutionActionMember(fieldKeyFormat)
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
		request.ActionType = oci_database.CreateExecutionActionDetailsActionTypeEnum(actionType.(string))
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

	if executionWindowId, ok := s.D.GetOkExists("execution_window_id"); ok {
		tmp := executionWindowId.(string)
		request.ExecutionWindowId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExecutionAction(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExecutionAction

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "executionaction", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseExecutionActionResourceCrud) Get() error {
	request := oci_database.GetExecutionActionRequest{}

	tmp := s.D.Id()
	request.ExecutionActionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExecutionAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExecutionAction
	return nil
}

func (s *DatabaseExecutionActionResourceCrud) Update() error {
	request := oci_database.UpdateExecutionActionRequest{}

	if actionMembers, ok := s.D.GetOkExists("action_members"); ok {
		interfaces := actionMembers.([]interface{})
		tmp := make([]oci_database.ExecutionActionMember, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_members", stateDataIndex)
			converted, err := s.mapToExecutionActionMember(fieldKeyFormat)
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

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.ExecutionActionId = &tmp

	if executionWindowId, ok := s.D.GetOkExists("execution_window_id"); ok {
		tmp := executionWindowId.(string)
		request.ExecutionWindowId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExecutionAction(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "executionaction", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseExecutionActionResourceCrud) Delete() error {
	request := oci_database.DeleteExecutionActionRequest{}

	tmp := s.D.Id()
	request.ExecutionActionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExecutionAction(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "executionaction", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExecutionActionResourceCrud) SetData() error {
	actionMembers := []interface{}{}
	for _, item := range s.Res.ActionMembers {
		actionMembers = append(actionMembers, ExecutionActionMemberToMap(item))
	}
	s.D.Set("action_members", actionMembers)

	s.D.Set("action_params", s.Res.ActionParams)

	s.D.Set("action_type", s.Res.ActionType)

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

	if s.Res.EstimatedTimeInMins != nil {
		s.D.Set("estimated_time_in_mins", *s.Res.EstimatedTimeInMins)
	}

	if s.Res.ExecutionActionOrder != nil {
		s.D.Set("execution_action_order", *s.Res.ExecutionActionOrder)
	}

	if s.Res.ExecutionWindowId != nil {
		s.D.Set("execution_window_id", *s.Res.ExecutionWindowId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalTimeTakenInMins != nil {
		s.D.Set("total_time_taken_in_mins", *s.Res.TotalTimeTakenInMins)
	}

	return nil
}

func (s *DatabaseExecutionActionResourceCrud) mapToExecutionActionMember(fieldKeyFormat string) (oci_database.ExecutionActionMember, error) {
	result := oci_database.ExecutionActionMember{}

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

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		tmp := status.(string)
		result.Status = &tmp
	}

	if totalTimeTakenInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "total_time_taken_in_mins")); ok {
		tmp := totalTimeTakenInMins.(int)
		result.TotalTimeTakenInMins = &tmp
	}

	return result, nil
}

func ExecutionActionMemberToMap(obj oci_database.ExecutionActionMember) map[string]interface{} {
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

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TotalTimeTakenInMins != nil {
		result["total_time_taken_in_mins"] = int(*obj.TotalTimeTakenInMins)
	}

	return result
}
