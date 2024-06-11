package functions

import (
	"reflect"
	"testing"
)

func TestBannerArt(t *testing.T) {
	tests := []struct {
		name string
		args string
		want map[int][]string
	}{
		{name: "standard", args: "../banners/standard.txt", want: BannerArt("../banners/standard.txt")},
		{name: "shadow", args: "../banners/shadow.txt", want: BannerArt("../banners/shadow.txt")},
		{name: "thinkertoy", args: "../banners/thinkertoy.txt", want: BannerArt("../banners/thinkertoy.txt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BannerArt(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BannerArt() = %v, want %v", got, tt.want)
			}
		})
	}
}
