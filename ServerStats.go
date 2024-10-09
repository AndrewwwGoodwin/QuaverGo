package QuaverGo

import "fmt"

type ServerStats struct {
	APIClient         *Client
	EndpointExtension string
}

func initServerStats(apiClient *Client) *ServerStats {
	return &ServerStats{APIClient: apiClient, EndpointExtension: "/server/"}
}

func (s *ServerStats) GetServerStats() (ServerStatsData, error) {
	var returnData ServerStatsData
	err := s.APIClient.fetchData(fmt.Sprintf("%s%sstats", s.APIClient.baseURL, s.EndpointExtension), &returnData)
	if err != nil {
		return ServerStatsData{}, err
	}
	return returnData, nil
}

type ServerStatsData struct {
	OnlineUsers  int `json:"online_users"`
	TotalMapsets int `json:"total_mapsets"`
	TotalScores  int `json:"total_scores"`
	TotalUsers   int `json:"total_users"`
}

// GetServerStatsByCountries not the hugest fan of how this is implemented. maybe some day ill fix it
func (s *ServerStats) GetServerStatsByCountries() (CountryStats, error) {
	var returnData serverStatsCountries
	err := s.APIClient.fetchData(fmt.Sprintf("%s%sstats/country", s.APIClient.baseURL, s.EndpointExtension), &returnData)
	if err != nil {
		return CountryStats{}, err
	}
	return returnData.Countries, nil
}

type serverStatsCountries struct {
	Countries CountryStats
}

