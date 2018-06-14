// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package integtest

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/database"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/oracle/oci-go-sdk/loadbalancer"
)

const (
	vcnDisplayName               = "GOSDK2_Test_Deps_VCN"
	subnetDisplayName1           = "GOSDK2_Test_Deps_Subnet1"
	subnetDisplayName2           = "GOSDK2_Test_Deps_Subnet2"
	instanceDisplayName          = "GOSDK2_Test_Deps_Instance"
	consoleHistoryDisplayName    = "GOSDK2_Test_Deps_ConsoleHistory"
	crossConnectDisplayName      = "GOSDK2_Test_Deps_CrossConnect"
	crossConnectGroupDisplayName = "GOSDK2_Test_Deps_CrossConnectGroup"
	peeringGatewaysDisplayName   = "GOSDK2_Test_Deps_LocalPeeringGateway"
	virtualCircuitDisplayName    = "GOSDK2_Test_Deps_VirtualCircuits"
	dbSystemDisplayName          = "GOSDK2_Test_Deps_DatabaseSystem"
	dbHomeDisplayName            = "GOSDK2_Test_Deps_DatabaseHome"
	databaseDisplayName          = "GOSDKDB"
	databasePassword             = "OraclE12--"
	dbBackupDisplayName          = "GOSDK2_Test_Deps_DatabaseBackup"
	loadbalancerDisplayName      = "GOSDK2_Test_Deps_Loadbalancer"
	volumeDisplayName            = "GOSDK2_Test_Deps_Volume"
	testUserDisplayName          = "GOSDK2_Test_Deps_TestUser"
	testGroupDisplayName         = "GOSDK2_Test_Deps_TestGroup"
	tagDisplayName               = "GOSDK2_Test_Deps_Tag"
	tagNamespaceDisplayName      = "GOSDK2_Test_Deps_TagNamespcae"
	testImageDisplayName         = "Oracle-Linux-7.4-2018.02.21-1"
)

// a helper method to either create a new vcn or get the one already exist
// this will be used by the test cases which depends on vcn
func createOrGetVcn(t *testing.T) core.Vcn {
	vcnItems := listVcns(t)

	for _, element := range vcnItems {
		if *element.DisplayName == vcnDisplayName {
			assert.NotEmpty(t, element.Id)

			// VCN already created, return it
			return element
		}
	}

	// create a new VCN
	// Notes: here will not destroy it. so for test cases depends on VCN can reuse it in next run
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16")
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(vcnDisplayName)
	request.DnsLabel = common.String("vcndns")

	r, err := c.CreateVcn(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r.Vcn)
	return r.Vcn
}

