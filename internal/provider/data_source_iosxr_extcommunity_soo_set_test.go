// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrExtcommunitySOOSet(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_extcommunity_soo_set.test", "rpl", "extcommunity-set soo SITE1\nend-set\n"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrExtcommunitySOOSetConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrExtcommunitySOOSetConfig() string {
	config := `resource "iosxr_extcommunity_soo_set" "test" {` + "\n"
	config += `	set_name = "SITE1"` + "\n"
	config += `	rpl = "extcommunity-set soo SITE1\nend-set\n"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_extcommunity_soo_set" "test" {
			set_name = "SITE1"
			depends_on = [iosxr_extcommunity_soo_set.test]
		}
	`
	return config
}
