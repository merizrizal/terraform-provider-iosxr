// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrL2VPN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrL2VPNConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_l2vpn.test", "description", "My L2VPN Description"),
					resource.TestCheckResourceAttr("iosxr_l2vpn.test", "router_id", "1.2.3.4"),
					resource.TestCheckResourceAttr("iosxr_l2vpn.test", "xconnect_groups.0.group_name", "P2P"),
				),
			},
			{
				ResourceName:  "iosxr_l2vpn.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn",
			},
		},
	})
}

func testAccIosxrL2VPNConfig_minimum() string {
	return `
	resource "iosxr_l2vpn" "test" {
	}
	`
}

func testAccIosxrL2VPNConfig_all() string {
	return `
	resource "iosxr_l2vpn" "test" {
		description = "My L2VPN Description"
		router_id = "1.2.3.4"
		xconnect_groups = [{
			group_name = "P2P"
		}]
	}
	`
}
