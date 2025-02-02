package useragent

var WIN_NT = [8]string{"5.1", "5.2", "6.0", "6.1", "6.2", "6.3", "6.4", "10.0"}

var WIN_BIT = [3]string{"Win64; x64", "WOW64", ""}

var MACINTOSH_VER = [19]string{
	"13_3_1", "12_6_5", "11_7_6", "10_15_7", "10_14_6", "10_13_6", "10_12_6", "10_11_6", "10_10_5",
	"10_9_5", "10_8_5", "10_7_5", "10_6_8", "10_5_8", "10_4_11", "10_3_9", "10_2_8", "10_1_5",
	"10_0_4",
}

var LINUX_VER = [13]string{
	"Ubuntu; Linux x86-64",
	"Ubuntu; Linux i686",
	"Ubuntu 24.04 LTS; Linux x86_64",
	"Ubuntu; Linux",
	"Debian; Linux",
	"Debian; Linux i686",
	"U; Linux armv5tejl",
	"Linux i686; nb-NO; rv:1.8.1.3",
	"FreeBSD amd64; rv:102.0",
	"FreeBSD amd64; rv:124.0",
	"U; FreeBSD amd64 3.99; rv:224.6",
	"FreeBSD amd64",
	"U; Linux i686",
}

var ANDROID_VER = [20]string{
	"9",
	"10",
	"11",
	"12",
	"12L",
	"13",
	"14",
	"14.53.0.40515",
	"14.53.0.40515",
	"14.52.0.40456",
	"14.50.0.40042",
	"14.51.0.40361",
	"14.44.0.39235",
	"15",
	"15.2.1",
	"15.3.1",
	"15.0",
	"15.1",
	"15.0.0.0",
	"16",
}

var HUAWEI = [35]string{
	// "HUAWEI U8825-1 Build/HuaweiU8825-1",
	// "HUAWEI G510-0200 Build/HuaweiG510-0200",
	// "HUAWEI G525-U00 Build/HuaweiG525-U00",
	// "HUAWEI G525-U00 Build/HuaweiG525-U00",
	// "HUAWEI U8950D Build/HuaweiU8950D",
	// "G7-L01 Build/HuaweiG7-L01",
	// "HUAWEI G700-U10 Build/HuaweiG700-U10",
	// "HUAWEI MT1-U06 Build/HuaweiMT1-U06",
	// "HUAWEI MT2-L01 Build/HuaweiMT2-L01; wv",
	// "U9200 Build/HuaweiU9200",
	// "HUAWEI P6-U06 Build/HuaweiP6-U06",
	"U8860 Build/HuaweiU8860",
	"HUAWEI U9508 Build/HuaweiU9508",
	"H30-U10 Build/HuaweiH30-U10",
	"HUAWEI G750-U10 Build/HuaweiG750-U10",
	"NEM-L51 Build/HONORNEM-L51",
	"KIW-L21 Build/HONORKIW-L21",
	"HUAWEI BLN-L24 Build/HUAWEIBLN-L24",
	"BLN-L22 Build/HONORBLN-L22",
	"PLK-AL10 Build/HONORPLK-AL10",
	"NEM-L21 Build/HONORNEM-L21",
	"LND-L29 Build/HONORLND-L29",
	"ATH-AL00 Build/HONORATH-AL00",
	"BND-L21 Build/HONORBND-L21",
	"FRD-L19 Build/HUAWEIFRD-L19",
	"PRA-AL00X Build/HONORPRA-AL00X",
	"PRA-AL00X Build/HONORPRA-AL00X",
	"VEN-L22 Build/HONORVEN-L22",
	"LLD-AL10 Build/HONORLLD-AL10",
	"LLD-AL20 Build/HONORLLD-AL20",
	"AUM-AL20 Build/HONORAUM-AL20",
	"EDI-AL10 Build/HUAWEIEDISON-AL10",
	"HUAWEI U8950-1 Build/HuaweiU8950-1",
	"RNE-L21 Build/HUAWEIRNE-L21",
	"EVR-AL00 Build/HUAWEIEVR-AL00",
	"HUAWEI NXT-AL10 Build/HUAWEINXT-AL10",
	"Nexus 6P",
	"LDN-LX2 Build/HUAWEILDN-LX2",
	"BAC-L21 Build/HUAWEIBAC-L21",
	"INE-LX2 Build/HUAWEIINE-LX2",
	"VTR-L29 Build/HUAWEIVTR-L29",
	"ELE-TL00 Build/HUAWEIELE-TL00",
	"ELE-L29 Build/HUAWEIELE-L29",
	"VOG-AL10 Build/HUAWEIVOG-AL10",
	"VOG-L29 Build/HUAWEIVOG-L29",
	"EVA-AL10 Build/HUAWEIEVA-AL10",
}

