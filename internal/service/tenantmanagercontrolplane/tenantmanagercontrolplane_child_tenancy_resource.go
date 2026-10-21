// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneChildTenancyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createTenantmanagercontrolplaneChildTenancyWithContext,
		ReadContext:   readTenantmanagercontrolplaneChildTenancyWithContext,
		DeleteContext: deleteTenantmanagercontrolplaneChildTenancyWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"home_region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional creation inputs, not computed outputs: the work request
			// supplies an ID but does not return these fields for state refresh.
			"features": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"governance_status": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"policy_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createTenantmanagercontrolplaneChildTenancyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &TenantmanagercontrolplaneChildTenancyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationClient()
	sync.WorkRequestClient = m.(*client.OracleClients).TenantmanagercontrolplaneWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readTenantmanagercontrolplaneChildTenancyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteTenantmanagercontrolplaneChildTenancyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type TenantmanagercontrolplaneChildTenancyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_tenantmanagercontrolplane.OrganizationClient
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_tenantmanagercontrolplane.WorkRequestClient
}

func (s *TenantmanagercontrolplaneChildTenancyResourceCrud) ID() string {
	// CreateChildTenancy returns a work request, not a ChildTenancy model.
	// The completed work request supplies the identifier stored in state.
	return s.D.Id()
}

func (s *TenantmanagercontrolplaneChildTenancyResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_tenantmanagercontrolplane.CreateChildTenancyRequest{}

	if adminEmail, ok := s.D.GetOkExists("admin_email"); ok {
		tmp := adminEmail.(string)
		request.AdminEmail = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if features, ok := s.D.GetOkExists("features"); ok {
		interfaces := features.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("features") {
			request.Features = tmp
		}
	}

	if governanceStatus, ok := s.D.GetOkExists("governance_status"); ok {
		request.GovernanceStatus = oci_tenantmanagercontrolplane.GovernanceStatusEnum(governanceStatus.(string))
	}

	if homeRegion, ok := s.D.GetOkExists("home_region"); ok {
		tmp := homeRegion.(string)
		request.HomeRegion = &tmp
	}

	if policyName, ok := s.D.GetOkExists("policy_name"); ok {
		tmp := policyName.(string)
		request.PolicyName = &tmp
	}

	if tenancyName, ok := s.D.GetOkExists("tenancy_name"); ok {
		tmp := tenancyName.(string)
		request.TenancyName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane")

	response, err := s.Client.CreateChildTenancy(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_tenantmanagercontrolplane.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(ctx,
		oci_tenantmanagercontrolplane.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "childtenancy") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getChildTenancyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane"), oci_tenantmanagercontrolplane.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *TenantmanagercontrolplaneChildTenancyResourceCrud) getChildTenancyFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_tenantmanagercontrolplane.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	childTenancyId, err := childTenancyWaitForWorkRequest(ctx, workId, "childtenancy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*childTenancyId)

	// The API has no child-tenancy GET operation to refresh a response model.
	// Creation is complete once the work request succeeds and supplies the ID.
	return nil
}

func childTenancyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "tenantmanagercontrolplane", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_tenantmanagercontrolplane.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func childTenancyWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_tenantmanagercontrolplane.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_tenantmanagercontrolplane.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "tenantmanagercontrolplane")
	retryPolicy.ShouldRetryOperation = childTenancyWorkRequestShouldRetryFunc(timeout)

	response := oci_tenantmanagercontrolplane.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_tenantmanagercontrolplane.OperationStatusInProgress),
			string(oci_tenantmanagercontrolplane.OperationStatusAccepted),
			string(oci_tenantmanagercontrolplane.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_tenantmanagercontrolplane.OperationStatusSucceeded),
			string(oci_tenantmanagercontrolplane.OperationStatusFailed),
			string(oci_tenantmanagercontrolplane.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_tenantmanagercontrolplane.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_tenantmanagercontrolplane.OperationStatusFailed || response.Status == oci_tenantmanagercontrolplane.OperationStatusCanceled {
		return nil, getErrorFromTenantmanagercontrolplaneChildTenancyWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromTenantmanagercontrolplaneChildTenancyWorkRequest(ctx context.Context, client *oci_tenantmanagercontrolplane.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_tenantmanagercontrolplane.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_tenantmanagercontrolplane.ListWorkRequestErrorsRequest{
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

func (s *TenantmanagercontrolplaneChildTenancyResourceCrud) SetData() error {
	return nil
}
