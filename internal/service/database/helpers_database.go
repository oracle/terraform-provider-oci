package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fileSystemConfigDiffFunc(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	// Access the current and proposed values
	oldVal, newVal := diff.GetChange("file_system_configuration_details")

	if oldVal == nil || newVal == nil {
		return nil
	}

	oldList := oldVal.([]interface{})
	newList := newVal.([]interface{})

	tmp1 := make([]map[string]interface{}, len(oldList))
	tmp2 := make([]map[string]interface{}, len(newList))

	for i := range oldList {
		map1 := oldList[i].(map[string]interface{})
		tmp1[i] = map1
	}

	for j := range newList {
		map2 := newList[j].(map[string]interface{})
		tmp2[j] = map2
	}

	//For each item in old list, check if it is present in new list using the mount_point as key
	//and if present, assign the file_system_size_gb value to the old list
	for i := range tmp1 {
		for j := range tmp2 {
			if tmp1[i]["mount_point"] == tmp2[j]["mount_point"] {
				tmp1[i]["file_system_size_gb"] = tmp2[j]["file_system_size_gb"]
			}
		}
	}

	if err := diff.SetNew("file_system_configuration_details", tmp1); err != nil {
		return err
	}

	return nil
}
