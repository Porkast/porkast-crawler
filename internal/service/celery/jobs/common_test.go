package jobs

import (
	"testing"
	"time"
)

func Test_randInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get random int between 5 and 20",
			args: args{
				min: 5,
				max: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := randInt(tt.args.min, tt.args.max)
			if result > tt.args.max || result < tt.args.min {
				t.Fatal("The random int not in 5 and 20")
			}
		})
	}
}

func Test_getRandomStartTime(t *testing.T) {
	var (
		gotTime time.Duration
	)

	gotTime = getRandomStartTime()
	if gotTime > time.Second * 20 || gotTime < time.Second * 5 {
		t.Fatal("Get random time not int 5 and 20 second")
	}

}
