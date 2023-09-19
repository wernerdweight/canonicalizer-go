package canonicalizer

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCanonicalizer_ReturnType(t *testing.T) {
	assertion := assert.New(t)

	assertion.IsType(string(""), New().Canonicalize("Some fancy text with 2 numbers or whatever..."))
}

func TestCanonicalizer_Canonicalize(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(
		"priserne-zlutoucky-kun-upel-dabelske-ody",
		New().Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy"),
	)
	assertion.Equal(
		"a-quick-brown-fox-jumps-over-the-lazy-dog",
		New().Canonicalize("A quick brown fox jumps over the lazy dog"),
	)
	assertion.Contains(
		[]string{
			"falsches-uben-von-xylophonmusik-qualt-jeden-grosseren-zwerg",
			"falsches-uben-von-xylophonmusik-qualt-jeden-groseren-zwerg",
		},
		New().Canonicalize("Falsches Üben von Xylophonmusik quält jeden größeren Zwerg"),
	)
	assertion.Contains(
		[]string{
			"le-coeur-decu-mais-l-ame-plutot-naive-louys-reva-de-crapauter-en-canoe-au-dela-des-iles-pres-du-malstrom-ou-brulent-les-novae",
			"le-coeur-decu-mais-lame-plutot-naive-louys-reva-de-crapauter-en-canoe-au-dela-des-iles-pres-du-malstrom-ou-brulent-les-novae",
		},
		New().Canonicalize("Le cœur déçu mais l'âme plutôt naïve, Louÿs rêva de crapaüter en canoë au delà des îles, près du mälström où brûlent les novæ."),
	)
	assertion.Equal(
		"krdel-stastnych-datlov-uci-pri-usti-vahu-mlkveho-kona-obhryzat-koru-a-zrat-cerstve-maso",
		New().Canonicalize("Kŕdeľ šťastných ďatľov učí pri ústí Váhu mĺkveho koňa obhrýzať kôru a žrať čerstvé mäso."),
	)
	assertion.Equal(
		"quel-fez-sghembo-copre-davanti",
		New().Canonicalize("Quel fez sghembo copre davanti"),
	)
	assertion.Equal(
		"pchnac-w-te-lodz-jeza-lub-osm-skrzyn-fig",
		New().Canonicalize("Pchnąć w tę łódź jeża lub ośm skrzyń fig"),
	)
	assertion.Equal(
		"v-cascach-juga-zil-by-citrus-da-no-falsivyj-ekzempljar",
		New().Canonicalize("В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!"),
	)
	assertion.Equal(
		"el-veloz-murcielago-hindu-comia-feliz-cardillo-y-kiwi-la-ciguena-tocaba-el-saxofon-detras-del-palenque-de-paja",
		New().Canonicalize("El veloz murciélago hindú comía feliz cardillo y kiwi. La cigüeña tocaba el saxofón detrás del palenque de paja"),
	)
}

