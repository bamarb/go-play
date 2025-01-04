package tests

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type Resp struct {
	Adm  string `json:"adm"`
	Iurl string `json:"iurl"`
}

func TestMarshal_plain(t *testing.T) {
	admHtml := `<div>
  <a href="https://example.com/track?adid=12345&crid=34567&li=786534353267">
  <img src="https://cdn.adster.tech/img/bjp-up.jpg" alt="bjp up ad">
  </a>
</div>`

	resp := Resp{Adm: admHtml, Iurl: "https://junk.com:8080/test/?junk=munk"}
	strWriter := strings.Builder{}
	err := json.NewEncoder(&strWriter).Encode(resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(strWriter.String())
	strWriter.Reset()
	enc := json.NewEncoder(&strWriter)
	enc.SetEscapeHTML(false)
	err = enc.Encode(resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", strWriter.String())
}

const (
	mkup     = "{\"ver\":\"1.2\",\"assets\":[{\"id\":1,\"required\":1,\"data\":{\"type\":12,\"value\":\"Test\"}},{\"id\":2,\"required\":1,\"title\":{\"text\":\"Test case for native ad response\",\"len\":32}},{\"id\":3,\"required\":1,\"img\":{\"type\":1,\"url\":\"http://localhost:8080/cdn/\"}},{\"id\":4,\"required\":1,\"img\":{\"type\":3,\"url\":\"http://localhost:8080/cdn/id-12345/bjp1.jpeg\",\"w\":728,\"h\":90}},{\"id\":6,\"required\":1,\"data\":{\"type\":2,\"value\":\"Test case for native ad response\"}}],\"link\":{\"url\":\"https://adster.tech/ref?ad=adsterx\",\"clicktrackers\":[\"https://logs.adster.tech/v1/pixel_click?actual_adunit=\\u0026ad_type=\\u0026ad_unit=\\u0026adv_name=\\u0026adv_vert=\\u0026advertiser_id=abcd\\u0026buy_type=\\u0026cb=\\u0026click_id=\\u0026client_timestamp=0\\u0026creative_id=000000000000000000000000\\u0026creative_text=\\u0026ctype=\\u0026derived_ecpm=10\\u0026ds_model=\\u0026ext=map%5B%5D\\u0026ga_id=\\u0026line_id=hijk\\u0026network=adster\\u0026order_id=efgh\\u0026partner_id=\\u0026placement_id=adster_Home_Screen_Unified\\u0026product_category=\\u0026product_subtype=\\u0026publisher_id=0\\u0026req_id=01HQJE1ZX3483R04BJ9WV0ND84\\u0026req_type=native\\u0026timestamp=0\\u0026user_agent=\\u0026user_id=hipster\"]},\"eventtrackers\":[{\"event\":1,\"method\":2,\"url\":\"https://logs.adster.tech/v1/pixel_imp?actual_adunit=\\u0026ad_type=\\u0026ad_unit=\\u0026adv_name=\\u0026adv_vert=\\u0026advertiser_id=abcd\\u0026buy_type=\\u0026cb=\\u0026client_timestamp=0\\u0026creative_id=000000000000000000000000\\u0026creative_text=\\u0026ctype=\\u0026derived_ecpm=10\\u0026ds_model=\\u0026ext=map%5B%5D\\u0026ga_id=\\u0026impression_id=\\u0026line_id=hijk\\u0026network=adster\\u0026order_id=efgh\\u0026partner_id=adsterx\\u0026placement_id=adster_Home_Screen_Unified\\u0026product_category=\\u0026product_subtype=\\u0026publisher_id=0\\u0026req_id=01HQJE1ZX3483R04BJ9WV0ND84\\u0026req_type=native\\u0026timestamp=1708946850721\\u0026user_agent=\\u0026user_id=hipster\"}]}"
	mkup_str = "\"{\"ver\":\"1.2\",\"assets\":[{\"id\":1,\"required\":1,\"data\":{\"type\":12,\"value\":\"Test\"}},{\"id\":2,\"required\":1,\"title\":{\"text\":\"Test case for native ad response\",\"len\":32}},{\"id\":3,\"required\":1,\"img\":{\"type\":1,\"url\":\"http://localhost:8080/cdn/\"}},{\"id\":4,\"required\":1,\"img\":{\"type\":3,\"url\":\"http://localhost:8080/cdn/id-12345/bjp1.jpeg\",\"w\":728,\"h\":90}},{\"id\":6,\"required\":1,\"data\":{\"type\":2,\"value\":\"Test case for native ad response\"}}],\"link\":{\"url\":\"https://adster.tech/ref?ad=adsterx\",\"clicktrackers\":[\"https://logs.adster.tech/v1/pixel_click?actual_adunit=\\u0026ad_type=\\u0026ad_unit=\\u0026adv_name=\\u0026adv_vert=\\u0026advertiser_id=abcd\\u0026buy_type=\\u0026cb=\\u0026click_id=\\u0026client_timestamp=0\\u0026creative_id=000000000000000000000000\\u0026creative_text=\\u0026ctype=\\u0026derived_ecpm=10\\u0026ds_model=\\u0026ext=map%5B%5D\\u0026ga_id=\\u0026line_id=hijk\\u0026network=adster\\u0026order_id=efgh\\u0026partner_id=\\u0026placement_id=adster_Home_Screen_Unified\\u0026product_category=\\u0026product_subtype=\\u0026publisher_id=0\\u0026req_id=01HQJE1ZX3483R04BJ9WV0ND84\\u0026req_type=native\\u0026timestamp=0\\u0026user_agent=\\u0026user_id=hipster\"]},\"eventtrackers\":[{\"event\":1,\"method\":2,\"url\":\"https://logs.adster.tech/v1/pixel_imp?actual_adunit=\\u0026ad_type=\\u0026ad_unit=\\u0026adv_name=\\u0026adv_vert=\\u0026advertiser_id=abcd\\u0026buy_type=\\u0026cb=\\u0026client_timestamp=0\\u0026creative_id=000000000000000000000000\\u0026creative_text=\\u0026ctype=\\u0026derived_ecpm=10\\u0026ds_model=\\u0026ext=map%5B%5D\\u0026ga_id=\\u0026impression_id=\\u0026line_id=hijk\\u0026network=adster\\u0026order_id=efgh\\u0026partner_id=adsterx\\u0026placement_id=adster_Home_Screen_Unified\\u0026product_category=\\u0026product_subtype=\\u0026publisher_id=0\\u0026req_id=01HQJE1ZX3483R04BJ9WV0ND84\\u0026req_type=native\\u0026timestamp=1708946850721\\u0026user_agent=\\u0026user_id=hipster\"}]}\""
)

func TestUnQuote(t *testing.T) {
	str := "\"I am A Quoted Str with Escapes\""
	ustr, err := strconv.Unquote(str)
	require.NoError(t, err)
	t.Logf("%s\n", ustr)
}

func TestUnQuote_jsonstr(t *testing.T) {
	ustr, err := strconv.Unquote("`" + mkup + "`")
	require.NoError(t, err)
	t.Logf("%s\n", ustr)
}
