package acctest

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"os"
	"path"

	"testing"
)

type ConfigFunc func(d *schema.ResourceData) (interface{}, error)

func init() {
	TestAccProvider = ProviderTestCopy(func(d *schema.ResourceData) (interface{}, error) {
		return GetTestClients(d), nil
	}).(*schema.Provider)

	TestAccProviders = map[string]terraform.ResourceProvider{
		"oci": TestAccProvider,
	}
}

func assertEnvAvailable(envVar string, t *testing.T) {
	if v := utils.GetEnvSettingWithBlankDefault(envVar); v == "" {
		t.Fatal("TF_VAR_" + envVar + " must be set for acceptance tests")
	}
}

func GetCompartmentIDForLegacyTests() string {
	var compartmentId string
	if compartmentId = utils.GetEnvSettingWithDefault("compartment_ocid", "compartment_ocid"); compartmentId == "compartment_ocid" {
		compartmentId = utils.GetRequiredEnvSetting("compartment_id_for_create")
	}
	return compartmentId
}

func LegacyTestProviderConfig() string {
	// Use the same config as the generated tests.
	config := ProviderTestConfig()

	// Add the 'compartment_id' used by the legacy tests.
	return config + `variable "compartment_id" {
		default = "` + GetCompartmentIDForLegacyTests() + `"
	}
	`
}

// Don't worry, this key is NOT a valid API key
var testPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,9F4D00DEF02B2B75

IbSQEhNjPeRt49jUhZbhAEaAIG4L9IokDksw/P/QdCPXzZT008xzYK/zmxkz7so1
ZwvIYHn07E0Ul6fIHR6kjw/+MD7AWluCN1FLHs3PHc4XF4THUCKFCC90FvGJ2PEs
kEh7oJ4azZA/PH51g4rSgWpYtH5B/S6ioE2eZ9jJ/prH+34pCuOpX4AvXEFl5zue
pjFm5FhsReAhZ/9eCvjgjIWDHKc7PRfinwSydVHQSzgDnuq+GTMzQh6eztS+EuAp
MLg7w0mazTqmPOuMT+mw9SHGaIePGzA9TcwB1y3QgkYsg3Ch20uN/sUymgQ4PEKI
njXLldWDYvFvv1Tv3/8IOjCEodQ4P/5oWz7msrLh3QF+EhF7lQPYO7132e9Hvz3C
hTmcygmVGrPCtOY1jzuqy+/Kmt4Gv8FQpSnO7i8wFvt5v0N26av18RO10CzYY1ut
EV6WvynimFUtg1Lo03cadh7bspNohSXfFLpbNTji5NwHrIa+UQqTw3h4/zSPZHJl
NwHwM2I8N5lcCsqmSbM01+uTRG3QZ5i1BS8fsArHaAcvPyLvOy4mZGKkpuNlLDXo
qrCCsb+0m9jHR2bzx5AGp4impdHm2Qi3vTV3dMe277wqKkU5qfd5yDbL2eTqAYzQ
hXpPmTjquOTNYdbvoNsOg4TCHZv7WCsGY0nNMPrRO7zXCDApA6cKDJzagbqhW5Zu
/yz7sDT2D3wzE2WXUbtIBLevXyF0OS3AL7AgfbcyAviByOfmEb7WCP9jmdCFaLwY
SgNh9AjeOgkEEr/cRg1kBAXt0kuE7By0w+/ODJHZYelG0wg5nxhseA9Kc596XIJl
NyjbL87CXGfXmMoSYYTA4rzbtCDMmee7xHtbWiYKF1VGxNaGkQ5nnZSJLhCaI6rH
AD0XYwxv92j4fIjHqonbY/dlIKPot1t3VRcdnebbZMjAcNZ63n+I/iVla3DJpWLO
1gT50A4H2uEAve+WWFWmDQe2rfg5wwUtVVkot+Tn3McB6RzNqgcs0c+7uNDnDcOB
WtQ1OfniE1TdoFCPfYcDw8ngimw7uMYwp4mZIYtwlk7Z5GFl4YpNQeLOgh368ao4
8HL7EnTZmiU5cMbuaA8cZmUbgBqiQY0DtLF22VquThi0QOeUMJxJ6N1QUPckD3AU
dikEn0gilOsDQ51fnOsgk9J2uCz8rd5bnyUXlIguj5pyz6S7agyYFhRrXessVzHd
3889QM9V82+px5mv4qCvMn6ReYOvC+KSY1hn4ljXsndOM+6hQzD5CZKeL948pXRn
G7nqbG9D44wLklOz6mkIvqLn3qxEFWapl9UK7yfzjoezGoqeNFweadZ10Kp2+Umu
Sa759/2YDCZLDzaVVoLDTHLzi9ejpAkUIXgEFaPNGzQ8DYiL8N2klRozLSlnDEMr
xTHuOMkklNO7SiTluAUBvXrjxfGqe/gwJOHxXQGHC8W6vyhR2BdVx9PKFVebWjlr
gzRMpGgWnjsaz0ldu3uO7ozRxZg8FgdToIzAIaTytpHKI8HvONvPJlYywOMC1gRi
KwX6p26xaVtCV8PbDpF3RHuEJV1NU6PDIhaIHhdL374BiX/KmcJ6yv7tbkczpK+V
-----END RSA PRIVATE KEY-----`

var testKeyFingerPrint = "b4:8a:7d:54:e6:81:04:b2:fa:ce:ba:55:34:dd:00:00"
var testTenancyOCID = "ocid1.tenancy.oc1..faketenancy"
var testUserOCID = "ocid1.user.oc1..fakeuser"

func WriteConfigFile() (string, string, error) {
	dataTpl := `[DEFAULT]
user=%s
fingerprint=%s
tenancy=%s
region=%s
[PROFILE1]
user=%s
fingerprint=%s
key_file=%s
passphrase=%s
[PROFILE2]
user=%s
[PROFILE3]
user=%s
key_file=%s
passphrase=%s
[PROFILE4]
key_file=%s
fingerprint=%s
tenancy=%s
region=%s
security_token_file=%s
`
	keyPath := path.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName, "oci_api_key.pem")
	configPath := path.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
	err := os.MkdirAll(path.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName), 0700)
	if err != nil {
		return "", "", err
	}
	err = utils.WriteTempFile(testPrivateKey, keyPath)
	if err != nil {
		return "", "", err
	}
	data := fmt.Sprintf(dataTpl, "invalid user", "invalid fingerprint", testTenancyOCID, "us-phoenix-1", testUserOCID, testKeyFingerPrint, keyPath, "password", "invalid user2",
		testUserOCID, keyPath, "password")
	err = utils.WriteTempFile(data, configPath)
	return keyPath, configPath, err
}
