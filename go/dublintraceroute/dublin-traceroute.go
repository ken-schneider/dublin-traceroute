/* SPDX-License-Identifier: BSD-2-Clause */

package dublintraceroute

import (
	"time"

	"github.com/ken-schneider/dublin-traceroute/go/dublintraceroute/results"
)

// default values and constants
const (
	DefaultReadTimeout = time.Millisecond * 3000
)

// DublinTraceroute is the common interface that every Dublin Traceroute
// probe type has to implement
type DublinTraceroute interface {
	Validate() error
	Traceroute() (*results.Results, error)
}
