// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_waas "github.com/oracle/oci-go-sdk/v56/waas"
)

func WaasHttpRedirectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createWaasHttpRedirect,
		Read:     readWaasHttpRedirect,
		Update:   updateWaasHttpRedirect,
		Delete:   deleteWaasHttpRedirect,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"host": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"query": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
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
			"response_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
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

func createWaasHttpRedirect(d *schema.ResourceData, m interface{}) error {
	sync := &WaasHttpRedirectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedirectClient()
	sync.WaasClient = m.(*client.OracleClients).WaasClient()

	return tfresource.CreateResource(d, sync)
}

func readWaasHttpRedirect(d *schema.ResourceData, m interface{}) error {
	sync := &WaasHttpRedirectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedirectClient()

	return tfresource.ReadResource(sync)
}

func updateWaasHttpRedirect(d *schema.ResourceData, m interface{}) error {
	sync := &WaasHttpRedirectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedirectClient()
	sync.WaasClient = m.(*client.OracleClients).WaasClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteWaasHttpRedirect(d *schema.ResourceData, m interface{}) error {
	sync := &WaasHttpRedirectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedirectClient()
	sync.WaasClient = m.(*client.OracleClients).WaasClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type WaasHttpRedirectResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waas.RedirectClient
	WaasClient             *oci_waas.WaasClient
	Res                    *oci_waas.HttpRedirect
	DisableNotFoundRetries bool
}

func (s *WaasHttpRedirectResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *WaasHttpRedirectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waas.LifecycleStatesCreating),
	}
}

func (s *WaasHttpRedirectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waas.LifecycleStatesActive),
	}
}

func (s *WaasHttpRedirectResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waas.LifecycleStatesDeleting),
	}
}

func (s *WaasHttpRedirectResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waas.LifecycleStatesDeleted),
	}
}

func (s *WaasHttpRedirectResourceCrud) Create() error {
	request := oci_waas.CreateHttpRedirectRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if responseCode, ok := s.D.GetOkExists("response_code"); ok {
		tmp := responseCode.(int)
		request.ResponseCode = &tmp
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target", 0)
			tmp, err := s.mapToHttpRedirectTarget(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Target = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.CreateHttpRedirect(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHttpRedirectFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas"), oci_waas.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *WaasHttpRedirectResourceCrud) getHttpRedirectFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_waas.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	httpRedirectId, err := httpRedirectWaitForWorkRequest(workId, "redirect",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WaasClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, httpRedirectId)
		_, cancelErr := s.WaasClient.CancelWorkRequest(context.Background(),
			oci_waas.CancelWorkRequestRequest{
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
	s.D.SetId(*httpRedirectId)

	return s.Get()
}

func httpRedirectWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "waas", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_waas.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func httpRedirectWaitForWorkRequest(wId *string, entityType string, action oci_waas.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waas.WaasClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "waas")
	retryPolicy.ShouldRetryOperation = httpRedirectWorkRequestShouldRetryFunc(timeout)

	response := oci_waas.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_waas.WorkRequestStatusValuesInProgress),
			string(oci_waas.WorkRequestStatusValuesAccepted),
			string(oci_waas.WorkRequestStatusValuesCanceling),
		},
		Target: []string{
			string(oci_waas.WorkRequestStatusValuesSucceeded),
			string(oci_waas.WorkRequestStatusValuesFailed),
			string(oci_waas.WorkRequestStatusValuesCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_waas.GetWorkRequestRequest{
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
	var workRequestErr error
	if len(response.Errors) > 0 {
		errorMessage := getErrorFromWaasHttpRedirectWorkRequest(response)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromWaasHttpRedirectWorkRequest(response oci_waas.GetWorkRequestResponse) string {
	allErrs := make([]string, 0)
	for _, wrkErr := range response.Errors {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage
}

func (s *WaasHttpRedirectResourceCrud) Get() error {
	request := oci_waas.GetHttpRedirectRequest{}

	tmp := s.D.Id()
	request.HttpRedirectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.GetHttpRedirect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HttpRedirect
	return nil
}

func (s *WaasHttpRedirectResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waas.UpdateHttpRedirectRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.HttpRedirectId = &tmp

	if responseCode, ok := s.D.GetOkExists("response_code"); ok {
		tmp := responseCode.(int)
		request.ResponseCode = &tmp
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target", 0)
			tmp, err := s.mapToHttpRedirectTarget(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Target = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.UpdateHttpRedirect(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHttpRedirectFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas"), oci_waas.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *WaasHttpRedirectResourceCrud) Delete() error {
	request := oci_waas.DeleteHttpRedirectRequest{}

	tmp := s.D.Id()
	request.HttpRedirectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.DeleteHttpRedirect(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := httpRedirectWaitForWorkRequest(workId, "waas",
		oci_waas.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WaasClient)
	return delWorkRequestErr
}

func (s *WaasHttpRedirectResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("response_code", s.Res.ResponseCode)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Target != nil {
		s.D.Set("target", []interface{}{HttpRedirectTargetToMap(s.Res.Target)})
	} else {
		s.D.Set("target", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *WaasHttpRedirectResourceCrud) mapToHttpRedirectTarget(fieldKeyFormat string) (oci_waas.HttpRedirectTarget, error) {
	result := oci_waas.HttpRedirectTarget{}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_waas.HttpRedirectTargetProtocolEnum(protocol.(string))
	}

	if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
		tmp := query.(string)
		result.Query = &tmp
	} else {
		tmp := ""
		result.Query = &tmp
	}

	return result, nil
}

func HttpRedirectTargetToMap(obj *oci_waas.HttpRedirectTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.Query != nil {
		result["query"] = string(*obj.Query)
	}

	return result
}

func (s *WaasHttpRedirectResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waas.ChangeHttpRedirectCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.HttpRedirectId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	_, err := s.Client.ChangeHttpRedirectCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
