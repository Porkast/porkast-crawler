package worker

import "testing"

func TestChannelUpdateByFeedLink(t *testing.T) {
	type args struct {
		feedLink string
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChannelUpdateByFeedLink(tt.args.feedLink)
		})
	}
}
