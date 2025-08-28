// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagedKafkaKafkaClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagedKafkaKafkaCluster,
		Read:     readManagedKafkaKafkaCluster,
		Update:   updateManagedKafkaKafkaCluster,
		Delete:   deleteManagedKafkaKafkaCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"access_subnets": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnets": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"broker_shape": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"node_count": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"ocpu_count": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"cluster_config_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_config_version": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cluster_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"coordination_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kafka_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"client_certificate_bundle": {
				Type:     schema.TypeString,
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

			// Computed
			"kafka_bootstrap_urls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createManagedKafkaKafkaCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.CreateResource(d, sync)
}

func readManagedKafkaKafkaCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.ReadResource(sync)
}

func updateManagedKafkaKafkaCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteManagedKafkaKafkaCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ManagedKafkaKafkaClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_managed_kafka.KafkaClusterClient
	Res                    *oci_managed_kafka.KafkaCluster
	DisableNotFoundRetries bool
}

func (s *ManagedKafkaKafkaClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ManagedKafkaKafkaClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterLifecycleStateCreating),
	}
}

func (s *ManagedKafkaKafkaClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterLifecycleStateActive),
	}
}

func (s *ManagedKafkaKafkaClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterLifecycleStateDeleting),
	}
}

func (s *ManagedKafkaKafkaClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterLifecycleStateDeleted),
	}
}

