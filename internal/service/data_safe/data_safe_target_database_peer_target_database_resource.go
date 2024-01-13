// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetDatabasePeerTargetDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeTargetDatabasePeerTargetDatabase,
		Read:     readDataSafeTargetDatabasePeerTargetDatabase,
		Update:   updateDataSafeTargetDatabasePeerTargetDatabase,
		Delete:   deleteDataSafeTargetDatabasePeerTargetDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"database_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"database_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AUTONOMOUS_DATABASE",
								"DATABASE_CLOUD_SERVICE",
								"INSTALLED_DATABASE",
							}, true),
						},
						"infrastructure_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"autonomous_database_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"db_system_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_addresses": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"listener_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vm_cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"target_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"dataguard_association_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"status": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"certificate_store_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_store_content": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"store_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
						"trust_store_content": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"database_unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeTargetDatabasePeerTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasePeerTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeTargetDatabasePeerTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasePeerTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeTargetDatabasePeerTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasePeerTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeTargetDatabasePeerTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasePeerTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeTargetDatabasePeerTargetDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.PeerTargetDatabase
	DisableNotFoundRetries bool
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) ID() string {
	peerTarget := *s.Res
	s1 := strconv.FormatInt(int64(*peerTarget.Key), 10)
	return GetTargetDatabasePeerTargetDatabaseCompositeId(s1, s.D.Get("target_database_id").(string))
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateCreating),
	}
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateActive),
		string(oci_data_safe.TargetDatabaseLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateDeleting),
	}
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateDeleted),
	}
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) Create() error {
	request := oci_data_safe.CreatePeerTargetDatabaseRequest{}

	if databaseDetails, ok := s.D.GetOkExists("database_details"); ok {
		if tmpList := databaseDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_details", 0)
			tmp, err := s.mapToDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseDetails = tmp
		}
	}

	if dataguardAssociationId, ok := s.D.GetOkExists("dataguard_association_id"); ok {
		tmp := dataguardAssociationId.(string)
		request.DataguardAssociationId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	if tlsConfig, ok := s.D.GetOkExists("tls_config"); ok {
		if tmpList := tlsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_config", 0)
			tmp, err := s.mapToTlsConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TlsConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreatePeerTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTargetDatabasePeerTargetDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) getTargetDatabasePeerTargetDatabaseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	targetDatabasePeerTargetDatabaseId, err := targetDatabasePeerTargetDatabaseWaitForWorkRequest(workId, "peer-target-database",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, targetDatabasePeerTargetDatabaseId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_data_safe.CancelWorkRequestRequest{
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
	s.D.SetId(*targetDatabasePeerTargetDatabaseId)

	return s.Get()
}

func targetDatabasePeerTargetDatabaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func targetDatabasePeerTargetDatabaseWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = targetDatabasePeerTargetDatabaseWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_safe.GetWorkRequestRequest{
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
				identifier = res.EntityUri
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeTargetDatabasePeerTargetDatabaseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeTargetDatabasePeerTargetDatabaseWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) Get() error {
	request := oci_data_safe.GetPeerTargetDatabaseRequest{}

	peerTargetDatabaseId, targetDatabaseId, err := parseTargetDatabasePeerTargetDatabaseCompositeId(s.D.Id())
	if err == nil {
		request.PeerTargetDatabaseId = &peerTargetDatabaseId
		request.TargetDatabaseId = &targetDatabaseId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetPeerTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PeerTargetDatabase
	return nil
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) Update() error {
	request := oci_data_safe.UpdatePeerTargetDatabaseRequest{}

	if databaseDetails, ok := s.D.GetOkExists("database_details"); ok {
		if tmpList := databaseDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_details", 0)
			tmp, err := s.mapToDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseDetails = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if peerTargetDatabaseId, ok := s.D.GetOkExists("key"); ok {
		tmp := peerTargetDatabaseId.(int)
		request.PeerTargetDatabaseId = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	if tlsConfig, ok := s.D.GetOkExists("tls_config"); ok {
		if tmpList := tlsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_config", 0)
			tmp, err := s.mapToTlsConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TlsConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdatePeerTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTargetDatabasePeerTargetDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) Delete() error {
	request := oci_data_safe.DeletePeerTargetDatabaseRequest{}

	if peerTargetDatabaseId, ok := s.D.GetOkExists("key"); ok {
		tmp := peerTargetDatabaseId.(int)
		request.PeerTargetDatabaseId = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeletePeerTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := targetDatabasePeerTargetDatabaseWaitForWorkRequest(workId, "peer-target-database",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) SetData() error {

	peerTargetDatabaseId, targetDatabaseId, err := parseTargetDatabasePeerTargetDatabaseCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &peerTargetDatabaseId)
		s.D.Set("target_database_id", &targetDatabaseId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DatabaseDetails != nil {
		databaseDetailsArray := []interface{}{}
		if databaseDetailsMap := DatabaseDetailsToMap(&s.Res.DatabaseDetails); databaseDetailsMap != nil {
			databaseDetailsArray = append(databaseDetailsArray, databaseDetailsMap)
		}
		s.D.Set("database_details", databaseDetailsArray)
	} else {
		s.D.Set("database_details", nil)
	}

	if s.Res.DatabaseUniqueName != nil {
		s.D.Set("database_unique_name", *s.Res.DatabaseUniqueName)
	}

	if s.Res.DataguardAssociationId != nil {
		s.D.Set("dataguard_association_id", *s.Res.DataguardAssociationId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Role != nil {
		s.D.Set("role", *s.Res.Role)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TlsConfig != nil {
		s.D.Set("tls_config", []interface{}{TlsConfigToMap(s.Res.TlsConfig)})
	} else {
		s.D.Set("tls_config", nil)
	}

	return nil
}

func GetTargetDatabasePeerTargetDatabaseCompositeId(peerTargetDatabaseId string, targetDatabaseId string) string {
	peerTargetDatabaseId = url.PathEscape(peerTargetDatabaseId)
	targetDatabaseId = url.PathEscape(targetDatabaseId)
	compositeId := "targetDatabases/" + targetDatabaseId + "/peerTargetDatabases/" + peerTargetDatabaseId
	return compositeId
}

func parseTargetDatabasePeerTargetDatabaseCompositeId(compositeId string) (peerTargetDatabaseId int, targetDatabaseId string, err error) {
	firstChar := compositeId[0:1]
	var compositeIdStr string
	if firstChar == "/" {
		compositeIdStr = trimLeftChar(compositeId)
	} else {
		compositeIdStr = compositeId
	}
	parts := strings.Split(compositeIdStr, "/")
	match, _ := regexp.MatchString("targetDatabases/.*/peerTargetDatabases/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	targetDatabaseId, _ = url.PathUnescape(parts[1])
	peerTargetDatabaseId, _ = strconv.Atoi(parts[3])

	return
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) mapToDatabaseDetails(fieldKeyFormat string) (oci_data_safe.DatabaseDetails, error) {
	var baseObject oci_data_safe.DatabaseDetails
	//discriminator
	databaseTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_type"))
	var databaseType string
	if ok {
		databaseType = databaseTypeRaw.(string)
	} else {
		databaseType = "" // default value
	}
	switch strings.ToLower(databaseType) {
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_data_safe.AutonomousDatabaseDetails{}
		if autonomousDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autonomous_database_id")); ok {
			tmp := autonomousDatabaseId.(string)
			details.AutonomousDatabaseId = &tmp
		}
		if infrastructureType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "infrastructure_type")); ok {
			details.InfrastructureType = oci_data_safe.InfrastructureTypeEnum(infrastructureType.(string))
		}
		baseObject = details
	case strings.ToLower("DATABASE_CLOUD_SERVICE"):
		details := oci_data_safe.DatabaseCloudServiceDetails{}
		if dbSystemId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_system_id")); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if listenerPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_port")); ok {
			tmp := listenerPort.(int)
			details.ListenerPort = &tmp
		}
		if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		if infrastructureType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "infrastructure_type")); ok {
			details.InfrastructureType = oci_data_safe.InfrastructureTypeEnum(infrastructureType.(string))
		}
		baseObject = details
	case strings.ToLower("INSTALLED_DATABASE"):
		details := oci_data_safe.InstalledDatabaseDetails{}
		if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if ipAddresses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_addresses")); ok {
			interfaces := ipAddresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ip_addresses")) {
				details.IpAddresses = tmp
			}
		}
		if listenerPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_port")); ok {
			tmp := listenerPort.(int)
			details.ListenerPort = &tmp
		}
		if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if infrastructureType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "infrastructure_type")); ok {
			details.InfrastructureType = oci_data_safe.InfrastructureTypeEnum(infrastructureType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown database_type '%v' was specified", databaseType)
	}
	return baseObject, nil
}

func PeerTargetDatabaseSummaryToMap(obj oci_data_safe.PeerTargetDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DatabaseUniqueName != nil {
		result["database_unique_name"] = string(*obj.DatabaseUniqueName)
	}

	if obj.DataguardAssociationId != nil {
		result["dataguard_association_id"] = string(*obj.DataguardAssociationId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Key != nil {
		result["key"] = int(*obj.Key)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Role != nil {
		result["role"] = string(*obj.Role)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseResourceCrud) mapToTlsConfig(fieldKeyFormat string) (oci_data_safe.TlsConfig, error) {
	result := oci_data_safe.TlsConfig{}

	if certificateStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_store_type")); ok {
		result.CertificateStoreType = oci_data_safe.TlsConfigCertificateStoreTypeEnum(certificateStoreType.(string))
	}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		tmp := keyStoreContent.(string)
		result.KeyStoreContent = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_data_safe.TlsConfigStatusEnum(status.(string))
	}

	if storePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "store_password")); ok {
		tmp := storePassword.(string)
		result.StorePassword = &tmp
	}

	if trustStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trust_store_content")); ok {
		tmp := trustStoreContent.(string)
		result.TrustStoreContent = &tmp
	}

	return result, nil
}
