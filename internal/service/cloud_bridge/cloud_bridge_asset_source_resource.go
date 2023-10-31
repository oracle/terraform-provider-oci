// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAssetSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudBridgeAssetSource,
		Read:     readCloudBridgeAssetSource,
		Update:   updateCloudBridgeAssetSource,
		Delete:   deleteCloudBridgeAssetSource,
		Schema: map[string]*schema.Schema{
			// Required
			"assets_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"discovery_credentials": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"secret_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"inventory_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"VMWARE",
				}, true),
			},
			"vcenter_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"are_historical_metrics_collected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"are_realtime_metrics_collected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"discovery_schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"replication_credentials": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"secret_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_details": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCloudBridgeAssetSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudBridgeAssetSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.ReadResource(sync)
}

func updateCloudBridgeAssetSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudBridgeAssetSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudBridgeAssetSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_bridge.DiscoveryClient
	WorkRequestClient      *oci_cloud_bridge.CommonClient
	Res                    *oci_cloud_bridge.AssetSource
	DisableNotFoundRetries bool
}

func (s *CloudBridgeAssetSourceResourceCrud) ID() string {
	assetSource := *s.Res
	return *assetSource.GetId()
}

func (s *CloudBridgeAssetSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_bridge.AssetSourceLifecycleStateCreating),
	}
}

func (s *CloudBridgeAssetSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AssetSourceLifecycleStateActive),
		string(oci_cloud_bridge.AssetSourceLifecycleStateNeedsAttention),
	}
}

func (s *CloudBridgeAssetSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_bridge.AssetSourceLifecycleStateDeleting),
	}
}

func (s *CloudBridgeAssetSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AssetSourceLifecycleStateDeleted),
	}
}

func (s *CloudBridgeAssetSourceResourceCrud) Create() error {
	request := oci_cloud_bridge.CreateAssetSourceRequest{}
	err := s.populateTopLevelPolymorphicCreateAssetSourceRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.CreateAssetSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAssetSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge"), oci_cloud_bridge.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *CloudBridgeAssetSourceResourceCrud) getAssetSourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_bridge.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	assetSourceId, err := assetSourceWaitForWorkRequest(workId, "ocbassetsource",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, assetSourceId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_cloud_bridge.CancelWorkRequestRequest{
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
	s.D.SetId(*assetSourceId)

	return s.Get()
}

func assetSourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "cloud_bridge", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_cloud_bridge.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func assetSourceWaitForWorkRequest(wId *string, entityType string, action oci_cloud_bridge.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_bridge.CommonClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_bridge")
	retryPolicy.ShouldRetryOperation = assetSourceWorkRequestShouldRetryFunc(timeout)

	response := oci_cloud_bridge.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_cloud_bridge.OperationStatusInProgress),
			string(oci_cloud_bridge.OperationStatusAccepted),
			string(oci_cloud_bridge.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_cloud_bridge.OperationStatusSucceeded),
			string(oci_cloud_bridge.OperationStatusFailed),
			string(oci_cloud_bridge.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_cloud_bridge.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_cloud_bridge.OperationStatusFailed || response.Status == oci_cloud_bridge.OperationStatusCanceled {
		return nil, getErrorFromCloudBridgeAssetSourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudBridgeAssetSourceWorkRequest(client *oci_cloud_bridge.CommonClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_bridge.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_cloud_bridge.ListWorkRequestErrorsRequest{
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

func (s *CloudBridgeAssetSourceResourceCrud) Get() error {
	request := oci_cloud_bridge.GetAssetSourceRequest{}

	tmp := s.D.Id()
	request.AssetSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.GetAssetSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AssetSource
	return nil
}

func (s *CloudBridgeAssetSourceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_bridge.UpdateAssetSourceRequest{}
	err := s.populateTopLevelPolymorphicUpdateAssetSourceRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.UpdateAssetSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAssetSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge"), oci_cloud_bridge.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *CloudBridgeAssetSourceResourceCrud) Delete() error {
	request := oci_cloud_bridge.DeleteAssetSourceRequest{}

	tmp := s.D.Id()
	request.AssetSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.DeleteAssetSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := assetSourceWaitForWorkRequest(workId, "ocbassetsource",
		oci_cloud_bridge.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *CloudBridgeAssetSourceResourceCrud) SetData() error {

	switch v := (*s.Res).(type) {
	case oci_cloud_bridge.VmWareAssetSource:
		s.D.Set("type", "VMWARE")

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.AreHistoricalMetricsCollected != nil {
			s.D.Set("are_historical_metrics_collected", *v.AreHistoricalMetricsCollected)
		}

		if v.AreRealtimeMetricsCollected != nil {
			s.D.Set("are_realtime_metrics_collected", *v.AreRealtimeMetricsCollected)
		}

		if v.DiscoveryCredentials != nil {
			s.D.Set("discovery_credentials", []interface{}{AssetSourceCredentialsToMap(v.DiscoveryCredentials)})
		} else {
			s.D.Set("discovery_credentials", nil)
		}

		if v.ReplicationCredentials != nil {
			s.D.Set("replication_credentials", []interface{}{AssetSourceCredentialsToMap(v.ReplicationCredentials)})
		} else {
			s.D.Set("replication_credentials", nil)
		}

		if v.VcenterEndpoint != nil {
			s.D.Set("vcenter_endpoint", *v.VcenterEndpoint)
		}

		if v.AssetsCompartmentId != nil {
			s.D.Set("assets_compartment_id", *v.AssetsCompartmentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DiscoveryScheduleId != nil {
			s.D.Set("discovery_schedule_id", *v.DiscoveryScheduleId)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.EnvironmentId != nil {
			s.D.Set("environment_id", *v.EnvironmentId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *CloudBridgeAssetSourceResourceCrud) mapToAssetSourceCredentials(fieldKeyFormat string) (oci_cloud_bridge.AssetSourceCredentials, error) {
	result := oci_cloud_bridge.AssetSourceCredentials{}

	if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
		tmp := secretId.(string)
		result.SecretId = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_cloud_bridge.AssetSourceCredentialsTypeEnum(type_.(string))
	}

	return result, nil
}

func AssetSourceCredentialsToMap(obj *oci_cloud_bridge.AssetSourceCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SecretId != nil {
		result["secret_id"] = string(*obj.SecretId)
	}

	result["type"] = string(obj.Type)

	return result
}

func AssetSourceSummaryToMap(obj oci_cloud_bridge.AssetSourceSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_cloud_bridge.VmWareAssetSourceSummary:
		result["type"] = "VMWARE"

		if v.VcenterEndpoint != nil {
			result["vcenter_endpoint"] = string(*v.VcenterEndpoint)
		}
		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.AssetsCompartmentId != nil {
			result["assets_compartment_id"] = *v.AssetsCompartmentId
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = *v.CompartmentId
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DisplayName != nil {
			result["display_name"] = *v.DisplayName
		}

		if v.EnvironmentId != nil {
			result["environment_id"] = *v.EnvironmentId
		}

		result["freeform_tags"] = v.FreeformTags

		if v.InventoryId != nil {
			result["inventory_id"] = *v.InventoryId
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = *v.LifecycleDetails
		}

		result["state"] = v.LifecycleState

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *CloudBridgeAssetSourceResourceCrud) populateTopLevelPolymorphicCreateAssetSourceRequest(request *oci_cloud_bridge.CreateAssetSourceRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("VMWARE"):
		details := oci_cloud_bridge.CreateVmWareAssetSourceDetails{}
		if areHistoricalMetricsCollected, ok := s.D.GetOkExists("are_historical_metrics_collected"); ok {
			tmp := areHistoricalMetricsCollected.(bool)
			details.AreHistoricalMetricsCollected = &tmp
		}
		if areRealtimeMetricsCollected, ok := s.D.GetOkExists("are_realtime_metrics_collected"); ok {
			tmp := areRealtimeMetricsCollected.(bool)
			details.AreRealtimeMetricsCollected = &tmp
		}
		if discoveryCredentials, ok := s.D.GetOkExists("discovery_credentials"); ok {
			if tmpList := discoveryCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "discovery_credentials", 0)
				tmp, err := s.mapToAssetSourceCredentials(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DiscoveryCredentials = &tmp
			}
		}
		if replicationCredentials, ok := s.D.GetOkExists("replication_credentials"); ok {
			if tmpList := replicationCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replication_credentials", 0)
				tmp, err := s.mapToAssetSourceCredentials(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ReplicationCredentials = &tmp
			}
		}
		if vcenterEndpoint, ok := s.D.GetOkExists("vcenter_endpoint"); ok {
			tmp := vcenterEndpoint.(string)
			details.VcenterEndpoint = &tmp
		}
		if assetsCompartmentId, ok := s.D.GetOkExists("assets_compartment_id"); ok {
			tmp := assetsCompartmentId.(string)
			details.AssetsCompartmentId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if discoveryScheduleId, ok := s.D.GetOkExists("discovery_schedule_id"); ok {
			tmp := discoveryScheduleId.(string)
			details.DiscoveryScheduleId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if environmentId, ok := s.D.GetOkExists("environment_id"); ok {
			tmp := environmentId.(string)
			details.EnvironmentId = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inventoryId, ok := s.D.GetOkExists("inventory_id"); ok {
			tmp := inventoryId.(string)
			details.InventoryId = &tmp
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.CreateAssetSourceDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *CloudBridgeAssetSourceResourceCrud) populateTopLevelPolymorphicUpdateAssetSourceRequest(request *oci_cloud_bridge.UpdateAssetSourceRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("VMWARE"):
		details := oci_cloud_bridge.UpdateVmWareAssetSourceDetails{}
		if areHistoricalMetricsCollected, ok := s.D.GetOkExists("are_historical_metrics_collected"); ok {
			tmp := areHistoricalMetricsCollected.(bool)
			details.AreHistoricalMetricsCollected = &tmp
		}
		if areRealtimeMetricsCollected, ok := s.D.GetOkExists("are_realtime_metrics_collected"); ok {
			tmp := areRealtimeMetricsCollected.(bool)
			details.AreRealtimeMetricsCollected = &tmp
		}
		if discoveryCredentials, ok := s.D.GetOkExists("discovery_credentials"); ok {
			if tmpList := discoveryCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "discovery_credentials", 0)
				tmp, err := s.mapToAssetSourceCredentials(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DiscoveryCredentials = &tmp
			}
		}
		if discoveryScheduleId, ok := s.D.GetOkExists("discovery_schedule_id"); ok {
			tmp := discoveryScheduleId.(string)
			details.DiscoveryScheduleId = &tmp
		}
		if replicationCredentials, ok := s.D.GetOkExists("replication_credentials"); ok {
			if tmpList := replicationCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replication_credentials", 0)
				tmp, err := s.mapToAssetSourceCredentials(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ReplicationCredentials = &tmp
			}
		}
		if vcenterEndpoint, ok := s.D.GetOkExists("vcenter_endpoint"); ok {
			tmp := vcenterEndpoint.(string)
			details.VcenterEndpoint = &tmp
		}
		tmp := s.D.Id()
		request.AssetSourceId = &tmp
		if assetsCompartmentId, ok := s.D.GetOkExists("assets_compartment_id"); ok {
			tmp := assetsCompartmentId.(string)
			details.AssetsCompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if discoveryScheduleId, ok := s.D.GetOkExists("discovery_schedule_id"); ok {
			tmp := discoveryScheduleId.(string)
			details.DiscoveryScheduleId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.UpdateAssetSourceDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *CloudBridgeAssetSourceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_bridge.ChangeAssetSourceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AssetSourceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.ChangeAssetSourceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
