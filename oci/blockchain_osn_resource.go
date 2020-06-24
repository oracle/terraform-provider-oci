// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	oci_blockchain "github.com/oracle/oci-go-sdk/blockchain"
	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterResource("oci_blockchain_osn", BlockchainOsnResource())
}

func BlockchainOsnResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("30m"),
			Update: getTimeoutDuration("30m"),
			Delete: getTimeoutDuration("30m"),
		},
		Create: createBlockchainOsn,
		Read:   readBlockchainOsn,
		Delete: deleteBlockchainOsn,
		Schema: map[string]*schema.Schema{
			// Required
			"ad": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"blockchain_platform_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ocpu_allocation_param": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ocpu_allocation_number": {
							Type:     schema.TypeFloat,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"osn_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBlockchainOsn(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainOsnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockchainPlatformClient()

	return CreateResource(d, sync)
}

func readBlockchainOsn(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainOsnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockchainPlatformClient()

	return ReadResource(sync)
}

func deleteBlockchainOsn(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainOsnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockchainPlatformClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type BlockchainOsnResourceCrud struct {
	BaseCrud
	Client                 *oci_blockchain.BlockchainPlatformClient
	Res                    *oci_blockchain.Osn
	DisableNotFoundRetries bool
}

func (s *BlockchainOsnResourceCrud) ID() string {
	return *s.Res.OsnKey
}

func (s *BlockchainOsnResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_blockchain.OsnLifecycleStateActive),
	}
}

func (s *BlockchainOsnResourceCrud) Create() error {
	request := oci_blockchain.CreateOsnRequest{}

	if ad, ok := s.D.GetOkExists("ad"); ok {
		request.Ad = oci_blockchain.AvailabilityDomainAdsEnum(ad.(string))
	}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	if ocpuAllocationParam, ok := s.D.GetOkExists("ocpu_allocation_param"); ok {
		if tmpList := ocpuAllocationParam.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ocpu_allocation_param", 0)
			tmp, err := s.mapToOcpuAllocationNumberParam(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OcpuAllocationParam = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.CreateOsn(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOsnFromWorkRequest(request.BlockchainPlatformId, workId, getRetryPolicy(s.DisableNotFoundRetries, "blockchain"), oci_blockchain.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BlockchainOsnResourceCrud) getOsnFromWorkRequest(blockchainPlatformId *string, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_blockchain.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Service not return osn directly from work request.
	// So to get the osn, we compare list before and after new osn creation
	listOsnBefore, err := getListOsnFromBlockChainPlatform(blockchainPlatformId, s.Client)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] listOsnBefore length: %v\n", len(listOsnBefore))
	// Wait until it finishes
	_, err = osnWaitForWorkRequest(workId, "instance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, blockchainPlatformId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_blockchain.DeleteWorkRequestRequest{
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
	listOsnAfter, err := getListOsnFromBlockChainPlatform(blockchainPlatformId, s.Client)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] listOsnAfter length: %v\n", len(listOsnAfter))
	osnId, err := difference(listOsnAfter, listOsnBefore)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] new osn keyId create: %v\n", *osnId)
	s.D.SetId(*osnId)

	return s.Get()
}

func osnWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "blockchain", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_blockchain.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func osnWaitForWorkRequest(wId *string, entityType string, action oci_blockchain.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_blockchain.BlockchainPlatformClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "blockchain")
	retryPolicy.ShouldRetryOperation = osnWorkRequestShouldRetryFunc(timeout)

	response := oci_blockchain.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_blockchain.WorkRequestStatusInProgress),
			string(oci_blockchain.WorkRequestStatusAccepted),
			string(oci_blockchain.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_blockchain.WorkRequestStatusSucceeded),
			string(oci_blockchain.WorkRequestStatusFailed),
			string(oci_blockchain.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_blockchain.GetWorkRequestRequest{
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

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error
	if response.Status == oci_blockchain.WorkRequestStatusFailed {
		errorMessage := getErrorFromBlockchainPlatformWorkRequest(response, client)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return identifier, workRequestErr
}

func (s *BlockchainOsnResourceCrud) Get() error {
	request := oci_blockchain.GetOsnRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	tmp := s.D.Id()
	request.OsnId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.GetOsn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Osn
	return nil
}

func (s *BlockchainOsnResourceCrud) Delete() error {
	// service not support delete yet
	return nil
}

func (s *BlockchainOsnResourceCrud) SetData() error {
	s.D.Set("ad", s.Res.Ad)

	if s.Res.OcpuAllocationParam != nil {
		s.D.Set("ocpu_allocation_param", []interface{}{OcpuAllocationNumberParamToMap(s.Res.OcpuAllocationParam)})
	} else {
		s.D.Set("ocpu_allocation_param", nil)
	}

	if s.Res.OsnKey != nil {
		s.D.Set("osn_key", *s.Res.OsnKey)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}

func (s *BlockchainOsnResourceCrud) mapToOcpuAllocationNumberParam(fieldKeyFormat string) (oci_blockchain.OcpuAllocationNumberParam, error) {
	result := oci_blockchain.OcpuAllocationNumberParam{}

	if ocpuAllocationNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpu_allocation_number")); ok {
		tmp := ocpuAllocationNumber.(float32)
		result.OcpuAllocationNumber = &tmp
	}

	return result, nil
}

func OsnSummaryToMap(obj oci_blockchain.OsnSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OsnKey != nil {
		result["osn_key"] = string(*obj.OsnKey)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}
