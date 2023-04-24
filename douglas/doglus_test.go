package douglas_test

import (
	"testing"

	"github.com/supernarsi/gotool/douglas"
)

func TestDouglasPeucker(t *testing.T) {
	tests := []struct {
		name  string
		input []douglas.Point
		want  []douglas.Point
	}{
		{name: "test1", input: []douglas.Point{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, want: []douglas.Point{{0, 0}, {3, 3}}},
		{
			name: "test2",
			input: []douglas.Point{
				{29.600307850964363, 106.50623532346269}, {29.600331225981076, 106.50615617160405}, {29.608226, 106.50092}, {29.608253, 106.500941}, {29.60821, 106.500963}, {29.608221, 106.500984},
			},
			want: []douglas.Point{{29.600307850964363, 106.50623532346269}, {29.608221, 106.500984}},
		},
		{
			name: "test3",
			input: []douglas.Point{
				{29.600307850964363, 106.50623532346269},
				{29.600331225981076, 106.50615617160405},
				{29.600376790384363, 106.50606931408845},
				{29.600421269416213, 106.50600073184677},
				{29.60048230817131, 106.50593376572789},
				{29.600783516101032, 106.5049902577372},
				{29.601142624261637, 106.504310853574},
				{29.601173925902685, 106.50425017717171},
				{29.601196135239515, 106.50420601078312},
				{29.601217723538145, 106.50414555317982},
				{29.60125011692011, 106.50407776815615},
				{29.601264900751662, 106.5040115374751},
				{29.60129771406872, 106.50394381230736},
				{29.601295332535148, 106.5038545795922},
				{29.601285457772367, 106.50378562568156},
				{29.60128441620573, 106.50372968328345},
				{29.60126616837989, 106.50366097815721},
				{29.60139287179201, 106.50230189910724},
				{29.601467274493455, 106.5023044955755},
				{29.60153162685326, 106.50231623406722},
				{29.60160257705776, 106.50231933881467},
				{29.601661384592404, 106.50231558343097},
				{29.60171759157669, 106.50230513795164},
				{29.60177236806361, 106.50230799252492},
				{29.601830434700627, 106.50230114636322},
				{29.60187792048587, 106.50229693170931},
				{29.601924133985502, 106.50231657546928},
				{29.6019576667084, 106.50231210079757},
				{29.602011747837608, 106.5023100300942},
				{29.60208274181027, 106.50230855856115},
				{29.60215864325679, 106.50228515305369},
				{29.602221205420562, 106.50228019149996},
				{29.602286461616437, 106.50227745343545},
				{29.60235026669692, 106.50227163452261},
				{29.60242170750834, 106.50229446018645},
				{29.60248864162448, 106.50230250988236},
				{29.60255219723208, 106.50229317150465},
				{29.602593222124604, 106.5022820172825},
				{29.602653939632052, 106.50228837170471},
				{29.602727467068892, 106.50227601295008},
				{29.60277536105935, 106.50225914624306},
				{29.602846777076458, 106.50226364683333},
				{29.602915864785892, 106.50229059002054},
				{29.602974616579914, 106.50228206891872},
				{29.603039548080044, 106.5023140068965},
				{29.603110412069285, 106.50232575572873},
				{29.603173549469037, 106.50233566961958},
				{29.6032256014258, 106.50230657978891},
				{29.603265720266556, 106.50231724012725},
				{29.60332122177841, 106.50230850946804},
				{29.60337527434301, 106.50229598011241},
				{29.603443536426447, 106.5022550467079},
				{29.603506857635885, 106.50223385384982},
				{29.603576859075783, 106.50224304031376},
				{29.60364572117968, 106.50224526757226},
				{29.603694596228266, 106.50224802210829},
				{29.60776295117472, 106.5018521013031},
				{29.607824762823174, 106.50183714964957},
				{29.60783254123113, 106.50176206513875},
				{29.607803067026328, 106.50170443617702},
				{29.60775461045737, 106.50166248912706},
				{29.607734161039208, 106.50160075298415},
				{29.6077412708345, 106.50152181994986},
				{29.60772873254477, 106.50145045311042},
				{29.607977402285677, 106.50085201049276},
				{29.6081231863431, 106.50077067250217},
			},
			want: []douglas.Point{
				{29.600307850964363, 106.50623532346269},
				{29.60129771406872, 106.50394381230736},
				{29.60139287179201, 106.50230189910724},
				{29.603173549469037, 106.50233566961958},
				{29.60776295117472, 106.5018521013031},
				{29.60772873254477, 106.50145045311042},
				{29.6081231863431, 106.50077067250217},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := douglas.DouglasPeucker(tt.input, 0.0001)
			if len(got) != len(tt.want) {
				t.Errorf("DouglasPeucker() = %v, want %v", got, tt.want)
			} else {
				for i, p := range got {
					if p.X != tt.want[i].X || p.Y != tt.want[i].Y {
						t.Errorf("DouglasPeucker() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}
