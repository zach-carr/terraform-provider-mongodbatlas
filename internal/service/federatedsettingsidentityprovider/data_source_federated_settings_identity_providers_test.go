package federatedsettingsidentityprovider_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/testutil/acc"
)

func TestAccFederatedSettingsIdentityProvidersDS_basic(t *testing.T) {
	var (
		resourceName        = "data.mongodbatlas_federated_settings_identity_providers.test"
		federatedSettingsID = os.Getenv("MONGODB_ATLAS_FEDERATION_SETTINGS_ID")
	)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.PreCheckFederatedSettings(t) },
		ProtoV6ProviderFactories: acc.TestAccProviderV6Factories,
		Steps: []resource.TestStep{
			{
				Config: testAccMongoDBAtlasDataSourceFederatedSettingsIdentityProvidersConfig(federatedSettingsID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMongoDBAtlasFederatedSettingsIdentityProvidersExists(resourceName),

					resource.TestCheckResourceAttrSet(resourceName, "federation_settings_id"),
					resource.TestCheckResourceAttrSet(resourceName, "results.#"),
					resource.TestCheckResourceAttrSet(resourceName, "results.0.acs_url"),
					resource.TestCheckResourceAttrSet(resourceName, "results.1.audience_claim.0"),
					resource.TestCheckResourceAttrSet(resourceName, "results.1.client_id"),
					resource.TestCheckResourceAttrSet(resourceName, "results.1.groups_claim"),
					resource.TestCheckResourceAttrSet(resourceName, "results.1.requested_scopes.0"),
					resource.TestCheckResourceAttrSet(resourceName, "results.1.user_claim"),
				),
			},
		},
	})
}

func testAccMongoDBAtlasDataSourceFederatedSettingsIdentityProvidersConfig(federatedSettingsID string) string {
	return fmt.Sprintf(`
		data "mongodbatlas_federated_settings_identity_providers" "test" {
			federation_settings_id = "%[1]s"
			page_num = 1
			items_per_page = 100
		}
`, federatedSettingsID)
}

func testAccCheckMongoDBAtlasFederatedSettingsIdentityProvidersExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}
		_, _, err := acc.ConnV2().FederatedAuthenticationApi.ListIdentityProviders(context.Background(), rs.Primary.Attributes["federation_settings_id"]).Execute()
		if err != nil {
			return fmt.Errorf("FederatedSettingsIdentityProviders (%s) does not exist", rs.Primary.ID)
		}
		return nil
	}
}
