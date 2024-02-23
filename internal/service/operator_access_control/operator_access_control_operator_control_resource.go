// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_operator_access_control "github.com/oracle/oci-go-sdk/v65/operatoraccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlOperatorControlResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOperatorAccessControlOperatorControl,
		Read:     readOperatorAccessControlOperatorControl,
		Update:   updateOperatorAccessControlOperatorControl,
		Delete:   deleteOperatorAccessControlOperatorControl,
		Schema: map[string]*schema.Schema{
			// Required
			"approver_groups_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_fully_pre_approved": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"operator_control_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"approvers_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"email_id_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"number_of_approvers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pre_approved_op_action_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"approval_required_op_action_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_default_operator_control": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_modified_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_creation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_modification": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOperatorAccessControlOperatorControl(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlClient()

	return tfresource.CreateResource(d, sync)
}

func readOperatorAccessControlOperatorControl(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlClient()

	return tfresource.ReadResource(sync)
}

func updateOperatorAccessControlOperatorControl(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOperatorAccessControlOperatorControl(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OperatorAccessControlOperatorControlResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_operator_access_control.OperatorControlClient
	Res                    *oci_operator_access_control.OperatorControl
	DisableNotFoundRetries bool
}

func (s *OperatorAccessControlOperatorControlResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OperatorAccessControlOperatorControlResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *OperatorAccessControlOperatorControlResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_operator_access_control.OperatorControlLifecycleStatesCreated),
		string(oci_operator_access_control.OperatorControlLifecycleStatesAssigned),
		string(oci_operator_access_control.OperatorControlLifecycleStatesUnassigned),
	}
}

func (s *OperatorAccessControlOperatorControlResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *OperatorAccessControlOperatorControlResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_operator_access_control.OperatorControlLifecycleStatesDeleted),
	}
}

