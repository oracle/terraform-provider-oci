// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package security_attribute

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_security_attribute "github.com/oracle/oci-go-sdk/v65/securityattribute"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func SecurityAttributeSecurityAttributeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("15m"),
			Update: tfresource.GetTimeoutDuration("15m"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createSecurityAttributeSecurityAttribute,
		Read:   readSecurityAttributeSecurityAttribute,
		Update: updateSecurityAttributeSecurityAttribute,
		Delete: deleteSecurityAttributeSecurityAttribute,
		Schema: map[string]*schema.Schema{
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"security_attribute_namespace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"is_retired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"validator": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"validator_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DEFAULT",
								"ENUM",
							}, true),
						},

						// Optional
						"values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_attribute_namespace_name": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSecurityAttributeSecurityAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.CreateResource(d, sync)
}

func readSecurityAttributeSecurityAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.ReadResource(sync)
}

func updateSecurityAttributeSecurityAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteSecurityAttributeSecurityAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type SecurityAttributeSecurityAttributeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_security_attribute.SecurityAttributeClient
	Res                    *oci_security_attribute.SecurityAttribute
	DisableNotFoundRetries bool
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) ID() string {
	return GetSecurityAttributeCompositeId(s.D.Get("name").(string), s.D.Get("security_attribute_namespace_id").(string))
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_security_attribute.SecurityAttributeLifecycleStateActive),
	}
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_security_attribute.SecurityAttributeLifecycleStateDeleting),
	}
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_security_attribute.SecurityAttributeLifecycleStateDeleted),
	}
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) Create() error {
	request := oci_security_attribute.CreateSecurityAttributeRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if securityAttributeNamespaceId, ok := s.D.GetOkExists("security_attribute_namespace_id"); ok {
		tmp := securityAttributeNamespaceId.(string)
		request.SecurityAttributeNamespaceId = &tmp
	}

	if validator, ok := s.D.GetOkExists("validator"); ok {
		if tmpList := validator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "validator", 0)
			tmp, err := s.mapToBaseSecurityAttributeValidator(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Validator = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	response, err := s.Client.CreateSecurityAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityAttribute
	return nil
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) getSecurityAttributeFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_security_attribute.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	securityAttributeId, err := securityAttributeWaitForWorkRequest(workId, "securityattributenamespace",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*securityAttributeId)

	return s.Get()
}

func securityAttributeWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "security_attribute", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_security_attribute.GetSecurityAttributeWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func securityAttributeWaitForWorkRequest(wId *string, entityType string, action oci_security_attribute.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_security_attribute.SecurityAttributeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "security_attribute")
	retryPolicy.ShouldRetryOperation = securityAttributeWorkRequestShouldRetryFunc(timeout)

	response := oci_security_attribute.GetSecurityAttributeWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_security_attribute.WorkRequestStatusInProgress),
			string(oci_security_attribute.WorkRequestStatusAccepted),
			string(oci_security_attribute.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_security_attribute.WorkRequestStatusSucceeded),
			string(oci_security_attribute.WorkRequestStatusFailed),
			string(oci_security_attribute.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetSecurityAttributeWorkRequest(context.Background(),
				oci_security_attribute.GetSecurityAttributeWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.SecurityAttributeWorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.SecurityAttributeWorkRequest.Resources {
		if res.Identifier != nil {
			identifier = res.Identifier
			break
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_security_attribute.WorkRequestStatusFailed || response.Status == oci_security_attribute.WorkRequestStatusCanceled {
		return nil, getErrorFromSecurityAttributeSecurityAttributeWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromSecurityAttributeSecurityAttributeWorkRequest(client *oci_security_attribute.SecurityAttributeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_security_attribute.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListSecurityAttributeWorkRequestErrors(context.Background(),
		oci_security_attribute.ListSecurityAttributeWorkRequestErrorsRequest{
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

func (s *SecurityAttributeSecurityAttributeResourceCrud) Get() error {
	request := oci_security_attribute.GetSecurityAttributeRequest{}

	if securityAttributeName, ok := s.D.GetOkExists("name"); ok {
		tmp := securityAttributeName.(string)
		request.SecurityAttributeName = &tmp
	}

	if securityAttributeNamespaceId, ok := s.D.GetOkExists("security_attribute_namespace_id"); ok {
		tmp := securityAttributeNamespaceId.(string)
		request.SecurityAttributeNamespaceId = &tmp
	}

	securityAttributeName, securityAttributeNamespaceId, err := parseSecurityAttributeCompositeId(s.D.Id())
	if err == nil {
		request.SecurityAttributeName = &securityAttributeName
		request.SecurityAttributeNamespaceId = &securityAttributeNamespaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	response, err := s.Client.GetSecurityAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityAttribute
	return nil
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) Update() error {
	request := oci_security_attribute.UpdateSecurityAttributeRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if isRetired, ok := s.D.GetOkExists("is_retired"); ok {
		tmp := isRetired.(bool)
		request.IsRetired = &tmp
	}

	if securityAttributeName, ok := s.D.GetOkExists("name"); ok {
		tmp := securityAttributeName.(string)
		request.SecurityAttributeName = &tmp
	}

	if securityAttributeNamespaceId, ok := s.D.GetOkExists("security_attribute_namespace_id"); ok {
		tmp := securityAttributeNamespaceId.(string)
		request.SecurityAttributeNamespaceId = &tmp
	}

	if validator, ok := s.D.GetOkExists("validator"); ok {
		if tmpList := validator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "validator", 0)
			tmp, err := s.mapToBaseSecurityAttributeValidator(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Validator = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	response, err := s.Client.UpdateSecurityAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityAttribute
	return nil
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) Delete() error {
	request := oci_security_attribute.DeleteSecurityAttributeRequest{}

	if securityAttributeName, ok := s.D.GetOkExists("name"); ok {
		tmp := securityAttributeName.(string)
		request.SecurityAttributeName = &tmp
	}

	if securityAttributeNamespaceId, ok := s.D.GetOkExists("security_attribute_namespace_id"); ok {
		tmp := securityAttributeNamespaceId.(string)
		request.SecurityAttributeNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	response, err := s.Client.DeleteSecurityAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := securityAttributeWaitForWorkRequest(workId, "securityattributenamespace",
		oci_security_attribute.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) SetData() error {

	securityAttributeName, securityAttributeNamespaceId, err := parseSecurityAttributeCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &securityAttributeName)
		s.D.Set("security_attribute_namespace_id", &securityAttributeNamespaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SecurityAttributeNamespaceId != nil {
		s.D.Set("security_attribute_namespace_id", *s.Res.SecurityAttributeNamespaceId)
	}

	if s.Res.SecurityAttributeNamespaceName != nil {
		s.D.Set("security_attribute_namespace_name", *s.Res.SecurityAttributeNamespaceName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	if s.Res.Validator != nil {
		validatorArray := []interface{}{}
		if validatorMap := BaseSecurityAttributeValidatorToMap(&s.Res.Validator); validatorMap != nil {
			validatorArray = append(validatorArray, validatorMap)
		}
		s.D.Set("validator", validatorArray)
	} else {
		s.D.Set("validator", nil)
	}

	return nil
}

func GetSecurityAttributeCompositeId(securityAttributeName string, securityAttributeNamespaceId string) string {
	securityAttributeName = url.PathEscape(securityAttributeName)
	securityAttributeNamespaceId = url.PathEscape(securityAttributeNamespaceId)
	compositeId := "securityAttributeNamespaces/" + securityAttributeNamespaceId + "/securityAttributes/" + securityAttributeName
	return compositeId
}

func parseSecurityAttributeCompositeId(compositeId string) (securityAttributeName string, securityAttributeNamespaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("securityAttributeNamespaces/.*/securityAttributes/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	securityAttributeNamespaceId, _ = url.PathUnescape(parts[1])
	securityAttributeName, _ = url.PathUnescape(parts[3])

	return
}

func (s *SecurityAttributeSecurityAttributeResourceCrud) mapToBaseSecurityAttributeValidator(fieldKeyFormat string) (oci_security_attribute.BaseSecurityAttributeValidator, error) {
	var baseObject oci_security_attribute.BaseSecurityAttributeValidator
	//discriminator
	validatorTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validator_type"))
	var validatorType string
	if ok {
		validatorType = validatorTypeRaw.(string)
	} else {
		validatorType = "" // default value
	}
	switch strings.ToLower(validatorType) {
	case strings.ToLower("DEFAULT"):
		details := oci_security_attribute.DefaultSecurityAttributeValidator{}
		baseObject = details
	case strings.ToLower("ENUM"):
		details := oci_security_attribute.EnumSecurityAttributeValidator{}
		if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
			interfaces := values.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
				details.Values = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown validator_type '%v' was specified", validatorType)
	}
	return baseObject, nil
}

func BaseSecurityAttributeValidatorToMap(obj *oci_security_attribute.BaseSecurityAttributeValidator) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_security_attribute.DefaultSecurityAttributeValidator:
		result["validator_type"] = "DEFAULT"
	case oci_security_attribute.EnumSecurityAttributeValidator:
		result["validator_type"] = "ENUM"

		result["values"] = v.Values
	default:
		log.Printf("[WARN] Received 'validator_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
