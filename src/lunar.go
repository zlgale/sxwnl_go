package src

import (
	"math"
	"strings"
)

//农历基础构件
var numCn = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"} //中文数字
var Gan = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var Zhi = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
var ShX = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
var XiZ = []string{"摩羯", "水瓶", "双鱼", "白羊", "金牛", "双子", "巨蟹", "狮子", "处女", "天秤", "天蝎", "射手"}
var yxmc = []string{"朔", "上弦", "望", "下弦"} //月相名称表
var jqmc = []string{"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪"}
var ymc = []string{"十一", "十二", "正", "二", "三", "四", "五", "六", "七", "八", "九", "十"} //月名称,建寅
var rmc = []string{"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十", "卅一"}

//纪年数据结构：数据用逗号分开，每7个描述一个年号，格式为:起始公元,使用年数,已用年数,朝代,朝号,皇帝,年号
var s = "-2069,45,0,夏,禹,,禹,-2024,10,0,夏,启,,启,-2014,25,0,夏,太康,,太康,-1986,14,0,夏,仲康,,仲康,-1972,28,0,夏,相,,相,-1944,2,0,夏,后羿,,后羿,-1942,38,0,夏,寒浞,,寒浞,-1904,21,0,夏,少康,,少康,-1883,17,0,夏,杼,,杼,-1866,26,0,夏,槐,,槐,-1840,18,0,夏,芒,,芒,-1822,16,0,夏,泄,,泄,-1806,59,0,夏,不降,,不降,-1747,21,0,夏,扃,,扃,-1726,21,0,夏,廑,,廑," +
	"-1705,31,0,夏,孔甲,,孔甲,-1674,11,0,夏,皋,,皋,-1663,11,0,夏,发,,发,-1652,53,0,夏,桀,,桀,-1599,11,0,商,商太祖,汤,商汤,-1588,1,0,商,商代王,太乙,商代王,-1587,2,0,商,哀王,子胜,外丙,-1585,4,0,商,懿王,子庸,仲壬,-1581,12,0,商,太宗,子至,太甲,-1569,29,0,商,昭王,子绚,沃丁,-1540,25,0,商,宣王,子辩,太庚,-1515,17,0,商,敬王,子高,小甲,-1498,13,0,商,元王,子密,雍己,-1485,75,0,商,中宗,子伷,太戊,-1410,11,0,商,孝成王,子庄,仲丁," +
	"-1399,15,0,商,思王,子发,外壬,-1384,9,0,商,前平王,子整,河亶甲,-1375,19,0,商,穆王,子滕,祖乙,-1356,16,0,商,桓王,子旦,祖辛,-1340,5,0,商,僖王,子逾,沃甲,-1335,9,0,商,庄王,子新,祖丁,-1326,6,0,商,顷王,子更,南庚,-1320,7,0,商,悼王,子和,阳甲,-1313,42,0,商,世祖,子旬,盘庚,-1271,21,0,商,章王,子颂,小辛,-1250,1,0,商,惠王,子敛,小乙,-1249,59,0,商,高宗,子昭,武丁,-1190,2,0,商,后平王,子跃,祖庚,-1188,33,0,商,世宗,子载,祖甲,-1155,8,0,商,甲宗,子先,廪辛," +
	"-1147,1,0,商,康祖,子嚣,庚丁,-1146,35,0,商,武祖,子瞿,武乙,-1111,11,0,商,匡王,子托,文丁,-1100,26,0,商,德王,子羡,帝乙,-1074,29,0,商,纣王,子寿,帝辛,-1045,4,0,西周,武王,姬发,武王,-1041,22,0,西周,成王,姬诵,成王,-1019,25,0,西周,康王,姬钊,康王,-994,19,0,西周,昭王,姬瑕,昭王,-975,54,0,西周,穆王,姬满,穆王,-921,23,0,西周,共王,姬繄,共王,-898,8,0,西周,懿王,姬囏,懿王,-890,6,0,西周,孝王,姬辟方,孝王,-884,8,0,西周,夷王,姬燮,夷王,-876,36,0,西周,厉王,姬胡,厉王," +
	"-840,14,0,西周,厉王,姬胡,共和,-826,46,0,西周,宣王,姬静,宣王,-780,11,0,西周,幽王,姬宫湦,幽王,-769,51,0,东周,平王,姬宜臼,平王,-718,23,0,东周,桓王,姬林,桓王,-695,15,0,东周,庄王,姬佗,庄王,-680,5,0,东周,釐王,姬胡齐,釐王,-675,25,0,东周,惠王,姬阆,惠王,-650,33,0,东周,襄王,姬郑,襄王,-617,6,0,东周,顷王,姬壬臣,顷王,-611,6,0,东周,匡王,姬班,匡王,-605,21,0,东周,定王,姬瑜,定王,-584,14,0,东周,简王,姬夷,简王,-570,27,0,东周,灵王,姬泄心,灵王,-543,24,0,东周,景王,姬贵,景王," +
	"-519,1,0,东周,悼王,姬勐,悼王,-518,44,0,东周,敬王,姬匄,敬王,-474,7,0,东周,元王,姬仁,元王,-467,27,0,东周,贞定王,姬介,贞定王,-440,1,0,东周,哀王-思王,姬去疾-姬叔,哀王-思王,-439,15,0,东周,考王,姬嵬,考王,-424,24,0,东周,威烈王,姬午,威烈王,-400,26,0,东周,安王,姬骄,安王,-374,7,0,东周,烈王,姬喜,烈王,-367,48,0,东周,显王,姬扁,显王,-319,6,0,东周,慎靓王,姬定,慎靓王,-313,8,0,东周,赧王,姬延,赧王,-305,56,0,战国-秦,昭襄王,嬴则,昭襄王,-249,1,0,战国-秦,孝文王,嬴柱,孝文王,-248,3,0,战国-秦,庄襄王,嬴子楚,庄襄王," +
	"-245,25,0,秦,嬴政,嬴政,嬴政,-220,12,0,秦,始皇帝,嬴政,始皇,-208,3,0,秦,二世皇帝,嬴胡亥,二世,-205,12,0,西汉,高帝,刘邦,高帝,-193,7,0,西汉,惠帝,刘盈,惠帝,-186,8,0,西汉,高后,吕雉,高后,-178,16,0,西汉,文帝,刘恒,文帝,-162,7,0,西汉,文帝,刘恒,后元,-155,7,0,西汉,景帝,刘启,景帝,-148,6,0,西汉,景帝,刘启,中元,-142,3,0,西汉,景帝,刘启,后元,-139,6,0,西汉,武帝,刘彻,建元,-133,6,0,西汉,武帝,刘彻,元光,-127,6,0,西汉,武帝,刘彻,元朔,-121,6,0,西汉,武帝,刘彻,元狩," +
	"-115,6,0,西汉,武帝,刘彻,元鼎,-109,6,0,西汉,武帝,刘彻,元封,-103,4,0,西汉,武帝,刘彻,太初,-99,4,0,西汉,武帝,刘彻,天汉,-95,4,0,西汉,武帝,刘彻,太始,-91,4,0,西汉,武帝,刘彻,征和,-87,2,0,西汉,武帝,刘彻,后元,-85,6,0,西汉,昭帝,刘弗陵,始元,-79,6,0,西汉,昭帝,刘弗陵,元凤,-73,1,0,西汉,昭帝,刘弗陵,元平,-72,4,0,西汉,宣帝,刘询,本始,-68,4,0,西汉,宣帝,刘询,地节,-64,4,0,西汉,宣帝,刘询,元康,-60,4,0,西汉,宣帝,刘询,神爵,-56,4,0,西汉,宣帝,刘询,五凤," +
	"-52,4,0,西汉,宣帝,刘询,甘露,-48,1,0,西汉,宣帝,刘询,黄龙,-47,5,0,西汉,元帝,刘奭,初元,-42,5,0,西汉,元帝,刘奭,永光,-37,5,0,西汉,元帝,刘奭,建昭,-32,1,0,西汉,元帝,刘奭,竟宁,-31,4,0,西汉,成帝,刘骜,建始,-27,4,0,西汉,成帝,刘骜,河平,-23,4,0,西汉,成帝,刘骜,阳朔,-19,4,0,西汉,成帝,刘骜,鸿嘉,-15,4,0,西汉,成帝,刘骜,永始,-11,4,0,西汉,成帝,刘骜,元延,-7,2,0,西汉,成帝,刘骜,绥和,-5,4,0,西汉,哀帝,刘欣,建平,-1,2,0,西汉,哀帝,刘欣,元寿," +
	"1,5,0,西汉,平帝,刘衍,元始,6,2,0,西汉,孺子婴,王莽摄政,居摄,8,1,0,西汉,孺子婴,王莽摄政,初始,9,5,0,新,,王莽,始建国,14,6,0,新,,王莽,天凤,20,3,0,新,,王莽,地皇,23,2,0,西汉,更始帝,刘玄,更始,25,31,0,东汉,光武帝,刘秀,建武,56,2,0,东汉,光武帝,刘秀,建武中元,58,18,0,东汉,明帝,刘庄,永平,76,8,0,东汉,章帝,刘炟,建初,84,3,0,东汉,章帝,刘炟,元和,87,2,0,东汉,章帝,刘炟,章和,89,16,0,东汉,和帝,刘肇,永元,105,1,0,东汉,和帝,刘肇,元兴," +
	"106,1,0,东汉,殇帝,刘隆,延平,107,7,0,东汉,安帝,刘祜,永初,114,6,0,东汉,安帝,刘祜,元初,120,1,0,东汉,安帝,刘祜,永宁,121,1,0,东汉,安帝,刘祜,建光,122,4,0,东汉,安帝,刘祜,延光,126,6,0,东汉,顺帝,刘保,永建,132,4,0,东汉,顺帝,刘保,阳嘉,136,6,0,东汉,顺帝,刘保,永和,142,2,0,东汉,顺帝,刘保,汉安,144,1,0,东汉,顺帝,刘保,建康,145,1,0,东汉,冲帝,刘炳,永嘉,146,1,0,东汉,质帝,刘缵,本初,147,3,0,东汉,桓帝,刘志,建和,150,1,0,东汉,桓帝,刘志,和平," +
	"151,2,0,东汉,桓帝,刘志,元嘉,153,2,0,东汉,桓帝,刘志,永兴,155,3,0,东汉,桓帝,刘志,永寿,158,9,0,东汉,桓帝,刘志,延熹,167,1,0,东汉,桓帝,刘志,永康,168,4,0,东汉,灵帝,刘宏,建宁,172,5,0,东汉,灵帝,刘宏,熹平,178,6,0,东汉,灵帝,刘宏,光和,184,6,0,东汉,灵帝,刘宏,中平,190,4,0,东汉,献帝,刘协,初平,194,2,0,东汉,献帝,刘协,兴平,196,24,0,东汉,献帝,刘协,建安,220,7,0,三国-魏,文帝,曹丕,黄初,227,6,0,三国-魏,明帝,曹叡,太和,233,4,0,三国-魏,明帝,曹叡,青龙," +
	"237,3,0,三国-魏,明帝,曹叡,景初,240,9,0,三国-魏,齐王,曹芳,正始,249,5,0,三国-魏,齐王,曹芳,嘉平,254,2,0,三国-魏,高贵乡公,曹髦,正元,256,4,0,三国-魏,高贵乡公,曹髦,甘露,260,4,0,三国-魏,元帝,曹奂,景元,264,1,0,三国-魏,元帝,曹奂,咸熙,265,10,0,西晋,武帝,司马炎,泰始,275,5,0,西晋,武帝,司马炎,咸宁,280,10,0,西晋,武帝,司马炎,太康,290,10,0,西晋,武帝,司马炎,太熙,300,1,0,西晋,惠帝,司马衷,永康,301,1,0,西晋,惠帝,司马衷,永宁,302,2,0,西晋,惠帝,司马衷,太安,304,2,0,西晋,惠帝,司马衷,永安," +
	"306,1,0,西晋,惠帝,司马衷,光熙,307,6,0,西晋,怀帝,司马炽,永嘉,313,4,0,西晋,愍帝,司马邺,建兴,317,1,0,东晋,元帝,司马睿,建武,318,4,0,东晋,元帝,司马睿,大兴,322,1,0,东晋,元帝,司马睿,永昌,323,3,0,东晋,明帝,司马绍,太宁,326,9,0,东晋,成帝,司马衍,咸和,335,8,0,东晋,成帝,司马衍,咸康,343,2,0,东晋,康帝,司马岳,建元,345,12,0,东晋,穆帝,司马聃,永和,357,5,0,东晋,穆帝,司马聃,升平,362,1,0,东晋,哀帝,司马丕,隆和,363,3,0,东晋,哀帝,司马丕,兴宁,366,5,0,东晋,海西公,司马奕,太和," +
	"371,2,0,东晋,简文帝,司马昱,咸安,373,3,0,东晋,孝武帝,司马曜,甯康,376,21,0,东晋,孝武帝,司马曜,太元,397,5,0,东晋,安帝,司马德宗,隆安,402,3,0,东晋,安帝,司马德宗,元兴,405,14,0,东晋,安帝,司马德宗,义熙,419,1,0,东晋,恭帝,司马德文,元熙,420,3,0,南朝/宋,武帝,刘裕,永初,423,2,0,南朝/宋,少帝,刘义符,景平,424,30,0,南朝/宋,文帝,刘義隆,元嘉,454,3,0,南朝/宋,孝武,帝刘骏,孝建,457,8,0,南朝/宋,孝武,帝刘骏,大明,465,1,0,南朝/宋,废帝,刘子业,永光,465,1,0,南朝/宋,废帝,刘子业,景和," +
	"465,7,0,南朝/宋,明帝,刘彧,泰始,472,1,0,南朝/宋,明帝,刘彧,泰豫,473,5,0,南朝/宋,废帝,刘昱,元徽,477,3,0,南朝/宋,顺帝,刘准,升明,479,4,0,南朝/齐,高帝,萧道成,建元,483,11,0,南朝/齐,武帝,萧赜,永明,494,1,0,南朝/齐,欎林王,萧昭业,隆昌,494,1,0,南朝/齐,海陵王,萧昭文,延兴,494,5,0,南朝/齐,明帝,萧鸾,建武,498,1,0,南朝/齐,明帝,萧鸾,永泰,499,3,0,南朝/齐,东昏侯,萧宝,中兴,501,2,0,南朝/齐,和帝,萧宝融,中兴,502,18,0,南朝/梁,武帝,萧衍,天监,520,8,0,南朝/梁,武帝,萧衍,普通,527,3,0,南朝/梁,武帝,萧衍,大通," +
	"529,6,0,南朝/梁,武帝,萧衍,中大通,535,12,0,南朝/梁,武帝,萧衍,大同,546,2,0,南朝/梁,武帝,萧衍,中大同,547,3,0,南朝/梁,武帝,萧衍,太清,550,2,0,南朝/梁,简文帝,萧纲,大宝,551,2,0,南朝/梁,豫章王,萧栋,天正,552,4,0,南朝/梁,元帝,萧绎,承圣,555,1,0,南朝/梁,贞阳侯,萧渊明,天成,555,2,0,南朝/梁,敬帝,萧方智,绍泰,556,2,0,南朝/梁,敬帝,萧方智,太平,557,3,0,南朝/陈,武帝,陈霸先,太平,560,7,0,南朝/陈,文帝,陈蒨,天嘉,566,1,0,南朝/陈,文帝,陈蒨,天康,567,2,0,南朝/陈,废帝,陈伯宗,光大,569,14,0,南朝/陈,宣帝,陈顼,太建," +
	"583,4,0,南朝/陈,后主,陈叔宝,至德,587,3,0,南朝/陈,后主,陈叔宝,祯明,555,8,0,南朝/后梁,宣帝,萧詧,大定,562,24,0,南朝/后梁,明帝,萧岿,天保,586,2,0,南朝/后梁,莒公,萧琮,广运,386,11,0,北朝/北魏,道武帝,拓跋圭,登国,396,3,0,北朝/北魏,道武帝,拓跋圭,皇始,398,7,0,北朝/北魏,道武帝,拓跋圭,天兴,404,6,0,北朝/北魏,道武帝,拓跋圭,天赐,409,5,0,北朝/北魏,明元帝,拓跋嗣,永兴,414,3,0,北朝/北魏,明元帝,拓跋嗣,神瑞,416,8,0,北朝/北魏,明元帝,拓跋嗣,泰常,424,5,0,北朝/北魏,太武帝,拓跋焘,始光,428,4,0,北朝/北魏,太武帝,拓跋焘,神麚,432,3,0,北朝/北魏,太武帝,拓跋焘,延和," +
	"435,6,0,北朝/北魏,太武帝,拓跋焘,太延,440,12,0,北朝/北魏,太武帝,拓跋焘,太平真君,451,2,0,北朝/北魏,太武帝,拓跋焘,正平,452,1,0,北朝/北魏,南安王,拓跋余,承平,452,3,0,北朝/北魏,文成帝,拓跋浚,兴安,454,2,0,北朝/北魏,文成帝,拓跋浚,兴光,455,5,0,北朝/北魏,文成帝,拓跋浚,太安,460,6,0,北朝/北魏,文成帝,拓跋浚,和平,466,2,0,北朝/北魏,献文帝,拓跋弘,天安,467,5,0,北朝/北魏,献文帝,拓跋弘,皇兴,471,6,0,北朝/北魏,教文帝,拓跋宏,延兴,476,1,0,北朝/北魏,孝文帝,拓跋宏,承明,477,23,0,北朝/北魏,孝文帝,拓跋宏,太和,500,4,0,北朝/北魏,宣武帝,元恪,景明,504,5,0,北朝/北魏,宣武帝,元恪,正始," +
	"508,5,0,北朝/北魏,宣武帝,元恪,永平,512,4,0,北朝/北魏,宣武帝,元恪,延昌,516,3,0,北朝/北魏,孝明帝,元诩,熙平,518,3,0,北朝/北魏,孝明帝,元诩,神龟,520,6,0,北朝/北魏,孝明帝,元诩,正光,525,3,0,北朝/北魏,孝明帝,元诩,孝昌,528,1,0,北朝/北魏,孝明帝,元诩,武泰,528,1,0,北朝/北魏,孝庄帝,元子攸,建义,528,3,0,北朝/北魏,孝庄帝,元子攸,永安,530,2,0,北朝/北魏,东海王,元晔,建明,531,2,0,北朝/北魏,节闵帝,元恭,普泰,531,2,0,北朝/北魏,安定王,元朗,中兴,532,1,0,北朝/北魏,孝武帝,元修,太昌,532,1,0,北朝/北魏,孝武帝,元修,永兴,532,3,0,北朝/北魏,孝武帝,元修,永熙," +
	"534,4,0,北朝/东魏,孝静帝,元善见,天平,538,2,0,北朝/东魏,孝静帝,元善见,元象,539,4,0,北朝/东魏,孝静帝,元善见,兴和,543,8,0,北朝/东魏,孝静帝,元善见,武定,535,17,0,北朝/西魏,文帝,元宝炬,大统,552,3,0,北朝/西魏,废帝,元钦,大统,554,3,0,北朝/西魏,恭帝,元廓,大统,550,10,0,北朝/北齐,文宣帝,高洋,天保,560,1,0,北朝/北齐,废帝,高殷,乾明,560,2,0,北朝/北齐,孝昭帝,高演,皇建,561,2,0,北朝/北齐,武成帝,高湛,太宁,562,4,0,北朝/北齐,武成帝,高湛,河清,565,5,0,北朝/北齐,温公,高纬,天统,570,7,0,北朝/北齐,温公,高纬,武平,576,2,0,北朝/北齐,温公,高纬,隆化," +
	"576,1,0,北朝/北齐,安德王,高延宗,德昌,577,1,0,北朝/北齐,幼主,高恒,承光,557,1,0,北朝/北周,闵帝,宇文觉,空,557,2,0,北朝/北周,明帝,宇文毓,空,559,2,0,北朝/北周,明帝,宇文毓,武成,561,5,0,北朝/北周,武帝,宇文邕,保定,566,7,0,北朝/北周,武帝,宇文邕,天和,572,7,0,北朝/北周,武帝,宇文邕,建德,578,1,0,北朝/北周,武帝,宇文邕,宣政,579,1,0,北朝/北周,宣帝,宇文贇,大成,579,2,0,北朝/北周,静帝,宇文衍,大象,581,1,0,北朝/北周,静帝,宇文衍,大定,581,20,0,隋,文帝,杨坚,开皇,601,4,0,隋,文帝,杨坚,仁寿,605,13,0,隋,炀帝,杨广,大业," +
	"617,2,0,隋,恭帝,杨侑,义宁,618,9,0,唐,高祖,李渊,武德,627,23,0,唐,太宗,李世民,贞观,650,6,0,唐,高宗,李治,永徽,656,6,0,唐,高宗,李治,显庆,661,3,0,唐,高宗,李治,龙朔,664,2,0,唐,高宗,李治,麟德,666,3,0,唐,高宗,李治,乾封,668,3,0,唐,高宗,李治,总章,670,5,0,唐,高宗,李治,咸亨,674,3,0,唐,高宗,李治,上元,676,4,0,唐,高宗,李治,仪凤,679,2,0,唐,高宗,李治,调露,680,2,0,唐,高宗,李治,永隆,681,2,0,唐,高宗,李治,开耀," +
	"682,2,0,唐,高宗,李治,永淳,683,1,0,唐,高宗,李治,弘道,684,1,0,唐,中宗,李显,嗣圣,684,1,0,唐,睿宗,李旦,文明,684,1,0,武周,则天后,武曌,光宅,685,4,0,武周,则天后,武曌,垂拱,689,1,0,武周,则天后,武曌,永昌,689,2,0,武周,则天后,武曌,载初,690,3,0,武周,则天后,武曌,天授,692,1,0,武周,则天后,武曌,如意,692,3,0,武周,则天后,武曌,长寿,694,1,0,武周,则天后,武曌,延载,695,1,0,武周,则天后,武曌,证圣,695,2,0,武周,则天后,武曌,天册万岁,696,1,0,武周,则天后,武曌,万岁登封," +
	"696,2,0,武周,则天后,武曌,万岁通天,697,1,0,武周,则天后,武曌,神功,698,3,0,武周,则天后,武曌,圣历,700,1,0,武周,则天后,武曌,久视,701,1,0,武周,则天后,武曌,大足,701,4,0,武周,则天后,武曌,长安,705,1,0,武周,则天后,李显,神龙,705,2,0,唐,中宗,李显,神龙,707,4,0,唐,中宗,李显,景龙,710,1,0,唐,温王,李重茂,唐隆,710,2,0,唐,睿宗,李旦,景云,712,1,0,唐,睿宗,李旦,太极,712,1,0,唐,睿宗,李旦,延和,712,2,0,唐,玄宗,李隆基,先天,713,29,0,唐,玄宗,李隆基,开元," +
	"742,15,0,唐,玄宗,李隆基,天宝,756,3,0,唐,肃宗,李亨,至德,758,3,0,唐,肃宗,李亨,乾元,760,3,0,唐,肃宗,李亨,上元,762,2,0,唐,肃宗,李亨,宝应,763,2,0,唐,代宗,李豫,广德,765,2,0,唐,肃宗,李亨,永泰,766,14,0,唐,肃宗,李亨,大历,780,4,0,唐,德宗,李适,建中,784,1,0,唐,德宗,李适,兴元,785,21,0,唐,德宗,李适,贞元,805,1,0,唐,顺宗,李诵,永贞,806,15,0,唐,宪宗,李纯,元和,821,4,0,唐,穆宗,李恒,长庆,825,3,0,唐,敬宗,李湛,宝历," +
	"827,9,0,唐,文宗,李昂,大和,836,5,0,唐,文宗,李昂,开成,841,6,0,唐,武宗,李炎,会昌,847,14,0,唐,宣宗,李忱,大中,860,15,0,唐,宣宗,李忱,咸通,874,6,0,唐,僖宗,李儇,乾符,880,2,0,唐,僖宗,李儇,广明,881,5,0,唐,僖宗,李儇,中和,885,4,0,唐,僖宗,李儇,光启,888,1,0,唐,僖宗,李儇,文德,889,1,0,唐,昭宗,李晔,龙纪,890,2,0,唐,昭宗,李晔,大顺,892,2,0,唐,昭宗,李晔,景福,894,5,0,唐,昭宗,李晔,乾宁,898,4,0,唐,昭宗,李晔,光化," +
	"901,4,0,唐,昭宗,李晔,天复,904,1,0,唐,昭宗,李晔,天佑,905,3,1,唐,昭宣帝,李祝,天佑,907,5,0,五代/梁,太祖,朱温,开平,911,2,0,五代/梁,太祖,朱温,乾化,913,1,0,五代/梁,庶人,朱友圭,凤历,913,3,2,五代/梁,末帝,朱友贞,乾化,915,7,0,五代/梁,末帝,朱友贞,贞明,921,3,0,五代/梁,末帝,朱友贞,龙德,923,4,0,五代/唐,庄宗,李存勗,同光,926,5,0,五代/唐,明宗,李嗣源,天成,930,4,0,五代/唐,明宗,李嗣源,长兴,934,1,0,五代/唐,闵帝,李从厚,应顺,934,3,0,五代/唐,潞王,李从珂,清泰,936,6,0,五代/晋,高祖,石敬瑭,天福," +
	"942,2,6,五代/晋,出帝,石重贵,天福,944,3,0,五代/晋,出帝,石重贵,开运,947,12,0,五代/汉,高祖,刘知远,天福,948,1,0,五代/汉,隐帝,刘承祐,乾祐,948,3,0,五代/汉,隐帝,刘承祐,乾祐,951,3,0,五代/周,太祖,郭威,广顺,954,1,0,五代/周,太祖,郭威,显德,954,6,0,五代/周,世宗,柴荣,显德,959,2,5,五代/周,恭帝,郭宗训,显德,960,4,0,北宋,太祖,赵匡胤,建隆,963,6,0,北宋,太祖,赵匡胤,乾德,968,9,0,北宋,太祖,赵匡胤,开宝,976,9,0,北宋,太宗,赵炅,太平兴国,984,4,0,北宋,太宗,赵炅,雍熙,988,2,0,北宋,太宗,赵炅,端拱," +
	"990,5,0,北宋,太宗,赵炅,淳化,995,3,0,北宋,太宗,赵炅,至道,998,6,0,北宋,真宗,赵恒,咸平,1004,4,0,北宋,真宗,赵恒,景德,1008,9,0,北宋,真宗,赵恒,大中祥符,1017,5,0,北宋,真宗,赵恒,天禧,1022,1,0,北宋,真宗,赵恒,乾兴,1023,10,0,北宋,仁宗,赵祯,天圣,1032,2,0,北宋,仁宗,赵祯,明道,1034,5,0,北宋,仁宗,赵祯,景祐,1038,3,0,北宋,仁宗,赵祯,宝元,1040,2,0,北宋,仁宗,赵祯,康定,1041,8,0,北宋,仁宗,赵祯,庆历,1049,6,0,北宋,仁宗,赵祯,皇祐,1054,3,0,北宋,仁宗,赵祯,至和," +
	"1056,8,0,北宋,仁宗,赵祯,嘉祐,1064,4,0,北宋,英宗,赵曙,治平,1068,10,0,北宋,神宗,赵顼,熙宁,1078,8,0,北宋,神宗,赵顼,元丰,1086,9,0,北宋,哲宗,赵煦,元祐,1094,5,0,北宋,哲宗,赵煦,绍圣,1098,3,0,北宋,哲宗,赵煦,元符,1101,1,0,北宋,徽宗,赵佶,建中靖国,1102,5,0,北宋,徽宗,赵佶,崇宁,1107,4,0,北宋,徽宗,赵佶,大观,1111,8,0,北宋,徽宗,赵佶,政和,1118,2,0,北宋,徽宗,赵佶,重和,1119,7,0,北宋,徽宗,赵佶,宣和,1126,2,0,北宋,钦宗,赵桓,靖康,1127,4,0,南宋,高宗,赵构,建炎," +
	"1131,32,0,南宋,高宗,赵构,绍兴,1163,2,0,南宋,孝宗,赵慎,隆兴,1165,9,0,南宋,孝宗,赵慎,乾道,1174,16,0,南宋,孝宗,赵慎,淳熙,1190,5,0,南宋,光宗,赵暴,绍熙,1195,6,0,南宋,宁宗,赵扩,庆元,1201,4,0,南宋,宁宗,赵扩,嘉泰,1205,3,0,南宋,宁宗,赵扩,开禧,1208,17,0,南宋,宁宗,赵扩,嘉定,1225,3,0,南宋,理宗,赵昀,宝庆,1228,6,0,南宋,理宗,赵昀,绍定,1234,3,0,南宋,理宗,赵昀,端平,1237,4,0,南宋,理宗,赵昀,嘉熙,1241,12,0,南宋,理宗,赵昀,淳祐,1253,6,0,南宋,理宗,赵昀,寶祐," +
	"1259,1,0,南宋,理宗,赵昀,开庆,1260,5,0,南宋,理宗,赵昀,景定,1265,10,0,南宋,度宗,赵禥,咸淳,1275,2,0,南宋,恭宗,赵(上“日”下“丝”),德祐 ,1276,3,0,南宋,端宗,赵(上“日”下“正”),景炎,1278,2,0,南宋,帝昺,赵昺,祥兴,1271,24,7,元,世祖,孛儿只斤·忽必烈,至元,1295,3,0,元,成宗,孛儿只斤·铁穆耳,元贞,1297,11,0,元,成宗,孛儿只斤·铁穆耳,大德,1308,4,0,元,武宗,孛儿只斤·海山,至大,1312,2,0,元,仁宗,孛儿只斤·爱育黎拔力八达,皇庆,1314,7,0,元,仁宗,孛儿只斤·愛育黎拔力八達,延祐,1321,3,0,元,英宗,孛儿只斤·宗硕德八剌,至治,1324,5,0,元,泰定帝,孛儿只斤·也孙铁木耳,泰定,1328,1,0,元,泰定帝,孛儿只斤·也孙铁木耳,至和," +
	"1328,1,0,元,幼主,孛儿只斤·阿速吉八,天顺,1328,3,0,元,文宗,孛儿只斤·图贴睦尔,天历,1330,3,0,元,文宗,孛儿只斤·图贴睦尔,至顺,1333,3,0,元,惠宗,孛儿只斤·妥镤贴睦尔,元统,1335,6,0,元,惠宗,孛儿只斤·妥镤贴睦尔,至元,1341,28,0,元,惠宗,孛儿只斤·妥镤贴睦尔,至正,1368,31,0,明,太祖,朱元璋,洪武,1399,4,0,明,惠帝,朱允溫,建文,1403,22,0,明,成祖,朱棣,永乐,1425,1,0,明,仁宗,朱高炽,洪熙,1426,10,0,明,宣宗,朱瞻基,宣德,1436,14,0,明,英宗,朱祁镇,正统,1450,7,0,明,代宗,朱祁钰,景泰,1457,8,0,明,英宗,朱祁镇,天顺,1465,23,0,明,宪宗,朱见深,成化," +
	"1488,18,0,明,孝宗,朱祐樘,弘治,1506,16,0,明,武宗,朱厚照,正德,1522,45,0,明,世宗,朱厚熜,嘉靖,1567,6,0,明,穆宗,朱载贺,隆庆,1573,48,0,明,神宗,朱翊钧,万历,1620,1,0,明,光宗,朱常洛,泰昌,1621,7,0,明,熹宗,朱同校,天启,1628,17,0,明,毅宗,朱由检,崇祯,1644,18,0,清,世祖,爱新觉罗福临,顺治,1662,61,0,清,圣祖,爱新觉罗玄烨,康熙,1723,13,0,清,世宗,爱新觉罗胤禛,雍正,1736,60,0,清,高宗,爱新觉罗弘历,乾隆,1796,25,0,清,仁宗,爱新觉罗颙琰,嘉庆,1821,30,0,清,宣宗,爱新觉罗旻宁,道光,1851,11,0,清,文宗,爱新觉罗奕詝,咸丰," +
	"1862,13,0,清,穆宗,爱新觉罗载淳,同治,1875,34,0,清,德宗,爱新觉罗载湉,光绪,1909,3,0,清,无朝,爱新觉罗溥仪,宣统,1912,37,0,近、现代,中华民国,,民国,1949,9999,1948,当代,中国,,公历纪元"

