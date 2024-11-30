// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrRouterBGPVRF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_auto", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_ip_address_ipv4_address", "14.14.14.14"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "rd_ip_address_index", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "default_information_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "default_metric", "125"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "timers_bgp_keepalive_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "timers_bgp_holdtime", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "bgp_router_id", "22.22.22.22"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "bfd_minimum_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "bfd_multiplier", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.neighbor_address", "10.1.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.remote_as", "65002.100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.use_neighbor_group", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.description", "My Neighbor Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.advertisement_interval_seconds", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.ignore_connected_check", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.ebgp_multihop_maximum_hop_count", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.bfd_minimum_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.bfd_multiplier", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.bfd_fast_detect", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.bfd_fast_detect_strict_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.bfd_fast_detect_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as", "65003"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as_no_prepend", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as_replace_as", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.local_as_dual_as", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.password", "12341C2713181F13253920"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.timers_keepalive_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.timers_holdtime", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.update_source", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf.test", "neighbors.0.ttl_security", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPVRFPrerequisitesConfig + testAccDataSourceIosxrRouterBGPVRFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPVRFPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
	lists = [
		{
			name = "neighbor-groups/neighbor-group"
			key = "neighbor-group-name"
			items = [
				{
					"neighbor-group-name" = "GROUP1"
				},
			]
		},
	]
}

`

func testAccDataSourceIosxrRouterBGPVRFConfig() string {
	config := `resource "iosxr_router_bgp_vrf" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	vrf_name = "VRF2"` + "\n"
	config += `	rd_auto = false` + "\n"
	config += `	rd_ip_address_ipv4_address = "14.14.14.14"` + "\n"
	config += `	rd_ip_address_index = 3` + "\n"
	config += `	default_information_originate = true` + "\n"
	config += `	default_metric = 125` + "\n"
	config += `	timers_bgp_keepalive_interval = 5` + "\n"
	config += `	timers_bgp_holdtime = "20"` + "\n"
	config += `	bgp_router_id = "22.22.22.22"` + "\n"
	config += `	bfd_minimum_interval = 10` + "\n"
	config += `	bfd_multiplier = 4` + "\n"
	config += `	neighbors = [{` + "\n"
	config += `		neighbor_address = "10.1.1.2"` + "\n"
	config += `		remote_as = "65002.100"` + "\n"
	config += `		use_neighbor_group = "GROUP1"` + "\n"
	config += `		description = "My Neighbor Description"` + "\n"
	config += `		advertisement_interval_seconds = 10` + "\n"
	config += `		ignore_connected_check = true` + "\n"
	config += `		ebgp_multihop_maximum_hop_count = 10` + "\n"
	config += `		bfd_minimum_interval = 10` + "\n"
	config += `		bfd_multiplier = 4` + "\n"
	config += `		bfd_fast_detect = true` + "\n"
	config += `		bfd_fast_detect_strict_mode = false` + "\n"
	config += `		bfd_fast_detect_disable = false` + "\n"
	config += `		local_as = "65003"` + "\n"
	config += `		local_as_no_prepend = true` + "\n"
	config += `		local_as_replace_as = true` + "\n"
	config += `		local_as_dual_as = true` + "\n"
	config += `		password = "12341C2713181F13253920"` + "\n"
	config += `		shutdown = false` + "\n"
	config += `		timers_keepalive_interval = 5` + "\n"
	config += `		timers_holdtime = "20"` + "\n"
	config += `		update_source = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		ttl_security = false` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_router_bgp_vrf" "test" {
			as_number = "65001"
			vrf_name = "VRF2"
			depends_on = [iosxr_router_bgp_vrf.test]
		}
	`
	return config
}
