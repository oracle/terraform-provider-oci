// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"context"
	"fmt"
	"net/http"

	"github.com/oracle/oci-go-sdk/common"
)

//VirtualNetworkClient a client for VirtualNetwork
type VirtualNetworkClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewVirtualNetworkClientWithConfigurationProvider Creates a new default VirtualNetwork client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewVirtualNetworkClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client VirtualNetworkClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = VirtualNetworkClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *VirtualNetworkClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "iaas", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *VirtualNetworkClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *VirtualNetworkClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkAddVirtualCircuitPublicPrefixes Adds one or more customer public IP prefixes to the specified public virtual circuit.
// Use this operation (and not UpdateVirtualCircuit)
// to add prefixes to the virtual circuit. Oracle must verify the customer's ownership
// of each prefix before traffic for that prefix will flow across the virtual circuit.
func (client VirtualNetworkClient) BulkAddVirtualCircuitPublicPrefixes(ctx context.Context, request BulkAddVirtualCircuitPublicPrefixesRequest, options ...common.RetryPolicyOption) (err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/virtualCircuits/{virtualCircuitId}/actions/bulkAddPublicPrefixes", request)
	if err != nil {
		return
	}

	return client.Call(ctx, &httpRequest, common.CallConfig{})
}

// BulkDeleteVirtualCircuitPublicPrefixes Removes one or more customer public IP prefixes from the specified public virtual circuit.
// Use this operation (and not UpdateVirtualCircuit)
// to remove prefixes from the virtual circuit. When the virtual circuit's state switches
// back to PROVISIONED, Oracle stops advertising the specified prefixes across the connection.
func (client VirtualNetworkClient) BulkDeleteVirtualCircuitPublicPrefixes(ctx context.Context, request BulkDeleteVirtualCircuitPublicPrefixesRequest, options ...common.RetryPolicyOption) (err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/virtualCircuits/{virtualCircuitId}/actions/bulkDeletePublicPrefixes", request)
	if err != nil {
		return
	}

	return client.Call(ctx, &httpRequest, common.CallConfig{})
}

