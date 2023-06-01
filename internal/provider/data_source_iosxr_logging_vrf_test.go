// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrLoggingVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrLoggingVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv4_addresses.0.ipv4_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv4_addresses.0.severity", "info"),
					resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv6_addresses.0.ipv6_address", "2001::1"),
					resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv6_addresses.0.severity", "info"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrLoggingVRFConfig = `

resource "iosxr_logging_vrf" "test" {
	vrf_name = "VRF1"
	host_ipv4_addresses = [{
		ipv4_address = "1.1.1.1"
		severity = "info"
	}]
	host_ipv6_addresses = [{
		ipv6_address = "2001::1"
		severity = "info"
	}]
}

data "iosxr_logging_vrf" "test" {
	vrf_name = "VRF1"
	depends_on = [iosxr_logging_vrf.test]
}
`
