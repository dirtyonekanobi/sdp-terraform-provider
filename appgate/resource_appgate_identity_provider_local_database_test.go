package appgate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccLocalDatabaseIdentityProviderBasic(t *testing.T) {
	resourceName := "appgate_local_database_identity_provider.local_database_test_resource"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLocalDatabaseIdentityProviderBasic(),
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(resourceName, "type", "LocalDatabase"),
					resource.TestCheckResourceAttr(resourceName, "min_password_length", "16"),
					resource.TestCheckResourceAttr(resourceName, "user_lockout_threshold", "10"),

					resource.TestCheckResourceAttrSet(resourceName, "admin_provider"),
					resource.TestCheckResourceAttrSet(resourceName, "block_local_dns_requests"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.#"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.0.attribute_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.0.claim_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.0.encrypted"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.0.list"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.1.attribute_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.1.claim_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.1.encrypted"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.1.list"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.2.attribute_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.2.claim_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.2.encrypted"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.2.list"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.3.attribute_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.3.claim_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.3.encrypted"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.3.list"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.4.attribute_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.4.claim_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.4.encrypted"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.4.list"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.5.attribute_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.5.claim_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.5.encrypted"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.5.list"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.6.attribute_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.6.claim_name"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.6.encrypted"),
					resource.TestCheckResourceAttrSet(resourceName, "claim_mappings.6.list"),
					resource.TestCheckResourceAttrSet(resourceName, "default"),
					resource.TestCheckResourceAttrSet(resourceName, "inactivity_timeout_minutes"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_pool_v4"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_pool_v6"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "notes"),
					resource.TestCheckResourceAttrSet(resourceName, "on_boarding_two_factor.#"),
					resource.TestCheckResourceAttrSet(resourceName, "on_demand_claim_mappings.#"),
					resource.TestCheckResourceAttrSet(resourceName, "tags.#"),
					resource.TestCheckResourceAttrSet(resourceName, "tags.3551210388"),
				),
			},
			{
				ResourceName:     resourceName,
				ImportState:      true,
				ImportStateCheck: testAccLocalDatabaseIdentityProviderImportStateCheckFunc(1),
			},
		},
	})
}

func testAccCheckLocalDatabaseIdentityProviderBasic() string {
	return fmt.Sprintf(`
data "appgate_ip_pool" "ip_v6_pool" {
  ip_pool_name = "default pool v6"
}

data "appgate_ip_pool" "ip_v4_pool" {
  ip_pool_name = "default pool v4"
}

resource "appgate_local_database_identity_provider" "local_database_test_resource" {
  notes      = "Built-in Identity Provider on local database."
  ip_pool_v4 = data.appgate_ip_pool.ip_v4_pool.id
  ip_pool_v6 = data.appgate_ip_pool.ip_v6_pool.id

  user_lockout_threshold = 10
  min_password_length = 16

  tags = [
    "builtin",
  ]
}
`)
}

func testAccLocalDatabaseIdentityProviderImportStateCheckFunc(expectedStates int) resource.ImportStateCheckFunc {
	return func(s []*terraform.InstanceState) error {
		if len(s) != expectedStates {
			return fmt.Errorf("expected %d states, got %d: %+v", expectedStates, len(s), s)
		}
		return nil
	}
}
