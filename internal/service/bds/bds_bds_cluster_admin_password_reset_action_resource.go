package bds

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BdsBdsClusterAdminPasswordResetActionResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Create: createBdsBdsClusterAdminPasswordResetAction,
		Update: updateBdsBdsClusterAdminPasswordResetAction,
		Read:   readBdsBdsClusterAdminPasswordResetAction,
		Delete: deleteBdsBdsClusterAdminPasswordResetAction,
		Schema: map[string]*schema.Schema{
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"current_cluster_admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"reset_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func createBdsBdsClusterAdminPasswordResetAction(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsClusterAdminPasswordResetActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	return tfresource.CreateResource(d, sync)
}

func updateBdsBdsClusterAdminPasswordResetAction(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsClusterAdminPasswordResetActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	return tfresource.CreateResource(d, sync)
}

func readBdsBdsClusterAdminPasswordResetAction(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteBdsBdsClusterAdminPasswordResetAction(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

type BdsBdsClusterAdminPasswordResetActionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	DisableNotFoundRetries bool
}

func (s *BdsBdsClusterAdminPasswordResetActionResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("BdsBdsClusterAdminPasswordResetActionResource-", BdsBdsClusterAdminPasswordResetActionResource(), s.D)
}

func (s *BdsBdsClusterAdminPasswordResetActionResourceCrud) Create() error {
	request := oci_bds.BdsInstanceResetPasswordRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	} else {
		return fmt.Errorf("bds_instance_id is required")
	}

	if service, ok := s.D.GetOkExists("service"); ok {
		serviceValue := strings.ToUpper(service.(string))
		request.Service = oci_bds.BdsInstanceResetPasswordDetailsServiceEnum(serviceValue)
	} else {
		return fmt.Errorf("service is required")
	}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		request.SecretId = &tmp
	}

	if currentPassword, ok := s.D.GetOkExists("current_cluster_admin_password"); ok {
		tmp := currentPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.BdsInstanceResetPassword(context.Background(), request)
	if err != nil {
		return err
	}

	if response.AdminPassword != nil {
		if err := s.D.Set("cluster_admin_password", *response.AdminPassword); err != nil {
			return fmt.Errorf("error setting cluster_admin_password: %v", err)
		}
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		if err := s.getBdsInstanceResetPasswordFromWorkRequest(
			workId,
			tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"),
			oci_bds.ActionTypesUpdated,
			s.D.Timeout(schema.TimeoutCreate),
		); err != nil {
			return err
		}
	} else {
		s.D.SetId(s.ID())
	}
	return nil
}

func (s *BdsBdsClusterAdminPasswordResetActionResourceCrud) getBdsInstanceResetPasswordFromWorkRequest(
	workId *string,
	retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum,
	timeout time.Duration,
) error {

	resetOpId, err := bdsInstanceResetPasswordWaitForWorkRequest(
		workId, "bds", actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	if resetOpId != nil {
		s.D.SetId(*resetOpId)
	} else {
		s.D.SetId(s.ID())
	}
	return nil
}

func bdsInstanceResetPasswordWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)

	return func(response oci_common.OCIOperationResponse) bool {
		if time.Now().After(stopTime) {
			return false
		}
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}
		if wr, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return wr.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceResetPasswordWaitForWorkRequest(
	wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient,
) (*string, error) {

	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceResetPasswordWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	state := &retry.StateChangeConf{
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
			response, err = client.GetWorkRequest(context.Background(),
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
	if _, e := state.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) &&
			res.ActionType == action {
			identifier = res.Identifier
			break
		}
	}

	if identifier == nil ||
		response.Status == oci_bds.OperationStatusFailed ||
		response.Status == oci_bds.OperationStatusCanceled {

		return nil, getErrorFromBdsBdsClusterAdminPasswordResetWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsClusterAdminPasswordResetWorkRequest(
	client *oci_bds.BdsClient, workId *string,
	retryPolicy *oci_common.RetryPolicy, entityType string,
	action oci_bds.ActionTypesEnum,
) error {

	response, err := client.ListWorkRequestErrors(context.Background(),
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

	return fmt.Errorf(
		"work request failed, workId: %s, entity: %s, action: %s. Errors: %s",
		*workId, entityType, action, strings.Join(allErrs, "\n"))
}

func (s *BdsBdsClusterAdminPasswordResetActionResourceCrud) SetData() error {
	if val, ok := s.D.GetOkExists("cluster_admin_password"); ok && val != nil {
		_ = s.D.Set("cluster_admin_password", val.(string))
	}
	return nil
}
