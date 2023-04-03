// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integration

import (
	// "bytes"
	"context"
	"fmt"
	"time"

	// "fmt"
	// "log"
	// "strings"
	// "time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	// "github.com/oracle/terraform-provider-oci/internal/utils"

	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	// oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_integration "github.com/oracle/oci-go-sdk/v65/integration"
)

func IntegrationPrivateEndpointOutboundConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createIntegrationPrivateEndpointOutboundConnection,
		Read:   readIntegrationPrivateEndpointOutboundConnection,
		Update: updateIntegrationPrivateEndpointOutboundConnection,
		// Delete: deletePEOC,
		Delete: deleteIntegrationPrivateEndpointOutboundConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"integration_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_integration.IntegrationInstanceLifecycleStateActive),
					string(oci_integration.IntegrationInstanceLifecycleStateInactive),
				}, true),
			},
		},
	}
}

func deletePEOC(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.OracleClients).IntegrationInstanceClient()

	request := oci_integration.ChangePrivateEndpointOutboundConnectionRequest{}

	if IntegrationInstanceId, ok := d.GetOkExists("integration_instance_id"); ok {
		tmp := IntegrationInstanceId.(string)
		request.IntegrationInstanceId = &tmp
	}

	details := oci_integration.ChangePrivateEndpointOutboundConnectionDetails{
		PrivateEndpointOutboundConnection: oci_integration.NoneOutboundConnection{},
	}

	request.ChangePrivateEndpointOutboundConnectionDetails = details

	_, err := c.ChangePrivateEndpointOutboundConnection(context.Background(), request)

	if err != nil {
		return err
	}

	// d.SetId("")

	return nil
}

func createIntegrationPrivateEndpointOutboundConnection(d *schema.ResourceData, m interface{}) error {
	instanceId := d.Get("integration_instance_id").(string)
	d.SetId(toPEId(instanceId))
	sync := &IntegrationPrivateEndpointOutboundConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	if error := tfresource.CreateResource(d, sync); error != nil {
		return error
	}

	return nil
}

func readIntegrationPrivateEndpointOutboundConnection(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationPrivateEndpointOutboundConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateIntegrationPrivateEndpointOutboundConnection(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIntegrationPrivateEndpointOutboundConnection(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationPrivateEndpointOutboundConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)

}

func toPEId(instanceId string) string {
	return fmt.Sprintf("%s/PE", instanceId)
}

type IntegrationPrivateEndpointOutboundConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_integration.IntegrationInstanceClient
	Res                    *PrivateEndpointOutboundConnection
	DisableNotFoundRetries bool
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) ID() string {
	return toPEId(*s.Res.IntegrationInstanceId)
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateUpdating),
	}
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateActive),
	}
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateUpdating),
	}
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateDeleted),
	}
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) Create() error {
	request := oci_integration.ChangePrivateEndpointOutboundConnectionRequest{}

	if IntegrationInstanceId, ok := s.D.GetOkExists("integration_instance_id"); ok {
		tmp := IntegrationInstanceId.(string)
		request.IntegrationInstanceId = &tmp
	}

	details := oci_integration.ChangePrivateEndpointOutboundConnectionDetails{}

	outboundConnection := oci_integration.PrivateEndpointOutboundConnection{}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		outboundConnection.SubnetId = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			outboundConnection.NsgIds = tmp
		}
	}

	details.PrivateEndpointOutboundConnection = outboundConnection

	request.ChangePrivateEndpointOutboundConnectionDetails = details

	response, err := s.Client.ChangePrivateEndpointOutboundConnection(context.Background(), request)

	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIntegrationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) getIntegrationInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_integration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	integrationInstanceId, err := integrationInstanceWaitForWorkRequest(workId, "integration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	s.D.SetId(toPEId(*integrationInstanceId))

	return s.Get()
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) Get() error {
	request := oci_integration.GetIntegrationInstanceRequest{}

	if instanceId, ok := s.D.GetOkExists("integration_instance_id"); ok {
		tmp := instanceId.(string)
		request.IntegrationInstanceId = &tmp
	}
	// tmp := s.D.IntegrationInstanceId()

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.GetIntegrationInstance(context.Background(), request)

	if err != nil {
		return err
	}

	if response.IntegrationInstance.PrivateEndpointOutboundConnection != nil {
		tmpId := toPEId(*response.IntegrationInstance.Id)

		s.Res = &PrivateEndpointOutboundConnection{
			&tmpId,
			response.IntegrationInstance.LifecycleState,
			response.IntegrationInstance.Id,
			response.IntegrationInstance.PrivateEndpointOutboundConnection.(oci_integration.PrivateEndpointOutboundConnection).SubnetId,
			response.IntegrationInstance.PrivateEndpointOutboundConnection.(oci_integration.PrivateEndpointOutboundConnection).NsgIds,
		}

	} else {
		s.Res = &PrivateEndpointOutboundConnection{
			LifecycleState: oci_integration.IntegrationInstanceLifecycleStateDeleted,
		}
	}

	return nil
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) Update() error {
	return nil
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) Delete() error {
	request := oci_integration.ChangePrivateEndpointOutboundConnectionRequest{}

	if IntegrationInstanceId, ok := s.D.GetOkExists("integration_instance_id"); ok {
		tmp := IntegrationInstanceId.(string)
		request.IntegrationInstanceId = &tmp
	}

	details := oci_integration.ChangePrivateEndpointOutboundConnectionDetails{}

	outboundConnection := oci_integration.NoneOutboundConnection{}

	details.PrivateEndpointOutboundConnection = outboundConnection

	request.ChangePrivateEndpointOutboundConnectionDetails = details

	_, err := s.Client.ChangePrivateEndpointOutboundConnection(context.Background(), request)
	return err
}

func (s *IntegrationPrivateEndpointOutboundConnectionResourceCrud) SetData() error {
	s.D.Set("integration_instance_id", s.Res.IntegrationInstanceId)
	s.D.Set("state", s.Res.LifecycleState)
	s.D.Set("nsg_ids", s.Res.NsgIds)
	s.D.Set("subnet_id", s.Res.SubnetId)
	return nil
}

type PrivateEndpointOutboundConnection struct {
	Id *string `mandatory:"true" json:"id"`

	LifecycleState oci_integration.IntegrationInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	IntegrationInstanceId *string `json:"integrationInstanceId"`

	// Customer Private Network VCN Subnet OCID. This is a required argument.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// One or more Network security group Ids. This is an optional argument.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}
