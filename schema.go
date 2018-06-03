// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dbnary

// Prefixes are iri prefixes
var Prefixes = []Prefix{
	{
		Name: "eng",
		URI:  "http://kaiko.getalp.org/dbnary/eng/",
		Key:  true,
	},
	{
		Name: "dbnary",
		URI:  "http://kaiko.getalp.org/dbnary#",
		Suffixes: []string{
			"gloss",
			"writtenForm",
			"Translation",
			"isTranslationOf",
			"targetLanguage",
			"usage",
			"partOfSpeech",
			"describes",
			"Page",
			"senseNumber",
			"Gloss",
			"rank",
			"synonym",
			"hyponym",
			"antonym",
			"hypernym",
			"meronym",
			"targetLanguageCode",
			"holonym",
			"troponym",
		},
	},
	{
		Name: "rdf",
		URI:  "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		Suffixes: []string{
			"type",
			"value",
			"predicate",
			"subject",
			"Statement",
			"object",
		},
	},
	{
		Name: "ontolex",
		URI:  "http://www.w3.org/ns/lemon/ontolex#",
		Suffixes: []string{
			"writtenRep",
			"Form",
			"LexicalEntry",
			"canonicalForm",
			"sense",
			"LexicalSense",
			"phoneticRep",
			"MultiWordExpression",
		},
	},
	{
		Name: "lexvo",
		URI:  "http://lexvo.org/id/iso639-3/",
		Suffixes: []string{
			"eng",
			"rus",
			"fin",
			"deu",
			"cmn",
			"fra",
			"spa",
			"por",
			"ita",
			"nld",
			"swe",
			"jpn",
			"pol",
			"hbs",
			"ell",
			"ces",
			"ron",
			"cat",
			"hun",
			"bul",
			"tur",
			"dan",
			"gle",
			"kor",
			"mri",
			"ara",
			"hye",
			"epo",
			"kat",
			"fas",
			"isl",
			"glg",
			"mkd",
			"heb",
			"lat",
			"vie",
			"ukr",
			"nor",
			"nob",
			"gla",
			"tgl",
			"lav",
			"est",
			"msa",
			"nno",
			"slv",
			"tha",
			"slk",
			"bel",
			"hin",
			"ind",
			"lit",
			"cym",
			"yue",
			"ido",
			"swa",
			"vol",
			"khm",
			"tel",
			"fao",
			"nrf",
			"kur",
			"yid",
			"sqi",
			"aze",
			"oci",
			"kaz",
			"mon",
			"urd",
			"nav",
			"eus",
			"afr",
			"ast",
			"uzb",
			"lao",
			"tgk",
			"ben",
			"kir",
			"glv",
			"mya",
			"nan",
			"wln",
			"mlt",
			"bak",
			"grc",
			"roh",
			"srd",
			"mar",
			"ltz",
			"ina",
			"tuk",
			"tam",
			"bre",
			"vec",
			"ang",
			"rup",
			"que",
			"fur",
			"dsb",
			"uig",
			"arz",
			"pus",
			"fry",
			"scn",
			"mal",
			"tat",
			"haw",
			"sco",
			"guj",
			"san",
			"asm",
			"ltg",
			"kan",
			"crh",
			"sin",
			"hak",
			"bod",
			"cor",
			"zza",
			"kal",
			"amh",
			"hsb",
			"zdj",
			"pan",
			"chr",
			"hat",
			"sme",
			"zul",
			"jav",
			"jbo",
			"nds",
			"hau",
			"ceb",
			"nov",
			"nep",
			"che",
			"oss",
			"tpi",
			"xal",
			"chu",
			"chv",
			"arg",
			"ckb",
			"arc",
			"dlm",
			"ary",
			"sun",
			"ori",
			"fro",
			"oji",
			"sah",
			"myv",
			"nap",
			"non",
			"xcl",
			"mlg",
			"nah",
			"frr",
			"alt",
			"sga",
			"abk",
			"ewe",
			"div",
			"got",
			"lld",
			"zha",
			"ota",
			"cdo",
			"tyv",
			"acw",
			"gsw",
			"csb",
			"rue",
			"lim",
			"yor",
			"ccc",
			"esu",
			"nya",
			"wol",
			"war",
			"ile",
			"som",
			"sot",
			"udm",
			"mwl",
			"cop",
			"lad",
			"snd",
			"iku",
			"dng",
			"ase",
			"ady",
			"grn",
			"twf",
			"wuu",
			"cos",
			"ist",
			"kum",
			"rom",
			"tir",
			"ain",
			"luy",
			"sms",
			"mel",
			"bar",
			"ava",
			"hil",
			"nci",
			"swb",
			"wym",
			"mnc",
			"syc",
			"kmr",
			"smo",
			"krc",
			"khb",
			"ilo",
			"lij",
			"sgs",
			"syl",
			"cjs",
			"bua",
			"lmo",
			"kik",
			"szl",
			"bal",
			"pdt",
			"sdc",
			"nmn",
			"zho",
			"kjh",
			"nog",
			"orv",
			"ton",
			"sdn",
			"lez",
			"xho",
			"fij",
			"pal",
			"lin",
			"mdf",
			"rap",
			"vro",
			"stq",
			"lkt",
			"frm",
			"pms",
			"liv",
			"yai",
			"inh",
			"frp",
			"pap",
			"osx",
			"mzn",
			"ext",
			"vep",
			"kin",
			"kbd",
			"mer",
			"luo",
			"ibo",
			"shn",
			"lus",
			"ovd",
			"koi",
			"smn",
			"egl",
			"enm",
			"tzm",
			"chm",
			"aym",
			"xmf",
			"ppl",
			"pjt",
			"dtp",
			"kam",
			"kaa",
			"krl",
			"ryu",
			"gag",
			"tah",
			"bdr",
			"vot",
			"shi",
			"goh",
			"sna",
			"otk",
			"sva",
			"pli",
			"pdc",
			"pam",
			"sne",
			"tly",
			"mic",
			"yua",
			"kpv",
			"tet",
			"cre",
			"akk",
			"evn",
			"kln",
			"pro",
			"new",
			"bcl",
			"akz",
			"chy",
			"lbe",
			"lug",
			"tsn",
			"sma",
			"tpn",
			"smj",
			"iba",
			"nau",
			"win",
			"sat",
			"lzz",
			"izh",
			"apw",
			"mns",
			"wmw",
			"txb",
			"kac",
			"unm",
			"nay",
			"mww",
			"kqr",
			"mvv",
			"ace",
			"pnb",
			"sux",
			"chk",
			"hop",
			"arn",
			"tli",
			"pih",
			"bam",
			"kab",
			"abq",
			"gmh",
			"dar",
			"ofs",
			"bis",
			"dzo",
			"srn",
			"kon",
			"xto",
			"niv",
			"ipk",
			"moe",
			"ban",
			"dlg",
			"prg",
			"abe",
			"mrw",
			"bjn",
			"vls",
			"rhg",
			"ood",
			"bqi",
			"ebk",
			"cha",
			"rgn",
			"xqa",
			"peo",
			"ruo",
			"mnw",
			"ale",
			"naq",
			"duo",
			"sce",
			"iii",
			"tab",
			"aka",
			"mrc",
			"pox",
			"ave",
			"nso",
			"avd",
			"gil",
			"ssw",
			"kld",
			"gez",
			"swg",
			"pcd",
			"mrj",
			"uga",
			"kok",
			"zun",
			"ksc",
			"sib",
			"bpy",
			"gan",
			"knb",
			"ill",
			"wrh",
			"tiw",
			"sei",
			"raj",
			"bbl",
			"zai",
			"ruq",
			"ems",
			"xpr",
			"kea",
			"cho",
			"hif",
			"aeb",
			"min",
			"rop",
			"apm",
			"yrk",
			"din",
			"dak",
			"udi",
			"mas",
			"dta",
			"amu",
			"jam",
			"vor",
			"orm",
			"aii",
			"xfa",
			"bth",
			"kca",
			"dtr",
			"wbp",
			"bol",
			"apc",
			"tzo",
			"mvi",
			"hrx",
			"tvl",
			"ket",
			"mai",
			"tih",
			"pcc",
			"otw",
			"ifk",
			"nuk",
			"xda",
			"fit",
			"drg",
			"atv",
			"kpy",
			"kim",
			"osa",
			"esh",
			"cic",
			"mul",
			"par",
			"gld",
			"sjd",
			"bho",
			"kas",
			"kri",
			"axm",
			"tig",
			"afb",
			"zgh",
			"pag",
			"gcf",
			"pua",
			"tfn",
			"phn",
			"pau",
			"brh",
			"mah",
			"lnd",
			"ckt",
			"str",
			"aqc",
			"atj",
			"moh",
			"ccp",
			"mad",
			"pot",
			"zea",
			"hit",
			"xnn",
			"pnt",
			"apj",
			"com",
			"nod",
			"aoz",
			"ven",
			"bla",
			"egy",
			"wls",
			"mnk",
			"yur",
			"chc",
			"bug",
			"jct",
			"mga",
			"tso",
			"hsn",
			"cjy",
			"pon",
			"chl",
			"ksw",
			"ful",
			"gml",
			"umb",
			"meu",
			"alq",
			"crk",
			"xpq",
			"olo",
			"yoi",
			"quc",
			"mhn",
			"sog",
			"kla",
			"kky",
			"hts",
			"bsy",
			"und",
			"rar",
			"nio",
			"niu",
			"tzj",
			"cak",
			"gni",
			"dgr",
			"sag",
			"arp",
			"cab",
			"nys",
			"odt",
			"cuk",
			"wni",
			"dum",
			"blt",
			"aae",
			"jra",
			"mfe",
			"paw",
			"ses",
			"ykg",
			"aar",
			"rmf",
			"tts",
			"xul",
			"oge",
			"sea",
			"chn",
			"gul",
			"tcy",
			"hid",
			"bor",
			"sou",
			"itl",
			"umu",
			"sje",
			"mtq",
			"wbl",
			"cim",
			"wam",
			"maz",
			"fkv",
			"rmn",
			"myp",
			"eve",
			"nwy",
			"ivv",
			"ote",
			"wrr",
			"xum",
			"wlc",
			"fud",
			"caa",
			"sac",
			"bem",
			"bsb",
			"osp",
			"zpq",
			"ttt",
			"zko",
			"mrv",
			"rad",
			"amm",
			"cbk",
			"gaa",
			"aln",
			"agr",
			"mop",
			"kjj",
			"kzt",
			"xug",
			"ajp",
			"pim",
			"kdr",
			"sas",
			"shh",
			"kha",
			"gin",
			"tsg",
			"mus",
			"kap",
			"ari",
			"aht",
			"shs",
			"acm",
			"ams",
			"asb",
			"kut",
			"ach",
			"rmy",
			"dty",
			"dbl",
			"hai",
			"yol",
			"sli",
			"pis",
			"mak",
			"xbc",
			"nia",
			"inz",
			"lou",
			"yij",
			"kis",
			"osc",
			"pqm",
			"ljp",
			"agx",
			"mbf",
			"ada",
			"yao",
			"xsr",
			"vma",
			"brg",
			"nus",
			"hac",
			"nrm",
			"pmh",
			"tqw",
			"kjg",
			"cpg",
			"yox",
			"jdt",
			"gue",
			"pkc",
			"fla",
			"wlm",
			"gmy",
			"sjt",
			"tbl",
			"csm",
			"coo",
			"ulk",
			"ctu",
			"mia",
			"kij",
			"sju",
			"yaq",
			"ddo",
			"kgp",
			"tyz",
			"ksk",
			"cja",
			"mcm",
			"cni",
			"mwp",
			"gdo",
			"yuf",
			"ude",
			"lud",
			"oma",
			"rif",
			"glk",
			"rut",
			"kaw",
			"tum",
			"mdh",
			"crj",
			"arq",
			"anq",
			"msn",
			"pgl",
			"wim",
			"ksi",
			"mrq",
			"nuj",
			"aau",
			"tpw",
			"kfr",
			"kva",
			"koy",
			"bkd",
			"tin",
			"tdd",
			"kyh",
			"jrb",
			"aaq",
			"yuy",
			"kmk",
			"doi",
			"sip",
			"miq",
			"akl",
			"mhx",
			"loz",
			"cas",
			"pov",
			"uum",
			"zav",
			"bew",
			"tsi",
			"pao",
			"bvr",
			"awa",
			"kwk",
			"mqm",
			"kzg",
			"okn",
			"pkp",
			"chg",
			"chp",
			"apy",
			"xkl",
			"psu",
			"crs",
			"pgn",
			"tcs",
			"clm",
			"tsd",
			"ett",
			"nst",
			"cku",
			"tta",
			"dwu",
			"tzh",
			"nhx",
			"xdc",
			"fon",
			"crl",
			"hus",
			"emb",
			"ssf",
			"poe",
			"apk",
			"lwg",
			"mlm",
			"tiy",
			"lzh",
			"qua",
			"ckv",
			"lui",
			"xdk",
			"krj",
			"shp",
			"fax",
			"nnt",
			"kmv",
			"car",
			"bej",
			"trv",
			"pmt",
			"cay",
			"ami",
			"cia",
			"rmq",
			"ibg",
			"ksd",
			"bdk",
			"aer",
			"pad",
			"gub",
			"pcm",
			"btx",
			"gug",
			"aaa",
			"xpm",
			"her",
			"taa",
			"lsd",
			"gut",
			"msk",
			"xnt",
			"bbc",
			"nbl",
			"lki",
			"wlo",
			"yuk",
			"gwi",
			"eka",
			"hmc",
			"mdr",
			"pah",
			"mov",
			"klb",
			"tae",
			"tus",
			"huz",
			"nij",
			"lbk",
			"yap",
			"ycn",
			"kyi",
			"sdo",
			"kbp",
			"rth",
			"skb",
			"ndd",
			"vkp",
			"xls",
			"srm",
			"bgs",
			"sky",
			"kmc",
			"bdy",
			"yag",
			"dgc",
			"abx",
			"btd",
			"mpm",
			"acy",
			"crx",
			"rtm",
			"bnn",
			"skp",
			"dbp",
			"ayl",
			"ryn",
			"bkm",
			"nfr",
			"dif",
			"wrz",
			"tbc",
			"urb",
			"tsu",
			"fia",
			"abl",
			"lep",
			"nrn",
			"xsv",
			"oko",
			"wau",
			"aho",
			"mhy",
			"mos",
			"chz",
			"cac",
			"cro",
			"ndo",
			"sjw",
			"lmw",
			"qui",
			"css",
			"crg",
			"vmb",
			"cle",
			"uby",
			"haa",
			"coc",
			"crd",
			"bno",
			"ing",
			"blc",
			"cmg",
			"idb",
			"lod",
			"tay",
			"aby",
			"ctg",
			"teo",
			"cpa",
			"ple",
			"bde",
			"lrc",
			"hoi",
			"mwf",
			"pwn",
			"cad",
			"dbj",
			"tee",
			"fan",
			"ute",
			"bxk",
			"rob",
			"cup",
			"pac",
			"apl",
			"mch",
			"zoc",
			"sgd",
			"nez",
			"guc",
			"smr",
			"djk",
			"mnv",
			"ess",
			"chf",
			"ngh",
			"kai",
			"hur",
			"kkz",
			"bsk",
			"lmc",
			"owl",
			"hdn",
			"kzj",
			"huv",
			"dny",
			"xrn",
			"tue",
			"ium",
			"enh",
			"xta",
			"kuz",
			"nde",
			"mpt",
			"ibb",
			"maq",
			"jao",
			"kqe",
			"hix",
			"tar",
			"xlu",
			"top",
			"smk",
			"mcb",
			"mni",
			"akr",
			"bll",
			"lua",
			"efi",
			"irk",
			"ngi",
			"wyi",
			"crn",
			"hnj",
			"cux",
			"tos",
			"kua",
			"azg",
			"acu",
			"tao",
			"mpj",
			"wyb",
			"dje",
			"hnn",
			"bya",
			"kge",
			"xeb",
			"urh",
			"kjq",
			"rol",
			"agm",
			"aca",
			"hub",
			"cce",
			"hvn",
			"nhn",
			"trs",
			"ttq",
			"cui",
			"nsq",
			"nsk",
			"ais",
			"gbm",
			"frk",
			"cts",
			"mse",
			"mph",
			"txg",
			"woe",
			"xsa",
			"suk",
			"sia",
			"bcr",
			"xpg",
			"har",
			"slr",
			"end",
			"gej",
			"sxr",
			"ian",
			"aly",
			"stp",
			"mez",
			"zku",
			"xsm",
			"bkl",
			"elx",
			"jup",
			"mir",
			"sxn",
			"ncz",
			"bcu",
			"aqp",
			"hax",
			"lut",
			"kvh",
			"zkg",
			"gdd",
			"mlv",
			"blm",
			"bdg",
			"ane",
			"one",
			"ims",
			"bjz",
			"mrn",
			"wya",
			"itv",
			"nbh",
			"oka",
			"xbr",
			"apn",
			"amn",
			"tnq",
			"ocu",
			"pak",
			"khw",
			"dig",
			"pok",
			"obi",
			"pyu",
			"xsy",
			"msb",
			"zpv",
			"kxd",
			"bpz",
			"kls",
			"mnb",
			"fab",
			"nlg",
			"dis",
			"ktn",
			"tbd",
			"kms",
			"cps",
			"scl",
			"mhl",
			"ibd",
			"csi",
			"ctm",
			"rys",
			"orh",
			"see",
			"yum",
			"jeb",
			"sob",
			"zaa",
			"mek",
			"sks",
			"mhq",
			"gun",
			"oui",
			"nbj",
			"oaa",
			"ojs",
			"cyo",
			"mjy",
			"fpe",
			"xnb",
			"kkh",
			"xgf",
			"bfa",
			"myx",
			"sel",
			"isi",
			"tkr",
			"gcr",
			"xav",
			"xcb",
			"ayn",
			"jiv",
			"bhp",
			"iwm",
			"aia",
			"xww",
			"urm",
			"cah",
			"ybh",
			"mhj",
			"cgc",
			"kpg",
			"mbl",
			"cir",
			"tbo",
			"xpu",
			"xce",
			"jmd",
			"aec",
			"apx",
			"sge",
			"bzj",
			"ksx",
			"njm",
			"lre",
			"ats",
			"zag",
			"sno",
			"zau",
			"oym",
			"plw",
			"mxe",
			"pni",
			"isd",
			"maa",
			"tpx",
			"gon",
			"zap",
			"wnw",
			"poi",
			"xsc",
			"ltc",
			"xtc",
			"bwi",
			"zoo",
			"cdm",
			"tkl",
			"tlv",
			"sae",
			"skh",
			"bgb",
			"ojp",
			"btw",
			"kys",
			"sln",
			"sjr",
			"akv",
			"tuo",
			"cbs",
			"aat",
			"arw",
			"haz",
			"uun",
			"xld",
			"tag",
			"mlu",
			"dus",
			"jac",
			"abd",
			"iyx",
			"for",
			"nwa",
			"ore",
			"obm",
			"tvo",
			"ali",
			"ahr",
			"zor",
			"alj",
			"bpr",
			"svm",
			"veo",
			"arl",
			"kdk",
			"suz",
			"pir",
			"khv",
			"bzq",
			"aud",
			"yux",
			"mng",
			"bxa",
			"alw",
			"ybe",
			"myy",
			"bgt",
			"tjg",
			"tmu",
			"bbr",
			"nym",
			"kee",
			"alk",
			"xsl",
			"kmb",
			"txh",
			"ata",
			"nha",
			"hoc",
			"urz",
			"kxa",
			"pnw",
			"mat",
			"ckx",
			"dua",
			"mzq",
			"brt",
			"pmf",
			"byn",
			"ilb",
			"dav",
			"kwd",
			"kdu",
			"cri",
			"dtd",
			"ssb",
			"cod",
			"mwr",
			"kbc",
			"agt",
			"kxo",
			"cpn",
			"aoa",
			"tmh",
			"ega",
			"dsn",
			"tpf",
			"bsn",
			"ski",
			"gor",
			"nxg",
			"bil",
			"iso",
			"bec",
			"obt",
			"anw",
			"yui",
			"hei",
			"juc",
			"tte",
			"alp",
			"eto",
			"alr",
			"qum",
			"slc",
			"mzs",
			"crw",
			"adw",
			"kiy",
			"sjm",
			"aui",
			"rej",
			"pri",
			"wao",
			"shy",
			"lil",
			"mqj",
			"kyl",
			"ktw",
			"bkn",
			"xaw",
			"bin",
			"api",
			"jhi",
			"peh",
			"mto",
			"uvl",
			"mfa",
			"gvf",
			"tcb",
			"pej",
			"jax",
			"sjo",
			"mbb",
			"kcg",
			"aak",
			"ljl",
			"mgr",
			"bdq",
			"kau",
			"tpy",
			"sgc",
			"tbi",
			"ani",
			"pre",
			"hvc",
			"lun",
			"gfk",
			"yka",
			"njo",
			"wiv",
			"gbu",
			"kuu",
			"cub",
			"tnc",
			"mid",
			"abm",
			"cbv",
			"are",
			"hss",
			"amg",
			"yug",
			"asu",
			"ewo",
			"ofu",
			"mog",
			"aum",
			"btu",
			"yuc",
			"bgc",
			"skd",
			"tri",
			"snn",
			"asn",
			"ppk",
			"apu",
			"mki",
			"hup",
			"lif",
			"xkn",
			"sub",
			"inm",
			"ayz",
			"guz",
			"ado",
			"teg",
			"bfs",
			"cks",
			"sri",
			"nee",
			"abn",
			"cjm",
			"mna",
			"kne",
			"vmy",
			"cin",
			"pos",
			"atw",
			"spp",
			"guh",
			"tew",
			"tft",
			"lmy",
			"hmv",
			"hmo",
			"sey",
			"mwv",
			"pmw",
			"txn",
			"wes",
			"squ",
			"kic",
			"rel",
			"too",
			"tnt",
			"mva",
			"xlc",
			"ncg",
			"tna",
			"con",
			"cul",
			"mjk",
			"max",
			"gvc",
			"xco",
			"piv",
			"dru",
			"kqn",
			"trf",
			"zca",
			"pga",
			"klq",
			"nyn",
			"coe",
			"blk",
			"sdh",
			"mfp",
			"ark",
			"isk",
			"bhj",
			"kxb",
			"otd",
			"ntw",
			"bfq",
			"wic",
			"tun",
			"mro",
			"tog",
			"yku",
			"smw",
			"sgw",
			"gqn",
			"asa",
			"txm",
			"stg",
			"mjm",
			"bky",
			"kry",
			"rmm",
			"xke",
			"zpf",
			"lgk",
			"bnd",
			"rak",
			"kpj",
			"alc",
			"nmw",
			"kqf",
			"nmc",
			"gnn",
			"ges",
			"mkf",
			"kem",
			"mpn",
			"vka",
			"abb",
			"xqt",
			"bap",
			"caf",
			"mqy",
			"kgk",
			"slu",
			"teu",
			"apt",
			"myh",
			"yak",
			"cng",
			"twu",
			"gae",
			"moz",
			"ret",
			"zae",
			"hux",
			"puw",
			"och",
			"crz",
			"zpl",
			"nho",
			"bhg",
			"svs",
			"xeu",
			"mbp",
			"hot",
			"lex",
			"vbb",
			"bxf",
			"aqm",
			"kfy",
			"oin",
			"sve",
			"doz",
			"imr",
			"ygr",
			"tsx",
			"wiy",
			"bkq",
			"ser",
			"sly",
			"psy",
			"til",
			"cso",
			"erg",
			"spe",
			"kip",
			"kwi",
			"kzx",
			"byd",
			"gur",
			"eya",
			"bcf",
			"tke",
			"srh",
			"bkz",
			"bnf",
			"bjc",
			"kgm",
			"pek",
			"met",
			"lvk",
			"goi",
			"nsz",
			"lis",
			"jbk",
			"tvm",
			"cwd",
			"der",
			"kxz",
			"lmb",
			"aap",
			"agw",
			"dyo",
			"jaa",
			"asx",
			"kbw",
			"muz",
			"trw",
			"thp",
			"dij",
			"pij",
			"kde",
			"vic",
			"bsu",
			"gob",
			"dby",
			"mxi",
			"yrl",
			"mfy",
			"aif",
			"xml",
			"bca",
			"bna",
			"mgq",
			"npl",
			"idi",
			"wms",
			"nyo",
			"kzf",
			"hmt",
			"kji",
			"jay",
			"dmr",
			"mhu",
			"hms",
			"mil",
			"rug",
			"arx",
			"ohu",
			"rjs",
			"hrc",
			"elk",
			"emw",
			"aig",
			"xkg",
			"bji",
			"osi",
			"btn",
			"umc",
			"kmj",
			"bhi",
			"slp",
			"hto",
			"wet",
			"gal",
			"mhi",
			"ifb",
			"hna",
			"kwf",
			"kii",
			"tro",
			"kfa",
			"ram",
			"brx",
			"lbw",
			"lra",
			"szv",
			"wed",
			"mup",
			"jka",
			"esi",
			"kpc",
			"mnd",
			"kju",
			"oty",
			"mkj",
			"kdj",
			"pbb",
			"agg",
			"sjk",
			"ghs",
			"dhv",
			"rir",
			"aoc",
			"ixl",
			"tea",
			"ppi",
			"gvj",
			"wru",
			"bpq",
			"mxz",
			"wmt",
			"stf",
			"tlc",
			"aey",
			"vme",
			"cog",
			"cyb",
			"kni",
			"bci",
			"bcd",
			"wat",
			"auw",
			"pnr",
			"khc",
			"xla",
			"tpa",
			"bue",
			"cap",
			"med",
			"msw",
			"lhu",
			"agd",
			"kts",
			"jmx",
			"auq",
			"xan",
			"iyo",
			"tbj",
			"hao",
			"nwr",
			"pnz",
			"cbt",
			"ddi",
			"mxb",
			"mne",
			"bqb",
			"wod",
			"xrw",
			"txs",
			"kpr",
			"wow",
			"res",
			"ojv",
			"reb",
			"yub",
			"dji",
			"khz",
			"guo",
			"ppu",
			"cbg",
			"oyd",
			"mrx",
			"nup",
			"boq",
			"gnd",
			"kxv",
			"tiv",
			"bzk",
			"tow",
			"hdy",
			"pbn",
			"kyt",
			"dva",
			"arh",
			"xmm",
			"bcn",
			"srw",
			"wne",
			"abp",
			"mrt",
			"dia",
			"irn",
			"vko",
			"mtp",
			"uln",
			"dmu",
			"cbi",
			"nll",
			"moc",
			"sry",
			"kti",
			"mvd",
			"tjs",
			"sda",
			"kek",
			"kul",
			"mod",
			"wal",
			"urt",
			"lag",
			"lti",
			"bmh",
			"bja",
			"mfb",
			"ttu",
			"nam",
			"hch",
			"hnd",
			"cji",
			"hhy",
			"www",
			"erk",
			"lsr",
			"sre",
			"apb",
			"aso",
			"bij",
			"abt",
			"kag",
			"spx",
			"not",
			"cta",
			"tdv",
			"kzp",
			"vah",
			"ktz",
			"zgr",
			"bgz",
			"ikw",
			"omn",
			"aah",
			"slm",
			"pab",
			"mgv",
			"tjm",
			"mbr",
			"heh",
			"sbh",
			"omx",
			"cuc",
			"kqw",
			"kio",
			"pil",
			"plv",
			"wrs",
			"shv",
			"tlr",
			"xbm",
			"crm",
			"kyq",
			"yii",
			"lue",
			"des",
			"gby",
			"scb",
			"blw",
			"ong",
			"kma",
			"piz",
			"bea",
			"kod",
			"xvs",
			"huu",
			"cav",
			"miz",
			"rof",
			"bcj",
			"iai",
			"tmq",
			"tcc",
			"kvr",
			"bdo",
			"zen",
			"dto",
			"unz",
			"mox",
			"hov",
			"pio",
			"nak",
			"sax",
			"bts",
			"bpg",
			"bwd",
			"kay",
			"pud",
			"bch",
			"aes",
			"sny",
			"oca",
			"niy",
			"num",
			"meo",
			"aai",
			"tav",
			"mda",
			"bny",
			"ylg",
			"xyy",
			"mnt",
			"khs",
			"ule",
			"cax",
			"skr",
			"att",
			"kmn",
			"jic",
			"ttv",
			"tuq",
			"sso",
			"aty",
			"quv",
			"dai",
			"gbr",
			"ibl",
			"enf",
			"pui",
			"cmi",
			"drn",
			"duk",
			"lib",
			"ndc",
			"aas",
			"mvn",
			"nab",
			"tuf",
			"grw",
			"csz",
			"ebo",
			"ues",
			"lcm",
			"blf",
			"chw",
			"cwg",
			"rah",
			"djd",
			"btz",
			"ojw",
			"tdt",
			"bbd",
			"bim",
			"kka",
			"myw",
			"rma",
			"poo",
			"puj",
			"kos",
			"srs",
			"big",
			"pay",
			"akm",
			"bob",
			"sgh",
			"lgt",
			"otn",
			"uiv",
			"xwa",
			"tkn",
			"bhl",
			"szp",
			"lah",
			"mmn",
			"kzi",
			"rac",
			"koa",
			"bmc",
			"los",
			"fuy",
			"khq",
			"uur",
			"pln",
			"btm",
			"bjk",
			"djm",
			"yml",
			"ywr",
			"xdm",
			"sja",
			"kak",
			"kbk",
			"dev",
			"emi",
			"buk",
			"bwf",
			"eme",
			"myu",
			"avt",
			"cao",
			"ign",
			"auu",
			"ykm",
			"men",
			"hya",
			"pka",
			"hea",
			"nux",
			"onw",
			"bkh",
			"kbt",
			"amk",
			"drl",
			"kjl",
			"bgv",
			"kdd",
			"bjt",
			"rai",
			"fun",
			"loj",
			"shk",
			"mmx",
			"svb",
			"nol",
			"alu",
			"trc",
			"kpx",
			"woc",
			"bhv",
			"khl",
			"haj",
			"ppn",
			"akg",
			"mty",
			"ncn",
			"fwa",
			"amq",
			"nxr",
			"ksb",
			"vkl",
			"bzd",
			"gri",
			"swp",
			"dgg",
			"nss",
			"cbc",
			"zkt",
			"dcc",
			"bjl",
			"aun",
			"pit",
			"iow",
			"fos",
			"wnk",
			"mcr",
			"pse",
			"cto",
			"rea",
			"kym",
			"ajg",
			"nem",
			"lmk",
			"htu",
			"mnp",
			"itx",
			"rro",
			"mot",
			"ttk",
			"kue",
			"lay",
			"dmw",
			"bzg",
			"ijc",
			"yer",
			"kvo",
			"ade",
			"huj",
			"uok",
			"ptu",
			"tdc",
			"lbx",
			"mtt",
			"upv",
			"val",
			"bei",
			"tfr",
			"rgr",
			"kxn",
			"bty",
			"cld",
			"ral",
			"buy",
			"hne",
			"unr",
			"asi",
			"kei",
			"pip",
			"txe",
			"pib",
			"mkv",
			"wca",
			"guu",
			"yup",
			"due",
			"dbn",
			"bkr",
			"ksf",
			"czn",
			"tbf",
			"beg",
			"mta",
			"kwh",
			"leh",
			"nmb",
			"mlp",
			"bco",
			"bdh",
			"xvn",
			"roe",
			"blz",
			"wab",
			"bzb",
			"man",
			"nfu",
			"xsi",
			"kmo",
			"acn",
			"prk",
			"vmg",
			"cam",
			"ese",
			"aua",
			"faf",
			"ass",
			"pww",
			"ppo",
			"bvy",
			"add",
			"snc",
			"alo",
			"eth",
			"roo",
			"bxe",
			"eno",
			"kog",
			"kbh",
			"lon",
			"ssg",
			"ddd",
			"kzh",
			"ape",
			"bln",
			"bdn",
			"boa",
			"bxg",
			"jig",
			"fut",
			"sbv",
			"tlk",
			"qyp",
			"brc",
			"smq",
			"gid",
			"mqn",
			"txa",
			"mxd",
			"pna",
			"tzn",
			"lgr",
			"zmi",
			"uta",
			"wkd",
			"aad",
			"mrh",
			"dim",
			"hgm",
			"sid",
			"tix",
			"gvl",
			"umo",
			"gay",
			"mpe",
			"sgt",
			"sam",
			"lam",
			"akt",
			"leu",
			"gum",
			"tve",
			"wig",
			"ddw",
			"ank",
			"acv",
			"agv",
			"wbv",
			"mok",
			"blb",
			"pia",
			"bhs",
			"boi",
			"gyd",
			"wad",
			"agf",
			"adz",
			"ktx",
			"dag",
			"scg",
			"amx",
			"cbd",
			"tip",
			"mui",
			"udj",
			"bum",
			"vmf",
			"tkw",
			"mxr",
			"bmr",
			"bdd",
			"mig",
			"utp",
			"bhq",
			"plh",
			"lml",
			"bnv",
			"bmk",
			"bld",
			"wuv",
			"mwe",
			"abr",
			"gcl",
			"kzu",
			"icr",
			"itb",
			"kkk",
			"pdo",
			"bfc",
			"mxm",
			"vnk",
			"rim",
			"bnz",
			"mwc",
			"sto",
			"rkb",
			"ndp",
			"bhw",
			"tlu",
			"tox",
			"itz",
			"miw",
			"emp",
			"nhu",
			"lub",
			"dil",
			"agn",
			"noe",
			"paq",
			"gdq",
			"bey",
			"wew",
			"xel",
			"bwu",
			"ars",
			"tct",
			"mpx",
			"duf",
			"noa",
			"frq",
			"mei",
			"sbf",
			"coj",
			"wah",
			"xpo",
			"gve",
			"rbb",
			"mhs",
			"dww",
			"ssq",
			"trn",
			"laj",
			"bhm",
			"skf",
			"mbc",
			"klx",
			"bnq",
			"plo",
			"tsj",
			"lew",
			"xib",
			"tbe",
			"knv",
			"bkt",
			"ilu",
			"nnh",
			"smp",
			"hoa",
			"kyz",
			"vai",
			"kpt",
			"gwd",
			"kjb",
			"bbn",
			"pwg",
			"ura",
			"awb",
			"aac",
			"bwa",
			"kvc",
			"brz",
			"ikt",
			"kck",
			"atu",
			"ssd",
			"plc",
			"asl",
			"ame",
			"bdp",
			"abs",
			"kpm",
			"mzp",
			"emn",
			"onn",
			"mvb",
			"kwe",
			"jod",
			"zmr",
			"nkr",
			"anm",
			"nsn",
			"xng",
			"aiw",
			"lgl",
			"ags",
			"mjg",
			"ttm",
			"avi",
			"lbb",
			"hmr",
			"mmq",
			"eky",
			"swu",
			"yuz",
			"wrg",
			"pma",
			"tbk",
			"fad",
			"djo",
			"sus",
			"ayd",
			"cof",
			"wro",
			"hih",
			"xem",
			"ubr",
			"amc",
			"skz",
			"prq",
			"mue",
			"xed",
			"ers",
			"ofo",
			"tgx",
			"pbh",
			"nil",
			"xmz",
			"bxh",
			"mjb",
			"jbt",
			"dob",
			"rin",
			"xct",
			"and",
			"bkj",
			"ulu",
			"ziw",
			"bzz",
			"ykn",
			"inb",
			"ved",
			"kje",
			"wlk",
			"ray",
			"trx",
			"jml",
			"aib",
			"tpj",
			"ctd",
			"kcn",
			"lev",
			"omc",
			"umi",
			"ate",
			"zne",
			"bra",
			"gva",
			"bag",
			"tan",
			"czk",
			"aot",
			"yey",
			"mbn",
			"tdi",
			"qun",
			"lmr",
			"kqb",
			"ysg",
			"nbp",
			"yrb",
			"ayh",
		},
	},
	{
		Name: "lexinfo",
		URI:  "http://www.lexinfo.net/ontology/2.0/lexinfo#",
		Suffixes: []string{
			"partOfSpeech",
			"noun",
			"verb",
			"adjective",
			"properNoun",
			"adverb",
			"interjection",
			"Interjection",
			"phraseologicalUnit",
			"Prefix",
			"prefix",
			"proverb",
			"pronunciation",
			"suffix",
			"Suffix",
			"pronoun",
			"Pronoun",
			"preposition",
			"Preposition",
			"numeral",
			"Numeral",
			"conjunction",
			"Conjunction",
			"determiner",
			"Determiner",
			"symbol",
			"Symbol",
			"Number",
			"infix",
			"Infix",
			"Particle",
			"particle",
			"Article",
			"article",
			"Affix",
			"affix",
			"postposition",
			"Postposition",
			"idiom",
		},
	},
	{
		Name: "olia",
		URI:  "http://purl.org/olia/olia.owl#",
		Suffixes: []string{
			"hasNumber",
			"Singular",
			"hasCountability",
			"Countable",
			"Uncountable",
			"Plural",
			"hasDegree",
			"Comparative",
		},
	},
	{
		Name: "dcterms",
		URI:  "http://purl.org/dc/terms/",
		Suffixes: []string{
			"language",
		},
	},
	{
		Name: "lime",
		URI:  "http://www.w3.org/ns/lemon/lime#",
		Suffixes: []string{
			"language",
		},
	},
	{
		Name: "skos",
		URI:  "http://www.w3.org/2004/02/skos/core#",
		Suffixes: []string{
			"definition",
		},
	},
	{
		Name: "vartrans",
		URI:  "http://www.w3.org/ns/lemon/vartrans#",
		Suffixes: []string{
			"lexicalRel",
		},
	},
	{
		Name: "decomp",
		URI:  "http://www.w3.org/ns/lemon/decomp#",
	},
	{
		Name: "dbetym",
		URI:  "http://etytree-virtuoso.wmflabs.org/dbnaryetymology#",
	},
	{
		Name: "rdfs",
		URI:  "http://www.w3.org/2000/01/rdf-schema#",
	},
	{
		Name: "synsem",
		URI:  "http://www.w3.org/ns/lemon/synsem#",
	},
	{
		Name: "xs",
		URI:  "http://www.w3.org/2001/XMLSchema#",
	},
	{
		Name: "wikt",
		URI:  "https://en.wiktionary.org/wiki/",
	},
}

