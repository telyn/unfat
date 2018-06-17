package dir

import "testing"

func TestPadShortName(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{{
		in:  "hi.wav",
		out: "HI      WAV",
	}, {
		in:  "godwhack.gif",
		out: "GODWHACKGIF",
	}, {
		in:  "whizz",
		out: "WHIZZ      ",
	}}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			res := padShortName(test.in)
			if res != test.out {
				t.Errorf("%q", res)
			}
		})
	}
}

func TestUnpadShortName(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{{
		in:  "HI      WAV",
		out: "hi.wav",
	}, {
		in:  "GODWHACKGIF",
		out: "godwhack.gif",
	}}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			res := padShortName(test.in)
			if res != test.out {
				//			t.Error(res)
			}
		})
	}
}
