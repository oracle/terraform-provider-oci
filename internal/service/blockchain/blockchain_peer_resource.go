// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_blockchain "github.com/oracle/oci-go-sdk/v56/blockchain"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func BlockchainPeerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
			Update: tfresource.GetTimeoutDuration("30m"),
			Delete: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createBlockchainPeer,
		Read:   readBlockchainPeer,
		Update: updateBlockchainPeer,
		Delete: deleteBlockchainPeer,
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
			"ocpu_allocation_param": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ocpu_allocation_number": {
							Type:             schema.TypeFloat,
							Required:         true,
							DiffSuppressFunc: utils.MonetaryDiffSuppress,
						},

						// Optional

						// Computed
					},
				},
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"alias": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_key": {
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

func createBlockchainPeer(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainPeerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.CreateResource(d, sync)
}

func readBlockchainPeer(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainPeerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.ReadResource(sync)
}

func updateBlockchainPeer(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainPeerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBlockchainPeer(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainPeerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BlockchainPeerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_blockchain.BlockchainPlatformClient
	Res                    *oci_blockchain.Peer
	DisableNotFoundRetries bool
}

func (s *BlockchainPeerResourceCrud) ID() string {
	return GetPeerCompositeId(s.D.Get("blockchain_platform_id").(string), *s.Res.PeerKey)
}

func (s *BlockchainPeerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_blockchain.PeerLifecycleStateActive),
	}
}

func (s *BlockchainPeerResourceCrud) Create() error {
	request := oci_blockchain.CreatePeerRequest{}

	if ad, ok := s.D.GetOkExists("ad"); ok {
		request.Ad = oci_blockchain.AvailabilityDomainAdsEnum(ad.(string))
	}

	if alias, ok := s.D.GetOkExists("alias"); ok {
		tmp := alias.(string)
		request.Alias = &tmp
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

	if role, ok := s.D.GetOkExists("role"); ok {
		request.Role = oci_blockchain.PeerRoleRoleEnum(role.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.CreatePeer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPeerFromWorkRequest(request.BlockchainPlatformId, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain"), oci_blockchain.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BlockchainPeerResourceCrud) getPeerFromWorkRequest(blockchainPlatformId *string, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_blockchain.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	peerId, err := peerWaitForWorkRequest(workId, "instance",
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

	if err != nil {
		return err
	}
	s.D.SetId(*peerId)

	return s.Get()
}

func peerWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "blockchain", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_blockchain.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func peerWaitForWorkRequest(wId *string, entityType string, action oci_blockchain.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_blockchain.BlockchainPlatformClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "blockchain")
	retryPolicy.ShouldRetryOperation = peerWorkRequestShouldRetryFunc(timeout)

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

	var subTypeKey *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				for _, subTypeDetail := range res.SubTypeDetails {
					subTypeKey = subTypeDetail.SubTypeKey
				}
				break
			}
		}
	}

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error
	if response.Status == oci_blockchain.WorkRequestStatusFailed {
		errorMessage := getErrorFromBlockchainPeerWorkRequest(client, wId, retryPolicy, entityType, action)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return subTypeKey, workRequestErr
}

func getErrorFromBlockchainPeerWorkRequest(client *oci_blockchain.BlockchainPlatformClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_blockchain.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_blockchain.ListWorkRequestErrorsRequest{
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

func (s *BlockchainPeerResourceCrud) Get() error {
	request := oci_blockchain.GetPeerRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	tmp := s.D.Id()
	request.PeerId = &tmp

	blockchainPlatformId, peerId, err := parsePeerCompositeId(s.D.Id())
	if err == nil {
		request.BlockchainPlatformId = &blockchainPlatformId
		request.PeerId = &peerId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.GetPeer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Peer
	return nil
}

func (s *BlockchainPeerResourceCrud) Update() error {
	request := oci_blockchain.UpdatePeerRequest{}

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

	tmp := s.D.Id()
	request.PeerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.UpdatePeer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, err = peerWaitForWorkRequest(workId, "instance",
		oci_blockchain.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *BlockchainPeerResourceCrud) Delete() error {
	request := oci_blockchain.DeletePeerRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	tmp := s.D.Id()
	request.PeerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.DeletePeer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := peerWaitForWorkRequest(workId, "instance",
		oci_blockchain.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BlockchainPeerResourceCrud) SetData() error {

	blockchainPlatformId, peerId, err := parsePeerCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("blockchain_platform_id", &blockchainPlatformId)
		s.D.SetId(peerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("ad", s.Res.Ad)

	if s.Res.Alias != nil {
		s.D.Set("alias", *s.Res.Alias)
	}

	if s.Res.Host != nil {
		s.D.Set("host", *s.Res.Host)
	}

	if s.Res.OcpuAllocationParam != nil {
		s.D.Set("ocpu_allocation_param", []interface{}{OcpuAllocationNumberParamToMap(s.Res.OcpuAllocationParam)})
	} else {
		s.D.Set("ocpu_allocation_param", nil)
	}

	if s.Res.PeerKey != nil {
		s.D.Set("peer_key", *s.Res.PeerKey)
	}

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}

func GetPeerCompositeId(blockchainPlatformId string, peerId string) string {
	blockchainPlatformId = url.PathEscape(blockchainPlatformId)
	peerId = url.PathEscape(peerId)
	compositeId := "blockchainPlatforms/" + blockchainPlatformId + "/peers/" + peerId
	return compositeId
}

func parsePeerCompositeId(compositeId string) (blockchainPlatformId string, peerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("blockchainPlatforms/.*/peers/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	blockchainPlatformId, _ = url.PathUnescape(parts[1])
	peerId, _ = url.PathUnescape(parts[3])
	return
}

func (s *BlockchainPeerResourceCrud) mapToOcpuAllocationNumberParam(fieldKeyFormat string) (oci_blockchain.OcpuAllocationNumberParam, error) {
	result := oci_blockchain.OcpuAllocationNumberParam{}

	if ocpuAllocationNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpu_allocation_number")); ok {
		tmp := float32(ocpuAllocationNumber.(float64))
		result.OcpuAllocationNumber = &tmp
	}

	return result, nil
}

func PeerSummaryToMap(obj oci_blockchain.PeerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PeerKey != nil {
		result["peer_key"] = string(*obj.PeerKey)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}
