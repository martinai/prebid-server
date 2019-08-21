package martinai

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "martinaitest", NewMartinaiBidder("https://bidder.martin.ai/bid/prebid"))
}