var HARMONYOS = [21]string{
	"HarmonyOS; DBR-W19; HMSCore 6.14.0.302",
	"HarmonyOS; MLD-AL00; HMSCore 6.13.0.352",
	"HarmonyOS; STG-AL00; HMSCore 6.14.0.302",
	"HarmonyOS; CTR-AL20; HMSCore 6.13.0.339",
	"HarmonyOS; FGD-AL00; HMSCore 6.13.0.342",
	"HarmonyOS; CTR-AL00; HMSCore 6.13.0.352",
	"HarmonyOS; ADA-AL00U; HMSCore 6.13.0.309",
	"HarmonyOS; CTR-AL00; HMSCore 6.14.0.302",
	"HarmonyOS; BRA-AL00; HMSCore 6.13.0.339",
	"HarmonyOS 4.0",
	"HarmonyOS 2.0",
	"HarmonyOS 3.0",
	"OXF-AN00 Build/HUAWEIOXF-AN00; HMSCore 6.13.0.302",
	"HarmonyOS/4.0.3.601; JAD-AL50; HMSCore 6.13.0.320 Build/AP31.240223.016.A3; x64; ARM64",
	"NOH-AN00 Build/HUAWEINOH-AN00",
	"PAL-AL00 Build/HUAWEIPAL-AL00",
	"GOT-W09 Build/HUAWEIGOT-W09",
	"ALT-AL00 Build/HUAWEIALT-AL00",
	"TET-AN00 Build/HUAWEITET-AN00",
	"TAS-AL00 Build/HUAWEITAS-AL00",
	"HarmonyOS/2.2.0;DBY-W09 Build/HUAWEIDBY-W09",
}

var GALAXY = [55]string{
	"SAMSUNG SM-A105F Build/PPR1.180610.011",
	"SAMSUNG SM-A205F Build/PPR1.180610.011",
	"SAMSUNG SM-A305F Build/PPR1.180610.011",
	"SAMSUNG SM-A505F Build/PPR1.180610.011",
	"SAMSUNG SM-A536W Build/TP1A.220624.014",
	"SAMSUNG SM-A536B Build/TP1A.220624.014",
	"SAMSUNG SM-A536E Build/TP1A.220624.014",
	"SAMSUNG SM-A546E Build/TP1A.220624.014",
	"SAMSUNG SM-A546W Build/TP1A.220624.014",
	"SAMSUNG SM-A546B Build/TP1A.220624.014",
	"SAMSUNG SM-A556E Build/UP1A.231005.007",
	"SAMSUNG SM-A556W Build/UP1A.231005.007",
	"SAMSUNG SM-A556B Build/UP1A.231005.007",
	"SAMSUNG SM-W2021",
	"SAMSUNG SM-E346B",
	"SAMSUNG SM-E146B",
	"SAMSUNG SM-G925F",
	"SAMSUNG SM-M346B2",
	"SAMSUNG SM-M346B1",
	"SM-M536S Build/SP1A.210812.016",
	"SM-M536S Build/SP1A.210812.016",
	"SM-A826S Build/RP1A.200720.012",
	"SM-M536S Build/SP1A.210812.016",
	"SAMSUNG SM-E045F",
	"SM-A606Y Build/QP1A.190711.020",
	"SM-A705FN Build/PPR1.180610.011",
	"SM-A707F Build/RP1A.200720.012",
	"SM-A715F Build/TP1A.220624.014",
	"SM-A716U Build/TP1A.220624.014",
	"SM-A725F Build/RP1A.200720.012",
	"SM-A736B Build/SP1A.210812.016",
	"SM-A805F Build/RP1A.200720.012",
	"SAMSUNG SM-A805N",
	"SM-G313M",
	"SM-G316ML",
	"SM-G316M Build/KTU84P",
	"SM-C5000 Build/MMB29M",
	"SM-C5018",
	"SM-C5010",
	"SM-C7108",
	"SM-C7100 Build/NMF26X",
	"SM-C7100 Build/M1AJQ",
	"SM-S911U Build/TP1A.220624.014",
	"SM-S918B Build/UP1A.231005.007",
	"SM-S918W Build/TP1A.220624.014",
	"SM-S926U Build/UP1A.231005.007",
	"SM-S9210 Build/UP1A.231005.007",
	"SM-S928U Build/UP1A.231005.007",
	"SM-S926U Build/UP1A.231005.007",
	"SM-F741U Build/UP1A.231005.007",
	"SM-F946B Build/TP1A.220624.014",
	"SM-F936U Build/SP2A.220305.013",
	"SM-F741U Build/UP1A.231005.007",
	"SM-F707B Build/RP1A.200720.012",
	"SM-F731B Build/UP1A.231005.007",
}