func (s *ManagedKafkaKafkaClusterResourceCrud) Create() error {
	request := oci_managed_kafka.CreateKafkaClusterRequest{}

	if accessSubnets, ok := s.D.GetOkExists("access_subnets"); ok {
		interfaces := accessSubnets.([]interface{})
		tmp := make([]oci_managed_kafka.SubnetSet, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "access_subnets", stateDataIndex)
			converted, err := s.mapToSubnetSet(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("access_subnets") {
			request.AccessSubnets = tmp
		}
	}

	if brokerShape, ok := s.D.GetOkExists("broker_shape"); ok {
		if tmpList := brokerShape.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "broker_shape", 0)
			tmp, err := s.mapToBrokerShape(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BrokerShape = &tmp
		}
	}

	if clientCertificateBundle, ok := s.D.GetOkExists("client_certificate_bundle"); ok {
		tmp := clientCertificateBundle.(string)
		request.ClientCertificateBundle = &tmp
	}

	if clusterConfigId, ok := s.D.GetOkExists("cluster_config_id"); ok {
		tmp := clusterConfigId.(string)
		request.ClusterConfigId = &tmp
	}

	if clusterConfigVersion, ok := s.D.GetOkExists("cluster_config_version"); ok {
		tmp := clusterConfigVersion.(int)
		request.ClusterConfigVersion = &tmp
	}

	if clusterType, ok := s.D.GetOkExists("cluster_type"); ok {
		request.ClusterType = oci_managed_kafka.KafkaClusterClusterTypeEnum(clusterType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if coordinationType, ok := s.D.GetOkExists("coordination_type"); ok {
		request.CoordinationType = oci_managed_kafka.KafkaClusterCoordinationTypeEnum(coordinationType.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kafkaVersion, ok := s.D.GetOkExists("kafka_version"); ok {
		tmp := kafkaVersion.(string)
		request.KafkaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.CreateKafkaCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getKafkaClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ManagedKafkaKafkaClusterResourceCrud) getKafkaClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_managed_kafka.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	kafkaClusterId, err := kafkaClusterWaitForWorkRequest(workId, "kafkacluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, kafkaClusterId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_managed_kafka.CancelWorkRequestRequest{
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
	s.D.SetId(*kafkaClusterId)

	return s.Get()
}

func kafkaClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "managed_kafka", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_managed_kafka.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func kafkaClusterWaitForWorkRequest(wId *string, entityType string, action oci_managed_kafka.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_managed_kafka.KafkaClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "managed_kafka")
	retryPolicy.ShouldRetryOperation = kafkaClusterWorkRequestShouldRetryFunc(timeout)

	response := oci_managed_kafka.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_managed_kafka.OperationStatusInProgress),
			string(oci_managed_kafka.OperationStatusAccepted),
			string(oci_managed_kafka.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_managed_kafka.OperationStatusSucceeded),
			string(oci_managed_kafka.OperationStatusFailed),
			string(oci_managed_kafka.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_managed_kafka.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_managed_kafka.OperationStatusFailed || response.Status == oci_managed_kafka.OperationStatusCanceled {
		return nil, getErrorFromManagedKafkaKafkaClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromManagedKafkaKafkaClusterWorkRequest(client *oci_managed_kafka.KafkaClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_managed_kafka.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_managed_kafka.ListWorkRequestErrorsRequest{
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

func (s *ManagedKafkaKafkaClusterResourceCrud) Get() error {
	request := oci_managed_kafka.GetKafkaClusterRequest{}

	tmp := s.D.Id()
	request.KafkaClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.GetKafkaCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KafkaCluster
	return nil
}

func (s *ManagedKafkaKafkaClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_managed_kafka.UpdateKafkaClusterRequest{}

	if accessSubnets, ok := s.D.GetOkExists("access_subnets"); ok {
		interfaces := accessSubnets.([]interface{})
		tmp := make([]oci_managed_kafka.SubnetSet, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "access_subnets", stateDataIndex)
			converted, err := s.mapToSubnetSet(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("access_subnets") {
			request.AccessSubnets = tmp
		}
	}

	if brokerShape, ok := s.D.GetOkExists("broker_shape"); ok {
		if tmpList := brokerShape.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "broker_shape", 0)
			tmp, err := s.mapToBrokerShape(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BrokerShape = &tmp
		}
	}

	if clientCertificateBundle, ok := s.D.GetOkExists("client_certificate_bundle"); ok {
		tmp := clientCertificateBundle.(string)
		request.ClientCertificateBundle = &tmp
	}

	if clusterConfigId, ok := s.D.GetOkExists("cluster_config_id"); ok {
		tmp := clusterConfigId.(string)
		request.ClusterConfigId = &tmp
	}

	if clusterConfigVersion, ok := s.D.GetOkExists("cluster_config_version"); ok {
		tmp := clusterConfigVersion.(int)
		request.ClusterConfigVersion = &tmp
	}

	if coordinationType, ok := s.D.GetOkExists("coordination_type"); ok {
		request.CoordinationType = oci_managed_kafka.KafkaClusterCoordinationTypeEnum(coordinationType.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.KafkaClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.UpdateKafkaCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKafkaClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ManagedKafkaKafkaClusterResourceCrud) Delete() error {
	request := oci_managed_kafka.DeleteKafkaClusterRequest{}

	tmp := s.D.Id()
	request.KafkaClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.DeleteKafkaCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := kafkaClusterWaitForWorkRequest(workId, "kafkacluster",
		oci_managed_kafka.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ManagedKafkaKafkaClusterResourceCrud) SetData() error {
	accessSubnets := []interface{}{}
	for _, item := range s.Res.AccessSubnets {
		accessSubnets = append(accessSubnets, SubnetSetToMap(item))
	}
	s.D.Set("access_subnets", accessSubnets)

	if s.Res.BrokerShape != nil {
		s.D.Set("broker_shape", []interface{}{BrokerShapeToMap(s.Res.BrokerShape)})
	} else {
		s.D.Set("broker_shape", nil)
	}

	if s.Res.ClientCertificateBundle != nil {
		s.D.Set("client_certificate_bundle", *s.Res.ClientCertificateBundle)
	}

	if s.Res.ClusterConfigId != nil {
		s.D.Set("cluster_config_id", *s.Res.ClusterConfigId)
	}

	if s.Res.ClusterConfigVersion != nil {
		s.D.Set("cluster_config_version", *s.Res.ClusterConfigVersion)
	}

	s.D.Set("cluster_type", s.Res.ClusterType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("coordination_type", s.Res.CoordinationType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	kafkaBootstrapUrls := []interface{}{}
	for _, item := range s.Res.KafkaBootstrapUrls {
		kafkaBootstrapUrls = append(kafkaBootstrapUrls, BootstrapUrlToMap(item))
	}
	s.D.Set("kafka_bootstrap_urls", kafkaBootstrapUrls)

	if s.Res.KafkaVersion != nil {
		s.D.Set("kafka_version", *s.Res.KafkaVersion)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecretId != nil {
		s.D.Set("secret_id", *s.Res.SecretId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func BootstrapUrlToMap(obj oci_managed_kafka.BootstrapUrl) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func (s *ManagedKafkaKafkaClusterResourceCrud) mapToBrokerShape(fieldKeyFormat string) (oci_managed_kafka.BrokerShape, error) {
	result := oci_managed_kafka.BrokerShape{}

	if nodeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_count")); ok {
		tmp := nodeCount.(int)
		result.NodeCount = &tmp
	}

	if ocpuCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpu_count")); ok {
		tmp := ocpuCount.(int)
		result.OcpuCount = &tmp
	}

	if storageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_size_in_gbs")); ok {
		tmp := storageSizeInGbs.(int)
		result.StorageSizeInGbs = &tmp
	}

	return result, nil
}

func BrokerShapeToMap(obj *oci_managed_kafka.BrokerShape) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NodeCount != nil {
		result["node_count"] = int(*obj.NodeCount)
	}

	if obj.OcpuCount != nil {
		result["ocpu_count"] = int(*obj.OcpuCount)
	}

	if obj.StorageSizeInGbs != nil {
		result["storage_size_in_gbs"] = int(*obj.StorageSizeInGbs)
	}

	return result
}

func KafkaClusterSummaryToMap(obj oci_managed_kafka.KafkaClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	accessSubnets := []interface{}{}
	for _, item := range obj.AccessSubnets {
		accessSubnets = append(accessSubnets, SubnetSetToMap(item))
	}
	result["access_subnets"] = accessSubnets

	if obj.BrokerShape != nil {
		result["broker_shape"] = []interface{}{BrokerShapeToMap(obj.BrokerShape)}
	}

	if obj.ClusterConfigId != nil {
		result["cluster_config_id"] = string(*obj.ClusterConfigId)
	}

	if obj.ClusterConfigVersion != nil {
		result["cluster_config_version"] = int(*obj.ClusterConfigVersion)
	}

	result["cluster_type"] = string(obj.ClusterType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["coordination_type"] = string(obj.CoordinationType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.KafkaVersion != nil {
		result["kafka_version"] = string(*obj.KafkaVersion)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *ManagedKafkaKafkaClusterResourceCrud) mapToSubnetSet(fieldKeyFormat string) (oci_managed_kafka.SubnetSet, error) {
	result := oci_managed_kafka.SubnetSet{}

	if subnets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnets")); ok {
		interfaces := subnets.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "subnets")) {
			result.Subnets = tmp
		}
	}

	return result, nil
}

func SubnetSetToMap(obj oci_managed_kafka.SubnetSet) map[string]interface{} {
	result := map[string]interface{}{}

	result["subnets"] = obj.Subnets

	return result
}

func (s *ManagedKafkaKafkaClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_managed_kafka.ChangeKafkaClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.KafkaClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.ChangeKafkaClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKafkaClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
