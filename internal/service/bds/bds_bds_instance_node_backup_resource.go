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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceNodeBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsInstanceNodeBackup,
		Read:     readBdsBdsInstanceNodeBackup,
		Delete:   deleteBdsBdsInstanceNodeBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"node_instance_id": {
				Type: schema.TypeString,
				//Required: true,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backup_config_id": {
				Type: schema.TypeString,
				//	Required: false,
				ForceNew: true,
				Optional: true,
			},
			"backup_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type: schema.TypeString,
				//	Required: true,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"level_type_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"level_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NODE_LEVEL",
								"NODE_TYPE_LEVEL",
							}, true),
						},

						// Optional
						"node_host_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"node_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBdsBdsInstanceNodeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.CreateResource(d, sync)
}

func readBdsBdsInstanceNodeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

/*func updateBdsBdsInstanceNodeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.UpdateResource(d, sync)
}*/

func deleteBdsBdsInstanceNodeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceNodeBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.NodeBackup
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) ID() string {
	return GetBdsInstanceNodeBackupCompositeId(s.D.Get("bds_instance_id").(string), *s.Res.Id)
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.NodeBackupLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.NodeBackupLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.NodeBackupLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.NodeBackupLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) Create() error {
	request := oci_bds.BackupNodeRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if backupType, ok := s.D.GetOkExists("backup_type"); ok {
		request.BackupType = oci_bds.NodeBackupBackupTypeEnum(backupType.(string))
	}

	if levelTypeDetails, ok := s.D.GetOkExists("level_type_details"); ok {
		if tmpList := levelTypeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "level_type_details", 0)
			tmp, err := s.mapToLevelTypeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LevelTypeDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.BackupNode(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceNodeBackupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))

}

func (s *BdsBdsInstanceNodeBackupResourceCrud) getBdsInstanceNodeBackupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {
	// Wait until it finishes
	bdsInstanceNodeBackupId, err := bdsInstanceNodeBackupWaitForWorkRequest(workId, "nodeBackup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}
	for _, res := range bdsInstanceNodeBackupId {
		s.D.SetId(res)
	}

	return s.Get()
}

func bdsInstanceNodeBackupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func bdsInstanceNodeBackupWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) ([]string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceNodeBackupWorkRequestShouldRetryFunc(timeout)

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
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	identifier := make([]string, 0)
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(*res.EntityType, entityType) {
			if res.ActionType == action {
				identifier = append(identifier, *res.Identifier)
			}
		}
	}
	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceNodeBackupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceNodeBackupWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
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
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) Get() error {
	request := oci_bds.GetNodeBackupRequest{}
	tmp := s.D.Id()
	request.NodeBackupId = &tmp
	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}
	bdsInstanceId, nodeBackupId, err := parseBdsInstanceNodeBackupCompositeId(s.D.Id())
	if err == nil {
		request.BdsInstanceId = &bdsInstanceId
		request.NodeBackupId = &nodeBackupId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetNodeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NodeBackup
	return nil
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) Delete() error {
	request := oci_bds.DeleteNodeBackupRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	tmp := s.D.Id()
	request.NodeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteNodeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceNodeBackupWaitForWorkRequest(workId, "nodeBackup",
		oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) SetData() error {

	bdsInstanceId, nodeBackupId, err := parseBdsInstanceNodeBackupCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
		s.D.SetId(nodeBackupId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}
	s.D.Set("backup_type", s.Res.BackupType)
	if s.Res.NodeInstanceId != nil {
		s.D.Set("node_instance_id", *s.Res.NodeInstanceId)
	}
	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("state", s.Res.LifecycleState)
	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.NodeBackupConfigId != nil {
		s.D.Set("backup_config_id", s.Res.NodeBackupConfigId)
	}
	return nil
}

func GetBdsInstanceNodeBackupCompositeId(bdsInstanceId string, nodeBackupId string) string {
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	nodeBackupId = url.PathEscape(nodeBackupId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/nodeBackups/" + nodeBackupId
	return compositeId
}

func parseBdsInstanceNodeBackupCompositeId(compositeId string) (bdsInstanceId string, nodeBackupId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/nodeBackups/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	nodeBackupId, _ = url.PathUnescape(parts[3])

	return
}

func (s *BdsBdsInstanceNodeBackupResourceCrud) mapToLevelTypeDetails(fieldKeyFormat string) (oci_bds.LevelTypeDetails, error) {
	var baseObject oci_bds.LevelTypeDetails
	//discriminator
	levelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "level_type"))
	var levelType string
	if ok {
		levelType = levelTypeRaw.(string)
	} else {
		levelType = "" // default value
	}
	switch strings.ToLower(levelType) {
	case strings.ToLower("NODE_LEVEL"):
		details := oci_bds.NodeLevelDetails{}
		if nodeHostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_host_name")); ok {
			tmp := nodeHostName.(string)
			details.NodeHostName = &tmp
		}
		baseObject = details
	case strings.ToLower("NODE_TYPE_LEVEL"):
		details := oci_bds.NodeTypeLevelDetails{}
		if nodeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_type")); ok {
			details.NodeType = oci_bds.NodeNodeTypeEnum(nodeType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown level_type '%v' was specified", levelType)
	}
	return baseObject, nil
}
