package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilecoinValidateAddress(t *testing.T) {
	validator := &Filecoin{}

	var mainnetCases = map[string]*Result{
		"f00":      {Success, true, FilID, ""},
		"f01":      {Success, true, FilID, ""},
		"f010":     {Success, true, FilID, ""},
		"f0150":    {Success, true, FilID, ""},
		"f0499":    {Success, true, FilID, ""},
		"f01024":   {Success, true, FilID, ""},
		"f01729":   {Success, true, FilID, ""},
		"f0999999": {Success, true, FilID, ""},
		"f15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq":                                              {Success, true, FilSecp256k1, ""},
		"f12fiakbhe2gwd5cnmrenekasyn6v5tnaxaqizq6a":                                              {Success, true, FilSecp256k1, ""},
		"f1wbxhu3ypkuo6eyp6hjx6davuelxaxrvwb2kuwva":                                              {Success, true, FilSecp256k1, ""},
		"f1xtwapqc6nh4si2hcwpr3656iotzmlwumogqbuaa":                                              {Success, true, FilSecp256k1, ""},
		"f1xcbgdhkgkwht3hrrnui3jdopeejsoatkzmoltqy":                                              {Success, true, FilSecp256k1, ""},
		"f17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy":                                              {Success, true, FilSecp256k1, ""},
		"f24vg6ut43yw2h2jqydgbg2xq7x6f4kub3bg6as6i":                                              {Success, true, FilActor, ""},
		"f25nml2cfbljvn4goqtclhifepvfnicv6g7mfmmvq":                                              {Success, true, FilActor, ""},
		"f2nuqrg7vuysaue2pistjjnt3fadsdzvyuatqtfei":                                              {Success, true, FilActor, ""},
		"f24dd4ox4c2vpf5vk5wkadgyyn6qtuvgcpxxon64a":                                              {Success, true, FilActor, ""},
		"f2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y":                                              {Success, true, FilActor, ""},
		"f3vvmn62lofvhjd2ugzca6sof2j2ubwok6cj4xxbfzz4yuxfkgobpihhd2thlanmsh3w2ptld2gqkn2jvlss4a": {Success, true, FilBLS, ""},
		"f3wmuu6crofhqmm3v4enos73okk2l366ck6yc4owxwbdtkmpk42ohkqxfitcpa57pjdcftql4tojda2poeruwa": {Success, true, FilBLS, ""},
		"f3s2q2hzhkpiknjgmf4zq3ejab2rh62qbndueslmsdzervrhapxr7dftie4kpnpdiv2n6tvkr743ndhrsw6d3a": {Success, true, FilBLS, ""},
		"f3q22fijmmlckhl56rn5nkyamkph3mcfu5ed6dheq53c244hfmnq2i7efdma3cj5voxenwiummf2ajlsbxc65a": {Success, true, FilBLS, ""},
		"f3u5zgwa4ael3vuocgc5mfgygo4yuqocrntuuhcklf4xzg5tcaqwbyfabxetwtj4tsam3pbhnwghyhijr5mixa": {Success, true, FilBLS, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[st