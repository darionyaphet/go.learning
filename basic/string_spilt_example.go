package main

import (
	"fmt"
	"strings"
)

func main() {
	tokens := "top_memberfree=1;top_nad=1;top_video_rew=0;pin_efs=orig;top_nid=0;web_logoc=blue;ls_play_continuous_order=2;top_keyword=0;top_recall=7;top_follow_reason=0;top_hqt=9;top_sj=2;top_root=1;top_root_few_topic=0;top_vdio_rew=0;top_an=1;top_gr_model=1;se_gi=1;top_mlt_model=4;top_billread=1;top_hca=0;top_multi_model=0;top_f_r_nb=1;top_retag=0;top_yhgc=1;top_gif=1;top_newfollow=0;top_nmt=0;top_tffrt=0;top_bill=0;top_feedre_itemcf=31;top_video_fix_position=4;top_tr=4;top_alt=0;top_billupdate1=3;top_ebook=0;top_login_card=2;top_nszt=0;top_dtmt=2;top_topic_feedre=21;top_billvideo=0;top_followtop=1;top_ntr=4;se_dt=0;top_free_content=-1;top_card=-1;top_lowup=1;se_tf=1;top_yc=1;top_cc_at=-1;top_is_gr=0;top_root_ac=1;top_user_gift=0;top_sjre=0;top_tag_isolation=1;top_root_mg=1;top_gr_topic_reweight=1;top_feedre_cpt=101;top_universalebook=1;pin_ef=a;top_billupdate=0;top_feedre=1;top_uit=0;top_billpic=0;top_nuc=0;top_tmt=0;top_adpar=0;top_gr_auto_model=0;web_ask_flow=exp_a;top_billab=1;top_feedre_rtt=41;top_30=0;top_feedtopiccard=0;top_nucc=0;tp_sft=a;top_root_web=0;top_tagore=1"
	//for _, token := range strings.Split(tokens, ";") {
	//	pair := strings.Split(token, "=")
	//	if pair[0] == "pin_efs" {
	//		fmt.Printf("pair %s = %s", pair[0], pair[1])
	//	}
	//}

	var m map[string]string
	var ss []string

	ss = strings.Split(tokens, ";")
	m = make(map[string]string)
	for _, pair := range ss {
		z := strings.Split(pair, "=")
		m[z[0]] = z[1]
	}

	fmt.Println(m)
}