var JNB = strings.Split(s, ",")
// TODO:获取年号

//精气
func qi_accurate(W float64) float64 {
	var t = S_aLon_t(W) * 36525
	return t - dt_T(t) + 8/24
}

//精朔
func so_accurate(W float64) float64 {
	var t = MS_aLon_t(W) * 36525
	return t - dt_T(t) + 8/24
}

//精气
func qi_accurate2(jd float64) float64 {
	var d = math.Pi / 12;
	var w = math.Floor((jd+293)/365.2422*24) * d
	var a = qi_accurate(w)
	if a-jd > 5 {
		return qi_accurate(w - d)
	}
	if a-jd < -5 {
		return qi_accurate(w + d)
	}
	return a
}

//精朔
func so_accurate2(jd float64) float64 {
	return so_accurate(math.Floor((jd+8)/29.5306) * math.Pi * 2)
}

/************************
  实气实朔计算器
  适用范围 -722年2月22日——1959年12月
  平气平朔计算使用古历参数进行计算
  定朔、定气计算使用开普勒椭圆轨道计算，同时考虑了光行差和力学时与UT1的时间差
  古代历算仅在晚期才使用开普勒方法计算，此前多采用一些修正表并插值得到，精度很低，与本程序中
的开普勒方法存在误差，造成朔日计算错误1千多个，这些错误使用一个修正表进行订正。同样，定气部分
也使用了相同的方法时行订正。
  平气朔表的算法(线性拟合)：
  气朔日期计算公式：D = k*n + b  , 式中n=0,1,2,3,...,N-1, N为该式适用的范围
  h表示k不变b允许的误差,如果b不变则k许可误差为h/N
  每行第1个参数为k,第2参数为b
  public中定义的成员可以直接使用
*************************/
var suoKB = []float64{ //朔直线拟合参数
	1457698.231017, 29.53067166, // -721-12-17 h=0.00032 古历·春秋
	1546082.512234, 29.53085106, // -479-12-11 h=0.00053 古历·战国
	1640640.735300, 29.53060000, // -221-10-31 h=0.01010 古历·秦汉
	1642472.151543, 29.53085439, // -216-11-04 h=0.00040 古历·秦汉

	1683430.509300, 29.53086148, // -104-12-25 h=0.00313 汉书·律历志(太初历)平气平朔
	1752148.041079, 29.53085097, //   85-02-13 h=0.00049 后汉书·律历志(四分历)
	//1807665.420323,29.53059851, //  237-02-12 h=0.00033 晋书·律历志(景初历)
	1807724.481520, 29.53059851, //  237-04-12 h=0.00033 晋书·律历志(景初历)
	1883618.114100, 29.53060000, //  445-01-24 h=0.00030 宋书·律历志(何承天元嘉历)
	1907360.704700, 29.53060000, //  510-01-26 h=0.00030 宋书·律历志(祖冲之大明历)
	1936596.224900, 29.53060000, //  590-02-10 h=0.01010 随书·律历志(开皇历)
	1939135.675300, 29.53060000, //  597-01-24 h=0.00890 随书·律历志(大业历)
	1947168.00,                  //  619-01-21
}

