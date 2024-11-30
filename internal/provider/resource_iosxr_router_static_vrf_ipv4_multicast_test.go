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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterStaticVRFIPv4Multicast(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "prefix_address", "100.0.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "prefix_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interfaces.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interfaces.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interfaces.0.distance_metric", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interfaces.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interfaces.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.address", "11.11.11.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.tag", "103"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.distance_metric", "144"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.bfd_fast_detect_minimum_interval", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_interface_addresses.0.bfd_fast_detect_multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_addresses.0.address", "100.0.2.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_addresses.0.description", "ip-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_addresses.0.tag", "104"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_addresses.0.distance_metric", "155"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_addresses.0.track", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "nexthop_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.vrf_name", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.distance_metric", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.address", "11.11.11.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.tag", "103"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.distance_metric", "144"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.address", "100.0.2.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.description", "ip-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.tag", "104"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.distance_metric", "155"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.track", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_vrf_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.metric", "10"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterStaticVRFIPv4MulticastConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterStaticVRFIPv4MulticastConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_static_vrf_ipv4_multicast.test",
		ImportState:   true,
		ImportStateId: "VRF2,100.0.1.0,24",
		Check:         resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrRouterStaticVRFIPv4MulticastConfig_minimum() string {
	config := `resource "iosxr_router_static_vrf_ipv4_multicast" "test" {` + "\n"
	config += `	vrf_name = "VRF2"` + "\n"
	config += `	prefix_address = "100.0.1.0"` + "\n"
	config += `	prefix_length = 24` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterStaticVRFIPv4MulticastConfig_all() string {
	config := `resource "iosxr_router_static_vrf_ipv4_multicast" "test" {` + "\n"
	config += `	vrf_name = "VRF2"` + "\n"
	config += `	prefix_address = "100.0.1.0"` + "\n"
	config += `	prefix_length = 24` + "\n"
	config += `	nexthop_interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		description = "interface-description"` + "\n"
	config += `		tag = 100` + "\n"
	config += `		distance_metric = 122` + "\n"
	config += `		permanent = true` + "\n"
	config += `		metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `	nexthop_interface_addresses = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/2"` + "\n"
	config += `		address = "11.11.11.1"` + "\n"
	config += `		description = "interface-description"` + "\n"
	config += `		tag = 103` + "\n"
	config += `		distance_metric = 144` + "\n"
	config += `		permanent = true` + "\n"
	config += `		metric = 10` + "\n"
	config += `		bfd_fast_detect_minimum_interval = 100` + "\n"
	config += `		bfd_fast_detect_multiplier = 3` + "\n"
	config += `		}]` + "\n"
	config += `	nexthop_addresses = [{` + "\n"
	config += `		address = "100.0.2.0"` + "\n"
	config += `		description = "ip-description"` + "\n"
	config += `		tag = 104` + "\n"
	config += `		distance_metric = 155` + "\n"
	config += `		track = "TRACK1"` + "\n"
	config += `		metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `	vrfs = [{` + "\n"
	config += `		vrf_name = "VRF1"` + "\n"
	config += `		nexthop_interfaces = [{` + "\n"
	config += `			interface_name = "GigabitEthernet0/0/0/3"` + "\n"
	config += `			description = "interface-description"` + "\n"
	config += `			tag = 100` + "\n"
	config += `			distance_metric = 122` + "\n"
	config += `			permanent = true` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `		nexthop_interface_addresses = [{` + "\n"
	config += `			interface_name = "GigabitEthernet0/0/0/4"` + "\n"
	config += `			address = "11.11.11.1"` + "\n"
	config += `			description = "interface-description"` + "\n"
	config += `			tag = 103` + "\n"
	config += `			distance_metric = 144` + "\n"
	config += `			permanent = true` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `		nexthop_addresses = [{` + "\n"
	config += `			address = "100.0.2.0"` + "\n"
	config += `			description = "ip-description"` + "\n"
	config += `			tag = 104` + "\n"
	config += `			distance_metric = 155` + "\n"
	config += `			track = "TRACK1"` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `		}]` + "\n"
	config += `}` + "\n"
	return config
}
