// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

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
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentAdminUserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createFusionAppsFusionEnvironmentAdminUserWithContext,
		ReadContext:   readFusionAppsFusionEnvironmentAdminUserWithContext,
		DeleteContext: deleteFusionAppsFusionEnvironmentAdminUserWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"email_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"email_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createFusionAppsFusionEnvironmentAdminUserWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FusionAppsFusionEnvironmentAdminUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readFusionAppsFusionEnvironmentAdminUserWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FusionAppsFusionEnvironmentAdminUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteFusionAppsFusionEnvironmentAdminUserWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FusionAppsFusionEnvironmentAdminUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type FusionAppsFusionEnvironmentAdminUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fusion_apps.FusionApplicationsClient
	Res                    *oci_fusion_apps.AdminUserCollection
	DisableNotFoundRetries bool
}

func (s *FusionAppsFusionEnvironmentAdminUserResourceCrud) ID() string {
	return GetFusionEnvironmentAdminUserCompositeId(s.D.Get("username").(string), s.D.Get("fusion_environment_id").(string))
}

func (s *FusionAppsFusionEnvironmentAdminUserResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_fusion_apps.CreateFusionEnvironmentAdminUserRequest{}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	if firstName, ok := s.D.GetOkExists("first_name"); ok {
		tmp := firstName.(string)
		request.FirstName = &tmp
	}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if lastName, ok := s.D.GetOkExists("last_name"); ok {
		tmp := lastName.(string)
		request.LastName = &tmp
	}

	if username, ok := s.D.GetOkExists("username"); ok {
		tmp := username.(string)
		request.Username = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.CreateFusionEnvironmentAdminUser(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.D.SetId(s.ID())
	return s.getFusionEnvironmentAdminUserFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FusionAppsFusionEnvironmentAdminUserResourceCrud) getFusionEnvironmentAdminUserFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fusion_apps.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fusionEnvironmentAdminUserId, err := fusionEnvironmentAdminUserWaitForWorkRequest(ctx, workId, "fusionenvironment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fusionEnvironmentAdminUserId)

	return s.GetWithContext(ctx)
}

func fusionEnvironmentAdminUserWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fusion_apps", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fusion_apps.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fusionEnvironmentAdminUserWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fusion_apps.FusionApplicationsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fusion_apps")
	retryPolicy.ShouldRetryOperation = fusionEnvironmentAdminUserWorkRequestShouldRetryFunc(timeout)

	response := oci_fusion_apps.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_fusion_apps.WorkRequestStatusInProgress),
			string(oci_fusion_apps.WorkRequestStatusAccepted),
			string(oci_fusion_apps.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_fusion_apps.WorkRequestStatusSucceeded),
			string(oci_fusion_apps.WorkRequestStatusFailed),
			string(oci_fusion_apps.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_fusion_apps.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fusion_apps.WorkRequestStatusFailed || response.Status == oci_fusion_apps.WorkRequestStatusCanceled {
		return nil, getErrorFromFusionAppsFusionEnvironmentAdminUserWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFusionAppsFusionEnvironmentAdminUserWorkRequest(ctx context.Context, client *oci_fusion_apps.FusionApplicationsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_fusion_apps.ListWorkRequestErrorsRequest{
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

func (s *FusionAppsFusionEnvironmentAdminUserResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fusion_apps.ListAdminUsersRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	_, fusionEnvironmentId, err := parseFusionEnvironmentAdminUserCompositeId(s.D.Id())
	if err == nil {
		request.FusionEnvironmentId = &fusionEnvironmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.ListAdminUsers(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.AdminUserCollection
	return nil
}

func (s *FusionAppsFusionEnvironmentAdminUserResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_fusion_apps.DeleteFusionEnvironmentAdminUserRequest{}

	if adminUsername, ok := s.D.GetOkExists("username"); ok {
		tmp := adminUsername.(string)
		request.AdminUsername = &tmp
	}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.DeleteFusionEnvironmentAdminUser(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fusionEnvironmentAdminUserWaitForWorkRequest(ctx, workId, "fusionenvironment",
		oci_fusion_apps.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *FusionAppsFusionEnvironmentAdminUserResourceCrud) SetData() error {

	adminUsername, fusionEnvironmentId, err := parseFusionEnvironmentAdminUserCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("username", &adminUsername)
		s.D.Set("fusion_environment_id", &fusionEnvironmentId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdminUserSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func GetFusionEnvironmentAdminUserCompositeId(adminUsername string, fusionEnvironmentId string) string {
	adminUsername = url.PathEscape(adminUsername)
	fusionEnvironmentId = url.PathEscape(fusionEnvironmentId)
	compositeId := "fusionEnvironments/" + fusionEnvironmentId + "/adminUsers/" + adminUsername
	return compositeId
}

func parseFusionEnvironmentAdminUserCompositeId(compositeId string) (adminUsername string, fusionEnvironmentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fusionEnvironments/.*/adminUsers/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fusionEnvironmentId, _ = url.PathUnescape(parts[1])
	adminUsername, _ = url.PathUnescape(parts[3])

	return
}

func AdminUserSummaryToMap(obj oci_fusion_apps.AdminUserSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmailAddress != nil {
		result["email_address"] = string(*obj.EmailAddress)
	}

	if obj.FirstName != nil {
		result["first_name"] = string(*obj.FirstName)
	}

	if obj.LastName != nil {
		result["last_name"] = string(*obj.LastName)
	}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}
