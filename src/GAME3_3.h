#ifndef NOX_PORT_GAME3_3
#define NOX_PORT_GAME3_3

#include "defs.h"

int nox_server_handler_PlayerDamage_4E17B0(int a1, int a2, int a3, int a4, int a5);
void nox_xxx_playerDecrementHPMana_4E20F0(int a1, int a2, float a3);
void nox_xxx_playerDamageItems_4E2180(int a1, int a2, int a3, int a4, float a5);
double sub_4E2220(int a1);
int sub_4E22A0(int a1, int a2, int a3, int a4, float a5, int a6);
int sub_4E2330(int a1, int a2, int a3, int a4, float a5, int a6);
int sub_4E23C0(int a1, int a2, int a3, int a4, int a5);
int sub_4E24B0(int a1, int a2, int a3, int a4, int a5);
int sub_4E24E0(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_damageFlammable_4E2520(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_damageBlackPowder_4E2560(int a1, int a2, int a3, int a4, int a5);
void nox_xxx_gameSetWallsDamage_4E25A0(int a1);
void nox_xxx_mapDamageUnitsAround_4E25B0(float* a1, float a2, float a3, int a4, int a5, nox_object_t* a6, nox_object_t* a7);
int nox_xxx_damageMonsterGen_4E27D0(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_allocClassArrayObjects_4E3360(unsigned int a1);
int nox_xxx_freeGameObjectClass_4E3420();
nox_object_t* nox_xxx_newObjectWithTypeInd_4E3450(int ind);
nox_object_t* nox_xxx_newObjectByTypeID_4E3810(char* id);
int nox_xxx_objectFreeMem_4E38A0(nox_object_t* a1);
char* nox_xxx_getUnitName_4E39D0(nox_object_t* a1);
char* nox_xxx_getUnitNameByThingType_4E3A80(int a1);
int nox_xxx_getNameId_4E3AA0(char* a1);
int nox_xxx_unitDefGetCount_4E3AC0();
int sub_4E3AD0(int a1);
nox_objectType_t* nox_xxx_objectTypeByInd_4E3B70(int ind);
int sub_4E3B80(int a1);
int nox_xxx_getUnitDefDd10_4E3BA0(int a1);
void sub_4E3BF0(nox_objectType_t* a1);
double sub_4E3CA0();
int sub_4E3CB0(float a1);
int nox_game_getQuestStage_4E3CC0();
void nox_game_setQuestStage_4E3CD0(int a1);
int nox_xxx_player_4E3CE0();
int sub_4E3D50();
short sub_4E3DD0();
void sub_4E4080(float a1);
double sub_4E40B0();
void sub_4E40C0(float a1);
double sub_4E40F0();
int sub_4E4100();
int sub_4E41B0(char* a1);
int sub_4E42C0(FILE* a1);
int sub_4E4390(FILE* a1);
FILE* sub_4E43F0(char* a1);
void nox_xxx_unitNeedSync_4E44F0(nox_object_t* a1);
int* sub_4E4500(int a1, int a2, int a3, int a4);
int nox_xxx_unitSetHP_4E4560(nox_object_t* obj, unsigned short amount);
int* nox_xxx_unitSetOnOff_4E4670(int a1, int a2);
void nox_xxx_unitRaise_4E46F0(nox_object_t* obj, float a2);
int* nox_xxx_unitUnsetXStatus_4E4780(int a1, int a2);
int* nox_xxx_unitSetXStatus_4E4800(int a1, int* a2);
int* nox_xxx_servMarkObjAnimFrame_4E4880(int a1, int a2);
int* nox_xxx_setUnitBuffFlags_4E48F0(int a1, int a2);
int* nox_xxx_modifSetItemAttrs_4E4990(nox_object_t* a1, int* a2);
double nox_xxx_objectGetMass_4E4A70(int a1);
int* nox_xxx_setNPCColor_4E4A90(nox_object_t* a1p, unsigned char a2, int a3);
int* nox_xxx_npcSetItemEquipFlags_4E4B20(int a1, nox_object_t* item, int a3);
int sub_4E4C00(nox_object_t* item);
int sub_4E4C10(nox_object_t* item);
int nox_object_getInitData_4E4C30(nox_object_t* item);
int sub_4E4C50(nox_object_t* item);
int sub_4E4C80(nox_object_t* item);
int sub_4E4C90(int a1, unsigned int a2);
int sub_4E4DE0();
int sub_4E4E50(int a1);
int sub_4E4ED0();
int sub_4E4EF0();
int sub_4E4F30(int a1);
int nox_xxx_playerResetImportantCtr_4E4F40(int a1);
int sub_4E4F80();
void sub_4E4FC0(int a1);
int nox_xxx_netSendPacket_4E5030(int a1, const void* a2, signed int a3, int a4, int a5, char a6);
int nox_xxx_importantCheckRate_4E52B0();
char* nox_xxx_playerKickDueToRate_4E5360(int a1);
int nox_xxx_netSendPacket1_4E5390(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_netClientSend2_4E53C0(int a1, const void* a2, int a3, int a4, int a5);
int nox_xxx_netSendPacket0_4E5420(int a1, const void* a2, signed int a3, int a4, int a5);
int sub_4E5450(int a1, char* a2, signed int a3, int a4, int a5);
void sub_4E54D0(int a1, int a2, int a3);
int sub_4E55A0(unsigned char a1, int a2);
int sub_4E55F0(unsigned char a1);
unsigned int nox_xxx_importantCheckRate2_4E5670(unsigned char a1);
void nox_xxx_netImportant_4E5770(unsigned char a1, int a2);
void nox_xxx_playerRemoveSpawnedStuff_4E5AD0(nox_object_t* a1);
int nox_xxx_isUnit_4E5B50(nox_object_t* a1p);
int sub_4E5B80(nox_object_t* a1p);
void sub_4E5BF0(int a1);
void nox_xxx_delayedDeleteObject_4E5CC0(nox_object_t* obj);
void nox_xxx_finalizeDeletingUnits_4E5EC0();
int sub_4E5F40(int a1);
void sub_4E5FC0(int a1);
void nox_xxx_playerCameraUnlock_4E6040(nox_object_t* player);
void nox_xxx_playerCameraFollow_4E6060(nox_object_t* player, nox_object_t* unitId);
nox_object_t* sub_4E6150(nox_playerInfo* a1);
int sub_4E6230();
nox_object_t* nox_xxx_playerObserverFindGoodSlave0_4E6280(nox_playerInfo* a1);
void nox_xxx_updatePlayerObserver_4E62F0(nox_object_t* a1);
int nox_xxx_playerGoObserver_4E6860(nox_playerInfo* pl, int a2, int a3);
void nox_xxx_playerLeaveObserver_0_4E6AA0(nox_playerInfo* pl);
int sub_4E6BD0(int a1);
double nox_xxx_calcDistance_4E6C00(nox_object_t* a1, nox_object_t* a2);
int sub_4E6CE0(float2* a1, float2* a2);
int nox_server_testTwoPointsAndDirection_4E6E50(float2* a1, int a2, float2* a3);
int sub_4E6EA0(int a1, float a2, int a3);
void sub_4E6EF0(int a1, uint32_t* a2);
void nox_xxx_unitMove_4E7010(nox_object_t* obj, float2* a2);
void nox_xxx_teleportToMB_4E7190(uint8_t* a1, float* a2);
void sub_4E71F0(int a1);
int nox_xxx_objectUnkUpdateCoords_4E7290(nox_object_t* a1);
int sub_4E7350(int a1);
void nox_xxx_spawnSomeBarrel_4E7470(int a1, int a2);
void sub_4E7540(nox_object_t* a1, nox_object_t* a2);
char nox_xxx_objectSetOn_4E75B0(nox_object_t* obj);
int nox_xxx_objectSetOff_4E7600(nox_object_t* obj);
int sub_4E7700(int a1);
int nox_xxx_inventoryGetFirst_4E7980(int a1);
int nox_xxx_inventoryGetNext_4E7990(int a1);
int sub_4E79B0(int a1);
char nox_xxx_unitFreeze_4E79C0(nox_object_t* obj, int a2);
char nox_xxx_unitUnFreeze_4E7A60(nox_object_t* obj, int a2);
void nox_xxx_unitBecomePet_4E7B00(int a1, int a2);
void nox_xxx_monsterRemoveMonitors_4E7B60(nox_object_t* a1, nox_object_t* a2);
int sub_4E7BC0(int a1);
int nox_xxx_unitIsCrown_4E7BE0(int a1);
int nox_xxx_unitIsGameball_4E7C30(int a1);
int nox_xxx_unitIsUnitTT_4E7C80(nox_object_t* a1, int a2);
int nox_xxx_unitCountSlaves_4E7CF0(int a1, int a2, int a3);
int nox_xxx_inventoryCountObjects_4E7D30(int a1, int a2);
int sub_4E7DE0(int a1, nox_object_t* item);
int sub_4E7EC0(int a1, nox_object_t* item);
char* nox_xxx_unitPostCreateNotify_4E7F10(nox_object_t* a1);
int nox_xxx_unitIsHostileMimic_4E7F90(int a1, int a2);
char* nox_xxx_monsterMarkUpdate_4E8020(int a1);
int sub_4E80C0(char a1);
char* sub_4E8110(int a1);
int sub_4E81D0(nox_object_t* a1p);
int sub_4E8290(char a1, short a2);
int sub_4E82C0(unsigned char a1, char a2, char a3, short a4);
char* sub_4E8310();
unsigned char* sub_4E8320(unsigned char a1);
void nox_xxx_fnFindCloseDoors_4E8340(float* a1, int a2);
int sub_4E8390(int a1);
unsigned char* nox_xxx_collideMonsterEventProc_4E83B0(int a1, int a2);
unsigned char* nox_xxx_collideMimic_4E83D0(int a1, int a2);
void nox_xxx_collidePlayer_4E8460(int a1, int a2);
void sub_4E86E0(int a1, float* a2);
void nox_xxx_collideProjectileGeneric_4E87B0(int a1, int a2);
void nox_xxx_collideProjectileSpark_4E8880(int a1, int a2);
int nox_xxx_doorGetSomeKey_4E8910(int a1, int a2);
int sub_4E8A10();
void nox_xxx_collideDoor_4E8AC0(int a2, int a3);
int nox_xxx_collidePickup_4E8DF0(int a1, int a2);
unsigned char* sub_4E8E50();
int sub_4E8E60();
bool nox_server_questMaybeWarp_4E8F60();
int sub_4E9010();
void nox_xxx_collideExit_4E9090(int a1, int a2, int a3);
void nox_xxx_collideDamage_4E9430(int a1, int a2);
uint32_t* nox_xxx_collideManadrain_4E9490(int a1, int a2);
void nox_xxx_spellFlyCollide_4E9500(int a1, int a2, float* a3);
void nox_xxx_collideBomb_4E96F0(int a1, int a2);
void nox_xxx_collideBoom_4E9770(int a1, int a2, float* a3);
void nox_xxx_collideDie_4E99B0(int unit, int a2);
char nox_xxx_collideGlyph_4E9A00(int a1, int a2);
int sub_4E9A30(int a1, int a2);
void nox_xxx_fireballCollide_4E9AC0(int a1, int a2);
void nox_xxx_collideChest_4E9C40(uint32_t* a1, int a2);
void nox_xxx_collideSulphurShot2_4E9D80(int a1, int a2, float* a3);
void nox_xxx_collideSulphurShot_4E9E50(int a1, int a2, int a3);
int* nox_xxx_collideDeathBall_4E9E90(int a1, int a2, float* a3);
void nox_xxx_collideDeathBallFragment_4E9FE0(int a1, int a2, float* a3);
void nox_xxx_collidePixie_4EA080(int a1, int a2, float* a3);
void nox_xxx_collideWallReflectSpark_4EA200(int a1, int a2, float2* a3);
void sub_4EA2C0(int a1, int a2);
void nox_xxx_collideSpark_4EA300(int a1, int a2, float* a3);
void nox_xxx_collideWebbing_4EA380(int a1, int a2);
void sub_4EA400(int a1, int a2);
void nox_xxx_pickupFlagCtf_4EA490(int a1, int a2);
int sub_4EA7A0(int a1);
short sub_4EA800(int a1, int a2);
uint32_t* sub_4EAAA0(int a1);
uint32_t* sub_4EAAD0(int a1, int a2);
int nox_xxx_collidePentagram_4EAB20(int a1);
void nox_xxx_collideSign_4EAB40(int a1, int a2);
void nox_xxx_collideTrapDoor_4EAB60(int a1, int a2);
uint32_t* sub_4EACA0(int a1, int a2);
int nox_xxx_collideSpellPedestal_4EAD20(int a1, int a2);
void nox_xxx_collideFist_4EADF0(int a1, int a2);
void nox_xxx_collideTeleportWake_4EAE30(int a1, int a2);
void nox_xxx_collideChakram_4EAF00(int a1, int a2, float* a3);
int sub_4EB250(int a1);
void sub_4EB340(float* a1, int a2);
void sub_4EB3E0(int a1);
void nox_xxx_collideArrow_4EB490(int a1, int a2);
void nox_xxx_collideHarpoon_4EB6A0(nox_object_t* a1, nox_object_t* a2);
void nox_xxx_collideMonsterArrow_4EB800(int a1, int a2);
void nox_xxx_collideBearTrap_4EB890(int* a1, int a2);
void nox_xxx_collidePoisonGasTrap_4EB910(int* a1, int a2);
int sub_4EB9B0(int a1, int a2);
uint32_t* nox_xxx_collideBall_4EBA00(int a1, int a2);
int sub_4EBB50(int a1, int a2);
short nox_xxx_collideHomeBase_4EBB80(int a1, int a2);
void nox_xxx_collideUndeadKiller_4EBD40(int a1, int a2, int a3);
void nox_xxx_collideMonsterGen_4EBE10(int a1, int a2);
void sub_4EBE40(int a1, int a2);
void nox_xxx_collideAnkhQuest_4EBF40(int a1, int a2);
void nox_xxx_unitSetOwner_4EC290(nox_object_t* obj1, nox_object_t* obj2);
void nox_xxx_unitClearOwner_4EC300(nox_object_t* obj);
int nox_xxx_playerObserverFindGoodSlave2_4EC3E0(int a1);
int nox_xxx_playerObserverFindGoodSlave_4EC420(int a1);
void nox_xxx_unitRemoveChild_4EC470(nox_object_t* a1);
void nox_xxx_unitTransferSlaves_4EC4B0(nox_object_t* a1);
int nox_xxx_unitHasThatParent_4EC4F0(int a1, int a2);
int nox_xxx_unitsHaveSameTeam_4EC520(int a1, int a2);
nox_object_t* nox_xxx_findParentChainPlayer_4EC580(nox_object_t* unit);
void sub_4EC5B0();
uint32_t* nox_xxx_respawnAdd_4EC5E0(nox_object_t* a1);
void sub_4EC6A0(int a1);
int nox_xxx_allocItemRespawnArray_4ECA60();
void sub_4ECA90();
int sub_4ECBD0(int a1);
int sub_4ECC00(const char** a1);
nox_object_t* nox_server_getObjectFromNetCode_4ECCB0(int a1);
int nox_server_netCodeCache_lookupObj_4ECD90(int a1);
int sub_4ECDE0(uint32_t* a1, int a2);
int sub_4ECE10(uint32_t* a1, int a2);
int nox_server_netCodeCache_initArray_4ECE50();
int nox_server_netCodeCache_addObj_4ECEA0(int a1);
int nox_server_netCodeCache_nextUnused_4ECEF0();
int sub_4ECF10(int a1);
int sub_4ECFA0(nox_object_t* a1);
int sub_4ECFE0();
int nox_xxx_netGetUnitByExtent_4ED020(int a1);
void sub_4ED050(int a1, int a2);
void sub_4ED0C0(nox_object_t* a1p, nox_object_t* object);
int nox_xxx_dropDefault_4ED290(nox_object_t* a1p, nox_object_t* a2p, float2* a3);
int sub_4ED500(int a1, int a2, float2* a3);
int nox_xxx_dropTrap_4ED580(int a1, int a2, float2* a3);
int nox_xxx_dropCrown_4ED5E0(int a1, int a2, int* a3);
int nox_xxx_dropTreasure_4ED710(int a1, int a2, int* a3);
int nox_xxx_drop_4ED790(nox_object_t* a1, nox_object_t* a2, float2* a3);
int nox_xxx_drop_4ED810(int a1, int a2, float* a3);
int nox_xxx_invForceDropItem_4ED930(int a1, uint32_t* a2);
float2* sub_4ED970(float a1, float2* a2, float2* a3);
uint32_t* nox_xxx_dropAllItems_4EDA40(uint32_t* a1);
int nox_xxx_unitDropCheckSome_4EDCD0(int a1, int a2);
void sub_4EDD00(int a1, int a2);
int sub_4EDDE0(int a1, uint32_t* a2, int* a3);
int nox_xxx_dropFood_4EDE50(int a1, int a2, int* a3);
void nox_xxx_chest_4EDF00(int a1, int a2);
double sub_4EE2A0(int a1);
int nox_xxx_dropAnkhTradable_4EE370(int a1, int a2, int* a3);
void nox_xxx_unitAdjustHP_4EE460(nox_object_t* unit, int dv);
void nox_xxx_mobInformOwnerHP_4EE4C0(nox_object_t* obj);
void nox_xxx_unitDamageClear_4EE5E0(nox_object_t* unit, int damageAmount);
void nox_xxx_unitHPsetOnMax_4EE6F0(int unit);
void nox_xxx_playerHP_4EE730(int a1);
short nox_xxx_unitGetHP_4EE780(nox_object_t* item);
short nox_xxx_unitGetMaxHP_4EE7A0(int a1);
int nox_xxx_unitSetMaxHP_4EE7C0(int a1, short a2);
int nox_xxx_activatePoison_4EE7E0(int a1, int a2, int a3);
void nox_xxx_updatePoison_4EE8F0(int a1, int a2);
void nox_xxx_removePoison_4EE9D0(int a1);
void nox_xxx_setSomePoisonData_4EEA90(int a1, int a2);
unsigned short nox_xxx_playerManaAdd_4EEB80(nox_object_t* unit, short amount);
uint32_t* nox_xxx_playerManaSub_4EEBF0(int unit, int amount);
short nox_xxx_unitGetOldMana_4EEC80(int unit);
short nox_xxx_playerGetMaxMana_4EECB0(int unit);
int nox_xxx_playerSetMaxMana_4EECD0(int unit, short amount);
uint32_t* nox_xxx_playerManaRefresh_4EECF0(int unit);
void nox_xxx_abilGivePlayerAll_4EED40(int a1, char a2, int a3);
int nox_xxx_plrReadVals_4EEDC0(nox_object_t* a1, int a2);
int sub_4EF140(int a1);
double nox_xxx_calcBoltDamage_4EF1E0(int a1, int a2);
void sub_4EF410(int a1, unsigned char a2);
void nox_xxx_set_god_4EF500(int a1);
char nox_xxx_getRespawnWeaponFlags_4EF580();
int sub_4EF6F0(int a1);
nox_object_t* nox_xxx_playerRespawnItem_4EF750(nox_object_t* a1, char* a2, int* a3, int a4, int a5);
char nox_xxx_playerMakeDefItems_4EF7D0(int a1, int a2, int a3);
int nox_xxx_netSendPlayerRespawn_4EFC30(int a1, char a2);
void nox_xxx_spellAwardAll2_4EFC80(nox_playerInfo* a1p);
void nox_xxx_spellAwardAll1_4EFD80(nox_playerInfo* a1p);
void nox_xxx_spellAwardAll3_4EFE10(nox_playerInfo* a1p);
char nox_xxx_unitInitPlayer_4EFE80(nox_object_t* a1);
int sub_4EFF10(int a1);
int nox_xxx_unitMonsterInit_4F0040(int a1);
uint32_t* nox_xxx_unitSparkInit_4F0390(int a1);
int nox_xxx_initFrog_4F03B0(int a1);
int* nox_xxx_initChest_4F0400(int a1);
uint32_t* nox_xxx_unitBoulderInit_4F0420(uint32_t* a1);
int sub_4F0450(int a1);
int sub_4F0490(int a1);
int nox_xxx_unitInitGold_4F04B0(int a1);
int* nox_xxx_breakInit_4F0570(int a1);
int nox_xxx_unitInitGenerator_4F0590(int a1);
uint32_t* nox_server_rewardgen_activateMarker_4F0720(int a1, unsigned int a2);
uint32_t* nox_xxx_rewardSpellBook_4F09F0(int a1, unsigned int a2);
int nox_server_rewardGen_pickRandomSlots_4F0B60(unsigned int a1);
uint32_t* nox_xxx_rewardAbilityBook_4F0C70(int a1);
uint32_t* nox_xxx_rewardFieldGuide_4F0D20(int a1, unsigned int a2);
uint32_t* nox_xxx_rewardMakeArmor_4F0E80(int a1, unsigned int a2);
int nox_xxx_rewardMakeWeapon_4F14E0(int a1, unsigned int a2);
uint32_t* nox_xxx_rewardMakePotion_4F1C40(int a1, unsigned int a2);
uint32_t* nox_xxx_createGem_4F1D30(int a1, unsigned int a2);
uint32_t* nox_xxx_createGem2_4F1F00(int a1, unsigned int a2);
void sub_4F2110();
int sub_4F2210();
int sub_4F24E0(int a1);
int sub_4F2530(int a1);
int sub_4F2570(int a1);
int sub_4F2590(int a1);
int sub_4F2700(int a1);
int sub_4F27A0(int a1);
int sub_4F27E0(int a1);
int sub_4F28C0(int a1);
int sub_4F2960(int a1);
int sub_4F2B20(int a1);
int sub_4F2B60(int a1);
int sub_4F2C30(int a1);
int nox_xxx_spell_4F2E70(int a1);
int sub_4F2EF0(int a1);
int nox_xxx_playerTryEquip_4F2F70(uint32_t* a1, nox_object_t* item);
int nox_xxx_playerTryDequip_4F2FB0(uint32_t* a1, const nox_object_t* object);
int nox_xxx_itemApplyEngageEffect_4F2FF0(nox_object_t* item, int a2);
int nox_xxx_itemApplyDisengageEffect_4F3030(const nox_object_t* object, int a2);
void nox_xxx_inventoryPutImpl_4F3070(int a1, nox_object_t* item, int a3);
bool nox_xxx_playerCheckStrength_4F3180(nox_object_t* a1, nox_object_t* item);
int nox_xxx_pickupDefault_4F31E0(nox_object_t* a1p, nox_object_t* item, int a3);
int nox_xxx_pickupFood_4F3350(int a1, int a2, int a3);
int sub_4F3400(int a1, int a2, int a3);
int nox_xxx_pickupUse_4F34D0(int a1, int a2, int a3);
int nox_xxx_pickupTrap_4F3510(int a1, int a2, int a3);
int nox_xxx_pickupTreasure_4F3580(int a1, int a2, int a3);
int nox_xxx_inventoryServPlace_4F36F0(nox_object_t* a1p, nox_object_t* a2p, int a3, int a4);
int nox_xxx_pickupPotion_4F37D0(int a1, int a2, int a3);
int nox_xxx_pickupAmmo_4F3B00(int a1, nox_object_t* item, int a3, int a4);
int nox_xxx_pickupSpellbook_4F3C60(int a1, int a2, int a3);
int nox_xxx_pickupAbilitybook_4F3CE0(int a1, int a2, int a3);
int sub_4F3DD0(int a1, int a2);
int nox_xxx_xfer_4F3E30(unsigned short a1, nox_object_t* a2, int a3);
int nox_xxx_servMapLoadPlaceObj_4F3F50(nox_object_t* a1, int a2, void* a3);
char sub_4F40A0(int a1);
int nox_xxx_readObjectOldVer_4F4170(int a1, int a2, int a3);
int nox_xxx_mapReadWriteObjData_4F4530(nox_object_t* a1, int a2);
int nox_xxx_XFerDefault_4F49A0(nox_object_t* a1, void* a2);
int nox_xxx_XFerSpellPagePedistal_4F4A20(int a1);
int nox_xxx_XFerReadable_4F4AB0(int a1);
int nox_xxx_XFerExit_4F4B90(int a1);
int nox_xxx_XFerDoor_4F4CB0(int a1);
int nox_xxx_unitTriggerXfer_4F4E50(float a1);
int nox_xxx_XFerHole_4F51D0(int a1);
int nox_xxx_XFerTransporter_4F5300(int a1);
int nox_xxx_XFerElevator_4F53D0(int a1);
int nox_xxx_XFerElevatorShaft_4F54A0(int a1);
int sub_4F5540(int a1);
int nox_xxx_XFerMover_4F5730(int a1);
int nox_xxx_XFerGlyph_4F5890(int a1);
int nox_xxx_XFerInvLight_4F5AA0(int* a1);
int nox_xxx_XFerSentry_4F5E50(int a1);

#endif // NOX_PORT_GAME3_3