var qiKB = []float64{ //气直线拟合参数
	1640650.479938, 15.21842500, // -221-11-09 h=0.01709 古历·秦汉
	1642476.703182, 15.21874996, // -216-11-09 h=0.01557 古历·秦汉

	1683430.515601, 15.218750011, // -104-12-25 h=0.01560 汉书·律历志(太初历)平气平朔 回归年=365.25000
	1752157.640664, 15.218749978, //   85-02-23 h=0.01559 后汉书·律历志(四分历) 回归年=365.25000
	1807675.003759, 15.218620279, //  237-02-22 h=0.00010 晋书·律历志(景初历) 回归年=365.24689
	1883627.765182, 15.218612292, //  445-02-03 h=0.00026 宋书·律历志(何承天元嘉历) 回归年=365.24670
	1907369.128100, 15.218449176, //  510-02-03 h=0.00027 宋书·律历志(祖冲之大明历) 回归年=365.24278
	1936603.140413, 15.218425000, //  590-02-17 h=0.00149 随书·律历志(开皇历) 回归年=365.24220
	1939145.524180, 15.218466998, //  597-02-03 h=0.00121 随书·律历志(大业历) 回归年=365.24321
	1947180.798300, 15.218524844, //  619-02-03 h=0.00052 新唐书·历志(戊寅元历)平气定朔 回归年=365.24460
	1964362.041824, 15.218533526, //  666-02-17 h=0.00059 新唐书·历志(麟德历) 回归年=365.24480
	1987372.340971, 15.218513908, //  729-02-16 h=0.00096 新唐书·历志(大衍历,至德历) 回归年=365.24433
	1999653.819126, 15.218530782, //  762-10-03 h=0.00093 新唐书·历志(五纪历) 回归年=365.24474
	2007445.469786, 15.218535181, //  784-02-01 h=0.00059 新唐书·历志(正元历,观象历) 回归年=365.24484
	2021324.917146, 15.218526248, //  822-02-01 h=0.00022 新唐书·历志(宣明历) 回归年=365.24463
	2047257.232342, 15.218519654, //  893-01-31 h=0.00015 新唐书·历志(崇玄历) 回归年=365.24447
	2070282.898213, 15.218425000, //  956-02-16 h=0.00149 旧五代·历志(钦天历) 回归年=365.24220
	2073204.872850, 15.218515221, //  964-02-16 h=0.00166 宋史·律历志(应天历) 回归年=365.24437
	2080144.500926, 15.218530782, //  983-02-16 h=0.00093 宋史·律历志(乾元历) 回归年=365.24474
	2086703.688963, 15.218523776, // 1001-01-31 h=0.00067 宋史·律历志(仪天历,崇天历) 回归年=365.24457
	2110033.182763, 15.218425000, // 1064-12-15 h=0.00669 宋史·律历志(明天历) 回归年=365.24220
	2111190.300888, 15.218425000, // 1068-02-15 h=0.00149 宋史·律历志(崇天历) 回归年=365.24220
	2113731.271005, 15.218515671, // 1075-01-30 h=0.00038 李锐补修(奉元历) 回归年=365.24438
	2120670.840263, 15.218425000, // 1094-01-30 h=0.00149 宋史·律历志 回归年=365.24220
	2123973.309063, 15.218425000, // 1103-02-14 h=0.00669 李锐补修(占天历) 回归年=365.24220
	2125068.997336, 15.218477932, // 1106-02-14 h=0.00056 宋史·律历志(纪元历) 回归年=365.24347
	2136026.312633, 15.218472436, // 1136-02-14 h=0.00088 宋史·律历志(统元历,乾道历,淳熙历) 回归年=365.24334
	2156099.495538, 15.218425000, // 1191-01-29 h=0.00149 宋史·律历志(会元历) 回归年=365.24220
	2159021.324663, 15.218425000, // 1199-01-29 h=0.00149 宋史·律历志(统天历) 回归年=365.24220
	2162308.575254, 15.218461742, // 1208-01-30 h=0.00146 宋史·律历志(开禧历) 回归年=365.24308
	2178485.706538, 15.218425000, // 1252-05-15 h=0.04606 淳祐历 回归年=365.24220
	2178759.662849, 15.218445786, // 1253-02-13 h=0.00231 会天历 回归年=365.24270
	2185334.020800, 15.218425000, // 1271-02-13 h=0.00520 宋史·律历志(成天历) 回归年=365.24220
	2187525.481425, 15.218425000, // 1277-02-12 h=0.00520 本天历 回归年=365.24220
	2188621.191481, 15.218437494, // 1280-02-13 h=0.00015 元史·历志(郭守敬授时历) 回归年=365.24250
	2322147.76,                   // 1645-09-21
}

