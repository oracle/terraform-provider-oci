// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreRemotePeeringConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreRemotePeeringConnection,
		Read:     readCoreRemotePeeringConnection,
		Update:   updateCoreRemotePeeringConnection,
		Delete:   deleteCoreRemotePeeringConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			// @CODEGEN peer_id and peer_region_name moved from computed to optional as they are required for the connect action
			"peer_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: utils.ValidateNotEmptyString(), //Don't allow empty string, it results in a terraform error when switching from valid value to empty string
			},
			"peer_region_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: utils.ValidateNotEmptyString(), //Don't allow empty string, it results in a terraform error when switching from valid value to empty string
			},

			// Computed
			"is_cross_tenancy_peering": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"peer_tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peering_status": {
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

func createCoreRemotePeeringConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRemotePeeringConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	//This needs to be here rather than in the Create() because we want the resource to finish provisioning and set to the statefile before we connect
	return sync.ConnectRemotePeeringConnection()
}

func readCoreRemotePeeringConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRemotePeeringConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreRemotePeeringConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRemotePeeringConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreRemotePeeringConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRemotePeeringConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreRemotePeeringConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.RemotePeeringConnection
	DisableNotFoundRetries bool
}

func (s *CoreRemotePeeringConnectionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreRemotePeeringConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.RemotePeeringConnectionLifecycleStateProvisioning),
	}
}

func (s *CoreRemotePeeringConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.RemotePeeringConnectionLifecycleStateAvailable),
	}
}

func (s *CoreRemotePeeringConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.RemotePeeringConnectionLifecycleStateTerminating),
	}
}

func (s *CoreRemotePeeringConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.RemotePeeringConnectionLifecycleStateTerminated),
	}
}

func (s *CoreRemotePeeringConnectionResourceCrud) ConnectRemotePeeringConnection() error {
	if s.Res == nil || s.Res.Id == nil {
		return fmt.Errorf("CreateRemotePeeringConnection returned a nil RemotePeeringConnection or a RemotePeeringConnection without an ID")
	}

	peerId, ok := s.D.GetOkExists("peer_id")
	if !ok {
		return nil
	}

	connectRequest := oci_core.ConnectRemotePeeringConnectionsRequest{}

	tmp := peerId.(string)
	connectRequest.PeerId = &tmp

	connectRequest.RemotePeeringConnectionId = s.Res.Id

	if peerRegionName, ok := s.D.GetOkExists("peer_region_name"); ok {
		tmp := peerRegionName.(string)
		connectRequest.PeerRegionName = &tmp
	}

	connectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ConnectRemotePeeringConnections(context.Background(), connectRequest)
	if err != nil {
		// we set peer_id to "" so that terraform detects a forceNew change on the next apply and the user can try the connection again
		s.D.Set("peer_id", "")
		return err
	}

	request := oci_core.GetRemotePeeringConnectionRequest{}

	tmpId := s.D.Id()
	request.RemotePeeringConnectionId = &tmpId

	request.RequestMetadata.RetryPolicy = getRemotePeeringConnectionRetryPolicy(s.D.Timeout(schema.TimeoutCreate))

	response, getError := s.Client.GetRemotePeeringConnection(context.Background(), request)
	if getError != nil {
		log.Printf("[DEBUG] Get error while waiting for peering connection to finish: %+v", getError)
		return getError
	}
	s.Res = &response.RemotePeeringConnection
	if response.RemotePeeringConnection.PeeringStatus != oci_core.RemotePeeringConnectionPeeringStatusPeered {
		s.D.Set("peer_id", "")
		return fmt.Errorf("unexpected Peering Status `%s` after trying to connect to the peer Remote Peering Connection (RPC). Make sure the peering_status of the peer RPC is not REVOKED", string(response.RemotePeeringConnection.PeeringStatus))
	}

	if err := s.SetData(); err != nil {
		return err
	}

	return nil
}

func (s *CoreRemotePeeringConnectionResourceCrud) Create() error {
	request := oci_core.CreateRemotePeeringConnectionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateRemotePeeringConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RemotePeeringConnection
	return nil
}

func (s *CoreRemotePeeringConnectionResourceCrud) Get() error {
	request := oci_core.GetRemotePeeringConnectionRequest{}

	tmp := s.D.Id()
	request.RemotePeeringConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetRemotePeeringConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RemotePeeringConnection
	return nil
}

func (s *CoreRemotePeeringConnectionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateRemotePeeringConnectionRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.RemotePeeringConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateRemotePeeringConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RemotePeeringConnection
	return nil
}

func (s *CoreRemotePeeringConnectionResourceCrud) Delete() error {
	request := oci_core.DeleteRemotePeeringConnectionRequest{}

	tmp := s.D.Id()
	request.RemotePeeringConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteRemotePeeringConnection(context.Background(), request)
	return err
}

func (s *CoreRemotePeeringConnectionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCrossTenancyPeering != nil {
		s.D.Set("is_cross_tenancy_peering", *s.Res.IsCrossTenancyPeering)
	}

	if s.Res.PeerId != nil {
		s.D.Set("peer_id", *s.Res.PeerId)
	}

	if s.Res.PeerRegionName != nil {
		s.D.Set("peer_region_name", *s.Res.PeerRegionName)
	}

	if s.Res.PeerTenancyId != nil {
		s.D.Set("peer_tenancy_id", *s.Res.PeerTenancyId)
	}

	s.D.Set("peering_status", s.Res.PeeringStatus)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func getRemotePeeringConnectionRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()
	// wait for peering status to not be Pending
	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if tfresource.ShouldRetry(response, false, "core", startTime) {
				return true
			}
			if getRemotePeeringConnectionResponse, ok := response.Response.(oci_core.GetRemotePeeringConnectionResponse); ok {
				if getRemotePeeringConnectionResponse.PeeringStatus == oci_core.RemotePeeringConnectionPeeringStatusPending {
					timeWaited := tfresource.GetElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return tfresource.GetRetryBackoffDuration(response, false, "core", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}

func (s *CoreRemotePeeringConnectionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeRemotePeeringConnectionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RemotePeeringConnectionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeRemotePeeringConnectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
