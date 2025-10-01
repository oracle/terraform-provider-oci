// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagedKafkaKafkaClusterSuperusersManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagedKafkaKafkaClusterSuperusersManagement,
		Read:     readManagedKafkaKafkaClusterSuperusersManagement,
		Update:   updateManagedKafkaKafkaClusterSuperusersManagement,
		Delete:   deleteManagedKafkaKafkaClusterSuperusersManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"kafka_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_superuser": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createManagedKafkaKafkaClusterSuperusersManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterSuperusersManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()
	sync.Res = &ManagedKafkaKafkaClusterSuperusersManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readManagedKafkaKafkaClusterSuperusersManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateManagedKafkaKafkaClusterSuperusersManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterSuperusersManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()
	sync.Res = &ManagedKafkaKafkaClusterSuperusersManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteManagedKafkaKafkaClusterSuperusersManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterSuperusersManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()
	sync.Res = &ManagedKafkaKafkaClusterSuperusersManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ManagedKafkaKafkaClusterSuperusersManagementResponse struct {
	enableResponse  *oci_managed_kafka.EnableSuperuserResponse
	disableResponse *oci_managed_kafka.DisableSuperuserResponse
}

type ManagedKafkaKafkaClusterSuperusersManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_managed_kafka.KafkaClusterClient
	Res                    *ManagedKafkaKafkaClusterSuperusersManagementResponse
	DisableNotFoundRetries bool
}

func (s *ManagedKafkaKafkaClusterSuperusersManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("ManagedKafkaKafkaClusterSuperusersManagementResource-", ManagedKafkaKafkaClusterSuperusersManagementResource(), s.D)
}

func (s *ManagedKafkaKafkaClusterSuperusersManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_superuser"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_managed_kafka.EnableSuperuserRequest{}

		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			request.CompartmentId = &tmp
		}

		if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
			tmp := kafkaClusterId.(string)
			request.KafkaClusterId = &tmp
		}

		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			request.SecretId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

		response, err := s.Client.EnableSuperuser(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getKafkaClusterSuperusersManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_managed_kafka.DisableSuperuserRequest{}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.DisableSuperuser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getKafkaClusterSuperusersManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *ManagedKafkaKafkaClusterSuperusersManagementResourceCrud) getKafkaClusterSuperusersManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_managed_kafka.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	kafkaClusterSuperusersManagementId, err := kafkaClusterSuperusersManagementWaitForWorkRequest(workId, "kafkacluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, kafkaClusterSuperusersManagementId)
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

	return nil
}

func kafkaClusterSuperusersManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func kafkaClusterSuperusersManagementWaitForWorkRequest(wId *string, entityType string, action oci_managed_kafka.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_managed_kafka.KafkaClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "managed_kafka")
	retryPolicy.ShouldRetryOperation = kafkaClusterSuperusersManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_managed_kafka.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
			if string(res.ActionType) == string(action) {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_managed_kafka.OperationStatusFailed || response.Status == oci_managed_kafka.OperationStatusCanceled {
		return nil, getErrorFromManagedKafkaKafkaClusterSuperusersManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromManagedKafkaKafkaClusterSuperusersManagementWorkRequest(client *oci_managed_kafka.KafkaClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_managed_kafka.ActionTypeEnum) error {
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

func (s *ManagedKafkaKafkaClusterSuperusersManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_superuser"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_managed_kafka.EnableSuperuserRequest{}

		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			request.CompartmentId = &tmp
		}

		if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
			tmp := kafkaClusterId.(string)
			request.KafkaClusterId = &tmp
		}

		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			request.SecretId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

		response, err := s.Client.EnableSuperuser(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getKafkaClusterSuperusersManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_managed_kafka.DisableSuperuserRequest{}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.DisableSuperuser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getKafkaClusterSuperusersManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *ManagedKafkaKafkaClusterSuperusersManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_superuser"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_managed_kafka.DisableSuperuserRequest{}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.DisableSuperuser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getKafkaClusterSuperusersManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *ManagedKafkaKafkaClusterSuperusersManagementResourceCrud) SetData() error {
	return nil
}
