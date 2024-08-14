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

func IntegrationCustomEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createIntegrationCustomEndpoint,
		Read:   readIntegrationCustomEndpoint,
		Update: updateIntegrationCustomEndpoint,
		Delete: deleteIntegrationCustomEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"integration_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_type": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"OCI",
				}, true),
			},
			"dns_zone_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_type": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ORACLE_MANAGED",
					"CUSTOMER_MANAGED",
				}, true),
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

func createIntegrationCustomEndpoint(d *schema.ResourceData, m interface{}) error {
	instanceId := d.Get("integration_instance_id").(string)
	d.SetId(toCustomEndpointId(instanceId))
	sync := &IntegrationCustomEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	if error := tfresource.CreateResource(d, sync); error != nil {
		return error
	}

	return nil
}

func readIntegrationCustomEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationCustomEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateIntegrationCustomEndpoint(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIntegrationCustomEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationCustomEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

func toCustomEndpointId(instanceId string) string {
	return fmt.Sprintf("%s/managed_custom_endpoint", instanceId)
}

type IntegrationCustomEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_integration.IntegrationInstanceClient
	Res                    *OracleManagedCustomEndpoint
	DisableNotFoundRetries bool
}

func (s *IntegrationCustomEndpointResourceCrud) ID() string {
	return toCustomEndpointId(*s.Res.IntegrationInstanceId)
}

func (s *IntegrationCustomEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateUpdating),
	}
}

func (s *IntegrationCustomEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateActive),
	}
}

func (s *IntegrationCustomEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateUpdating),
	}
}

func (s *IntegrationCustomEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateDeleted),
	}
}

func (s *IntegrationCustomEndpointResourceCrud) Create() error {
	request := oci_integration.AddOracleManagedCustomEndpointRequest{}

	if IntegrationInstanceId, ok := s.D.GetOkExists("integration_instance_id"); ok {
		tmp := IntegrationInstanceId.(string)
		request.IntegrationInstanceId = &tmp
	}

	details := oci_integration.AddOracleManagedCustomEndpointDetails{}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		details.Hostname = &tmp
	}

	if dnsType, ok := s.D.GetOkExists("dns_type"); ok {
		tmp := dnsType.(oci_integration.AddOracleManagedCustomEndpointDetailsDnsTypeEnum)
		details.DnsType = tmp
	}

	if dnsZoneName, ok := s.D.GetOkExists("dns_zone_name"); ok {
		tmp := dnsZoneName.(string)
		details.DnsZoneName = &tmp
	}

	request.AddOracleManagedCustomEndpointDetails = details

	response, err := s.Client.AddOracleManagedCustomEndpoint(context.Background(), request)

	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIntegrationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IntegrationCustomEndpointResourceCrud) getIntegrationInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_integration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	integrationInstanceId, err := integrationInstanceWaitForWorkRequest(workId, "integration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	s.D.SetId(toCustomEndpointId(*integrationInstanceId))

	return s.Get()
}

func (s *IntegrationCustomEndpointResourceCrud) Get() error {
	request := oci_integration.GetIntegrationInstanceRequest{}

	if instanceId, ok := s.D.GetOkExists("integration_instance_id"); ok {
		tmp := instanceId.(string)
		request.IntegrationInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.GetIntegrationInstance(context.Background(), request)

	if err != nil {
		return err
	}

	if response.IntegrationInstance.CustomEndpoint != nil {
		tmpId := toCustomEndpointId(*response.IntegrationInstance.Id)

		s.Res = &OracleManagedCustomEndpoint{
			&tmpId,
			response.IntegrationInstance.LifecycleState,
			response.IntegrationInstance.Id,
			response.IntegrationInstance.CustomEndpoint.Hostname,
			(*string)(&response.IntegrationInstance.CustomEndpoint.DnsType),
			response.IntegrationInstance.CustomEndpoint.DnsZoneName,
			(*string)(&response.IntegrationInstance.CustomEndpoint.ManagedType),
		}

	} else {
		s.Res = &OracleManagedCustomEndpoint{
			LifecycleState: oci_integration.IntegrationInstanceLifecycleStateDeleted,
		}
	}

	return nil
}

func (s *IntegrationCustomEndpointResourceCrud) Update() error {
	return nil
}

func (s *IntegrationCustomEndpointResourceCrud) Delete() error {
	request := oci_integration.RemoveOracleManagedCustomEndpointRequest{}

	if IntegrationInstanceId, ok := s.D.GetOkExists("integration_instance_id"); ok {
		tmp := IntegrationInstanceId.(string)
		request.IntegrationInstanceId = &tmp
	}

	response, err := s.Client.RemoveOracleManagedCustomEndpoint(context.Background(), request)

	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIntegrationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete))
}

func (s *IntegrationCustomEndpointResourceCrud) SetData() error {
	s.D.Set("integration_instance_id", s.Res.IntegrationInstanceId)
	s.D.Set("hostname", s.Res.Hostname)
	s.D.Set("dns_type", s.Res.DnsType)
	s.D.Set("dns_zone_name", s.Res.DnsZoneName)
	s.D.Set("managed_type", s.Res.ManagedType)

	return nil
}

type OracleManagedCustomEndpoint struct {
	Id *string `mandatory:"true" json:"id"`

	LifecycleState oci_integration.IntegrationInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	IntegrationInstanceId *string `json:"integrationInstanceId"`

	Hostname *string `mandatory:"true" json:"hostname"`

	DnsType     *string `mandatory:"true" json:"DnsType"`
	DnsZoneName *string `mandatory:"true" json:"DnsZoneName"`
	ManagedType *string `mandatory:"true" json:"ManagedType"`
}