// CountryStats is nasty and probably doesn't contain all the possible return countries here.
// I think I should write a custom unmarshaler that makes this return integers instead of strings as well
// problem for me later in life
type CountryStats struct {
	Unknown string `json:""`
	Ad      string `json:"ad"`
	Ae      string `json:"ae"`
	Af      string `json:"af"`
	Ag      string `json:"ag"`
	Ai      string `json:"ai"`
	Al      string `json:"al"`
	Am      string `json:"am"`
	Ao      string `json:"ao"`
	Ar      string `json:"ar"`
	As      string `json:"as"`
	At      string `json:"at"`
	Au      string `json:"au"`
	Aw      string `json:"aw"`
	Ax      string `json:"ax"`
	Az      string `json:"az"`
	Ba      string `json:"ba"`
	Bb      string `json:"bb"`
	Bd      string `json:"bd"`
	Be      string `json:"be"`
	Bf      string `json:"bf"`
	Bg      string `json:"bg"`
	Bh      string `json:"bh"`
	Bj      string `json:"bj"`
	Bl      string `json:"bl"`
	Bm      string `json:"bm"`
	Bn      string `json:"bn"`
	Bo      string `json:"bo"`
	Bq      string `json:"bq"`
	Br      string `json:"br"`
	Bs      string `json:"bs"`
	Bt      string `json:"bt"`
	Bw      string `json:"bw"`
	By      string `json:"by"`
	Bz      string `json:"bz"`
	Ca      string `json:"ca"`
	Cd      string `json:"cd"`
	Ch      string `json:"ch"`
	Ci      string `json:"ci"`
	Cl      string `json:"cl"`
	Cm      string `json:"cm"`
	Cn      string `json:"cn"`
	Co      string `json:"co"`
	Cr      string `json:"cr"`
	Cu      string `json:"cu"`
	Cv      string `json:"cv"`
	Cw      string `json:"cw"`
	Cx      string `json:"cx"`
	Cy      string `json:"cy"`
	Cz      string `json:"cz"`
	De      string `json:"de"`
	Dk      string `json:"dk"`
	Dm      string `json:"dm"`
	Do      string `json:"do"`
	Dz      string `json:"dz"`
	Ec      string `json:"ec"`
	Ee      string `json:"ee"`
	Eg      string `json:"eg"`
	Es      string `json:"es"`
	Et      string `json:"et"`
	Fi      string `json:"fi"`
	Fj      string `json:"fj"`
	Fm      string `json:"fm"`
	Fo      string `json:"fo"`
	Fr      string `json:"fr"`
	Ga      string `json:"ga"`
	Gb      string `json:"gb"`
	Gd      string `json:"gd"`
	Ge      string `json:"ge"`
	Gf      string `json:"gf"`
	Gg      string `json:"gg"`
	Gh      string `json:"gh"`
	Gi      string `json:"gi"`
	Gl      string `json:"gl"`
	Gn      string `json:"gn"`
	Gp      string `json:"gp"`
	Gq      string `json:"gq"`
	Gr      string `json:"gr"`
	Gt      string `json:"gt"`
	Gu      string `json:"gu"`
	Gy      string `json:"gy"`
	Hk      string `json:"hk"`
	Hm      string `json:"hm"`
	Hn      string `json:"hn"`
	Hr      string `json:"hr"`
	Ht      string `json:"ht"`
	Hu      string `json:"hu"`
	Id      string `json:"id"`
	Ie      string `json:"ie"`
	Il      string `json:"il"`
	Im      string `json:"im"`
	In      string `json:"in"`
	Iq      string `json:"iq"`
	Ir      string `json:"ir"`
	Is      string `json:"is"`
	It      string `json:"it"`
	Je      string `json:"je"`
	Jm      string `json:"jm"`
	Jo      string `json:"jo"`
	Jp      string `json:"jp"`
	Ke      string `json:"ke"`
	Kg      string `json:"kg"`
	Kh      string `json:"kh"`
	Kn      string `json:"kn"`
	Kp      string `json:"kp"`
	Kr      string `json:"kr"`
	Kw      string `json:"kw"`
	Ky      string `json:"ky"`
	Kz      string `json:"kz"`
	La      string `json:"la"`
	Lb      string `json:"lb"`
	Lc      string `json:"lc"`
	Li      string `json:"li"`
	Lk      string `json:"lk"`
	Ls      string `json:"ls"`
	Lt      string `json:"lt"`
	Lu      string `json:"lu"`
	Lv      string `json:"lv"`
	Ly      string `json:"ly"`
	Ma      string `json:"ma"`
	Mc      string `json:"mc"`
	Md      string `json:"md"`
	Me      string `json:"me"`
	Mf      string `json:"mf"`
	Mg      string `json:"mg"`
	Mh      string `json:"mh"`
	Mk      string `json:"mk"`
	Mm      string `json:"mm"`
	Mn      string `json:"mn"`
	Mo      string `json:"mo"`
	Mp      string `json:"mp"`
	Mq      string `json:"mq"`
	Mt      string `json:"mt"`
	Mu      string `json:"mu"`
	Mv      string `json:"mv"`
	Mw      string `json:"mw"`
	Mx      string `json:"mx"`
	My      string `json:"my"`
	Mz      string `json:"mz"`
	Na      string `json:"na"`
	Nc      string `json:"nc"`
	Ne      string `json:"ne"`
	Ng      string `json:"ng"`
	Ni      string `json:"ni"`
	Nl      string `json:"nl"`
	No      string `json:"no"`
	Np      string `json:"np"`
	Nu      string `json:"nu"`
	Nz      string `json:"nz"`
	Om      string `json:"om"`
	Pa      string `json:"pa"`
	Pe      string `json:"pe"`
	Pf      string `json:"pf"`
	Pg      string `json:"pg"`
	Ph      string `json:"ph"`
	Pk      string `json:"pk"`
	Pl      string `json:"pl"`
	Pm      string `json:"pm"`
	Pr      string `json:"pr"`
	Ps      string `json:"ps"`
	Pt      string `json:"pt"`
	Pw      string `json:"pw"`
	Py      string `json:"py"`
	Qa      string `json:"qa"`
	Re      string `json:"re"`
	Ro      string `json:"ro"`
	Rs      string `json:"rs"`
	Ru      string `json:"ru"`
	Sa      string `json:"sa"`
	Sb      string `json:"sb"`
	Sc      string `json:"sc"`
	Sd      string `json:"sd"`
	Se      string `json:"se"`
	Sg      string `json:"sg"`
	Si      string `json:"si"`
	Sk      string `json:"sk"`
	Sl      string `json:"sl"`
	Sm      string `json:"sm"`
	Sn      string `json:"sn"`
	So      string `json:"so"`
	Sr      string `json:"sr"`
	St      string `json:"st"`
	Sv      string `json:"sv"`
	Sx      string `json:"sx"`
	Sy      string `json:"sy"`
	Tc      string `json:"tc"`
	Td      string `json:"td"`
	Tg      string `json:"tg"`
	Th      string `json:"th"`
	Tj      string `json:"tj"`
	Tl      string `json:"tl"`
	Tn      string `json:"tn"`
	To      string `json:"to"`
	Total   string `json:"total"`
	Tr      string `json:"tr"`
	Tt      string `json:"tt"`
	Tw      string `json:"tw"`
	Tz      string `json:"tz"`
	Ua      string `json:"ua"`
	Ug      string `json:"ug"`
	Um      string `json:"um"`
	Us      string `json:"us"`
	Uy      string `json:"uy"`
	Uz      string `json:"uz"`
	Va      string `json:"va"`
	Vc      string `json:"vc"`
	Ve      string `json:"ve"`
	Vg      string `json:"vg"`
	Vi      string `json:"vi"`
	Vn      string `json:"vn"`
	Vu      string `json:"vu"`
	Xk      string `json:"xk"`
	Xx      string `json:"xx"`
	Ye      string `json:"ye"`
	Yt      string `json:"yt"`
	Za      string `json:"za"`
	Zm      string `json:"zm"`
	Zw      string `json:"zw"`
}