func so_low(W float64) float64 { //低精度定朔计算,在2000年至600，误差在2小时以内(仍比古代日历精准很多)
	var v = 7771.37714500204
	var t = (W + 1.08472) / v
	t -= (-0.0000331*t*t +
		0.10976*math.Cos(0.785+8328.6914*t) +
		0.02224*math.Cos(0.187+7214.0629*t) -
		0.03342*math.Cos(4.669+628.3076*t))/v +
		(32*(t+1.8)*(t+1.8)-20)/86400/36525
	return t*36525 + 8/24
}

func qi_low(W float64) float64 { //最大误差小于30分钟，平均5分
	var t, L float64
	var v = 628.3319653318
	t = (W - 4.895062166) / v                                                                             //第一次估算,误差2天以内
	t -= (53*t*t + 334116*math.Cos(4.67+628.307585*t) + 2061*math.Cos(2.678+628.3076*t)*t) / v / 10000000 //第二次估算,误差2小时以内

	L = 48950621.66 + 6283319653.318*t + 53*t*t + //平黄经
		334166*math.Cos(4.669257+628.307585*t) + //地球椭圆轨道级数展开
		3489*math.Cos(4.6261+1256.61517*t) + //地球椭圆轨道级数展开
		2060.6*math.Cos(2.67823+628.307585*t)*t - //一次泊松项
		994 - 834*math.Sin(2.1824-33.75705*t) //光行差与章动修正

	t -= (L/10000000-W)/628.332 + (32*(t+1.8)*(t+1.8)-20)/86400/36525
	return t*36525 + 8/24
}

