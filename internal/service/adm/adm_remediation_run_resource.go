// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRunResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAdmRemediationRun,
		Read:     readAdmRemediationRun,
		Update:   updateAdmRemediationRun,
		Delete:   deleteAdmRemediationRun,
		Schema: map[string]*schema.Schema{
			// Required
			"remediation_recipe_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
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
			"display_name": {
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
			"current_stage_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remediation_run_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_finished": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
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

func createAdmRemediationRun(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
		if err != nil {
			log.Printf("error doing a Get() after compartment update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment update: %v", err)
		}
	}
	return nil
}

func readAdmRemediationRun(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

func updateAdmRemediationRun(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAdmRemediationRun(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AdmRemediationRunResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_adm.ApplicationDependencyManagementClient
	Res                    *oci_adm.RemediationRun
	DisableNotFoundRetries bool
}

func (s *AdmRemediationRunResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AdmRemediationRunResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_adm.RemediationRunLifecycleStateAccepted),
	}
}

func (s *AdmRemediationRunResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_adm.RemediationRunLifecycleStateAccepted),
		string(oci_adm.RemediationRunLifecycleStateInProgress),
		string(oci_adm.RemediationRunLifecycleStateSucceeded),
		string(oci_adm.RemediationRunLifecycleStateFailed),
		string(oci_adm.RemediationRunLifecycleStateCanceled),
	}
}

func (s *AdmRemediationRunResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_adm.RemediationRunLifecycleStateDeleting),
	}
}

func (s *AdmRemediationRunResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_adm.RemediationRunLifecycleStateDeleted),
	}
}

func (s *AdmRemediationRunResourceCrud) Create() error {
	request := oci_adm.CreateRemediationRunRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if remediationRecipeId, ok := s.D.GetOkExists("remediation_recipe_id"); ok {
		tmp := remediationRecipeId.(string)
		request.RemediationRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.CreateRemediationRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RemediationRun
	return nil
}

func (s *AdmRemediationRunResourceCrud) Get() error {
	request := oci_adm.GetRemediationRunRequest{}

	tmp := s.D.Id()
	request.RemediationRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.GetRemediationRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RemediationRun
	return nil
}

func (s *AdmRemediationRunResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists(""); ok && s.D.HasChange("") {
		err := s.CancelRemediationRun()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_adm.UpdateRemediationRunRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.RemediationRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.UpdateRemediationRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RemediationRun
	return nil
}

func (s *AdmRemediationRunResourceCrud) Delete() error {
	request := oci_adm.DeleteRemediationRunRequest{}

	err := s.Get()
	if err != nil {
		return err
	}

	if s.Res.LifecycleState == oci_adm.RemediationRunLifecycleStateInProgress ||
		s.Res.LifecycleState == oci_adm.RemediationRunLifecycleStateAccepted {
		err = s.CancelRemediationRun()
		if err != nil {
			return err
		}
		conditionFunc := func() bool {
			return (s.Res.LifecycleState == oci_adm.RemediationRunLifecycleStateCanceled ||
				s.Res.LifecycleState == oci_adm.RemediationRunLifecycleStateFailed ||
				s.Res.LifecycleState == oci_adm.RemediationRunLifecycleStateSucceeded)
		}
		err = tfresource.WaitForResourceCondition(s, conditionFunc, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	tmp := s.D.Id()
	request.RemediationRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	_, err = s.Client.DeleteRemediationRun(context.Background(), request)
	return err
}

func (s *AdmRemediationRunResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("current_stage_type", s.Res.CurrentStageType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RemediationRecipeId != nil {
		s.D.Set("remediation_recipe_id", *s.Res.RemediationRecipeId)
	}

	s.D.Set("remediation_run_source", s.Res.RemediationRunSource)

	stages := []interface{}{}
	for _, item := range s.Res.Stages {
		stages = append(stages, StageSummaryToMap(item))
	}
	s.D.Set("stages", stages)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *AdmRemediationRunResourceCrud) CancelRemediationRun() error {
	request := oci_adm.CancelRemediationRunRequest{}

	idTmp := s.D.Id()
	request.RemediationRunId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	_, err := s.Client.CancelRemediationRun(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func RemediationRunSummaryToMap(obj oci_adm.RemediationRunSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["current_stage_type"] = string(obj.CurrentStageType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.RemediationRecipeId != nil {
		result["remediation_recipe_id"] = string(*obj.RemediationRecipeId)
	}

	result["remediation_run_source"] = string(obj.RemediationRunSource)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func StageSummaryToMap(obj oci_adm.StageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *AdmRemediationRunResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_adm.ChangeRemediationRunCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RemediationRunId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	_, err := s.Client.ChangeRemediationRunCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
