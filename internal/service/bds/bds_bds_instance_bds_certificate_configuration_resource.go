// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceBdsCertificateConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBdsBdsInstanceBdsCertificateConfigurationWithContext,
		ReadContext:   readBdsBdsInstanceBdsCertificateConfigurationWithContext,
		UpdateContext: updateBdsBdsInstanceBdsCertificateConfigurationWithContext,
		DeleteContext: deleteBdsBdsInstanceBdsCertificateConfigurationWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"certificate_authority_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"is_default_configuration": {
				Type:     schema.TypeBool,
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
			"time_last_refreshed_or_generated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"issue_certificate_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"renew_certificate_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"set_default_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_missing_nodes_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func createBdsBdsInstanceBdsCertificateConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCertificateConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readBdsBdsInstanceBdsCertificateConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCertificateConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteBdsBdsInstanceBdsCertificateConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCertificateConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BdsBdsInstanceBdsCertificateConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsCertificateConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) ID() string {
	return GetBdsInstanceBdsCertificateConfigurationCompositeId(*s.Res.Id, s.D.Get("bds_instance_id").(string))
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.BdsCertificateConfigurationLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.BdsCertificateConfigurationLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.BdsCertificateConfigurationLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.BdsCertificateConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_bds.CreateBdsCertificateConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if certificateAuthorityId, ok := s.D.GetOkExists("certificate_authority_id"); ok {
		tmp := certificateAuthorityId.(string)
		request.CertificateAuthorityId = &tmp
	}

	if certificateType, ok := s.D.GetOkExists("certificate_type"); ok {
		request.CertificateType = oci_bds.BdsCertificateConfigurationTypeEnum(certificateType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsCertificateConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceBdsCertificateConfigurationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) UpdateWithContext(ctx context.Context) error {
	return s.GetWithContext(ctx)
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) IssueCertificate(ctx context.Context) error {
	request := oci_bds.GenerateBdsCertificateRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	details := oci_bds.GenerateBdsCertificateDetails{}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		details.ClusterAdminPassword = &tmp
	}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		details.SecretId = &tmp
	}

	bdsCertificateConfigurationId, _, err := parseBdsInstanceBdsCertificateConfigurationCompositeId(s.D.Id())
	if err != nil {
		return err
	}

	configDetails := oci_bds.ConfigLevelManageBdsCertificateDetails{
		CertificateConfigurationId: &bdsCertificateConfigurationId,
	}

	if isMissingNodesOnly, ok := s.D.GetOkExists("is_missing_nodes_only"); ok {
		tmp := isMissingNodesOnly.(bool)
		configDetails.IsMissingNodesOnly = &tmp
	}

	details.ManageCertificateLevelTypeDetails = configDetails
	request.GenerateBdsCertificateDetails = details
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GenerateBdsCertificate(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	val := s.D.Get("issue_certificate_trigger")
	s.D.Set("issue_certificate_trigger", val)

	return s.waitForBdsCertificateConfigurationActionWorkRequest(ctx, workId, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) RenewCertificate(ctx context.Context) error {
	request := oci_bds.RenewBdsCertificateRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	details := oci_bds.RenewBdsCertificateDetails{}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		details.ClusterAdminPassword = &tmp
	}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		details.SecretId = &tmp
	}

	bdsCertificateConfigurationId, _, err := parseBdsInstanceBdsCertificateConfigurationCompositeId(s.D.Id())
	if err != nil {
		return err
	}

	configDetails := oci_bds.ConfigLevelManageBdsCertificateDetails{
		CertificateConfigurationId: &bdsCertificateConfigurationId,
	}

	details.ManageCertificateLevelTypeDetails = configDetails
	request.RenewBdsCertificateDetails = details
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.RenewBdsCertificate(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	val := s.D.Get("renew_certificate_trigger")
	s.D.Set("renew_certificate_trigger", val)

	return s.waitForBdsCertificateConfigurationActionWorkRequest(ctx, workId, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) SetDefaultCertificateConfiguration(ctx context.Context) error {
	request := oci_bds.SetDefaultBdsCertificateConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsCertificateConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCertificateConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCertificateConfigurationId = &bdsCertificateConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	} else {
		log.Printf("[WARN] SetDefaultCertificateConfiguration() unable to parse current ID: %s", s.D.Id())
	}

	details := oci_bds.SetDefaultBdsCertificateConfigurationDetails{}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		details.ClusterAdminPassword = &tmp
	}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		details.SecretId = &tmp
	}

	request.SetDefaultBdsCertificateConfigurationDetails = details
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.SetDefaultBdsCertificateConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	val := s.D.Get("set_default_trigger")
	s.D.Set("set_default_trigger", val)

	return s.getBdsInstanceBdsCertificateConfigurationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"),
		oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) waitForBdsCertificateConfigurationActionWorkRequest(ctx context.Context, workId *string, timeout time.Duration) error {
	_, err := bdsInstanceBdsCertificateConfigurationWaitForWorkRequest(ctx,
		workId, "bds", oci_bds.ActionTypesUpdated, timeout, s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}

	return s.GetWithContext(ctx)
}
func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) getBdsInstanceBdsCertificateConfigurationFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceBdsCertificateConfigurationId, err := bdsInstanceBdsCertificateConfigurationWaitForWorkRequest(ctx, workId, "bdscertificateconfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(GetBdsInstanceBdsCertificateConfigurationCompositeId(*bdsInstanceBdsCertificateConfigurationId, s.D.Get("bds_instance_id").(string)))

	return s.GetWithContext(ctx)
}

func bdsInstanceBdsCertificateConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceBdsCertificateConfigurationWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceBdsCertificateConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_bds.GetWorkRequestRequest{
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
		if res.EntityType == nil {
			continue
		}
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			identifier = res.Identifier
			break
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceBdsCertificateConfigurationWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceBdsCertificateConfigurationWorkRequest(ctx context.Context, client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_bds.ListWorkRequestErrorsRequest{
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

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_bds.GetBdsCertificateConfigurationRequest{}

	if bdsCertificateConfigurationId, ok := s.D.GetOkExists("bds_certificate_configuration_id"); ok {
		tmp := bdsCertificateConfigurationId.(string)
		request.BdsCertificateConfigurationId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsCertificateConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCertificateConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCertificateConfigurationId = &bdsCertificateConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetBdsCertificateConfiguration(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsCertificateConfiguration
	return nil
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_bds.DeleteBdsCertificateConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsCertificateConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCertificateConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCertificateConfigurationId = &bdsCertificateConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteBdsCertificateConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceBdsCertificateConfigurationWaitForWorkRequest(ctx, workId, "bdscertificateconfig",
		oci_bds.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceBdsCertificateConfigurationResourceCrud) SetData() error {

	_, bdsInstanceId, err := parseBdsInstanceBdsCertificateConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BdsInstanceId != nil {
		s.D.Set("bds_instance_id", *s.Res.BdsInstanceId)
	}

	if s.Res.CertificateAuthorityId != nil {
		s.D.Set("certificate_authority_id", *s.Res.CertificateAuthorityId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsDefaultConfiguration != nil {
		s.D.Set("is_default_configuration", *s.Res.IsDefaultConfiguration)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRefreshedOrGenerated != nil {
		s.D.Set("time_last_refreshed_or_generated", s.Res.TimeLastRefreshedOrGenerated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func GetBdsInstanceBdsCertificateConfigurationCompositeId(bdsCertificateConfigurationId string, bdsInstanceId string) string {
	bdsCertificateConfigurationId = url.PathEscape(bdsCertificateConfigurationId)
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/bdsCertificateConfigurations/" + bdsCertificateConfigurationId
	return compositeId
}

func isActiveTrigger(v string) bool {
	val := strings.TrimSpace(strings.ToLower(v))
	return val != "" && val != "false"
}

func updateBdsBdsInstanceBdsCertificateConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCertificateConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	runIssue := false
	runRenew := false
	runSetDefault := false

	if triggerVal, ok := sync.D.GetOkExists("issue_certificate_trigger"); ok {
		if isActiveTrigger(triggerVal.(string)) && (sync.D.HasChange("issue_certificate_trigger") || sync.D.HasChange("is_missing_nodes_only")) {
			runIssue = true
		}
	}

	if triggerVal, ok := sync.D.GetOkExists("renew_certificate_trigger"); ok && sync.D.HasChange("renew_certificate_trigger") {
		if isActiveTrigger(triggerVal.(string)) {
			runRenew = true
		}
	}

	if triggerVal, ok := sync.D.GetOkExists("set_default_trigger"); ok && sync.D.HasChange("set_default_trigger") {
		if isActiveTrigger(triggerVal.(string)) {
			runSetDefault = true
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	if runIssue {
		if err := sync.IssueCertificate(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if runRenew {
		if err := sync.RenewCertificate(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if runSetDefault {
		if err := sync.SetDefaultCertificateConfiguration(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	return nil
}

func parseBdsInstanceBdsCertificateConfigurationCompositeId(compositeId string) (bdsCertificateConfigurationId string, bdsInstanceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/bdsCertificateConfigurations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	bdsCertificateConfigurationId, _ = url.PathUnescape(parts[3])

	return
}
