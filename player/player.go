package player

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"github.com/colevscode/composal/track"

	"github.com/go-mix/mix"
	"github.com/go-mix/mix/bind"
	"github.com/go-mix/mix/bind/spec"
)

var (
	sampleHz = float64(48000)
	playSpec = spec.AudioSpec{
		Freq:     sampleHz,
		Format:   spec.AudioS16,
		Channels: 2,
	}
	totalTime  time.Duration
	debugging  = false
	prefixPath string
)

func Setup(prefix string, debug bool) {
	debugging = debug
	prefixPath = prefix

	mix.Debug(debugging)
	bind.UseOutputString("wav")
	mix.Configure(playSpec)
	mix.SetSoundsPath(strings.TrimRight(prefixPath, "/") + "/")
	if debugging {
		fmt.Fprintf(os.Stderr, "Mix, pid:%v, spec:%v\n", os.Getpid(), playSpec)
	}
}

func AddTrack(track *track.Track) {
	for s := 0; s < len(track.Pattern); s++ {
		p := track.Pattern[s]
		if debugging {
			fmt.Fprintf(os.Stderr, "Setfire %s at: %d for: %d vol: %d pan: %d\n",
				track.Samples[p.Sample],
				int(p.Begin*time.Millisecond),
				int(p.Duration*time.Millisecond),
				int(p.Volume), int(p.Pan))
		}

		mix.SetFire(track.Samples[p.Sample],
			p.Begin*time.Millisecond,
			p.Duration*time.Microsecond,
			p.Volume, p.Pan)

		totalTime = time.Duration(math.Max(float64(totalTime),
			float64(p.Begin*time.Millisecond+p.Duration*time.Millisecond)))

	}
}

func Render(start time.Duration, output io.Writer) {
	trailingTime := time.Duration(time.Second)
	fmt.Fprintf(os.Stderr, "Rendering audio. Total time %v\n", totalTime+trailingTime)

	defer Teardown()
	mix.Start()
	mix.OutputStart(totalTime+trailingTime, output)
	mix.OutputContinueTo(totalTime + trailingTime)
	mix.OutputClose()
}

func Teardown() {
	mix.Teardown()
	totalTime = 0
}