const (
	ID_eng                         = 0
	ID_dbnary                      = 1
	ID_dbnary_gloss                = 0
	ID_dbnary_writtenForm          = 1
	ID_dbnary_Translation          = 2
	ID_dbnary_isTranslationOf      = 3
	ID_dbnary_targetLanguage       = 4
	ID_dbnary_usage                = 5
	ID_dbnary_partOfSpeech         = 6
	ID_dbnary_describes            = 7
	ID_dbnary_Page                 = 8
	ID_dbnary_senseNumber          = 9
	ID_dbnary_Gloss                = 10
	ID_dbnary_rank                 = 11
	ID_dbnary_synonym              = 12
	ID_dbnary_hyponym              = 13
	ID_dbnary_antonym              = 14
	ID_dbnary_hypernym             = 15
	ID_dbnary_meronym              = 16
	ID_dbnary_targetLanguageCode   = 17
	ID_dbnary_holonym              = 18
	ID_dbnary_troponym             = 19
	ID_rdf                         = 2
	ID_rdf_type                    = 0
	ID_rdf_value                   = 1
	ID_rdf_predicate               = 2
	ID_rdf_subject                 = 3
	ID_rdf_Statement               = 4
	ID_rdf_object                  = 5
	ID_ontolex                     = 3
	ID_ontolex_writtenRep          = 0
	ID_ontolex_Form                = 1
	ID_ontolex_LexicalEntry        = 2
	ID_ontolex_canonicalForm       = 3
	ID_ontolex_sense               = 4
	ID_ontolex_LexicalSense        = 5
	ID_ontolex_phoneticRep         = 6
	ID_ontolex_MultiWordExpression = 7
	ID_lexvo                       = 4
	ID_lexvo_eng                   = 0
	ID_lexvo_rus                   = 1
	ID_lexvo_fin                   = 2
	ID_lexvo_deu                   = 3
	ID_lexvo_cmn                   = 4
	ID_lexvo_fra                   = 5
	ID_lexvo_spa                   = 6
	ID_lexvo_por                   = 7
	ID_lexvo_ita                   = 8
	ID_lexvo_nld                   = 9
	ID_lexvo_swe                   = 10
	ID_lexvo_jpn                   = 11
	ID_lexvo_pol                   = 12
	ID_lexvo_hbs                   = 13
	ID_lexvo_ell                   = 14
	ID_lexvo_ces                   = 15
	ID_lexvo_ron                   = 16
	ID_lexvo_cat                   = 17
	ID_lexvo_hun                   = 18
	ID_lexvo_bul                   = 19
	ID_lexvo_tur                   = 20
	ID_lexvo_dan                   = 21
	ID_lexvo_gle                   = 22
	ID_lexvo_kor                   = 23
	ID_lexvo_mri                   = 24
	ID_lexvo_ara                   = 25
	ID_lexvo_hye                   = 26
	ID_lexvo_epo                   = 27
	ID_lexvo_kat                   = 28
	ID_lexvo_fas                   = 29
	ID_lexvo_isl                   = 30
	ID_lexvo_glg                   = 31
	ID_lexvo_mkd                   = 32
	ID_lexvo_heb                   = 33
	ID_lexvo_lat                   = 34
	ID_lexvo_vie                   = 35
	ID_lexvo_ukr                   = 36
	ID_lexvo_nor                   = 37
	ID_lexvo_nob                   = 38
	ID_lexvo_gla                   = 39
	ID_lexvo_tgl                   = 40
	ID_lexvo_lav                   = 41
	ID_lexvo_est                   = 42
	ID_lexvo_msa                   = 43
	ID_lexvo_nno                   = 44
	ID_lexvo_slv                   = 45
	ID_lexvo_tha                   = 46
	ID_lexvo_slk                   = 47
	ID_lexvo_bel                   = 48
	ID_lexvo_hin                   = 49
	ID_lexvo_ind                   = 50
	ID_lexvo_lit                   = 51
	ID_lexvo_cym                   = 52
	ID_lexvo_yue                   = 53
	ID_lexvo_ido                   = 54
	ID_lexvo_swa                   = 55
	ID_lexvo_vol                   = 56
	ID_lexvo_khm                   = 57
	ID_lexvo_tel                   = 58
	ID_lexvo_fao                   = 59
	ID_lexvo_nrf                   = 60
	ID_lexvo_kur                   = 61
	ID_lexvo_yid                   = 62
	ID_lexvo_sqi                   = 63
	ID_lexvo_aze                   = 64
	ID_lexvo_oci                   = 65
	ID_lexvo_kaz                   = 66
	ID_lexvo_mon                   = 67
	ID_lexvo_urd                   = 68
	ID_lexvo_nav                   = 69
	ID_lexvo_eus                   = 70
	ID_lexvo_afr                   = 71
	ID_lexvo_ast                   = 72
	ID_lexvo_uzb                   = 73
	ID_lexvo_lao                   = 74
	ID_lexvo_tgk                   = 75
	ID_lexvo_ben                   = 76
	ID_lexvo_kir                   = 77
	ID_lexvo_glv                   = 78
	ID_lexvo_mya                   = 79
	ID_lexvo_nan                   = 80
	ID_lexvo_wln                   = 81
	ID_lexvo_mlt                   = 82
	ID_lexvo_bak                   = 83
	ID_lexvo_grc                   = 84
	ID_lexvo_roh                   = 85
	ID_lexvo_srd                   = 86
	ID_lexvo_mar                   = 87
	ID_lexvo_ltz                   = 88
	ID_lexvo_ina                   = 89
	ID_lexvo_tuk                   = 90
	ID_lexvo_tam                   = 91
	ID_lexvo_bre                   = 92
	ID_lexvo_vec                   = 93
	ID_lexvo_ang                   = 94
	ID_lexvo_rup                   = 95
	ID_lexvo_que                   = 96
	ID_lexvo_fur                   = 97
	ID_lexvo_dsb                   = 98
	ID_lexvo_uig                   = 99
	ID_lexvo_arz                   = 100
	ID_lexvo_pus                   = 101
	ID_lexvo_fry                   = 102
	ID_lexvo_scn                   = 103
	ID_lexvo_mal                   = 104
	ID_lexvo_tat                   = 105
	ID_lexvo_haw                   = 106
	ID_lexvo_sco                   = 107
	ID_lexvo_guj                   = 108
	ID_lexvo_san                   = 109
	ID_lexvo_asm                   = 110
	ID_lexvo_ltg                   = 111
	ID_lexvo_kan                   = 112
	ID_lexvo_crh                   = 113
	ID_lexvo_sin                   = 114
	ID_lexvo_hak                   = 115
	ID_lexvo_bod                   = 116
	ID_lexvo_cor                   = 117
	ID_lexvo_zza                   = 118
	ID_lexvo_kal                   = 119
	ID_lexvo_amh                   = 120
	ID_lexvo_hsb                   = 121
	ID_lexvo_zdj                   = 122
	ID_lexvo_pan                   = 123
	ID_lexvo_chr                   = 124
	ID_lexvo_hat                   = 125
	ID_lexvo_sme                   = 126
	ID_lexvo_zul                   = 127
	ID_lexvo_jav                   = 128
	ID_lexvo_jbo                   = 129
	ID_lexvo_nds                   = 130
	ID_lexvo_hau                   = 131
	ID_lexvo_ceb                   = 132
	ID_lexvo_nov                   = 133
	ID_lexvo_nep                   = 134
	ID_lexvo_che                   = 135
	ID_lexvo_oss                   = 136
	ID_lexvo_tpi                   = 137
	ID_lexvo_xal                   = 138
	ID_lexvo_chu                   = 139
	ID_lexvo_chv                   = 140
	ID_lexvo_arg                   = 141
	ID_lexvo_ckb                   = 142
	ID_lexvo_arc                   = 143
	ID_lexvo_dlm                   = 144
	ID_lexvo_ary                   = 145
	ID_lexvo_sun                   = 146
	ID_lexvo_ori                   = 147
	ID_lexvo_fro                   = 148
	ID_lexvo_oji                   = 149
	ID_lexvo_sah                   = 150
	ID_lexvo_myv                   = 151
	ID_lexvo_nap                   = 152
	ID_lexvo_non                   = 153
	ID_lexvo_xcl                   = 154
	ID_lexvo_mlg                   = 155
	ID_lexvo_nah                   = 156
	ID_lexvo_frr                   = 157
	ID_lexvo_alt                   = 158
	ID_lexvo_sga                   = 159
	ID_lexvo_abk                   = 160
	ID_lexvo_ewe                   = 161
	ID_lexvo_div                   = 162
	ID_lexvo_got                   = 163
	ID_lexvo_lld                   = 164
	ID_lexvo_zha                   = 165
	ID_lexvo_ota                   = 166
	ID_lexvo_cdo                   = 167
	ID_lexvo_tyv                   = 168
	ID_lexvo_acw                   = 169
	ID_lexvo_gsw                   = 170
	ID_lexvo_csb                   = 171
	ID_lexvo_rue                   = 172
	ID_lexvo_lim                   = 173
	ID_lexvo_yor                   = 174
	ID_lexvo_ccc                   = 175
	ID_lexvo_esu                   = 176
	ID_lexvo_nya                   = 177
	ID_lexvo_wol                   = 178
	ID_lexvo_war                   = 179
	ID_lexvo_ile                   = 180
	ID_lexvo_som                   = 181
	ID_lexvo_sot                   = 182
	ID_lexvo_udm                   = 183
	ID_lexvo_mwl                   = 184
	ID_lexvo_cop                   = 185
	ID_lexvo_lad                   = 186
	ID_lexvo_snd                   = 187
	ID_lexvo_iku                   = 188
	ID_lexvo_dng                   = 189
	ID_lexvo_ase                   = 190
	ID_lexvo_ady                   = 191
	ID_lexvo_grn                   = 192
	ID_lexvo_twf                   = 193
	ID_lexvo_wuu                   = 194
	ID_lexvo_cos                   = 195
	ID_lexvo_ist                   = 196
	ID_lexvo_kum                   = 197
	ID_lexvo_rom                   = 198
	ID_lexvo_tir                   = 199
	ID_lexvo_ain                   = 200
	ID_lexvo_luy                   = 201
	ID_lexvo_sms                   = 202
	ID_lexvo_mel                   = 203
	ID_lexvo_bar                   = 204
	ID_lexvo_ava                   = 205
	ID_lexvo_hil                   = 206
	ID_lexvo_nci                   = 207
	ID_lexvo_swb                   = 208
	ID_lexvo_wym                   = 209
	ID_lexvo_mnc                   = 210
	ID_lexvo_syc                   = 211
	ID_lexvo_kmr                   = 212
	ID_lexvo_smo                   = 213
	ID_lexvo_krc                   = 214
	ID_lexvo_khb                   = 215
	ID_lexvo_ilo                   = 216
	ID_lexvo_lij                   = 217
	ID_lexvo_sgs                   = 218
	ID_lexvo_syl                   = 219
	ID_lexvo_cjs                   = 220
	ID_lexvo_bua                   = 221
	ID_lexvo_lmo                   = 222
	ID_lexvo_kik                   = 223
	ID_lexvo_szl                   = 224
	ID_lexvo_bal                   = 225
	ID_lexvo_pdt                   = 226
	ID_lexvo_sdc                   = 227
	ID_lexvo_nmn                   = 228
	ID_lexvo_zho                   = 229
	ID_lexvo_kjh                   = 230
	ID_lexvo_nog                   = 231
	ID_lexvo_orv                   = 232
	ID_lexvo_ton                   = 233
	ID_lexvo_sdn                   = 234
	ID_lexvo_lez                   = 235
	ID_lexvo_xho                   = 236
	ID_lexvo_fij                   = 237
	ID_lexvo_pal                   = 238
	ID_lexvo_lin                   = 239
	ID_lexvo_mdf                   = 240
	ID_lexvo_rap                   = 241
	ID_lexvo_vro                   = 242
	ID_lexvo_stq                   = 243
	ID_lexvo_lkt                   = 244
	ID_lexvo_frm                   = 245
	ID_lexvo_pms                   = 246
	ID_lexvo_liv                   = 247
	ID_lexvo_yai                   = 248
	ID_lexvo_inh                   = 249
	ID_lexvo_frp                   = 250
	ID_lexvo_pap                   = 251
	ID_lexvo_osx                   = 252
	ID_lexvo_mzn                   = 253
	ID_lexvo_ext                   = 254
	ID_lexvo_vep                   = 255
	ID_lexvo_kin                   = 256
	ID_lexvo_kbd                   = 257
	ID_lexvo_mer                   = 258
	ID_lexvo_luo                   = 259
	ID_lexvo_ibo                   = 260
	ID_lexvo_shn                   = 261
	ID_lexvo_lus                   = 262
	ID_lexvo_ovd                   = 263
	ID_lexvo_koi                   = 264
	ID_lexvo_smn                   = 265
	ID_lexvo_egl                   = 266
	ID_lexvo_enm                   = 267
	ID_lexvo_tzm                   = 268
	ID_lexvo_chm                   = 269
	ID_lexvo_aym                   = 270
	ID_lexvo_xmf                   = 271
	ID_lexvo_ppl                   = 272
	ID_lexvo_pjt                   = 273
	ID_lexvo_dtp                   = 274
	ID_lexvo_kam                   = 275
	ID_lexvo_kaa                   = 276
	ID_lexvo_krl                   = 277
	ID_lexvo_ryu                   = 278
	ID_lexvo_gag                   = 279
	ID_lexvo_tah                   = 280
	ID_lexvo_bdr                   = 281
	ID_lexvo_vot                   = 282
	ID_lexvo_shi                   = 283
	ID_lexvo_goh                   = 284
	ID_lexvo_sna                   = 285
	ID_lexvo_otk                   = 286
	ID_lexvo_sva                   = 287
	ID_lexvo_pli                   = 288
	ID_lexvo_pdc                   = 289
	ID_lexvo_pam                   = 290
	ID_lexvo_sne                   = 291
	ID_lexvo_tly                   = 292
	ID_lexvo_mic                   = 293
	ID_lexvo_yua                   = 294
	ID_lexvo_kpv                   = 295
	ID_lexvo_tet                   = 296
	ID_lexvo_cre                   = 297
	ID_lexvo_akk                   = 298
	ID_lexvo_evn                   = 299
	ID_lexvo_kln                   = 300
	ID_lexvo_pro                   = 301
	ID_lexvo_new                   = 302
	ID_lexvo_bcl                   = 303
	ID_lexvo_akz                   = 304
	ID_lexvo_chy                   = 305
	ID_lexvo_lbe                   = 306
	ID_lexvo_lug                   = 307
	ID_lexvo_tsn                   = 308
	ID_lexvo_sma                   = 309
	ID_lexvo_tpn                   = 310
	ID_lexvo_smj                   = 311
	ID_lexvo_iba                   = 312
	ID_lexvo_nau                   = 313
	ID_lexvo_win                   = 314
	ID_lexvo_sat                   = 315
	ID_lexvo_lzz                   = 316
	ID_lexvo_izh                   = 317
	ID_lexvo_apw                   = 318
	ID_lexvo_mns                   = 319
	ID_lexvo_wmw                   = 320
	ID_lexvo_txb                   = 321
	ID_lexvo_kac                   = 322
	ID_lexvo_unm                   = 323
	ID_lexvo_nay                   = 324
	ID_lexvo_mww                   = 325
	ID_lexvo_kqr                   = 326
	ID_lexvo_mvv                   = 327
	ID_lexvo_ace                   = 328
	ID_lexvo_pnb                   = 329
	ID_lexvo_sux                   = 330
	ID_lexvo_chk                   = 331
	ID_lexvo_hop                   = 332
	ID_lexvo_arn                   = 333
	ID_lexvo_tli                   = 334
	ID_lexvo_pih                   = 335
	ID_lexvo_bam                   = 336
	ID_lexvo_kab                   = 337
	ID_lexvo_abq                   = 338
	ID_lexvo_gmh                   = 339
	ID_lexvo_dar                   = 340
	ID_lexvo_ofs                   = 341
	ID_lexvo_bis                   = 342
	ID_lexvo_dzo                   = 343
	ID_lexvo_srn                   = 344
	ID_lexvo_kon                   = 345
	ID_lexvo_xto                   = 346
	ID_lexvo_niv                   = 347
	ID_lexvo_ipk                   = 348
	ID_lexvo_moe                   = 349
	ID_lexvo_ban                   = 350
	ID_lexvo_dlg                   = 351
	ID_lexvo_prg                   = 352
	ID_lexvo_abe                   = 353
	ID_lexvo_mrw                   = 354
	ID_lexvo_bjn                   = 355
	ID_lexvo_vls                   = 356
	ID_lexvo_rhg                   = 357
	ID_lexvo_ood                   = 358
	ID_lexvo_bqi                   = 359
	ID_lexvo_ebk                   = 360
	ID_lexvo_cha                   = 361
	ID_lexvo_rgn                   = 362
	ID_lexvo_xqa                   = 363
	ID_lexvo_peo                   = 364
	ID_lexvo_ruo                   = 365
	ID_lexvo_mnw                   = 366
	ID_lexvo_ale                   = 367
	ID_lexvo_naq                   = 368
	ID_lexvo_duo                   = 369
	ID_lexvo_sce                   = 370
	ID_lexvo_iii                   = 371
	ID_lexvo_tab                   = 372
	ID_lexvo_aka                   = 373
	ID_lexvo_mrc                   = 374
	ID_lexvo_pox                   = 375
	ID_lexvo_ave                   = 376
	ID_lexvo_nso                   = 377
	ID_lexvo_avd                   = 378
	ID_lexvo_gil                   = 379
	ID_lexvo_ssw                   = 380
	ID_lexvo_kld                   = 381
	ID_lexvo_gez                   = 382
	ID_lexvo_swg                   = 383
	ID_lexvo_pcd                   = 384
	ID_lexvo_mrj                   = 385
	ID_lexvo_uga                   = 386
	ID_lexvo_kok                   = 387
	ID_lexvo_zun                   = 388
	ID_lexvo_ksc                   = 389
	ID_lexvo_sib                   = 390
	ID_lexvo_bpy                   = 391
	ID_lexvo_gan                   = 392
	ID_lexvo_knb                   = 393
	ID_lexvo_ill                   = 394
	ID_lexvo_wrh                   = 395
	ID_lexvo_tiw                   = 396
	ID_lexvo_sei                   = 397
	ID_lexvo_raj                   = 398
	ID_lexvo_bbl                   = 399
	ID_lexvo_zai                   = 400
	ID_lexvo_ruq                   = 401
	ID_lexvo_ems                   = 402
	ID_lexvo_xpr                   = 403
	ID_lexvo_kea                   = 404
	ID_lexvo_cho                   = 405
	ID_lexvo_hif                   = 406
	ID_lexvo_aeb                   = 407
	ID_lexvo_min                   = 408
	ID_lexvo_rop                   = 409
	ID_lexvo_apm                   = 410
	ID_lexvo_yrk                   = 411
	ID_lexvo_din                   = 412
	ID_lexvo_dak                   = 413
	ID_lexvo_udi                   = 414
	ID_lexvo_mas                   = 415
	ID_lexvo_dta                   = 416
	ID_lexvo_amu                   = 417
	ID_lexvo_jam                   = 418
	ID_lexvo_vor                   = 419
	ID_lexvo_orm                   = 420
	ID_lexvo_aii                   = 421
	ID_lexvo_xfa                   = 422
	ID_lexvo_bth                   = 423
	ID_lexvo_kca                   = 424
	ID_lexvo_dtr                   = 425
	ID_lexvo_wbp                   = 426
	ID_lexvo_bol                   = 427
	ID_lexvo_apc                   = 428
	ID_lexvo_tzo                   = 429
	ID_lexvo_mvi                   = 430
	ID_lexvo_hrx                   = 431
	ID_lexvo_tvl                   = 432
	ID_lexvo_ket                   = 433
	ID_lexvo_mai                   = 434
	ID_lexvo_tih                   = 435
	ID_lexvo_pcc                   = 436
	ID_lexvo_otw                   = 437
	ID_lexvo_ifk                   = 438
	ID_lexvo_nuk                   = 439
	ID_lexvo_xda                   = 440
	ID_lexvo_fit                   = 441
	ID_lexvo_drg                   = 442
	ID_lexvo_atv                   = 443
	ID_lexvo_kpy                   = 444
	ID_lexvo_kim                   = 445
	ID_lexvo_osa                   = 446
	ID_lexvo_esh                   = 447
	ID_lexvo_cic                   = 448
	ID_lexvo_mul                   = 449
	ID_lexvo_par                   = 450
	ID_lexvo_gld                   = 451
	ID_lexvo_sjd                   = 452
	ID_lexvo_bho                   = 453
	ID_lexvo_kas                   = 454
	ID_lexvo_kri                   = 455
	ID_lexvo_axm                   = 456
	ID_lexvo_tig                   = 457
	ID_lexvo_afb                   = 458
	ID_lexvo_zgh                   = 459
	ID_lexvo_pag                   = 460
	ID_lexvo_gcf                   = 461
	ID_lexvo_pua                   = 462
	ID_lexvo_tfn                   = 463
	ID_lexvo_phn                   = 464
	ID_lexvo_pau                   = 465
	ID_lexvo_brh                   = 466
	ID_lexvo_mah                   = 467
	ID_lexvo_lnd                   = 468
	ID_lexvo_ckt                   = 469
	ID_lexvo_str                   = 470
	ID_lexvo_aqc                   = 471
	ID_lexvo_atj                   = 472
	ID_lexvo_moh                   = 473
	ID_lexvo_ccp                   = 474
	ID_lexvo_mad                   = 475
	ID_lexvo_pot                   = 476
	ID_lexvo_zea                   = 477
	ID_lexvo_hit                   = 478
	ID_lexvo_xnn                   = 479
	ID_lexvo_pnt                   = 480
	ID_lexvo_apj                   = 481
	ID_lexvo_com                   = 482
	ID_lexvo_nod                   = 483
	ID_lexvo_aoz                   = 484
	ID_lexvo_ven                   = 485
	ID_lexvo_bla                   = 486
	ID_lexvo_egy                   = 487
	ID_lexvo_wls                   = 488
	ID_lexvo_mnk                   = 489
	ID_lexvo_yur                   = 490
	ID_lexvo_chc                   = 491
	ID_lexvo_bug                   = 492
	ID_lexvo_jct                   = 493
	ID_lexvo_mga                   = 494
	ID_lexvo_tso                   = 495
	ID_lexvo_hsn                   = 496
	ID_lexvo_cjy                   = 497
	ID_lexvo_pon                   = 498
	ID_lexvo_chl                   = 499
	ID_lexvo_ksw                   = 500
	ID_lexvo_ful                   = 501
	ID_lexvo_gml                   = 502
	ID_lexvo_umb                   = 503
	ID_lexvo_meu                   = 504
	ID_lexvo_alq                   = 505
	ID_lexvo_crk                   = 506
	ID_lexvo_xpq                   = 507
	ID_lexvo_olo                   = 508
	ID_lexvo_yoi                   = 509
	ID_lexvo_quc                   = 510
	ID_lexvo_mhn                   = 511
	ID_lexvo_sog                   = 512
	ID_lexvo_kla                   = 513
	ID_lexvo_kky                   = 514
	ID_lexvo_hts                   = 515
	ID_lexvo_bsy                   = 516
	ID_lexvo_und                   = 517
	ID_lexvo_rar                   = 518
	ID_lexvo_nio                   = 519
	ID_lexvo_niu                   = 520
	ID_lexvo_tzj                   = 521
	ID_lexvo_cak                   = 522
	ID_lexvo_gni                   = 523
	ID_lexvo_dgr                   = 524
	ID_lexvo_sag                   = 525
	ID_lexvo_arp                   = 526
	ID_lexvo_cab                   = 527
	ID_lexvo_nys                   = 528
	ID_lexvo_odt                   = 529
	ID_lexvo_cuk                   = 530
	ID_lexvo_wni                   = 531
	ID_lexvo_dum                   = 532
	ID_lexvo_blt                   = 533
	ID_lexvo_aae                   = 534
	ID_lexvo_jra                   = 535
	ID_lexvo_mfe                   = 536
	ID_lexvo_paw                   = 537
	ID_lexvo_ses                   = 538
	ID_lexvo_ykg                   = 539
	ID_lexvo_aar                   = 540
	ID_lexvo_rmf                   = 541
	ID_lexvo_tts                   = 542
	ID_lexvo_xul                   = 543
	ID_lexvo_oge                   = 544
	ID_lexvo_sea                   = 545
	ID_lexvo_chn                   = 546
	ID_lexvo_gul                   = 547
	ID_lexvo_tcy                   = 548
	ID_lexvo_hid                   = 549
	ID_lexvo_bor                   = 550
	ID_lexvo_sou                   = 551
	ID_lexvo_itl                   = 552
	ID_lexvo_umu                   = 553
	ID_lexvo_sje                   = 554
	ID_lexvo_mtq                   = 555
	ID_lexvo_wbl                   = 556
	ID_lexvo_cim                   = 557
	ID_lexvo_wam                   = 558
	ID_lexvo_maz                   = 559
	ID_lexvo_fkv                   = 560
	ID_lexvo_rmn                   = 561
	ID_lexvo_myp                   = 562
	ID_lexvo_eve                   = 563
	ID_lexvo_nwy                   = 564
	ID_lexvo_ivv                   = 565
	ID_lexvo_ote                   = 566
	ID_lexvo_wrr                   = 567
	ID_lexvo_xum                   = 568
	ID_lexvo_wlc                   = 569
	ID_lexvo_fud                   = 570
	ID_lexvo_caa                   = 571
	ID_lexvo_sac                   = 572
	ID_lexvo_bem                   = 573
	ID_lexvo_bsb                   = 574
	ID_lexvo_osp                   = 575
	ID_lexvo_zpq                   = 576
	ID_lexvo_ttt                   = 577
	ID_lexvo_zko                   = 578
	ID_lexvo_mrv                   = 579
	ID_lexvo_rad                   = 580
	ID_lexvo_amm                   = 581
	ID_lexvo_cbk                   = 582
	ID_lexvo_gaa                   = 583
	ID_lexvo_aln                   = 584
	ID_lexvo_agr                   = 585
	ID_lexvo_mop                   = 586
	ID_lexvo_kjj                   = 587
	ID_lexvo_kzt                   = 588
	ID_lexvo_xug                   = 589
	ID_lexvo_ajp                   = 590
	ID_lexvo_pim                   = 591
	ID_lexvo_kdr                   = 592
	ID_lexvo_sas                   = 593
	ID_lexvo_shh                   = 594
	ID_lexvo_kha                   = 595
	ID_lexvo_gin                   = 596
	ID_lexvo_tsg                   = 597
	ID_lexvo_mus                   = 598
	ID_lexvo_kap                   = 599
	ID_lexvo_ari                   = 600
	ID_lexvo_aht                   = 601
	ID_lexvo_shs                   = 602
	ID_lexvo_acm                   = 603
	ID_lexvo_ams                   = 604
	ID_lexvo_asb                   = 605
	ID_lexvo_kut                   = 606
	ID_lexvo_ach                   = 607
	ID_lexvo_rmy                   = 608
	ID_lexvo_dty                   = 609
	ID_lexvo_dbl                   = 610
	ID_lexvo_hai                   = 611
	ID_lexvo_yol                   = 612
	ID_lexvo_sli                   = 613
	ID_lexvo_pis                   = 614
	ID_lexvo_mak                   = 615
	ID_lexvo_xbc                   = 616
	ID_lexvo_nia                   = 617
	ID_lexvo_inz                   = 618
	ID_lexvo_lou                   = 619
	ID_lexvo_yij                   = 620
	ID_lexvo_kis                   = 621
	ID_lexvo_osc                   = 622
	ID_lexvo_pqm                   = 623
	ID_lexvo_ljp                   = 624
	ID_lexvo_agx                   = 625
	ID_lexvo_mbf                   = 626
	ID_lexvo_ada                   = 627
	ID_lexvo_yao                   = 628
	ID_lexvo_xsr                   = 629
	ID_lexvo_vma                   = 630
	ID_lexvo_brg                   = 631
	ID_lexvo_nus                   = 632
	ID_lexvo_hac                   = 633
	ID_lexvo_nrm                   = 634
	ID_lexvo_pmh                   = 635
	ID_lexvo_tqw                   = 636
	ID_lexvo_kjg                   = 637
	ID_lexvo_cpg                   = 638
	ID_lexvo_yox                   = 639
	ID_lexvo_jdt                   = 640
	ID_lexvo_gue                   = 641
	ID_lexvo_pkc                   = 642
	ID_lexvo_fla                   = 643
	ID_lexvo_wlm                   = 644
	ID_lexvo_gmy                   = 645
	ID_lexvo_sjt                   = 646
	ID_lexvo_tbl                   = 647
	ID_lexvo_csm                   = 648
	ID_lexvo_coo                   = 649
	ID_lexvo_ulk                   = 650
	ID_lexvo_ctu                   = 651
	ID_lexvo_mia                   = 652
	ID_lexvo_kij                   = 653
	ID_lexvo_sju                   = 654
	ID_lexvo_yaq                   = 655
	ID_lexvo_ddo                   = 656
	ID_lexvo_kgp                   = 657
	ID_lexvo_tyz                   = 658
	ID_lexvo_ksk                   = 659
	ID_lexvo_cja                   = 660
	ID_lexvo_mcm                   = 661
	ID_lexvo_cni                   = 662
	ID_lexvo_mwp                   = 663
	ID_lexvo_gdo                   = 664
	ID_lexvo_yuf                   = 665
	ID_lexvo_ude                   = 666
	ID_lexvo_lud                   = 667
	ID_lexvo_oma                   = 668
	ID_lexvo_rif                   = 669
	ID_lexvo_glk                   = 670
	ID_lexvo_rut                   = 671
	ID_lexvo_kaw                   = 672
	ID_lexvo_tum                   = 673
	ID_lexvo_mdh                   = 674
	ID_lexvo_crj                   = 675
	ID_lexvo_arq                   = 676
	ID_lexvo_anq                   = 677
	ID_lexvo_msn                   = 678
	ID_lexvo_pgl                   = 679
	ID_lexvo_wim                   = 680
	ID_lexvo_ksi                   = 681
	ID_lexvo_mrq                   = 682
	ID_lexvo_nuj                   = 683
	ID_lexvo_aau                   = 684
	ID_lexvo_tpw                   = 685
	ID_lexvo_kfr                   = 686
	ID_lexvo_kva                   = 687
	ID_lexvo_koy                   = 688
	ID_lexvo_bkd                   = 689
	ID_lexvo_tin                   = 690
	ID_lexvo_tdd                   = 691
	ID_lexvo_kyh                   = 692
	ID_lexvo_jrb                   = 693
	ID_lexvo_aaq                   = 694
	ID_lexvo_yuy                   = 695
	ID_lexvo_kmk                   = 696
	ID_lexvo_doi                   = 697
	ID_lexvo_sip                   = 698
	ID_lexvo_miq                   = 699
	ID_lexvo_akl                   = 700
	ID_lexvo_mhx                   = 701
	ID_lexvo_loz                   = 702
	ID_lexvo_cas                   = 703
	ID_lexvo_pov                   = 704
	ID_lexvo_uum                   = 705
	ID_lexvo_zav                   = 706
	ID_lexvo_bew                   = 707
	ID_lexvo_tsi                   = 708
	ID_lexvo_pao                   = 709
	ID_lexvo_bvr                   = 710
	ID_lexvo_awa                   = 711
	ID_lexvo_kwk                   = 712
	ID_lexvo_mqm                   = 713
	ID_lexvo_kzg                   = 714
	ID_lexvo_okn                   = 715
	ID_lexvo_pkp                   = 716
	ID_lexvo_chg                   = 717
	ID_lexvo_chp                   = 718
	ID_lexvo_apy                   = 719
	ID_lexvo_xkl                   = 720
	ID_lexvo_psu                   = 721
	ID_lexvo_crs                   = 722
	ID_lexvo_pgn                   = 723
	ID_lexvo_tcs                   = 724
	ID_lexvo_clm                   = 725
	ID_lexvo_tsd                   = 726
	ID_lexvo_ett                   = 727
	ID_lexvo_nst                   = 728
	ID_lexvo_cku                   = 729
	ID_lexvo_tta                   = 730
	ID_lexvo_dwu                   = 731
	ID_lexvo_tzh                   = 732
	ID_lexvo_nhx                   = 733
	ID_lexvo_xdc                   = 734
	ID_lexvo_fon                   = 735
	ID_lexvo_crl                   = 736
	ID_lexvo_hus                   = 737
	ID_lexvo_emb                   = 738
	ID_lexvo_ssf                   = 739
	ID_lexvo_poe                   = 740
	ID_lexvo_apk                   = 741
	ID_lexvo_lwg                   = 742
	ID_lexvo_mlm                   = 743
	ID_lexvo_tiy                   = 744
	ID_lexvo_lzh                   = 745
	ID_lexvo_qua                   = 746
	ID_lexvo_ckv                   = 747
	ID_lexvo_lui                   = 748
	ID_lexvo_xdk                   = 749
	ID_lexvo_krj                   = 750
	ID_lexvo_shp                   = 751
	ID_lexvo_fax                   = 752
	ID_lexvo_nnt                   = 753
	ID_lexvo_kmv                   = 754
	ID_lexvo_car                   = 755
	ID_lexvo_bej                   = 756
	ID_lexvo_trv                   = 757
	ID_lexvo_pmt                   = 758
	ID_lexvo_cay                   = 759
	ID_lexvo_ami                   = 760
	ID_lexvo_cia                   = 761
	ID_lexvo_rmq                   = 762
	ID_lexvo_ibg                   = 763
	ID_lexvo_ksd                   = 764
	ID_lexvo_bdk                   = 765
	ID_lexvo_aer                   = 766
	ID_lexvo_pad                   = 767
	ID_lexvo_gub                   = 768
	ID_lexvo_pcm                   = 769
	ID_lexvo_btx                   = 770
	ID_lexvo_gug                   = 771
	ID_lexvo_aaa                   = 772
	ID_lexvo_xpm                   = 773
	ID_lexvo_her                   = 774
	ID_lexvo_taa                   = 775
	ID_lexvo_lsd                   = 776
	ID_lexvo_gut                   = 777
	ID_lexvo_msk                   = 778
	ID_lexvo_xnt                   = 779
	ID_lexvo_bbc                   = 780
	ID_lexvo_nbl                   = 781
	ID_lexvo_lki                   = 782
	ID_lexvo_wlo                   = 783
	ID_lexvo_yuk                   = 784
	ID_lexvo_gwi                   = 785
	ID_lexvo_eka                   = 786
	ID_lexvo_hmc                   = 787
	ID_lexvo_mdr                   = 788
	ID_lexvo_pah                   = 789
	ID_lexvo_mov                   = 790
	ID_lexvo_klb                   = 791
	ID_lexvo_tae                   = 792
	ID_lexvo_tus                   = 793
	ID_lexvo_huz                   = 794
	ID_lexvo_nij                   = 795
	ID_lexvo_lbk                   = 796
	ID_lexvo_yap                   = 797
	ID_lexvo_ycn                   = 798
	ID_lexvo_kyi                   = 799
	ID_lexvo_sdo                   = 800
	ID_lexvo_kbp                   = 801
	ID_lexvo_rth                   = 802
	ID_lexvo_skb                   = 803
	ID_lexvo_ndd                   = 804
	ID_lexvo_vkp                   = 805
	ID_lexvo_xls                   = 806
	ID_lexvo_srm                   = 807
	ID_lexvo_bgs                   = 808
	ID_lexvo_sky                   = 809
	ID_lexvo_kmc                   = 810
	ID_lexvo_bdy                   = 811
	ID_lexvo_yag                   = 812
	ID_lexvo_dgc                   = 813
	ID_lexvo_abx                   = 814
	ID_lexvo_btd                   = 815
	ID_lexvo_mpm                   = 816
	ID_lexvo_acy                   = 817
	ID_lexvo_crx                   = 818
	ID_lexvo_rtm                   = 819
	ID_lexvo_bnn                   = 820
	ID_lexvo_skp                   = 821
	ID_lexvo_dbp                   = 822
	ID_lexvo_ayl                   = 823
	ID_lexvo_ryn                   = 824
	ID_lexvo_bkm                   = 825
	ID_lexvo_nfr                   = 826
	ID_lexvo_dif                   = 827
	ID_lexvo_wrz                   = 828
	ID_lexvo_tbc                   = 829
	ID_lexvo_urb                   = 830
	ID_lexvo_tsu                   = 831
	ID_lexvo_fia                   = 832
	ID_lexvo_abl                   = 833
	ID_lexvo_lep                   = 834
	ID_lexvo_nrn                   = 835
	ID_lexvo_xsv                   = 836
	ID_lexvo_oko                   = 837
	ID_lexvo_wau                   = 838
	ID_lexvo_aho                   = 839
	ID_lexvo_mhy                   = 840
	ID_lexvo_mos                   = 841
	ID_lexvo_chz                   = 842
	ID_lexvo_cac                   = 843
	ID_lexvo_cro                   = 844
	ID_lexvo_ndo                   = 845
	ID_lexvo_sjw                   = 846
	ID_lexvo_lmw                   = 847
	ID_lexvo_qui                   = 848
	ID_lexvo_css                   = 849
	ID_lexvo_crg                   = 850
	ID_lexvo_vmb                   = 851
	ID_lexvo_cle                   = 852
	ID_lexvo_uby                   = 853
	ID_lexvo_haa                   = 854
	ID_lexvo_coc                   = 855
	ID_lexvo_crd                   = 856
	ID_lexvo_bno                   = 857
	ID_lexvo_ing                   = 858
	ID_lexvo_blc                   = 859
	ID_lexvo_cmg                   = 860
	ID_lexvo_idb                   = 861
	ID_lexvo_lod                   = 862
	ID_lexvo_tay                   = 863
	ID_lexvo_aby                   = 864
	ID_lexvo_ctg                   = 865
	ID_lexvo_teo                   = 866
	ID_lexvo_cpa                   = 867
	ID_lexvo_ple                   = 868
	ID_lexvo_bde                   = 869
	ID_lexvo_lrc                   = 870
	ID_lexvo_hoi                   = 871
	ID_lexvo_mwf                   = 872
	ID_lexvo_pwn                   = 873
	ID_lexvo_cad                   = 874
	ID_lexvo_dbj                   = 875
	ID_lexvo_tee                   = 876
	ID_lexvo_fan                   = 877
	ID_lexvo_ute                   = 878
	ID_lexvo_bxk                   = 879
	ID_lexvo_rob                   = 880
	ID_lexvo_cup                   = 881
	ID_lexvo_pac                   = 882
	ID_lexvo_apl                   = 883
	ID_lexvo_mch                   = 884
	ID_lexvo_zoc                   = 885
	ID_lexvo_sgd                   = 886
	ID_lexvo_nez                   = 887
	ID_lexvo_guc                   = 888
	ID_lexvo_smr                   = 889
	ID_lexvo_djk                   = 890
	ID_lexvo_mnv                   = 891
	ID_lexvo_ess                   = 892
	ID_lexvo_chf                   = 893
	ID_lexvo_ngh                   = 894
	ID_lexvo_kai                   = 895
	ID_lexvo_hur                   = 896
	ID_lexvo_kkz                   = 897
	ID_lexvo_bsk                   = 898
	ID_lexvo_lmc                   = 899
	ID_lexvo_owl                   = 900
	ID_lexvo_hdn                   = 901
	ID_lexvo_kzj                   = 902
	ID_lexvo_huv                   = 903
	ID_lexvo_dny                   = 904
	ID_lexvo_xrn                   = 905
	ID_lexvo_tue                   = 906
	ID_lexvo_ium                   = 907
	ID_lexvo_enh                   = 908
	ID_lexvo_xta                   = 909
	ID_lexvo_kuz                   = 910
	ID_lexvo_nde                   = 911
	ID_lexvo_mpt                   = 912
	ID_lexvo_ibb                   = 913
	ID_lexvo_maq                   = 914
	ID_lexvo_jao                   = 915
	ID_lexvo_kqe                   = 916
	ID_lexvo_hix                   = 917
	ID_lexvo_tar                   = 918
	ID_lexvo_xlu                   = 919
	ID_lexvo_top                   = 920
	ID_lexvo_smk                   = 921
	ID_lexvo_mcb                   = 922
	ID_lexvo_mni                   = 923
	ID_lexvo_akr                   = 924
	ID_lexvo_bll                   = 925
	ID_lexvo_lua                   = 926
	ID_lexvo_efi                   = 927
	ID_lexvo_irk                   = 928
	ID_lexvo_ngi                   = 929
	ID_lexvo_wyi                   = 930
	ID_lexvo_crn                   = 931
	ID_lexvo_hnj                   = 932
	ID_lexvo_cux                   = 933
	ID_lexvo_tos                   = 934
	ID_lexvo_kua                   = 935
	ID_lexvo_azg                   = 936
	ID_lexvo_acu                   = 937
	ID_lexvo_tao                   = 938
	ID_lexvo_mpj                   = 939
	ID_lexvo_wyb                   = 940
	ID_lexvo_dje                   = 941
	ID_lexvo_hnn                   = 942
	ID_lexvo_bya                   = 943
	ID_lexvo_kge                   = 944
	ID_lexvo_xeb                   = 945
	ID_lexvo_urh                   = 946
	ID_lexvo_kjq                   = 947
	ID_lexvo_rol                   = 948
	ID_lexvo_agm                   = 949
	ID_lexvo_aca                   = 950
	ID_lexvo_hub                   = 951
	ID_lexvo_cce                   = 952
	ID_lexvo_hvn                   = 953
	ID_lexvo_nhn                   = 954
	ID_lexvo_trs                   = 955
	ID_lexvo_ttq                   = 956
	ID_lexvo_cui                   = 957
	ID_lexvo_nsq                   = 958
	ID_lexvo_nsk                   = 959
	ID_lexvo_ais                   = 960
	ID_lexvo_gbm                   = 961
	ID_lexvo_frk                   = 962
	ID_lexvo_cts                   = 963
	ID_lexvo_mse                   = 964
	ID_lexvo_mph                   = 965
	ID_lexvo_txg                   = 966
	ID_lexvo_woe                   = 967
	ID_lexvo_xsa                   = 968
	ID_lexvo_suk                   = 969
	ID_lexvo_sia                   = 970
	ID_lexvo_bcr                   = 971
	ID_lexvo_xpg                   = 972
	ID_lexvo_har                   = 973
	ID_lexvo_slr                   = 974
	ID_lexvo_end                   = 975
	ID_lexvo_gej                   = 976
	ID_lexvo_sxr                   = 977
	ID_lexvo_ian                   = 978
	ID_lexvo_aly                   = 979
	ID_lexvo_stp                   = 980
	ID_lexvo_mez                   = 981
	ID_lexvo_zku                   = 982
	ID_lexvo_xsm                   = 983
	ID_lexvo_bkl                   = 984
	ID_lexvo_elx                   = 985
	ID_lexvo_jup                   = 986
	ID_lexvo_mir                   = 987
	ID_lexvo_sxn                   = 988
	ID_lexvo_ncz                   = 989
	ID_lexvo_bcu                   = 990
	ID_lexvo_aqp                   = 991
	ID_lexvo_hax                   = 992
	ID_lexvo_lut                   = 993
	ID_lexvo_kvh                   = 994
	ID_lexvo_zkg                   = 995
	ID_lexvo_gdd                   = 996
	ID_lexvo_mlv                   = 997
	ID_lexvo_blm                   = 998
	ID_lexvo_bdg                   = 999
	ID_lexvo_ane                   = 1000
	ID_lexvo_one                   = 1001
	ID_lexvo_ims                   = 1002
	ID_lexvo_bjz                   = 1003
	ID_lexvo_mrn                   = 1004
	ID_lexvo_wya                   = 1005
	ID_lexvo_itv                   = 1006
	ID_lexvo_nbh                   = 1007
	ID_lexvo_oka                   = 1008
	ID_lexvo_xbr                   = 1009
	ID_lexvo_apn                   = 1010
	ID_lexvo_amn                   = 1011
	ID_lexvo_tnq                   = 1012
	ID_lexvo_ocu                   = 1013
	ID_lexvo_pak                   = 1014
	ID_lexvo_khw                   = 1015
	ID_lexvo_dig                   = 1016
	ID_lexvo_pok                   = 1017
	ID_lexvo_obi                   = 1018
	ID_lexvo_pyu                   = 1019
	ID_lexvo_xsy                   = 1020
	ID_lexvo_msb                   = 1021
	ID_lexvo_zpv                   = 1022
	ID_lexvo_kxd                   = 1023
	ID_lexvo_bpz                   = 1024
	ID_lexvo_kls                   = 1025
	ID_lexvo_mnb                   = 1026
	ID_lexvo_fab                   = 1027
	ID_lexvo_nlg                   = 1028
	ID_lexvo_dis                   = 1029
	ID_lexvo_ktn                   = 1030
	ID_lexvo_tbd                   = 1031
	ID_lexvo_kms                   = 1032
	ID_lexvo_cps                   = 1033
	ID_lexvo_scl                   = 1034
	ID_lexvo_mhl                   = 1035
	ID_lexvo_ibd                   = 1036
	ID_lexvo_csi                   = 1037
	ID_lexvo_ctm                   = 1038
	ID_lexvo_rys                   = 1039
	ID_lexvo_orh                   = 1040
	ID_lexvo_see                   = 1041
	ID_lexvo_yum                   = 1042
	ID_lexvo_jeb                   = 1043
	ID_lexvo_sob                   = 1044
	ID_lexvo_zaa                   = 1045
	ID_lexvo_mek                   = 1046
	ID_lexvo_sks                   = 1047
	ID_lexvo_mhq                   = 1048
	ID_lexvo_gun                   = 1049
	ID_lexvo_oui                   = 1050
	ID_lexvo_nbj                   = 1051
	ID_lexvo_oaa                   = 1052
	ID_lexvo_ojs                   = 1053
	ID_lexvo_cyo                   = 1054
	ID_lexvo_mjy                   = 1055
	ID_lexvo_fpe                   = 1056
	ID_lexvo_xnb                   = 1057
	ID_lexvo_kkh                   = 1058
	ID_lexvo_xgf                   = 1059
	ID_lexvo_bfa                   = 1060
	ID_lexvo_myx                   = 1061
	ID_lexvo_sel                   = 1062
	ID_lexvo_isi                   = 1063
	ID_lexvo_tkr                   = 1064
	ID_lexvo_gcr                   = 1065
	ID_lexvo_xav                   = 1066
	ID_lexvo_xcb                   = 1067
	ID_lexvo_ayn                   = 1068
	ID_lexvo_jiv                   = 1069
	ID_lexvo_bhp                   = 1070
	ID_lexvo_iwm                   = 1071
	ID_lexvo_aia                   = 1072
	ID_lexvo_xww                   = 1073
	ID_lexvo_urm                   = 1074
	ID_lexvo_cah                   = 1075
	ID_lexvo_ybh                   = 1076
	ID_lexvo_mhj                   = 1077
	ID_lexvo_cgc                   = 1078
	ID_lexvo_kpg                   = 1079
	ID_lexvo_mbl                   = 1080
	ID_lexvo_cir                   = 1081
	ID_lexvo_tbo                   = 1082
	ID_lexvo_xpu                   = 1083
	ID_lexvo_xce                   = 1084
	ID_lexvo_jmd                   = 1085
	ID_lexvo_aec                   = 1086
	ID_lexvo_apx                   = 1087
	ID_lexvo_sge                   = 1088
	ID_lexvo_bzj                   = 1089
	ID_lexvo_ksx                   = 1090
	ID_lexvo_njm                   = 1091
	ID_lexvo_lre                   = 1092
	ID_lexvo_ats                   = 1093
	ID_lexvo_zag                   = 1094
	ID_lexvo_sno                   = 1095
	ID_lexvo_zau                   = 1096
	ID_lexvo_oym                   = 1097
	ID_lexvo_plw                   = 1098
	ID_lexvo_mxe                   = 1099
	ID_lexvo_pni                   = 1100
	ID_lexvo_isd                   = 1101
	ID_lexvo_maa                   = 1102
	ID_lexvo_tpx                   = 1103
	ID_lexvo_gon                   = 1104
	ID_lexvo_zap                   = 1105
	ID_lexvo_wnw                   = 1106
	ID_lexvo_poi                   = 1107
	ID_lexvo_xsc                   = 1108
	ID_lexvo_ltc                   = 1109
	ID_lexvo_xtc                   = 1110
	ID_lexvo_bwi                   = 1111
	ID_lexvo_zoo                   = 1112
	ID_lexvo_cdm                   = 1113
	ID_lexvo_tkl                   = 1114
	ID_lexvo_tlv                   = 1115
	ID_lexvo_sae                   = 1116
	ID_lexvo_skh                   = 1117
	ID_lexvo_bgb                   = 1118
	ID_lexvo_ojp                   = 1119
	ID_lexvo_btw                   = 1120
	ID_lexvo_kys                   = 1121
	ID_lexvo_sln                   = 1122
	ID_lexvo_sjr                   = 1123
	ID_lexvo_akv                   = 1124
	ID_lexvo_tuo                   = 1125
	ID_lexvo_cbs                   = 1126
	ID_lexvo_aat                   = 1127
	ID_lexvo_arw                   = 1128
	ID_lexvo_haz                   = 1129
	ID_lexvo_uun                   = 1130
	ID_lexvo_xld                   = 1131
	ID_lexvo_tag                   = 1132
	ID_lexvo_mlu                   = 1133
	ID_lexvo_dus                   = 1134
	ID_lexvo_jac                   = 1135
	ID_lexvo_abd                   = 1136
	ID_lexvo_iyx                   = 1137
	ID_lexvo_for                   = 1138
	ID_lexvo_nwa                   = 1139
	ID_lexvo_ore                   = 1140
	ID_lexvo_obm                   = 1141
	ID_lexvo_tvo                   = 1142
	ID_lexvo_ali                   = 1143
	ID_lexvo_ahr                   = 1144
	ID_lexvo_zor                   = 1145
	ID_lexvo_alj                   = 1146
	ID_lexvo_bpr                   = 1147
	ID_lexvo_svm                   = 1148
	ID_lexvo_veo                   = 1149
	ID_lexvo_arl                   = 1150
	ID_lexvo_kdk                   = 1151
	ID_lexvo_suz                   = 1152
	ID_lexvo_pir                   = 1153
	ID_lexvo_khv                   = 1154
	ID_lexvo_bzq                   = 1155
	ID_lexvo_aud                   = 1156
	ID_lexvo_yux                   = 1157
	ID_lexvo_mng                   = 1158
	ID_lexvo_bxa                   = 1159
	ID_lexvo_alw                   = 1160
	ID_lexvo_ybe                   = 1161
	ID_lexvo_myy                   = 1162
	ID_lexvo_bgt                   = 1163
	ID_lexvo_tjg                   = 1164
	ID_lexvo_tmu                   = 1165
	ID_lexvo_bbr                   = 1166
	ID_lexvo_nym                   = 1167
	ID_lexvo_kee                   = 1168
	ID_lexvo_alk                   = 1169
	ID_lexvo_xsl                   = 1170
	ID_lexvo_kmb                   = 1171
	ID_lexvo_txh                   = 1172
	ID_lexvo_ata                   = 1173
	ID_lexvo_nha                   = 1174
	ID_lexvo_hoc                   = 1175
	ID_lexvo_urz                   = 1176
	ID_lexvo_kxa                   = 1177
	ID_lexvo_pnw                   = 1178
	ID_lexvo_mat                   = 1179
	ID_lexvo_ckx                   = 1180
	ID_lexvo_dua                   = 1181
	ID_lexvo_mzq                   = 1182
	ID_lexvo_brt                   = 1183
	ID_lexvo_pmf                   = 1184
	ID_lexvo_byn                   = 1185
	ID_lexvo_ilb                   = 1186
	ID_lexvo_dav                   = 1187
	ID_lexvo_kwd                   = 1188
	ID_lexvo_kdu                   = 1189
	ID_lexvo_cri                   = 1190
	ID_lexvo_dtd                   = 1191
	ID_lexvo_ssb                   = 1192
	ID_lexvo_cod                   = 1193
	ID_lexvo_mwr                   = 1194
	ID_lexvo_kbc                   = 1195
	ID_lexvo_agt                   = 1196
	ID_lexvo_kxo                   = 1197
	ID_lexvo_cpn                   = 1198
	ID_lexvo_aoa                   = 1199
	ID_lexvo_tmh                   = 1200
	ID_lexvo_ega                   = 1201
	ID_lexvo_dsn                   = 1202
	ID_lexvo_tpf                   = 1203
	ID_lexvo_bsn                   = 1204
	ID_lexvo_ski                   = 1205
	ID_lexvo_gor                   = 1206
	ID_lexvo_nxg                   = 1207
	ID_lexvo_bil                   = 1208
	ID_lexvo_iso                   = 1209
	ID_lexvo_bec                   = 1210
	ID_lexvo_obt                   = 1211
	ID_lexvo_anw                   = 1212
	ID_lexvo_yui                   = 1213
	ID_lexvo_hei                   = 1214
	ID_lexvo_juc                   = 1215
	ID_lexvo_tte                   = 1216
	ID_lexvo_alp                   = 1217
	ID_lexvo_eto                   = 1218
	ID_lexvo_alr                   = 1219
	ID_lexvo_qum                   = 1220
	ID_lexvo_slc                   = 1221
	ID_lexvo_mzs                   = 1222
	ID_lexvo_crw                   = 1223
	ID_lexvo_adw                   = 1224
	ID_lexvo_kiy                   = 1225
	ID_lexvo_sjm                   = 1226
	ID_lexvo_aui                   = 1227
	ID_lexvo_rej                   = 1228
	ID_lexvo_pri                   = 1229
	ID_lexvo_wao                   = 1230
	ID_lexvo_shy                   = 1231
	ID_lexvo_lil                   = 1232
	ID_lexvo_mqj                   = 1233
	ID_lexvo_kyl                   = 1234
	ID_lexvo_ktw                   = 1235
	ID_lexvo_bkn                   = 1236
	ID_lexvo_xaw                   = 1237
	ID_lexvo_bin                   = 1238
	ID_lexvo_api                   = 1239
	ID_lexvo_jhi                   = 1240
	ID_lexvo_peh                   = 1241
	ID_lexvo_mto                   = 1242
	ID_lexvo_uvl                   = 1243
	ID_lexvo_mfa                   = 1244
	ID_lexvo_gvf                   = 1245
	ID_lexvo_tcb                   = 1246
	ID_lexvo_pej                   = 1247
	ID_lexvo_jax                   = 1248
	ID_lexvo_sjo                   = 1249
	ID_lexvo_mbb                   = 1250
	ID_lexvo_kcg                   = 1251
	ID_lexvo_aak                   = 1252
	ID_lexvo_ljl                   = 1253
	ID_lexvo_mgr                   = 1254
	ID_lexvo_bdq                   = 1255
	ID_lexvo_kau                   = 1256
	ID_lexvo_tpy                   = 1257
	ID_lexvo_sgc                   = 1258
	ID_lexvo_tbi                   = 1259
	ID_lexvo_ani                   = 1260
	ID_lexvo_pre                   = 1261
	ID_lexvo_hvc                   = 1262
	ID_lexvo_lun                   = 1263
	ID_lexvo_gfk                   = 1264
	ID_lexvo_yka                   = 1265
	ID_lexvo_njo                   = 1266
	ID_lexvo_wiv                   = 1267
	ID_lexvo_gbu                   = 1268
	ID_lexvo_kuu                   = 1269
	ID_lexvo_cub                   = 1270
	ID_lexvo_tnc                   = 1271
	ID_lexvo_mid                   = 1272
	ID_lexvo_abm                   = 1273
	ID_lexvo_cbv                   = 1274
	ID_lexvo_are                   = 1275
	ID_lexvo_hss                   = 1276
	ID_lexvo_amg                   = 1277
	ID_lexvo_yug                   = 1278
	ID_lexvo_asu                   = 1279
	ID_lexvo_ewo                   = 1280
	ID_lexvo_ofu                   = 1281
	ID_lexvo_mog                   = 1282
	ID_lexvo_aum                   = 1283
	ID_lexvo_btu                   = 1284
	ID_lexvo_yuc                   = 1285
	ID_lexvo_bgc                   = 1286
	ID_lexvo_skd                   = 1287
	ID_lexvo_tri                   = 1288
	ID_lexvo_snn                   = 1289
	ID_lexvo_asn                   = 1290
	ID_lexvo_ppk                   = 1291
	ID_lexvo_apu                   = 1292
	ID_lexvo_mki                   = 1293
	ID_lexvo_hup                   = 1294
	ID_lexvo_lif                   = 1295
	ID_lexvo_xkn                   = 1296
	ID_lexvo_sub                   = 1297
	ID_lexvo_inm                   = 1298
	ID_lexvo_ayz                   = 1299
	ID_lexvo_guz                   = 1300
	ID_lexvo_ado                   = 1301
	ID_lexvo_teg                   = 1302
	ID_lexvo_bfs                   = 1303
	ID_lexvo_cks                   = 1304
	ID_lexvo_sri                   = 1305
	ID_lexvo_nee                   = 1306
	ID_lexvo_abn                   = 1307
	ID_lexvo_cjm                   = 1308
	ID_lexvo_mna                   = 1309
	ID_lexvo_kne                   = 1310
	ID_lexvo_vmy                   = 1311
	ID_lexvo_cin                   = 1312
	ID_lexvo_pos                   = 1313
	ID_lexvo_atw                   = 1314
	ID_lexvo_spp                   = 1315
	ID_lexvo_guh                   = 1316
	ID_lexvo_tew                   = 1317
	ID_lexvo_tft                   = 1318
	ID_lexvo_lmy                   = 1319
	ID_lexvo_hmv                   = 1320
	ID_lexvo_hmo                   = 1321
	ID_lexvo_sey                   = 1322
	ID_lexvo_mwv                   = 1323
	ID_lexvo_pmw                   = 1324
	ID_lexvo_txn                   = 1325
	ID_lexvo_wes                   = 1326
	ID_lexvo_squ                   = 1327
	ID_lexvo_kic                   = 1328
	ID_lexvo_rel                   = 1329
	ID_lexvo_too                   = 1330
	ID_lexvo_tnt                   = 1331
	ID_lexvo_mva                   = 1332
	ID_lexvo_xlc                   = 1333
	ID_lexvo_ncg                   = 1334
	ID_lexvo_tna                   = 1335
	ID_lexvo_con                   = 1336
	ID_lexvo_cul                   = 1337
	ID_lexvo_mjk                   = 1338
	ID_lexvo_max                   = 1339
	ID_lexvo_gvc                   = 1340
	ID_lexvo_xco                   = 1341
	ID_lexvo_piv                   = 1342
	ID_lexvo_dru                   = 1343
	ID_lexvo_kqn                   = 1344
	ID_lexvo_trf                   = 1345
	ID_lexvo_zca                   = 1346
	ID_lexvo_pga                   = 1347
	ID_lexvo_klq                   = 1348
	ID_lexvo_nyn                   = 1349
	ID_lexvo_coe                   = 1350
	ID_lexvo_blk                   = 1351
	ID_lexvo_sdh                   = 1352
	ID_lexvo_mfp                   = 1353
	ID_lexvo_ark                   = 1354
	ID_lexvo_isk                   = 1355
	ID_lexvo_bhj                   = 1356
	ID_lexvo_kxb                   = 1357
	ID_lexvo_otd                   = 1358
	ID_lexvo_ntw                   = 1359
	ID_lexvo_bfq                   = 1360
	ID_lexvo_wic                   = 1361
	ID_lexvo_tun                   = 1362
	ID_lexvo_mro                   = 1363
	ID_lexvo_tog                   = 1364
	ID_lexvo_yku                   = 1365
	ID_lexvo_smw                   = 1366
	ID_lexvo_sgw                   = 1367
	ID_lexvo_gqn                   = 1368
	ID_lexvo_asa                   = 1369
	ID_lexvo_txm                   = 1370
	ID_lexvo_stg                   = 1371
	ID_lexvo_mjm                   = 1372
	ID_lexvo_bky                   = 1373
	ID_lexvo_kry                   = 1374
	ID_lexvo_rmm                   = 1375
	ID_lexvo_xke                   = 1376
	ID_lexvo_zpf                   = 1377
	ID_lexvo_lgk                   = 1378
	ID_lexvo_bnd                   = 1379
	ID_lexvo_rak                   = 1380
	ID_lexvo_kpj                   = 1381
	ID_lexvo_alc                   = 1382
	ID_lexvo_nmw                   = 1383
	ID_lexvo_kqf                   = 1384
	ID_lexvo_nmc                   = 1385
	ID_lexvo_gnn                   = 1386
	ID_lexvo_ges                   = 1387
	ID_lexvo_mkf                   = 1388
	ID_lexvo_kem                   = 1389
	ID_lexvo_mpn                   = 1390
	ID_lexvo_vka                   = 1391
	ID_lexvo_abb                   = 1392
	ID_lexvo_xqt                   = 1393
	ID_lexvo_bap                   = 1394
	ID_lexvo_caf                   = 1395
	ID_lexvo_mqy                   = 1396
	ID_lexvo_kgk                   = 1397
	ID_lexvo_slu                   = 1398
	ID_lexvo_teu                   = 1399
	ID_lexvo_apt                   = 1400
	ID_lexvo_myh                   = 1401
	ID_lexvo_yak                   = 1402
	ID_lexvo_cng                   = 1403
	ID_lexvo_twu                   = 1404
	ID_lexvo_gae                   = 1405
	ID_lexvo_moz                   = 1406
	ID_lexvo_ret                   = 1407
	ID_lexvo_zae                   = 1408
	ID_lexvo_hux                   = 1409
	ID_lexvo_puw                   = 1410
	ID_lexvo_och                   = 1411
	ID_lexvo_crz                   = 1412
	ID_lexvo_zpl                   = 1413
	ID_lexvo_nho                   = 1414
	ID_lexvo_bhg                   = 1415
	ID_lexvo_svs                   = 1416
	ID_lexvo_xeu                   = 1417
	ID_lexvo_mbp                   = 1418
	ID_lexvo_hot                   = 1419
	ID_lexvo_lex                   = 1420
	ID_lexvo_vbb                   = 1421
	ID_lexvo_bxf                   = 1422
	ID_lexvo_aqm                   = 1423
	ID_lexvo_kfy                   = 1424
	ID_lexvo_oin                   = 1425
	ID_lexvo_sve                   = 1426
	ID_lexvo_doz                   = 1427
	ID_lexvo_imr                   = 1428
	ID_lexvo_ygr                   = 1429
	ID_lexvo_tsx                   = 1430
	ID_lexvo_wiy                   = 1431
	ID_lexvo_bkq                   = 1432
	ID_lexvo_ser                   = 1433
	ID_lexvo_sly                   = 1434
	ID_lexvo_psy                   = 1435
	ID_lexvo_til                   = 1436
	ID_lexvo_cso                   = 1437
	ID_lexvo_erg                   = 1438
	ID_lexvo_spe                   = 1439
	ID_lexvo_kip                   = 1440
	ID_lexvo_kwi                   = 1441
	ID_lexvo_kzx                   = 1442
	ID_lexvo_byd                   = 1443
	ID_lexvo_gur                   = 1444
	ID_lexvo_eya                   = 1445
	ID_lexvo_bcf                   = 1446
	ID_lexvo_tke                   = 1447
	ID_lexvo_srh                   = 1448
	ID_lexvo_bkz                   = 1449
	ID_lexvo_bnf                   = 1450
	ID_lexvo_bjc                   = 1451
	ID_lexvo_kgm                   = 1452
	ID_lexvo_pek                   = 1453
	ID_lexvo_met                   = 1454
	ID_lexvo_lvk                   = 1455
	ID_lexvo_goi                   = 1456
	ID_lexvo_nsz                   = 1457
	ID_lexvo_lis                   = 1458
	ID_lexvo_jbk                   = 1459
	ID_lexvo_tvm                   = 1460
	ID_lexvo_cwd                   = 1461
	ID_lexvo_der                   = 1462
	ID_lexvo_kxz                   = 1463
	ID_lexvo_lmb                   = 1464
	ID_lexvo_aap                   = 1465
	ID_lexvo_agw                   = 1466
	ID_lexvo_dyo                   = 1467
	ID_lexvo_jaa                   = 1468
	ID_lexvo_asx                   = 1469
	ID_lexvo_kbw                   = 1470
	ID_lexvo_muz                   = 1471
	ID_lexvo_trw                   = 1472
	ID_lexvo_thp                   = 1473
	ID_lexvo_dij                   = 1474
	ID_lexvo_pij                   = 1475
	ID_lexvo_kde                   = 1476
	ID_lexvo_vic                   = 1477
	ID_lexvo_bsu                   = 1478
	ID_lexvo_gob                   = 1479
	ID_lexvo_dby                   = 1480
	ID_lexvo_mxi                   = 1481
	ID_lexvo_yrl                   = 1482
	ID_lexvo_mfy                   = 1483
	ID_lexvo_aif                   = 1484
	ID_lexvo_xml                   = 1485
	ID_lexvo_bca                   = 1486
	ID_lexvo_bna                   = 1487
	ID_lexvo_mgq                   = 1488
	ID_lexvo_npl                   = 1489
	ID_lexvo_idi                   = 1490
	ID_lexvo_wms                   = 1491
	ID_lexvo_nyo                   = 1492
	ID_lexvo_kzf                   = 1493
	ID_lexvo_hmt                   = 1494
	ID_lexvo_kji                   = 1495
	ID_lexvo_jay                   = 1496
	ID_lexvo_dmr                   = 1497
	ID_lexvo_mhu                   = 1498
	ID_lexvo_hms                   = 1499
	ID_lexvo_mil                   = 1500
	ID_lexvo_rug                   = 1501
	ID_lexvo_arx                   = 1502
	ID_lexvo_ohu                   = 1503
	ID_lexvo_rjs                   = 1504
	ID_lexvo_hrc                   = 1505
	ID_lexvo_elk                   = 1506
	ID_lexvo_emw                   = 1507
	ID_lexvo_aig                   = 1508
	ID_lexvo_xkg                   = 1509
	ID_lexvo_bji                   = 1510
	ID_lexvo_osi                   = 1511
	ID_lexvo_btn                   = 1512
	ID_lexvo_umc                   = 1513
	ID_lexvo_kmj                   = 1514
	ID_lexvo_bhi                   = 1515
	ID_lexvo_slp                   = 1516
	ID_lexvo_hto                   = 1517
	ID_lexvo_wet                   = 1518
	ID_lexvo_gal                   = 1519
	ID_lexvo_mhi                   = 1520
	ID_lexvo_ifb                   = 1521
	ID_lexvo_hna                   = 1522
	ID_lexvo_kwf                   = 1523
	ID_lexvo_kii                   = 1524
	ID_lexvo_tro                   = 1525
	ID_lexvo_kfa                   = 1526
	ID_lexvo_ram                   = 1527
	ID_lexvo_brx                   = 1528
	ID_lexvo_lbw                   = 1529
	ID_lexvo_lra                   = 1530
	ID_lexvo_szv                   = 1531
	ID_lexvo_wed                   = 1532
	ID_lexvo_mup                   = 1533
	ID_lexvo_jka                   = 1534
	ID_lexvo_esi                   = 1535
	ID_lexvo_kpc                   = 1536
	ID_lexvo_mnd                   = 1537
	ID_lexvo_kju                   = 1538
	ID_lexvo_oty                   = 1539
	ID_lexvo_mkj                   = 1540
	ID_lexvo_kdj                   = 1541
	ID_lexvo_pbb                   = 1542
	ID_lexvo_agg                   = 1543
	ID_lexvo_sjk                   = 1544
	ID_lexvo_ghs                   = 1545
	ID_lexvo_dhv                   = 1546
	ID_lexvo_rir                   = 1547
	ID_lexvo_aoc                   = 1548
	ID_lexvo_ixl                   = 1549
	ID_lexvo_tea                   = 1550
	ID_lexvo_ppi                   = 1551
	ID_lexvo_gvj                   = 1552
	ID_lexvo_wru                   = 1553
	ID_lexvo_bpq                   = 1554
	ID_lexvo_mxz                   = 1555
	ID_lexvo_wmt                   = 1556
	ID_lexvo_stf                   = 1557
	ID_lexvo_tlc                   = 1558
	ID_lexvo_aey                   = 1559
	ID_lexvo_vme                   = 1560
	ID_lexvo_cog                   = 1561
	ID_lexvo_cyb                   = 1562
	ID_lexvo_kni                   = 1563
	ID_lexvo_bci                   = 1564
	ID_lexvo_bcd                   = 1565
	ID_lexvo_wat                   = 1566
	ID_lexvo_auw                   = 1567
	ID_lexvo_pnr                   = 1568
	ID_lexvo_khc                   = 1569
	ID_lexvo_xla                   = 1570
	ID_lexvo_tpa                   = 1571
	ID_lexvo_bue                   = 1572
	ID_lexvo_cap                   = 1573
	ID_lexvo_med                   = 1574
	ID_lexvo_msw                   = 1575
	ID_lexvo_lhu                   = 1576
	ID_lexvo_agd                   = 1577
	ID_lexvo_kts                   = 1578
	ID_lexvo_jmx                   = 1579
	ID_lexvo_auq                   = 1580
	ID_lexvo_xan                   = 1581
	ID_lexvo_iyo                   = 1582
	ID_lexvo_tbj                   = 1583
	ID_lexvo_hao                   = 1584
	ID_lexvo_nwr                   = 1585
	ID_lexvo_pnz                   = 1586
	ID_lexvo_cbt                   = 1587
	ID_lexvo_ddi                   = 1588
	ID_lexvo_mxb                   = 1589
	ID_lexvo_mne                   = 1590
	ID_lexvo_bqb                   = 1591
	ID_lexvo_wod                   = 1592
	ID_lexvo_xrw                   = 1593
	ID_lexvo_txs                   = 1594
	ID_lexvo_kpr                   = 1595
	ID_lexvo_wow                   = 1596
	ID_lexvo_res                   = 1597
	ID_lexvo_ojv                   = 1598
	ID_lexvo_reb                   = 1599
	ID_lexvo_yub                   = 1600
	ID_lexvo_dji                   = 1601
	ID_lexvo_khz                   = 1602
	ID_lexvo_guo                   = 1603
	ID_lexvo_ppu                   = 1604
	ID_lexvo_cbg                   = 1605
	ID_lexvo_oyd                   = 1606
	ID_lexvo_mrx                   = 1607
	ID_lexvo_nup                   = 1608
	ID_lexvo_boq                   = 1609
	ID_lexvo_gnd                   = 1610
	ID_lexvo_kxv                   = 1611
	ID_lexvo_tiv                   = 1612
	ID_lexvo_bzk                   = 1613
	ID_lexvo_tow                   = 1614
	ID_lexvo_hdy                   = 1615
	ID_lexvo_pbn                   = 1616
	ID_lexvo_kyt                   = 1617
	ID_lexvo_dva                   = 1618
	ID_lexvo_arh                   = 1619
	ID_lexvo_xmm                   = 1620
	ID_lexvo_bcn                   = 1621
	ID_lexvo_srw                   = 1622
	ID_lexvo_wne                   = 1623
	ID_lexvo_abp                   = 1624
	ID_lexvo_mrt                   = 1625
	ID_lexvo_dia                   = 1626
	ID_lexvo_irn                   = 1627
	ID_lexvo_vko                   = 1628
	ID_lexvo_mtp                   = 1629
	ID_lexvo_uln                   = 1630
	ID_lexvo_dmu                   = 1631
	ID_lexvo_cbi                   = 1632
	ID_lexvo_nll                   = 1633
	ID_lexvo_moc                   = 1634
	ID_lexvo_sry                   = 1635
	ID_lexvo_kti                   = 1636
	ID_lexvo_mvd                   = 1637
	ID_lexvo_tjs                   = 1638
	ID_lexvo_sda                   = 1639
	ID_lexvo_kek                   = 1640
	ID_lexvo_kul                   = 1641
	ID_lexvo_mod                   = 1642
	ID_lexvo_wal                   = 1643
	ID_lexvo_urt                   = 1644
	ID_lexvo_lag                   = 1645
	ID_lexvo_lti                   = 1646
	ID_lexvo_bmh                   = 1647
	ID_lexvo_bja                   = 1648
	ID_lexvo_mfb                   = 1649
	ID_lexvo_ttu                   = 1650
	ID_lexvo_nam                   = 1651
	ID_lexvo_hch                   = 1652
	ID_lexvo_hnd                   = 1653
	ID_lexvo_cji                   = 1654
	ID_lexvo_hhy                   = 1655
	ID_lexvo_www                   = 1656
	ID_lexvo_erk                   = 1657
	ID_lexvo_lsr                   = 1658
	ID_lexvo_sre                   = 1659
	ID_lexvo_apb                   = 1660
	ID_lexvo_aso                   = 1661
	ID_lexvo_bij                   = 1662
	ID_lexvo_abt                   = 1663
	ID_lexvo_kag                   = 1664
	ID_lexvo_spx                   = 1665
	ID_lexvo_not                   = 1666
	ID_lexvo_cta                   = 1667
	ID_lexvo_tdv                   = 1668
	ID_lexvo_kzp                   = 1669
	ID_lexvo_vah                   = 1670
	ID_lexvo_ktz                   = 1671
	ID_lexvo_zgr                   = 1672
	ID_lexvo_bgz                   = 1673
	ID_lexvo_ikw                   = 1674
	ID_lexvo_omn                   = 1675
	ID_lexvo_aah                   = 1676
	ID_lexvo_slm                   = 1677
	ID_lexvo_pab                   = 1678
	ID_lexvo_mgv                   = 1679
	ID_lexvo_tjm                   = 1680
	ID_lexvo_mbr                   = 1681
	ID_lexvo_heh                   = 1682
	ID_lexvo_sbh                   = 1683
	ID_lexvo_omx                   = 1684
	ID_lexvo_cuc                   = 1685
	ID_lexvo_kqw                   = 1686
	ID_lexvo_kio                   = 1687
	ID_lexvo_pil                   = 1688
	ID_lexvo_plv                   = 1689
	ID_lexvo_wrs                   = 1690
	ID_lexvo_shv                   = 1691
	ID_lexvo_tlr                   = 1692
	ID_lexvo_xbm                   = 1693
	ID_lexvo_crm                   = 1694
	ID_lexvo_kyq                   = 1695
	ID_lexvo_yii                   = 1696
	ID_lexvo_lue                   = 1697
	ID_lexvo_des                   = 1698
	ID_lexvo_gby                   = 1699
	ID_lexvo_scb                   = 1700
	ID_lexvo_blw                   = 1701
	ID_lexvo_ong                   = 1702
	ID_lexvo_kma                   = 1703
	ID_lexvo_piz                   = 1704
	ID_lexvo_bea                   = 1705
	ID_lexvo_kod                   = 1706
	ID_lexvo_xvs                   = 1707
	ID_lexvo_huu                   = 1708
	ID_lexvo_cav                   = 1709
	ID_lexvo_miz                   = 1710
	ID_lexvo_rof                   = 1711
	ID_lexvo_bcj                   = 1712
	ID_lexvo_iai                   = 1713
	ID_lexvo_tmq                   = 1714
	ID_lexvo_tcc                   = 1715
	ID_lexvo_kvr                   = 1716
	ID_lexvo_bdo                   = 1717
	ID_lexvo_zen                   = 1718
	ID_lexvo_dto                   = 1719
	ID_lexvo_unz                   = 1720
	ID_lexvo_mox                   = 1721
	ID_lexvo_hov                   = 1722
	ID_lexvo_pio                   = 1723
	ID_lexvo_nak                   = 1724
	ID_lexvo_sax                   = 1725
	ID_lexvo_bts                   = 1726
	ID_lexvo_bpg                   = 1727
	ID_lexvo_bwd                   = 1728
	ID_lexvo_kay                   = 1729
	ID_lexvo_pud                   = 1730
	ID_lexvo_bch                   = 1731
	ID_lexvo_aes                   = 1732
	ID_lexvo_sny                   = 1733
	ID_lexvo_oca                   = 1734
	ID_lexvo_niy                   = 1735
	ID_lexvo_num                   = 1736
	ID_lexvo_meo                   = 1737
	ID_lexvo_aai                   = 1738
	ID_lexvo_tav                   = 1739
	ID_lexvo_mda                   = 1740
	ID_lexvo_bny                   = 1741
	ID_lexvo_ylg                   = 1742
	ID_lexvo_xyy                   = 1743
	ID_lexvo_mnt                   = 1744
	ID_lexvo_khs                   = 1745
	ID_lexvo_ule                   = 1746
	ID_lexvo_cax                   = 1747
	ID_lexvo_skr                   = 1748
	ID_lexvo_att                   = 1749
	ID_lexvo_kmn                   = 1750
	ID_lexvo_jic                   = 1751
	ID_lexvo_ttv                   = 1752
	ID_lexvo_tuq                   = 1753
	ID_lexvo_sso                   = 1754
	ID_lexvo_aty                   = 1755
	ID_lexvo_quv                   = 1756
	ID_lexvo_dai                   = 1757
	ID_lexvo_gbr                   = 1758
	ID_lexvo_ibl                   = 1759
	ID_lexvo_enf                   = 1760
	ID_lexvo_pui                   = 1761
	ID_lexvo_cmi                   = 1762
	ID_lexvo_drn                   = 1763
	ID_lexvo_duk                   = 1764
	ID_lexvo_lib                   = 1765
	ID_lexvo_ndc                   = 1766
	ID_lexvo_aas                   = 1767
	ID_lexvo_mvn                   = 1768
	ID_lexvo_nab                   = 1769
	ID_lexvo_tuf                   = 1770
	ID_lexvo_grw                   = 1771
	ID_lexvo_csz                   = 1772
	ID_lexvo_ebo                   = 1773
	ID_lexvo_ues                   = 1774
	ID_lexvo_lcm                   = 1775
	ID_lexvo_blf                   = 1776
	ID_lexvo_chw                   = 1777
	ID_lexvo_cwg                   = 1778
	ID_lexvo_rah                   = 1779
	ID_lexvo_djd                   = 1780
	ID_lexvo_btz                   = 1781
	ID_lexvo_ojw                   = 1782
	ID_lexvo_tdt                   = 1783
	ID_lexvo_bbd                   = 1784
	ID_lexvo_bim                   = 1785
	ID_lexvo_kka                   = 1786
	ID_lexvo_myw                   = 1787
	ID_lexvo_rma                   = 1788
	ID_lexvo_poo                   = 1789
	ID_lexvo_puj                   = 1790
	ID_lexvo_kos                   = 1791
	ID_lexvo_srs                   = 1792
	ID_lexvo_big                   = 1793
	ID_lexvo_pay                   = 1794
	ID_lexvo_akm                   = 1795
	ID_lexvo_bob                   = 1796
	ID_lexvo_sgh                   = 1797
	ID_lexvo_lgt                   = 1798
	ID_lexvo_otn                   = 1799
	ID_lexvo_uiv                   = 1800
	ID_lexvo_xwa                   = 1801
	ID_lexvo_tkn                   = 1802
	ID_lexvo_bhl                   = 1803
	ID_lexvo_szp                   = 1804
	ID_lexvo_lah                   = 1805
	ID_lexvo_mmn                   = 1806
	ID_lexvo_kzi                   = 1807
	ID_lexvo_rac                   = 1808
	ID_lexvo_koa                   = 1809
	ID_lexvo_bmc                   = 1810
	ID_lexvo_los                   = 1811
	ID_lexvo_fuy                   = 1812
	ID_lexvo_khq                   = 1813
	ID_lexvo_uur                   = 1814
	ID_lexvo_pln                   = 1815
	ID_lexvo_btm                   = 1816
	ID_lexvo_bjk                   = 1817
	ID_lexvo_djm                   = 1818
	ID_lexvo_yml                   = 1819
	ID_lexvo_ywr                   = 1820
	ID_lexvo_xdm                   = 1821
	ID_lexvo_sja                   = 1822
	ID_lexvo_kak                   = 1823
	ID_lexvo_kbk                   = 1824
	ID_lexvo_dev                   = 1825
	ID_lexvo_emi                   = 1826
	ID_lexvo_buk                   = 1827
	ID_lexvo_bwf                   = 1828
	ID_lexvo_eme                   = 1829
	ID_lexvo_myu                   = 1830
	ID_lexvo_avt                   = 1831
	ID_lexvo_cao                   = 1832
	ID_lexvo_ign                   = 1833
	ID_lexvo_auu                   = 1834
	ID_lexvo_ykm                   = 1835
	ID_lexvo_men                   = 1836
	ID_lexvo_hya                   = 1837
	ID_lexvo_pka                   = 1838
	ID_lexvo_hea                   = 1839
	ID_lexvo_nux                   = 1840
	ID_lexvo_onw                   = 1841
	ID_lexvo_bkh                   = 1842
	ID_lexvo_kbt                   = 1843
	ID_lexvo_amk                   = 1844
	ID_lexvo_drl                   = 1845
	ID_lexvo_kjl                   = 1846
	ID_lexvo_bgv                   = 1847
	ID_lexvo_kdd                   = 1848
	ID_lexvo_bjt                   = 1849
	ID_lexvo_rai                   = 1850
	ID_lexvo_fun                   = 1851
	ID_lexvo_loj                   = 1852
	ID_lexvo_shk                   = 1853
	ID_lexvo_mmx                   = 1854
	ID_lexvo_svb                   = 1855
	ID_lexvo_nol                   = 1856
	ID_lexvo_alu                   = 1857
	ID_lexvo_trc                   = 1858
	ID_lexvo_kpx                   = 1859
	ID_lexvo_woc                   = 1860
	ID_lexvo_bhv                   = 1861
	ID_lexvo_khl                   = 1862
	ID_lexvo_haj                   = 1863
	ID_lexvo_ppn                   = 1864
	ID_lexvo_akg                   = 1865
	ID_lexvo_mty                   = 1866
	ID_lexvo_ncn                   = 1867
	ID_lexvo_fwa                   = 1868
	ID_lexvo_amq                   = 1869
	ID_lexvo_nxr                   = 1870
	ID_lexvo_ksb                   = 1871
	ID_lexvo_vkl                   = 1872
	ID_lexvo_bzd                   = 1873
	ID_lexvo_gri                   = 1874
	ID_lexvo_swp                   = 1875
	ID_lexvo_dgg                   = 1876
	ID_lexvo_nss                   = 1877
	ID_lexvo_cbc                   = 1878
	ID_lexvo_zkt                   = 1879
	ID_lexvo_dcc                   = 1880
	ID_lexvo_bjl                   = 1881
	ID_lexvo_aun                   = 1882
	ID_lexvo_pit                   = 1883
	ID_lexvo_iow                   = 1884
	ID_lexvo_fos                   = 1885
	ID_lexvo_wnk                   = 1886
	ID_lexvo_mcr                   = 1887
	ID_lexvo_pse                   = 1888
	ID_lexvo_cto                   = 1889
	ID_lexvo_rea                   = 1890
	ID_lexvo_kym                   = 1891
	ID_lexvo_ajg                   = 1892
	ID_lexvo_nem                   = 1893
	ID_lexvo_lmk                   = 1894
	ID_lexvo_htu                   = 1895
	ID_lexvo_mnp                   = 1896
	ID_lexvo_itx                   = 1897
	ID_lexvo_rro                   = 1898
	ID_lexvo_mot                   = 1899
	ID_lexvo_ttk                   = 1900
	ID_lexvo_kue                   = 1901
	ID_lexvo_lay                   = 1902
	ID_lexvo_dmw                   = 1903
	ID_lexvo_bzg                   = 1904
	ID_lexvo_ijc                   = 1905
	ID_lexvo_yer                   = 1906
	ID_lexvo_kvo                   = 1907
	ID_lexvo_ade                   = 1908
	ID_lexvo_huj                   = 1909
	ID_lexvo_uok                   = 1910
	ID_lexvo_ptu                   = 1911
	ID_lexvo_tdc                   = 1912
	ID_lexvo_lbx                   = 1913
	ID_lexvo_mtt                   = 1914
	ID_lexvo_upv                   = 1915
	ID_lexvo_val                   = 1916
	ID_lexvo_bei                   = 1917
	ID_lexvo_tfr                   = 1918
	ID_lexvo_rgr                   = 1919
	ID_lexvo_kxn                   = 1920
	ID_lexvo_bty                   = 1921
	ID_lexvo_cld                   = 1922
	ID_lexvo_ral                   = 1923
	ID_lexvo_buy                   = 1924
	ID_lexvo_hne                   = 1925
	ID_lexvo_unr                   = 1926
	ID_lexvo_asi                   = 1927
	ID_lexvo_kei                   = 1928
	ID_lexvo_pip                   = 1929
	ID_lexvo_txe                   = 1930
	ID_lexvo_pib                   = 1931
	ID_lexvo_mkv                   = 1932
	ID_lexvo_wca                   = 1933
	ID_lexvo_guu                   = 1934
	ID_lexvo_yup                   = 1935
	ID_lexvo_due                   = 1936
	ID_lexvo_dbn                   = 1937
	ID_lexvo_bkr                   = 1938
	ID_lexvo_ksf                   = 1939
	ID_lexvo_czn                   = 1940
	ID_lexvo_tbf                   = 1941
	ID_lexvo_beg                   = 1942
	ID_lexvo_mta                   = 1943
	ID_lexvo_kwh                   = 1944
	ID_lexvo_leh                   = 1945
	ID_lexvo_nmb                   = 1946
	ID_lexvo_mlp                   = 1947
	ID_lexvo_bco                   = 1948
	ID_lexvo_bdh                   = 1949
	ID_lexvo_xvn                   = 1950
	ID_lexvo_roe                   = 1951
	ID_lexvo_blz                   = 1952
	ID_lexvo_wab                   = 1953
	ID_lexvo_bzb                   = 1954
	ID_lexvo_man                   = 1955
	ID_lexvo_nfu                   = 1956
	ID_lexvo_xsi                   = 1957
	ID_lexvo_kmo                   = 1958
	ID_lexvo_acn                   = 1959
	ID_lexvo_prk                   = 1960
	ID_lexvo_vmg                   = 1961
	ID_lexvo_cam                   = 1962
	ID_lexvo_ese                   = 1963
	ID_lexvo_aua                   = 1964
	ID_lexvo_faf                   = 1965
	ID_lexvo_ass                   = 1966
	ID_lexvo_pww                   = 1967
	ID_lexvo_ppo                   = 1968
	ID_lexvo_bvy                   = 1969
	ID_lexvo_add                   = 1970
	ID_lexvo_snc                   = 1971
	ID_lexvo_alo                   = 1972
	ID_lexvo_eth                   = 1973
	ID_lexvo_roo                   = 1974
	ID_lexvo_bxe                   = 1975
	ID_lexvo_eno                   = 1976
	ID_lexvo_kog                   = 1977
	ID_lexvo_kbh                   = 1978
	ID_lexvo_lon                   = 1979
	ID_lexvo_ssg                   = 1980
	ID_lexvo_ddd                   = 1981
	ID_lexvo_kzh                   = 1982
	ID_lexvo_ape                   = 1983
	ID_lexvo_bln                   = 1984
	ID_lexvo_bdn                   = 1985
	ID_lexvo_boa                   = 1986
	ID_lexvo_bxg                   = 1987
	ID_lexvo_jig                   = 1988
	ID_lexvo_fut                   = 1989
	ID_lexvo_sbv                   = 1990
	ID_lexvo_tlk                   = 1991
	ID_lexvo_qyp                   = 1992
	ID_lexvo_brc                   = 1993
	ID_lexvo_smq                   = 1994
	ID_lexvo_gid                   = 1995
	ID_lexvo_mqn                   = 1996
	ID_lexvo_txa                   = 1997
	ID_lexvo_mxd                   = 1998
	ID_lexvo_pna                   = 1999
	ID_lexvo_tzn                   = 2000
	ID_lexvo_lgr                   = 2001
	ID_lexvo_zmi                   = 2002
	ID_lexvo_uta                   = 2003
	ID_lexvo_wkd                   = 2004
	ID_lexvo_aad                   = 2005
	ID_lexvo_mrh                   = 2006
	ID_lexvo_dim                   = 2007
	ID_lexvo_hgm                   = 2008
	ID_lexvo_sid                   = 2009
	ID_lexvo_tix                   = 2010
	ID_lexvo_gvl                   = 2011
	ID_lexvo_umo                   = 2012
	ID_lexvo_gay                   = 2013
	ID_lexvo_mpe                   = 2014
	ID_lexvo_sgt                   = 2015
	ID_lexvo_sam                   = 2016
	ID_lexvo_lam                   = 2017
	ID_lexvo_akt                   = 2018
	ID_lexvo_leu                   = 2019
	ID_lexvo_gum                   = 2020
	ID_lexvo_tve                   = 2021
	ID_lexvo_wig                   = 2022
	ID_lexvo_ddw                   = 2023
	ID_lexvo_ank                   = 2024
	ID_lexvo_acv                   = 2025
	ID_lexvo_agv                   = 2026
	ID_lexvo_wbv                   = 2027
	ID_lexvo_mok                   = 2028
	ID_lexvo_blb                   = 2029
	ID_lexvo_pia                   = 2030
	ID_lexvo_bhs                   = 2031
	ID_lexvo_boi                   = 2032
	ID_lexvo_gyd                   = 2033
	ID_lexvo_wad                   = 2034
	ID_lexvo_agf                   = 2035
	ID_lexvo_adz                   = 2036
	ID_lexvo_ktx                   = 2037
	ID_lexvo_dag                   = 2038
	ID_lexvo_scg                   = 2039
	ID_lexvo_amx                   = 2040
	ID_lexvo_cbd                   = 2041
	ID_lexvo_tip                   = 2042
	ID_lexvo_mui                   = 2043
	ID_lexvo_udj                   = 2044
	ID_lexvo_bum                   = 2045
	ID_lexvo_vmf                   = 2046
	ID_lexvo_tkw                   = 2047
	ID_lexvo_mxr                   = 2048
	ID_lexvo_bmr                   = 2049
	ID_lexvo_bdd                   = 2050
	ID_lexvo_mig                   = 2051
	ID_lexvo_utp                   = 2052
	ID_lexvo_bhq                   = 2053
	ID_lexvo_plh                   = 2054
	ID_lexvo_lml                   = 2055
	ID_lexvo_bnv                   = 2056
	ID_lexvo_bmk                   = 2057
	ID_lexvo_bld                   = 2058
	ID_lexvo_wuv                   = 2059
	ID_lexvo_mwe                   = 2060
	ID_lexvo_abr                   = 2061
	ID_lexvo_gcl                   = 2062
	ID_lexvo_kzu                   = 2063
	ID_lexvo_icr                   = 2064
	ID_lexvo_itb                   = 2065
	ID_lexvo_kkk                   = 2066
	ID_lexvo_pdo                   = 2067
	ID_lexvo_bfc                   = 2068
	ID_lexvo_mxm                   = 2069
	ID_lexvo_vnk                   = 2070
	ID_lexvo_rim                   = 2071
	ID_lexvo_bnz                   = 2072
	ID_lexvo_mwc                   = 2073
	ID_lexvo_sto                   = 2074
	ID_lexvo_rkb                   = 2075
	ID_lexvo_ndp                   = 2076
	ID_lexvo_bhw                   = 2077
	ID_lexvo_tlu                   = 2078
	ID_lexvo_tox                   = 2079
	ID_lexvo_itz                   = 2080
	ID_lexvo_miw                   = 2081
	ID_lexvo_emp                   = 2082
	ID_lexvo_nhu                   = 2083
	ID_lexvo_lub                   = 2084
	ID_lexvo_dil                   = 2085
	ID_lexvo_agn                   = 2086
	ID_lexvo_noe                   = 2087
	ID_lexvo_paq                   = 2088
	ID_lexvo_gdq                   = 2089
	ID_lexvo_bey                   = 2090
	ID_lexvo_wew                   = 2091
	ID_lexvo_xel                   = 2092
	ID_lexvo_bwu                   = 2093
	ID_lexvo_ars                   = 2094
	ID_lexvo_tct                   = 2095
	ID_lexvo_mpx                   = 2096
	ID_lexvo_duf                   = 2097
	ID_lexvo_noa                   = 2098
	ID_lexvo_frq                   = 2099
	ID_lexvo_mei                   = 2100
	ID_lexvo_sbf                   = 2101
	ID_lexvo_coj                   = 2102
	ID_lexvo_wah                   = 2103
	ID_lexvo_xpo                   = 2104
	ID_lexvo_gve                   = 2105
	ID_lexvo_rbb                   = 2106
	ID_lexvo_mhs                   = 2107
	ID_lexvo_dww                   = 2108
	ID_lexvo_ssq                   = 2109
	ID_lexvo_trn                   = 2110
	ID_lexvo_laj                   = 2111
	ID_lexvo_bhm                   = 2112
	ID_lexvo_skf                   = 2113
	ID_lexvo_mbc                   = 2114
	ID_lexvo_klx                   = 2115
	ID_lexvo_bnq                   = 2116
	ID_lexvo_plo                   = 2117
	ID_lexvo_tsj                   = 2118
	ID_lexvo_lew                   = 2119
	ID_lexvo_xib                   = 2120
	ID_lexvo_tbe                   = 2121
	ID_lexvo_knv                   = 2122
	ID_lexvo_bkt                   = 2123
	ID_lexvo_ilu                   = 2124
	ID_lexvo_nnh                   = 2125
	ID_lexvo_smp                   = 2126
	ID_lexvo_hoa                   = 2127
	ID_lexvo_kyz                   = 2128
	ID_lexvo_vai                   = 2129
	ID_lexvo_kpt                   = 2130
	ID_lexvo_gwd                   = 2131
	ID_lexvo_kjb                   = 2132
	ID_lexvo_bbn                   = 2133
	ID_lexvo_pwg                   = 2134
	ID_lexvo_ura                   = 2135
	ID_lexvo_awb                   = 2136
	ID_lexvo_aac                   = 2137
	ID_lexvo_bwa                   = 2138
	ID_lexvo_kvc                   = 2139
	ID_lexvo_brz                   = 2140
	ID_lexvo_ikt                   = 2141
	ID_lexvo_kck                   = 2142
	ID_lexvo_atu                   = 2143
	ID_lexvo_ssd                   = 2144
	ID_lexvo_plc                   = 2145
	ID_lexvo_asl                   = 2146
	ID_lexvo_ame                   = 2147
	ID_lexvo_bdp                   = 2148
	ID_lexvo_abs                   = 2149
	ID_lexvo_kpm                   = 2150
	ID_lexvo_mzp                   = 2151
	ID_lexvo_emn                   = 2152
	ID_lexvo_onn                   = 2153
	ID_lexvo_mvb                   = 2154
	ID_lexvo_kwe                   = 2155
	ID_lexvo_jod                   = 2156
	ID_lexvo_zmr                   = 2157
	ID_lexvo_nkr                   = 2158
	ID_lexvo_anm                   = 2159
	ID_lexvo_nsn                   = 2160
	ID_lexvo_xng                   = 2161
	ID_lexvo_aiw                   = 2162
	ID_lexvo_lgl                   = 2163
	ID_lexvo_ags                   = 2164
	ID_lexvo_mjg                   = 2165
	ID_lexvo_ttm                   = 2166
	ID_lexvo_avi                   = 2167
	ID_lexvo_lbb                   = 2168
	ID_lexvo_hmr                   = 2169
	ID_lexvo_mmq                   = 2170
	ID_lexvo_eky                   = 2171
	ID_lexvo_swu                   = 2172
	ID_lexvo_yuz                   = 2173
	ID_lexvo_wrg                   = 2174
	ID_lexvo_pma                   = 2175
	ID_lexvo_tbk                   = 2176
	ID_lexvo_fad                   = 2177
	ID_lexvo_djo                   = 2178
	ID_lexvo_sus                   = 2179
	ID_lexvo_ayd                   = 2180
	ID_lexvo_cof                   = 2181
	ID_lexvo_wro                   = 2182
	ID_lexvo_hih                   = 2183
	ID_lexvo_xem                   = 2184
	ID_lexvo_ubr                   = 2185
	ID_lexvo_amc                   = 2186
	ID_lexvo_skz                   = 2187
	ID_lexvo_prq                   = 2188
	ID_lexvo_mue                   = 2189
	ID_lexvo_xed                   = 2190
	ID_lexvo_ers                   = 2191
	ID_lexvo_ofo                   = 2192
	ID_lexvo_tgx                   = 2193
	ID_lexvo_pbh                   = 2194
	ID_lexvo_nil                   = 2195
	ID_lexvo_xmz                   = 2196
	ID_lexvo_bxh                   = 2197
	ID_lexvo_mjb                   = 2198
	ID_lexvo_jbt                   = 2199
	ID_lexvo_dob                   = 2200
	ID_lexvo_rin                   = 2201
	ID_lexvo_xct                   = 2202
	ID_lexvo_and                   = 2203
	ID_lexvo_bkj                   = 2204
	ID_lexvo_ulu                   = 2205
	ID_lexvo_ziw                   = 2206
	ID_lexvo_bzz                   = 2207
	ID_lexvo_ykn                   = 2208
	ID_lexvo_inb                   = 2209
	ID_lexvo_ved                   = 2210
	ID_lexvo_kje                   = 2211
	ID_lexvo_wlk                   = 2212
	ID_lexvo_ray                   = 2213
	ID_lexvo_trx                   = 2214
	ID_lexvo_jml                   = 2215
	ID_lexvo_aib                   = 2216
	ID_lexvo_tpj                   = 2217
	ID_lexvo_ctd                   = 2218
	ID_lexvo_kcn                   = 2219
	ID_lexvo_lev                   = 2220
	ID_lexvo_omc                   = 2221
	ID_lexvo_umi                   = 2222
	ID_lexvo_ate                   = 2223
	ID_lexvo_zne                   = 2224
	ID_lexvo_bra                   = 2225
	ID_lexvo_gva                   = 2226
	ID_lexvo_bag                   = 2227
	ID_lexvo_tan                   = 2228
	ID_lexvo_czk                   = 2229
	ID_lexvo_aot                   = 2230
	ID_lexvo_yey                   = 2231
	ID_lexvo_mbn                   = 2232
	ID_lexvo_tdi                   = 2233
	ID_lexvo_qun                   = 2234
	ID_lexvo_lmr                   = 2235
	ID_lexvo_kqb                   = 2236
	ID_lexvo_ysg                   = 2237
	ID_lexvo_nbp                   = 2238
	ID_lexvo_yrb                   = 2239
	ID_lexvo_ayh                   = 2240
	ID_lexinfo                     = 5
	ID_lexinfo_partOfSpeech        = 0
	ID_lexinfo_noun                = 1
	ID_lexinfo_verb                = 2
	ID_lexinfo_adjective           = 3
	ID_lexinfo_properNoun          = 4
	ID_lexinfo_adverb              = 5
	ID_lexinfo_interjection        = 6
	ID_lexinfo_Interjection        = 7
	ID_lexinfo_phraseologicalUnit  = 8
	ID_lexinfo_Prefix              = 9
	ID_lexinfo_prefix              = 10
	ID_lexinfo_proverb             = 11
	ID_lexinfo_pronunciation       = 12
	ID_lexinfo_suffix              = 13
	ID_lexinfo_Suffix              = 14
	ID_lexinfo_pronoun             = 15
	ID_lexinfo_Pronoun             = 16
	ID_lexinfo_preposition         = 17
	ID_lexinfo_Preposition         = 18
	ID_lexinfo_numeral             = 19
	ID_lexinfo_Numeral             = 20
	ID_lexinfo_conjunction         = 21
	ID_lexinfo_Conjunction         = 22
	ID_lexinfo_determiner          = 23
	ID_lexinfo_Determiner          = 24
	ID_lexinfo_symbol              = 25
	ID_lexinfo_Symbol              = 26
	ID_lexinfo_Number              = 27
	ID_lexinfo_infix               = 28
	ID_lexinfo_Infix               = 29
	ID_lexinfo_Particle            = 30
	ID_lexinfo_particle            = 31
	ID_lexinfo_Article             = 32
	ID_lexinfo_article             = 33
	ID_lexinfo_Affix               = 34
	ID_lexinfo_affix               = 35
	ID_lexinfo_postposition        = 36
	ID_lexinfo_Postposition        = 37
	ID_lexinfo_idiom               = 38
	ID_olia                        = 6
	ID_olia_hasNumber              = 0
	ID_olia_Singular               = 1
	ID_olia_hasCountability        = 2
	ID_olia_Countable              = 3
	ID_olia_Uncountable            = 4
	ID_olia_Plural                 = 5
	ID_olia_hasDegree              = 6
	ID_olia_Comparative            = 7
	ID_dcterms                     = 7
	ID_dcterms_language            = 0
	ID_lime                        = 8
	ID_lime_language               = 0
	ID_skos                        = 9
	ID_skos_definition             = 0
	ID_vartrans                    = 10
	ID_vartrans_lexicalRel         = 0
	ID_decomp                      = 11
	ID_dbetym                      = 12
	ID_rdfs                        = 13
	ID_synsem                      = 14
	ID_xs                          = 15
	ID_wikt                        = 16
)