func TestCanonicalizer_Suffix(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(
		"priserne-zlutoucky-kun-upel-dabelske-ody-ending",
		NewWithMaxLength(80).CanonicalizeWithSuffix(
			"Příšerně žluťoučký kůň úpěl ďábelské ódy",
			"-ending",
		),
	)
	assertion.Equal(
		"a-quick-brown-fox-jumps-over-the-lazy-dog-ending",
		NewWithMaxLength(80).CanonicalizeWithSuffix("A quick brown fox jumps over the lazy dog", "-ending"),
	)
	assertion.Contains(
		[]string{
			"falsches-uben-von-xylophonmusik-qualt-jeden-grosseren-zwerg-ending",
			"falsches-uben-von-xylophonmusik-qualt-jeden-groseren-zwerg-ending",
		},
		NewWithMaxLength(80).CanonicalizeWithSuffix(
			"Falsches Üben von Xylophonmusik quält jeden größeren Zwerg",
			"-ending",
		),
	)
	assertion.Contains(
		[]string{
			"le-coeur-decu-mais-l-ame-plutot-naive-louys-reva-de-crapauter-en-canoe-au-ending",
			"le-coeur-decu-mais-lame-plutot-naive-louys-reva-de-crapauter-en-canoe-au-ending",
		},
		NewWithMaxLength(80).CanonicalizeWithSuffix(
			"Le cœur déçu mais l'âme plutôt naïve, Louÿs rêva de crapaüter en canoë au delà des îles, près du mälström où brûlent les novæ.",
			"-ending",
		),
	)
	assertion.Equal(
		"krdel-stastnych-datlov-uci-pri-usti-vahu-mlkveho-kona-obhryzat-koru-a-zra-ending",
		NewWithMaxLength(80).CanonicalizeWithSuffix(
			"Kŕdeľ šťastných ďatľov učí pri ústí Váhu mĺkveho koňa obhrýzať kôru a žrať čerstvé mäso.",
			"-ending",
		),
	)
	assertion.Equal(
		"quel-fez-sghembo-copre-davanti-ending",
		NewWithMaxLength(80).CanonicalizeWithSuffix("Quel fez sghembo copre davanti", "-ending"),
	)
	assertion.Equal(
		"pchnac-w-te-lodz-jeza-lub-osm-skrzyn-fig-ending",
		NewWithMaxLength(80).CanonicalizeWithSuffix("Pchnąć w tę łódź jeża lub ośm skrzyń fig", "-ending"),
	)
	assertion.Equal(
		"v-cascach-juga-zil-by-citrus-da-no-falsivyj-ekzempljar-ending",
		NewWithMaxLength(80).CanonicalizeWithSuffix(
			"В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!",
			"-ending",
		),
	)
}

func TestCanonicalizer_Separator(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(
		"priserne zlutoucky kun upel dabelske ody",
		New().CanonicalizeWithSeparator(
			"Příšerně žluťoučký kůň úpěl ďábelské ódy",
			" ",
		),
	)
	assertion.Equal(
		"a quick brown fox jumps over the lazy dog",
		New().CanonicalizeWithSeparator("A quick brown fox jumps over the lazy dog", " "),
	)
	assertion.Contains(
		[]string{
			"falsches uben von xylophonmusik qualt jeden grosseren zwerg",
			"falsches uben von xylophonmusik qualt jeden groseren zwerg",
		},
		New().CanonicalizeWithSeparator(
			"Falsches Üben von Xylophonmusik quält jeden größeren Zwerg",
			" ",
		),
	)
	assertion.Contains(
		[]string{
			"le coeur decu mais l ame plutot naive louys reva de crapauter en canoe au dela des iles pres du malstrom ou brulent les novae",
			"le coeur decu mais lame plutot naive louys reva de crapauter en canoe au dela des iles pres du malstrom ou brulent les novae",
		},
		New().CanonicalizeWithSeparator(
			"Le cœur déçu mais l'âme plutôt naïve, Louÿs rêva de crapaüter en canoë au delà des îles, près du mälström où brûlent les novæ.",
			" ",
		),
	)
	assertion.Equal(
		"krdel stastnych datlov uci pri usti vahu mlkveho kona obhryzat koru a zrat cerstve maso",
		New().CanonicalizeWithSeparator(
			"Kŕdeľ šťastných ďatľov učí pri ústí Váhu mĺkveho koňa obhrýzať kôru a žrať čerstvé mäso.",
			" ",
		),
	)
	assertion.Equal(
		"quel fez sghembo copre davanti",
		New().CanonicalizeWithSeparator("Quel fez sghembo copre davanti", " "),
	)
	assertion.Equal(
		"pchnac w te lodz jeza lub osm skrzyn fig",
		New().CanonicalizeWithSeparator("Pchnąć w tę łódź jeża lub ośm skrzyń fig", " "),
	)
	assertion.Equal(
		"v cascach juga zil by citrus da no falsivyj ekzempljar",
		New().CanonicalizeWithSeparator(
			"В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!",
			" ",
		),
	)
}

