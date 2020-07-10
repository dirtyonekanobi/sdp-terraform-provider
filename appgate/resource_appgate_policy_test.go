package appgate

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPolicyBasic(t *testing.T) {
	resourceName := "appgate_policy.test_policy"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPolicyBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "policy-test"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "expression", "return true;\n"),
					resource.TestCheckResourceAttr(resourceName, "notes", "terraform policy notes"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2876187004", "api-created"),
					resource.TestCheckResourceAttr(resourceName, "tags.535570215", "terraform"),
					resource.TestCheckResourceAttr(resourceName, "tamper_proofing", "true"),
				),
			},
		},
	})
}

func testAccCheckPolicyBasic() string {
	return fmt.Sprintf(`
resource "appgate_policy" "test_policy" {
    name = "policy-test"
    notes = "terraform policy notes"
    tags = [
      "terraform",
      "api-created"
    ]
    disabled = false

    expression = <<-EOF
    return true;
  EOF
}
`)
}

func testAccCheckPolicyExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		token := testAccProvider.Meta().(*Client).Token
		api := testAccProvider.Meta().(*Client).API.PoliciesApi

		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		_, _, err := api.PoliciesIdGet(context.Background(), rs.Primary.ID).Authorization(token).Execute()
		if err != nil {
			return fmt.Errorf("error fetching policy with resource %s. %s", resource, err)
		}
		return nil
	}
}

func testAccCheckPolicyDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "appgate_policy" {
			continue
		}

		token := testAccProvider.Meta().(*Client).Token
		api := testAccProvider.Meta().(*Client).API.PoliciesApi

		_, _, err := api.PoliciesIdGet(context.Background(), rs.Primary.ID).Authorization(token).Execute()
		if err == nil {
			return fmt.Errorf("policy still exists, %+v", err)
		}
	}
	return nil
}