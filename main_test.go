package main

import (
	"github.com/go-check/check"
	"testing"
)

type testSuite struct {}

func (t *testSuite) TestStripQuery(c *check.C) {
	var tcs = []struct{
		i, o, err string
	}{
		// Generic URL remains unchanged.
		{
			i: "https://turntbot.com/this/is?a=test",
			o: "https://turntbot.com/this/is?a=test",
		},
		// Non-URL remains unchanged.
		{
			i: "non-url",
			o: "non-url",
		},
		// UTM params get scrubbed.
		{
			i: "https://turntbot.com/this/is?a=test&utm_content=buffercf3b2&utm_medium=social&utm_source=facebook.com&utm_campaign=buffer",
			o: "https://turntbot.com/this/is?a=test",
		},
		// Spotify si param gets scrubbed.
		{
			i: "https://open.spotify.com/track/1xNcBAoUw8Hz6LqK2jt4Ff?si=INdLibJ9R4GBprQGmy1t5g",
			o: "https://open.spotify.com/track/1xNcBAoUw8Hz6LqK2jt4Ff",
		},
		// Amazon product URL gets wiped.
		{
			i: "https://www.amazon.com/Callaway-Supersoft-Balls-Dozen-Finish/dp/B07MB6N6SG/ref=sr_1_3?dchild=1&keywords=pink+golf+balls&qid=1604118163&sr=8-3",
			// Note: the ref param is in the path itself, and I'll allow it for now.
			o: "https://www.amazon.com/Callaway-Supersoft-Balls-Dozen-Finish/dp/B07MB6N6SG/ref=sr_1_3",
		},
		// Amazon search URL gets mostly wiped.
		{
			i: "https://www.amazon.com/s?k=portable+air+conditioner&i=garden&crid=FVHRJQPEH0VG&sprefix=portable+air+co%2Caps%2C226&ref=nb_sb_ss_c_2_15_ts-do-p",
			// Note: the ref param is in the path itself, and I'll allow it for now.
			o: "https://www.amazon.com/s?k=portable+air+conditioner",
		},
		// Invalid URL.
		{
			i: "Test\n",
			// Note: the ref param is in the path itself, and I'll allow it for now.
			o: "",
			err: ".*invalid control character in URL",
		},
	}

	for _, tc := range tcs {
		var o, err = stripQuery(tc.i)

		c.Check(o, check.Equals, tc.o)

		if tc.err == "" {
			c.Check(err, check.IsNil)
		} else {
			c.Check(err, check.ErrorMatches, tc.err)
		}
	}
}

var _ = check.Suite(&testSuite{})

// Hook up runner.
func Test(t *testing.T) { check.TestingT(t) }