func qi_high(W float64) float64 { //较高精度气
	var t = S_aLon_t2(W) * 36525
	t = t - dt_T(t) + 8/24
	var v = math.Mod(t+0.5, 1) * 86400
	if v < 1200 || v > 86400-1200 {
		t = S_aLon_t(W)*36525 - dt_T(t) + 8/24
	}
	return t
}

func so_high(W float64) float64 { //较高精度朔
	var t = MS_aLon_t2(W) * 36525
	t = t - dt_T(t) + 8/24
	var v = math.Mod(t+0.5, 1) * 86400
	if v < 1800 || v > 86400-1800 {
		t = MS_aLon_t(W)*36525 - dt_T(t) + 8/24
	}
	return t
}

func jieya(s string) string { //气朔解压缩
	var o = "0000000000"
	var o2 = o + o
	s = strings.Replace(s, "J", "00", -1)
	s = strings.Replace(s, "I", "000", -1)
	s = strings.Replace(s, "H", "0000", -1)
	s = strings.Replace(s, "G", "00000", -1)
	s = strings.Replace(s, "t", "02", -1)
	s = strings.Replace(s, "s", "002", -1)
	s = strings.Replace(s, "r", "0002", -1)
	s = strings.Replace(s, "q", "00002", -1)
	s = strings.Replace(s, "p", "000002", -1)
	s = strings.Replace(s, "o", "0000002", -1)
	s = strings.Replace(s, "n", "00000002", -1)
	s = strings.Replace(s, "m", "000000002", -1)
	s = strings.Replace(s, "l", "0000000002", -1)
	s = strings.Replace(s, "k", "01", -1)
	s = strings.Replace(s, "j", "0101", -1)
	s = strings.Replace(s, "i", "001", -1)
	s = strings.Replace(s, "h", "001001", -1)
	s = strings.Replace(s, "g", "0001", -1)
	s = strings.Replace(s, "f", "00001", -1)
	s = strings.Replace(s, "e", "000001", -1)
	s = strings.Replace(s, "d", "0000001", -1)
	s = strings.Replace(s, "c", "00000001", -1)
	s = strings.Replace(s, "b", "000000001", -1)
	s = strings.Replace(s, "a", "0000000001", -1)
	s = strings.Replace(s, "A", o2+o2+o2, -1)
	s = strings.Replace(s, "B", o2+o2+o, -1)
	s = strings.Replace(s, "C", o2+o2, -1)
	s = strings.Replace(s, "D", o2+o, -1)
	s = strings.Replace(s, "E", o2, -1)
	s = strings.Replace(s, "F", o, -1)
	return s
}

func SSQInit() (string, string) { //实朔实气计算器

	var suoS, qiS string
	//  619-01-21开始16598个朔日修正表 d0=1947168
	suoS = "EqoFscDcrFpmEsF2DfFideFelFpFfFfFiaipqti1ksttikptikqckstekqttgkqttgkqteksttikptikq2fjstgjqttjkqttgkqt"
	suoS += "ekstfkptikq2tijstgjiFkirFsAeACoFsiDaDiADc1AFbBfgdfikijFifegF1FhaikgFag1E2btaieeibggiffdeigFfqDfaiBkF"
	suoS += "1kEaikhkigeidhhdiegcFfakF1ggkidbiaedksaFffckekidhhdhdikcikiakicjF1deedFhFccgicdekgiFbiaikcfi1kbFibef"
	suoS += "gEgFdcFkFeFkdcfkF1kfkcickEiFkDacFiEfbiaejcFfffkhkdgkaiei1ehigikhdFikfckF1dhhdikcfgjikhfjicjicgiehdik"
	suoS += "cikggcifgiejF1jkieFhegikggcikFegiegkfjebhigikggcikdgkaFkijcfkcikfkcifikiggkaeeigefkcdfcfkhkdgkegieid"
	suoS += "hijcFfakhfgeidieidiegikhfkfckfcjbdehdikggikgkfkicjicjF1dbidikFiggcifgiejkiegkigcdiegfggcikdbgfgefjF1"
	suoS += "kfegikggcikdgFkeeijcfkcikfkekcikdgkabhkFikaffcfkhkdgkegbiaekfkiakicjhfgqdq2fkiakgkfkhfkfcjiekgFebicg"
	suoS += "gbedF1jikejbbbiakgbgkacgiejkijjgigfiakggfggcibFifjefjF1kfekdgjcibFeFkijcfkfhkfkeaieigekgbhkfikidfcje"
	suoS += "aibgekgdkiffiffkiakF1jhbakgdki1dj1ikfkicjicjieeFkgdkicggkighdF1jfgkgfgbdkicggfggkidFkiekgijkeigfiski"
	suoS += "ggfaidheigF1jekijcikickiggkidhhdbgcfkFikikhkigeidieFikggikhkffaffijhidhhakgdkhkijF1kiakF1kfheakgdkif"
	suoS += "iggkigicjiejkieedikgdfcggkigieeiejfgkgkigbgikicggkiaideeijkefjeijikhkiggkiaidheigcikaikffikijgkiahi1"
	suoS += "hhdikgjfifaakekighie1hiaikggikhkffakicjhiahaikggikhkijF1kfejfeFhidikggiffiggkigicjiekgieeigikggiffig"
	suoS += "gkidheigkgfjkeigiegikifiggkidhedeijcfkFikikhkiggkidhh1ehigcikaffkhkiggkidhh1hhigikekfiFkFikcidhh1hit"
	suoS += "cikggikhkfkicjicghiediaikggikhkijbjfejfeFhaikggifikiggkigiejkikgkgieeigikggiffiggkigieeigekijcijikgg"
	suoS += "ifikiggkideedeijkefkfckikhkiggkidhh1ehijcikaffkhkiggkidhh1hhigikhkikFikfckcidhh1hiaikgjikhfjicjicgie"
	suoS += "hdikcikggifikigiejfejkieFhegikggifikiggfghigkfjeijkhigikggifikiggkigieeijcijcikfksikifikiggkidehdeij"
	suoS += "cfdckikhkiggkhghh1ehijikifffffkhsFngErD1pAfBoDd1BlEtFqA2AqoEpDqElAEsEeB2BmADlDkqBtC1FnEpDqnEmFsFsAFn"
	suoS += "llBbFmDsDiCtDmAB2BmtCgpEplCpAEiBiEoFqFtEqsDcCnFtADnFlEgdkEgmEtEsCtDmADqFtAFrAtEcCqAE1BoFqC1F1DrFtBmF"
	suoS += "tAC2ACnFaoCgADcADcCcFfoFtDlAFgmFqBq2bpEoAEmkqnEeCtAE1bAEqgDfFfCrgEcBrACfAAABqAAB1AAClEnFeCtCgAADqDoB"
	suoS += "mtAAACbFiAAADsEtBqAB2FsDqpFqEmFsCeDtFlCeDtoEpClEqAAFrAFoCgFmFsFqEnAEcCqFeCtFtEnAEeFtAAEkFnErAABbFkAD"
	suoS += "nAAeCtFeAfBoAEpFtAABtFqAApDcCGJ"

	//1645-09-23开始7567个节气修正表
	qiS = "FrcFs22AFsckF2tsDtFqEtF1posFdFgiFseFtmelpsEfhkF2anmelpFlF1ikrotcnEqEq2FfqmcDsrFor22FgFrcgDscFs22FgEe"
	qiS += "FtE2sfFs22sCoEsaF2tsD1FpeE2eFsssEciFsFnmelpFcFhkF2tcnEqEpFgkrotcnEqrEtFermcDsrE222FgBmcmr22DaEfnaF22"
	qiS += "2sD1FpeForeF2tssEfiFpEoeFssD1iFstEqFppDgFstcnEqEpFg11FscnEqrAoAF2ClAEsDmDtCtBaDlAFbAEpAAAAAD2FgBiBqo"
	qiS += "BbnBaBoAAAAAAAEgDqAdBqAFrBaBoACdAAf1AACgAAAeBbCamDgEifAE2AABa1C1BgFdiAAACoCeE1ADiEifDaAEqAAFe1AcFbcA"
	qiS += "AAAAF1iFaAAACpACmFmAAAAAAAACrDaAAADG0"

	SB := jieya(suoS) //定朔修正表解压
	QB := jieya(qiS)  //定气修正表解压
	return SB, QB
}

