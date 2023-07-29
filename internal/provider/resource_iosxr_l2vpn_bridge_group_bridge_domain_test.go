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

func TestAccIosxrL2VPNBridgeGroupBridgeDomain(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn_bridge_group_bridge_domain.test", "bridge_domain_name", "BD123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn_bridge_group_bridge_domain.test", "evis.0.vpn_id", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn_bridge_group_bridge_domain.test", "vnis.0.vni_id", "1234"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_l2vpn_bridge_group_bridge_domain.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/bridge/groups/group[group-name=BG123]/bridge-domains/bridge-domain[bridge-domain-name=BD123]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_minimum() string {
	config := `resource "iosxr_l2vpn_bridge_group_bridge_domain" "test" {` + "\n"
	config += `	bridge_group_name = "BG123"` + "\n"
	config += `	bridge_domain_name = "BD123"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_all() string {
	config := `resource "iosxr_l2vpn_bridge_group_bridge_domain" "test" {` + "\n"
	config += `	bridge_group_name = "BG123"` + "\n"
	config += `	bridge_domain_name = "BD123"` + "\n"
	config += `	evis = [{` + "\n"
	config += `		vpn_id = 1234` + "\n"
	config += `		}]` + "\n"
	config += `	vnis = [{` + "\n"
	config += `		vni_id = 1234` + "\n"
	config += `		}]` + "\n"
	config += `}` + "\n"
	return config
}
