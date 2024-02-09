// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vn_monitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_vn_monitoring "github.com/oracle/oci-go-sdk/v65/vnmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VnMonitoringPathAnalysiResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createVnMonitoringPathAnalysi,
		Read:     readVnMonitoringPathAnalysi,
		Delete:   deleteVnMonitoringPathAnalysi,
		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ADHOC_QUERY",
					"PERSISTED_QUERY",
				}, true),
			},

			// Optional
			"cache_control": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"destination_endpoint": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_INSTANCE",
								"IP_ADDRESS",
								"LOAD_BALANCER",
								"LOAD_BALANCER_LISTENER",
								"NETWORK_LOAD_BALANCER",
								"NETWORK_LOAD_BALANCER_LISTENER",
								"ON_PREM",
								"SUBNET",
								"VLAN",
								"VNIC",
							}, true),
						},

						// Optional
						"address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"listener_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"network_load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vnic_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"path_analyzer_test_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"protocol_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ICMP",
								"TCP",
								"UDP",
							}, true),
						},

						// Optional
						"destination_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"icmp_code": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"icmp_type": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"source_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"query_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_bi_directional_analysis": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"source_endpoint": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_INSTANCE",
								"IP_ADDRESS",
								"LOAD_BALANCER",
								"LOAD_BALANCER_LISTENER",
								"NETWORK_LOAD_BALANCER",
								"NETWORK_LOAD_BALANCER_LISTENER",
								"ON_PREM",
								"SUBNET",
								"VLAN",
								"VNIC",
							}, true),
						},

						// Optional
						"address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"listener_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"network_load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vnic_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Computed
		},
	}
}

func createVnMonitoringPathAnalysi(d *schema.ResourceData, m interface{}) error {
	sync := &VnMonitoringPathAnalysiResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VnMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readVnMonitoringPathAnalysi(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteVnMonitoringPathAnalysi(d *schema.ResourceData, m interface{}) error {
	return nil
}

type VnMonitoringPathAnalysiResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_vn_monitoring.VnMonitoringClient
	DisableNotFoundRetries bool
}

func (s *VnMonitoringPathAnalysiResourceCrud) ID() string {
	return s.D.Id()
}

func (s *VnMonitoringPathAnalysiResourceCrud) Create() error {
	request := oci_vn_monitoring.GetPathAnalysisRequest{}
	err := s.populateTopLevelPolymorphicGetPathAnalysisRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring")

	response, err := s.Client.GetPathAnalysis(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_vn_monitoring.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_vn_monitoring.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "vn_monitoring") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getPathAnalysiFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring"), oci_vn_monitoring.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *VnMonitoringPathAnalysiResourceCrud) getPathAnalysiFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_vn_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	pathAnalysiId, err := pathAnalysiWaitForWorkRequest(workId, "vn_monitoring",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*pathAnalysiId)

	return nil
}

func pathAnalysiWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "vn_monitoring", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_vn_monitoring.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func pathAnalysiWaitForWorkRequest(wId *string, entityType string, action oci_vn_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_vn_monitoring.VnMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "vn_monitoring")
	retryPolicy.ShouldRetryOperation = pathAnalysiWorkRequestShouldRetryFunc(timeout)

	response := oci_vn_monitoring.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_vn_monitoring.OperationStatusInProgress),
			string(oci_vn_monitoring.OperationStatusAccepted),
			string(oci_vn_monitoring.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_vn_monitoring.OperationStatusSucceeded),
			string(oci_vn_monitoring.OperationStatusFailed),
			string(oci_vn_monitoring.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_vn_monitoring.GetWorkRequestRequest{
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

	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if response.Status == oci_vn_monitoring.OperationStatusFailed || response.Status == oci_vn_monitoring.OperationStatusCanceled {
		return nil, getErrorFromVnMonitoringPathAnalysiWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return wId, nil
}

func getErrorFromVnMonitoringPathAnalysiWorkRequest(client *oci_vn_monitoring.VnMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_vn_monitoring.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_vn_monitoring.ListWorkRequestErrorsRequest{
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

func (s *VnMonitoringPathAnalysiResourceCrud) SetData() error {
	return nil
}

func (s *VnMonitoringPathAnalysiResourceCrud) mapToEndpoint(fieldKeyFormat string) (oci_vn_monitoring.Endpoint, error) {
	var baseObject oci_vn_monitoring.Endpoint
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("COMPUTE_INSTANCE"):
		details := oci_vn_monitoring.ComputeInstanceEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if vnicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_id")); ok {
			tmp := vnicId.(string)
			details.VnicId = &tmp
		}
		baseObject = details
	case strings.ToLower("IP_ADDRESS"):
		details := oci_vn_monitoring.IpAddressEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		baseObject = details
	case strings.ToLower("LOAD_BALANCER"):
		details := oci_vn_monitoring.LoadBalancerEndpoint{}
		if loadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "load_balancer_id")); ok {
			tmp := loadBalancerId.(string)
			details.LoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("LOAD_BALANCER_LISTENER"):
		details := oci_vn_monitoring.LoadBalancerListenerEndpoint{}
		if listenerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_id")); ok {
			tmp := listenerId.(string)
			details.ListenerId = &tmp
		}
		if loadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "load_balancer_id")); ok {
			tmp := loadBalancerId.(string)
			details.LoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("NETWORK_LOAD_BALANCER"):
		details := oci_vn_monitoring.NetworkLoadBalancerEndpoint{}
		if networkLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_load_balancer_id")); ok {
			tmp := networkLoadBalancerId.(string)
			details.NetworkLoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("NETWORK_LOAD_BALANCER_LISTENER"):
		details := oci_vn_monitoring.NetworkLoadBalancerListenerEndpoint{}
		if listenerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_id")); ok {
			tmp := listenerId.(string)
			details.ListenerId = &tmp
		}
		if networkLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_load_balancer_id")); ok {
			tmp := networkLoadBalancerId.(string)
			details.NetworkLoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("ON_PREM"):
		details := oci_vn_monitoring.OnPremEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		baseObject = details
	case strings.ToLower("SUBNET"):
		details := oci_vn_monitoring.SubnetEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	case strings.ToLower("VLAN"):
		details := oci_vn_monitoring.VlanEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok {
			tmp := vlanId.(string)
			details.VlanId = &tmp
		}
		baseObject = details
	case strings.ToLower("VNIC"):
		details := oci_vn_monitoring.VnicEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if vnicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_id")); ok {
			tmp := vnicId.(string)
			details.VnicId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *VnMonitoringPathAnalysiResourceCrud) mapToProtocolParameters(fieldKeyFormat string) (oci_vn_monitoring.ProtocolParameters, error) {
	var baseObject oci_vn_monitoring.ProtocolParameters
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ICMP"):
		details := oci_vn_monitoring.IcmpProtocolParameters{}
		if icmpCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icmp_code")); ok {
			tmp := icmpCode.(int)
			details.IcmpCode = &tmp
		}
		if icmpType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icmp_type")); ok {
			tmp := icmpType.(int)
			details.IcmpType = &tmp
		}
		baseObject = details
	case strings.ToLower("TCP"):
		details := oci_vn_monitoring.TcpProtocolParameters{}
		if destinationPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_port")); ok {
			tmp := destinationPort.(int)
			details.DestinationPort = &tmp
		}
		if sourcePort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_port")); ok {
			tmp := sourcePort.(int)
			details.SourcePort = &tmp
		}
		baseObject = details
	case strings.ToLower("UDP"):
		details := oci_vn_monitoring.UdpProtocolParameters{}
		if destinationPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_port")); ok {
			tmp := destinationPort.(int)
			details.DestinationPort = &tmp
		}
		if sourcePort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_port")); ok {
			tmp := sourcePort.(int)
			details.SourcePort = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *VnMonitoringPathAnalysiResourceCrud) mapToQueryOptions(fieldKeyFormat string) (oci_vn_monitoring.QueryOptions, error) {
	result := oci_vn_monitoring.QueryOptions{}

	if isBiDirectionalAnalysis, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_bi_directional_analysis")); ok {
		tmp := isBiDirectionalAnalysis.(bool)
		result.IsBiDirectionalAnalysis = &tmp
	}

	return result, nil
}