var IPHONE = [32]string{"15_1_1", "15_2", "15_3", "16_0", "14_2", "14_0", "14_4", "15_4_1", "15_3", "15_1", "13.6",
	"13_2_3", "13_4", "13_0", "13_1_3", "13_2_3", "13_1", "16_0", "16_1_1", "16_4_1", "17_0",
	"17_1", "17_2", "17_3", "17_4", "18_0", "18_3", "18_4", "18_5", "18_6", "18_7", "19_2_3",
}

var XIAOMI = [12]string{
	"Xiaomi 14 Build/UKQ1.230804.001",
	"23127PN0CG Build/UKQ1.230804.001",
	"23116PN5BC Build/UKQ1.230804.001",
	"Xiaomi 14 Pro Build/UKQ1.230804.001",
	"Xiaomi 14 Ultra Build/UKQ1.231003.002",
	"Xiaomi 13 Ultra Build/TKQ1.221114.001",
	"Xiaomi 13T Pro Build/TP1A.220624.014",
	"22061218C Build/UKQ1.230804.001",
	"2308CPXD0C Build/UKQ1.230804.001",
	"24053PY09C Build/UKQ1.240116.001",
	"23046PNC9C Build/UP1A.231005.007",
	"23046PNC9C Build/TP1A.220624.014",
}

var REDMI = [20]string{
	"Redmi Note 12 Pro 5G Build/SP1A.210812.016",
	"Redmi 6 Build/O11019",
	"Redmi Y2 Build/OPM1.171019.011",
	"Redmi Y1 Lite Build/N2G47H",
	"Redmi Y1 Build/N2G47H",
	"Redmi S2 Build/OPM1.171019.011",
	"Redmi Note 5A Prime Build/N2G47H",
	"Redmi Note 4X Build/MMB29M",
	"M2012K11AC Build/SKQ1.220213.001",
	"M2012K11AC Build/RKQ1.200826.002",
	"M2012K10C Build/TP1A.220624.014",
	"M2012K11C",
	"22041211AC",
	"Redmi K40 Pro+ Build/RKQ1.201112.002",
	"22011211C Build/TP1A.220624.014",
	"Redmi K50 Ultra Build/TKQ1.220829.002",
	"2407FRK8EC Build/UP1A.231005.007",
	"23117RK66C Build/UKQ1.230804.001",
	"23078RKD5C Build/TP1A.220624.014",
	"22127RK46C Build/TKQ1.220905.001",
}

var APPLE_WEB_KIT = [2]string{
	"AppleWebKit/537.36 (KHTML, like Gecko)",
	"AppleWebKit/605.1.15 (KHTML, like Gecko)",
}

var EDGE = [20]string{
	"Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0",
	"Chrome/124.0.6300.198 Safari/537.36 Edg/126.0.2364.61",
	"Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0 Config/91.2.2513.17",
	"Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.113",
	"Chrome/126.0.0.0 Safari/537.36 Edg/126.0.0.0 Agency/97.8.7535.73",
	"Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.113",
	"Chrome/125.0.6305.194 Safari/537.36 Edg/126.0.2337.69",
	"Chrome/126.0.0.0 Safari/537.36 Edg/126.0.0.0 ;Build/0220",
	"Chrome/126.0.6265.214 Safari/537.36 Edg/125.0.2325.85",
	"Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0 Viewer/99.9.8984.89",
	"Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.51",
	"Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0 Viewer/99.9.8984.89",
	"Chrome/126.0.6265.214 Safari/537.36 Edg/125.0.2325.85",
	"Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0 Viewer/99.9.8984.89",
	"Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.51",
	"Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0 Viewer/99.9.8984.89",
	"Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0 GLS/100.10.9989.100",
	"Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0 OneOutlook/1.2024.214.400",
	"Chrome/125.0.6314.196 Safari/537.36 Edg/124.0.2236.66",
	"Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.",
}