func TestCanonicalizer_Callbacks(t *testing.T) {
	assertion := assert.New(t)

	canonicalizer := NewWithCallbacks(
		func(str string) string {
			return strings.NewReplacer("a", "x", "e", "y").Replace(str)
		},
		strings.ToUpper,
	)

	assertion.Equal(
		"PRISYRNE-ZLUTOUCKY-KUN-UPEL-DABYLSKE-ODY",
		canonicalizer.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy"),
	)
	assertion.Equal(
		"A-QUICK-BROWN-FOX-JUMPS-OVYR-THY-LXZY-DOG",
		canonicalizer.Canonicalize("A quick brown fox jumps over the lazy dog"),
	)
	assertion.Contains(
		[]string{
			"FXLSCHYS-UBYN-VON-XYLOPHONMUSIK-QUALT-JYDYN-GROSSYRYN-ZWYRG",
			"FXLSCHYS-UBYN-VON-XYLOPHONMUSIK-QUALT-JYDYN-GROSYRYN-ZWYRG",
		},
		canonicalizer.Canonicalize("Falsches Üben von Xylophonmusik quält jeden größeren Zwerg"),
	)
	assertion.Contains(
		[]string{
			"LY-COEUR-DECU-MXIS-L-AMY-PLUTOT-NXIVY-LOUYS-REVX-DY-CRXPXUTYR-YN-CXNOE-XU-DYLA-DYS-ILYS-PRES-DU-MALSTROM-OU-BRULYNT-LYS-NOVAE",
			"LY-COEUR-DECU-MXIS-LAMY-PLUTOT-NXIVY-LOUYS-REVX-DY-CRXPXUTYR-YN-CXNOE-XU-DYLA-DYS-ILYS-PRES-DU-MALSTROM-OU-BRULYNT-LYS-NOVAE",
		},
		canonicalizer.Canonicalize(
			"Le cœur déçu mais l'âme plutôt naïve, Louÿs rêva de crapaüter en canoë au delà des îles, près du mälström où brûlent les novæ.",
		),
	)
	assertion.Equal(
		"KRDYL-STXSTNYCH-DXTLOV-UCI-PRI-USTI-VAHU-MLKVYHO-KONX-OBHRYZXT-KORU-X-ZRXT-CYRSTVE-MASO",
		canonicalizer.Canonicalize(
			"Kŕdeľ šťastných ďatľov učí pri ústí Váhu mĺkveho koňa obhrýzať kôru a žrať čerstvé mäso.",
		),
	)
	assertion.Equal(
		"QUYL-FYZ-SGHYMBO-COPRY-DXVXNTI",
		canonicalizer.Canonicalize("Quel fez sghembo copre davanti"),
	)
	assertion.Equal(
		"PCHNAC-W-TE-LODZ-JYZX-LUB-OSM-SKRZYN-FIG",
		canonicalizer.Canonicalize("Pchnąć w tę łódź jeża lub ośm skrzyń fig"),
	)
	assertion.Equal(
		"V-CASCACH-JUGA-ZIL-BY-CITRUS-DA-NO-FALSIVYJ-EKZEMPLJAR",
		canonicalizer.Canonicalize("В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!"),
	)
}

func TestCanonicalizer_CallbackSetters(t *testing.T) {
	assertion := assert.New(t)

	canonicalizer := New()

	assertion.Equal(
		"priserne-zlutoucky-kun-upel-dabelske-ody",
		canonicalizer.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy"),
	)
	canonicalizer.SetBeforeCallback(func(str string) string {
		return strings.NewReplacer("a", "e", "e", "a").Replace(str)
	})
	assertion.Equal(
		"prisarne-zlutoucky-kun-upel-dabalske-ody",
		canonicalizer.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy"),
	)
	canonicalizer.SetAfterCallback(strings.ToUpper)
	assertion.Equal(
		"PRISARNE-ZLUTOUCKY-KUN-UPEL-DABALSKE-ODY",
		canonicalizer.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy"),
	)
	canonicalizer.SetBeforeCallback(nil)
	assertion.Equal(
		"PRISERNE-ZLUTOUCKY-KUN-UPEL-DABELSKE-ODY",
		canonicalizer.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy"),
	)
	canonicalizer.SetAfterCallback(nil)
	assertion.Equal(
		"priserne-zlutoucky-kun-upel-dabelske-ody",
		canonicalizer.Canonicalize("Příšerně žluťoučký kůň úpěl ďábelské ódy"),
	)
}
