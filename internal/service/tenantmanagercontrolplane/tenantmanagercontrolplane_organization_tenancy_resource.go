// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneOrganizationTenancyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: importTenantmanagercontrolplaneOrganizationTenancy,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createTenantmanagercontrolplaneOrganizationTenancy,
		Read:     readTenantmanagercontrolplaneOrganizationTenancy,
		Delete:   deleteTenantmanagercontrolplaneOrganizationTenancy,
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
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"governance_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policy_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"is_approved_for_transfer": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_joined": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_left": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createTenantmanagercontrolplaneOrganizationTenancy(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneOrganizationTenancyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationClient()
	sync.WorkRequestClient = m.(*client.OracleClients).TenantmanagercontrolplaneWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readTenantmanagercontrolplaneOrganizationTenancy(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneOrganizationTenancyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationClient()

	return tfresource.ReadResource(sync)
}

func deleteTenantmanagercontrolplaneOrganizationTenancy(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneOrganizationTenancyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationClient()

	return tfresource.DeleteResource(d, sync)
}

// importTenantmanagercontrolplaneOrganizationTenancy parses a composite import id of the form
// "organizations/{organizationId}/tenancies/{tenancyId}" (the GetOrganizationTenancy URI shape),
// because reading the resource requires the organization OCID in addition to the tenancy OCID.
func importTenantmanagercontrolplaneOrganizationTenancy(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) == 4 && parts[0] == "organizations" && parts[2] == "tenancies" && parts[1] != "" && parts[3] != "" {
		d.Set("organization_id", parts[1])
		d.SetId(parts[3])
		return []*schema.ResourceData{d}, nil
	}
	return nil, fmt.Errorf("import id %q is not valid; expected the form organizations/{organizationId}/tenancies/{tenancyId}", d.Id())
}

type TenantmanagercontrolplaneOrganizationTenancyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_tenantmanagercontrolplane.OrganizationClient
	WorkRequestClient      *oci_tenantmanagercontrolplane.WorkRequestClient
	Res                    *oci_tenantmanagercontrolplane.OrganizationTenancy
	DisableNotFoundRetries bool
}

func (s *TenantmanagercontrolplaneOrganizationTenancyResourceCrud) ID() string {
	return *s.Res.TenancyId
}

func (s *TenantmanagercontrolplaneOrganizationTenancyResourceCrud) Create() error {
	request := oci_tenantmanagercontrolplane.CreateChildTenancyRequest{}

	if adminEmail, ok := s.D.GetOkExists("admin_email"); ok {
		tmp := adminEmail.(string)
		request.AdminEmail = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyName, ok := s.D.GetOkExists("tenancy_name"); ok {
		tmp := tenancyName.(string)
		request.TenancyName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane")

	response, err := s.Client.CreateChildTenancy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOrganizationTenancyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane"), oci_tenantmanagercontrolplane.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *TenantmanagercontrolplaneOrganizationTenancyResourceCrud) getOrganizationTenancyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_tenantmanagercontrolplane.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	organizationTenancyId, err := organizationTenancyWaitForWorkRequest(workId, "tenancy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	if organizationTenancyId == nil {
		return fmt.Errorf("work request %s succeeded but did not return a child tenancy identifier", *workId)
	}
	s.D.SetId(*organizationTenancyId)

	return s.Get()
}

// organizationTenancyWorkRequestIsDone reports whether the work request should be considered
// terminal. In addition to the normal terminal states, the Organizations service can leave a
// CREATE_CHILD_TENANCY work request in IN_PROGRESS at 100% completion without ever setting
// timeFinished or transitioning to SUCCEEDED; treat that as done so we don't poll until timeout.
func organizationTenancyWorkRequestIsDone(response oci_tenantmanagercontrolplane.GetWorkRequestResponse) bool {
	if response.TimeFinished != nil {
		return true
	}
	if response.PercentComplete != nil && *response.PercentComplete >= 100 {
		return true
	}
	return false
}

func organizationTenancyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

		// Only stop if the work request is done (finished, or stuck IN_PROGRESS at 100%)
		if workRequestResponse, ok := response.Response.(oci_tenantmanagercontrolplane.GetWorkRequestResponse); ok {
			return !organizationTenancyWorkRequestIsDone(workRequestResponse)
		}
		return false
	}
}

