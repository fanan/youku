package youku

import (
	"testing"
)

var segs_data = []byte(`
{
            "hd2": [{
                "no": "0",
                "size": "28988540",
                "seconds": "201",
                "k": "20035fbded6b551424116846",
                "k2": "1cf5a205bbdf6b9ab"},
            {
                "no": "1",
                "size": "32039755",
                "seconds": "205",
                "k": "21aa697159af2fb224116846",
                "k2": "16be0e83cec9ea719"},
            {
                "no": "2",
                "size": "30474129",
                "seconds": "194",
                "k": "c94cfda55500512a28289df9",
                "k2": "17d9c67b46e0b106f"},
            {
                "no": "3",
                "size": "31209175",
                "seconds": "210",
                "k": "8713a9c00807e4b424116846",
                "k2": "151d2e04d40141f72"},
            {
                "no": "4",
                "size": "25173393",
                "seconds": "183",
                "k": "6e15fe8f028e925b28289df9",
                "k2": "12470c2900093417a"},
            {
                "no": "5",
                "size": "33971981",
                "seconds": "205",
                "k": "e3d9bfc2d3a09eb0261d031f",
                "k2": "160924c87124f7f1c"},
            {
                "no": "6",
                "size": "30767218",
                "seconds": "207",
                "k": "d6ca464550045c8724116846",
                "k2": "129c034d060491e1f"},
            {
                "no": "7",
                "size": "26160554",
                "seconds": "198",
                "k": "9dab1ebfbac3bab528289df9",
                "k2": "1f0fa43a565389cb5"},
            {
                "no": "8",
                "size": "34634426",
                "seconds": "204",
                "k": "6f906f01d7e0b40428289df9",
                "k2": "1a9f850063972551c"},
            {
                "no": "9",
                "size": "28708587",
                "seconds": "198",
                "k": "ef88e96b3028eed6261d031f",
                "k2": "1997bd663cdf939c4"},
            {
                "no": "10",
                "size": "30068897",
                "seconds": "197",
                "k": "0e8069cbcde6d24d24116846",
                "k2": "1de7f9708e6118858"},
            {
                "no": "11",
                "size": "28422030",
                "seconds": "205",
                "k": "19c84ba047815a1124116846",
                "k2": "1c31ee2f1e09347fb"},
            {
                "no": "12",
                "size": "15909990",
                "seconds": "138",
                "k": "6b563b6bb47f55f024116846",
                "k2": "1386e1f1e5480870c"},
            {
                "no": "13",
                "size": "30179696",
                "seconds": "158",
                "k": "1a19c84488e3f4c0261d031f",
                "k2": "18d84cdfe1f89bbcf"}],
            "mp4": [{
                "no": "0",
                "size": "27779128",
                "seconds": "406",
                "k": "5a5677af5476188b261d031f",
                "k2": "1c79e28d959fbbb62"},
            {
                "no": "1",
                "size": "28424739",
                "seconds": "404",
                "k": "1858e8d8001f6212261d031f",
                "k2": "1467b1c068a01e5af"},
            {
                "no": "2",
                "size": "27127311",
                "seconds": "388",
                "k": "2ac239ff59a2e0f5261d031f",
                "k2": "1ab6bcea5f5125294"},
            {
                "no": "3",
                "size": "26708039",
                "seconds": "415",
                "k": "262c62351dddf02124116846",
                "k2": "1b47c5bdeb6bd2c12"},
            {
                "no": "4",
                "size": "29662635",
                "seconds": "415",
                "k": "d404d517e940fdaa28289df9",
                "k2": "114a76b35dbf2d479"},
            {
                "no": "5",
                "size": "22907490",
                "seconds": "331",
                "k": "5e9bb1bf630e32bf28289df9",
                "k2": "1d1b5b98624cf5514"},
            {
                "no": "6",
                "size": "23647045",
                "seconds": "345",
                "k": "88d25bbb65a76610261d031f",
                "k2": "1762b05b949067ee0"}],
            "flv": [{
                "no": "0",
                "size": "14944141",
                "seconds": "419",
                "k": "9200da37a75ea64d261d031f",
                "k2": "1cfcc29e7d74768bd"},
            {
                "no": "1",
                "size": "13861574",
                "seconds": "380",
                "k": "be08526b22965d5428289df9",
                "k2": "1b10091360b9d969c"},
            {
                "no": "2",
                "size": "14669825",
                "seconds": "400",
                "k": "ccb6e7bc134690c924116846",
                "k2": "1c267590e5b86dffa"},
            {
                "no": "3",
                "size": "13278309",
                "seconds": "402",
                "k": "e64ad2078368afb924116846",
                "k2": "16d4033565807128a"},
            {
                "no": "4",
                "size": "14513459",
                "seconds": "382",
                "k": "a9620d9fc6e6334a261d031f",
                "k2": "1418a4d1fc75b99e0"},
            {
                "no": "5",
                "size": "13242353",
                "seconds": "366",
                "k": "84e740f631bc035028289df9",
                "k2": "1c0550841fe7e26f8"},
            {
                "no": "6",
                "size": "12569182",
                "seconds": "355",
                "k": "f8da2eea3e835d8328289df9",
                "k2": "1aea869f6a0894291"}]
        }
`)

func TestSegsParse(t *testing.T) {
	segs := new(Segs)
	err := segs.Parse(&segs_data)
	if err != nil {
		t.Error(err)
	} else {
		logger.Println(len(segs.Hd2), len(segs.Mp4), len(segs.Flv))
	}
}