func calc(jd float64, qs string) float64 { //jd应靠近所要取得的气朔日,qs='气'时，算节气的儒略日
	jd += 2451545
	var i int
	var D float64
	var n string
	var B = suoKB
	pc := 14.0
	if qs == "气" {
		B = qiKB
		pc = 7.0
	}
	var f1 = B[0] - pc
	var f2 = B[len(B)-1] - pc
	var f3 = 2436935.0

	if jd < f1 || jd >= f3 { //平气朔表中首个之前，使用现代天文算法。1960.1.1以后，使用现代天文算法 (这一部分调用了qi_high和so_high,所以需星历表支持)
		if qs == "气" {
			return math.Floor(qi_high(math.Floor((jd+pc-2451259)/365.2422*24)*math.Pi/12) + 0.5) //2451259是1999.3.21,太阳视黄经为0,春分.定气计算
		} else {
			return math.Floor(so_high(math.Floor((jd+pc-2451551)/29.5306)*math.Pi*2) + 0.5) //2451551是2000.1.7的那个朔日,黄经差为0.定朔计算
		}
	}

	if jd >= f1 && jd < f2 { //平气或平朔
		for i = 0; i < len(B); i += 2 {
			if jd+pc < B[i+2] {
				break
			}

		}
		D = B[i] + B[i+1]*math.Floor((jd+pc-B[i])/B[i+1])
		D = math.Floor(D + 0.5)
		if D == 1683460 {
			D++
		} //如果使用太初历计算-103年1月24日的朔日,结果得到的是23日,这里修正为24日(实历)。修正后仍不影响-103的无中置闰。如果使用秦汉历，得到的是24日，本行D不会被执行。}
		return D - 2451545
	}

	SB, QB := SSQInit()

	if jd >= f2 && jd < f3 { //定气或定朔
		if qs == "气" {
			D = math.Floor(qi_low(math.Floor((jd+pc-2451259)/365.2422*24)*math.Pi/12) + 0.5) //2451259是1999.3.21,太阳视黄经为0,春分.定气计算
			n = string([]rune(QB)[ int(math.Floor((jd-f2)/365.2422*24)):1 ])                 //找定气修正值
		} else {
			D = math.Floor(so_low(math.Floor((jd+pc-2451551)/29.5306)*math.Pi*2) + 0.5) //2451551是2000.1.7的那个朔日,黄经差为0.定朔计算
			n = string([]rune(SB)[ int(math.Floor((jd-f2)/29.5306)):1 ])                //找定朔修正值
		}
		if n == "1" {
			return D + 1
		}
		if n == "2" {
			return D - 1
		}
		return D
	}
	return 0
}

var ZQ []float64 //中气表,其中.liqiu是节气立秋的儒略日,计算三伏时用到
var HS []float64 //合朔表
var dx []float64 ////各月大小
var ym []int     //各月名称
var leap = 0     //闰月位置
//排月序(生成实际年历),在调用calcY()后得到以下数据
//时间系统全部使用北京时，即使是天象时刻的输出，也是使用北京时
//如果天象的输出不使用北京时，会造成显示混乱，更严重的是无法与古历比对
func calcY(jd float64) { //农历排月序计算,可定出农历,有效范围：两个冬至之间(冬至一 <= d < 冬至二)

	var i int
	var W, w float64

	//该年的气
	W = Int2((jd-355+183)/365.2422)*365.2422 + 355 //355是2000.12冬至,得到较靠近jd的冬至估计值
	if calc(W, "气") > jd {
		W -= 365.2422
	}

	//var ZQpe1,ZQpe2 float64

	for i := 0; i < 25; i++ {
		//25个节气时刻(北京时间),从冬至开始到下一个冬至以后
		ZQ = append(ZQ, calc(W+15.2184*float64(i), "气"))
	}
	//ZQpe1=calc(W-15.2,"气")
	//ZQpe2=calc(W-30.4,"气") //补算二气,确保一年中所有月份的“气”全部被计算在内

	//今年"首朔"的日月黄经差w
	w = calc(ZQ[0], "朔") //求较靠近冬至的朔日
	if w > ZQ[0] {
		w -= 29.53
	}

	//该年所有朔,包含14个月的始末
	for i := 0; i < 15; i++ {
		HS = append(HS,calc(w+29.5306*float64(i), "朔"))
	}

	//月大小
	leap = 0
	for i := 0; i < 14; i++ {

		dx = append(dx, HS[i+1]-HS[i]) //月大小
		ym = append(ym,i)                      //月序初始化
	}

	//-721年至-104年的后九月及月建问题,与朔有关，与气无关
	var YY = Int2((ZQ[0]+10+180)/365.2422) + 2000 //确定年份
	if YY >= -721 && YY <= -104 {
		var ns = make([]int, 12)
		var yy float64
		for i := 0; i < 3; i++ {
			yy = YY + float64(i) - 1
			//颁行历年首, 闰月名称, 月建
			if yy >= -721 {
				ns[i] = int(calc(1457698-float64(J2000)+Int2(0.342+(yy+721)*12.368422)*29.5306, "朔"))
				//ns[i+3] = "十三"
				ns[i+6] = 2 //春秋历,ly为-722.12.17
			}
			if yy >= -479 {
				ns[i] = int(calc(1546083-float64(J2000)+Int2(0.500+(yy+479)*12.368422)*29.5306, "朔"))
				//ns[i+3] = "十三"
				ns[i+6] = 2 //战国历,ly为-480.12.11
			}
			if yy >= -220 {
				ns[i] = int(calc(1640641-float64(J2000)+Int2(0.866+(yy+220)*12.369000)*29.5306, "朔"))
				//ns[i+3] = "后九"
				ns[i+6] = 11 //秦汉历,ly为-221.10.31
			}
		}
		var nn, f1 int
		var nsnn float64
		for i := 0; i < 14; i++ {
			for nn = 2; nn >= 0; nn-- {
				if HS[i] >= float64(ns[nn]) {
					break
				}

			}
			f1 = int(Int2((HS[i] - nsnn + 15) / 29.5306)) //该月积数
			if f1 < 12 {
				ym[i] = (f1 + int(ns[nn+6])) % 12
			} else {
				ym[i] = ns[nn+3]
			}
		}
		return
	}

	//无中气置闰法确定闰月,(气朔结合法,数据源需有冬至开始的的气和朔)
	if HS[13] <= ZQ[24] { //第13月的月末没有超过冬至(不含冬至),说明今年含有13个月
		for i = 1; HS[i+1] > ZQ[2*i] && i < 13; i++ { //在13个月中找第1个没有中气的月份

		}
		leap = i
		for ; i < 14; i++ {
			ym[i]--
		}
	}

	//名称转换(月建别名)
	for i = 0; i < 14; i++ {
		var Dm = HS[i] + float64(J2000)
		var v2 = ym[i]   //Dm初一的儒略日,v2为月建序号
		var mc = v2 % 12 //月建对应的默认月名称：建子十一,建丑十二,建寅为正……
		if Dm >= 1724360 && Dm <= 1729794 {
			mc = (v2 + 1) % 12 //  8.01.15至 23.12.02 建子为十二,其它顺推
		} else if Dm >= 1807724 && Dm <= 1808699 {
			mc = (v2 + 1) % 12 //237.04.12至239.12.13 建子为十二,其它顺推
		} else if Dm >= 1999349 && Dm <= 1999467 {
			mc = (v2 + 2) % 12 //761.12.02至762.03.30 建子为正月,其它顺推
		} else if Dm >= 1973067 && Dm <= 1977052 {
			if v2%12 == 0 {
				mc = 2
			}
			if v2 == 2 {
				mc = 2
			}

		}                //689.12.18至700.11.15 建子为正月,建寅为一月,其它不变

		if Dm == 1729794 || Dm == 1808699 {
			mc = 12
			//mc = "拾贰" //239.12.13及23.12.02均为十二月,为避免两个连续十二月，此处改名
		}

		ym[i] = mc
	}
}

