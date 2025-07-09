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

func BdsBdsInstanceNodeBackupConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsInstanceNodeBackupConfiguration,
		Read:     readBdsBdsInstanceNodeBackupConfiguration,
		Update:   updateBdsBdsInstanceNodeBackupConfiguration,
		Delete:   deleteBdsBdsInstanceNodeBackupConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"level_type_details": {
				Type:     schema.TypeList,
				Required: true,
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
			"schedule": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"backup_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"number_of_backups_to_retain": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBdsBdsInstanceNodeBackupConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.CreateResource(d, sync)
}

func readBdsBdsInstanceNodeBackupConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

func updateBdsBdsInstanceNodeBackupConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBdsBdsInstanceNodeBackupConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceNodeBackupConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.NodeBackupConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) ID() string {
	return GetBdsInstanceNodeBackupConfigurationCompositeId(s.D.Get("bds_instance_id").(string), *s.Res.Id)
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.NodeBackupConfigurationLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.NodeBackupConfigurationLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.NodeBackupConfigurationLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.NodeBackupConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) Create() error {
	request := oci_bds.CreateNodeBackupConfigurationRequest{}

	if backupType, ok := s.D.GetOkExists("backup_type"); ok {
		request.BackupType = oci_bds.NodeBackupBackupTypeEnum(backupType.(string))
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
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

	if numberOfBackupsToRetain, ok := s.D.GetOkExists("number_of_backups_to_retain"); ok {
		tmp := numberOfBackupsToRetain.(int)
		request.NumberOfBackupsToRetain = &tmp
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateNodeBackupConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceNodeBackupConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) getBdsInstanceNodeBackupConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceNodeBackupConfigurationId, err := bdsInstanceNodeBackupConfigurationWaitForWorkRequest(workId, "backupConfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceNodeBackupConfigurationId)

	return s.Get()
}

func bdsInstanceNodeBackupConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func bdsInstanceNodeBackupConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceNodeBackupConfigurationWorkRequestShouldRetryFunc(timeout)

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

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(*res.EntityType, entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}
	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceNodeBackupConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceNodeBackupConfigurationWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
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

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) Get() error {
	request := oci_bds.GetNodeBackupConfigurationRequest{}
	tmp := s.D.Id()
	request.NodeBackupConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsInstanceId, nodeBackupConfigurationId, err := parseBdsInstanceNodeBackupConfigurationCompositeId(s.D.Id())

	if err == nil {
		request.BdsInstanceId = &bdsInstanceId
		request.NodeBackupConfigurationId = &nodeBackupConfigurationId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetNodeBackupConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NodeBackupConfiguration
	return nil
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) Update() error {
	request := oci_bds.UpdateNodeBackupConfigurationRequest{}

	if backupType, ok := s.D.GetOkExists("backup_type"); ok {
		request.BackupType = oci_bds.NodeBackupBackupTypeEnum(backupType.(string))
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
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

	tmp := s.D.Id()
	request.NodeBackupConfigurationId = &tmp

	if numberOfBackupsToRetain, ok := s.D.GetOkExists("number_of_backups_to_retain"); ok {
		tmp := numberOfBackupsToRetain.(int)
		request.NumberOfBackupsToRetain = &tmp
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateNodeBackupConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceNodeBackupConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) Delete() error {
	request := oci_bds.DeleteNodeBackupConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	tmp := s.D.Id()
	request.NodeBackupConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteNodeBackupConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceNodeBackupConfigurationWaitForWorkRequest(workId, "backupConfig",
		oci_bds.ActionTypesInProgress, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) SetData() error {

	bdsInstanceId, nodeBackupConfigurationId, err := parseBdsInstanceNodeBackupConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
		s.D.SetId(nodeBackupConfigurationId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.BdsInstanceId != nil {
		s.D.Set("bds_instance_id", *s.Res.BdsInstanceId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LevelTypeDetails != nil {
		levelTypeDetailsArray := []interface{}{}
		if levelTypeDetailsMap := LevelTypeDetailsToMap(&s.Res.LevelTypeDetails); levelTypeDetailsMap != nil {
			levelTypeDetailsArray = append(levelTypeDetailsArray, levelTypeDetailsMap)
		}
		s.D.Set("level_type_details", levelTypeDetailsArray)
	} else {
		s.D.Set("level_type_details", nil)
	}

	if s.Res.NumberOfBackupsToRetain != nil {
		s.D.Set("number_of_backups_to_retain", *s.Res.NumberOfBackupsToRetain)
	}

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	return nil
}

func GetBdsInstanceNodeBackupConfigurationCompositeId(bdsInstanceId string, nodeBackupConfigurationId string) string {
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	nodeBackupConfigurationId = url.PathEscape(nodeBackupConfigurationId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/nodeBackupConfigurations/" + nodeBackupConfigurationId
	return compositeId
}

func parseBdsInstanceNodeBackupConfigurationCompositeId(compositeId string) (bdsInstanceId string, nodeBackupConfigurationId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/nodeBackupConfigurations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	nodeBackupConfigurationId, _ = url.PathUnescape(parts[3])

	return
}

func (s *BdsBdsInstanceNodeBackupConfigurationResourceCrud) mapToLevelTypeDetails(fieldKeyFormat string) (oci_bds.LevelTypeDetails, error) {
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

func LevelTypeDetailsToMap(obj *oci_bds.LevelTypeDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_bds.NodeLevelDetails:
		result["level_type"] = "NODE_LEVEL"

		if v.NodeHostName != nil {
			result["node_host_name"] = string(*v.NodeHostName)
		}
	case oci_bds.NodeTypeLevelDetails:
		result["level_type"] = "NODE_TYPE_LEVEL"

		result["node_type"] = string(v.NodeType)
	default:
		log.Printf("[WARN] Received 'level_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
