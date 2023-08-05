package utils

import "testing"

func TestUnitGetSDKServiceName(t *testing.T) {
	type args struct {
		clientServiceName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		"client Name oci_adm.ApplicationDependencyManagementClient",
		args{"oci_adm.ApplicationDependencyManagementClient"},
		"adm",
	},
		{
			"service friendly name Name object_storage",
			args{"object_storage"},
			"objectstorage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSDKServiceName(tt.args.clientServiceName); got != tt.want {
				t.Errorf("GetSDKServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
