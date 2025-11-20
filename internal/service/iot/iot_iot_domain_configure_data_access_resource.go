// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotIotDomainConfigureDataAccessResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotIotDomainConfigureDataAccessWithContext,
		ReadContext:   readIotIotDomainConfigureDataAccessWithContext,
		DeleteContext: deleteIotIotDomainConfigureDataAccessWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"iot_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"APEX",
					"DIRECT",
					"ORDS",
				}, true),
			},

			// Optional
			"db_allow_listed_identity_group_names": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"db_allowed_identity_domain_host": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_workspace_admin_initial_password": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			// Computed
		},
	}
}

func createIotIotDomainConfigureDataAccessWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainConfigureDataAccessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotIotDomainConfigureDataAccessWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteIotIotDomainConfigureDataAccessWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type IotIotDomainConfigureDataAccessResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.IotDomain
	DisableNotFoundRetries bool
}

func (s *IotIotDomainConfigureDataAccessResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotIotDomainConfigureDataAccessResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.ConfigureIotDomainDataAccessRequest{}
	err := s.populateTopLevelPolymorphicConfigureIotDomainDataAccessRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.ConfigureIotDomainDataAccess(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotDomainConfigureDataAccessFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IotIotDomainConfigureDataAccessResourceCrud) getIotDomainConfigureDataAccessFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_iot.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	iotDomainId, err := iotDomainConfigureDataAccessWaitForWorkRequest(ctx, workId, "iotdomain",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*iotDomainId)

	return s.GetWithContext(ctx)
}

func (s *IotIotDomainConfigureDataAccessResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotDomainRequest{}

	tmp := s.D.Id()
	request.IotDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetIotDomain(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.IotDomain
	return nil
}

func iotDomainConfigureDataAccessWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "iot", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_iot.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func iotDomainConfigureDataAccessWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_iot.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_iot.IotClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "iot")
	retryPolicy.ShouldRetryOperation = iotDomainConfigureDataAccessWorkRequestShouldRetryFunc(timeout)

	response := oci_iot.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_iot.OperationStatusInProgress),
			string(oci_iot.OperationStatusAccepted),
			string(oci_iot.OperationStatusWaiting),
		},
		Target: []string{
			string(oci_iot.OperationStatusSucceeded),
			string(oci_iot.OperationStatusFailed),
			string(oci_iot.OperationStatusNeedsAttention),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_iot.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_iot.OperationStatusFailed || response.Status == oci_iot.OperationStatusNeedsAttention {
		return nil, getErrorFromIotIotDomainConfigureDataAccessWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIotIotDomainConfigureDataAccessWorkRequest(ctx context.Context, client *oci_iot.IotClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_iot.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_iot.ListWorkRequestErrorsRequest{
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

func (s *IotIotDomainConfigureDataAccessResourceCrud) SetData() error {
	return nil
}

func (s *IotIotDomainConfigureDataAccessResourceCrud) populateTopLevelPolymorphicConfigureIotDomainDataAccessRequest(request *oci_iot.ConfigureIotDomainDataAccessRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("APEX"):
		details := oci_iot.ApexDataAccessDetails{}
		if dbWorkspaceAdminInitialPassword, ok := s.D.GetOkExists("db_workspace_admin_initial_password"); ok {
			tmp := dbWorkspaceAdminInitialPassword.(string)
			details.DbWorkspaceAdminInitialPassword = &tmp
		}
		if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
			tmp := iotDomainId.(string)
			request.IotDomainId = &tmp
		}
		request.ConfigureIotDomainDataAccessDetails = details
	case strings.ToLower("DIRECT"):
		details := oci_iot.DirectDataAccessDetails{}
		interfaces := s.D.Get("db_allow_listed_identity_group_names").([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		details.DbAllowListedIdentityGroupNames = tmp
		if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
			tmp := iotDomainId.(string)
			request.IotDomainId = &tmp
		}
		request.ConfigureIotDomainDataAccessDetails = details
	case strings.ToLower("ORDS"):
		details := oci_iot.OrdsDataAccessDetails{}
		if dbAllowedIdentityDomainHost, ok := s.D.GetOkExists("db_allowed_identity_domain_host"); ok {
			tmp := dbAllowedIdentityDomainHost.(string)
			details.DbAllowedIdentityDomainHost = &tmp
		}
		if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
			tmp := iotDomainId.(string)
			request.IotDomainId = &tmp
		}
		request.ConfigureIotDomainDataAccessDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
