#ifndef NOX_PORT_GAME5_2
#define NOX_PORT_GAME5_2

#include "defs.h"


void nullsub_2();
int __thiscall sub_42CC50(LPVOID* this);
void nullsub_33(_DWORD a1, _DWORD a2, _DWORD a3);
void nullsub_3(void);
BOOL nox_xxx_testCPUID2_444D90();
int sub_448640(void);
void nullsub_4(_DWORD a1, _DWORD a2, _DWORD a3, _DWORD a4);
int nox_xxx_j_inventoryNameSignInit_467460(void);
void nullsub_5(void);
void nullsub_6(void);
int nullsub_8(int a1, int a2);
void nullsub_9(_DWORD a1);
void nullsub_10(_DWORD a1);
void nullsub_11(void);
void nox_xxx_j_resetNPCRenderData_49A2E0(void);
void nox_alloc_npcs_2();
void nullsub_12(void);
void nullsub_13(void);
void nullsub_14(void);
void nullsub_15(void);
void nullsub_16(void);
void nullsub_17(void);
void nullsub_18(void);
void nullsub_19(void);
void nullsub_20(void);
void nullsub_22(void);
void nullsub_36(void);
void nullsub_23(void);
void nullsub_35(_DWORD a1, _DWORD a2);
void nullsub_24(_DWORD a1);
void nullsub_34(_DWORD a1, _DWORD a2, _DWORD a3);
void nox_xxx_set_sage(_DWORD a1);
void nox_xxx_j_allocHitArray_511840(void);
void nullsub_27(_DWORD a1);
void nullsub_28(_DWORD a1);
void nullsub_29(void);
void nullsub_30(_DWORD a1);
int nox_xxx_netSendReadPacket_5528B0(unsigned int a1, char a2);
int nox_xxx_servNetInitialPackets_552A80(unsigned int id, char flags);
int sub_552E70(unsigned int a1);
int sub_552F20(unsigned int a1);
unsigned int sub_552FD0(int a1);
int nox_xxx_netBigSwitch_553210(unsigned int id, unsigned char* packet, int packetSz, int a4);
int sub_553D10();
int sub_553D30(int a1);
int sub_553D80();
int sub_553D90();
int sub_553DB0();
int sub_553DA0();
__int64 sub_553DC0(int a1);
unsigned int sub_553E10(int a1);
unsigned int sub_553E90(int a1);
unsigned int sub_553E50(int a1);
int sub_553ED0(int a3);
void sub_553F40(int a1, int a2);
void sub_553FC0(int a1, int a2);
void nox_xxx_netCountData_554030(int a1, int a2);
unsigned int nox_server_makeServerInfoPacket_554040(const char* inBuf, int inSz, char* out);
int* nox_xxx_findPlayerID_5541D0(int a1);
int nox_xxx_net_getIP_554200(unsigned int a1);
char* sub_554230();
int sub_554240(int a1);
unsigned int sub_554290();
int sub_554300();
int nox_xxx_netInit_554380(nox_net_struct_arg_t* narg);
__int16 sub_5545A0();
int nox_xxx_netStructReadPackets_5545B0(unsigned int a1);
int nox_server_netClose_5546A0(unsigned int a1);
int nox_xxx_netPreStructToFull_5546F0(nox_net_struct_arg_t* narg);
int sub_554760(int a1, char* cp, int hostshort, int a4, int a5);
int sub_5549F0(unsigned int a1);
int sub_554A50(unsigned int a1);
void nox_xxx_lobbyMakePacket_554AA0(uint16_t hostshort, const char* payload, int payloadSz, unsigned int ticks);
int nox_xxx_createSocketLocal_554B40(uint16_t hostshort);
int sub_554D10();
int nox_xxx_sendLobbyPacket_554C80(uint16_t hostshort, char* buf, int a3);
int sub_554D70(char a1);
int sub_554FF0();
int nox_client_sendToServer_555010(unsigned int addr, uint16_t port, char* buf, int sz);
int nox_client_sendJoinGame_5550A0(unsigned int addr, uint16_t port, char* data);
int sub_5550D0(int a1, uint16_t hostshort, char* buf);
int sub_555130(unsigned int a1, const void* a2, signed int a3);
int sub_5551F0(unsigned int a1, char a2, int a3);
int sub_555250(unsigned int a1, _DWORD* a2);
int sub_555290(unsigned int a1, _DWORD* a2);
int nox_xxx_netSend_5552D0(unsigned int a1, char a2, int a3);
int sub_555360(unsigned int a1, unsigned __int8 a2, int a3);
void nullsub_31(_DWORD a1);
int* sub_565360(int a1, _WORD* a2, int* a3, unsigned int a4, int a5, int a6);
unsigned int sub_5654A0(int a1, unsigned __int8* a2, int* a3, unsigned int a4, int a5, int a6);
int sub_56F1C0();
int nox_xxx_protect_56F240();
int sub_56F250();
int nox_xxx_protectionCreateStructForInt_56F280(int a1, int a2);
int sub_56F2F0(_DWORD* a1);
_DWORD* sub_56F3B0();
int nox_xxx_protectionCreateInt_56F400(int a1);
int nox_xxx_protectionCreateFloat_56F440(int a1);
int nox_xxx_protectionCreateStructForFloat_56F480(int a1, int a2);
int sub_56F4F0(int* a1);
int sub_56F510(int a1);
_DWORD* sub_56F590(int a1);
int nox_xxx_protectData_56F5C0();
_DWORD* sub_56F6F0(int a1);
void sub_56F720(int* a1, int* a2);
_DWORD* sub_56F780(int a1, int a2);
_DWORD* nox_xxx_playerUpdateNetBuffs_56F7D0(int a1, int a2);
_DWORD* sub_56F820(int a1, unsigned __int8 a2);
_DWORD* nox_xxx_protectPlayerHPMana_56F870(int a1, unsigned __int16 a2);
_DWORD* sub_56F8C0(int a1, float a2);
_DWORD* sub_56F920(int a1, int a2);
_DWORD* nox_xxx_protectMana_56F9E0(int a1, __int16 a2);
_DWORD* sub_56FA40(int a1, float a2);
int nox_xxx_protectionStringCRC_56FAC0(int* a1, unsigned int a2);
int nox_xxx_protectionStringCRCLen_56FAE0(int* a1, unsigned int a2);
int sub_56FB00(int* a1, unsigned int a2, int a3);
int sub_56FB60(nox_object_t* item);
int nox_xxx_protect_56FBF0(int a1, nox_object_t* item);
int nox_xxx_protect_56FC50(int a1, const nox_object_t* object);
int sub_56FCB0(int a1, int a2);
int nox_xxx_playerAwardSpellProtection_56FCE0(int a1, int a2, int a3);
int sub_56FD50(int a1, void* a2, int a3);
void nox_xxx_cryptXor_56FDD0(char key, unsigned char* p, int n);
void nox_xxx_cryptXorDst_56FE00(char key, unsigned char* src, int n, unsigned char* dst);
double nox_xxx_unkDoubleSmth_56FE30();
void sub_56FF00(int a1);
int sub_56FF80(int a1, int a2);
unsigned int nox_xxx_netGetUnitCodeServ_578AC0(nox_object_t* item);
unsigned int nox_xxx_netGetUnitCodeCli_578B00(int a1);
int nox_xxx_netClearHighBit_578B30(__int16 a1);
int nox_xxx_packetDynamicUnitCode_578B40(int a1);
unsigned int nox_xxx_netTestHighBit_578B70(unsigned int a1);
_DWORD* sub_578B80();
unsigned int sub_578BA0(unsigned int a1);
int sub_578BB0(void** a1, int a2, unsigned __int8* a3, int a4);
void sub_578BD0(LPVOID lpMem);
_DWORD* sub_578BF0();
int nox_xxx_nxzDecompress_578C10(_DWORD* a1, _BYTE* a2, _DWORD* a3, unsigned int a4, _DWORD* a5);
int sub_578C30(int a1);
void sub_578C40(LPVOID lpMem);
int sub_578C60();
int sub_578C90(int a1);
int sub_578CD0();
char* nox_xxx_GetEndgameDialog_578D80();
char sub_578DE0(char a1);
unsigned __int8 sub_578DF0();
void sub_578E00();
LPVOID nox_xxx_waypointGetList_579860();
int nox_xxx_waypointNext_579870(int a1);
int sub_579890();
int sub_5798A0(int a1);
unsigned int nox_xxx_waypoint_5798C0();
_DWORD* nox_xxx_waypointNew_5798F0(float a1, float a2);
float* nox_xxx_waypointNewNotMap_579970(int a1, float a2, float a3);
char* nox_xxx_waypoint_5799C0();
char* sub_579A30();
float* sub_579AD0(float a1, float a2);
void sub_579B30(LPVOID lpMem);
_DWORD* sub_579C00();
_DWORD* nox_server_getWaypointById_579C40(int a1);
int sub_579C60(int a1);
_DWORD* sub_579C80(int a1);
int sub_579CA0();
int sub_579D20();
void nox_xxx_waypointDeleteAll_579DD0();
const char* nox_xxx_waypointByName_579E30(const char* a1);
_DWORD* sub_579E70();
void sub_579E90(int a1);
BOOL sub_579EE0(int a1, unsigned __int8 a2);
int nox_xxx_waypoint_579F00(_DWORD* a1, int a2);
int nox_xxx_playerCanTalkMB_57A160(int a1);
int nox_xxx_giant_57A190(int a1);
char* sub_57A1B0(__int16 a1);
char sub_57A1E0(int* a1, const char* a2, int* a3, char a4, __int16 a5);
int sub_57A3F0(char* a1, int a2, int a3, int a4);
void sub_57A4D0(wchar_t* a1, int a2, int a3, int a4);
BOOL sub_57A620(unsigned __int8 a1, const wchar_t** a2, int a3, int a4);
int sub_57A950(char* a1);
int sub_57A9F0(const char* a1, const char* a2);
char sub_57AAA0(const char* a1, char* a2, int* a3);
int* sub_57ADF0(int* a1);
int sub_57AE30(const char* a1);
int nox_xxx_playerCheckSpellClass_57AEA0(int a1, int a2);
BOOL sub_57AEE0(int a1, int a2);
int nox_xxx_get_57AF20();
void sub_57B0A0();
__int64 nox_xxx___Getcvt_57B180();
int sub_57B190(unsigned __int16 a1, unsigned __int16 a2);
void nox_xxx_loadBaseValues_57B200();
float* sub_57B350();
float* nox_xxx_plrGetMaxVarsPtr_57B360(int a1);
char sub_57B370(int a1, unsigned __int8 a2, int a3);
BOOL nox_xxx_playerClassCanUseItem_57B3D0(nox_object_t* item, char a2);
int nox_xxx_client_57B400(int a1);
BOOL sub_57B450(nox_drawable* a1p);
int sub_57B4D0(int a1);
char sub_57B500(int a1, int a2, char a3);
char sub_57B630(int a1, int a2, int a3);
float2* sub_57B770(float2* a1, float2* a2);
int nox_xxx_collideReflect_57B810(float* a1, int a2);
BOOL nox_xxx_map_57B850(float2* a1, float* a2, float2* a3);
int sub_57B920(void* a1);
char sub_57B930(int a1, int a2, int a3, unsigned int a4);
char nox_xxx_cliGenerateAlias_57B9A0(int a1, int a2, int a3, unsigned int a4);
int sub_57BA10(int a1, __int16 a2, __int16 a3, int a4);
int sub_57BA30(int2* a1, int2* a2, int4* a3);
int nox_xxx_getToken_57BBC0(FILE* f, char* buf, int bufSz);
int nox_xxx_mapNxzDecompress_57BC50(char* a1, char* a2);
int nox_xxx_mapFile_57BDD0(LPVOID lpMem, int a2);
BOOL sub_57BF20(int a1, int a2);
BOOL sub_57BF40(int a1, int a2);
int sub_57BF60();
int sub_57BF70();
unsigned int sub_57BF80();
int nox_xxx_allocGroupRelatedArrays_57BFB0();
void sub_57C000();
int sub_57C030();
int nox_server_getFirstMapGroup_57C080();
int nox_server_getNextMapGroup_57C090(int a1);
int nox_server_scriptGetGroup_57C0A0(int a1);
int nox_server_mapLoadAddGroup_57C0C0(char* a1, int a2, char a3);
int sub_57C130(_DWORD* a1, int a2);
int sub_57C1E0(_DWORD* a1, int a2);
void* nox_server_scriptGetMapGroupByName_57C280(const char* a1, int a2);
void* sub_57C330();
void* sub_57C360();
void sub_57C370(_QWORD* a1);
void sub_57C390(_QWORD* a1);
int nox_server_addNewMapGroup_57C3B0(int a1);
void* nox_xxx_getDebugData_57C3E0();
void* nox_xxx_nextDebugObject_57C3F0(void* a1);
BOOL nox_xxx_allocDebugDataArray_57C410();
void sub_57C440();
int sub_57C460();
int sub_57C490(const char* a1);
int sub_57C500(const char* a1, const char* a2);
void sub_57C5A0(const char* a1);
void sub_57C650(float* a1, float* a2, float* a3);
void sub_57C790(float4* a1, float2* a2, float2* a3, float a4);
BOOL nox_xxx_mathPointOnTheLine_57C8A0(float4* a1, float2* a2, float2* a3);
char nox_xxx_mapTraceRayImpl_57C9A0(int a1, int a2, float* a3, float* a4, char a5);
float2* sub_57CD30(float4* a1, float2* a2, float2* a3);
float2* sub_57CD70(float4* a1, float2* a2, float2* a3);
int sub_57CDB0(int2* a1, float* a2, float2* a3);
void __thiscall sub_57D150(LPVOID* this);
int __thiscall sub_57D1C0(void** this, int a2, unsigned __int8* a3, int a4);
void __thiscall sub_57DD70(LPVOID* this);
_DWORD* __thiscall sub_57DD90(_DWORD* this);
void __thiscall sub_57DDC0(LPVOID* this);
int __thiscall sub_57DDD0(void** this);
unsigned int sub_57DDE0(int a1, int a2);
int __thiscall sub_57DEA0(_DWORD* this, _WORD* a2);
_DWORD* __thiscall sub_57DF00(_DWORD* this);
void __thiscall sub_57DF70(LPVOID* this);
int __thiscall sub_57DFC0(_DWORD* this, int* a2);
int __thiscall sub_57E2C0(_DWORD** this);
unsigned int __thiscall sub_57E3F0(_DWORD** this, unsigned int a2, int a3);
int __thiscall sub_57E4C0(_DWORD** this, unsigned int a2, int a3, unsigned int a4, unsigned int a5);
int __thiscall sub_57E7D0(_DWORD** this);
_DWORD* __thiscall sub_57E8A0(_DWORD* this);
void __thiscall sub_57E910(LPVOID* this);
int __thiscall sub_57E970(void** this);
_DWORD* __thiscall sub_57E9A0(_DWORD* this);
void __thiscall sub_57EA00(LPVOID* this);
int __thiscall sub_57EA60(int this);
int __thiscall nox_xxx_nxzDecompressImpl_57EA80(_DWORD* this, _BYTE* a2, _DWORD* a3, unsigned int a4, _DWORD* a5);
_DWORD* __thiscall sub_57F160(int this, int a2, int a3);
char sub_57F1D0(float2* a1);
int sub_57F2A0(float2* a1, int a2, int a3);
int SEH_40E260(int a1, int a2, int a3, int a4);
int SEH_40E320(int a1, int a2, int a3, int a4);
int SEH_556500(int a1, int a2, int a3, int a4);
int SEH_556570(int a1, int a2, int a3, int a4);
int SEH_5565E0(int a1, int a2, int a3, int a4);
int SEH_557C70(int a1, int a2, int a3, int a4);
int SEH_5704C0(int a1, int a2, int a3, int a4);
int SEH_570570(int a1, int a2, int a3, int a4);
void sub_57FCD0();
int SEH_570B90(int a1, int a2, int a3, int a4);
int SEH_571810(int a1, int a2, int a3, int a4);
int SEH_571C0E(int a1, int a2, int a3, int a4);
int SEH_572045(int a1, int a2, int a3, int a4);
int SEH_57211C(int a1, int a2, int a3, int a4);
int SEH_57222A(int a1, int a2, int a3, int a4);
int SEH_572333(int a1, int a2, int a3, int a4);
int SEH_572442(int a1, int a2, int a3, int a4);
int SEH_572A06(int a1, int a2, int a3, int a4);
int SEH_572BD6(int a1, int a2, int a3, int a4);
int SEH_573153(int a1, int a2, int a3, int a4);
int SEH_57330C(int a1, int a2, int a3, int a4);
int SEH_573401(int a1, int a2, int a3, int a4);
int SEH_57366C(int a1, int a2, int a3, int a4);
int SEH_57381A(int a1, int a2, int a3, int a4);
int SEH_57390B(int a1, int a2, int a3, int a4);
int SEH_5739D2(int a1, int a2, int a3, int a4);
int SEH_574410(int a1, int a2, int a3, int a4);
int SEH_5755F0(int a1, int a2, int a3, int a4);
int SEH_575E60(int a1, int a2, int a3, int a4);
int SEH_576A20(int a1, int a2, int a3, int a4);
int SEH_576DA0(int a1, int a2, int a3, int a4);
int SEH_577060(int a1, int a2, int a3, int a4);
int SEH_577100(int a1, int a2, int a3, int a4);
int SEH_577760(int a1, int a2, int a3, int a4);
int SEH_577F90(int a1, int a2, int a3, int a4);
int SEH_578240(int a1, int a2, int a3, int a4);
int SEH_578370(int a1, int a2, int a3, int a4);
void sub_5800B0();
int SEH_578810(int a1, int a2, int a3, int a4);
int SEH_57D1C0(int a1, int a2, int a3, int a4);
int SEH_57DF00(int a1, int a2, int a3, int a4);
int SEH_57DF70(int a1, int a2, int a3, int a4);
int SEH_57E8A0(int a1, int a2, int a3, int a4);
int SEH_57E910(int a1, int a2, int a3, int a4);
int SEH_57EA00(int a1, int a2, int a3, int a4);

void nox_client_setOnLobbyServer_555000(int (*fnc)(const char*, uint16_t, const char*, const char*));
int __thiscall sub_558800(int (**this)(void));
int __thiscall sub_558810(int (**this)(void));

#endif // NOX_PORT_GAME5_2
