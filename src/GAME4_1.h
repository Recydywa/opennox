#ifndef NOX_PORT_GAME4_1
#define NOX_PORT_GAME4_1

#include "defs.h"

int sub_5098A0();
int sub_5099B0();
void nox_server_checkVictory_509A60();
void sub_509C30(nox_playerInfo* pl);
int* sub_509CB0();
int sub_509CF0(const char* a1, char a2, int a3);
int sub_509D80(int a1);
int nox_xxx_xferDirectionToAngle_509E00(uint32_t* a1);
int nox_xxx_xferIndexedDirection_509E20(int a1, int2* a2);
int nox_xxx_mathDirection4ToAngle_509E90(int a1);
int nox_xxx_math_509EA0(int a1);
int nox_xxx_math_509ED0(float2* a1);
void nox_xxx_utilNormalizeVector_509F20(float2* a1);
int sub_509FF0(int a1);
int nox_xxx_monsterActionIsCondition_50A010(int a1);
int nox_xxx_mobActionGet_50A020(int a1);
int sub_50A040(int a1);
int nox_xxx_monsterIsActionScheduled_50A090(int a1, int a2);
int nox_xxx_checkMobAction_50A0D0(nox_object_t* a1, int a2);
int nox_xxx_monsterPopAction_50A160(nox_object_t* a1p);
#define nox_xxx_monsterPushAction_50A260(obj, a2) nox_xxx_monsterPushAction_50A260_impl(obj, a2, __FILE__, __LINE__)
void* nox_xxx_monsterPushAction_50A260_impl(nox_object_t* a1p, int a2, char* file, int line);
int* nox_xxx_monsterAction_50A360(int a1, int a2);
void nox_xxx_monsterClearActionStack_50A3A0(int a1);
int nox_xxx_monsterCallDieFn_50A3D0(uint32_t* a1);
void nox_xxx_unitUpdateMonster_50A5C0(nox_object_t* a1);
char nox_xxx_updateNPCAnimData_50A850(nox_object_t* a1);
int nox_xxx_mobAction_50A910(nox_object_t* a1);
int nox_xxx_minimapFirstMonster_50AAE0();
int nox_xxx_minimapNextMonster_50AB10();
short sub_50AB50(int a1, int a2);
int nox_xxx_allocVisitNodesArray_50AB90();
void sub_50ABF0();
int sub_50AC20(int a3, uint16_t* a2);
void sub_50AE80(int a1, int a2);
int sub_50AEA0(int a1, float2* a2, uint32_t* a3);
int nox_xxx_aiTestUnitDangerous_50B2C0(int a1);
void sub_50B500();
int sub_50B510();
int sub_50B520();
int nox_xxx_mapTraceObstacles_50B580(nox_object_t* a1, float4* a3);
int sub_50B810(int a1, float* a2);
int sub_50B870(int a1, int a2, int a3);
int nox_xxx_pathfind_preCheckWalls2_50B8A0(int a1, int a2, int a3);
unsigned int sub_50B8E0(int a1, int a2, int a3);
int sub_50B950(int a1, int a2, int a3);
int nox_xxx_genPathToPoint_50B9A0(int a1, int a2, int a3, float* a4);
int sub_50C320(int a1, int a2, uint32_t* a3);
int sub_50C830(int a1, int a2, int a3);
void nox_xxx_pathfind_preCheckWalls_50C8D0(int a1, int2* a2);
int nox_xxx_appendWorkPath_50C990(int a1, int a2, int a3);
int nox_xxx_generateRetreatPath_50CA00(int a1, int a2, int a3, float* a4);
int sub_50CA60(int a4, int a2, int a3);
void sub_50CAC0(int a1, int a2);
int nox_xxx_pathFindStatus_50CAF0();
int sub_50CB00();
void* sub_50CB10();
int sub_50CB20(int a1, float* a2);
void nox_xxx_audioAddAIInteresting_50CD40(int a1, nox_object_t* a2, float2* a3);
void sub_50D1C0();
int nox_xxx_creatureSetDetailedPath_50D220(int a1, int a2);
int sub_50D2A0(int a1, int a2);
int sub_50D2E0(int a1);
int nox_xxx_creatureActuallyMove_50D3B0(float* a1);
int nox_xxx_creatureSetMovePath_50D5A0(int a1);
int nox_xxx_allocMonsterRelatedArrays_50D780();
void sub_50D7E0();
void sub_50D820();
void sub_50D8D0();
int sub_50D960();
int sub_50DE10();
uint32_t* sub_50DE60(int a1);
int sub_50DE80(int a1, float* a2);
void sub_50DFB0(float* a1, int a2);
int sub_50E030(int a1, uint32_t* a2);
int sub_50E110(int a1);
void sub_50E140(int a1);
int sub_50E1B0(int a1);
void sub_50E1E0(int object);
void sub_50E210(nox_object_t* a1);
int nox_xxx_registerShopClasses_50E2A0();
int nox_xxx_deleteShopInventories_50E300();
int sub_50E360();
int nox_xxx_shopGetItemCost_50E3D0(int a1, int a2, float a3);
int sub_50E820(int a1, int a2);
uint32_t* nox_xxx_createShopStruct_50E870();
void nox_xxx_loadShopItems_50E970(int a1);
float* nox_xxx_addItemToShopSession_50EE00(int a1, float a2);
int sub_50EEC0(uint32_t* a1);
int sub_50F2B0(int a1, uint32_t* a2);
void sub_50F3A0(uint32_t* a1);
int sub_50F450(int a1);
int sub_50F490(uint32_t* a1, int a2);
void nox_xxx_shopExit_50F4C0(uint32_t* a1);
int nox_xxx_sendEndTradeMsg_50F560(int a1);
void nox_xxx_tradeAccept_50F5A0(int a1, int a2);
int sub_50F6B0(int a1);
int sub_50F6E0(int a1);
int sub_50F720(int a1, uint32_t* a2);
void sub_50F790(int a1, int a2, int* a3);
uint32_t* sub_50F7F0(int a1, int a2);
int nox_xxx_tradeP2PUpdStuff_50FA00(int a1, uint32_t* a2);
uint32_t* sub_50FB90(uint32_t* a1);
int sub_50FD20(uint32_t* a1, int a2);
int nox_xxx_tradeP2PAddOfferMB_50FE20(int a1, int a2);
int sub_50FF90(int a1, int a2, int a3);
uint32_t* sub_50FFE0(uint32_t* a1, int a2);
void sub_510000(int a1);
int nox_xxx_getSomeShopData_5103A0(int a1, int a2);
int sub_5103F0(int a1, int a2);
uint32_t* sub_5108D0(int a1, int a2, int a3);
uint32_t* sub_510AE0(int* a1, int a2, uint32_t* a3);
unsigned short sub_510D10(int* a1, int a2, int a3, unsigned int a4);
void nox_xxx_shopCancelSession_510DC0(uint32_t* a1);
int sub_510DE0(int a1, int a2);
void sub_510E20(int a1);
void sub_510E50();
signed int nox_xxx_updateSentryGlobe_510E60(int a1);
uint32_t* nox_xxx_sentryUpdateList_510FD0(uint32_t* a1);
void nox_xxx_sentry_511020(int a1, int a2);
void sub_511100(int a1);
int sub_511250(int a1, float* a2);
void nox_xxx_updateSprings_5113A0();
int nox_xxx_unitSetDecayTime_511660(nox_object_t* a1, int a2);
int nox_xxx_decay_5116F0(nox_object_t* item);
void nox_xxx_decay_511750();
int nox_xxx_decayDestroy_5117B0();
char sub_5117F0(nox_object_t* a1);
void nox_xxx_unit_511810(nox_object_t* a1);
void nox_xxx_collisions_511850();
int nox_xxx_updateObjectsVelocity_5118A0(float step);
nox_object_t* nox_xxx_script_511C50(int a1);
int sub_511CB0(uint32_t* a1, int a2);
int sub_511CE0(uint32_t* a1, int a2);
int sub_511D20();
int nox_xxx_scriptPrepareFoundUnit_511D70(nox_object_t* obj);
int sub_511DC0();
int sub_511DE0(nox_object_t* a1);
int sub_511E20();
void nox_xxx_wallOpen_511F80(int a1);
void nox_xxx_wallClose_512070(int a1);
void nox_xxx_wallToggle_512160(int a1);
int nox_xxx_wallPreDestroyByPtr_5122C0(int a1);
float* nox_xxx_monsterLookAt_5125A0(nox_object_t* obj, int a2);
uint32_t* sub_5130E0(int a1, uint32_t* a2);
int sub_513280(int a1, int* a2);
short nox_xxx_enchantUnit_513390(int a1, int* a2);
void nox_xxx_monsterWalkTo_514110(nox_object_t* obj, float x, float y);
void nox_xxx_monsterGoPatrol_515680(int a1, int a2);
void nox_xxx_unitHunt_5157A0(nox_object_t* obj);
void nox_xxx_unitIdle_515820(nox_object_t* obj);
void nox_xxx_unitSetFollow_5158C0(nox_object_t* obj1, nox_object_t* obj2);
int nox_xxx_monsterSetAgressiveness_515980(int a1, uint32_t* a2);
void nox_xxx_monsterActionMelee_515A30(int a1, float2* a2);
void nox_xxx_monsterMissileAttack_515B80(int a1, uint32_t* a2);
int sub_515C80(int a1, uint8_t* a2);
void nox_xxx_mobSetFightTarg_515D30(int a1, int a2);
int sub_515E20(int a1, uint32_t* a2);
int sub_515EB0(int a1, uint32_t* a2);
void nox_server_scriptFleeFrom_515F70(int a1, uint32_t* a2);
void sub_516090(int a1, uint32_t* a2);
void sub_516A80(int a1, int a2, int a3, int a4);
unsigned int sub_516D00(nox_object_t* a1);
int nox_xxx_allocPendingOwnsArray_516EE0();
int sub_516F10();
void sub_516F30();
uint32_t* sub_516F90(int a1, int a2);
void sub_516FC0();
int nox_xxx_readStr_517090(FILE* a1, uint8_t* a2);
int sub_517140(FILE* a1);
int nox_xxx_servParseMonsterDef_517170(FILE* a1, const char* a2);
uint32_t* nox_xxx_monsterListFree_5174F0();
int nox_xxx_monsterList_517520();
uint32_t* nox_xxx_monsterDefByTT_517560(int a1);
int sub_517590(float a1, float a2);
int nox_xxx_roundCoord_5175E0(float a1, int a2);
short nox_xxx_unitCreateMissileSmth_517640(nox_object_t* a1);
int nox_xxx_addObjToMapMB_517780(int a1, int a2, int a3, int a4);
short sub_517870(nox_object_t* a1);
unsigned short* sub_5178E0(int a1, unsigned short* a2);
void nox_xxx_moveUpdateSpecial_517970(nox_object_t* unit);
void nox_xxx_waypointMapRegister_5179B0(int a1);
int sub_517A70(int a1);
int sub_517AE0();
void sub_517B30();
void nox_xxx_unitsGetInCircle_517F90(float2* a1, float a2, void* a3, nox_object_t* a4);
void nox_xxx_unitsGetInNotFarFn_518000(float* a1, int a2);
int sub_518040(int arg0, float a2, int arg8, int a4);
void sub_5180B0(int a1, int a2);
void sub_518170(void* a1p, float a2, void* a3, nox_object_t* a4p);
int sub_518460(float2* a1, unsigned char a2, int a3);
int sub_518550(int* a1, int* a2, unsigned char a3, int a4);
int sub_518740(float2* a1, unsigned char a2);
int sub_518770();
int nox_xxx_netSendPhantomPlrMb_5187E0(int a1, int a2);
int nox_xxx_netSendSimpleObj_5188A0(int a1, int a2);
int nox_xxx_netSendComplexObject_518960(int a1, uint32_t* a2, int a3);
int nox_xxx_netSpriteUpdate_518AE0(int a1, int a2, uint32_t* a3);
int nox_xxx_netPlayerObjSend_518C30(nox_object_t* a1, nox_object_t* a2, int a3, signed int a4);
int nox_xxx_netUpdate_518EE0(nox_object_t* obj);
int nox_xxx_netPlayerObjSendCamera_519330(int a1);
char nox_xxx_unitAroundPlayerFn_5193B0(uint32_t* a1, int a2);
char nox_xxx_netSendObjects2Plr_519410(int a1, int a2);
int sub_519660(int a1, int a2);
int sub_519710(int a1);
void sub_519760(int a1, float* a2);
int nox_xxx_netMapSendClear_519830(int a1, char a2);
int nox_xxx_netMapSendStop_519870();
void nox_xxx_mapSendCancelAll_5198B0(unsigned char a1);
int sub_519930(int a1);
void* sub_51A100();
int sub_51A130();
int sub_51A500(int a1);
char* sub_51A550();
void nox_xxx_spawnHecubahQuest_51A5A0(int* a1);
void nox_xxx_spawnNecroQuest_51A7A0(int* a1);
int nox_xxx_getQuestStage_51A930();
int sub_51A940(int a1);
int sub_51A950();
void nox_xxx_playerResetControlBuffer_51AC30(int a1);
void sub_51B810(nox_object_t* a1);
char sub_51B860(int a1);
void nox_xxx_updateFallLogic_51B870(nox_object_t* a1);
void sub_51D0E0();
int sub_51D0F0(char a1);
int sub_51D100(int a1);
uint32_t* sub_51D120(float* a1);
float* sub_51D1A0(float2* a1);
int sub_51D2C0(int a1, int a2);
int sub_51D300(int a1, int a2, char a3);
float2* sub_51D3F0(float2* a1, float2* a2);
int nox_xxx_tileGetDefByName_51D4D0(char* a1);
int nox_xxx_tileCheckImage_51D540(int a1);
int nox_xxx_tileCheckImageVari_51D570(int a1);
int nox_xxx_tile_51D5C0(int a1);
float* sub_51D5E0(float* a1);
int sub_51D8F0(float2* a1);
int sub_51D9C0(int a1, int a2, int a3, int a4, int a5);

void* nox_xxx_pathFind_50BA00(int a1, int a2, float* a3, float* a4, int (*a5)(int, int, int), int a6);
void sub_517B70(float2* a1, void (*a2)(uint32_t, int), int a3);
void nox_xxx_getUnitsInRect_517C10(float4* a1, void (*a2)(float*, int), void* a3);
void nox_xxx_getUnitsInRectAdvImpl_517DC0(float4* a1, void (*a2)(int, int), int a3);
void nox_xxx_getUnitsInRectAdv_517ED0(float4* a1, void (*a2)(float*, int), int a3);
int* nox_xxx_secretWallCheckUnits_517F00(float* a1, int (*a2)(int*, int), int a3);

#endif // NOX_PORT_GAME4_1