func (s *OperatorAccessControlOperatorControlResourceCrud) Create() error {
	request := oci_operator_access_control.CreateOperatorControlRequest{}

	if approverGroupsList, ok := s.D.GetOkExists("approver_groups_list"); ok {
		interfaces := approverGroupsList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("approver_groups_list") {
			request.ApproverGroupsList = tmp
		}
	}

	if approversList, ok := s.D.GetOkExists("approvers_list"); ok {
		interfaces := approversList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("approvers_list") {
			request.ApproversList = tmp
		}
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

	if emailIdList, ok := s.D.GetOkExists("email_id_list"); ok {
		interfaces := emailIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("email_id_list") {
			request.EmailIdList = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isFullyPreApproved, ok := s.D.GetOkExists("is_fully_pre_approved"); ok {
		tmp := isFullyPreApproved.(bool)
		request.IsFullyPreApproved = &tmp
	}

	if numberOfApprovers, ok := s.D.GetOkExists("number_of_approvers"); ok {
		tmp := numberOfApprovers.(int)
		request.NumberOfApprovers = &tmp
	}

	if operatorControlName, ok := s.D.GetOkExists("operator_control_name"); ok {
		tmp := operatorControlName.(string)
		request.OperatorControlName = &tmp
	}

	if preApprovedOpActionList, ok := s.D.GetOkExists("pre_approved_op_action_list"); ok {
		interfaces := preApprovedOpActionList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("pre_approved_op_action_list") {
			request.PreApprovedOpActionList = tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_operator_access_control.ResourceTypesEnum(resourceType.(string))
	}

	if systemMessage, ok := s.D.GetOkExists("system_message"); ok {
		tmp := systemMessage.(string)
		request.SystemMessage = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	response, err := s.Client.CreateOperatorControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperatorControl
	return nil
}

func (s *OperatorAccessControlOperatorControlResourceCrud) Get() error {
	request := oci_operator_access_control.GetOperatorControlRequest{}

	tmp := s.D.Id()
	request.OperatorControlId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	response, err := s.Client.GetOperatorControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperatorControl
	return nil
}

func (s *OperatorAccessControlOperatorControlResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_operator_access_control.UpdateOperatorControlRequest{}

	if approverGroupsList, ok := s.D.GetOkExists("approver_groups_list"); ok {
		interfaces := approverGroupsList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("approver_groups_list") {
			request.ApproverGroupsList = tmp
		}
	}

	if approversList, ok := s.D.GetOkExists("approvers_list"); ok {
		interfaces := approversList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("approvers_list") {
			request.ApproversList = tmp
		}
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

	if emailIdList, ok := s.D.GetOkExists("email_id_list"); ok {
		interfaces := emailIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("email_id_list") {
			request.EmailIdList = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isFullyPreApproved, ok := s.D.GetOkExists("is_fully_pre_approved"); ok {
		tmp := isFullyPreApproved.(bool)
		request.IsFullyPreApproved = &tmp
	}

	if numberOfApprovers, ok := s.D.GetOkExists("number_of_approvers"); ok {
		tmp := numberOfApprovers.(int)
		request.NumberOfApprovers = &tmp
	}

	tmp := s.D.Id()
	request.OperatorControlId = &tmp

	if operatorControlName, ok := s.D.GetOkExists("operator_control_name"); ok {
		tmp := operatorControlName.(string)
		request.OperatorControlName = &tmp
	}

	if preApprovedOpActionList, ok := s.D.GetOkExists("pre_approved_op_action_list"); ok {
		interfaces := preApprovedOpActionList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("pre_approved_op_action_list") {
			request.PreApprovedOpActionList = tmp
		}
	}

	if systemMessage, ok := s.D.GetOkExists("system_message"); ok {
		tmp := systemMessage.(string)
		request.SystemMessage = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	response, err := s.Client.UpdateOperatorControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperatorControl
	return nil
}

func (s *OperatorAccessControlOperatorControlResourceCrud) Delete() error {
	request := oci_operator_access_control.DeleteOperatorControlRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.OperatorControlId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	_, err := s.Client.DeleteOperatorControl(context.Background(), request)
	return err
}

func (s *OperatorAccessControlOperatorControlResourceCrud) SetData() error {
	s.D.Set("approval_required_op_action_list", s.Res.ApprovalRequiredOpActionList)

	s.D.Set("approver_groups_list", s.Res.ApproverGroupsList)

	s.D.Set("approvers_list", s.Res.ApproversList)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("email_id_list", s.Res.EmailIdList)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsDefaultOperatorControl != nil {
		s.D.Set("is_default_operator_control", *s.Res.IsDefaultOperatorControl)
	}

	if s.Res.IsFullyPreApproved != nil {
		s.D.Set("is_fully_pre_approved", *s.Res.IsFullyPreApproved)
	}

	if s.Res.LastModifiedInfo != nil {
		s.D.Set("last_modified_info", *s.Res.LastModifiedInfo)
	}

	if s.Res.NumberOfApprovers != nil {
		s.D.Set("number_of_approvers", *s.Res.NumberOfApprovers)
	}

	if s.Res.OperatorControlName != nil {
		s.D.Set("operator_control_name", *s.Res.OperatorControlName)
	}

	s.D.Set("pre_approved_op_action_list", s.Res.PreApprovedOpActionList)

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemMessage != nil {
		s.D.Set("system_message", *s.Res.SystemMessage)
	}

	if s.Res.TimeOfCreation != nil {
		s.D.Set("time_of_creation", s.Res.TimeOfCreation.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.TimeOfModification != nil {
		s.D.Set("time_of_modification", s.Res.TimeOfModification.String())
	}

	return nil
}

func OperatorControlSummaryToMap(obj oci_operator_access_control.OperatorControlSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsFullyPreApproved != nil {
		result["is_fully_pre_approved"] = bool(*obj.IsFullyPreApproved)
	}

	if obj.NumberOfApprovers != nil {
		result["number_of_approvers"] = int(*obj.NumberOfApprovers)
	}

	if obj.OperatorControlName != nil {
		result["operator_control_name"] = string(*obj.OperatorControlName)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeOfCreation != nil {
		result["time_of_creation"] = obj.TimeOfCreation.String()
	}

	if obj.TimeOfDeletion != nil {
		result["time_of_deletion"] = obj.TimeOfDeletion.String()
	}

	if obj.TimeOfModification != nil {
		result["time_of_modification"] = obj.TimeOfModification.String()
	}

	return result
}

func (s *OperatorAccessControlOperatorControlResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_operator_access_control.ChangeOperatorControlCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OperatorControlId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	_, err := s.Client.ChangeOperatorControlCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
