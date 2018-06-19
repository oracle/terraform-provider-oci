// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"time"

	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func LocalPeeringGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createLocalPeeringGateway,
		Read:     readLocalPeeringGateway,
		Update:   updateLocalPeeringGateway,
		Delete:   deleteLocalPeeringGateway,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			// @CODEGEN we use peer_id to do the connect action
			"peer_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateNotEmptyString(), //Don't allow empty string, it results in a terraform error when switching from valid value to empty string
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cross_tenancy_peering": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"peer_advertised_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peering_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peering_status_details": {
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

func createLocalPeeringGateway(d *schema.ResourceData, m interface{}) error {
	sync := &LocalPeeringGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	err := crud.CreateResource(d, sync)
	if err != nil {
		return err
	}
	//This needs to be here rather than in the Create() because we want the resource to finish provisioning and set to the statefile before we connect
	return sync.ConnectLocalPeeringGateway()
}

func readLocalPeeringGateway(d *schema.ResourceData, m interface{}) error {
	sync := &LocalPeeringGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateLocalPeeringGateway(d *schema.ResourceData, m interface{}) error {
	sync := &LocalPeeringGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteLocalPeeringGateway(d *schema.ResourceData, m interface{}) error {
	sync := &LocalPeeringGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type LocalPeeringGatewayResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.LocalPeeringGateway
	DisableNotFoundRetries bool
}

func (s *LocalPeeringGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LocalPeeringGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.LocalPeeringGatewayLifecycleStateProvisioning),
	}
}

func (s *LocalPeeringGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.LocalPeeringGatewayLifecycleStateAvailable),
	}
}

func (s *LocalPeeringGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.LocalPeeringGatewayLifecycleStateTerminating),
	}
}

func (s *LocalPeeringGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.LocalPeeringGatewayLifecycleStateTerminated),
	}
}

func (s *LocalPeeringGatewayResourceCrud) ConnectLocalPeeringGateway() error {
	if s.Res == nil || s.Res.Id == nil {
		return fmt.Errorf("CreateLocalPeeringGateway returned a nil LocalPeeringGateway or a LocalPeeringGateway without an ID")
	}

	if peerId, ok := s.D.GetOkExists("peer_id"); ok {
		connectRequest := oci_core.ConnectLocalPeeringGatewaysRequest{}

		tmp := peerId.(string)
		connectRequest.PeerId = &tmp

		connectRequest.LocalPeeringGatewayId = s.Res.Id

		connectRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

		_, err := s.Client.ConnectLocalPeeringGateways(context.Background(), connectRequest)
		if err != nil {
			// we set peer_id to "" so that terraform detects a forceNew change on the next apply and the user can try the connection again
			s.D.Set("peer_id", "")
			return err
		}

		// wait for peering status to not be Pending
		waitForPeerStatusPolicy := &common.RetryPolicy{
			ShouldRetryOperation:  waitForLPGPeeringStatusShouldRetry(s.D.Timeout(schema.TimeoutCreate)),
			NextDuration:          nextDuration,
			MaximumNumberAttempts: 0,
		}
		request := oci_core.GetLocalPeeringGatewayRequest{}

		tmpId := s.D.Id()
		request.LocalPeeringGatewayId = &tmpId

		request.RequestMetadata.RetryPolicy = waitForPeerStatusPolicy

		response, getError := s.Client.GetLocalPeeringGateway(context.Background(), request)
		if getError != nil {
			log.Printf("[DEBUG] Get error while waiting for peering connection to finish: %+v", getError)
			return getError
		}
		s.Res = &response.LocalPeeringGateway
		if response.LocalPeeringGateway.PeeringStatus != oci_core.LocalPeeringGatewayPeeringStatusPeered {
			s.D.Set("peer_id", "")
			return fmt.Errorf("unexpected Peering Status `%s` after trying to connect to the peer Local Peering Gateway", string(response.LocalPeeringGateway.PeeringStatus))
		}
		s.SetData()
	}
	return nil
}

func (s *LocalPeeringGatewayResourceCrud) Create() error {
	request := oci_core.CreateLocalPeeringGatewayRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateLocalPeeringGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LocalPeeringGateway
	return nil
}

func (s *LocalPeeringGatewayResourceCrud) Get() error {
	request := oci_core.GetLocalPeeringGatewayRequest{}

	tmp := s.D.Id()
	request.LocalPeeringGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetLocalPeeringGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LocalPeeringGateway
	return nil
}

func (s *LocalPeeringGatewayResourceCrud) Update() error {
	request := oci_core.UpdateLocalPeeringGatewayRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.LocalPeeringGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateLocalPeeringGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LocalPeeringGateway
	return nil
}

func (s *LocalPeeringGatewayResourceCrud) Delete() error {
	request := oci_core.DeleteLocalPeeringGatewayRequest{}

	tmp := s.D.Id()
	request.LocalPeeringGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteLocalPeeringGateway(context.Background(), request)
	return err
}

func (s *LocalPeeringGatewayResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.IsCrossTenancyPeering != nil {
		s.D.Set("is_cross_tenancy_peering", *s.Res.IsCrossTenancyPeering)
	}

	if s.Res.PeerAdvertisedCidr != nil {
		s.D.Set("peer_advertised_cidr", *s.Res.PeerAdvertisedCidr)
	}

	s.D.Set("peering_status", s.Res.PeeringStatus)

	if s.Res.PeeringStatusDetails != nil {
		s.D.Set("peering_status_details", *s.Res.PeeringStatusDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

}

func waitForLPGPeeringStatusShouldRetry(timeout time.Duration) func(response common.OCIOperationResponse) bool {
	return func(response common.OCIOperationResponse) bool {
		if shouldRetry(response, false, "core") {
			return true
		}
		if getLocalPeeringGatewayResponse, ok := response.Response.(oci_core.GetLocalPeeringGatewayResponse); ok {
			if getLocalPeeringGatewayResponse.PeeringStatus == oci_core.LocalPeeringGatewayPeeringStatusPending {
				timeWaited := getTimeWaited(response.AttemptNumber)
				return timeWaited < timeout
			}
		}
		return false
	}
}