// ConnectLocalPeeringGateways Connects this local peering gateway (LPG) to another one in the same region.
// This operation must be called by the VCN administrator who is designated as
// the *requestor* in the peering relationship. The *acceptor* must implement
// an Identity and Access Management (IAM) policy that gives the requestor permission
// to connect to LPGs in the acceptor's compartment. Without that permission, this
// operation will fail. For more information, see
// [VCN Peering](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/VCNpeering.htm).
func (client VirtualNetworkClient) ConnectLocalPeeringGateways(ctx context.Context, request ConnectLocalPeeringGatewaysRequest, options ...common.RetryPolicyOption) (response ConnectLocalPeeringGatewaysResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/localPeeringGateways/{localPeeringGatewayId}/actions/connect", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateCpe Creates a new virtual Customer-Premises Equipment (CPE) object in the specified compartment. For
// more information, see [IPSec VPNs](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPsec.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want
// the CPE to reside. Notice that the CPE doesn't have to be in the same compartment as the IPSec
// connection or other Networking Service components. If you're not sure which compartment to
// use, put the CPE in the same compartment as the DRG. For more information about
// compartments and access control, see [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You must provide the public IP address of your on-premises router. See
// [Configuring Your On-Premises Router for an IPSec VPN](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/configuringCPE.htm).
// You may optionally specify a *display name* for the CPE, otherwise a default is provided. It does not have to
// be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateCpe(ctx context.Context, request CreateCpeRequest, options ...common.RetryPolicyOption) (response CreateCpeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/cpes", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateCrossConnect Creates a new cross-connect. Oracle recommends you create each cross-connect in a
// CrossConnectGroup so you can use link aggregation
// with the connection.
// After creating the `CrossConnect` object, you need to go the FastConnect location
// and request to have the physical cable installed. For more information, see
// [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).
// For the purposes of access control, you must provide the OCID of the
// compartment where you want the cross-connect to reside. If you're
// not sure which compartment to use, put the cross-connect in the
// same compartment with your VCN. For more information about
// compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the cross-connect.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateCrossConnect(ctx context.Context, request CreateCrossConnectRequest, options ...common.RetryPolicyOption) (response CreateCrossConnectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/crossConnects", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateCrossConnectGroup Creates a new cross-connect group to use with Oracle Cloud Infrastructure
// FastConnect. For more information, see
// [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).
// For the purposes of access control, you must provide the OCID of the
// compartment where you want the cross-connect group to reside. If you're
// not sure which compartment to use, put the cross-connect group in the
// same compartment with your VCN. For more information about
// compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the cross-connect group.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateCrossConnectGroup(ctx context.Context, request CreateCrossConnectGroupRequest, options ...common.RetryPolicyOption) (response CreateCrossConnectGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/crossConnectGroups", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateDhcpOptions Creates a new set of DHCP options for the specified VCN. For more information, see
// DhcpOptions.
// For the purposes of access control, you must provide the OCID of the compartment where you want the set of
// DHCP options to reside. Notice that the set of options doesn't have to be in the same compartment as the VCN,
// subnets, or other Networking Service components. If you're not sure which compartment to use, put the set
// of DHCP options in the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the set of DHCP options, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateDhcpOptions(ctx context.Context, request CreateDhcpOptionsRequest, options ...common.RetryPolicyOption) (response CreateDhcpOptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/dhcps", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateDrg Creates a new Dynamic Routing Gateway (DRG) in the specified compartment. For more information,
// see [Dynamic Routing Gateways (DRGs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingDRGs.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want
// the DRG to reside. Notice that the DRG doesn't have to be in the same compartment as the VCN,
// the DRG attachment, or other Networking Service components. If you're not sure which compartment
// to use, put the DRG in the same compartment as the VCN. For more information about compartments
// and access control, see [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the DRG, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateDrg(ctx context.Context, request CreateDrgRequest, options ...common.RetryPolicyOption) (response CreateDrgResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/drgs", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateDrgAttachment Attaches the specified DRG to the specified VCN. A VCN can be attached to only one DRG at a time,
// and vice versa. The response includes a `DrgAttachment` object with its own OCID. For more
// information about DRGs, see
// [Dynamic Routing Gateways (DRGs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingDRGs.htm).
// You may optionally specify a *display name* for the attachment, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// For the purposes of access control, the DRG attachment is automatically placed into the same compartment
// as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
func (client VirtualNetworkClient) CreateDrgAttachment(ctx context.Context, request CreateDrgAttachmentRequest, options ...common.RetryPolicyOption) (response CreateDrgAttachmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/drgAttachments", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateIPSecConnection Creates a new IPSec connection between the specified DRG and CPE. For more information, see
// [IPSec VPNs](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPsec.htm).
// In the request, you must include at least one static route to the CPE object (you're allowed a maximum
// of 10). For example: 10.0.8.0/16.
// For the purposes of access control, you must provide the OCID of the compartment where you want the
// IPSec connection to reside. Notice that the IPSec connection doesn't have to be in the same compartment
// as the DRG, CPE, or other Networking Service components. If you're not sure which compartment to
// use, put the IPSec connection in the same compartment as the DRG. For more information about
// compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the IPSec connection, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// After creating the IPSec connection, you need to configure your on-premises router
// with tunnel-specific information returned by
// GetIPSecConnectionDeviceConfig.
// For each tunnel, that operation gives you the IP address of Oracle's VPN headend and the shared secret
// (that is, the pre-shared key). For more information, see
// [Configuring Your On-Premises Router for an IPSec VPN](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/configuringCPE.htm).
// To get the status of the tunnels (whether they're up or down), use
// GetIPSecConnectionDeviceStatus.
func (client VirtualNetworkClient) CreateIPSecConnection(ctx context.Context, request CreateIPSecConnectionRequest, options ...common.RetryPolicyOption) (response CreateIPSecConnectionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/ipsecConnections", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateInternetGateway Creates a new Internet Gateway for the specified VCN. For more information, see
// [Connectivity to the Internet](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIGs.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the Internet
// Gateway to reside. Notice that the Internet Gateway doesn't have to be in the same compartment as the VCN or
// other Networking Service components. If you're not sure which compartment to use, put the Internet
// Gateway in the same compartment with the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the Internet Gateway, otherwise a default is provided. It
// does not have to be unique, and you can change it. Avoid entering confidential information.
// For traffic to flow between a subnet and an Internet Gateway, you must create a route rule accordingly in
// the subnet's route table (for example, 0.0.0.0/0 > Internet Gateway). See
// UpdateRouteTable.
// You must specify whether the Internet Gateway is enabled when you create it. If it's disabled, that means no
// traffic will flow to/from the internet even if there's a route rule that enables that traffic. You can later
// use UpdateInternetGateway to easily disable/enable
// the gateway without changing the route rule.
func (client VirtualNetworkClient) CreateInternetGateway(ctx context.Context, request CreateInternetGatewayRequest, options ...common.RetryPolicyOption) (response CreateInternetGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/internetGateways", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateLocalPeeringGateway Creates a new local peering gateway (LPG) for the specified VCN.
func (client VirtualNetworkClient) CreateLocalPeeringGateway(ctx context.Context, request CreateLocalPeeringGatewayRequest, options ...common.RetryPolicyOption) (response CreateLocalPeeringGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/localPeeringGateways", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreatePrivateIp Creates a secondary private IP for the specified VNIC.
// For more information about secondary private IPs, see
// [IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPaddresses.htm).
func (client VirtualNetworkClient) CreatePrivateIp(ctx context.Context, request CreatePrivateIpRequest, options ...common.RetryPolicyOption) (response CreatePrivateIpResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/privateIps", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateRouteTable Creates a new route table for the specified VCN. In the request you must also include at least one route
// rule for the new route table. For information on the number of rules you can have in a route table, see
// [Service Limits](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm). For general information about route
// tables in your VCN and the types of targets you can use in route rules,
// see [Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the route
// table to reside. Notice that the route table doesn't have to be in the same compartment as the VCN, subnets,
// or other Networking Service components. If you're not sure which compartment to use, put the route
// table in the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the route table, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateRouteTable(ctx context.Context, request CreateRouteTableRequest, options ...common.RetryPolicyOption) (response CreateRouteTableResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/routeTables", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateSecurityList Creates a new security list for the specified VCN. For more information
// about security lists, see [Security Lists](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).
// For information on the number of rules you can have in a security list, see
// [Service Limits](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the security
// list to reside. Notice that the security list doesn't have to be in the same compartment as the VCN, subnets,
// or other Networking Service components. If you're not sure which compartment to use, put the security
// list in the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the security list, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateSecurityList(ctx context.Context, request CreateSecurityListRequest, options ...common.RetryPolicyOption) (response CreateSecurityListResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/securityLists", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateSubnet Creates a new subnet in the specified VCN. You can't change the size of the subnet after creation,
// so it's important to think about the size of subnets you need before creating them.
// For more information, see [VCNs and Subnets](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVCNs.htm).
// For information on the number of subnets you can have in a VCN, see
// [Service Limits](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the subnet
// to reside. Notice that the subnet doesn't have to be in the same compartment as the VCN, route tables, or
// other Networking Service components. If you're not sure which compartment to use, put the subnet in
// the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs,
// see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally associate a route table with the subnet. If you don't, the subnet will use the
// VCN's default route table. For more information about route tables, see
// [Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm).
// You may optionally associate a security list with the subnet. If you don't, the subnet will use the
// VCN's default security list. For more information about security lists, see
// [Security Lists](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).
// You may optionally associate a set of DHCP options with the subnet. If you don't, the subnet will use the
// VCN's default set. For more information about DHCP options, see
// [DHCP Options](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingDHCP.htm).
// You may optionally specify a *display name* for the subnet, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// You can also add a DNS label for the subnet, which is required if you want the Internet and
// VCN Resolver to resolve hostnames for instances in the subnet. For more information, see
// [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).
func (client VirtualNetworkClient) CreateSubnet(ctx context.Context, request CreateSubnetRequest, options ...common.RetryPolicyOption) (response CreateSubnetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/subnets", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateVcn Creates a new Virtual Cloud Network (VCN). For more information, see
// [VCNs and Subnets](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVCNs.htm).
// For the VCN you must specify a single, contiguous IPv4 CIDR block. Oracle recommends using one of the
// private IP address ranges specified in [RFC 1918](https://tools.ietf.org/html/rfc1918) (10.0.0.0/8,
// 172.16/12, and 192.168/16). Example: 172.16.0.0/16. The CIDR block can range from /16 to /30, and it
// must not overlap with your on-premises network. You can't change the size of the VCN after creation.
// For the purposes of access control, you must provide the OCID of the compartment where you want the VCN to
// reside. Consult an Oracle Cloud Infrastructure administrator in your organization if you're not sure which
// compartment to use. Notice that the VCN doesn't have to be in the same compartment as the subnets or other
// Networking Service components. For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the VCN, otherwise a default is provided. It does not have to
// be unique, and you can change it. Avoid entering confidential information.
// You can also add a DNS label for the VCN, which is required if you want the instances to use the
// Interent and VCN Resolver option for DNS in the VCN. For more information, see
// [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).
// The VCN automatically comes with a default route table, default security list, and default set of DHCP options.
// The OCID for each is returned in the response. You can't delete these default objects, but you can change their
// contents (that is, change the route rules, security list rules, and so on).
// The VCN and subnets you create are not accessible until you attach an Internet Gateway or set up an IPSec VPN
// or FastConnect. For more information, see
// [Overview of the Networking Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/overview.htm).
func (client VirtualNetworkClient) CreateVcn(ctx context.Context, request CreateVcnRequest, options ...common.RetryPolicyOption) (response CreateVcnResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/vcns", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateVirtualCircuit Creates a new virtual circuit to use with Oracle Cloud
// Infrastructure FastConnect. For more information, see
// [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).
// For the purposes of access control, you must provide the OCID of the
// compartment where you want the virtual circuit to reside. If you're
// not sure which compartment to use, put the virtual circuit in the
// same compartment with the DRG it's using. For more information about
// compartments and access control, see
// [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see
// [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the virtual circuit.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// **Important:** When creating a virtual circuit, you specify a DRG for
// the traffic to flow through. Make sure you attach the DRG to your
// VCN and confirm the VCN's routing sends traffic to the DRG. Otherwise
// traffic will not flow. For more information, see
// [Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm).
func (client VirtualNetworkClient) CreateVirtualCircuit(ctx context.Context, request CreateVirtualCircuitRequest, options ...common.RetryPolicyOption) (response CreateVirtualCircuitResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/virtualCircuits", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteCpe Deletes the specified CPE object. The CPE must not be connected to a DRG. This is an asynchronous
// operation. The CPE's `lifecycleState` will change to TERMINATING temporarily until the CPE is completely
// removed.
func (client VirtualNetworkClient) DeleteCpe(ctx context.Context, request DeleteCpeRequest, options ...common.RetryPolicyOption) (response DeleteCpeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/cpes/{cpeId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteCrossConnect Deletes the specified cross-connect. It must not be mapped to a
// VirtualCircuit.
func (client VirtualNetworkClient) DeleteCrossConnect(ctx context.Context, request DeleteCrossConnectRequest, options ...common.RetryPolicyOption) (response DeleteCrossConnectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/crossConnects/{crossConnectId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteCrossConnectGroup Deletes the specified cross-connect group. It must not contain any
// cross-connects, and it cannot be mapped to a
// VirtualCircuit.
func (client VirtualNetworkClient) DeleteCrossConnectGroup(ctx context.Context, request DeleteCrossConnectGroupRequest, options ...common.RetryPolicyOption) (response DeleteCrossConnectGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/crossConnectGroups/{crossConnectGroupId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteDhcpOptions Deletes the specified set of DHCP options, but only if it's not associated with a subnet. You can't delete a
// VCN's default set of DHCP options.
// This is an asynchronous operation. The state of the set of options will switch to TERMINATING temporarily
// until the set is completely removed.
func (client VirtualNetworkClient) DeleteDhcpOptions(ctx context.Context, request DeleteDhcpOptionsRequest, options ...common.RetryPolicyOption) (response DeleteDhcpOptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/dhcps/{dhcpId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteDrg Deletes the specified DRG. The DRG must not be attached to a VCN or be connected to your on-premise
// network. Also, there must not be a route table that lists the DRG as a target. This is an asynchronous
// operation. The DRG's `lifecycleState` will change to TERMINATING temporarily until the DRG is completely
// removed.
func (client VirtualNetworkClient) DeleteDrg(ctx context.Context, request DeleteDrgRequest, options ...common.RetryPolicyOption) (response DeleteDrgResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/drgs/{drgId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteDrgAttachment Detaches a DRG from a VCN by deleting the corresponding `DrgAttachment`. This is an asynchronous
// operation. The attachment's `lifecycleState` will change to DETACHING temporarily until the attachment
// is completely removed.
func (client VirtualNetworkClient) DeleteDrgAttachment(ctx context.Context, request DeleteDrgAttachmentRequest, options ...common.RetryPolicyOption) (response DeleteDrgAttachmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/drgAttachments/{drgAttachmentId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteIPSecConnection Deletes the specified IPSec connection. If your goal is to disable the IPSec VPN between your VCN and
// on-premises network, it's easiest to simply detach the DRG but keep all the IPSec VPN components intact.
// If you were to delete all the components and then later need to create an IPSec VPN again, you would
// need to configure your on-premises router again with the new information returned from
// CreateIPSecConnection.
// This is an asynchronous operation. The connection's `lifecycleState` will change to TERMINATING temporarily
// until the connection is completely removed.
func (client VirtualNetworkClient) DeleteIPSecConnection(ctx context.Context, request DeleteIPSecConnectionRequest, options ...common.RetryPolicyOption) (response DeleteIPSecConnectionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/ipsecConnections/{ipscId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteInternetGateway Deletes the specified Internet Gateway. The Internet Gateway does not have to be disabled, but
// there must not be a route table that lists it as a target.
// This is an asynchronous operation. The gateway's `lifecycleState` will change to TERMINATING temporarily
// until the gateway is completely removed.
func (client VirtualNetworkClient) DeleteInternetGateway(ctx context.Context, request DeleteInternetGatewayRequest, options ...common.RetryPolicyOption) (response DeleteInternetGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/internetGateways/{igId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteLocalPeeringGateway Deletes the specified local peering gateway (LPG).
// This is an asynchronous operation; the local peering gateway's `lifecycleState` changes to TERMINATING temporarily
// until the local peering gateway is completely removed.
func (client VirtualNetworkClient) DeleteLocalPeeringGateway(ctx context.Context, request DeleteLocalPeeringGatewayRequest, options ...common.RetryPolicyOption) (response DeleteLocalPeeringGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/localPeeringGateways/{localPeeringGatewayId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeletePrivateIp Unassigns and deletes the specified private IP. You must
// specify the object's OCID. The private IP address is returned to
// the subnet's pool of available addresses.
// This operation cannot be used with primary private IPs, which are
// automatically unassigned and deleted when the VNIC is terminated.
// **Important:** If a secondary private IP is the
// [target of a route rule](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm#privateip),
// unassigning it from the VNIC causes that route rule to blackhole and the traffic
// will be dropped.
func (client VirtualNetworkClient) DeletePrivateIp(ctx context.Context, request DeletePrivateIpRequest, options ...common.RetryPolicyOption) (response DeletePrivateIpResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/privateIps/{privateIpId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteRouteTable Deletes the specified route table, but only if it's not associated with a subnet. You can't delete a
// VCN's default route table.
// This is an asynchronous operation. The route table's `lifecycleState` will change to TERMINATING temporarily
// until the route table is completely removed.
func (client VirtualNetworkClient) DeleteRouteTable(ctx context.Context, request DeleteRouteTableRequest, options ...common.RetryPolicyOption) (response DeleteRouteTableResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/routeTables/{rtId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteSecurityList Deletes the specified security list, but only if it's not associated with a subnet. You can't delete
// a VCN's default security list.
// This is an asynchronous operation. The security list's `lifecycleState` will change to TERMINATING temporarily
// until the security list is completely removed.
func (client VirtualNetworkClient) DeleteSecurityList(ctx context.Context, request DeleteSecurityListRequest, options ...common.RetryPolicyOption) (response DeleteSecurityListResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/securityLists/{securityListId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteSubnet Deletes the specified subnet, but only if there are no instances in the subnet. This is an asynchronous
// operation. The subnet's `lifecycleState` will change to TERMINATING temporarily. If there are any
// instances in the subnet, the state will instead change back to AVAILABLE.
func (client VirtualNetworkClient) DeleteSubnet(ctx context.Context, request DeleteSubnetRequest, options ...common.RetryPolicyOption) (response DeleteSubnetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/subnets/{subnetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteVcn Deletes the specified VCN. The VCN must be empty and have no attached gateways. This is an asynchronous
// operation. The VCN's `lifecycleState` will change to TERMINATING temporarily until the VCN is completely
// removed.
func (client VirtualNetworkClient) DeleteVcn(ctx context.Context, request DeleteVcnRequest, options ...common.RetryPolicyOption) (response DeleteVcnResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/vcns/{vcnId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteVirtualCircuit Deletes the specified virtual circuit.
// **Important:** If you're using FastConnect via a provider,
// make sure to also terminate the connection with
// the provider, or else the provider may continue to bill you.
func (client VirtualNetworkClient) DeleteVirtualCircuit(ctx context.Context, request DeleteVirtualCircuitRequest, options ...common.RetryPolicyOption) (response DeleteVirtualCircuitResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/virtualCircuits/{virtualCircuitId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetCpe Gets the specified CPE's information.
func (client VirtualNetworkClient) GetCpe(ctx context.Context, request GetCpeRequest, options ...common.RetryPolicyOption) (response GetCpeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/cpes/{cpeId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetCrossConnect Gets the specified cross-connect's information.
func (client VirtualNetworkClient) GetCrossConnect(ctx context.Context, request GetCrossConnectRequest, options ...common.RetryPolicyOption) (response GetCrossConnectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnects/{crossConnectId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetCrossConnectGroup Gets the specified cross-connect group's information.
func (client VirtualNetworkClient) GetCrossConnectGroup(ctx context.Context, request GetCrossConnectGroupRequest, options ...common.RetryPolicyOption) (response GetCrossConnectGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnectGroups/{crossConnectGroupId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetCrossConnectLetterOfAuthority Gets the Letter of Authority for the specified cross-connect.
func (client VirtualNetworkClient) GetCrossConnectLetterOfAuthority(ctx context.Context, request GetCrossConnectLetterOfAuthorityRequest, options ...common.RetryPolicyOption) (response GetCrossConnectLetterOfAuthorityResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnects/{crossConnectId}/letterOfAuthority", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetCrossConnectStatus Gets the status of the specified cross-connect.
func (client VirtualNetworkClient) GetCrossConnectStatus(ctx context.Context, request GetCrossConnectStatusRequest, options ...common.RetryPolicyOption) (response GetCrossConnectStatusResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnects/{crossConnectId}/status", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetDhcpOptions Gets the specified set of DHCP options.
func (client VirtualNetworkClient) GetDhcpOptions(ctx context.Context, request GetDhcpOptionsRequest, options ...common.RetryPolicyOption) (response GetDhcpOptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dhcps/{dhcpId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetDrg Gets the specified DRG's information.
func (client VirtualNetworkClient) GetDrg(ctx context.Context, request GetDrgRequest, options ...common.RetryPolicyOption) (response GetDrgResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/drgs/{drgId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetDrgAttachment Gets the information for the specified `DrgAttachment`.
func (client VirtualNetworkClient) GetDrgAttachment(ctx context.Context, request GetDrgAttachmentRequest, options ...common.RetryPolicyOption) (response GetDrgAttachmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/drgAttachments/{drgAttachmentId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetFastConnectProviderService Gets the specified provider service.
// For more information, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).
func (client VirtualNetworkClient) GetFastConnectProviderService(ctx context.Context, request GetFastConnectProviderServiceRequest, options ...common.RetryPolicyOption) (response GetFastConnectProviderServiceResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fastConnectProviderServices/{providerServiceId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetIPSecConnection Gets the specified IPSec connection's basic information, including the static routes for the
// on-premises router. If you want the status of the connection (whether it's up or down), use
// GetIPSecConnectionDeviceStatus.
func (client VirtualNetworkClient) GetIPSecConnection(ctx context.Context, request GetIPSecConnectionRequest, options ...common.RetryPolicyOption) (response GetIPSecConnectionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/ipsecConnections/{ipscId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetIPSecConnectionDeviceConfig Gets the configuration information for the specified IPSec connection. For each tunnel, the
// response includes the IP address of Oracle's VPN headend and the shared secret.
func (client VirtualNetworkClient) GetIPSecConnectionDeviceConfig(ctx context.Context, request GetIPSecConnectionDeviceConfigRequest, options ...common.RetryPolicyOption) (response GetIPSecConnectionDeviceConfigResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/ipsecConnections/{ipscId}/deviceConfig", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetIPSecConnectionDeviceStatus Gets the status of the specified IPSec connection (whether it's up or down).
func (client VirtualNetworkClient) GetIPSecConnectionDeviceStatus(ctx context.Context, request GetIPSecConnectionDeviceStatusRequest, options ...common.RetryPolicyOption) (response GetIPSecConnectionDeviceStatusResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/ipsecConnections/{ipscId}/deviceStatus", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetInternetGateway Gets the specified Internet Gateway's information.
func (client VirtualNetworkClient) GetInternetGateway(ctx context.Context, request GetInternetGatewayRequest, options ...common.RetryPolicyOption) (response GetInternetGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/internetGateways/{igId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetLocalPeeringGateway Gets the specified local peering gateway's information.
func (client VirtualNetworkClient) GetLocalPeeringGateway(ctx context.Context, request GetLocalPeeringGatewayRequest, options ...common.RetryPolicyOption) (response GetLocalPeeringGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/localPeeringGateways/{localPeeringGatewayId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetPrivateIp Gets the specified private IP. You must specify the object's OCID.
// Alternatively, you can get the object by using
// ListPrivateIps
// with the private IP address (for example, 10.0.3.3) and subnet OCID.
func (client VirtualNetworkClient) GetPrivateIp(ctx context.Context, request GetPrivateIpRequest, options ...common.RetryPolicyOption) (response GetPrivateIpResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/privateIps/{privateIpId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetRouteTable Gets the specified route table's information.
func (client VirtualNetworkClient) GetRouteTable(ctx context.Context, request GetRouteTableRequest, options ...common.RetryPolicyOption) (response GetRouteTableResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/routeTables/{rtId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetSecurityList Gets the specified security list's information.
func (client VirtualNetworkClient) GetSecurityList(ctx context.Context, request GetSecurityListRequest, options ...common.RetryPolicyOption) (response GetSecurityListResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/securityLists/{securityListId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetSubnet Gets the specified subnet's information.
func (client VirtualNetworkClient) GetSubnet(ctx context.Context, request GetSubnetRequest, options ...common.RetryPolicyOption) (response GetSubnetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/subnets/{subnetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetVcn Gets the specified VCN's information.
func (client VirtualNetworkClient) GetVcn(ctx context.Context, request GetVcnRequest, options ...common.RetryPolicyOption) (response GetVcnResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/vcns/{vcnId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetVirtualCircuit Gets the specified virtual circuit's information.
func (client VirtualNetworkClient) GetVirtualCircuit(ctx context.Context, request GetVirtualCircuitRequest, options ...common.RetryPolicyOption) (response GetVirtualCircuitResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/virtualCircuits/{virtualCircuitId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetVnic Gets the information for the specified virtual network interface card (VNIC).
// You can get the VNIC OCID from the
// ListVnicAttachments
// operation.
func (client VirtualNetworkClient) GetVnic(ctx context.Context, request GetVnicRequest, options ...common.RetryPolicyOption) (response GetVnicResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/vnics/{vnicId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListCpes Lists the Customer-Premises Equipment objects (CPEs) in the specified compartment.
func (client VirtualNetworkClient) ListCpes(ctx context.Context, request ListCpesRequest, options ...common.RetryPolicyOption) (response ListCpesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/cpes", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListCrossConnectGroups Lists the cross-connect groups in the specified compartment.
func (client VirtualNetworkClient) ListCrossConnectGroups(ctx context.Context, request ListCrossConnectGroupsRequest, options ...common.RetryPolicyOption) (response ListCrossConnectGroupsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnectGroups", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListCrossConnectLocations Lists the available FastConnect locations for cross-connect installation. You need
// this information so you can specify your desired location when you create a cross-connect.
func (client VirtualNetworkClient) ListCrossConnectLocations(ctx context.Context, request ListCrossConnectLocationsRequest, options ...common.RetryPolicyOption) (response ListCrossConnectLocationsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnectLocations", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListCrossConnects Lists the cross-connects in the specified compartment. You can filter the list
// by specifying the OCID of a cross-connect group.
func (client VirtualNetworkClient) ListCrossConnects(ctx context.Context, request ListCrossConnectsRequest, options ...common.RetryPolicyOption) (response ListCrossConnectsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnects", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListCrossconnectPortSpeedShapes Lists the available port speeds for cross-connects. You need this information
// so you can specify your desired port speed (that is, shape) when you create a
// cross-connect.
func (client VirtualNetworkClient) ListCrossconnectPortSpeedShapes(ctx context.Context, request ListCrossconnectPortSpeedShapesRequest, options ...common.RetryPolicyOption) (response ListCrossconnectPortSpeedShapesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/crossConnectPortSpeedShapes", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListDhcpOptions Lists the sets of DHCP options in the specified VCN and specified compartment.
// The response includes the default set of options that automatically comes with each VCN,
// plus any other sets you've created.
func (client VirtualNetworkClient) ListDhcpOptions(ctx context.Context, request ListDhcpOptionsRequest, options ...common.RetryPolicyOption) (response ListDhcpOptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/dhcps", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListDrgAttachments Lists the `DrgAttachment` objects for the specified compartment. You can filter the
// results by VCN or DRG.
func (client VirtualNetworkClient) ListDrgAttachments(ctx context.Context, request ListDrgAttachmentsRequest, options ...common.RetryPolicyOption) (response ListDrgAttachmentsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/drgAttachments", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListDrgs Lists the DRGs in the specified compartment.
func (client VirtualNetworkClient) ListDrgs(ctx context.Context, request ListDrgsRequest, options ...common.RetryPolicyOption) (response ListDrgsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/drgs", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListFastConnectProviderServices Lists the service offerings from supported providers. You need this
// information so you can specify your desired provider and service
// offering when you create a virtual circuit.
// For the compartment ID, provide the OCID of your tenancy (the root compartment).
// For more information, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).
func (client VirtualNetworkClient) ListFastConnectProviderServices(ctx context.Context, request ListFastConnectProviderServicesRequest, options ...common.RetryPolicyOption) (response ListFastConnectProviderServicesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fastConnectProviderServices", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListFastConnectProviderVirtualCircuitBandwidthShapes Gets the list of available virtual circuit bandwidth levels for a provider.
// You need this information so you can specify your desired bandwidth level (shape) when you create a virtual circuit.
// For more information about virtual circuits, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).
func (client VirtualNetworkClient) ListFastConnectProviderVirtualCircuitBandwidthShapes(ctx context.Context, request ListFastConnectProviderVirtualCircuitBandwidthShapesRequest, options ...common.RetryPolicyOption) (response ListFastConnectProviderVirtualCircuitBandwidthShapesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fastConnectProviderServices/{providerServiceId}/virtualCircuitBandwidthShapes", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListIPSecConnections Lists the IPSec connections for the specified compartment. You can filter the
// results by DRG or CPE.
func (client VirtualNetworkClient) ListIPSecConnections(ctx context.Context, request ListIPSecConnectionsRequest, options ...common.RetryPolicyOption) (response ListIPSecConnectionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/ipsecConnections", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListInternetGateways Lists the Internet Gateways in the specified VCN and the specified compartment.
func (client VirtualNetworkClient) ListInternetGateways(ctx context.Context, request ListInternetGatewaysRequest, options ...common.RetryPolicyOption) (response ListInternetGatewaysResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/internetGateways", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListLocalPeeringGateways Lists the local peering gateways (LPGs) for the specified VCN and compartment
// (the LPG's compartment).
func (client VirtualNetworkClient) ListLocalPeeringGateways(ctx context.Context, request ListLocalPeeringGatewaysRequest, options ...common.RetryPolicyOption) (response ListLocalPeeringGatewaysResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/localPeeringGateways", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListPrivateIps Lists the PrivateIp objects based
// on one of these filters:
//   - Subnet OCID.
//   - VNIC OCID.
//   - Both private IP address and subnet OCID: This lets
//   you get a `privateIP` object based on its private IP
//   address (for example, 10.0.3.3) and not its OCID. For comparison,
//   GetPrivateIp
//   requires the OCID.
// If you're listing all the private IPs associated with a given subnet
// or VNIC, the response includes both primary and secondary private IPs.
func (client VirtualNetworkClient) ListPrivateIps(ctx context.Context, request ListPrivateIpsRequest, options ...common.RetryPolicyOption) (response ListPrivateIpsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/privateIps", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListRouteTables Lists the route tables in the specified VCN and specified compartment. The response
// includes the default route table that automatically comes with each VCN, plus any route tables
// you've created.
func (client VirtualNetworkClient) ListRouteTables(ctx context.Context, request ListRouteTablesRequest, options ...common.RetryPolicyOption) (response ListRouteTablesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/routeTables", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListSecurityLists Lists the security lists in the specified VCN and compartment.
func (client VirtualNetworkClient) ListSecurityLists(ctx context.Context, request ListSecurityListsRequest, options ...common.RetryPolicyOption) (response ListSecurityListsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/securityLists", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListSubnets Lists the subnets in the specified VCN and the specified compartment.
func (client VirtualNetworkClient) ListSubnets(ctx context.Context, request ListSubnetsRequest, options ...common.RetryPolicyOption) (response ListSubnetsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/subnets", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListVcns Lists the Virtual Cloud Networks (VCNs) in the specified compartment.
func (client VirtualNetworkClient) ListVcns(ctx context.Context, request ListVcnsRequest, options ...common.RetryPolicyOption) (response ListVcnsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/vcns", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListVirtualCircuitBandwidthShapes The deprecated operation lists available bandwidth levels for virtual circuits. For the compartment ID, provide the OCID of your tenancy (the root compartment).
func (client VirtualNetworkClient) ListVirtualCircuitBandwidthShapes(ctx context.Context, request ListVirtualCircuitBandwidthShapesRequest, options ...common.RetryPolicyOption) (response ListVirtualCircuitBandwidthShapesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/virtualCircuitBandwidthShapes", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListVirtualCircuitPublicPrefixes Lists the public IP prefixes and their details for the specified
// public virtual circuit.
func (client VirtualNetworkClient) ListVirtualCircuitPublicPrefixes(ctx context.Context, request ListVirtualCircuitPublicPrefixesRequest, options ...common.RetryPolicyOption) (response ListVirtualCircuitPublicPrefixesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/virtualCircuits/{virtualCircuitId}/publicPrefixes", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListVirtualCircuits Lists the virtual circuits in the specified compartment.
func (client VirtualNetworkClient) ListVirtualCircuits(ctx context.Context, request ListVirtualCircuitsRequest, options ...common.RetryPolicyOption) (response ListVirtualCircuitsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/virtualCircuits", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateCpe Updates the specified CPE's display name.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateCpe(ctx context.Context, request UpdateCpeRequest, options ...common.RetryPolicyOption) (response UpdateCpeResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/cpes/{cpeId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateCrossConnect Updates the specified cross-connect.
func (client VirtualNetworkClient) UpdateCrossConnect(ctx context.Context, request UpdateCrossConnectRequest, options ...common.RetryPolicyOption) (response UpdateCrossConnectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/crossConnects/{crossConnectId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateCrossConnectGroup Updates the specified cross-connect group's display name.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateCrossConnectGroup(ctx context.Context, request UpdateCrossConnectGroupRequest, options ...common.RetryPolicyOption) (response UpdateCrossConnectGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/crossConnectGroups/{crossConnectGroupId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateDhcpOptions Updates the specified set of DHCP options. You can update the display name or the options
// themselves. Avoid entering confidential information.
// Note that the `options` object you provide replaces the entire existing set of options.
func (client VirtualNetworkClient) UpdateDhcpOptions(ctx context.Context, request UpdateDhcpOptionsRequest, options ...common.RetryPolicyOption) (response UpdateDhcpOptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/dhcps/{dhcpId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateDrg Updates the specified DRG's display name. Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateDrg(ctx context.Context, request UpdateDrgRequest, options ...common.RetryPolicyOption) (response UpdateDrgResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/drgs/{drgId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateDrgAttachment Updates the display name for the specified `DrgAttachment`.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateDrgAttachment(ctx context.Context, request UpdateDrgAttachmentRequest, options ...common.RetryPolicyOption) (response UpdateDrgAttachmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/drgAttachments/{drgAttachmentId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateIPSecConnection Updates the display name for the specified IPSec connection.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateIPSecConnection(ctx context.Context, request UpdateIPSecConnectionRequest, options ...common.RetryPolicyOption) (response UpdateIPSecConnectionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/ipsecConnections/{ipscId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateInternetGateway Updates the specified Internet Gateway. You can disable/enable it, or change its display name.
// Avoid entering confidential information.
// If the gateway is disabled, that means no traffic will flow to/from the internet even if there's
// a route rule that enables that traffic.
func (client VirtualNetworkClient) UpdateInternetGateway(ctx context.Context, request UpdateInternetGatewayRequest, options ...common.RetryPolicyOption) (response UpdateInternetGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/internetGateways/{igId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateLocalPeeringGateway Updates the specified local peering gateway (LPG).
func (client VirtualNetworkClient) UpdateLocalPeeringGateway(ctx context.Context, request UpdateLocalPeeringGatewayRequest, options ...common.RetryPolicyOption) (response UpdateLocalPeeringGatewayResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/localPeeringGateways/{localPeeringGatewayId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdatePrivateIp Updates the specified private IP. You must specify the object's OCID.
// Use this operation if you want to:
//   - Move a secondary private IP to a different VNIC in the same subnet.
//   - Change the display name for a secondary private IP.
//   - Change the hostname for a secondary private IP.
// This operation cannot be used with primary private IPs.
// To update the hostname for the primary IP on a VNIC, use
// UpdateVnic.
func (client VirtualNetworkClient) UpdatePrivateIp(ctx context.Context, request UpdatePrivateIpRequest, options ...common.RetryPolicyOption) (response UpdatePrivateIpResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/privateIps/{privateIpId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateRouteTable Updates the specified route table's display name or route rules.
// Avoid entering confidential information.
// Note that the `routeRules` object you provide replaces the entire existing set of rules.
func (client VirtualNetworkClient) UpdateRouteTable(ctx context.Context, request UpdateRouteTableRequest, options ...common.RetryPolicyOption) (response UpdateRouteTableResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/routeTables/{rtId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateSecurityList Updates the specified security list's display name or rules.
// Avoid entering confidential information.
// Note that the `egressSecurityRules` or `ingressSecurityRules` objects you provide replace the entire
// existing objects.
func (client VirtualNetworkClient) UpdateSecurityList(ctx context.Context, request UpdateSecurityListRequest, options ...common.RetryPolicyOption) (response UpdateSecurityListResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/securityLists/{securityListId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateSubnet Updates the specified subnet's display name. Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateSubnet(ctx context.Context, request UpdateSubnetRequest, options ...common.RetryPolicyOption) (response UpdateSubnetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/subnets/{subnetId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateVcn Updates the specified VCN's display name.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateVcn(ctx context.Context, request UpdateVcnRequest, options ...common.RetryPolicyOption) (response UpdateVcnResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/vcns/{vcnId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateVirtualCircuit Updates the specified virtual circuit. This can be called by
// either the customer who owns the virtual circuit, or the
// provider (when provisioning or de-provisioning the virtual
// circuit from their end). The documentation for
// UpdateVirtualCircuitDetails
// indicates who can update each property of the virtual circuit.
// **Important:** If the virtual circuit is working and in the
// PROVISIONED state, updating any of the network-related properties
// (such as the DRG being used, the BGP ASN, and so on) will cause the virtual
// circuit's state to switch to PROVISIONING and the related BGP
// session to go down. After Oracle re-provisions the virtual circuit,
// its state will return to PROVISIONED. Make sure you confirm that
// the associated BGP session is back up. For more information
// about the various states and how to test connectivity, see
// [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).
// To change the list of public IP prefixes for a public virtual circuit,
// use BulkAddVirtualCircuitPublicPrefixes
// and
// BulkDeleteVirtualCircuitPublicPrefixes.
// Updating the list of prefixes does NOT cause the BGP session to go down. However,
// Oracle must verify the customer's ownership of each added prefix before
// traffic for that prefix will flow across the virtual circuit.
func (client VirtualNetworkClient) UpdateVirtualCircuit(ctx context.Context, request UpdateVirtualCircuitRequest, options ...common.RetryPolicyOption) (response UpdateVirtualCircuitResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/virtualCircuits/{virtualCircuitId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateVnic Updates the specified VNIC.
func (client VirtualNetworkClient) UpdateVnic(ctx context.Context, request UpdateVnicRequest, options ...common.RetryPolicyOption) (response UpdateVnicResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/vnics/{vnicId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}
