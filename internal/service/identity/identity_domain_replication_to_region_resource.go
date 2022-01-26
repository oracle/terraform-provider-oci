// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"
)

func IdentityDomainReplicationToRegionResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainReplicationToRegion,
		Read:     readIdentityDomainReplicationToRegion,
		Delete:   deleteIdentityDomainReplicationToRegion,
		Schema: map[string]*schema.Schema{
			// Required
			"domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"replica_region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createIdentityDomainReplicationToRegion(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainReplicationToRegionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainReplicationToRegion(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIdentityDomainReplicationToRegion(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainReplicationToRegionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Domain
	DisableNotFoundRetries bool
}

func (s *IdentityDomainReplicationToRegionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainReplicationToRegionResourceCrud) Create() error {
	request := oci_identity.EnableReplicationToRegionRequest{}

	if domainId, ok := s.D.GetOkExists("domain_id"); ok {
		tmp := domainId.(string)
		request.DomainId = &tmp
	}

	if replicaRegion, ok := s.D.GetOkExists("replica_region"); ok {
		tmp := replicaRegion.(string)
		request.ReplicaRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.EnableReplicationToRegion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDomainReplicationToRegionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "domain"), oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IdentityDomainReplicationToRegionResourceCrud) getDomainReplicationToRegionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_identity.IamWorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	domainReplicationToRegionId, err := domainReplicationToRegionWaitForWorkRequest(workId, "domain",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*domainReplicationToRegionId)

	return s.Get()
}

func (s *IdentityDomainReplicationToRegionResourceCrud) Get() error {
	request := oci_identity.GetDomainRequest{}

	tmp := s.D.Id()
	request.DomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Domain
	return nil
}

func domainReplicationToRegionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "identity", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_identity.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func domainReplicationToRegionWaitForWorkRequest(wId *string, entityType string, action oci_identity.IamWorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_identity.IdentityClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "identity")
	retryPolicy.ShouldRetryOperation = domainReplicationToRegionWorkRequestShouldRetryFunc(timeout)

	response := oci_identity.GetIamWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_identity.IamWorkRequestStatusAccepted),
			string(oci_identity.IamWorkRequestStatusInProgress),
			string(oci_identity.IamWorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_identity.IamWorkRequestStatusSucceeded),
			string(oci_identity.IamWorkRequestStatusFailed),
			string(oci_identity.IamWorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetIamWorkRequest(context.Background(),
				oci_identity.GetIamWorkRequestRequest{
					IamWorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			return response, string(response.Status), err
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
	if identifier == nil || response.Status == oci_identity.IamWorkRequestStatusFailed || response.Status == oci_identity.IamWorkRequestStatusCanceled {
		return nil, getErrorFromIdentityDomainReplicaWorkRequest(client, wId, retryPolicy, response.OperationType)
	}

	return identifier, nil
}

func getErrorFromIdentityDomainReplicaWorkRequest(client *oci_identity.IdentityClient, workId *string, retryPolicy *oci_common.RetryPolicy, operationType oci_identity.IamWorkRequestOperationTypeEnum) error {

	errorMessage, err := getErrorMessageFromIdentityDomainReplicaWorkRequest(client, workId, retryPolicy)
	if err != nil {
		return err
	}

	workRequestErr := fmt.Errorf("oci_identity_domain: iam work request did not succeed, workId: %s, action: %s. ErrorMessage: %s", *workId, operationType, errorMessage)
	return workRequestErr
}

func getErrorMessageFromIdentityDomainReplicaWorkRequest(client *oci_identity.IdentityClient, workId *string, retryPolicy *oci_common.RetryPolicy) (string, error) {
	errorMessage := ""
	response, err := client.ListIamWorkRequestErrors(context.Background(),
		oci_identity.ListIamWorkRequestErrorsRequest{
			IamWorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})

	if err != nil {
		return errorMessage, err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage = strings.Join(allErrs, "\n")

	return errorMessage, nil
}

func (s *IdentityDomainReplicationToRegionResourceCrud) SetData() error {
	return nil
}

func (s *IdentityDomainResourceCrud) GetDomain() error {
	request := oci_identity.GetDomainRequest{}

	tmp := s.D.Id()
	request.DomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Domain
	return nil
}