func (s *VnMonitoringPathAnalysiResourceCrud) populateTopLevelPolymorphicGetPathAnalysisRequest(request *oci_vn_monitoring.GetPathAnalysisRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	if cacheControl, ok := s.D.GetOkExists("cache_control"); ok {
		tmp := cacheControl.(string)
		request.CacheControl = &tmp
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ADHOC_QUERY"):
		details := oci_vn_monitoring.AdhocGetPathAnalysisDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if destinationEndpoint, ok := s.D.GetOkExists("destination_endpoint"); ok {
			if tmpList := destinationEndpoint.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "destination_endpoint", 0)
				tmp, err := s.mapToEndpoint(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DestinationEndpoint = tmp
			}
		}
		if protocol, ok := s.D.GetOkExists("protocol"); ok {
			tmp := protocol.(int)
			details.Protocol = &tmp
		}
		if protocolParameters, ok := s.D.GetOkExists("protocol_parameters"); ok {
			if tmpList := protocolParameters.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "protocol_parameters", 0)
				tmp, err := s.mapToProtocolParameters(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ProtocolParameters = tmp
			}
		}
		if queryOptions, ok := s.D.GetOkExists("query_options"); ok {
			if tmpList := queryOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_options", 0)
				tmp, err := s.mapToQueryOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.QueryOptions = &tmp
			}
		}
		if sourceEndpoint, ok := s.D.GetOkExists("source_endpoint"); ok {
			if tmpList := sourceEndpoint.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_endpoint", 0)
				tmp, err := s.mapToEndpoint(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.SourceEndpoint = tmp
			}
		}
		request.GetPathAnalysisDetails = details
	case strings.ToLower("PERSISTED_QUERY"):
		details := oci_vn_monitoring.PersistedGetPathAnalysisDetails{}
		if pathAnalyzerTestId, ok := s.D.GetOkExists("path_analyzer_test_id"); ok {
			tmp := pathAnalyzerTestId.(string)
			details.PathAnalyzerTestId = &tmp
		}
		request.GetPathAnalysisDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
