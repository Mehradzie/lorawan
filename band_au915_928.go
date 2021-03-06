// +build au_915_928

package band

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUplinkAndDownlinkChannels(t *testing.T) {
	Convey("Given a testtable for uplink", t, func() {
		testTable := []struct {
			Channel   int
			Frequency int
			DataRates []int
		}{
			{Channel: 0, Frequency: 915200000, DataRates: []int{0, 1, 2, 3}},
			{Channel: 63, Frequency: 927800000, DataRates: []int{0, 1, 2, 3}},
			{Channel: 64, Frequency: 915900000, DataRates: []int{4}},
			{Channel: 71, Frequency: 927100000, DataRates: []int{4}},
		}

		for _, test := range testTable {
			Convey(fmt.Sprintf("Then channel %d must have frequency %d and data rates %v", test.Channel, test.Frequency, test.DataRates), func() {
				So(UplinkChannelConfiguration[test.Channel].Frequency, ShouldEqual, test.Frequency)
				So(UplinkChannelConfiguration[test.Channel].DataRates, ShouldResemble, test.DataRates)
			})
		}
	})

	Convey("Given a testtable for downlink", t, func() {
		testTable := []struct {
			Frequency    int
			DataRate     int
			ExpFrequency int
			Err          error
		}{
			{Frequency: 915900000, ExpFrequency: 927500000, DataRate: 4},
			{Frequency: 915900000, DataRate: 3, Err: errors.New("lorawan/band: could not get channel number for frequency: 914900000, data rate: 4")},
			{Frequency: 915200000, DataRate: 3, ExpFrequency: 923300000},
		}

		for _, test := range testTable {
			Convey(fmt.Sprintf("Then frequency: %d and data rate: %d must return frequency: %d or error: %v", test.Frequency, test.DataRate, test.ExpFrequency, test.Err), func() {
				freq, err := GetRX1Frequency(test.Frequency, test.DataRate)

				if test.Err != nil {
					So(err, ShouldResemble, test.Err)
				} else {
					So(err, ShouldBeNil)
					So(freq, ShouldEqual, test.ExpFrequency)
				}
			})
		}
	})
}

func TestGetDataRate(t *testing.T) {
	Convey("When iterating over all data rates", t, func() {
		notImplemented := DataRate{}
		for i, d := range DataRateConfiguration {
			if d == notImplemented {
				continue
			}

			Convey(fmt.Sprintf("Then %v should be DR%d (test %d)", d, expected, i), func() {
				dr, err := GetDataRate(d)
				So(err, ShouldBeNil)
				So(dr, ShouldEqual, expected)
			})
		}
	})
}
