// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceMeshMeshResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceMeshMesh,
		Read:     readServiceMeshMesh,
		Update:   updateServiceMeshMesh,
		Delete:   deleteServiceMeshMesh,
		Schema: map[string]*schema.Schema{
			// Required
			"certificate_authorities": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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
			"display_name": {
				Type:     schema.TypeString,
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
			"mtls": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"minimum": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
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

func createServiceMeshMesh(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshMeshResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceMeshMesh(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshMeshResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

func updateServiceMeshMesh(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshMeshResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteServiceMeshMesh(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshMeshResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceMeshMeshResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_mesh.ServiceMeshClient
	Res                    *oci_service_mesh.Mesh
	DisableNotFoundRetries bool
}

func (s *ServiceMeshMeshResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceMeshMeshResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_service_mesh.MeshLifecycleStateCreating),
	}
}

func (s *ServiceMeshMeshResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_service_mesh.MeshLifecycleStateActive),
	}
}

func (s *ServiceMeshMeshResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_service_mesh.MeshLifecycleStateDeleting),
	}
}

func (s *ServiceMeshMeshResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_service_mesh.MeshLifecycleStateDeleted),
	}
}

func (s *ServiceMeshMeshResourceCrud) Create() error {
	request := oci_service_mesh.CreateMeshRequest{}

	if certificateAuthorities, ok := s.D.GetOkExists("certificate_authorities"); ok {
		interfaces := certificateAuthorities.([]interface{})
		tmp := make([]oci_service_mesh.CertificateAuthority, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_authorities", stateDataIndex)
			converted, err := s.mapToCertificateAuthority(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("certificate_authorities") {
			request.CertificateAuthorities = tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if mtls, ok := s.D.GetOkExists("mtls"); ok {
		if tmpList := mtls.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mtls", 0)
			tmp, err := s.mapToMeshMutualTransportLayerSecurity(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Mtls = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.CreateMesh(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMeshFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ServiceMeshMeshResourceCrud) getMeshFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_service_mesh.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	meshId, err := meshWaitForWorkRequest(workId, "mesh",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, meshId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_service_mesh.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*meshId)

	return s.Get()
}

func meshWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "service_mesh", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_service_mesh.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func meshWaitForWorkRequest(wId *string, entityType string, action oci_service_mesh.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_service_mesh.ServiceMeshClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "service_mesh")
	retryPolicy.ShouldRetryOperation = meshWorkRequestShouldRetryFunc(timeout)

	response := oci_service_mesh.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_service_mesh.OperationStatusInProgress),
			string(oci_service_mesh.OperationStatusAccepted),
			string(oci_service_mesh.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_service_mesh.OperationStatusSucceeded),
			string(oci_service_mesh.OperationStatusFailed),
			string(oci_service_mesh.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_service_mesh.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_service_mesh.OperationStatusFailed || response.Status == oci_service_mesh.OperationStatusCanceled {
		return nil, getErrorFromServiceMeshMeshWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromServiceMeshMeshWorkRequest(client *oci_service_mesh.ServiceMeshClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_service_mesh.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_service_mesh.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *ServiceMeshMeshResourceCrud) Get() error {
	request := oci_service_mesh.GetMeshRequest{}

	tmp := s.D.Id()
	request.MeshId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.GetMesh(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Mesh
	return nil
}

func (s *ServiceMeshMeshResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_service_mesh.UpdateMeshRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.MeshId = &tmp

	if mtls, ok := s.D.GetOkExists("mtls"); ok {
		if tmpList := mtls.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mtls", 0)
			tmp, err := s.mapToMeshMutualTransportLayerSecurity(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Mtls = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.UpdateMesh(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMeshFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ServiceMeshMeshResourceCrud) Delete() error {
	request := oci_service_mesh.DeleteMeshRequest{}

	tmp := s.D.Id()
	request.MeshId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.DeleteMesh(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := meshWaitForWorkRequest(workId, "mesh",
		oci_service_mesh.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ServiceMeshMeshResourceCrud) SetData() error {
	certificateAuthorities := []interface{}{}
	for _, item := range s.Res.CertificateAuthorities {
		certificateAuthorities = append(certificateAuthorities, CertificateAuthorityToMap(item))
	}
	s.D.Set("certificate_authorities", certificateAuthorities)

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Mtls != nil {
		s.D.Set("mtls", []interface{}{MeshMutualTransportLayerSecurityToMap(s.Res.Mtls)})
	} else {
		s.D.Set("mtls", nil)
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

func (s *ServiceMeshMeshResourceCrud) mapToCertificateAuthority(fieldKeyFormat string) (oci_service_mesh.CertificateAuthority, error) {
	result := oci_service_mesh.CertificateAuthority{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func CertificateAuthorityToMap(obj oci_service_mesh.CertificateAuthority) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *ServiceMeshMeshResourceCrud) mapToMeshMutualTransportLayerSecurity(fieldKeyFormat string) (oci_service_mesh.MeshMutualTransportLayerSecurity, error) {
	result := oci_service_mesh.MeshMutualTransportLayerSecurity{}

	if minimum, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minimum")); ok {
		result.Minimum = oci_service_mesh.MutualTransportLayerSecurityModeEnum(minimum.(string))
	}

	return result, nil
}

func MeshMutualTransportLayerSecurityToMap(obj *oci_service_mesh.MeshMutualTransportLayerSecurity) map[string]interface{} {
	result := map[string]interface{}{}

	result["minimum"] = string(obj.Minimum)

	return result
}

func MeshSummaryToMap(obj oci_service_mesh.MeshSummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Mtls != nil {
		result["mtls"] = []interface{}{MeshMutualTransportLayerSecurityToMap(obj.Mtls)}
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

func (s *ServiceMeshMeshResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_service_mesh.ChangeMeshCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MeshId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.ChangeMeshCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMeshFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
