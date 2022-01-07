package nox

/*
#include "GAME1_2.h"
#include "GAME2_1.h"
#include "GAME3_1.h"
#include "GAME3_2.h"
#include "GAME4_1.h"
#include "GAME4_3.h"
#include "GAME5.h"
#include "GAME5_2.h"
#include "client__video__draw_common.h"
#include "client__shell__noxworld.h"
#include "client__system__parsecmd.h"
#include "common__magic__comguide.h"
#include "common__net_list.h"
#include "client__gui__guicon.h"
#include "client__gui__guisave.h"
#include "client__gui__chaticon.h"
#include "client__gui__guirank.h"
#include "client__gui__guisumn.h"
#include "client__gui__guiquit.h"
#include "client__gui__gui_ctf.h"
#include "client__gui__guiobs.h"
#include "client__gui__guitrade.h"
#include "client__gui__guiinput.h"
#include "client__shell__mainmenu.h"
#include "client__shell__selchar.h"
#include "client__shell__selcolor.h"
#include "client__shell__wolapi__wolchat.h"
#include "client__shell__wolapi__wollogin.h"
#include "common__telnet__telnetd.h"
#include "server__network__playback.h"
#include "server__network__sdecode.h"
#include "client__drawable__drawable.h"
extern unsigned int nox_game_createOrJoin_815048;
extern unsigned int nox_client_gui_flag_815132;
extern unsigned int nox_client_gui_flag_1556112;
extern unsigned int nox_client_renderGUI_80828;
extern unsigned int nox_xxx_xxxRenderGUI_587000_80832;
extern unsigned int dword_5d4594_1548524;
extern unsigned int dword_5d4594_2650652;
extern void* dword_5d4594_1548532;
extern unsigned int nox_server_switchToWP_1548664;
extern unsigned int dword_5d4594_1548704;
extern unsigned int dword_5d4594_1556144;
extern unsigned int dword_5d4594_1563064;
extern unsigned int dword_5d4594_251744;
extern unsigned int dword_5d4594_815052;
extern unsigned int nox_console_waitSysOpPass;
extern unsigned int dword_5d4594_1049508;
extern nox_draw_viewport_t nox_draw_viewport;
extern unsigned char nox_net_lists_buf[2048];

int sub_4EDD70();
void sub_426060();
void sub_417160();
void sub_4D3130();
void sub_502100();
void sub_5524C0();
char* sub_4DB160();
void sub_4D2160();
void sub_4D22B0();
void sub_4D2230();
void sub_4DBA30(int a1);
void sub_50AFA0();
void nox_console_sendSysOpPass_4409D0(wchar_t* a1);
unsigned int*  nox_xxx_netUseMap_4DEE00(const char* a1, int a2);
char* nox_xxx_getSomeMapName_4D0CF0();
int  nox_server_loadMapFile_4CF5F0(char* a1, int a2);
void  nox_xxx_addDebugEntry_511590(void* a1, void* a2);
int nox_xxx_mapLoadRequired_4DCC80();
int  sub_4EF660(nox_object_t* a1p);
void  sub_500510(const char* a1);

void nox_xxx_getUnitsInRect_517C10_go(nox_object_t* obj, void* payload);
*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/noxworld-dev/xwis"

	"nox/v1/common"
	"nox/v1/common/alloc"
	"nox/v1/common/console"
	"nox/v1/common/discover"
	"nox/v1/common/env"
	noxflags "nox/v1/common/flags"
	"nox/v1/common/log"
	"nox/v1/common/memmap"
	"nox/v1/common/object"
	"nox/v1/common/types"
	"nox/v1/common/unit/ai"
)

var (
	useXWIS = true
	gameLog = log.New("game")
)

func init() {
	configBoolPtr("network.xwis.register", "NOX_XWIS", true, &useXWIS)
}

var (
	nox_game_playState_811372 int
)

//export nox_xxx_gameGetPlayState_4356B0
func nox_xxx_gameGetPlayState_4356B0() C.int {
	return C.int(gameGetPlayState())
}
func gameGetPlayState() int {
	return nox_game_playState_811372
}

func gameSetPlayState(st int) {
	gameLog.Println("play state:", st)
	nox_game_playState_811372 = st
}

func nox_game_setMovieFile_4CB230(name string, out *C.char) bool {
	cname := C.CString(name)
	defer StrFree(cname)
	return C.nox_game_setMovieFile_4CB230(cname, out) != 0
}

func nox_game_rollLogoAndStart_4AB1F0() bool {
	if isServer || isServerQuest {
		// FIXME: switch to server state directly
		return startServer()
	}
	pathP, freePath := alloc.Malloc(128)
	defer freePath()
	path := (*C.char)(pathP)
	gamePopState()
	if noxflags.HasGame(noxflags.GameFlag26) || !nox_game_setMovieFile_4CB230("WWLogo.vqa", path) {
		nox_game_rollIntroAndStart_4AB170()
		return true
	}
	C.sub_4B0300(path)
	C.sub_4B0640((*[0]byte)(C.nox_game_rollIntroAndStart_4AB170))
	if err := drawGeneral_4B0340(0); err != nil {
		videoLog.Println(err)
	}
	return true
}

//export nox_game_rollIntroAndStart_4AB170
func nox_game_rollIntroAndStart_4AB170() C.int {
	path, freePath := alloc.Malloc(128)
	defer freePath()
	if C.sub_578DF0()&0x80 != 0 || noxflags.HasGame(noxflags.GameFlag26) || !nox_game_setMovieFile_4CB230("Intro.vqa", (*C.char)(path)) {
		nox_game_rollNoxLogoAndStart_4AB0F0()
		return 1
	}
	C.sub_4B0300((*C.char)(path))
	C.sub_4B0640((*[0]byte)(C.nox_game_rollNoxLogoAndStart_4AB0F0))
	if err := drawGeneral_4B0340(1); err != nil {
		videoLog.Println(err)
	}
	C.sub_578DE0(C.char(C.sub_578DF0() | 0x80))
	return 1
}

//export nox_game_rollNoxLogoAndStart_4AB0F0
func nox_game_rollNoxLogoAndStart_4AB0F0() C.int {
	path, freePath := alloc.Malloc(128)
	defer freePath()
	if noxflags.HasGame(noxflags.GameFlag26) || !nox_game_setMovieFile_4CB230("NoxLogo.vqa", (*C.char)(path)) {
		if C.nox_game_showLegal_4CC4E0() == 0 {
			nox_xxx_setContinueMenuOrHost_43DDD0(0)
			C.nox_client_gui_flag_815132 = 0
		}
		return 1
	}
	C.sub_4B0300((*C.char)(path))
	C.sub_4B0640((*[0]byte)(C.nox_game_showLegal_4CC4E0))
	if err := drawGeneral_4B0340(0); err != nil {
		videoLog.Println(err)
	}
	return 1
}

func startServer() bool {
	C.nox_game_createOrJoin_815048 = 1
	setEngineFlag(NOX_ENGINE_FLAG_ADMIN)
	resetEngineFlag(NOX_ENGINE_FLAG_GODMODE)
	noxflags.SetGame(noxflags.GameOnline)
	if !isServerQuest {
		noxflags.SetGame(noxflags.GameNotQuest)
	} else {
		noxflags.UnsetGame(noxflags.GameNotQuest)
	}
	noxflags.UnsetGame(noxflags.GameModeCoop)
	C.sub_461440(0)
	C.sub_4D6F40(0)
	C.sub_4D6F90(0)
	if !isServerQuest {
		nox_xxx_setQuest_4D6F60(0)
	} else {
		nox_xxx_setQuest_4D6F60(1)
		C.nox_xxx_cliShowHideTubes_470AA0(1)
	}
	C.sub_4D6F80(0)
	if C.sub_473670() == 0 {
		C.nox_client_toggleMap_473610()
	}
	if !isServerQuest {
		C.nox_xxx_cliShowHideTubes_470AA0(0)
	}
	C.nox_xxx_cliSetMinimapZoom_472520(2300)
	C.sub_43AF50(0)
	if C.nox_xxx_parseGamedataBinPre_4D1630() == 0 {
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		C.nox_client_gui_flag_815132 = 0
		return false
	}
	nox_xxx_serverHost_43B4D0()
	return true
}

func getServerName() string {
	return C.GoString(C.nox_xxx_serverOptionsGetServername_40A4C0())
}

func getServerMap() string {
	return C.GoString(C.nox_xxx_mapGetMapName_409B40())
}

func getServerCurPlayers() int {
	return cntPlayers()
}

func getServerMaxPlayers() int {
	return int(C.nox_xxx_servGetPlrLimit_409FA0())
}

func nox_xxx_isQuest_4D6F50() bool {
	return memmap.Uint32(0x5D4594, 1556160) != 0
}

func nox_xxx_setQuest_4D6F60(v int) {
	*memmap.PtrUint32(0x5D4594, 1556160) = uint32(v)
}

func sub_4D6F30() uint32 {
	return memmap.Uint32(0x5D4594, 1556156)
}

func sub_43AF30() uint32 {
	return uint32(C.dword_5d4594_815052)
}

//export nox_xxx_set_god_4EF500
func nox_xxx_set_god_4EF500(v C.int) {
	serverCheatGod(v != 0)
}

var srvReg struct {
	cur xwisInfoShort
	srv *discover.RegServer
}

//export nox_xxx_serverHost_43B4D0
func nox_xxx_serverHost_43B4D0() {
	isHost := C.nox_game_createOrJoin_815048 != 0
	gameLog.Println("host:", isHost)
	if isHost {
		// host
		C.nox_client_xxx_switchChatMap_43B510()
		C.nox_client_guiXxx_43A9D0()
		nox_client_guiXxxDestroy_4A24A0()
		C.nox_xxx_gameSetAudioFadeoutMb_501AC0(0)
	} else {
		// join
		C.sub_43B630()
		nox_client_createSockAndJoin_43B440()
	}
	C.sub_49FF20()
	if isHost {
		maybeRegisterGameOnline()
	} else if srvReg.srv != nil {
		srvReg.srv.Close()
		srvReg.srv = nil
	}
}

type xwisInfoShort struct {
	Name       string
	Map        string
	Flags      xwis.GameFlags
	Access     xwis.Access
	Players    int
	MaxPlayers int
	Level      int
}

func getGameInfoXWIS() xwisInfoShort {
	var access xwis.Access
	if acc := sub_416640(); acc[100]&0x20 != 0 {
		access = xwis.AccessPrivate
	} else if acc[100]&0x10 != 0 {
		access = xwis.AccessClosed
	}
	return xwisInfoShort{
		Name:       getServerName(),
		Map:        getServerMap(),
		Flags:      xwis.GameFlags(noxflags.GetGame()),
		Players:    getServerCurPlayers(),
		MaxPlayers: getServerMaxPlayers(),
		Level:      nox_game_getQuestStage_4E3CC0(),
		Access:     access,
	}
}

func xwisIsQuest(info *xwis.GameInfo) bool {
	flags := int(info.Flags) | int(info.MapType)
	return flags&int(xwis.MapTypeQuest) != 0
}

func (v xwisInfoShort) XWIS() xwis.GameInfo {
	info := xwis.GameInfo{
		Name:       v.Name,
		Map:        v.Map,
		Resolution: xwis.Res1024x768,
		Flags:      v.Flags,
		Access:     v.Access,
		Players:    v.Players,
		MaxPlayers: v.MaxPlayers,
	}
	if xwisIsQuest(&info) {
		info.FragLimit = v.Level
	}
	return info
}

func maybeRegisterGameOnline() {
	if !useXWIS || !noxflags.HasGame(noxflags.GameOnline) || env.IsE2E() {
		return
	}
	info := getGameInfoXWIS()
	if srvReg.srv == nil {
		srvReg.cur = info
		srvReg.srv = discover.NewServer(info.XWIS())
		return
	}
	if srvReg.cur == info {
		return
	}
	srvReg.cur = info
	srvReg.srv.Update(info.XWIS())
}

func noxGetViewport() *C.nox_draw_viewport_t {
	return &C.nox_draw_viewport
}

func nox_xxx_getSomeCoods_435670() types.Point {
	vp := noxGetViewport()
	return types.Point{
		X: int(vp.field_6),
		Y: int(vp.field_7),
	}
}

func initGameSession435CC0() error {
	ctx := context.Background()
	C.sub_445450()
	C.sub_45DB90()
	C.sub_41D1A0(0)
	C.nox_xxx_initTime_435570()
	C.nox_client_gui_flag_1556112 = 0

	if C.nox_alloc_drawable_init(5000) == 0 {
		return errors.New("nox_alloc_drawable_init failed")
	}

	sub_49A8E0_init()

	if C.nox_xxx_allocArrayHealthChanges_49A5F0() == 0 {
		return errors.New("nox_xxx_allocArrayHealthChanges_49A5F0 failed")
	}

	if C.nox_xxx_parseThingBinClient_44C840_read_things() == nil {
		return errors.New("nox_xxx_parseThingBinClient_44C840_read_things failed")
	}

	if C.nox_xxx_loadGuides_427070() == 0 {
		return errors.New("nox_xxx_loadGuides_427070 failed")
	}

	if C.sub_494F00() == 0 {
		return errors.New("sub_494F00 failed")
	}

	if !isDedicatedServer {
		if err := nox_game_guiInit_473680(); err != nil {
			return fmt.Errorf("game gui init failed: %w", err)
		}
	}

	C.nox_alloc_npcs()
	if C.nox_xxx_loadReflSheild_499360() == 0 {
		return errors.New("nox_xxx_loadReflSheild_499360 failed")
	}

	C.nox_xxx_initSight_485F80()
	if C.sub_4960B0() == 0 {
		return errors.New("sub_4960B0 failed")
	}

	initDrawableLists()

	if C.nox_xxx_allocArrayDrawableFX_495AB0() == 0 {
		return errors.New("nox_xxx_allocArrayDrawableFX_495AB0 failed")
	}

	if C.nox_xxx_allocClassListFriends_495980() == 0 {
		return errors.New("nox_xxx_allocClassListFriends_495980 failed")
	}

	C.sub_4958F0()
	C.nox_xxx_gameSetCliConnected_43C720(0)
	noxflags.SetGame(noxflags.GameFlag24)
	if noxflags.HasGame(noxflags.GameHost) {
		if !isDedicatedServer {
			C.nox_xxx_netPlayerIncomingServ_4DDF60(31)
		}
	} else {
		nox_xxx_netSendIncomingClient_43CB00()
	}
	gameSetCliDrawFunc(clientDraw)
	gameSetPlayState(3)
	*memmap.PtrUint32(0x587000, 85720) = 1
	sz := videoGetWindowSize()
	vp := noxGetViewport()
	vp.x1 = 0
	vp.y1 = 0
	vp.x2 = C.int(sz.W - 1)
	vp.y2 = C.int(sz.H - 1)
	vp.width = C.int(sz.W)
	vp.height = C.int(sz.H)
	vp.field_10 = 0
	vp.field_11 = 0
	vp.field_12 = 0
	v1 := C.nox_video_getCutSize_4766D0()
	nox_draw_setCutSize_476700(v1, 0)
	if noxflags.HasGame(noxflags.GameModeCoop) {
		C.sub_41CC00((*C.char)(memmap.PtrOff(0x85B3FC, 10984)))
	} else if nox_xxx_isQuest_4D6F50() || sub_4D6F70() {
		if nox_xxx_isQuest_4D6F50() || sub_4D6F70() {
			C.sub_460380()
			C.nox_xxx_cliPrepareGameplay1_460E60()
			C.nox_xxx_cliPrepareGameplay2_4721D0()
		}
		if !noxflags.HasGame(noxflags.GameHost) {
			C.sub_41CC00((*C.char)(memmap.PtrOff(0x85B3FC, 10984)))
		}
	}
	if !isServer {
		C.nox_xxx_plrLoad_41A480((*C.char)(memmap.PtrOff(0x85B3FC, 10984)))
	}
	if isServer && !isDedicatedServer {
		getPlayers()[0].GoObserver(false, true)
	}
	execConsoleCmd(ctx, "execrul autoexec.rul")
	if isServer {
		execServerCmd("set cycle on")
		for _, cmd := range serverExec {
			if len(cmd) == 0 {
				continue
			}
			execServerCmd(cmd)
		}
	}
	C.sub_4951C0()
	C.sub_465DE0(0)
	return nil
}

//export nox_xxx_netServerCmd_440950_empty
func nox_xxx_netServerCmd_440950_empty() {
	nox_xxx_netServerCmd_440950(0, "")
}

//export nox_server_parseCmdText_443C80
func nox_server_parseCmdText_443C80(cstr *C.wchar_t, _ C.int) C.int {
	cmd := GoWString(cstr)
	if cmd == "" {
		return 0
	}
	res := execConsoleCmd(context.Background(), cmd)
	return C.int(bool2int(res))
}

func execConsoleCmd(ctx context.Context, cmd string) bool { // nox_server_parseCmdText_443C80
	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		return false
	}
	if consoleWaitSysOpPass {
		nox_console_sendSysOpPass_4409D0(cmd)
		consoleWaitSysOpPass = false
		return true
	}
	if noxflags.HasGame(noxflags.GameHost) {
		ctx = console.AsServer(ctx)
	} else {
		ctx = console.AsClient(ctx)
	}
	if getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING) {
		ctx = console.AsDedicated(ctx)
	}
	return noxConsole.Exec(ctx, cmd)
}

func execServerCmd(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		return
	}
	ctx := context.Background()
	ctx = console.AsServer(ctx)
	if getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING) {
		ctx = console.AsDedicated(ctx)
	}
	ctx = console.WithCheats(ctx)
	noxConsole.Exec(ctx, cmd)
}

func serverCmdLoadMap(name string) {
	if len(name) == 0 {
		return
	}
	execServerCmd("load " + name)
}

func nox_xxx_getSomeMapName_4D0CF0() string {
	return GoString(C.nox_xxx_getSomeMapName_4D0CF0())
}

func nox_server_currentMapGetFilename_409B30() string {
	return C.GoString((*C.char)(memmap.PtrOff(0x5D4594, 2598188)))
}

func nox_xxx_mapFilenameGetSolo_4DB260() string {
	return C.GoString((*C.char)(memmap.PtrOff(0x5D4594, 1559960)))
}

func nox_xxx_servInitialMapLoad_4D17F0() bool {
	C.sub_4E79B0(0)
	if nox_server_currentMapGetFilename_409B30() == "" {
		nox_xxx_gameSetMapPath_409D70("tutorial.map")
	}
	C.nox_xxx_netMapSendStop_519870()
	if !nox_xxx_mapExitAndCheckNext_4D1860_server() {
		return false
	}
	if debugMainloop {
		log.Println("gameStateFunc = nox_xxx_gameTick_4D2580_server")
	}
	nox_xxx_setGameState_43DDF0(nox_xxx_gameTick_4D2580_server)
	C.nox_netlist_resetAllInList_40EE90(1)
	noxflags.SetGame(noxflags.GameFlag18)
	C.nox_xxx_netGameSettings_4DEF00()
	C.nox_server_gameUnsetMapLoad_40A690()
	return true
}

var (
	gameTickMu    sync.Mutex
	gameTickHooks []func()
)

func addGameTickHook(fnc func()) {
	gameTickMu.Lock()
	gameTickHooks = append(gameTickHooks, fnc)
	gameTickMu.Unlock()
}

func runGameTickHooks() {
	gameTickMu.Lock()
	defer gameTickMu.Unlock()
	for i, fnc := range gameTickHooks {
		fnc()
		gameTickHooks[i] = nil
	}
	gameTickHooks = gameTickHooks[:0]
}

func nox_xxx_gameTick_4D2580_server() bool {
	defer runGameTickHooks()
	v0 := platformTicks()
	v2 := false
	if C.dword_5d4594_2650652 == 0 {
		C.nox_netlist_resetAllInList_40EE90(1)
	} else {
		v4 := int(C.nox_xxx_rateGet_40A6C0())
		if C.sub_416650() != 0 && sub_41E2F0() == 8 {
			v2 = true
		}
		if v4 == 1 || noxflags.HasGame(8) || gameFrame()%uint32(v4) == 1 {
			C.nox_netlist_resetAllInList_40EE90(1)
		}
	}
	C.sub_502100()
	C.sub_5524C0()
	C.nox_xxx_netMaybeSendAll_552460()
	if getEngineFlag(NOX_ENGINE_FLAG_REPLAY_READ) {
		nox_xxx_replayTickMB_4D3580_net_playback(true)
	}
	if getEngineFlag(NOX_ENGINE_FLAG_LOG_BAND) {
		if v0-memmap.Uint64(0x5D4594, 1548684) > 1000 {
			C.sub_4D3130()
			*memmap.PtrUint64(0x5D4594, 1548684) = v0
		}
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		C.nox_telnet_tick_578FC0()
	}
	if noxflags.HasGame(8) {
		nox_xxx_gameTick_4D2580_server_A(v2)
	} else if C.dword_5d4594_1548524 == 0 {
		if !nox_xxx_gameTick_4D2580_server_B(v0) {
			return true
		}
	} else if C.nox_xxx_check_flag_aaa_43AF70() == 0 || !v2 {
		if !nox_xxx_gameTick_4D2580_server_C() {
			return false
		}
	}
	nox_xxx_gameTick_4D2580_server_E()
	return true
}

func nox_xxx_gameTick_4D2580_server_A(v2 bool) {
	if memmap.Uint64(0x5D4594, 1548676) == 0 {
		*memmap.PtrUint64(0x5D4594, 1548676) = platformTicks() + 10000
		C.nox_xxx_guiServerOptionsHide_4597E0(0)
		if C.nox_xxx_check_flag_aaa_43AF70() == 1 {
			if !noxflags.HasGame(128) {
				C.nox_xxx_net_4263C0()
				C.sub_40DF90()
				C.sub_4264D0()
			}
			if unsafe.Slice((*byte)(unsafe.Pointer(C.sub_4165B0())), 58)[57] == 0 {
				C.sub_41D650()
			}
		}
		netLog.Println("EndGame")
		C.sub_46DCC0()
	}
	nox_server_netMaybeSendInitialPackets_4DEB30()
	nox_xxx_netlist_4DEB50()
	if platformTicks() <= memmap.Uint64(0x5D4594, 1548676) || v2 {
		return
	}
	*memmap.PtrUint32(0x5D4594, 1548676) = 0
	*memmap.PtrUint32(0x5D4594, 1548680) = 0
	C.dword_5d4594_1548524 = 1
	sub_416170(12)
	noxflags.UnsetGame(8)
	for _, u := range getPlayerUnits() {
		u.dropAllItems()
		C.nox_xxx_playerMakeDefItems_4EF7D0(C.int(uintptr(unsafe.Pointer(u.CObj()))), 1, 0)
	}
	C.nox_xxx_unitsNewAddToList_4DAC00()
	if noxflags.HasGame(512) {
		return
	}
	v7 := nox_xxx_cliGamedataGet_416590(0)
	if v7[57] != 0 {
		v8 := GoString(C.sub_409B80()) + ".map"
		nox_xxx_gameSetMapPath_409D70(v8)
		v7[57] = 0
	} else if C.sub_4D0D70() != 0 {
		C.nox_xxx_loadMapCycle_4D0A30()
		v11 := nox_xxx_getSomeMapName_4D0CF0()
		if v11 != "" {
			v12 := nox_server_currentMapGetFilename_409B30()
			if strings.ToLower(v11) == strings.ToLower(v12) {
				v11 = nox_xxx_getSomeMapName_4D0CF0()
			}
			if v11 != "" {
				nox_xxx_gameSetMapPath_409D70(v11)
			}
		}
	}
}

func nox_xxx_cliGamedataGet_416590(v int) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(C.nox_xxx_cliGamedataGet_416590(C.int(v)))), 60)
}

func nox_xxx_gameTick_4D2580_server_D() {
	pl := getPlayerByInd(31)
	if pl == nil {
		return
	}
	u := pl.UnitC()
	if u == nil || C.sub_4DCC10(u.CObj()) != 1 {
		return
	}
	v23 := false
	v24 := (u.Flags16()>>15)&1 != 0
	if !v24 {
		scriptOnEvent("MapShutdown")
		noxflags.SetGame(0x8000000)
		v26 := GoString(C.sub_4DB160())
		v23 = C.nox_xxx_saveDoAutosaveMB_4DB370_savegame(internCStr(v26)) != 0
		noxflags.UnsetGame(0x8000000)
		if !v23 && noxflags.HasGame(2) {
			v35 := strMan.GetStringInFile("GUISave.c:SaveErrorTitle", "C:\\NoxPost\\src\\Server\\System\\server.c")
			NewDialogWindow(nil, v35, v35, 33, nil, nil)
		}
	}
	v28 := uintptr(C.sub_4DB1C0())
	if v24 || !v23 {
		if v28 != 0 && !v23 {
			u.SetPos(asPointf(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(v28 + 700)) + 80)))
		}
	} else if v28 != 0 {
		v30 := GoString(*(**C.char)(unsafe.Pointer(v28 + 700)))
		v31, err := nox_client_checkSaveMapExistsTmp(v30)
		if err == nil && v31 != "" {
			C.nox_xxx_gameSetSwitchSolo_4DB220(1)
			C.nox_xxx_gameSetNoMPFlag_4DB230(1)
			C.nox_xxx_gameSetSoloSavePath_4DB270(internCStr(v31))
		} else {
			gameLog.Println(err)
		}
		switchMap(v30)
	}
	sub_4DB170(0, C.int(v28), 0)
}

func nox_xxx_gameTick_4D2580_server_C() bool {
	C.sub_4EDD70()
	sub_417160()
	C.sub_4573B0()
	if checkGameplayFlags(2) && !noxflags.HasGame(49152) &&
		checkGameplayFlags(4) && !noxflags.HasGame(128) {
		C.sub_4181F0(1)
	}
	if noxflags.HasGame(noxflags.GameModeQuest) && C.nox_xxx_check_flag_aaa_43AF70() == 1 && !noxflags.HasGame(128) {
		C.sub_4264D0()
	}
	noxflags.SetGame(0x8000000)
	sub_43F140(500)
	v13 := nox_xxx_mapExitAndCheckNext_4D1860_server()
	sub_43F1A0()
	noxflags.UnsetGame(0x8000000)
	v37 := getServerMap()
	if !v13 {
		v14 := strMan.GetStringInFile("MapAccessError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		PrintToPlayers(fmt.Sprintf(v14, v37))
		//v36 := strMan.GetStringInFile("Error", "C:\\NoxPost\\src\\Server\\System\\server.c")
		//v15 := strMan.GetStringInFile("MapCorrupt", "C:\\NoxPost\\src\\Server\\System\\server.c")
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		return false
	}
	v39 := v37 + ".map"
	v18 := C.nox_xxx_mapCrcGetMB_409B00()
	C.nox_xxx_netUseMap_4DEE00(internCStr(v39), v18)
	if C.nox_xxx_check_flag_aaa_43AF70() == 1 {
		C.sub_416690()
		if noxflags.HasGame(128) {
			if noxflags.HasGame(0x8000) {
				C.sub_419030(1)
			}
			noxflags.UnsetGame(49152)
		} else {
			C.sub_426060()
			C.sub_41D6C0()
			inputSetKeyTimeoutLegacy(15)
		}
	}
	C.nox_xxx_guiServerOptionsHide_4597E0(0)
	return true
}

func nox_game_guiInit_473680() error {
	*memmap.PtrPtr(0x5D4594, 1096420) = unsafe.Pointer(nox_xxx_gLoadImg("CursorBitmap").C())
	if C.sub_455C30() == 0 {
		return errors.New("sub_455C30 failed")
	}
	if C.sub_456070() == 0 {
		return errors.New("sub_456070 failed")
	}
	if C.sub_470710() == 0 {
		return errors.New("sub_470710 failed")
	}
	if C.nox_xxx_guiHealthManaInit_4714E0() == 0 {
		return errors.New("nox_xxx_guiHealthManaInit_4714E0 failed")
	}
	if C.nox_xxx_bookInit_45B9D0() == 0 {
		return errors.New("nox_xxx_bookInit_45B9D0 failed")
	}
	if C.sub_476E20() == nil {
		return errors.New("sub_476E20 failed")
	}
	if C.sub_4BFAD0() == 0 {
		return errors.New("sub_4BFAD0 failed")
	}
	tmp := C.nox_xxx_wndCreateInventoryMB_465E00()
	*memmap.PtrUint32(0x5D4594, 1096328) = uint32(tmp)
	if tmp == 0 {
		return errors.New("nox_xxx_wndCreateInventoryMB_465E00 failed")
	}
	if C.nox_game_initOptionsInGame_4ADAD0() == 0 {
		return errors.New("nox_game_initOptionsInGame_4ADAD0 failed")
	}
	if C.sub_48D000() == 0 {
		return errors.New("sub_48D000 failed")
	}
	if C.sub_4C3760() == 0 {
		return errors.New("sub_4C3760 failed")
	}
	if C.nox_savegame_sub_46C730() == 0 {
		return errors.New("nox_savegame_sub_46C730 failed")
	}
	if C.sub_4C09D0() == 0 {
		return errors.New("sub_4C09D0 failed")
	}
	if C.sub_478110() == 0 {
		return errors.New("sub_478110 failed")
	}
	if C.sub_49B3E0() == 0 {
		return errors.New("sub_49B3E0 failed")
	}
	if C.sub_4BFC90() == 0 {
		return errors.New("sub_4BFC90 failed")
	}
	if C.nox_gui_itemAmount_init_4BFEF0() == 0 {
		return errors.New("nox_gui_itemAmount_init_4BFEF0 failed")
	}
	if C.sub_4799A0() == 0 {
		return errors.New("sub_4799A0 failed")
	}
	if C.nox_xxx_cliPrepareGameplay1_460E60() == 0 {
		return errors.New("nox_xxx_cliPrepareGameplay1_460E60 failed")
	}
	vsz := videoGetWindowSize()
	*memmap.PtrPtr(0x5D4594, 1096352) = unsafe.Pointer(C.nox_gui_console_Create_450C70(C.int(vsz.W), C.int(vsz.H)))
	if memmap.Uint32(0x5D4594, 1096352) == 0 {
		return errors.New("nox_gui_console_Create_450C70 failed")
	}
	if C.sub_46A730() == nil {
		return errors.New("sub_46A730 failed")
	}
	if C.sub_44E560() == nil {
		return errors.New("sub_44E560 failed")
	}
	if C.sub_4C3500() == 0 {
		return errors.New("sub_4C3500 failed")
	}
	tmp = C.nox_xxx_guiDrawRank_46E870()
	*memmap.PtrUint32(0x5D4594, 1096340) = uint32(tmp)
	if tmp == 0 {
		return errors.New("nox_xxx_guiDrawRank_46E870 failed")
	}
	tmp = C.nox_xxx_guiMotdLoad_4465C0()
	*memmap.PtrUint32(0x5D4594, 1096324) = uint32(tmp)
	if tmp == 0 {
		return errors.New("nox_xxx_guiMotdLoad_4465C0 failed")
	}
	if C.nox_xxx_guiSummonCreatureLoad_4C1D80() == 0 {
		return errors.New("nox_xxx_guiSummonCreatureLoad_4C1D80 failed")
	}
	if C.nox_xxx_wndLoadQuitMenu_445790() == 0 {
		return errors.New("nox_xxx_wndLoadQuitMenu_445790 failed")
	}
	if C.sub_4AB260() == 0 {
		return errors.New("sub_4AB260 failed")
	}
	if C.nox_xxx_guiChatIconLoad_445650() == 0 {
		return errors.New("nox_xxx_guiChatIconLoad_445650 failed")
	}
	if C.sub_4C3390() == 0 {
		return errors.New("sub_4C3390 failed")
	}
	if C.sub_48C980() == 0 {
		return errors.New("sub_48C980 failed")
	}
	C.nox_gui_console_Enable_450BE0()
	C.sub_4AB4A0(0)
	C.sub_4AB4D0(0)
	if C.nox_client_renderGUI_80828 == 0 || getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING) {
		C.sub_4721A0(0)
		C.sub_460EA0(0)
		C.nox_window_set_visible_unk5(0)
		C.sub_45D500(0)
		C.sub_455A00(0)
		C.sub_455F10(0)
	}
	return nil
}

func nox_xxx_mapFindPlayerStart_4F7AB0(a2 *Unit) types.Pointf {
	p, freeCp := alloc.Malloc(unsafe.Sizeof(C.float2{}))
	defer freeCp()
	cp := (*C.float2)(p)
	C.nox_xxx_mapFindPlayerStart_4F7AB0(cp, a2.CObj())
	return types.Pointf{
		X: float32(cp.field_0),
		Y: float32(cp.field_4),
	}
}

func nox_xxx_mapExitAndCheckNext_4D1860_server() bool {
	if noxflags.HasGame(2) {
		nox_client_setCursorType_477610(10)
	}
	C.sub_4D22B0()
	nox_xxx_netMsgFadeBegin_4D9800(false, true)
	scriptOnEvent("MapExit")
	v2 := getServerMap()
	C.sub_500510(internCStr(v2))
	C.nox_xxx_mapSwitchLevel_4D12E0(1)
	for _, obj := range getObjects() {
		obj.SetFlags16(obj.Flags16() | 0x80000000)
	}
	for _, obj := range getObjectsUpdatable2() {
		obj.SetFlags16(obj.Flags16() | 0x80000000)
	}
	if noxflags.HasGame(noxflags.GameModeCoop) {
		C.nox_xxx_spellEnableAll_424BD0()
		C.sub_4537F0()
	}
	var v6 bool
	if C.nox_xxx_gameIsSwitchToSolo_4DB240() != 0 {
		v5 := nox_xxx_mapFilenameGetSolo_4DB260()
		v6 = nox_server_loadMapFile_4CF5F0(v5, false)
	} else {
		v7p := unsafe.Pointer(C.sub_4165B0())
		v7 := unsafe.Slice((*byte)(v7p), 58)
		if noxflags.HasGame(noxflags.GameOnline) {
			if v7[57] == 0 {
				v63, freeV63 := alloc.Bytes(60)
				defer freeV63()
				copy(v63, GoStringS(v7))
				v59 := *(*uint16)(unsafe.Pointer(&v7[52]))
				v8 := C.sub_459870()
				C.sub_57A1E0((*C.int)(unsafe.Pointer(&v63[0])), internCStr("user.rul"), (*C.int)(unsafe.Pointer(v8)), 3, C.short(v59))
				v9 := C.sub_459870()
				C.sub_57AAA0(internCStr("user.rul"), (*C.char)(v7p), (*C.int)(unsafe.Pointer(v9)))
			}
			C.nox_xxx_spellEnableAll_424BD0()
			C.sub_4537F0()
		}
		v10 := nox_server_currentMapGetFilename_409B30()
		v6 = nox_server_loadMapFile_4CF5F0(v10, false)
		if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(128) {
			v13 := 0
			if C.nox_xxx_gamePlayIsAnyPlayers_40A8A0() != 0 {
				v12 := noxflags.GetGame()
				if C.sub_40A180(C.short(v12)) != 0 {
					v13 = 1
				}
			}
			C.sub_40A1F0(C.int(v13))
			C.sub_40A250()
			v60 := noxflags.GetGame()
			v58 := C.sub_459870()
			v14 := nox_xxx_cliGamedataGet_416590(0)
			C.sub_57A1E0((*C.int)(unsafe.Pointer(&v14[0])), internCStr("user.rul"), (*C.int)(unsafe.Pointer(v58)), 3, C.short(v60))
			C.sub_4D2230()
		}
	}
	if !v6 {
		*memmap.PtrUint32(0x5D4594, 1548520) = 1
		return false
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		C.nox_xxx_netMapSendPrepair_519EB0_net_mapsend()
	}
	C.nox_xxx_unitsNewAddToList_4DAC00()
	for _, k := range getPlayerUnits() {
		C.sub_4EF660(k.CObj())
		v61 := nox_xxx_mapFindPlayerStart_4F7AB0(k)
		if noxflags.HasGame(128) {
			if C.nox_xxx_getTeamCounter_417DD0() != 0 {
				if !checkGameplayFlags(2) && !noxflags.HasGame(0x8000) {
					v17 := unsafe.Pointer(C.nox_xxx_clientGetTeamColor_418AB0(C.int(k.team())))
					if v17 != nil {
						v61 = randomReachablePointAround(50.0, asPointf(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v17, 72)), 56)))
					}
				}
			}
		}
		k.SetPos(v61)
		if !noxflags.HasGame(512) {
			plx := k.ControllingPlayer()
			plx.field_2136 = 0
			plx.field_2140 = 0
			plx.field_2144 = C.uint(gameFrame())
			C.nox_xxx_netReportLesson_4D8EF0(k.CObj())
		}
	}
	C.sub_50AFA0()
	C.sub_50B520()
	*memmap.PtrUint32(0x5D4594, 1548528) = 0
	C.dword_5d4594_1548532 = nil
	if memmap.Uint32(0x5D4594, 1548708) == 0 {
		*memmap.PtrUint32(0x5D4594, 1548708) = uint32(getObjectTypeByID("PlayerStart").Ind())
	}
	if memmap.Uint32(0x5D4594, 1548712) == 0 {
		*memmap.PtrUint32(0x5D4594, 1548712) = uint32(getObjectTypeByID("GameBallStart").Ind())
	}
	var starts struct {
		playerN int
		flagN   int
		ballN   int
	}
	for _, v22 := range getObjects() {
		if v22.Class().Has(0x10000000) {
			v22.field_34 = 0
			C.dword_5d4594_1548532 = unsafe.Pointer(v22.CObj())
			starts.flagN++
		}
		ti := uint32(v22.objTypeInd())
		if ti == memmap.Uint32(0x5D4594, 1548712) {
			starts.ballN++
		}
		if ti == memmap.Uint32(0x5D4594, 1548708) {
			starts.playerN++
		}
	}
	if starts.playerN == 0 {
		v24 := strMan.GetStringInFile("StartingPositionError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		PrintToPlayers(v24)
		return false
	}
	if noxflags.HasGame(96) && starts.flagN < 2 {
		v24 := strMan.GetStringInFile("FlagCountError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		PrintToPlayers(v24)
		return false
	}
	if noxflags.HasGame(64) && starts.ballN < 1 {
		v24 := strMan.GetStringInFile("BallStartCountError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		PrintToPlayers(v24)
		return false
	}

	if checkGameplayFlags(4) || noxflags.HasGame(0x8000) {
		C.sub_417D00()
	}
	for l := C.nox_xxx_getDebugData_57C3E0(); l != nil; l = C.nox_xxx_nextDebugObject_57C3F0(l) {
		switch str := GoString((*C.char)(l)); str {
		case "spring":
			str2 := GoString((*C.char)(unsafe.Add(l, 80)))
			str2 = strings.TrimSpace(str2)
			if sub := strings.Fields(str2); len(sub) >= 2 {
				v28, _ := strconv.Atoi(sub[0])
				v29 := getObjectByInd(v28)
				v31, _ := strconv.Atoi(sub[1])
				v32 := getObjectByInd(v31)
				if v29 != nil && v32 != nil {
					C.nox_xxx_addDebugEntry_511590(unsafe.Pointer(v29.CObj()), unsafe.Pointer(v32.CObj()))
				}
			}
		case "SentryGlobe":
			str2 := GoString((*C.char)(unsafe.Add(l, 80)))
			str2 = strings.TrimSpace(str2)
			if sub := strings.Fields(str2); len(sub) >= 3 {
				v34, _ := strconv.Atoi(sub[0])
				if v35 := getObjectByInd(v34).AsUnit(); v35 != nil {
					v36 := unsafe.Slice((*float32)(v35.updateDataPtr()), 3)
					v38, _ := strconv.ParseFloat(sub[1], 64)
					v36[0] = float32(v38 * memmap.Float64(0x581450, 10008))
					v36[1] = v36[0]
					v40, _ := strconv.ParseFloat(sub[2], 64)
					v36[2] = float32(v40 * memmap.Float64(0x581450, 10000) / float64(gameFPS()))
				}
			}
		}
	}
	C.dword_5d4594_1548524 = 0
	v41 := getServerMap()
	C.sub_500510(internCStr(v41))
	if C.nox_xxx_gameIsSwitchToSolo_4DB240() == 0 {
		C.nox_xxx_resetMapInit_4FC570(1)
	}
	noxflags.SetGame(0x8000000)
	sub_43F140(500)
	v42 := C.nox_xxx_gameIsSwitchToSolo_4DB240()
	C.sub_4DBA30(v42)
	sub_43F1A0()
	noxflags.UnsetGame(0x8000000)
	C.sub_4FC580(1)
	if mapSwitchWPName != "" {
		if wp := getWaypointByID(mapSwitchWPName); wp != nil {
			gameLog.Printf("moving player to waypoint: %q", mapSwitchWPName)
			wpos := wp.Pos()
			for _, u := range getPlayerUnits() {
				u.SetPos(wpos)
			}
		} else {
			gameLog.Printf("cannot find map waypoint %q!", mapSwitchWPName)
		}
	}
	mapSwitchWPName = ""
	if sub_4DCC00() {
		for _, m := range getPlayerUnits() {
			for _, np := range m.GetOwned516() {
				var np *Object = np
				if C.nox_xxx_isUnit_4E5B50(np.CObj()) != 0 {
					n := np.AsUnit()
					ud := n.updateDataMonster()
					v61 := randomReachablePointAround(50.0, m.Pos())
					n.SetPos(v61)
					ud.field_97 = 0
					n.clearActionStack()
					n.field_130 = 0
					ud.current_enemy = nil
					ud.field_304 = 0
					*(*byte)(unsafe.Add(unsafe.Pointer(ud), 1129)) = 0
					ud.field_0 = 0xDEADFACE
					p := m.Pos()
					n.monsterPushAction(ai.ACTION_ESCORT, p.X, p.Y, m)
					if n.Class().Has(2) && n.field_3&0x30 != 0 {
						C.nox_xxx_setNPCColor_4E4A90(n.CObj(), 0, C.int(ud.field_519))
					}
				} else if m.Class().Has(1) && C.sub_4E5B80(m.CObj()) != 0 {
					C.sub_4E81D0(m.CObj())
				}
			}
		}
		C.sub_4DCBF0(0)
	}
	for _, obj := range getObjects() {
		obj.SetFlags16(obj.Flags16() & 0x7FFFFFFF)
	}
	for _, obj := range getObjectsUpdatable2() {
		obj.SetFlags16(obj.Flags16() & 0x7FFFFFFF)
	}
	if noxflags.HasGame(16) && checkGameplayFlags(4) {
		C.sub_4D2160()
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		nox_server_questMapNextLevel()
	}
	if noxflags.HasGame(noxflags.GameModeCoop) && C.nox_xxx_mapLoadRequired_4DCC80() == 0 {
		C.sub_4DB130(internCStr(common.SaveAuto))
		sub_4DB170(1, 0, 30)
	}
	nox_xxx_mapLoadOrSaveMB_4DCC70(0)
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_413980(30)
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		sub_4D7140(gameFrame() + gameFPS())
	} else {
		sub_4D7140(0)
	}
	nox_ticks_reset_416D40()
	C.sub_40A970()
	C.dword_5d4594_1548704 = 0
	if noxflags.HasGame(noxflags.GameModeQuest) {
		C.sub_4D71E0(0)
		if C.sub_4D72C0() == 1 {
			C.sub_4D72B0(0)
			v57 := C.sub_4D72C0()
			C.sub_4D7280(255, C.char(v57))
		}
	}
	C.sub_413850()
	return true
}

func nox_xxx_mapLoadOrSaveMB_4DCC70(v int) {
	C.nox_xxx_mapLoadOrSaveMB_4DCC70(C.int(v))
}

func sub_4D7140(a1 uint32) {
	C.dword_5d4594_1556144 = C.uint(a1)
}

func sub_4DCC00() bool {
	return C.dword_5d4594_1563064 != 0
}

func sub_413980(a1 int) {
	if a1 != 0 {
		sub_413A00(1)
		C.dword_5d4594_251744 = C.uint(a1 + int(nox_xxx_bookGet_430B40_get_mouse_prev_seq()))
	} else {
		C.dword_5d4594_251744 = 0
	}
}

func sub_4139C0() {
	if C.dword_5d4594_251744 != 0 {
		if int(nox_xxx_bookGet_430B40_get_mouse_prev_seq()) >= int(C.dword_5d4594_251744) {
			C.dword_5d4594_251744 = 0
			if C.sub_450560() == 0 {
				sub_413A00(0)
			}
		}
	}
}

func sub_417160() {
	for _, it := range getPlayers() {
		C.nox_xxx_playerUnsetStatus_417530(it.C(), 16)
	}
}

func randomReachablePointAround(a1 float32, v3 types.Pointf) types.Pointf { // sub_4ED970
	var v10 [4]float32
	v10[0] = v3.X
	v10[1] = v3.Y
	v9 := a1 * 0.015625
	v11 := float32(noxRndCounter1.FloatClamp(-math.Pi, math.Pi))
	for v5 := 0; v5 < 64; v5++ {
		v6 := v11 + 1.8849558
		v11 = v6
		v10[2] = float32(math.Cos(float64(v6)))*a1 + v3.X
		v10[3] = float32(math.Sin(float64(v11)))*a1 + v3.Y
		if nox_xxx_mapTraceRay_535250_00(&v10, 1) {
			return types.Pointf{
				X: v10[2], Y: v10[3],
			}
		}
		a1 = a1 - v9
	}
	return v3
}

func nox_xxx_mapTraceRay_535250_00(a1 *[4]float32, a4 byte) bool {
	a1p, a1Free := alloc.Malloc(16)
	defer a1Free()
	a1c := (*[4]float32)(a1p)
	*a1c = *a1
	res := C.nox_xxx_mapTraceRay_535250((*C.float4)(unsafe.Pointer(a1c)), nil, nil, C.char(a4)) != 0
	*a1 = *a1c
	return res
}

//export nox_xxx_mapTraceObstacles_50B580
func nox_xxx_mapTraceObstacles_50B580(from *C.nox_object_t, a2 *C.float4) C.int {
	if nox_xxx_mapTraceObstacles(asUnitC(from), types.Pointf{
		X: float32(a2.field_0),
		Y: float32(a2.field_4),
	}, types.Pointf{
		X: float32(a2.field_8),
		Y: float32(a2.field_C),
	}) {
		return 1
	}
	return 0
}

func nox_xxx_mapTraceObstacles(from *Unit, p1, p2 types.Pointf) bool { // nox_xxx_mapTraceObstacles_50B580
	rect := types.RectFromPointsf(p1, p2)

	pp, ppFree := alloc.Malloc(16)
	defer ppFree()
	p := (*[4]float32)(pp)
	p[0], p[1] = p1.X, p1.Y
	p[2], p[3] = p2.X, p2.Y

	searching := true
	getUnitsInRect(rect, func(obj *Object) {
		if !searching {
			return
		}
		if from.CObj() == obj.CObj() {
			return
		}
		if obj.Class().HasAny(object.MaskUnits) {
			u2 := obj.AsUnit()
			if from.isEnemyTo(u2) {
				return
			}
		} else if !obj.Class().HasAny(object.ClassImmobile | object.ClassObstacle) {
			return
		}
		if (obj.Flags16()&0x48 != 0) || obj.Class().Has(object.ClassDoor) {
			return
		}
		pos := obj.Pos()
		sh := obj.getShape()
		switch sh.kind {
		case shapeKindCircle:
			a3p, a3Free := alloc.Malloc(8)
			defer a3Free()
			a3 := unsafe.Slice((*float32)(a3p), 2)
			if C.nox_xxx_mathPointOnTheLine_57C8A0((*C.float4)(pp), (*C.float2)(unsafe.Pointer(&obj.x)), (*C.float2)(a3p)) != 0 {
				dx := a3[0] - pos.X
				dy := a3[1] - pos.Y
				if dy*dy+dx*dx <= sh.circle.R2 {
					searching = false
				}
			}
		case shapeKindBox:
			a2p, a2Free := alloc.Malloc(16)
			defer a2Free()
			a2 := unsafe.Slice((*float32)(a2p), 4)

			v12 := sh.box.LeftTop + pos.X
			v5 := sh.box.LeftBottom + pos.Y
			a2[0] = v12
			v13 := v5
			v6 := sh.box.LeftBottom2 + pos.X
			a2[1] = v13
			v9 := v6
			v7 := sh.box.LeftTop2 + pos.Y
			a2[2] = v9
			v10 := v7
			v8 := sh.box.RightTop + pos.X
			a2[3] = v10
			v11 := sh.box.RightBottom + pos.Y
			xx := sh.box.RightBottom2 + pos.X
			yy := sh.box.RightTop2 + pos.Y
			if C.sub_427980((*C.float4)(pp), (*C.float4)(a2p)) != 0 {
				searching = false
				return
			}
			a2[0] = v12
			a2[1] = v13
			a2[2] = v8
			a2[3] = v11
			if C.sub_427980((*C.float4)(pp), (*C.float4)(a2p)) != 0 {
				searching = false
				return
			}
			a2[0] = xx
			a2[1] = yy
			a2[2] = v8
			a2[3] = v11
			if C.sub_427980((*C.float4)(pp), (*C.float4)(a2p)) != 0 {
				searching = false
				return
			}
			a2[0] = xx
			a2[1] = yy
			a2[2] = v9
			a2[3] = v10
			if C.sub_427980((*C.float4)(pp), (*C.float4)(a2p)) != 0 {
				searching = false
			}
		}
	})
	return searching
}

var (
	// TODO: port nox_xxx_getUnitsInRect_517C10 instead
	getUnitsInRectStack []func(obj *Object)
)

//export nox_xxx_getUnitsInRect_517C10_go
func nox_xxx_getUnitsInRect_517C10_go(obj *C.nox_object_t, _ unsafe.Pointer) {
	getUnitsInRectStack[len(getUnitsInRectStack)-1](asObjectC(obj))
}

func getUnitsInRect(rect types.Rectf, fnc func(obj *Object)) {
	getUnitsInRectStack = append(getUnitsInRectStack, fnc)
	defer func() {
		getUnitsInRectStack = getUnitsInRectStack[:len(getUnitsInRectStack)-1]
	}()
	_ = nox_xxx_getUnitsInRect_517C10_go
	nox_xxx_getUnitsInRect_517C10(rect, C.nox_xxx_getUnitsInRect_517C10_go, nil)
}

func nox_xxx_getUnitsInRect_517C10(rect types.Rectf, fnc unsafe.Pointer, payload unsafe.Pointer) {
	rp, rpFree := alloc.Malloc(16)
	defer rpFree()
	rr := (*types.Rectf)(rp)
	*rr = rect
	C.nox_xxx_getUnitsInRect_517C10((*C.float4)(rp), (*[0]byte)(fnc), payload)
}

func nox_xxx_calcDistance_4E6C00(obj1, obj2 noxObject) float32 {
	return float32(C.nox_xxx_calcDistance_4E6C00(obj1.CObj(), obj2.CObj()))
}

func nox_xxx_unitCanInteractWith_5370E0(u, obj noxObject, a3 byte) bool {
	return C.nox_xxx_unitCanInteractWith_5370E0(u.CObj(), obj.CObj(), C.char(a3)) != 0
}

func sub_473840() {
	C.nox_gui_console_Disable_450BF0()
	C.nox_gui_itemAmount_free_4C03E0()
	sub_46CCB0()
	C.sub_4AE3B0()
	C.sub_48D450()
	C.sub_4C4220()
	C.nox_xxx_closeP2PTradeWnd_4C12A0()
	C.sub_4BFD10()
	C.sub_49B490()
	C.sub_478F80()
	C.sub_479D10()
	C.sub_4AB470()
	C.sub_4C34A0()
	C.sub_445770()
	C.sub_456240()
	C.sub_455EE0()
	sub_460E90()
	C.sub_4505E0()
	C.sub_46A860()
	C.sub_49C7A0()
	C.nox_xxx_guiServerOptionsHide_4597E0(1)
	C.sub_467980()
	sub_46C5D0()
	C.nox_client_renderGUI_80828 = C.nox_xxx_xxxRenderGUI_587000_80832
}

func sub_460E90() {
	if sub_460D40() {
		C.sub_460D50()
	}
}

func sub_460D40() bool {
	return C.dword_5d4594_1049508 != 0
}

const netListsBufSize = 2048

func nox_xxx_netlist_4DEB50() {
	if !noxflags.HasGame(2) {
		return
	}
	if getEngineFlag(NOX_ENGINE_FLAG_REPLAY_READ) {
		nox_xxx_replayTickMB_4D3580_net_playback(false)
		C.nox_netlist_resetByInd_40ED10(31, 0)
	} else if !isDedicatedServer {
		buf := nox_netlist_copyPacketList_40ED60(31, 0)
		if len(buf) != 0 {
			dst := unsafe.Slice((*byte)(unsafe.Pointer(&C.nox_net_lists_buf[0])), netListsBufSize)
			n := copy(dst, buf)
			C.nox_xxx_netOnPacketRecvServ_51BAD0_net_sdecode(31, &C.nox_net_lists_buf[0], C.int(n))
		}
		C.nox_netlist_resetByInd_40ED10(31, 0)
	}
}

func nox_xxx_netOnPacketRecvServ_51BAD0_net_sdecode(ind int, data []byte) int {
	cdata, cfree := alloc.Bytes(uintptr(len(data)))
	defer cfree()
	return int(C.nox_xxx_netOnPacketRecvServ_51BAD0_net_sdecode(C.int(ind), (*C.uchar)(unsafe.Pointer(&cdata[0])), C.int(len(data))))
}

var (
	nox_game_state_arr []int
)

const (
	gameStateMovies      = 10
	gameStateMainMenu    = 100
	gameStateOptions     = 300
	gameStateOnlineOrLAN = 400
	gameStateCharSelect  = 500
	gameStateClassSelect = 600
	gameStateColorSelect = 700
	gameStateWolLogin    = 1700
	gameStateWolChat     = 1900
	gameStateServerList  = 10000
	gameStateXxx         = 1915
)

//export nox_game_addStateCode_43BDD0
func nox_game_addStateCode_43BDD0(code C.int) {
	gameAddStateCode(int(code))
}

//export nox_game_getStateCode_43BE10
func nox_game_getStateCode_43BE10() C.int {
	return C.int(gameGetStateCode())
}

//export nox_game_decStateInd_43BDC0
func nox_game_decStateInd_43BDC0() {
	gamePopState()
}

func gameAddStateCode(code int) {
	if gameGetStateCode() == code {
		return
	}
	nox_game_state_arr = append(nox_game_state_arr, code)
	gameLog.Println("game state code:", code)
}

func gameGetStateCode() int {
	if len(nox_game_state_arr) == 0 {
		return -1
	}
	return nox_game_state_arr[len(nox_game_state_arr)-1]
}

func gamePopState() {
	if len(nox_game_state_arr) > 0 {
		nox_game_state_arr = nox_game_state_arr[:len(nox_game_state_arr)-1]
	}
	gameLog.Println("game state code:", gameGetStateCode())
}

func gamePopStateUntil(code int) {
	for gameGetStateCode() != code {
		gamePopState()
	}
}

//export nox_game_switchStates_43C0A0
func nox_game_switchStates_43C0A0() C.int { // switch game states
	switch gameGetStateCode() {
	case gameStateMovies:
		if !nox_game_rollLogoAndStart_4AB1F0() {
			return 0
		}
		return 1
	case gameStateMainMenu: // main menu
		if C.nox_game_showMainMenu_4A1C00() == 0 {
			return 0
		}
		return 1
	case gameStateOptions: // main menu options
		if C.nox_game_showOptions_4AA6B0() == 0 {
			return 0
		}
		return 1
	case gameStateOnlineOrLAN: // online or LAN
		if C.nox_game_showOnlineOrLAN_413800() == 0 {
			return 0
		}
		return 1
	case gameStateCharSelect: // character selection
		C.sub_4A7A70(1)
		if C.nox_game_showSelChar_4A4DB0() == 0 {
			return 0
		}
		return 1
	case gameStateClassSelect:
		if C.nox_game_showSelClass_4A4840() == 0 {
			return 0
		}
		return 1
	case gameStateColorSelect:
		if C.nox_game_showSelColor_4A5D00() == 0 {
			return 0
		}
		return 1
	case gameStateWolLogin:
		if C.nox_game_showWolLogin_44A560() == 0 {
			return 0
		}
		return 1
	case gameStateWolChat:
		if C.nox_game_showWolChat_447620() == 0 {
			return 0
		}
		return 1
	case gameStateServerList: // list of servers
		if C.nox_xxx_check_flag_aaa_43AF70() == 1 && C.sub_40E0B0() == 0 {
			sub_41E300(9)
			C.sub_41F4B0()
			C.sub_41EC30()
			C.sub_446490(0)
			C.nox_xxx____setargv_4_44B000()
			sub_44A400()
			return 1
		}
		fallthrough
	case gameStateXxx:
		if C.nox_game_showGameSel_4379F0() == 0 {
			return 0
		}
		return 1
	}
	return 1
}

//export nox_game_checkStateSwitch_43C1E0
func nox_game_checkStateSwitch_43C1E0() {
	if gameGetStateCode() < 0 {
		return
	}
	p := nox_gui_findAnimationForDest_43C520(gameGetStateCode())
	if p != nil {
		p.Func12()
		p.field_13 = (*[0]byte)(C.nox_game_switchStates_43C0A0)
		gamePopState()
	}
}

//export nox_game_checkStateOptions_43C220
func nox_game_checkStateOptions_43C220() {
	if gameGetStateCode() < 0 {
		return
	}
	p := nox_gui_findAnimationForDest_43C520(gameGetStateCode())
	if p != nil {
		p.Func12()
		p.field_13 = (*[0]byte)(C.nox_game_showOptions_4AA6B0)
	}
}

//export nox_game_checkStateWol_43C260
func nox_game_checkStateWol_43C260() {
	if gameGetStateCode() < 0 {
		return
	}
	p := nox_gui_findAnimationForDest_43C520(gameGetStateCode())
	if p != nil {
		if gameGetStateCode() == gameStateWolLogin {
			C.nox_game_showWolLogin_44A560()
		} else {
			p.Func12()
			p.field_13 = (*[0]byte)(C.nox_game_showWolLogin_44A560)
		}
	}
	gamePopStateUntil(gameStateWolLogin)
}

//export nox_game_checkStateMenu_43C2F0
func nox_game_checkStateMenu_43C2F0() {
	if gameGetStateCode() < 0 {
		return
	}
	p := nox_gui_findAnimationForDest_43C520(gameGetStateCode())
	if p != nil {
		if gameGetStateCode() == gameStateMainMenu {
			C.nox_game_showMainMenu_4A1C00()
		} else {
			p.Func12()
			p.field_13 = (*[0]byte)(C.nox_game_showMainMenu_4A1C00)
		}
	}
	gamePopStateUntil(gameStateMainMenu)
}
