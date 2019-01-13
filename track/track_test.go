package track

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	testTrackJSON := `{
            "type": "midi",
            "samples": {
              "kick1": "kick1.wav",
              "kick2": "kick2.wav",
              "marac": "maracas.wav",
              "snare": "snare.wav",
              "hitom": "hightom.wav",
              "lotom": "tom1.wav",
              "clhat": "cl_hihat.wav"
            },
            "pattern": [
              {
                "begin": 0,
                "duration": 1000,
                "sample": "kick1",
                "volume": 1,
                "pan": 1
              }
            ]
		  }`

	var testTrack Track
	err := json.Unmarshal([]byte(testTrackJSON), &testTrack)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, "snare.wav", testTrack.Samples["snare"])
	assert.Equal(t, time.Duration(1000), testTrack.Pattern[0].Duration)
}
