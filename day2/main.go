package main

var boxes = []string{
	"bvhfawknyoqsudzrpgslecmtkj", "bpufawcnyoqxldzrpgsleimtkj",
	"bvhfawcnyoqxqdzrplsleimtkf", "bvhoagcnyoqxudzrpgsleixtkj",
	"bxvfgwcnyoqxudzrpgsleimtkj", "bvqfawcngoqxudzrpgsleiktkj",
	"bvhfawcnmoqxuyzrpgsleimtkp", "bvheawcnyomxsdzrpgsleimtkj",
	"bcdfawcnyoqxudzrpgsyeimtkj", "bvhpawcnyoqxudzrpgsteimtkz",
	"bxhfawcnyozxudzrpgsleimtoj", "bvhfdwcnyozxudzrposleimtkj",
	"bvpfawcnyotxudzrpgsleimtkq", "bvhfpwccyoqxudzrpgslkimtkj",
	"bvhfawcnyoqxudirpgsreimtsj", "bvhfawcnyoqxudzppgbleemtkj",
	"bvhzawcnyoqxudqrpgslvimtkj", "bvhfawclyoqxudirpgsleimtka",
	"bvhgawfnyoqxudzrpguleimtkj", "bvhfazcnytqxudzrpgslvimtkj",
	"bvhfawcnygxxudzrpgjleimtkj", "bxhfawcnyoqxudzipgsleimtxj",
	"bvhptwcnyoqxudzrpgsleimtmj", "bzhfawcgyooxudzrpgsleimtkj",
	"bvhjlwcnyokxudzrpgsleimtkj", "bvhfawcnyoqxudbrmgslesmtkj",
	"bvhfawcnysixudzwpgsleimtkj", "bvhflwcnymqxxdzrpgsleimtkj",
	"bvifawcnyoyxudzrpgsleimtvj", "bvhfawcnyofxudlrpgsheimtkj",
	"bvhbawcmyoqxudzrpggleimtkj", "bhhxgwcnyoqxudzrpgsleimtkj",
	"bvhfawgnyoqxbdzrpgsleimfkj", "bvhfawcnyoqxudcrngsleimykj",
	"bvhfawcnyofxudzrpgslebgtkj", "bvhfaocnybqxudzapgsleimtkj",
	"bvhxawcnyodxudzrpfsleimtkj", "bchfawcnyoqxudrrtgsleimtkj",
	"bvhfawcqyoqxudzdpgsltimtkj", "bvhfawknyoqxudzrpnsleimtbj",
	"cihfawcnyoqxudirpgsleimtkj", "bvlfawpnyoqxudzrpgslgimtkj",
	"bulfawcnyoqbudzrpgsleimtkj", "bvhfajcnyoqkudzrpgsoeimtkj",
	"bvhrakcnyoqxudzrpgsleimjkj", "bvbftwcnyoqxuvzrpgsleimtkj",
	"bvhfhwcnyoqxudzrpgslelmtbj", "bvhyawcntoqxudzrpgsleimtuj",
	"xvhuawcnyoqxuqzrpgsleimtkj", "pvhfawcnyoqxudzdpglleimtkj",
	"bvhfawsnyoqxudzrpgvlefmtkj", "bvhfawcnyoqxudzrpgepeiwtkj",
	"bvhfawcnyoqxudzrphsleittkr", "dvhfawcnyoqxudzrpkslzimtkj",
	"bvhfawpnyoqxudzrpgmlcimtkj", "bvhsawcnyzqxudzrpgsaeimtkj",
	"bdhfawcnyoqxudzrpasleiwtkj", "bvhfawbnyoqxpdbrpgsleimtkj",
	"mvhfawwnyoqxujzrpgsleimtkj", "bvafawcnyoyxudzrpgsleidtkj",
	"bvhyawcnyoqxudztpgzleimtkj", "besfawcnyoqxudzrpgsleimdkj",
	"bvhfawcnyoqxudrrpgsjeimjkj", "xvhfkwcnyoqxudzcpgsleimtkj",
	"bvhfawcnyeqdudzrpgzleimtkj", "bvhfuwcnybqxudzrpgsleimttj",
	"lvhfawcnyoqhudzdpgsleimtkj", "bvhfawcnyoqxudzrpgslevwtnj",
	"bvhfadcnzoqxxdzrpgsleimtkj", "bvsfawcnyoqxpdzrpgileimtkj",
	"bzhfaycnyoqxudzrpgsxeimtkj", "bwhfdwcnyoqxudzrpgsleimtkz",
	"bvhfawcnyoqxudzrpgsjlimtkm", "bvhfawcnyoqxudsrwgsleimtlj",
	"bbhfalynyoqxudzrpgsleimtkj", "bvhfawcnyeqxudzrpglleimtkr",
	"bvhfawnnboqxurzrpgsleimtkj", "yvhfawcnyoqxudzrpgslzimtpj",
	"bvhfjwcnyoqxqdxrpgsleimtkj", "bthfawcnyoqfudzrpgslhimtkj",
	"bvhfawchuoqxudzqpgsleimtkj", "bvhfawcndoqxudzrugsleimrkj",
	"bvhfawcnnoqxjdzrpgsleidtkj", "bvhpawcnyoqkudzrpgsleimtzj",
	"bvhfaiinyoqxudzopgsleimtkj", "bvhfawcnyxqxuizrigsleimtkj",
	"bvnfawcnyoqxudzqpgsleimbkj", "bvnfawcnyoeyudzrpgsleimtkj",
	"bvhfawcnyoqxudarpgsieimtoj", "bthcawcnyoqxudlrpgsleimtkj",
	"bvhfnwcnyozxudzrpgsleomtkj", "bpwfawcnyoqxudzrpgskeimtkj",
	"bvhfapcnyoqxudnrpgsxeimtkj", "bvhfdwcnyoqxubzrxgsleimtkj",
	"fvhfawcnyoqxjdzrpgsleirtkj", "bvhfawcneoqxudzrvzsleimtkj",
	"bvhaawcnyoqxudzrpgsleimtex", "bvhfawcnyojxudvrpgsleimckj",
	"bvlfawcnyoqxddzrpgsleimtko", "bvhfawclfoqxudzrpgsleiktkj",
	"bvhfawciyobxudzrpgkleimtkj", "bvhfpwcnyoqxudzrpgsqeimtkd",
	"bvhyawcnyyqxudzrkgsleimtkj", "bvhfawcncoqxudzrphsaeimtkj",
	"bvhfawmnyoqxudzrpgifeimtkj", "bvhfawcjyoqxudzjpgszeimtkj",
	"bohfawcnwoqxudzrpgsleimwkj", "bvhfaucnyoqxudzrpgfluimtkj",
	"bvhfawlnyoqgudzrpgwleimtkj", "bmhfawcnyoqxndzrpgsleymtkj",
	"bvhfawcngoqxudzrpzxleimtkj", "bihfawcnyoqxudrrpgsleimokj",
	"lvhfawcnylqxudzrpgsleintkj", "bvhfawcnyoqvugzrqgsleimtkj",
	"bvhfawcnyoqxudzgpgslqimtij", "bvhfawcnyoqludzrpgslnimtcj",
	"hvhfawcnyolxudzrpgsmeimtkj", "nvhfawcdkoqxudzrpgsleimtkj",
	"bvhfawcnyoqxkdzrggsneimtkj", "bvhfawnnyoqxudzrpgqleibtkj",
	"bvhfawyuyoqxudzrhgsleimtkj", "wvhfbwcnyoqxtdzrpgsleimtkj",
	"bvhfawcnyoqxedzzpgoleimtkj", "bvhfawcnioqxunzrpgsleimtnj",
	"bvhfawctyoqxudzrpgsldkmtkj", "bvhfawcnyonxudzrpgsleitpkj",
	"bvefawcnyoqaudzhpgsleimtkj", "bvhfawcnyxqxudzrpgslelmbkj",
	"bvhfamrnyoqxudzrpgsleimgkj", "bvhfaqcnyoqxudzrpgsaeimekj",
	"bvhfawcnyoqcidzrpgsleimvkj", "bvhfawcnnorxudzrpgsmeimtkj",
	"bvroawcnyoqxudzrpgsleiwtkj", "bvhfwwcnyoqxudzrpaslewmtkj",
	"bvsfawcnyoqxudzcpgszeimtkj", "bkhfmwcnyoqjudzrpgsleimtkj",
	"bvtfawcnyoqxudzrcgslecmtkj", "bvhfawcnypzxudzrpgsleimtkv",
	"bvhfawcnyoqzudzrfgtleimtkj", "bvhpawcnyoqxudhrpgsleimtko",
	"tvhfawcnyoqxudzxpfsleimtkj", "bvhfawccyofxudzrpqsleimtkj",
	"bvnfawtnyoqxuzzrpgsleimtkj", "bvhfamcnuwqxudzrpgsleimtkj",
	"bvhfawcfyoqxudjrpgsleimrkj", "bvhpalcnyoqxudzrpgslexmtkj",
	"bvhfawcnjsqxudzlpgsleimtkj", "bvhfafcnioqxydzrpgsleimtkj",
	"bvzfawcnyxqxudzgpgsleimtkj", "bvhzawcnyoqxudzrpgslewctkj",
	"bvhiawcnhoqrudzrpgsleimtkj", "bvhfawcnyoqxuszrggslenmtkj",
	"bvhfowcnyoqxudzrptseeimtkj", "behfawfnyoqxudzrpgsleimlkj",
	"lvhfawcnyoqxudsrpgvleimtkj", "bvhfawnnyaqxudzrpgsqeimtkj",
	"lvhfawcnfoqxvdzrpgsleimtkj", "svhxawcnyoqxudzrpqsleimtkj",
	"bvhfawqnfoqxudzrpgsleimkkj", "bvhfafcnyoqcudzrpgsleimtcj",
	"bvhfyfcntoqxudzrpgsleimtkj", "bvhfpwcnyoqxudzrpgsleimumj",
	"bvhfawccyoqxudzrqgrleimtkj", "bvhfawqnyoqxudzbpgsleimkkj",
	"bvhflwcnyoqxudzrpxsleemtkj", "bvhfawcnyoqxuezrpgslehrtkj",
	"bvhfawceyoqxudzrpgsleimswj", "bvhfawcncohgudzrpgsleimtkj",
	"bahfawcnyoqxgdzrpgsleamtkj", "yvhfawcnyoqxudzrppslrimtkj",
	"fvhfawcmyoqxudzrpgskeimtkj", "bvylawsnyoqxudzrpgsleimtkj",
	"bvhfswcnyyqxedzrpgsleimtkj", "fvrfawcnyoqxudzrpgzleimtkj",
	"bvhfawcnyoqxuvzrpgslermtks", "bvhkawccyoqxudzcpgsleimtkj",
	"bvhfaobnyoqxudzrprsleimtkj", "bvbfawcnyoqxudirpgsleimhkj",
	"bvhfawcnyoqxudzvpgsueimtgj", "bvhxawcnyoqxudzrpgsleimtgi",
	"svhfawcjyoqxuszrpgsleimtkj", "bvnfawcnyoeyudzrpgsldimtkj",
	"bvhfawcnyoqxuhzrpgsleimcki", "bvhfvwcnyoqxudzizgsleimtkj",
	"bvhfapznyohxudzrpgsleimtkj", "bvhfaelnyosxudzrpgsleimtkj",
	"xvhfawcnmoqxuhzrpgsleimtkj", "bjhfawcnyaqxutzrpgsleimtkj",
	"bvhfawcnyohxudzrpgslgnmtkj", "bvhfawcnyoqxudzrppsreimtkx",
	"fvhfapcnyoqyudzrpgsleimtkj", "qvhfafcnyoqxudorpgsleimtkj",
	"bvhfawcnyoqxedzrwgsleimtvj", "bvhfawgnyoqxudzupgqleimtkj",
	"bvhfowctyoqxudzrpgbleimtkj", "bvhwawcnyoqxudzapgslvimtkj",
	"bvhfadcnyoqxudzrugsleimtuj", "bvhfawcnyosxudzlpgsleamtkj",
	"bvhfawcnywqxuqzrpgsloimtkj", "bvhfawcnyoqxumzrpgvlfimtkj",
	"bvhfawcgyoqxbdzrpgsleomtkj", "bvhfahcnyoqwudzrfgsleimtkj",
	"gvbfawcnyrqxudzrpgsleimtkj", "svhfawcnyoqxudlrpgsleimtkx",
	"avhfafcnyoqxuhzrpgsleimtkj", "bvhfawcsyoqxuazrpgsleimtej",
	"bvofawcnyoqxudzrpgsteimtkf", "bvhfajcnyoqxudzqpgszeimtkj",
	"bvhfawcsyoqxudzrmgsleiktkj", "mvhfawcnyoqxudzrpgkluimtkj",
	"bvhfawcnhoqxudzrpgslwhmtkj", "bmhaawsnyoqxudzrpgsleimtkj",
	"bvhfawcnyoqxudzhpgsleimhyj", "bvhfxwcnyoqxsdzypgsleimtkj",
	"bvhpawcyyoqxuczrpgsleimtkj", "bvomawcnyovxudzrpgsleimtkj",
	"bvhfawcnjvqxudzrpgsleimtkt", "nvhfawcnyqqxudzrpgsleittkj",
	"bvhiawcnyzqxudzrpysleimtkj", "bvhdawcnyoqxukzrpgsleimtuj",
	"bvhfawcnyyxxudzrpgslzimtkj", "hvhfawcnyoqxudzupgslemmtkj",
	"byhfawknyoqxudzrpgsleimtkb", "bvhfawcnyoqxudzrpasleihakj",
	"bvafahcnyaqxudzrpgsleimtkj", "bkhfawcnyoqxudzrpgllepmtkj",
	"bghfawcnycqxuzzrpgsleimtkj", "bvhfawcnyoqxudzrbgeleimtkl",
	"bvhfascnyoqgudzrpgsveimtkj", "bvhfawnnyoqxudzrpgsleimtdl",
	"bvhqawcnyoqxudzrpgsleimgrj", "bvhsawdwyoqxudzrpgsleimtkj",
	"bvhfawcnyoqxudzrpgaleipttj", "bvhfawcnrlqxudzrbgsleimtkj",
	"bvhfdwcnyoqxudzqpcsleimtkj", "bvhfawcnyoqxudzopgslexmokj",
	"bvhfawcoyoqxudzrpghlewmtkj", "bvhfozcnykqxudzrpgsleimtkj",
	"bvhfawcnyoqxuvzrpgslrimtkr", "bvhfrncnyoqrudzrpgsleimtkj",
	"bvhfawcnyocxuizrpgslefmtkj", "bvhfawywyoqxudzrpgsleimxkj",
	"bvhfawcnyoqxugzrpgslrimtij", "bvtfawcnyoqxudzcpgsleimtfj",
	"bvhfawcnyoqxuzzspgsleimtkz", "bvhfawcnzoqxvdzrpgslsimtkj",
	"bvhfzwcnyoqxudzrpgslenmhkj", "bvhfkccnyoqxudzrpgzleimtkj",
	"bvhfawcnyoqzudzrpgslhimwkj", "bzhfawvnyooxudzrpgsleimtkj",
}

func main1() {
	var twos = 0
	var threes = 0

	for _, id := range boxes {
		var letters = make(map[rune]int)

		for _, letter := range id {
			letters[letter]++
		}

		var found2, found3 = false, false

		for _, count := range letters {
			if count == 2 && !found2 {
				twos++
				found2 = true
			} else if count == 3 && !found3 {
				threes++
				found3 = true
			}
		}
	}

	println(twos, "*", threes, "=", twos*threes)
}

func main() {
	for _, id1 := range boxes {
		for _, id2 := range boxes {
			var differ, differIndex = 0, 0

			for i, rune1 := range id1 { // Compare the letters
				if byte(rune1) != id2[i] {
					differ++
					differIndex = i
				}
			}
			if differ == 1 {
				println(id1, id2, "-difference", id1[:differIndex]+id1[differIndex+1:])
			}
		}
	}
}
