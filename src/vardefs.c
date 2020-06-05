#include "client/shell/wolapi/wolchat.h"
#include "client/system/parsecmd.h"
#include "common/object/modifier.h"
#include "server/object/die/die.h"
#include "server/object/pickdrop/pickup.h"
#include "server/system/cscrfunc.h"

#include "client/gui/gamewin/psscript.h"
#include "common/wolapi/wolstate.h"

#include "client/drawable/update/dball.h"
#include "client/drawable/update/drainup.h"
#include "client/drawable/update/healup.h"
#include "client/drawable/update/manabomb.h"
#include "client/drawable/update/mmislup.h"
#include "client/drawable/update/mtailup.h"
#include "client/drawable/update/sparklup.h"
#include "client/drawable/update/telwake.h"
#include "client/drawable/update/vortexup.h"

#include "proto.h"

void sub_57F634();
void nullsub_45();
void nullsub_46();
void nullsub_47();
void nullsub_48();
void nullsub_62();
void nullsub_65();
void nullsub_66();
void nullsub_67();
void nullsub_68();
void nullsub_69();
void nullsub_70();
_DWORD dword_5d4594_3798784 = 0;
_DWORD dword_5d4594_1599636 = 0;
_DWORD dword_5d4594_3799540 = 0;
_DWORD dword_5d4594_4668 = 0;
_DWORD dword_5d4594_1046492 = 0;
_DWORD dword_587000_133480 = &byte_5D4594[1049220];
_DWORD dword_5d4594_830224 = 0;
_DWORD dword_5d4594_814980 = 0;
_DWORD dword_5d4594_3804664 = 0;
_DWORD dword_5d4594_1045516 = 0;
_DWORD dword_5d4594_3804656 = 0;
_DWORD dword_5d4594_3804672 = 0;
_DWORD dword_5d4594_1045684 = 0;
_DWORD dword_5d4594_1307784 = 0;
_DWORD dword_5d4594_3798844 = 0;
_DWORD dword_5d4594_2650652 = 0;
_DWORD dword_5d4594_1217448 = 0;
_DWORD dword_5d4594_1320940 = 0;
_DWORD dword_5d4594_815004 = 0;
_DWORD dword_5d4594_815012 = 0;
_DWORD dword_5d4594_1062480 = 0;
_DWORD dword_5d4594_832484 = 0;
_DWORD dword_5d4594_3798796 = 0;
_DWORD dword_5d4594_825760 = 0;
_DWORD dword_587000_80828 = 0x1;
_DWORD dword_5d4594_810632 = 0;
_DWORD dword_5d4594_787148 = 0;
_DWORD dword_5d4594_2386944 = 0;
_DWORD dword_5d4594_1064916 = 0;
_DWORD dword_5d4594_3801780 = 0;
_DWORD dword_5d4594_2516348 = 0;
_DWORD dword_5d4594_3799536 = 0;
_DWORD dword_5d4594_1063116 = 0;
_DWORD dword_5d4594_1090120 = 0;
_DWORD dword_5d4594_2516328 = 0;
_DWORD dword_5d4594_1189584 = 0;
_DWORD dword_5d4594_1049504 = 0;
_DWORD dword_5d4594_1046868 = 0;
_DWORD dword_5d4594_1082856 = 0;
_DWORD dword_5d4594_1193380 = 0;
_DWORD dword_587000_26048 = 0x9;
_DWORD dword_5d4594_3799624 = 0;
_DWORD dword_5d4594_1316972 = 0;
_DWORD dword_5d4594_1123524 = 0;
_DWORD dword_5d4594_1064192 = 0;
_DWORD dword_5d4594_1045604 = 0;
_DWORD dword_5d4594_2386176 = 0;
_DWORD dword_587000_197424 = 0xffffffff;
_DWORD dword_5d4594_1064816 = 0;
_DWORD dword_5d4594_529316 = 0;
_DWORD dword_5d4594_1309820 = 0;
_DWORD dword_5d4594_1062508 = 0;
_DWORD dword_5d4594_1049864 = 0;
_DWORD dword_5d4594_829480 = 0;
_DWORD dword_5d4594_1062512 = 0;
_DWORD dword_5d4594_1217464 = 0;
_DWORD dword_5d4594_830248 = 0;
_DWORD dword_5d4594_1599656 = 0;
_DWORD dword_5d4594_1049500 = 0;
_DWORD dword_5d4594_3798840 = 0;
_DWORD dword_5d4594_1309812 = 0;
_DWORD dword_5d4594_754036 = 0;
_DWORD dword_587000_155144 = &byte_5D4594[1193204];
_DWORD dword_5d4594_741676 = 0;
_DWORD dword_5d4594_1047516 = 0;
_DWORD dword_5d4594_1308084 = 0;
_DWORD dword_5d4594_815000 = 0;
_DWORD dword_5d4594_3798824 = 0;
_DWORD dword_5d4594_1307780 = 0;
_DWORD dword_5d4594_3798820 = 0;
_DWORD dword_5d4594_1193516 = 0;
_DWORD dword_5d4594_1090048 = 0;
_DWORD dword_5d4594_816364 = 0;
_DWORD dword_5d4594_1049508 = 0;
_DWORD dword_5d4594_1049796 = 0;
_DWORD dword_5d4594_829492 = 0;
_DWORD dword_5d4594_1098576 = 0;
_DWORD dword_5d4594_1321040 = 0;
_DWORD dword_5d4594_2487236 = 0;
_DWORD dword_5d4594_3799432 = 0;
_DWORD dword_5d4594_830124 = 0;
_DWORD dword_587000_87412 = 0xffffffff;
_DWORD dword_5d4594_1548476 = 0;
_DWORD dword_5d4594_2660652 = 0;
_DWORD dword_5d4594_1045636 = 0;
_DWORD dword_587000_87404 = 0x2;
_DWORD dword_5d4594_3798804 = 0;
_DWORD dword_5d4594_832476 = 0;
_DWORD dword_5d4594_1570272 = 0;
_DWORD dword_5d4594_1305684 = 0;
_DWORD dword_5d4594_2386940 = 0;
_DWORD dword_5d4594_2523752 = 0;
_DWORD dword_5d4594_815132 = 0;
_DWORD dword_5d4594_1321228 = 0;
_DWORD dword_5d4594_815048 = 0;
_DWORD dword_5d4594_829488 = 0;
_DWORD dword_5d4594_1063636 = 0;
_DWORD dword_587000_197428 = 0xffffffff;
_DWORD dword_587000_60352 = 0x1;
_DWORD dword_5d4594_833716 = 0;
_DWORD dword_5d4594_1319228 = 0;
_DWORD dword_5d4594_10956 = 0;
_DWORD dword_5d4594_1062456 = 0;
_DWORD dword_5d4594_816348 = 0;
_DWORD dword_5d4594_2650672 = 0;
_DWORD dword_5d4594_3798696 = 0;
_DWORD dword_5d4594_2523756 = 0;
_DWORD dword_5d4594_1307764 = 0;
_DWORD dword_5d4594_1098628 = 0;
_DWORD dword_5d4594_529340 = 0;
_DWORD dword_587000_149232 = 0x456;
_DWORD dword_5d4594_1046872 = 0;
_DWORD dword_5d4594_1046496 = 0;
_DWORD dword_5d4594_830108 = 0;
_DWORD dword_5d4594_1523028 = 0;
_DWORD dword_5d4594_1316708 = 0;
_DWORD dword_5d4594_1064888 = 0;
_DWORD dword_5d4594_3798636 = 0;
_DWORD dword_5d4594_251572 = 0;
_DWORD dword_5d4594_1556860 = 0;
_DWORD dword_5d4594_1303452 = 0;
_DWORD dword_5d4594_2386224 = 0;
_DWORD dword_5d4594_815016 = 0;
_DWORD dword_5d4594_2386152 = 0;
_DWORD dword_5d4594_1311144 = 0;
_DWORD dword_5d4594_1569728 = 0;
_DWORD dword_5d4594_1197312 = 0;
_DWORD dword_5d4594_831236 = 0;
_DWORD dword_5d4594_3799484 = 0;
_DWORD dword_5d4594_1064296 = 0;
_DWORD dword_5d4594_1047540 = 0;
_DWORD dword_5d4594_814984 = 0;
_DWORD dword_5d4594_1320964 = 0;
_DWORD dword_5d4594_2523900 = 0;
_DWORD dword_5d4594_1049800 = 0;
_DWORD dword_5d4594_251568 = 0;
_DWORD dword_5d4594_1309720 = 0;
_DWORD dword_5d4594_1569752 = 0;
_DWORD dword_5d4594_1321244 = 0;
_DWORD dword_5d4594_1046936 = 0;
_DWORD dword_5d4594_2487908 = 0;
_DWORD dword_5d4594_831088 = 0;
_DWORD dword_5d4594_1107036 = 0;
_DWORD dword_5d4594_1321248 = 0;
_DWORD dword_5d4594_1599640 = 0;
_DWORD dword_5d4594_3798836 = 0;
_DWORD dword_5d4594_1565512 = 0;
_DWORD dword_5d4594_815028 = 0;
_DWORD dword_5d4594_816456 = 0;
_DWORD dword_5d4594_1301848 = 0;
_DWORD dword_5d4594_829508 = 0;
_DWORD dword_5d4594_1316712 = 0;
_DWORD dword_5d4594_1313800 = 0;
_DWORD dword_5d4594_1321232 = 0;
_DWORD dword_5d4594_815024 = 0;
_DWORD dword_5d4594_830104 = 0;
_DWORD dword_587000_80820 = 0x1;
_DWORD dword_5d4594_1045528 = 0;
_DWORD dword_5d4594_787156 = 0;
_DWORD dword_5d4594_1301792 = 0;
_DWORD dword_587000_154956 = 0x1;
_DWORD dword_5d4594_1313796 = 0;
_DWORD dword_5d4594_1522624 = 0;
_DWORD dword_5d4594_815020 = 0;
_DWORD dword_5d4594_1305680 = 0;
_DWORD dword_5d4594_810640 = 0;
_DWORD dword_5d4594_815032 = 0;
_DWORD dword_5d4594_1311148 = 0;
_DWORD dword_5d4594_787164 = 0;
_DWORD dword_5d4594_816440 = 0;
_DWORD dword_5d4594_1301800 = 0;
_DWORD dword_5d4594_1193360 = 0;
_DWORD dword_5d4594_1522628 = 0;
_DWORD dword_5d4594_3679320 = 0;
_DWORD dword_5d4594_1049520 = 0;
_DWORD dword_5d4594_1309748 = 0;
_DWORD dword_5d4594_816432 = 0;
_DWORD dword_5d4594_1599540 = 0;
_DWORD dword_5d4594_251548 = 0;
_DWORD dword_5d4594_1556844 = 0;
_DWORD dword_5d4594_1045464 = 0;
_DWORD dword_5d4594_1045484 = 0;
_DWORD dword_5d4594_1049856 = 0;
_DWORD dword_5d4594_1064856 = 0;
_DWORD dword_5d4594_823692 = 0;
_DWORD dword_5d4594_3799524 = 0;
_DWORD dword_5d4594_1522600 = 0;
_DWORD dword_587000_132132 = 0x1;
_DWORD dword_5d4594_754056 = 0;
_DWORD dword_5d4594_1217456 = 0;
_DWORD dword_5d4594_1193348 = 0;
_DWORD dword_5d4594_1047520 = 0;
_DWORD dword_5d4594_3798632 = 0;
_DWORD dword_5d4594_534808 = 0;
_DWORD dword_5d4594_1197324 = 0;
_DWORD dword_5d4594_3821640 = 0;
_DWORD dword_5d4594_1599576 = 0;
_DWORD dword_5d4594_3821636 = 0;
_DWORD dword_5d4594_1599596 = 0;
_DWORD dword_5d4594_1045480 = 0;
_DWORD dword_5d4594_3807156 = 0;
_DWORD dword_5d4594_816452 = 0;
_DWORD dword_5d4594_2487556 = 0;
_DWORD dword_5d4594_833712 = 0;
_DWORD dword_5d4594_1307776 = 0;
_DWORD dword_5d4594_1599588 = 0;
_DWORD dword_5d4594_1064912 = 0;
_DWORD dword_5d4594_3798640 = 0;
_DWORD dword_5d4594_1523024 = 0;
_DWORD dword_5d4594_2618912 = 0;
_DWORD dword_5d4594_2516344 = 0;
_DWORD dword_5d4594_1046508 = 0;
_DWORD dword_5d4594_1123532 = 0;
_DWORD dword_5d4594_1193384 = 0;
_DWORD dword_5d4594_1045536 = 0;
_DWORD dword_587000_166704 = 0x6;
_DWORD dword_5d4594_1319256 = 0;
_DWORD dword_5d4594_251540 = 0;
_DWORD dword_5d4594_371692 = 0;
_DWORD dword_5d4594_3798832 = 0;
_DWORD dword_5d4594_1197320 = 0;
_DWORD dword_5d4594_3799492 = 0;
_DWORD dword_5d4594_1301844 = 0;
_DWORD dword_5d4594_3798808 = 0;
_DWORD dword_5d4594_1062524 = 0;
_DWORD dword_5d4594_3835396 = 0;
_DWORD dword_5d4594_1062528 = 0;
_DWORD dword_5d4594_2388648 = 0;
_DWORD dword_587000_154952 = 0x1;
_DWORD dword_5d4594_1312480 = 0;
_DWORD dword_5d4594_1522604 = 0;
_DWORD dword_5d4594_1096640 = 0;
_DWORD dword_5d4594_3798816 = 0;
_DWORD dword_5d4594_1045532 = 0;
_DWORD dword_5d4594_528256 = 0;
_DWORD dword_5d4594_1568024 = 0;
_DWORD dword_5d4594_1599564 = 0;
_DWORD dword_5d4594_1308088 = 0;
_DWORD dword_5d4594_3801804 = 0;
_DWORD dword_5d4594_816492 = 0;
_DWORD dword_5d4594_826028 = 0;
_DWORD dword_5d4594_2487524 = 0;
_DWORD dword_5d4594_1189592 = 0;
_DWORD dword_5d4594_1522612 = 0;
_DWORD dword_5d4594_814624 = 0;
_DWORD dword_5d4594_1599660 = 0;
_DWORD dword_5d4594_829484 = 0;
_DWORD dword_5d4594_3821964 = 0;
_DWORD dword_5d4594_1096516 = 0;
_DWORD dword_5d4594_833720 = 0;
_DWORD dword_5d4594_1045468 = 0;
_DWORD dword_5d4594_2487904 = 0;
_DWORD dword_587000_127004 = &byte_5D4594[1045324];
_DWORD dword_5d4594_2386184 = 0;
_DWORD dword_5d4594_815088 = 0;
_DWORD dword_5d4594_2495920 = 0;
_DWORD dword_5d4594_1046924 = 0;
_DWORD dword_5d4594_1312484 = 0;
_DWORD dword_5d4594_1189596 = 0;
_DWORD dword_5d4594_3798156 = 0;
_DWORD dword_587000_60356 = 0x1;
_DWORD dword_5d4594_3624 = 0;
_DWORD dword_5d4594_1217460 = 0;
_DWORD dword_5d4594_1307716 = 0;
_DWORD dword_587000_93156 = 0x1;
_DWORD dword_5d4594_830140 = 0;
_DWORD dword_5d4594_2523780 = 0;
_DWORD dword_5d4594_2650676 = 0;
_DWORD dword_5d4594_1321252 = 0;
_DWORD dword_5d4594_1522632 = 0;
_DWORD dword_5d4594_531648 = 0;
_DWORD dword_5d4594_787208 = 0;
_DWORD dword_5d4594_1045432 = 0;
_DWORD dword_5d4594_815044 = 0;
_DWORD dword_5d4594_831232 = 0;
_DWORD dword_5d4594_787152 = 0;
_DWORD dword_5d4594_816488 = 0;
_DWORD dword_5d4594_2487712 = 0;
_DWORD dword_5d4594_2386496 = 0;
_DWORD dword_5d4594_529328 = 0;
_DWORD dword_5d4594_2386180 = 0;
_DWORD dword_5d4594_3843628 = 0;
_DWORD dword_5d4594_2386212 = 0;
_DWORD dword_5d4594_3798720 = 0;
_DWORD dword_5d4594_1319060 = 0;
_DWORD dword_5d4594_1047548 = 0;
_DWORD dword_5d4594_816340 = 0;
_DWORD dword_5d4594_1049512 = 0;
_DWORD dword_5d4594_1311140 = 0;
_DWORD dword_5d4594_3799544 = 0;
_DWORD dword_587000_264364 = 0x41b80000;
_DWORD dword_5d4594_1047932 = 0;
_DWORD dword_5d4594_2523912 = 0;
_DWORD dword_5d4594_816368 = 0;
_DWORD dword_5d4594_1062476 = 0;
_DWORD dword_5d4594_1090276 = 0;
_DWORD dword_5d4594_830116 = 0;
_DWORD dword_5d4594_1049496 = 0;
_DWORD dword_587000_93164 = &byte_5D4594[816244];
_DWORD dword_5d4594_1098624 = 0;
_DWORD dword_5d4594_1569672 = 0;
_DWORD dword_5d4594_3632 = 0;
_DWORD dword_5d4594_1321240 = 0;
_DWORD dword_5d4594_1197316 = 0;
_DWORD dword_5d4594_1599624 = 0;
_DWORD dword_5d4594_2386576 = 0;
_DWORD dword_5d4594_830148 = 0;
_DWORD dword_5d4594_1569756 = 0;
_DWORD dword_5d4594_1320968 = 0;
_DWORD dword_5d4594_2488604 = 0;
_QWORD qword_581450_9544 = 0x3fe0000000000000;
_QWORD qword_581450_9568 = 0x0;
_QWORD qword_5d4594_1567940 = 0;
_QWORD qword_581450_10256 = 0x3f847ae147ae147b;
_QWORD qword_581450_10392 = 0xbf847ae147ae147b;
_QWORD qword_581450_9552 = 0x40f0000000000000;
_QWORD qword_581450_9512 = 0x3ff0000000000000;
_QWORD qword_5d4594_815068 = 0;
_QWORD qword_5d4594_814956 = 0;
_QWORD qword_581450_10176 = 0x3fb99999a0000000;
_QWORD qword_5d4594_815724 = 0;
_DWORD dword_5d4594_1316408 = 0;
_DWORD dword_5d4594_1522620 = 0;
_DWORD dword_5d4594_2488660 = 0;
_DWORD dword_587000_292492 = 0x42c80000;
_DWORD dword_5d4594_805844 = 0;
_DWORD dword_5d4594_251608 = 0;
_DWORD dword_5d4594_1309732 = 0;
_DWORD dword_5d4594_251600 = 0;
_DWORD dword_5d4594_1321204 = 0;
_DWORD dword_5d4594_1599632 = 0;
_DWORD dword_5d4594_2386500 = 0;
_DWORD dword_5d4594_2649712 = 0;
_DWORD dword_5d4594_1307204 = 0;
_DWORD dword_5d4594_741284 = 0;
_DWORD dword_587000_312820 = 0x3f800000;
_DWORD dword_5d4594_1046532 = 0;
_DWORD dword_5d4594_3798828 = 0;
_DWORD dword_5d4594_1309728 = 0;
_DWORD dword_5d4594_1312460 = 0;
_DWORD dword_5d4594_1309836 = 0;
_DWORD dword_5d4594_1320992 = 0;
_DWORD dword_5d4594_1599476 = 0;
