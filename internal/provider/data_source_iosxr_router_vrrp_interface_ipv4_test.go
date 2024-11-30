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

func TestAccDataSourceIosxrRouterVRRPInterfaceIPv4(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "priority", "250"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "name", "TEST"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "text_authentication", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "timer_advertisement_seconds", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "timer_force", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "preempt_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "preempt_delay", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "accept_mode_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "track_interfaces.0.interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "track_interfaces.0.priority_decrement", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "track_objects.0.object_name", "OBJECT"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "track_objects.0.priority_decrement", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_vrrp_interface_ipv4.test", "bfd_fast_detect_peer_ipv4", "33.33.33.3"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterVRRPInterfaceIPv4Config(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrRouterVRRPInterfaceIPv4Config() string {
	config := `resource "iosxr_router_vrrp_interface_ipv4" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	vrrp_id = 123` + "\n"
	config += `	version = 2` + "\n"
	config += `	address = "1.1.1.1"` + "\n"
	config += `	priority = 250` + "\n"
	config += `	name = "TEST"` + "\n"
	config += `	text_authentication = "7"` + "\n"
	config += `	timer_advertisement_seconds = 123` + "\n"
	config += `	timer_force = false` + "\n"
	config += `	preempt_disable = false` + "\n"
	config += `	preempt_delay = 255` + "\n"
	config += `	accept_mode_disable = false` + "\n"
	config += `	track_interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		priority_decrement = 12` + "\n"
	config += `	}]` + "\n"
	config += `	track_objects = [{` + "\n"
	config += `		object_name = "OBJECT"` + "\n"
	config += `		priority_decrement = 22` + "\n"
	config += `	}]` + "\n"
	config += `	bfd_fast_detect_peer_ipv4 = "33.33.33.3"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_router_vrrp_interface_ipv4" "test" {
			interface_name = "GigabitEthernet0/0/0/1"
			vrrp_id = 123
			version = 2
			depends_on = [iosxr_router_vrrp_interface_ipv4.test]
		}
	`
	return config
}