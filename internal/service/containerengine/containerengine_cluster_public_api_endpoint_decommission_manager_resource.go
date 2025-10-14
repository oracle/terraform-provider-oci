// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterPublicApiEndpointDecommissionManagerResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createContainerengineClusterPublicApiEndpointDecommissionManager,
		Read:   readContainerengineClusterPublicApiEndpointDecommissionManager,
		Update: updateContainerengineClusterPublicApiEndpointDecommissionManager,
		Delete: deleteContainerengineClusterPublicApiEndpointDecommissionManager,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_public_api_endpoint_decommissioned": {
				Type:     schema.TypeBool,
				Required: true,
			},
			// Optional
			"rollback_deadline_delay": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
		},
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error {
			isDecom, _ := diff.GetOk("is_public_api_endpoint_decommissioned")
			rollbackDelaySet, _ := diff.GetOk("rollback_deadline_delay")
			if !isDecom.(bool) && rollbackDelaySet != "" {
				return fmt.Errorf("rollback_deadline_delay cannot be specified when is_public_api_endpoint_decommissioned is false")
			}
			return nil
		},
	}
}

func createContainerengineClusterPublicApiEndpointDecommissionManager(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	return tfresource.CreateResource(d, sync)
}

func readContainerengineClusterPublicApiEndpointDecommissionManager(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineClusterPublicApiEndpointDecommissionManager(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	return tfresource.UpdateResource(d, sync)
}

// delete is an no-op
func deleteContainerengineClusterPublicApiEndpointDecommissionManager(d *schema.ResourceData, m interface{}) error {
	return nil
}

type ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud struct {
	tfresource.BaseCrud
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.PublicApiEndpointDecommissionStatus

	DisableNotFoundRetries bool
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("ContainerengineClusterPublicApiEndpointDecommissionManagerResource-", ContainerengineClusterPublicApiEndpointDecommissionManagerResource(), s.D)
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud) Get() error {
	clusterId, _ := s.D.GetOkExists("cluster_id")
	tmpClusterId := clusterId.(string)
	requestGet := oci_containerengine.GetPublicApiEndpointDecommissionStatusRequest{}
	requestGet.ClusterId = &tmpClusterId
	requestGet.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err := s.Client.GetPublicApiEndpointDecommissionStatus(context.Background(), requestGet)
	if err != nil {
		return err
	}
	s.Res = &responseGet.PublicApiEndpointDecommissionStatus
	return nil
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud) Create() error {
	var operation bool
	if is_public_api_endpoint_decommissioned, ok := s.D.GetOkExists("is_public_api_endpoint_decommissioned"); ok {
		operation = is_public_api_endpoint_decommissioned.(bool)
	}

	clusterId, _ := s.D.GetOkExists("cluster_id")
	tmpClusterId := clusterId.(string)

	if operation {
		// do public api endpoint decommission
		request := oci_containerengine.StartPublicApiEndpointDecommissionRequest{}
		request.ClusterId = &tmpClusterId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
		response, err := s.Client.StartPublicApiEndpointDecommission(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		// Wait for work request to finish
		_, waitErr := clusterWaitForWorkRequest(workId, "cluster",
			oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
		if waitErr != nil {
			return waitErr
		}
	} else {
		// raise rollback request for public api endpoint decommission
		request := oci_containerengine.RollbackPublicApiEndpointDecommissionRequest{}
		request.ClusterId = &tmpClusterId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
		response, err := s.Client.RollbackPublicApiEndpointDecommission(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		// Wait for work request to finish
		_, waitErr := clusterWaitForWorkRequest(workId, "cluster",
			oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
		if waitErr != nil {
			return waitErr
		}
	}

	// store in state file in case of failure
	if setDataErr := s.SetData(); setDataErr != nil {
		log.Printf("[ERROR] error setting data before clusterWaitForWorkRequest() error: %v", setDataErr)
	}

	// if rollback deadline delay is specified, raise request for extendRollback deadline delay
	if rollbackDeadlineDelay, ok := s.D.GetOkExists("rollback_deadline_delay"); ok {
		requestExtend := oci_containerengine.ExtendEndpointDecommissionRollbackDeadlineRequest{}
		delay := rollbackDeadlineDelay.(string)
		requestExtend.RollbackDeadlineDelay = &delay
		requestExtend.ClusterId = &tmpClusterId

		requestExtend.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
		_, err := s.Client.ExtendEndpointDecommissionRollbackDeadline(context.Background(), requestExtend)
		if err != nil {
			return err
		}
		if setDataErr := s.SetData(); setDataErr != nil {
			log.Printf("[ERROR] error setting data before clusterWaitForWorkRequest() error: %v", setDataErr)
		}
	}
	return nil
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud) Update() error {
	clusterId, _ := s.D.GetOkExists("cluster_id")
	tmpClusterId := clusterId.(string)
	// if no change, do nothing
	if s.D.HasChange("is_public_api_endpoint_decommissioned") {
		var operation bool
		if is_public_api_endpoint_decommissioned, ok := s.D.GetOkExists("is_public_api_endpoint_decommissioned"); ok {
			operation = is_public_api_endpoint_decommissioned.(bool)
		}
		if operation {
			// do public api endpoint decommission
			request := oci_containerengine.StartPublicApiEndpointDecommissionRequest{}
			request.ClusterId = &tmpClusterId
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
			response, err := s.Client.StartPublicApiEndpointDecommission(context.Background(), request)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId

			// Wait for work request to finish
			_, waitErr := clusterWaitForWorkRequest(workId, "cluster",
				oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
			if waitErr != nil {
				return waitErr
			}
		} else {
			// raise rollback request for public api endpoint decommission
			request := oci_containerengine.RollbackPublicApiEndpointDecommissionRequest{}
			request.ClusterId = &tmpClusterId
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
			response, err := s.Client.RollbackPublicApiEndpointDecommission(context.Background(), request)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId
			// Wait for work request to finish
			_, waitErr := clusterWaitForWorkRequest(workId, "cluster",
				oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
			if waitErr != nil {
				return waitErr
			}
		}
		// store in state file in case of failure
		if setDataErr := s.SetData(); setDataErr != nil {
			log.Printf("[ERROR] error setting data before clusterWaitForWorkRequest() error: %v", setDataErr)
		}
	}

	// if rollback deadline delay is changed, raise request for extendRollback deadline delay
	if rollbackDeadlineDelay, ok := s.D.GetOkExists("rollback_deadline_delay"); ok && s.D.HasChange("rollback_deadline_delay") && rollbackDeadlineDelay != "" {
		requestExtend := oci_containerengine.ExtendEndpointDecommissionRollbackDeadlineRequest{}
		delay := rollbackDeadlineDelay.(string)
		requestExtend.RollbackDeadlineDelay = &delay
		requestExtend.ClusterId = &tmpClusterId

		requestExtend.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
		_, err := s.Client.ExtendEndpointDecommissionRollbackDeadline(context.Background(), requestExtend)
		if err != nil {
			return err
		}
		if setDataErr := s.SetData(); setDataErr != nil {
			log.Printf("[ERROR] error setting data before clusterWaitForWorkRequest() error: %v", setDataErr)
		}
	}
	return nil
}

func commonStringPointer(val interface{}) *string {
	if str, ok := val.(string); ok {
		return &str
	}
	return nil
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionManagerResourceCrud) SetData() error {
	clusterId, _ := s.D.GetOkExists("cluster_id")
	tmpClusterId := clusterId.(string)
	requestGet := oci_containerengine.GetPublicApiEndpointDecommissionStatusRequest{}
	requestGet.ClusterId = &tmpClusterId
	requestGet.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, err := s.Client.GetPublicApiEndpointDecommissionStatus(context.Background(), requestGet)
	if err != nil {
		return err
	}
	s.Res = &responseGet.PublicApiEndpointDecommissionStatus

	d := s.D

	// is_public_api_endpoint_decommissioned -- from status logic
	var status string
	if s.Res != nil {
		status = string(s.Res.Status)
	}
	// Map status to boolean according to your logic
	var isDecommissioned bool
	switch status {
	case "DECOMMISSIONED", "ROLLING_BACK", "ROLLBACK_FAILED", "COMPLETED":
		isDecommissioned = true
	case "DECOMMISSION_FAILED", "IN_PROGRESS", "PENDING":
		isDecommissioned = false
	default:
		isDecommissioned = false // optionally handle unknowns
	}
	d.Set("is_public_api_endpoint_decommissioned", isDecommissioned)

	// rollback_deadline_delay is passthrough
	if val, ok := d.GetOkExists("rollback_deadline_delay"); ok {
		d.Set("rollback_deadline_delay", val)
	}
	return nil
}
