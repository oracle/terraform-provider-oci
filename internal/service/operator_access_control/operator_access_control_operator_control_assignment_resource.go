// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v58/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func OperatorAccessControlOperatorControlAssignmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOperatorAccessControlOperatorControlAssignment,
		Read:     readOperatorAccessControlOperatorControlAssignment,
		Update:   updateOperatorAccessControlOperatorControlAssignment,
		Delete:   deleteOperatorAccessControlOperatorControlAssignment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_enforced_always": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"operator_control_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"is_auto_approve_during_maintenance": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_log_forwarded": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"remote_syslog_server_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_syslog_server_ca_cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_syslog_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"time_assignment_from": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},
			"time_assignment_to": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},

			// Computed
			"assigner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"detachment_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_code": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"error_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_assignment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unassigner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOperatorAccessControlOperatorControlAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlAssignmentClient()

	return tfresource.CreateResource(d, sync)
}

func readOperatorAccessControlOperatorControlAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlAssignmentClient()

	return tfresource.ReadResource(sync)
}

func updateOperatorAccessControlOperatorControlAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlAssignmentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOperatorAccessControlOperatorControlAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlAssignmentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OperatorAccessControlOperatorControlAssignmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_operator_access_control.OperatorControlAssignmentClient
	Res                    *oci_operator_access_control.OperatorControlAssignment
	DisableNotFoundRetries bool
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) CreatedPending() []string {
	return []string{string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesCreated)}
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesApplied),
	}
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesDeleting),
	}
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesDeleted),
	}
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) Create() error {
	request := oci_operator_access_control.CreateOperatorControlAssignmentRequest{}

	if comment, ok := s.D.GetOkExists("comment"); ok {
		tmp := comment.(string)
		request.Comment = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoApproveDuringMaintenance, ok := s.D.GetOkExists("is_auto_approve_during_maintenance"); ok {
		tmp := isAutoApproveDuringMaintenance.(bool)
		request.IsAutoApproveDuringMaintenance = &tmp
	}

	if isEnforcedAlways, ok := s.D.GetOkExists("is_enforced_always"); ok {
		tmp := isEnforcedAlways.(bool)
		request.IsEnforcedAlways = &tmp
	}

	if isLogForwarded, ok := s.D.GetOkExists("is_log_forwarded"); ok {
		tmp := isLogForwarded.(bool)
		request.IsLogForwarded = &tmp
	}

	if operatorControlId, ok := s.D.GetOkExists("operator_control_id"); ok {
		tmp := operatorControlId.(string)
		request.OperatorControlId = &tmp
	}

	if remoteSyslogServerAddress, ok := s.D.GetOkExists("remote_syslog_server_address"); ok {
		tmp := remoteSyslogServerAddress.(string)
		request.RemoteSyslogServerAddress = &tmp
	}

	if remoteSyslogServerCACert, ok := s.D.GetOkExists("remote_syslog_server_ca_cert"); ok {
		tmp := remoteSyslogServerCACert.(string)
		request.RemoteSyslogServerCACert = &tmp
	}

	if remoteSyslogServerPort, ok := s.D.GetOkExists("remote_syslog_server_port"); ok {
		tmp := remoteSyslogServerPort.(int)
		request.RemoteSyslogServerPort = &tmp
	}

	if resourceCompartmentId, ok := s.D.GetOkExists("resource_compartment_id"); ok {
		tmp := resourceCompartmentId.(string)
		request.ResourceCompartmentId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_operator_access_control.ResourceTypesEnum(resourceType.(string))
	}

	if timeAssignmentFrom, ok := s.D.GetOkExists("time_assignment_from"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAssignmentFrom.(string))
		if err != nil {
			return err
		}
		request.TimeAssignmentFrom = &oci_common.SDKTime{Time: tmp}
	}

	if timeAssignmentTo, ok := s.D.GetOkExists("time_assignment_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAssignmentTo.(string))
		if err != nil {
			return err
		}
		request.TimeAssignmentTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	response, err := s.Client.CreateOperatorControlAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperatorControlAssignment
	return nil
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) Get() error {
	request := oci_operator_access_control.GetOperatorControlAssignmentRequest{}

	tmp := s.D.Id()
	request.OperatorControlAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	response, err := s.Client.GetOperatorControlAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperatorControlAssignment
	return nil
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_operator_access_control.UpdateOperatorControlAssignmentRequest{}

	if comment, ok := s.D.GetOkExists("comment"); ok {
		tmp := comment.(string)
		request.Comment = &tmp
	}

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

	if isAutoApproveDuringMaintenance, ok := s.D.GetOkExists("is_auto_approve_during_maintenance"); ok {
		tmp := isAutoApproveDuringMaintenance.(bool)
		request.IsAutoApproveDuringMaintenance = &tmp
	}

	if isEnforcedAlways, ok := s.D.GetOkExists("is_enforced_always"); ok {
		tmp := isEnforcedAlways.(bool)
		request.IsEnforcedAlways = &tmp
	}

	if isLogForwarded, ok := s.D.GetOkExists("is_log_forwarded"); ok {
		tmp := isLogForwarded.(bool)
		request.IsLogForwarded = &tmp
	}

	tmp := s.D.Id()
	request.OperatorControlAssignmentId = &tmp

	if remoteSyslogServerAddress, ok := s.D.GetOkExists("remote_syslog_server_address"); ok {
		tmp := remoteSyslogServerAddress.(string)
		request.RemoteSyslogServerAddress = &tmp
	}

	if remoteSyslogServerCACert, ok := s.D.GetOkExists("remote_syslog_server_ca_cert"); ok {
		tmp := remoteSyslogServerCACert.(string)
		request.RemoteSyslogServerCACert = &tmp
	}

	if remoteSyslogServerPort, ok := s.D.GetOkExists("remote_syslog_server_port"); ok {
		tmp := remoteSyslogServerPort.(int)
		request.RemoteSyslogServerPort = &tmp
	}

	if timeAssignmentFrom, ok := s.D.GetOkExists("time_assignment_from"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAssignmentFrom.(string))
		if err != nil {
			return err
		}
		request.TimeAssignmentFrom = &oci_common.SDKTime{Time: tmp}
	}

	if timeAssignmentTo, ok := s.D.GetOkExists("time_assignment_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAssignmentTo.(string))
		if err != nil {
			return err
		}
		request.TimeAssignmentTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	response, err := s.Client.UpdateOperatorControlAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperatorControlAssignment
	return nil
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) Delete() error {
	request := oci_operator_access_control.DeleteOperatorControlAssignmentRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.OperatorControlAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	_, err := s.Client.DeleteOperatorControlAssignment(context.Background(), request)
	return err
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) SetData() error {
	if s.Res.AssignerId != nil {
		s.D.Set("assigner_id", *s.Res.AssignerId)
	}

	if s.Res.Comment != nil {
		s.D.Set("comment", *s.Res.Comment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DetachmentDescription != nil {
		s.D.Set("detachment_description", *s.Res.DetachmentDescription)
	}

	if s.Res.ErrorCode != nil {
		s.D.Set("error_code", *s.Res.ErrorCode)
	}

	if s.Res.ErrorMessage != nil {
		s.D.Set("error_message", *s.Res.ErrorMessage)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoApproveDuringMaintenance != nil {
		s.D.Set("is_auto_approve_during_maintenance", *s.Res.IsAutoApproveDuringMaintenance)
	}

	if s.Res.IsEnforcedAlways != nil {
		s.D.Set("is_enforced_always", *s.Res.IsEnforcedAlways)
	}

	if s.Res.IsLogForwarded != nil {
		s.D.Set("is_log_forwarded", *s.Res.IsLogForwarded)
	}

	if s.Res.OperatorControlId != nil {
		s.D.Set("operator_control_id", *s.Res.OperatorControlId)
	}

	if s.Res.RemoteSyslogServerAddress != nil {
		s.D.Set("remote_syslog_server_address", *s.Res.RemoteSyslogServerAddress)
	}

	if s.Res.RemoteSyslogServerCACert != nil {
		s.D.Set("remote_syslog_server_ca_cert", *s.Res.RemoteSyslogServerCACert)
	}

	if s.Res.RemoteSyslogServerPort != nil {
		s.D.Set("remote_syslog_server_port", *s.Res.RemoteSyslogServerPort)
	}

	if s.Res.ResourceCompartmentId != nil {
		s.D.Set("resource_compartment_id", *s.Res.ResourceCompartmentId)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAssignmentFrom != nil {
		s.D.Set("time_assignment_from", s.Res.TimeAssignmentFrom.Format(time.RFC3339Nano))
	}

	if s.Res.TimeAssignmentTo != nil {
		s.D.Set("time_assignment_to", s.Res.TimeAssignmentTo.Format(time.RFC3339Nano))
	}

	if s.Res.TimeOfAssignment != nil {
		s.D.Set("time_of_assignment", s.Res.TimeOfAssignment.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.UnassignerId != nil {
		s.D.Set("unassigner_id", *s.Res.UnassignerId)
	}

	return nil
}

func OperatorControlAssignmentSummaryToMap(obj oci_operator_access_control.OperatorControlAssignmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ErrorCode != nil {
		result["error_code"] = int(*obj.ErrorCode)
	}

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEnforcedAlways != nil {
		result["is_enforced_always"] = bool(*obj.IsEnforcedAlways)
	}

	if obj.IsLogForwarded != nil {
		result["is_log_forwarded"] = bool(*obj.IsLogForwarded)
	}

	if obj.OperatorControlId != nil {
		result["operator_control_id"] = string(*obj.OperatorControlId)
	}

	if obj.RemoteSyslogServerAddress != nil {
		result["remote_syslog_server_address"] = string(*obj.RemoteSyslogServerAddress)
	}

	if obj.RemoteSyslogServerPort != nil {
		result["remote_syslog_server_port"] = int(*obj.RemoteSyslogServerPort)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeAssignmentFrom != nil {
		result["time_assignment_from"] = obj.TimeAssignmentFrom.String()
	}

	if obj.TimeAssignmentTo != nil {
		result["time_assignment_to"] = obj.TimeAssignmentTo.String()
	}

	if obj.TimeOfAssignment != nil {
		result["time_of_assignment"] = obj.TimeOfAssignment.String()
	}

	return result
}

func (s *OperatorAccessControlOperatorControlAssignmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_operator_access_control.ChangeOperatorControlAssignmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OperatorControlAssignmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "operator_access_control")

	_, err := s.Client.ChangeOperatorControlAssignmentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
