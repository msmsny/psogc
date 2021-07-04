package status

import (
	"bytes"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	testCases := map[string]struct {
		name   string
		level  string
		output string
	}{
		"fomar5": {
			name:  "fomar",
			level: "5",
			output: heredoc.Doc(`
			name:         fomar
			level:           5
			HP:             47
			TP:            111
			Attack:         30
			Defense:        14
			MindStrength:   70
			Accuracy:       66.0
			Evasion:        53
			`),
		},
		"fomar100": {
			name:  "fomar",
			level: "100",
			output: heredoc.Doc(`
			name:         fomar
			level:         100
			HP:            523
			TP:            883
			Attack:        380
			Defense:       121
			MindStrength:  490
			Accuracy:      111.0
			Evasion:       381
			`),
		},
		"fomar200": {
			name:  "fomar",
			level: "200",
			output: heredoc.Doc(`
			name:         fomar
			level:         200
			HP:           1175
			TP:           1783
			Attack:        750
			Defense:       321
			MindStrength:  990
			Accuracy:      138.0
			Evasion:       551
			`),
		},
		"fomarAll": {
			name: "fomar",
			output: heredoc.Doc(`
			name: fomar, level:   5, HP:   47, TP:  111, Attack:   30, Defense:  14, MindStrength:   70, Accuracy:  66.0, Evasion:  53
			name: fomar, level:  10, HP:   71, TP:  153, Attack:   47, Defense:  19, MindStrength:   93, Accuracy:  68.0, Evasion:  75
			name: fomar, level:  15, HP:   94, TP:  193, Attack:   64, Defense:  23, MindStrength:  115, Accuracy:  71.0, Evasion:  98
			name: fomar, level:  20, HP:  120, TP:  234, Attack:   82, Defense:  27, MindStrength:  137, Accuracy:  73.0, Evasion: 121
			name: fomar, level:  25, HP:  145, TP:  273, Attack:  102, Defense:  29, MindStrength:  158, Accuracy:  75.0, Evasion: 144
			name: fomar, level:  30, HP:  168, TP:  310, Attack:  124, Defense:  33, MindStrength:  178, Accuracy:  77.0, Evasion: 165
			name: fomar, level:  35, HP:  192, TP:  351, Attack:  146, Defense:  40, MindStrength:  200, Accuracy:  79.0, Evasion: 186
			name: fomar, level:  40, HP:  214, TP:  391, Attack:  168, Defense:  47, MindStrength:  222, Accuracy:  82.0, Evasion: 206
			name: fomar, level:  45, HP:  236, TP:  435, Attack:  190, Defense:  53, MindStrength:  246, Accuracy:  84.0, Evasion: 225
			name: fomar, level:  50, HP:  255, TP:  478, Attack:  209, Defense:  59, MindStrength:  270, Accuracy:  87.0, Evasion: 243
			name: fomar, level:  55, HP:  281, TP:  531, Attack:  224, Defense:  66, MindStrength:  300, Accuracy:  89.0, Evasion: 260
			name: fomar, level:  60, HP:  307, TP:  582, Attack:  241, Defense:  71, MindStrength:  329, Accuracy:  92.0, Evasion: 278
			name: fomar, level:  65, HP:  333, TP:  624, Attack:  257, Defense:  76, MindStrength:  352, Accuracy:  94.0, Evasion: 297
			name: fomar, level:  70, HP:  359, TP:  667, Attack:  272, Defense:  81, MindStrength:  376, Accuracy:  96.0, Evasion: 314
			name: fomar, level:  75, HP:  385, TP:  708, Attack:  287, Defense:  89, MindStrength:  398, Accuracy:  99.0, Evasion: 324
			name: fomar, level:  80, HP:  413, TP:  745, Attack:  302, Defense:  94, MindStrength:  418, Accuracy: 101.0, Evasion: 336
			name: fomar, level:  85, HP:  440, TP:  780, Attack:  320, Defense: 101, MindStrength:  436, Accuracy: 104.0, Evasion: 347
			name: fomar, level:  90, HP:  466, TP:  814, Attack:  340, Defense: 107, MindStrength:  454, Accuracy: 106.0, Evasion: 359
			name: fomar, level:  95, HP:  494, TP:  849, Attack:  360, Defense: 114, MindStrength:  472, Accuracy: 109.0, Evasion: 370
			name: fomar, level: 100, HP:  523, TP:  883, Attack:  380, Defense: 121, MindStrength:  490, Accuracy: 111.0, Evasion: 381
			name: fomar, level: 105, HP:  555, TP:  928, Attack:  397, Defense: 131, MindStrength:  515, Accuracy: 113.0, Evasion: 393
			name: fomar, level: 110, HP:  587, TP:  973, Attack:  417, Defense: 141, MindStrength:  540, Accuracy: 114.0, Evasion: 404
			name: fomar, level: 115, HP:  620, TP: 1018, Attack:  437, Defense: 151, MindStrength:  565, Accuracy: 115.0, Evasion: 416
			name: fomar, level: 120, HP:  653, TP: 1063, Attack:  457, Defense: 161, MindStrength:  590, Accuracy: 117.0, Evasion: 428
			name: fomar, level: 125, HP:  685, TP: 1108, Attack:  475, Defense: 171, MindStrength:  615, Accuracy: 118.0, Evasion: 439
			name: fomar, level: 130, HP:  717, TP: 1153, Attack:  493, Defense: 181, MindStrength:  640, Accuracy: 119.0, Evasion: 451
			name: fomar, level: 135, HP:  751, TP: 1198, Attack:  512, Defense: 191, MindStrength:  665, Accuracy: 120.0, Evasion: 462
			name: fomar, level: 140, HP:  784, TP: 1243, Attack:  530, Defense: 201, MindStrength:  690, Accuracy: 121.0, Evasion: 474
			name: fomar, level: 145, HP:  816, TP: 1288, Attack:  548, Defense: 211, MindStrength:  715, Accuracy: 123.0, Evasion: 485
			name: fomar, level: 150, HP:  848, TP: 1333, Attack:  567, Defense: 221, MindStrength:  740, Accuracy: 124.0, Evasion: 496
			name: fomar, level: 155, HP:  881, TP: 1378, Attack:  585, Defense: 231, MindStrength:  765, Accuracy: 125.0, Evasion: 506
			name: fomar, level: 160, HP:  914, TP: 1423, Attack:  603, Defense: 241, MindStrength:  790, Accuracy: 127.0, Evasion: 517
			name: fomar, level: 165, HP:  946, TP: 1468, Attack:  622, Defense: 251, MindStrength:  815, Accuracy: 128.0, Evasion: 527
			name: fomar, level: 170, HP:  978, TP: 1513, Attack:  640, Defense: 261, MindStrength:  840, Accuracy: 130.0, Evasion: 537
			name: fomar, level: 175, HP: 1012, TP: 1558, Attack:  658, Defense: 271, MindStrength:  865, Accuracy: 131.0, Evasion: 547
			name: fomar, level: 180, HP: 1045, TP: 1603, Attack:  676, Defense: 281, MindStrength:  890, Accuracy: 133.0, Evasion: 551
			name: fomar, level: 185, HP: 1077, TP: 1648, Attack:  695, Defense: 291, MindStrength:  915, Accuracy: 134.0, Evasion: 551
			name: fomar, level: 190, HP: 1109, TP: 1693, Attack:  711, Defense: 301, MindStrength:  940, Accuracy: 136.0, Evasion: 551
			name: fomar, level: 195, HP: 1142, TP: 1738, Attack:  731, Defense: 311, MindStrength:  965, Accuracy: 137.0, Evasion: 551
			name: fomar, level: 200, HP: 1175, TP: 1783, Attack:  750, Defense: 321, MindStrength:  990, Accuracy: 138.0, Evasion: 551
			`),
		},
	}

	for testCase, tt := range testCases {
		t.Run(testCase, func(t *testing.T) {
			buf := &bytes.Buffer{}
			output := &output{writer: buf}
			cmd := NewStatusCommand(withOutput(output))
			require.NoError(t, cmd.Flags().Set("name", tt.name))
			if tt.level != "" {
				require.NoError(t, cmd.Flags().Set("level", tt.level))
			}
			require.NoError(t, cmd.Execute())
			assert.Equal(t, tt.output, buf.String())
		})
	}
}
