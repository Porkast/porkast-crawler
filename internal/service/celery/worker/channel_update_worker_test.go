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
        {
        	name: "update feed by feed link",
        	args: args{
        		feedLink: "https://feed.xyzfm.space/byhkljlbep9j",
        	},
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChannelUpdateByFeedLink(tt.args.feedLink)
		})
	}
}
