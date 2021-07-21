// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
	oci_core "github.com/oracle/oci-go-sdk/v45/core"
	oci_identity "github.com/oracle/oci-go-sdk/v45/identity"
)

const (
	shape           = "VM.Standard2.1"
	cidrBlockSubnet = "10.0.0.0/24"
	cidrBlockVcn    = "10.0.0.0/16"
)

var imageIdMap = map[string]string{
	"us-phoenix-1":   "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a",
	"us-ashburn-1":   "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va",
	"eu-frankfurt-1": "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna",
	"uk-london-1":    "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq",
}

func createVolumeInRegion(clients *OracleClients, region string) (string, error) {
	compartment := getEnvSettingWithBlankDefault("compartment_ocid")

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.identityClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&identityClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	identityClient.SetRegion(region)
	listAvailabilityDomainsResponse, err := identityClient.ListAvailabilityDomains(context.Background(),
		oci_identity.ListAvailabilityDomainsRequest{
			CompartmentId: &compartment,
		})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to get the domain name with the error %v", err)
	}
	domain := listAvailabilityDomainsResponse.Items[0].Name

	createVolumeResponse, err := blockStorageClient.CreateVolume(context.Background(), oci_core.CreateVolumeRequest{
		CreateVolumeDetails: oci_core.CreateVolumeDetails{
			AvailabilityDomain: domain,
			CompartmentId:      &compartment,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: getRetryPolicy(false, "core"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to create source Volume with the error %v", err)
	}
	retryPolicy := getRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = conditionShouldRetry(time.Duration(10*time.Minute), volumeAvailableWaitCondition, "core", false)

	_, err = blockStorageClient.GetVolume(context.Background(), oci_core.GetVolumeRequest{
		VolumeId: createVolumeResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for volumeAvailableWaitCondition failed for %s resource with error %v", *createVolumeResponse.Id, err)
	} else {
		log.Printf("[INFO] end of waitTillCondition for resource %s ", *createVolumeResponse.Id)
	}

	return *createVolumeResponse.Id, nil
}

func createVolumeBackupInRegion(clients *OracleClients, region string, volumeId *string) (string, error) {

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	createVolumeBackupResponse, err := blockStorageClient.CreateVolumeBackup(context.Background(), oci_core.CreateVolumeBackupRequest{
		CreateVolumeBackupDetails: oci_core.CreateVolumeBackupDetails{
			VolumeId: volumeId,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to create source VolumeBackup with the error %v", err)

	}

	retryPolicy := getRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = conditionShouldRetry(time.Duration(10*time.Minute), volumeBackupAvailableWaitCondition, "core", false)
	_, err = blockStorageClient.GetVolumeBackup(context.Background(), oci_core.GetVolumeBackupRequest{
		VolumeBackupId: createVolumeBackupResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for volumeBackupAvailableWaitCondition failed for %s resource with error %v", *createVolumeBackupResponse.Id, err)
	} else {
		log.Printf("[INFO] end of waitTillCondition for resource %s ", *createVolumeBackupResponse.Id)
	}
	return *createVolumeBackupResponse.Id, nil

}

func deleteVolumeInRegion(clients *OracleClients, region string, volumeId string) error {

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if volumeId != "" {
		deleteVolumeRequest := oci_core.DeleteVolumeRequest{}
		deleteVolumeRequest.VolumeId = &volumeId

		_, err := blockStorageClient.DeleteVolume(context.Background(), deleteVolumeRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source Volume %s resource with error %v", *deleteVolumeRequest.VolumeId, err)
		}
	}

	return nil
}

func deleteVolumeBackupInRegion(clients *OracleClients, region string, volumeBackupId string) error {

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if volumeBackupId != "" {
		deleteVolumeBackupRequest := oci_core.DeleteVolumeBackupRequest{}
		deleteVolumeBackupRequest.VolumeBackupId = &volumeBackupId

		_, err := blockStorageClient.DeleteVolumeBackup(context.Background(), deleteVolumeBackupRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source VolumeBackup %s resource with error %v", *deleteVolumeBackupRequest.VolumeBackupId, err)
		}
	}

	return nil
}

func volumeAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if volumeResponse, ok := response.Response.(oci_core.GetVolumeResponse); ok {
		return volumeResponse.LifecycleState != oci_core.VolumeLifecycleStateAvailable
	}

	return false
}

func volumeBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if volumeBackupResponse, ok := response.Response.(oci_core.GetVolumeBackupResponse); ok {
		return volumeBackupResponse.LifecycleState != oci_core.VolumeBackupLifecycleStateAvailable
	}

	return false
}

func createBootVolumeInRegion(clients *OracleClients, region string) (string, string, error) {
	compartment := getEnvSettingWithBlankDefault("compartment_ocid")

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", "", fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return "", "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.identityClient().ConfigurationProvider())
	if err != nil {
		return "", "", fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&identityClient.BaseClient)
	if err != nil {
		return "", "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	identityClient.SetRegion(region)
	listAvailabilityDomainsResponse, err := identityClient.ListAvailabilityDomains(context.Background(),
		oci_identity.ListAvailabilityDomainsRequest{
			CompartmentId: &compartment,
		})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to get the domain name with the error %v", err)
	}
	domain := listAvailabilityDomainsResponse.Items[0].Name

	// Create subnet
	networkClient, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(*clients.virtualNetworkClient().ConfigurationProvider())

	if err != nil {
		return "", "", fmt.Errorf("[WARN] cannot configure client for the source region %v", err)
	}

	cidrBlockVcn := cidrBlockVcn
	networkClient.SetRegion(region)
	createVcnResponse, err := networkClient.CreateVcn(context.Background(), oci_core.CreateVcnRequest{
		CreateVcnDetails: oci_core.CreateVcnDetails{
			CidrBlock:     &cidrBlockVcn,
			CompartmentId: &compartment,
		}})

	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to create source VCN with the error %v", err)
	}

	cidrBlockSubnet := cidrBlockSubnet

	createSubnetResponse, err := networkClient.CreateSubnet(context.Background(), oci_core.CreateSubnetRequest{
		CreateSubnetDetails: oci_core.CreateSubnetDetails{
			CompartmentId: &compartment,
			CidrBlock:     &cidrBlockSubnet,
			VcnId:         createVcnResponse.Id,
		},
	})

	computeClient, err := oci_core.NewComputeClientWithConfigurationProvider(*clients.computeClient().ConfigurationProvider())
	if err != nil {
		return "", "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	shape := shape
	computeClient.SetRegion(region)
	imageId := imageIdMap[region]
	createInstanceResponse, err := computeClient.LaunchInstance(context.Background(), oci_core.LaunchInstanceRequest{
		LaunchInstanceDetails: oci_core.LaunchInstanceDetails{
			AvailabilityDomain: domain,
			CompartmentId:      &compartment,
			Shape:              &shape,
			SubnetId:           createSubnetResponse.Id,
			ImageId:            &imageId,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: getRetryPolicy(false, "core"),
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to create source Instance with the error %v", err)
	}
	instanceId := createInstanceResponse.Id

	retryPolicyInstance := getRetryPolicy(false, "core")
	retryPolicyInstance.ShouldRetryOperation = conditionShouldRetry(time.Duration(10*time.Minute), instanceAvailableWaitCondition, "core", false)

	_, err = computeClient.GetInstance(context.Background(), oci_core.GetInstanceRequest{
		InstanceId: instanceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicyInstance,
		},
	})

	listBootVolumeReq, err := computeClient.ListBootVolumeAttachments(context.Background(), oci_core.ListBootVolumeAttachmentsRequest{
		AvailabilityDomain: domain,
		CompartmentId:      &compartment,
		InstanceId:         instanceId,
	})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to ListBootVolumeAttachments with the error %v", err)
	}

	bootVolumeId := listBootVolumeReq.Items[0].BootVolumeId

	retryPolicy := getRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = conditionShouldRetry(time.Duration(10*time.Minute), bootVolumeAvailableWaitCondition, "core", false)

	_, err = blockStorageClient.GetBootVolume(context.Background(), oci_core.GetBootVolumeRequest{
		BootVolumeId: bootVolumeId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] wait for bootVolumeAvailableWaitCondition failed for %s resource with error %v", *bootVolumeId, err)
	} else {
		log.Printf("[INFO] end of waitTillCondition for resource %s ", *bootVolumeId)
	}

	return *instanceId, *bootVolumeId, nil
}

func createBootVolumeBackupInRegion(clients *OracleClients, region string, bootVolumeId *string) (string, error) {

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	createBootVolumeBackupResponse, err := blockStorageClient.CreateBootVolumeBackup(context.Background(), oci_core.CreateBootVolumeBackupRequest{
		CreateBootVolumeBackupDetails: oci_core.CreateBootVolumeBackupDetails{
			BootVolumeId: bootVolumeId,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to create source BootVolumeBackup with the error %v", err)

	}

	retryPolicy := getRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = conditionShouldRetry(time.Duration(10*time.Minute), bootVolumeBackupAvailableWaitCondition, "core", false)
	_, err = blockStorageClient.GetBootVolumeBackup(context.Background(), oci_core.GetBootVolumeBackupRequest{
		BootVolumeBackupId: createBootVolumeBackupResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for bootVolumeBackupAvailableWaitCondition failed for %s resource with error %v", *createBootVolumeBackupResponse.Id, err)
	} else {
		log.Printf("[INFO] end of waitTillCondition for resource %s ", *createBootVolumeBackupResponse.Id)
	}
	return *createBootVolumeBackupResponse.Id, nil

}

func deleteBootVolumeInRegion(clients *OracleClients, region string, bootVolumeId string) error {

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if bootVolumeId != "" {
		deleteBootVolumeRequest := oci_core.DeleteBootVolumeRequest{}
		deleteBootVolumeRequest.BootVolumeId = &bootVolumeId

		_, err := blockStorageClient.DeleteBootVolume(context.Background(), deleteBootVolumeRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source BootVolume %s resource with error %v", *deleteBootVolumeRequest.BootVolumeId, err)
		}
	}

	return nil
}

func deleteBootVolumeBackupInRegion(clients *OracleClients, region string, bootVolumeBackupId string) error {

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if bootVolumeBackupId != "" {
		deleteBootVolumeBackupRequest := oci_core.DeleteBootVolumeBackupRequest{}
		deleteBootVolumeBackupRequest.BootVolumeBackupId = &bootVolumeBackupId

		_, err := blockStorageClient.DeleteBootVolumeBackup(context.Background(), deleteBootVolumeBackupRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source BootVolumeBackup %s resource with error %v", *deleteBootVolumeBackupRequest.BootVolumeBackupId, err)
		}
	}

	return nil
}

func terminateInstanceInRegion(clients *OracleClients, region string, instanceId string) error {
	computeClient, err := oci_core.NewComputeClientWithConfigurationProvider(*clients.computeClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&computeClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}

	computeClient.SetRegion(region)

	if instanceId != "" {
		terminateInstanceRequest := oci_core.TerminateInstanceRequest{}
		terminateInstanceRequest.InstanceId = &instanceId

		_, err := computeClient.TerminateInstance(context.Background(), terminateInstanceRequest)
		if err != nil {
			return fmt.Errorf("failed to terminate instance %s resource with error %v", *terminateInstanceRequest.InstanceId, err)
		}
	}

	return nil
}

func bootVolumeAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if bootVolumeResponse, ok := response.Response.(oci_core.GetBootVolumeResponse); ok {
		return bootVolumeResponse.LifecycleState != oci_core.BootVolumeLifecycleStateAvailable
	}

	return false
}

func instanceAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if instanceResponse, ok := response.Response.(oci_core.GetInstanceResponse); ok {
		return instanceResponse.LifecycleState != oci_core.InstanceLifecycleStateRunning
	}

	return false
}

func bootVolumeBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if bootVolumeBackupResponse, ok := response.Response.(oci_core.GetBootVolumeBackupResponse); ok {
		return bootVolumeBackupResponse.LifecycleState != oci_core.BootVolumeBackupLifecycleStateAvailable
	}

	return false
}
