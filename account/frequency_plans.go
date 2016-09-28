// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/util"
)

// FrequencyPlans
func FrequencyPlans(server string) (map[string]FrequencyPlan, error) {

	var plans map[string]FrequencyPlan
	err := util.GET(server, auth.Public, "/frequency-plans", &plans)
	if err != nil {
		return nil, err
	}

	// fill in the names
	for key, plan := range plans {
		plan.Name = key
		plans[key] = plan
	}

	return plans, nil
}