func organizationTenancyWaitForWorkRequest(wId *string, entityType string, action oci_tenantmanagercontrolplane.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_tenantmanagercontrolplane.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "tenantmanagercontrolplane")
	retryPolicy.ShouldRetryOperation = organizationTenancyWorkRequestShouldRetryFunc(timeout)

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
			response, err = client.GetWorkRequest(context.Background(),
				oci_tenantmanagercontrolplane.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(organizationTenancyWorkRequestEffectiveStatus(response)), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	// Determine success/failure from the work request status (the Organizations service uses
	// actionType RELATED for affected resources, so we cannot rely on a CREATED action).
	switch organizationTenancyWorkRequestEffectiveStatus(response) {
	case oci_tenantmanagercontrolplane.OperationStatusFailed, oci_tenantmanagercontrolplane.OperationStatusCanceled:
		return nil, getErrorFromTenantmanagercontrolplaneOrganizationTenancyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	// The work request response contains the resources affected by the operation. Match by entity
	// type only ("childTenancy" contains "tenancy"); the identifier is the child tenancy OCID.
	var identifier *string
	for _, res := range response.Resources {
		if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			identifier = res.Identifier
			break
		}
	}

	return identifier, nil
}

// organizationTenancyWorkRequestEffectiveStatus maps a work request that is stuck IN_PROGRESS at
// 100% completion to SUCCEEDED; otherwise it returns the reported status.
func organizationTenancyWorkRequestEffectiveStatus(response oci_tenantmanagercontrolplane.GetWorkRequestResponse) oci_tenantmanagercontrolplane.OperationStatusEnum {
	if response.Status == oci_tenantmanagercontrolplane.OperationStatusInProgress && response.PercentComplete != nil && *response.PercentComplete >= 100 {
		return oci_tenantmanagercontrolplane.OperationStatusSucceeded
	}
	return response.Status
}

func getErrorFromTenantmanagercontrolplaneOrganizationTenancyWorkRequest(client *oci_tenantmanagercontrolplane.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_tenantmanagercontrolplane.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
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

func (s *TenantmanagercontrolplaneOrganizationTenancyResourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetOrganizationTenancyRequest{}

	if organizationId, ok := s.D.GetOkExists("organization_id"); ok {
		tmp := organizationId.(string)
		request.OrganizationId = &tmp
	}

	tmp := s.D.Id()
	request.TenancyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane")

	response, err := s.Client.GetOrganizationTenancy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OrganizationTenancy
	return nil
}

// Delete intentionally does NOT terminate the child tenancy. Child tenancy termination is not a
// reliable self-service operation through the organization API (for example, the parent
// organization cannot terminate a child that is OPTED_OUT of governance), and in some cases must be
// performed from within the child tenancy itself or via an Oracle Support request. Destroying this
// resource therefore only removes it from Terraform state; the underlying tenancy is left untouched.
func (s *TenantmanagercontrolplaneOrganizationTenancyResourceCrud) Delete() error {
	log.Printf("[WARN] Destroying oci_tenantmanagercontrolplane_organization_tenancy %q removes it from Terraform state only; the child tenancy is NOT terminated. Terminate it from within the child tenancy (Governance & Administration -> Tenancy Management -> Terminate) or via an Oracle Support request.", s.D.Id())
	return nil
}

func (s *TenantmanagercontrolplaneOrganizationTenancyResourceCrud) SetData() error {
	s.D.Set("governance_status", s.Res.GovernanceStatus)

	if s.Res.IsApprovedForTransfer != nil {
		s.D.Set("is_approved_for_transfer", *s.Res.IsApprovedForTransfer)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
		// The service does not return the original create inputs; tenancy_name mirrors the
		// returned name, so set it too to avoid a spurious diff (notably after import).
		s.D.Set("tenancy_name", *s.Res.Name)
	}

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeJoined != nil {
		s.D.Set("time_joined", s.Res.TimeJoined.String())
	}

	if s.Res.TimeLeft != nil {
		s.D.Set("time_left", s.Res.TimeLeft.String())
	}

	return nil
}