/*********************************
以下是月历表的具体实现方法
*********************************/

/*********************************
=====以下是公历、农历、回历综合日历计算类=====

  Lunar：日历计算物件
  使用条件：事先引用eph.js天文历算文件

  实例创建： var lun = new Lunar();
一、 yueLiCalc(By,Bm)方法
·功能：算出该月每日的详信息
·入口参数：
  By是年(公历)
  Bm是月(公历)
·返回：
  lun.w0= (Bd0 + J2000 +1)%7; //本月第一天的星期
  lun.y  公历年份
  lun.m  公历月分
  lun.d0 月首儒略日数
  lun.dn 月天数
  lun.Ly   干支纪年
  lun.ShX  该年对应的生肖
  lun.nianhao 年号纪年
  lun.lun[] 各日信息(对象),日对象尊守此脚本程序开始的注释中定义

二、yueLiHTML(By,Bm)方法
·功能：算出该月每日的详细信息，并给出HTML月历表
·入口参数：
  By是年(公历)
  Bm是月(公历)
·返回：
  yueLiCalc(By,Bm)返回的信息
  lun.pg0 年份信息
  lun.pg1 月历表
  lun.pg2 月相节气表

**********************************/

//返回公历某一个月的'公农回'三合历
// By是年(公历)  Bm是月(公历)
//func yueLiCalc(By, Bm int) {
//	var month Month
//	//var i,j,k,
//	var Bd0, Bdn int
//	//日历物件初始化
//	var JD Time
//	JD.h = 12
//	JD.m = 0
//	JD.s = 0.1
//	JD.Y = float64(By)
//	JD.M = float64(Bm)
//	JD.D = 1
//	Bd0 = int(Int2(toJD(JD)) - float64(J2000)) //公历月首,中午
//	JD.M++
//	if JD.M > 12 {
//		JD.Y++
//		JD.M = 1
//		Bdn = int(Int2(toJD(JD)) - float64(J2000) - float64(Bd0)) //本月天数(公历)
//	}
//
//	var w0 = math.Mod(float64(Bd0)+J2000+1+7000000, 7) //本月第一天的星期
//	month.y = By                                       //公历年份
//	month.m = Bm                                       //公历月分
//	month.d0 = Bd0
//	month.dn = Bdn
//
//	//所属公历年对应的农历干支纪年
//	var c = By - 1984 + 12000
//	month.Ly = Gan[c%10] + Zhi[c%12] //干支纪年
//	month.ShX = ShX[c%12]            //该年对应的生肖
//	//month.nianhao = getNH(By)
//
//	//var D,w,ob,ob2;
//
//	//提取各日信息
//
//	for i := 0; i < Bdn; i++ {
//		var ob Day
//		ob.d0 = float64(Bd0) + float64(i)               //儒略日,北京时12:00
//		ob.di = i                                       //公历月内日序数
//		ob.y = By                                       //公历年
//		ob.m = Bm                                       //公历月
//		ob.dn = Bdn                                     //公历月天数
//		ob.week0 = int(w0)                              //月首的星期
//		ob.week = int(math.Mod(w0+float64(i), 7))       //当前日的星期
//		ob.weeki = int(Int2((w0 + float64(i)) / 7))     //本日所在的周序号
//		ob.weekN = int(Int2((w0+float64(Bdn)-1)/7) + 1) //本月的总周数
//		setFromJD(float64(ob.d0) + J2000)
//		ob.d = int(JD.D) //公历日名称
//
//		//农历月历
//		if len(ZQ) != 0 || ob.d0 < ZQ[0] || ob.d0 >= ZQ[24] { //如果d0已在计算农历范围内则不再计算
//			calcY(float64(ob.d0))
//		}
//
//		var mk = int(Int2((ob.d0 - HS[0]) / 30))
//		if mk < 13 && HS[mk+1] <= ob.d0 { //农历所在月的序数
//			mk++
//		}
//
//		ob.Ldi = ob.d0 - HS[mk]    //距农历月首的编移量,0对应初一
//		ob.Ldc = rmc[int(ob.Ldi)]  //农历日名称
//		ob.cur_dz = ob.d0 - ZQ[0]  //距冬至的天数
//		ob.cur_xz = ob.d0 - ZQ[12] //距夏至的天数
//		ob.cur_lq = ob.d0 - ZQ[15] //距立秋的天数
//		ob.cur_mz = ob.d0 - ZQ[11] //距芒种的天数
//		ob.cur_xs = ob.d0 - ZQ[13] //距小暑的天数
//		if ob.d0 == HS[mk] || ob.d0 == float64(Bd0) { //月的信息
//			ob.Lmc = ym[mk].(string)                        //月名称
//			ob.Ldn = dx[mk]                                 //月大小
//			ob.Lleap = (SSQ.leap && SSQ.leap == mk)?'闰':''; //闰状况
//			ob.Lmc2 = mk < 13?SSQ.ym[mk+1]:"未知";            //下个月名称,判断除夕时要用到
//		} else {
//			ob2 = this.lun[i-1];
//			ob.Lmc = ob2.Lmc, ob.Ldn  = ob2.Ldn;
//			ob.Lleap = ob2.Lleap, ob.Lmc2 = ob2.Lmc2;
//		}
//		var qk = int2((ob.d0 - SSQ.ZQ[0] - 7) / 15.2184);
//		if (qk < 23 && ob.d0 >= SSQ.ZQ[qk+1]) qk++; //节气的取值范围是0-23
//		if (ob.d0 == SSQ.ZQ[qk]) ob.Ljq = obb.jqmc[qk];
//		else ob.Ljq = '';
//		ob.yxmc = ob.yxjd = ob.yxsj = ''; //月相名称,月相时刻(儒略日),月相时间串
//		ob.jqmc = ob.jqjd = ob.jqsj = ''; //定气名称,节气时刻(儒略日),节气时间串
//
//		//干支纪年处理
//		//以立春为界定年首
//		D = SSQ.ZQ[3] + (ob.d0 < SSQ.ZQ[3]?-365:0) + 365.25*16 - 35; //以立春为界定纪年
//		ob.Lyear = Math.floor(D/365.2422 + 0.5);                     //农历纪年(10进制,1984年起算)
//		//以下几行以正月初一定年首
//		D = SSQ.HS[2]; //一般第3个月为春节
//		for (j = 0; j < 14; j++){ //找春节
//			if (SSQ.ym[j] != '正' || SSQ.leap == j && j)
//			continue;
//			D = SSQ.HS[j];
//			if (ob.d0 < D) {
//				D -= 365;
//				break;
//			} //无需再找下一个正月
//		}
//		D = D + 5810;                             //计算该年春节与1984年平均春节(立春附近)相差天数估计
//		ob.Lyear0 = Math.floor(D/365.2422 + 0.5); //农历纪年(10进制,1984年起算)
//
//		D = ob.Lyear + 12000;
//		ob.Lyear2 = obb.Gan[D%10] + obb.Zhi[D%12]; //干支纪年(立春)
//		D = ob.Lyear0 + 12000;
//		ob.Lyear3 = obb.Gan[D%10] + obb.Zhi[D%12]; //干支纪年(正月)
//		ob.Lyear4 = ob.Lyear0 + 1984 + 2698;       //黄帝纪年
//
//		//纪月处理,1998年12月7(大雪)开始连续进行节气计数,0为甲子
//		mk = int2((ob.d0 - SSQ.ZQ[0]) / 30.43685);
//		if (mk < 12 && ob.d0 >= SSQ.ZQ[2*mk+1]) mk++; //相对大雪的月数计算,mk的取值范围0-12
//
//		D = mk + int2((SSQ.ZQ[12]+390)/365.2422)*12 + 900000; //相对于1998年12月7(大雪)的月数,900000为正数基数
//		ob.Lmonth = D % 12;
//		ob.Lmonth2 = obb.Gan[D%10] + obb.Zhi[D%12];
//
//		//纪日,2000年1月7日起算
//		D = ob.d0 - 6 + 9000000;
//		ob.Lday2 = obb.Gan[D%10] + obb.Zhi[D%12];
//
//		//星座
//		mk = int2((ob.d0 - SSQ.ZQ[0] - 15) / 30.43685);
//		if (mk < 11 && ob.d0 >= SSQ.ZQ[2*mk+2]) mk++; //星座所在月的序数,(如果j=13,ob.d0不会超过第14号中气)
//		ob.XiZ = obb.XiZ[(mk+12)%12] + '座';
//		//回历
//		oba.getHuiLi(ob.d0, ob);
//		//节日
//		ob.A = ob.B = ob.C = '';
//		ob.Fjia = 0;
//		oba.getDayName(ob, ob); //公历
//		obb.getDayName(ob, ob); //农历
//	}
//
//	//以下是月相与节气的处理
//	var d, xn, jd2 = Bd0 + dt_T(Bd0) - 8/24;
//	//月相查找
//	w = XL.MS_aLon(jd2/36525, 10, 3);
//	w = int2((w-0.78)/Math.PI*2) * Math.PI / 2;
//	do{
//		d = obb.so_accurate(w);
//		D = int2(d+0.5);
//		xn = int2(w/pi2*4+4000000.01)%4;
//		w += pi2/4;
//		if (D>=Bd0+Bdn) break;
//		if (D<Bd0) continue;
//		ob = this.lun[D-Bd0];
//		ob.yxmc = obb.yxmc[xn]; //取得月相名称
//		ob.yxjd = d;
//		ob.yxsj = JD.timeStr(d);
//	}
//	while(D+5 < Bd0+Bdn);
//
//	//节气查找
//	w = XL.S_aLon(jd2/36525, 3);
//	w = int2((w-0.13)/pi2*24) * pi2 / 24;
//	do{
//		d = obb.qi_accurate(w);
//		D = int2(d+0.5);
//		xn = int2(w/pi2*24+24000006.01)%24;
//		w += pi2/24;
//		if (D>=Bd0+Bdn) break;
//		if (D<Bd0) continue;
//		ob = this.lun[D-Bd0];
//		ob.jqmc = obb.jqmc[xn]; //取得节气名称
//		ob.jqjd = d;
//		ob.jqsj = JD.timeStr(d);
//	}
//	while(D+12 < Bd0+Bdn);
//}

