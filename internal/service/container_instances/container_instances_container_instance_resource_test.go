package container_instances

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestContainerVolumeToMapContainerOciFssVolume(t *testing.T) {
	name := "volumeFss"
	exportId := "exportId"
	mountTargetId := "mountTargetId"
	mountOption := "nfsvers"
	mountOptionValue := "4.1"
	subnetId := "subnetId"
	isEncryptedInTransit := true

	volume := oci_container_instances.ContainerOciFssVolume{
		Name: &name,
		Export: oci_container_instances.OciFssExportId{
			Id: &exportId,
		},
		MountCommand: &oci_container_instances.OciFssMountCommand{
			MountOptions: []oci_container_instances.OciFssMountOption{{
				Option: &mountOption,
				Value:  &mountOptionValue,
			}},
		},
		MountTarget: oci_container_instances.OciFssMountTargetId{
			Id: &mountTargetId,
		},
		Security: oci_container_instances.OciFssSysSecurity{
			IsEncryptedInTransit: &isEncryptedInTransit,
		},
		SubnetId: &subnetId,
	}

	result := ContainerVolumeToMap(volume, nil, 0)
	if result == nil {
		t.Fatal("expected FSS volume to map to state")
	}
	assertEqual(t, result["volume_type"], "OCI_FSS_FILE_SYSTEM")
	assertEqual(t, result["name"], name)
	assertEqual(t, result["subnet_id"], subnetId)

	export := firstNestedMap(t, result, "export")
	assertEqual(t, export["id"], exportId)
	assertEqual(t, export["oci_fss_export_type"], "OCID")

	mountCommand := firstNestedMap(t, result, "mount_command")
	mountOptionMap := firstNestedMap(t, mountCommand, "mount_options")
	assertEqual(t, mountOptionMap["option"], mountOption)
	assertEqual(t, mountOptionMap["value"], mountOptionValue)

	mountTarget := firstNestedMap(t, result, "mount_target")
	assertEqual(t, mountTarget["id"], mountTargetId)
	assertEqual(t, mountTarget["oci_fss_mount_target_type"], "OCID")

	security := firstNestedMap(t, result, "security")
	assertEqual(t, security["auth"], "SYS")
	assertEqual(t, security["is_encrypted_in_transit"], true)
}

func TestMapToCreateContainerOciFssVolumeDetailsIncludesEncryption(t *testing.T) {
	resourceData := schema.TestResourceDataRaw(t, ContainerInstancesContainerInstanceResource().Schema, map[string]interface{}{
		"volumes": []interface{}{map[string]interface{}{
			"name":        "volumeFss",
			"volume_type": "OCI_FSS_FILE_SYSTEM",
			"export": []interface{}{map[string]interface{}{
				"id":                  "exportId",
				"oci_fss_export_type": "OCID",
			}},
			"mount_target": []interface{}{map[string]interface{}{
				"id":                        "mountTargetId",
				"oci_fss_mount_target_type": "OCID",
			}},
			"security": []interface{}{map[string]interface{}{
				"auth":                    "SYS",
				"is_encrypted_in_transit": true,
			}},
		}},
	})
	crud := ContainerInstancesContainerInstanceResourceCrud{BaseCrud: tfresource.BaseCrud{D: resourceData}}

	details, err := crud.mapToCreateContainerVolumeDetails("volumes.0.%s")
	if err != nil {
		t.Fatalf("unexpected error mapping FSS volume: %v", err)
	}
	fssVolume, ok := details.(oci_container_instances.CreateContainerOciFssVolumeDetails)
	if !ok {
		t.Fatalf("expected CreateContainerOciFssVolumeDetails, got %T", details)
	}
	security, ok := fssVolume.Security.(oci_container_instances.CreateOciFssSysSecurityDetails)
	if !ok {
		t.Fatalf("expected CreateOciFssSysSecurityDetails, got %T", fssVolume.Security)
	}
	if security.IsEncryptedInTransit == nil || !*security.IsEncryptedInTransit {
		t.Fatalf("expected is_encrypted_in_transit to be true, got %v", security.IsEncryptedInTransit)
	}
}

func firstNestedMap(t *testing.T, result map[string]interface{}, key string) map[string]interface{} {
	t.Helper()

	list, ok := result[key].([]interface{})
	if !ok {
		t.Fatalf("expected %s to be []interface{}, got %T", key, result[key])
	}
	if len(list) != 1 {
		t.Fatalf("expected %s to have 1 item, got %d", key, len(list))
	}
	item, ok := list[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected %s item to be map[string]interface{}, got %T", key, list[0])
	}
	return item
}

func assertEqual(t *testing.T, got interface{}, want interface{}) {
	t.Helper()

	if got != want {
		t.Fatalf("got %v (%T), want %v (%T)", got, got, want, want)
	}
}
