// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func createVolumeInRegion(clients *OracleClients, region string) (string, error) {
	compartment := getEnvSettingWithBlankDefault("compartment_ocid")

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient.ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = configureClient(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.identityClient.ConfigurationProvider())
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

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient.ConfigurationProvider())
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

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient.ConfigurationProvider())
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

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.blockstorageClient.ConfigurationProvider())
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