func GetDayBySolar(_year int, _month int, _day int) Day {
	if _month > 12 || _month <= 0 {

	}

	var num = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if _day <= 0 || _day > num[_month] {
		var year = _year
		if year < 0 {
			year++ //公元前润年
		}
		if !(year%4 == 0 && _month == 2 && _day == 29) || year == 0 { //公元记年法没有0年，公元从1年开始，公元前从-1开始
			//throw LunarException(ErrorCode_DateError)
		}
	}

	var month Month
	var Bd0, Bdn float64

	var t Time
	t.h = 12
	t.m = 0
	t.s = 0.1
	t.Y = float64(_year)
	t.M = float64(_month)
	t.D = 1

	//公历月首的儒略日,中午;
	Bd0 = Int2(toJD(t)) - J2000

	t.M++

	if t.M > 12 {
		t.Y++
		t.M = 1
	}

	//本月天数(公历);
	Bdn = Int2(toJD(t)) - J2000 - Bd0

	//本月第一天的星期;
	var w0 = math.Mod(Bd0+J2000+1+7000000, 7)

	////所属公历年对应的农历干支纪年
	var c = _year - 1984 + 12000

	//年天干地支
	month.yearGan = c % 10
	month.yearZhi = c % 12

	//年生肖
	month.ShX = c % 12

	var i = float64(_day - 1)
	//提取各日信息
	var day Day
	day.d0 = Bd0 + i                   //儒略日,北京时12:00
	day.di = i                         //公历月内日序数
	day.y = _year                      //公历年
	day.m = _month                     //公历月
	day.dn = Bdn                       //公历月天数
	day.week0 = w0                     //月首的星期
	day.week = math.Mod(w0+i, 7)       //当前日的星期
	day.weeki = Int2((w0 + i) / 7)     //本日所在的周序号
	day.weekN = Int2((w0+Bdn-1)/7) + 1 //本月的总周数

	time := DD(day.d0 + J2000) //儒略日数转公历
	day.d = time.D             //公历日名称

	//如果d0已在计算农历范围内则不再计算
	if len(ZQ) == 0 || day.d0 < ZQ[0] || day.d0 >= ZQ[24] {
		calcY(day.d0)
	}

	var mk = int(Int2((day.d0 - HS[0]) / 30))
	//农历所在月的序数;
	if mk < 13 && HS[mk+1] <= day.d0 {
		mk++
	}

	day.Ldi = day.d0 - HS[mk] //距农历月首的编移量,0对应初一

	day.cur_dz = day.d0 - ZQ[0]  //距冬至的天数
	day.cur_xz = day.d0 - ZQ[12] //距夏至的天数
	day.cur_lq = day.d0 - ZQ[15] //距立秋的天数
	day.cur_mz = day.d0 - ZQ[11] //距芒种的天数
	day.cur_xs = day.d0 - ZQ[13] //距小暑的天数

	if day.d0 == HS[mk] || day.d0 == Bd0 { //月的信息
		day.Lmc = ym[mk]                    //月名称
		day.Ldn = dx[mk]                    //月大小
		day.Lleap = leap != 0 && leap == mk //闰状况
		//day.Lmc2 = mk < 13 ? mk + 1 : -1; //下个月名称,判断除夕时要用到
		if mk < 13 {
			day.Lmc2 = mk + 1 //下个月名称,判断除夕时要用到
		} else {
			day.Lmc2 = -1
		}

	} else {
		var day2 = GetDayBySolar(_year, _month, _day-1)
		day.Lmc = day2.Lmc
		day.Ldn = day2.Ldn
		day.Lleap = day2.Lleap
		day.Lmc2 = day2.Lmc2
	}
	var qk = int(Int2((day.d0 - ZQ[0] - 7) / 15.2184))

	//节气的取值范围是0-23
	if qk < 23 && day.d0 >= ZQ[qk+1] {
		qk++
	}

	day.qk = -1
	if day.d0 == ZQ[qk] {
		day.qk = qk
	}

	//day.yxmc = day.yxjd = day.yxsj = '';//月相名称,月相时刻(儒略日),月相时间串
	//day.jqmc = day.jqjd = day.jqsj = '';//定气名称,节气时刻(儒略日),节气时间串

	//干支纪年处理
	//以立春为界定年首

	if day.d0 < ZQ[3] {
		day.d0 = -365
	} else {
		day.d0 = 0
	}

	var D = ZQ[3] + day.d0 + 365.25*16 - 35           //以立春为界定纪年
	day.Lyear = math.Floor(float64(D/365.2422) + 0.5) //农历纪年(10进制,1984年起算)
	//以下几行以正月初一定年首
	D = HS[2] //一般第3个月为春节
	for j := 0; j < 14; j++ { //找春节
		//正月的index = 3
		if ym[j] != 2 {
			continue
		}
		D = HS[j]
		if day.d0 < D {
			D -= 365
			break
		} //无需再找下一个正月
	}
	D = D + 5810                               //计算该年春节与1984年平均春节(立春附近)相差天数估计
	day.Lyear0 = math.Floor(D/365.2422 + 0.5); //农历纪年(10进制,1984年起算)

	//干支纪年(立春)
	D = day.Lyear + 12000
	day.Lyear2.tg = int(D) % 10
	day.Lyear2.dz = int(D) % 12

	//干支纪年(正月)
	D = day.Lyear0 + 12000
	day.Lyear3.tg = int(D) % 10
	day.Lyear3.dz = int(D) % 12

	//干支纪年(正月)
	day.Lyear4 = day.Lyear0 + 1984 + 2698 //黄帝纪年

	//纪月处理,1998年12月7(大雪)开始连续进行节气计数,0为甲子
	//一年固定有12个中气日，平均每个中气日间隔30.43685天
	mk = int(Int2((day.d0 - ZQ[0]) / 30.43685))
	//相对大雪的月数计算,mk的取值范围0-12
	if mk < 12 && day.d0 >= ZQ[2*mk+1] {
		mk++
	}

	D = float64(mk) + Int2((ZQ[12]+390)/365.2422)*12 + 900000 //相对于1998年12月7(大雪)的月数,900000为正数基数
	day.Lmonth = int(D) % 12
	day.Lmonth2.tg = int(D) % 10
	day.Lmonth2.dz = int(D) % 12

	////纪日,2000年1月7日起算
	D = day.d0 - 6 + 9000000
	day.Lday2.tg = int(D) % 10
	day.Lday2.dz = int(D) % 12

	//星座
	mk = int(Int2((day.d0 - ZQ[0] - 15) / 30.43685))
	//星座所在月的序数,(如果j=13,day.d0不会超过第14号中气)
	if mk < 11 && day.d0 >= ZQ[2*mk+2] {
		mk++
	}
	day.XiZ = (mk + 12) % 12

	//以下是月相与节气的处理
	var d, xn float64
	var jd2 = Bd0 + dt_T(Bd0) - (8.0 / 24.0)
	//月相查找
	var w = MS_aLon(jd2/36525, 10, 3)
	w = Int2((w-0.78)/math.Pi*2) * math.Pi / 2

	for {
		d = so_accurate(w)
		D = Int2(d + 0.5)
		xn = math.Mod(Int2(w/pi2*4+4000000.01), 4)
		w += pi2 / 4
		if D >= Bd0+Bdn {
			break
		}
		if D < Bd0 {
			continue
		}
		if int(D-Bd0) == _day-1 {
			day.yxmc = xn //取得月相名称
			day.yxjd = d
			day.yxsj = timeStr(d)
			break
		}
		if D+5 < Bd0+Bdn {
			break
		}
	}

	//节气查找
	w = S_aLon(jd2/36525, 3)
	w = Int2((w-0.13)/pi2*24) * pi2 / 24

	for {
		d = qi_accurate(w)
		D = Int2(d + 0.5)
		xn = math.Mod(w/pi2*24+24000006.01, 24)
		w += pi2 / 24
		if D >= Bd0+Bdn {
			break
		}

		if D < Bd0 {
			continue
		}

		if int(D-Bd0) == _day-1 {
			day.jqmc = xn //取得节气名称
			day.jqjd = d
			day.jqsj = timeStr(d)
			break
			if D+12 < Bd0+Bdn {
				break
			}
		}
	}

	//获取准确节气的时间
	jd2 = ZQ[0] + dt_T(ZQ[0]) - (8.0 / 24.0)
	w = S_aLon(jd2/36525, 3)
	w = Int2((w-0.13)/pi2*24) * pi2 / 24

	//for ( it := 0; it != len(ZQ); it++){
	//while (true)
	//{
	//d = qi_accurate(w);
	//D = int2(d + 0.5);
	//xn = int2(w / pi2 * 24 + 24000006.01) % 24;
	//w += pi2 / 24;
	//if (D < *it) continue;
	//break;
	//
	//}
	//Time t1 = JD::JD2DD(d);
	//Time t2 = JD::JD2DD(*it);
	//
	//t2.h = t1.h;
	//t2.m = t1.m;
	//t2.s = t1.s;
	//
	//
	//auto jd = JD::toJD(t2);
	//Time t3 = JD::JD2DD(jd + J2000);
	//
	//day.cur_jq.push_back(jd);
	//day.cur_cn.push_back((int)xn);
	//}

	return day
}