func listVcns(t *testing.T) []core.Vcn {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.ListVcnsRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	r, err := c.ListVcns(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	return r.Items
}

func createOrGetSubnet(t *testing.T) core.Subnet {
	return createOrGetSubnetWithDetails(
		t,
		common.String(subnetDisplayName1),
		common.String("10.0.0.0/24"),
		common.String("subnetdns1"),
		common.String(validAD()))
}

func createOrGetSubnetWithDetails(t *testing.T, displayName *string, cidrBlock *string, dnsLabel *string, availableDomain *string) core.Subnet {
	subnets := listSubnets(t)

	if displayName == nil {
		displayName = common.String(subnetDisplayName1)
	}

	// check if the subnet has already been created
	for _, element := range subnets {
		if *element.DisplayName == *displayName {
			// find the subnet, return it
			assert.NotEmpty(t, element)
			return element
		}
	}

	// create a new subnet
	vcn := createOrGetVcn(t)
	request := core.CreateSubnetRequest{}
	request.AvailabilityDomain = availableDomain
	request.CompartmentId = common.String(getCompartmentID())
	request.CidrBlock = cidrBlock
	request.VcnId = vcn.Id
	request.DisplayName = displayName
	request.DnsLabel = dnsLabel

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	r, err := c.CreateSubnet(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Subnet)

	getSubnet := func() (interface{}, error) {
		getReq := core.GetSubnetRequest{
			SubnetId: r.Id,
		}

		getResp, err := c.GetSubnet(context.Background(), getReq)

		if err != nil {
			return nil, err
		}

		return getResp, nil
	}

	// wait for lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getSubnet,
			checkLifecycleState(string(core.SubnetLifecycleStateAvailable)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	// update the security rules
	getReq := core.GetSecurityListRequest{
		SecurityListId: common.String(r.SecurityListIds[0]),
	}

	getResp, err := c.GetSecurityList(context.Background(), getReq)
	failIfError(t, err)

	portRange := core.PortRange{
		Max: common.Int(1521),
		Min: common.Int(1521),
	}
	newRules := append(getResp.IngressSecurityRules, core.IngressSecurityRule{
		Protocol: common.String("6"), // TCP
		Source:   common.String("0.0.0.0/0"),
		TcpOptions: &core.TcpOptions{
			DestinationPortRange: &portRange,
		},
	})

	updateReq := core.UpdateSecurityListRequest{
		SecurityListId: common.String(r.SecurityListIds[0]),
	}

	updateReq.IngressSecurityRules = newRules

	_, err = c.UpdateSecurityList(context.Background(), updateReq)
	failIfError(t, err)

	return r.Subnet
}

func listSubnets(t *testing.T) []core.Subnet {
	vcn := createOrGetVcn(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.ListSubnetsRequest{
		CompartmentId: common.String(getCompartmentID()),
		VcnId:         vcn.Id,
	}

	r, err := c.ListSubnets(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	return r.Items
}

// list the available images in the test compartment
func listImages(t *testing.T) []core.Image {
	return listImagesByDisplayName(t, nil)
}

func listImagesByDisplayName(t *testing.T, displayName *string) []core.Image {
	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.ListImagesRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	if displayName != nil {
		request.DisplayName = displayName
	}

	r, err := c.ListImages(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)

	return r.Items
}

// list the boot volumes in the specified compartment and Availability Domain.
func listBootVolumes(t *testing.T) []core.BootVolume {
	// this line make sure boot volumes is created
	createOrGetInstance(t)

	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.ListBootVolumesRequest{
		AvailabilityDomain: common.String(validAD()),
		CompartmentId:      common.String(getCompartmentID()),
	}

	r, err := c.ListBootVolumes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)
	return r.Items
}

func listBootVolumeAttachments(t *testing.T) []core.BootVolumeAttachment {
	// this line make sure boot volumes is created
	createOrGetInstance(t)

	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.ListBootVolumeAttachmentsRequest{
		AvailabilityDomain: common.String(validAD()),
		CompartmentId:      common.String(getCompartmentID()),
	}

	r, err := c.ListBootVolumeAttachments(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r.Items)
	return r.Items
}

func listShapes(t *testing.T) []core.Shape {
	return listShapesForImage(t, nil)
}

func listShapesForImage(t *testing.T, imageID *string) []core.Shape {
	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	if imageID == nil {
		images := listImages(t)

		assert.NotEmpty(t, images)
		assert.NotEqual(t, len(images), 0)
		imageID = images[0].Id
	}

	request := core.ListShapesRequest{
		CompartmentId: common.String(getCompartmentID()),
		ImageId:       imageID,
	}

	r, err := c.ListShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)
	return r.Items
}

func createOrGetInstance(t *testing.T) core.Instance {
	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	listRequest := core.ListInstancesRequest{}
	listRequest.CompartmentId = common.String(getCompartmentID())
	listRequest.AvailabilityDomain = common.String(validAD())
	listRequest.LifecycleState = core.InstanceLifecycleStateRunning

	listResp, err := c.ListInstances(context.Background(), listRequest)
	failIfError(t, err)
	assert.NotEmpty(t, listResp)

	// check if instance exists or not
	for _, element := range listResp.Items {
		if *element.DisplayName == instanceDisplayName {
			return element
		}
	}

	// create a new instance
	createRequest := core.LaunchInstanceRequest{}
	createRequest.CompartmentId = common.String(getCompartmentID())
	createRequest.DisplayName = common.String(instanceDisplayName)
	createRequest.AvailabilityDomain = common.String(validAD())
	createRequest.SubnetId = createOrGetSubnet(t).Id

	// search image by display name to make integration test running more relaible
	// i.e. ServiceLimitExeceed error etc...
	images := listImagesByDisplayName(t, common.String(testImageDisplayName))
	assert.NotEmpty(t, images)
	createRequest.ImageId = images[0].Id

	shapes := listShapesForImage(t, createRequest.ImageId)
	assert.NotEmpty(t, shapes)
	createRequest.Shape = shapes[0].Shape

	createResp, err := c.LaunchInstance(context.Background(), createRequest)
	assert.NotEmpty(t, createResp, fmt.Sprint(createResp))
	failIfError(t, err)

	// get new created instance
	getInstance := func() (interface{}, error) {
		c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
		if clerr != nil {
			return nil, clerr
		}

		request := core.GetInstanceRequest{
			InstanceId: createResp.Instance.Id,
		}

		readResp, err := c.GetInstance(context.Background(), request)

		if err != nil {
			return nil, clerr
		}

		return readResp, err
	}

	// wait for instance lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getInstance,
			checkLifecycleState(string(core.InstanceLifecycleStateRunning)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	return createResp.Instance
}

func createOrGetCrossConnectGroup(t *testing.T) core.CrossConnectGroup {
	crossConnectGrpups := listCrossConnectGroups(t)

	// if connect group exist, return it
	for _, element := range crossConnectGrpups {
		if *element.DisplayName == crossConnectGroupDisplayName {
			return element
		}
	}

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// create a new one
	createRequest := core.CreateCrossConnectGroupRequest{}
	createRequest.CompartmentId = common.String(getCompartmentID())
	createRequest.DisplayName = common.String(crossConnectGroupDisplayName)

	createResp, err := c.CreateCrossConnectGroup(context.Background(), createRequest)
	failIfError(t, err)

	assert.NotEmpty(t, createResp)
	return createResp.CrossConnectGroup
}

func listCrossConnectGroups(t *testing.T) []core.CrossConnectGroup {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListCrossConnectGroupsRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(crossConnectGroupDisplayName)

	r, err := c.ListCrossConnectGroups(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	return r.Items
}

func createOrGetCrossConnect(t *testing.T) core.CrossConnect {
	crossConnects := listCrossConnects(t)

	// if connect group exist, return it
	for _, element := range crossConnects {
		if *element.DisplayName == crossConnectDisplayName {
			return element
		}
	}

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// create a new one
	request := core.CreateCrossConnectRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	crossConnectGroup := createOrGetCrossConnectGroup(t)
	request.CrossConnectGroupId = crossConnectGroup.Id
	request.DisplayName = common.String(crossConnectDisplayName)

	locations := listCrossConnectLocations(t)
	assert.NotEmpty(t, locations)
	request.LocationName = locations[0].Name
	request.PortSpeedShapeName = common.String("10 Gbps")

	resp, err := c.CreateCrossConnect(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	return resp.CrossConnect
}

func listCrossConnects(t *testing.T) []core.CrossConnect {
	crossConnectGroup := createOrGetCrossConnectGroup(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListCrossConnectsRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(crossConnectDisplayName)
	request.CrossConnectGroupId = crossConnectGroup.Id

	resp, err := c.ListCrossConnects(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func listCrossConnectLocations(t *testing.T) []core.CrossConnectLocation {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListCrossConnectLocationsRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListCrossConnectLocations(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func getCrossConnectStatus(t *testing.T) core.CrossConnectStatus {
	crossConnect := createOrGetCrossConnect(t)
	assert.NotEmpty(t, crossConnect)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.GetCrossConnectStatusRequest{}
	request.CrossConnectId = crossConnect.Id

	resp, err := c.GetCrossConnectStatus(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.CrossConnectStatus
}

func listCrossConnectPortSpeedShapes(t *testing.T) []core.CrossConnectPortSpeedShape {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListCrossconnectPortSpeedShapesRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListCrossconnectPortSpeedShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func getCrossConnectLetterOfAuthority(t *testing.T) core.LetterOfAuthority {
	crossConnect := createOrGetCrossConnect(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.GetCrossConnectLetterOfAuthorityRequest{}
	request.CrossConnectId = crossConnect.Id

	resp, err := c.GetCrossConnectLetterOfAuthority(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.LetterOfAuthority
}

func getFastConnectProviderServices(t *testing.T) core.FastConnectProviderService {
	prividerServices := listFastConnectProviderServices(t)
	assert.NotEqual(t, len(prividerServices), 0)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.GetFastConnectProviderServiceRequest{}
	request.ProviderServiceId = prividerServices[0].Id

	resp, err := c.GetFastConnectProviderService(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.FastConnectProviderService
}

func listFastConnectProviderServices(t *testing.T) []core.FastConnectProviderService {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListFastConnectProviderServicesRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListFastConnectProviderServices(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Items)
	return resp.Items
}

func listFastConnectProviderVirtualCircuitBandwidthShapes(t *testing.T) []core.VirtualCircuitBandwidthShape {
	providerService := getFastConnectProviderServices(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListFastConnectProviderVirtualCircuitBandwidthShapesRequest{}
	request.ProviderServiceId = providerService.Id

	resp, err := c.ListFastConnectProviderVirtualCircuitBandwidthShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Items)
	return resp.Items
}

func listVirtualCircuitBandwidthShapes(t *testing.T) []core.VirtualCircuitBandwidthShape {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListVirtualCircuitBandwidthShapesRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListVirtualCircuitBandwidthShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Items)
	return resp.Items
}

func createOrGetInstanceConsoleConnection(t *testing.T) core.InstanceConsoleConnection {
	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	instance := createOrGetInstance(t)

	listRequest := core.ListInstanceConsoleConnectionsRequest{}
	listRequest.CompartmentId = common.String(getCompartmentID())
	listRequest.InstanceId = instance.Id

	listResp, err := c.ListInstanceConsoleConnections(context.Background(), listRequest)
	failIfError(t, err)
	assert.NotEmpty(t, listResp)

	if listResp.Items != nil && len(listResp.Items) != 0 {
		if listResp.Items[0].LifecycleState == core.InstanceConsoleConnectionLifecycleStateActive {
			return listResp.Items[0]
		}
	}

	// create a new one
	createRequest := core.CreateInstanceConsoleConnectionRequest{}
	createRequest.InstanceId = instance.Id

	// get the public key
	buffer, err := readTestPubKey()
	failIfError(t, err)
	createRequest.PublicKey = common.String(string(buffer))

	createResp, err := c.CreateInstanceConsoleConnection(context.Background(), createRequest)
	failIfError(t, err)
	assert.NotEmpty(t, createResp)

	getInstanceConsoleConnection := func() (interface{}, error) {
		getReq := core.GetInstanceConsoleConnectionRequest{
			InstanceConsoleConnectionId: createResp.Id,
		}

		getResp, err := c.GetInstanceConsoleConnection(context.Background(), getReq)

		if err != nil {
			return nil, err
		}

		return getResp, nil
	}

	// wait for instance lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getInstanceConsoleConnection,
			checkLifecycleState(string(core.InstanceConsoleConnectionLifecycleStateActive)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	return createResp.InstanceConsoleConnection
}

func getConsoleHistory(t *testing.T, historyID *string) core.ConsoleHistory {
	assert.NotEmpty(t, historyID)

	// create a new console history
	request := core.GetConsoleHistoryRequest{
		InstanceConsoleHistoryId: historyID,
	}

	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	resp, err := c.GetConsoleHistory(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.ConsoleHistory)
	return resp.ConsoleHistory
}

func captureOrGetConsoleHistory(t *testing.T) core.ConsoleHistory {
	consoleHistories := listConsoleHistories(t)

	for _, element := range consoleHistories {
		assert.NotEmpty(t, element)
		if *element.DisplayName == consoleHistoryDisplayName &&
			element.LifecycleState == core.ConsoleHistoryLifecycleStateSucceeded {
			// find it, return
			return element
		}
	}

	// create a new console history
	instance := createOrGetInstance(t)
	request := core.CaptureConsoleHistoryRequest{}
	request.InstanceId = instance.Id
	request.DisplayName = common.String(consoleHistoryDisplayName)

	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	resp, err := c.CaptureConsoleHistory(context.Background(), request)
	failIfError(t, err)

	getConsoleHistory := func() (interface{}, error) {
		getReq := core.GetConsoleHistoryRequest{
			InstanceConsoleHistoryId: resp.Id,
		}

		getResp, err := c.GetConsoleHistory(context.Background(), getReq)

		if err != nil {
			return nil, err
		}

		return getResp, nil
	}

	// wait for instance lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getConsoleHistory,
			checkLifecycleState(string(core.ConsoleHistoryLifecycleStateSucceeded)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.ConsoleHistory)
	return resp.ConsoleHistory
}

func listConsoleHistories(t *testing.T) []core.ConsoleHistory {
	instance := createOrGetInstance(t)

	request := core.ListConsoleHistoriesRequest{
		CompartmentId:      common.String(getCompartmentID()),
		AvailabilityDomain: common.String(validAD()),
		InstanceId:         instance.Id,
	}

	c, clerr := core.NewComputeClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	resp, err := c.ListConsoleHistories(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	return resp.Items
}

func createOrGetLocalPeeringGateway(t *testing.T) core.LocalPeeringGateway {
	gateways := listLocalPeeringGateways(t)

	for _, element := range gateways {
		assert.NotEmpty(t, element)
		if *element.DisplayName == peeringGatewaysDisplayName {
			return element
		}
	}

	// create a new one
	vcn := createOrGetVcn(t)
	assert.NotEmpty(t, vcn)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.CreateLocalPeeringGatewayRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(peeringGatewaysDisplayName)
	request.VcnId = vcn.Id

	resp, err := c.CreateLocalPeeringGateway(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.LocalPeeringGateway
}

func listLocalPeeringGateways(t *testing.T) []core.LocalPeeringGateway {
	vnc := createOrGetVcn(t)
	assert.NotEmpty(t, vnc)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListLocalPeeringGatewaysRequest{
		CompartmentId: common.String(getCompartmentID()),
		VcnId:         vnc.Id,
	}

	resp, err := c.ListLocalPeeringGateways(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func createOrGetVirtualCircuit(t *testing.T) core.VirtualCircuit {
	virtualCircuits := listVirtualCircuits(t)

	for _, element := range virtualCircuits {
		assert.NotEmpty(t, element)
		if *element.DisplayName == virtualCircuitDisplayName {
			return element
		}
	}

	// create a new one
	request := core.CreateVirtualCircuitRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(virtualCircuitDisplayName)
	request.Type = "PRIVATE"

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	resp, err := c.CreateVirtualCircuit(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.VirtualCircuit
}

func listVirtualCircuits(t *testing.T) []core.VirtualCircuit {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListVirtualCircuitsRequest{
		CompartmentId: common.String(getCompartmentID()),
		DisplayName:   common.String(virtualCircuitDisplayName),
	}

	resp, err := c.ListVirtualCircuits(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func createOrGetLoadBalancer(t *testing.T) *string {
	loadbalancers := listActiveLoadBalancers(t)

	for _, element := range loadbalancers {
		assert.NotEmpty(t, element)
		if *element.DisplayName == loadbalancerDisplayName {
			// found it, return
			return element.Id
		}
	}

	// create new load balancer
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := loadbalancer.CreateLoadBalancerRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(loadbalancerDisplayName)

	subnet1 := createOrGetSubnet(t)

	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, err)
	req := identity.ListAvailabilityDomainsRequest{}
	req.CompartmentId = common.String(getCompartmentID())
	response, err := identityClient.ListAvailabilityDomains(context.Background(), req)
	failIfError(t, err)
	availableDomain := response.Items[1].Name

	subnet2 := createOrGetSubnetWithDetails(t, common.String(subnetDisplayName2), common.String("10.0.1.0/24"), common.String("subnetdns2"), availableDomain)
	request.SubnetIds = []string{*subnet1.Id, *subnet2.Id}

	shapes := listLoadBalancerShapes(t)
	request.ShapeName = shapes[0].Name

	resp, err := c.CreateLoadBalancer(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp.OpcRequestId)

	// get new created loadbalancer
	getTestLoadBalancer := func() *loadbalancer.LoadBalancer {
		loadbalancers = listActiveLoadBalancers(t)
		for _, element := range loadbalancers {
			assert.NotEmpty(t, element)
			if *element.DisplayName == loadbalancerDisplayName {
				// found it, return
				return &element
			}
		}

		return nil
	}

	// use to check the lifecycle of new load balancer
	getLoadBalancerCheck := func() (interface{}, error) {
		testLoadBalancer := getTestLoadBalancer()
		if testLoadBalancer != nil {
			return testLoadBalancer, nil
		}

		return loadbalancer.LoadBalancer{}, nil
	}

	// wait for instance lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getLoadBalancerCheck,
			checkLifecycleState(string(loadbalancer.LoadBalancerLifecycleStateActive)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	newCreatedLoadBalancer := getTestLoadBalancer()
	assert.NotEmpty(t, newCreatedLoadBalancer)
	assert.NotEmpty(t, newCreatedLoadBalancer.Id)
	return newCreatedLoadBalancer.Id
}

func listActiveLoadBalancers(t *testing.T) []loadbalancer.LoadBalancer {
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := loadbalancer.ListLoadBalancersRequest{
		CompartmentId:  common.String(getCompartmentID()),
		DisplayName:    common.String(loadbalancerDisplayName),
		LifecycleState: loadbalancer.LoadBalancerLifecycleStateActive,
	}

	r, err := c.ListLoadBalancers(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	return r.Items
}

func listLoadBalancerShapes(t *testing.T) []loadbalancer.LoadBalancerShape {
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListShapesRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	r, err := c.ListShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)
	assert.NotEqual(t, len(r.Items), 0)
	return r.Items
}

func createOrGetVolume(t *testing.T) core.Volume {
	volumes := listVolumes(t)

	for _, element := range volumes {
		assert.NotEmpty(t, element)
		if *element.DisplayName == volumeDisplayName &&
			element.LifecycleState == core.VolumeLifecycleStateAvailable {
			// found it, return
			return element
		}
	}

	// create a new one
	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.CreateVolumeRequest{}
	request.AvailabilityDomain = common.String(validAD())
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(volumeDisplayName)

	r, err := c.CreateVolume(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Volume)

	getVolume := func() (interface{}, error) {
		getReq := core.GetVolumeRequest{
			VolumeId: r.Volume.Id,
		}

		getResp, err := c.GetVolume(context.Background(), getReq)

		if err != nil {
			return nil, err
		}

		return getResp, nil
	}

	// wait for lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getVolume,
			checkLifecycleState(string(core.VolumeLifecycleStateAvailable)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	return r.Volume
}

func listVolumes(t *testing.T) []core.Volume {
	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	request := core.ListVolumesRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	r, err := c.ListVolumes(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, r)
	return r.Items
}

func createTestUser(t *testing.T, name *string) identity.User {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	req := identity.CreateUserRequest{}
	req.CompartmentId = common.String(getTenancyID())
	req.Name = name
	req.Description = common.String("GoSDK Test User")
	req.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rsp, err := c.CreateUser(context.Background(), req)
	verifyResponseIsValid(t, rsp, err)
	failIfError(t, err)
	return rsp.User
}

func createOrGetTestGroup(t *testing.T) identity.Group {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	listReq := identity.ListGroupsRequest{
		CompartmentId:   common.String(getTenancyID()),
		Limit:           common.Int(500),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}

	listResp, err := c.ListGroups(context.Background(), listReq)
	verifyResponseIsValid(t, listResp, err)
	failIfError(t, err)

	for _, group := range listResp.Items {
		if *group.Name == testGroupDisplayName {
			// found test group, return it
			return group
		}
	}

	group := createTestGroup(t, common.String(testGroupDisplayName))
	return group
}

func createTestGroup(t *testing.T, name *string) identity.Group {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	req := identity.CreateGroupRequest{}
	req.CompartmentId = common.String(getTenancyID())
	req.Name = name
	req.Description = common.String("Go SDK Test Group")
	req.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rsp, err := c.CreateGroup(context.Background(), req)
	verifyResponseIsValid(t, rsp, err)
	failIfError(t, err)
	return rsp.Group
}

func getDatabase(t *testing.T) (*database.DatabaseSummary, error) {
	dbHome, err := getDbHome(t)
	failIfError(t, err)

	c, clerr := getDatabaseClient()
	failIfError(t, clerr)

	listReq := database.ListDatabasesRequest{
		CompartmentId: common.String(getCompartmentID()),
		DbHomeId:      dbHome.Id,
	}

	resp, err := c.ListDatabases(context.Background(), listReq)
	failIfError(t, err)

	for _, element := range resp.Items {
		if *element.DbName == databaseDisplayName &&
			element.LifecycleState == database.DatabaseSummaryLifecycleStateAvailable {
			return &element, nil
		}
	}

	return nil, errors.New("cannot find the dbhome")
}

func createOrGetDBSystem(t *testing.T) *string {
	c, clerr := getDatabaseClient()
	failIfError(t, clerr)

	listReq := database.ListDbSystemsRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	r, err := c.ListDbSystems(context.Background(), listReq)
	failIfError(t, err)

	for _, dbSystem := range r.Items {
		if *dbSystem.DisplayName == dbSystemDisplayName &&
			dbSystem.LifecycleState == database.DbSystemSummaryLifecycleStateAvailable {
			return dbSystem.Id
		}
	}

	return createDBSystem(t, dbSystemDisplayName, databaseDisplayName)
}

func createDBSystem(t *testing.T, dbSystemName string, databaseName string) *string {
	c, clerr := getDatabaseClient()
	failIfError(t, clerr)
	// create a new db system
	request := database.LaunchDbSystemRequest{}
	details := database.LaunchDbSystemDetails{}
	details.AvailabilityDomain = common.String(validAD())
	details.CompartmentId = common.String(getCompartmentID())
	details.CpuCoreCount = common.Int(2)
	details.DatabaseEdition = "ENTERPRISE_EDITION"
	details.DisplayName = common.String(dbSystemName)
	details.Shape = common.String("BM.DenseIO1.36") // this shape will not get service limit error for now

	buffer, err := readTestPubKey()
	failIfError(t, err)
	details.SshPublicKeys = []string{string(buffer)}

	subnet := createOrGetSubnet(t)
	details.SubnetId = subnet.Id
	details.Hostname = common.String("test")

	details.DbHome = &database.CreateDbHomeDetails{
		DbVersion:   common.String("11.2.0.4"),
		DisplayName: common.String(dbHomeDisplayName),
		Database: &database.CreateDatabaseDetails{
			DbName:        common.String(databaseName),
			AdminPassword: common.String(databasePassword),
		},
	}

	request.LaunchDbSystemDetails = details
	resp, err := c.LaunchDbSystem(context.Background(), request)
	failIfError(t, err)

	getDBSystem := func() (interface{}, error) {
		getReq := database.GetDbSystemRequest{
			DbSystemId: resp.Id,
		}

		getResp, err := c.GetDbSystem(context.Background(), getReq)

		if err != nil {
			return nil, err
		}

		return getResp, nil
	}

	// wait for lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getDBSystem,
			checkLifecycleState(string(database.DbSystemSummaryLifecycleStateAvailable)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	return resp.DbSystem.Id
}

func getDbHome(t *testing.T) (*database.DbHomeSummary, error) {
	dbSystemID := createOrGetDBSystem(t)

	c, clerr := getDatabaseClient()
	failIfError(t, clerr)

	listDbHomeReq := database.ListDbHomesRequest{
		CompartmentId: common.String(getCompartmentID()),
		DbSystemId:    dbSystemID,
	}

	r, err := c.ListDbHomes(context.Background(), listDbHomeReq)
	failIfError(t, err)

	for _, element := range r.Items {
		if *element.DisplayName == dbHomeDisplayName &&
			element.LifecycleState == database.DbHomeSummaryLifecycleStateAvailable {
			return &element, nil
		}
	}

	return nil, errors.New("cannot find the dbhome")
}

func createOrGetDatabaseBackup(t *testing.T) *string {
	c, clerr := getDatabaseClient()
	failIfError(t, clerr)

	listReq := database.ListBackupsRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	listResp, err := c.ListBackups(context.Background(), listReq)
	failIfError(t, err)

	for _, element := range listResp.Items {
		if *element.DisplayName == dbBackupDisplayName &&
			element.LifecycleState == database.BackupSummaryLifecycleStateActive {
			return element.Id
		}
	}

	return createDBBackup(t)
}

func createDBBackup(t *testing.T) *string {
	db, err := getDatabase(t)
	failIfError(t, err)
	c, clerr := getDatabaseClient()
	failIfError(t, clerr)

	// create a backup
	req := database.CreateBackupRequest{}
	req.DatabaseId = db.Id
	req.DisplayName = common.String(dbBackupDisplayName)
	r, err := c.CreateBackup(context.Background(), req)
	failIfError(t, err)
	getBackup := func() (interface{}, error) {
		getReq := database.GetBackupRequest{
			BackupId: r.Id,
		}

		getResp, err := c.GetBackup(context.Background(), getReq)

		if err != nil {
			return nil, err
		}

		return getResp, nil
	}

	// wait for lifecyle become running
	failIfError(
		t,
		retryUntilTrueOrError(
			getBackup,
			checkLifecycleState(string(database.BackupSummaryLifecycleStateActive)),
			time.Tick(10*time.Second),
			time.After((5*time.Minute))))

	return r.Id
}

func createOrGetDataGuardAssociation(t *testing.T) *string {
	c, clerr := getDatabaseClient()
	failIfError(t, clerr)
	db, err := getDatabase(t)
	failIfError(t, err)

	dbsystemID := createDBSystem(t, "GOSDK2_Test_Deps_PeerDbSystem", "DB2")

	defer func() {
		// clean up
		fmt.Println("Deleting DBSystem")
		if dbsystemID != nil {
			rDelete := database.TerminateDbSystemRequest{
				DbSystemId: dbsystemID,
			}

			delRes, err := c.TerminateDbSystem(context.Background(), rDelete)
			failIfError(t, err)
			assert.NotEmpty(t, delRes.OpcRequestId)
		}
	}()

	listReq := database.ListDataGuardAssociationsRequest{
		DatabaseId: db.Id,
	}

	listResp, err := c.ListDataGuardAssociations(context.Background(), listReq)
	failIfError(t, err)

	for _, element := range listResp.Items {
		if element.LifecycleState == database.DataGuardAssociationSummaryLifecycleStateAvailable {
			return element.Id
		}
	}

	// create a new one
	req := database.CreateDataGuardAssociationRequest{
		DatabaseId: db.Id,
	}

	details := database.CreateDataGuardAssociationToExistingDbSystemDetails{
		ProtectionMode:        database.CreateDataGuardAssociationDetailsProtectionModePerformance,
		TransportType:         database.CreateDataGuardAssociationDetailsTransportTypeAsync,
		DatabaseAdminPassword: common.String(databasePassword),
		PeerDbSystemId:        dbsystemID,
	}

	req.CreateDataGuardAssociationDetails = details

	r, err := c.CreateDataGuardAssociation(context.Background(), req)
	failIfError(t, err)
	return r.Id
}

func createOrGetUser(t *testing.T) identity.User {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	listReq := identity.ListUsersRequest{
		CompartmentId: common.String(getTenancyID()),
		Limit:         common.Int(500), // not ideal here, but easy and reduce number of requests
	}

	listResp, err := c.ListUsers(context.Background(), listReq)
	verifyResponseIsValid(t, listResp, err)
	failIfError(t, err)

	for _, user := range listResp.Items {
		if *user.Name == testUserDisplayName {
			// found test user, return it
			return user
		}
	}

	user := createTestUser(t, common.String(testUserDisplayName))
	return user
}

func createOrGetFreeformTags(t *testing.T) map[string]string {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	tagNamespaceID := createOrGetTagNamespace(t)

	listReq := identity.ListTagsRequest{
		TagNamespaceId: tagNamespaceID,
	}

	listResp, err := c.ListTags(context.Background(), listReq)
	failIfError(t, err)

	for _, element := range listResp.Items {
		if *element.Name == tagDisplayName {
			return element.FreeformTags
		}
	}

	req := identity.CreateTagRequest{
		TagNamespaceId: tagNamespaceID,
	}

	req.Name = common.String(tagDisplayName)
	req.Description = common.String("GOSDK Test Tag")
	req.FreeformTags = map[string]string{"GOSDKTagKey": "GOSDKTagValue"}

	resp, err := c.CreateTag(context.Background(), req)
	failIfError(t, err)

	return resp.FreeformTags
}

func createOrGetTagNamespace(t *testing.T) *string {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	listReq := identity.ListTagNamespacesRequest{
		CompartmentId: common.String(getTenancyID()),
		Limit:         common.Int(500),
	}

	listResp, err := c.ListTagNamespaces(context.Background(), listReq)
	failIfError(t, err)

	for _, element := range listResp.Items {
		if *element.Name == tagNamespaceDisplayName {
			return element.Id
		}
	}

	req := identity.CreateTagNamespaceRequest{}
	req.CompartmentId = common.String(getTenancyID())
	req.Name = common.String(tagNamespaceDisplayName)
	req.Description = common.String("GSDK Test Tag Namespace")

	resp, err := c.CreateTagNamespace(context.Background(), req)
	failIfError(t, err)
	assert.NotEmpty(t, resp)

	return resp.Id
}