var OPERA = [21]string{
	"Chrome/125.0.0.0 Safari/537.36 OPR/110.0.5130.23",
	"Chrome/125.0.0.0 Safari/537.36 OPR/110.0.5130.35",
	"Chrome/125.0.0.0 Safari/537.36 OPR/110.0.5130.39",
	"Chrome/125.0.0.0 Safari/537.36 OPR/110.0.5130.49",
	"Chrome/125.0.0.0 Safari/537.36 OPR/110.0.5130.64",
	"Chrome/125.0.0.0 Safari/537.36 OPR/110.0.5130.66",
	"Chrome/125.0.0.0 Safari/537.36 OPR/110.0.5130.82",
	"Chrome/125.0.0.0 Safari/537.36 OPR/111.0.5168.25",
	"Chrome/125.0.0.0 Safari/537.36 OPR/111.0.5168.43",
	"Chrome/125.0.0.0 Safari/537.36 OPR/111.0.5168.55",
	"Chrome/125.0.0.0 Safari/537.36 OPR/111.0.5168.61",
	"Chrome/126.0.0.0 Safari/537.36 OPR/112.0.5197.24",
	"Chrome/126.0.0.0 Safari/537.36 OPR/112.0.5197.25",
	"Chrome/126.0.0.0 Safari/537.36 OPR/112.0.5197.30",
	"Chrome/126.0.0.0 Safari/537.36 OPR/112.0.5197.39",
	"Chrome/126.0.0.0 Safari/537.36 OPR/112.0.5197.53",
	"Chrome/127.0.0.0 Safari/537.36 OPR/113.0.5230.31",
	"Chrome/127.0.0.0 Safari/537.36 OPR/113.0.5230.32",
	"Chrome/127.0.0.0 Safari/537.36 OPR/113.0.5230.47",
	"Chrome/127.0.0.0 Safari/537.36 OPR/113.0.5230.55",
	"Chrome/127.0.0.0 Safari/537.36 OPR/113.0.5230.62",
}

var CHROME = [26]string{
	"126.0.8295.25 Safari/537.36",
	"127.0.6533.64 Safari/537.36",
	"127.0.4824.29 Safari/537.36",
	"126.0.6478.134 Safari/537.36",
	"126.0.6478.186 Safari/537.36",
	"126.0.6478.114 Safari/537.36.0",
	"126.0.6478.178 Safari/537.36",
	"125.0.6422.77 Safari/537.36",
	"125.0.6422.137 Safari/537.36",
	"125.0.0.0 Safari/537.36 Viewer/99.9.8782.87",
	"125.0.6422.165 Safari/537.36",
	"125.0.6422.41 Safari/537.36",
	"125.0.6422.78 Safari/537.36",
	"124.0.6367.179 Safari/537.36",
	"124.0.0.0 Safari/537.36",
	"124.0.6330.194 Safari/537.36",
	"124.0.6284.211 Safari/537.36",
	"124.0.6367.225 Safari/537.36",
	"123.0.6312.40 Safari/537.36",
	"123.0.0.0 Safari/537.36",
	"123.0.6312.99 Safari/537.36",
	"123.0.6268.217 Safari/537.36",
	"123.0.6312.86 Safari/537.36",
	"123.0.6312.112 Safari/537.36",
	"123.0.6312.134 Safari/537.36",
	"123.0.6312.118 Safari/537.36",
}

var FIREFOX = [20]string{
	"Firefox/101.0",
	"Firefox/129.4",
	"Firefox/129.7",
	"Firefox/129.9",
	"Firefox/129.2",
	"Firefox/129.0",
	"Firefox/129.0.2",
	"Firefox/130.3",
	"Firefox/130.1",
	"Firefox/130.0",
	"Firefox/131.7",
	"Firefox/131.0",
	"Firefox/128.0",
	"Firefox/128.7",
	"Firefox/128.1",
	"Firefox/128.9",
	"Firefox/128.4",
	"Firefox/127.0",
	"Firefox/127.6",
	"Firefox/126.0",
}

var SAFARI = [25]string{
	"Version/17.6 Safari/605",
	"Version/17.5 Safari/605.1.15",
	"Version/17.0.96 Safari/614.31.14",
	"Version/17.3.1 Safari/605.1.15",
	"Version/17.0.7 Safari/605.1.15",
	"Version/17.4 Safari/618.3.5",
	"Version/17.0.6 Safari/605.1.15",
	"Version/17.0.3 Safari/605.1.15",
	"Version/17.4.1 Safari/605.1.15",
	"Version/17.0.8 Safari/605.1.15",
	"Version/17.3.87 Safari/618.35",
	"Version/17.3 Safari/617.35",
	"Version/17.7.74 Safari/617.26",
	"Version/17.5 Safari/605.1.15 Ddg/17.5",
	"Version/16.0 Safari/605.1.15",
	"Version/16.5 Safari/605.1.15 Ddg/17.6",
	"Version/16.6 Safari/605.1.1",
	"Version/16.5 Safari/605.1.15 Ddg/18.0",
	"Version/16.5 Safari/605.1.15 Ddg/16.7",
	"Version/15.4 Safari/605.1.15",
	"Version/15.2 Safari/605.1.15",
	"Version/15.7 Safari/605.1.15",
	"Version/15.0 Safari/605.1.15",
	"Version/15.5 Safari/605.1.15",
	"Version/15.6.4 Safari/605.1.15",
}
