package provider

import (
	"testing"
)

func TestUnitAuthType(t *testing.T) {
	cp := &ConfigProvider{}
	type testFormat struct {
		name     string
		gotError bool
	}
	tests := []testFormat{
		{
			name:     "Test AuthType",
			gotError: true,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if _, err := cp.AuthType(); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
}

func TestUnitUserOCID(t *testing.T) {
	type testFormat struct {
		name     string
		cp       *ConfigProvider
		gotError bool
		userOcid string
	}
	tests := []testFormat{
		{
			name:     "Test error in UserOCID",
			cp:       &ConfigProvider{D: &ociPluginProvider{}},
			gotError: true,
			userOcid: "",
		},
		{
			name:     "Test success",
			cp:       &ConfigProvider{D: &ociPluginProvider{userOcid: testUserOCID}},
			gotError: false,
			userOcid: testUserOCID,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		response, err := test.cp.UserOCID()
		if err != nil {
			if !test.gotError {
				t.Errorf("Output error - %q which is not equal to expected error - %t and got userOcid value - %s", err, test.gotError, response)
			}
		} else if test.gotError {
			t.Errorf("Expected an error but there is no error returned")
		} else if test.userOcid != response {
			t.Errorf("Expected tenancyOcid value - %s but got the value - %s\n", test.userOcid, response)
		}
	}
}
