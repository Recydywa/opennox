package nox

/*
#include "GAME2_2.h"
#include "GAME3.h"
#include "GAME3_1.h"
#include "win.h"
#include "client__gui__gadgets__listbox.h"
*/
import "C"
import (
	"image"
	"unsafe"

	"nox/v1/client/gui"
	"nox/v1/common/alloc"
	noxcolor "nox/v1/common/color"
)

type guiWidgetData interface {
	cWidgetData() unsafe.Pointer
}

type rawWidgetData struct {
	Ptr unsafe.Pointer
}

func (d rawWidgetData) cWidgetData() unsafe.Pointer {
	return unsafe.Pointer(d.Ptr)
}

type radioButtonData C.nox_radioButton_data

func (d *radioButtonData) cWidgetData() unsafe.Pointer {
	return unsafe.Pointer(d)
}

type sliderData C.nox_slider_data

func (d *sliderData) cWidgetData() unsafe.Pointer {
	return unsafe.Pointer(d)
}

type scrollListBoxData C.nox_scrollListBox_data

func (d *scrollListBoxData) cWidgetData() unsafe.Pointer {
	return unsafe.Pointer(d)
}

type entryFieldData C.nox_entryField_data

func (d *entryFieldData) cWidgetData() unsafe.Pointer {
	return unsafe.Pointer(d)
}

type staticTextData C.nox_staticText_data

func (d *staticTextData) cWidgetData() unsafe.Pointer {
	return unsafe.Pointer(d)
}

func guiNewWidget(typ string, parent *Window, status gui.StatusFlags, px, py, w, h int, draw *WindowData, data guiWidgetData) *Window {
	draw.win = parent.C()
	iparent := unsafePtrToInt(unsafe.Pointer(parent.C()))
	udraw := unsafe.Pointer(draw.C())
	switch typ {
	case "PUSHBUTTON":
		draw.style |= C.int(gui.StylePushButton)
		return newButtonOrCheckbox(parent, status, px, py, w, h, draw)
	case "RADIOBUTTON":
		tdata, _ := data.(*radioButtonData)
		draw.style |= C.int(gui.StyleRadioButton)
		return newRadioButton(parent, status, px, py, w, h, draw, tdata)
	case "CHECKBOX":
		draw.style |= C.int(gui.StyleCheckBox)
		return newButtonOrCheckbox(parent, status, px, py, w, h, draw)
	case "VERTSLIDER":
		tdata, _ := data.(*sliderData)
		draw.style |= C.int(gui.StyleVertSlider)
		return asWindow(C.nox_gui_newSlider_4B4EE0(iparent, C.int(status), C.int(px), C.int(py), C.int(w), C.int(h), (*C.uint)(udraw), (*C.float)(unsafe.Pointer(tdata))))
	case "HORZSLIDER":
		tdata, _ := data.(*sliderData)
		draw.style |= C.int(gui.StyleHorizSlider)
		return asWindow(C.nox_gui_newSlider_4B4EE0(iparent, C.int(status), C.int(px), C.int(py), C.int(w), C.int(h), (*C.uint)(udraw), (*C.float)(unsafe.Pointer(tdata))))
	case "SCROLLLISTBOX":
		tdata, _ := data.(*scrollListBoxData)
		draw.style |= C.int(gui.StyleScrollListBox)
		return asWindow(C.nox_gui_newScrollListBox_4A4310(iparent, C.int(status), C.int(px), C.int(py), C.int(w), C.int(h), dataPtrToInt(draw), (*C.nox_scrollListBox_data)(unsafe.Pointer(tdata))))
	case "ENTRYFIELD":
		tdata, _ := data.(*entryFieldData)
		draw.style |= C.int(gui.StyleEntryField)
		return asWindow(C.nox_gui_newEntryField_488500(iparent, C.int(status), C.int(px), C.int(py), C.int(w), C.int(h), dataPtrToInt(draw), (*C.ushort)(unsafe.Pointer(tdata))))
	case "STATICTEXT":
		tdata, _ := data.(*staticTextData)
		draw.style |= C.int(gui.StyleStaticText)
		return newStaticText(parent, status, px, py, w, h, draw, tdata)
	case "PROGRESSBAR":
		draw.style |= C.int(gui.StyleProgressBar)
		return asWindow(C.nox_gui_newProgressBar_4CAF10(iparent, C.int(status), C.int(px), C.int(py), C.int(w), C.int(h), (*C.uint)(udraw)))
	}
	return nil
}

func tempDrawData() (*WindowData, func()) {
	p, free := alloc.Calloc(1, unsafe.Sizeof(WindowData{}))
	d := (*WindowData)(p)
	d.SetHighlightColor(noxcolor.RGBColor(255, 255, 255))
	d.SetTextColor(noxcolor.RGBColor(200, 200, 200))
	d.SetEnabledColor(gui.ColorTransparent)
	d.SetDisabledColor(gui.ColorTransparent)
	d.SetSelectedColor(gui.ColorTransparent)
	d.SetBackgroundColor(gui.ColorTransparent)
	return d, free
}

func NewStaticText(par *Window, id uint, px, py, w, h int, center, f2 bool, text string) *Window {
	draw, dfree := tempDrawData()
	defer dfree()
	*draw = *par.DrawData()

	draw.win = par.C()
	draw.style |= C.int(gui.StyleStaticText)
	status := gui.StatusSmoothText | gui.StatusNoFocus

	win := newStaticText(par, status, px, py, w, h, draw, &staticTextData{
		text:    internWStr(text),
		center:  C.uint(bool2int(center)),
		field_2: C.uint(bool2int(f2)),
	})
	win.SetID(id)
	if par != nil {
		par.Func94(22, uintptr(id), 0)
	}
	return win
}

//export nox_gui_newStaticText_489300
func nox_gui_newStaticText_489300(par *C.nox_window, status C.int, px, py, w, h C.int, draw *C.nox_window_data, data *C.nox_staticText_data) *C.nox_window {
	return newStaticText(asWindow(par), gui.StatusFlags(status), int(px), int(py), int(w), int(h), asWindowData(draw), (*staticTextData)(unsafe.Pointer(data))).C()
}

func newStaticText(par *Window, status gui.StatusFlags, px, py, w, h int, draw *WindowData, data *staticTextData) *Window { // nox_gui_newStaticText_489300
	style := draw.StyleFlags()
	style &= 0xFFFFFBFF
	draw.SetStyleFlags(style)
	if style&gui.StyleStaticText == 0 {
		return nil
	}
	_ = nox_xxx_wndStaticProcPre_489390
	win := newWindowRaw(par, status, px, py, w, h, C.nox_xxx_wndStaticProcPre_489390)
	if !win.Flags().Has(gui.StatusImage) {
		win.SetAllFuncs(C.nox_xxx_wndStaticProc_489420, C.nox_xxx_wndStaticDrawNoImage_488D00, nil)
	} else {
		_ = nox_xxx_wndStaticDrawWithImage_489550
		win.SetAllFuncs(C.nox_xxx_wndStaticProc_489420, C.nox_xxx_wndStaticDrawWithImage_489550, nil)
	}
	if draw.win == nil {
		draw.win = win.C()
	}
	win.CopyDrawData(draw)
	datap, _ := alloc.Malloc(unsafe.Sizeof(staticTextData{}))
	wdata := (*staticTextData)(datap)
	*wdata = *data
	win.widget_data = unsafe.Pointer(wdata)
	return win
}

const (
	guiEventStaticTextPref    = 0x4000
	guiEventStaticTextSetText = guiEventStaticTextPref + 1
	guiEventStaticTextGetText = guiEventStaticTextPref + 2
)

//export nox_xxx_wndStaticProcPre_489390
func nox_xxx_wndStaticProcPre_489390(cwin *C.nox_window, ev C.int, a3, a4 C.int) C.int {
	win := asWindow(cwin)
	switch ev {
	case 2:
		alloc.Free(win.widget_data)
		win.widget_data = nil
	case guiEventStaticTextSetText:
		if a3 != 0 {
			ptr := (*C.wchar_t)(unsafe.Pointer(uintptr(a3)))
			data := (*staticTextData)(win.widget_data)
			data.text = ptr
		}
	case guiEventStaticTextGetText:
		data := (*staticTextData)(win.widget_data)
		return C.int(uintptr(unsafe.Pointer(data.text)))
	}
	return 0
}

//export nox_xxx_wndStaticDrawWithImage_489550
func nox_xxx_wndStaticDrawWithImage_489550(cwin *C.nox_window, cdraw *C.nox_window_data) C.int {
	win := asWindow(cwin)
	draw := asWindowData(cdraw)
	r := noxrend

	data := (*staticTextData)(unsafe.Pointer(win.widget_data))

	wpos := win.GlobalPos()
	wsz := win.Size()
	fnt := draw.Font()
	fh := r.FontHeight(fnt)
	if win.Flags().Has(gui.StatusSmoothText) {
		r.text.smooth = true
	}
	if bg := draw.BackgroundImage(); bg != nil {
		r.DrawImageAt(bg, wpos.Add(draw.ImagePoint()))
	}
	var img *Image
	if win.Flags().Has(gui.StatusEnabled) {
		img = draw.EnabledImage()
	} else {
		img = draw.DisabledImage()
	}
	if img != nil {
		r.DrawImageAt(img, wpos.Add(draw.ImagePoint()))
	}
	if str := GoWString(data.text); str != "" {
		r.Data().SetTextColor(draw.TextColorRaw())
		y0 := wpos.Y + wsz.H/2 - fh/2
		if data.center != 0 {
			tsz := r.GetStringSizeWrapped(fnt, str, 0)
			x0 := wpos.X + (wsz.W-tsz.W)/2
			rect := image.Rect(x0, y0, x0+wsz.W, y0)
			r.DrawStringWrapped(fnt, str, rect)
		} else {
			x0 := wpos.X + 2
			rect := image.Rect(x0, y0, x0+wsz.W, y0)
			r.DrawStringWrapped(fnt, str, rect)
		}
	}
	r.text.smooth = false
	return 1
}

func NewHorizontalSlider(par *Window, id uint, px, py, w, h int, min, max int) *Window {
	draw, dfree := tempDrawData()
	defer dfree()
	*draw = *par.DrawData()

	draw.win = par.C()
	draw.style |= C.int(gui.StyleHorizSlider)
	draw.SetHighlightColor(gui.ColorTransparent)
	draw.SetDisabledColor(gui.ColorTransparent)
	draw.SetEnabledColor(noxcolor.RGBColor(230, 165, 65))
	draw.SetSelectedColor(noxcolor.RGBColor(230, 165, 65))
	status := gui.StatusEnabled | gui.StatusNoFocus

	datap, dataFree := alloc.Calloc(1, unsafe.Sizeof(sliderData{}))
	defer dataFree()
	data := (*sliderData)(datap)
	data.field_0 = C.uint(min)
	data.field_1 = C.uint(max)
	data.field_2 = 0
	data.field_3 = 0

	iparent := unsafePtrToInt(unsafe.Pointer(par.C()))
	win := asWindow(C.nox_gui_newSlider_4B4EE0(iparent, C.int(status), C.int(px), C.int(py), C.int(w), C.int(h),
		(*C.uint)(unsafe.Pointer(draw)), (*C.float)(unsafe.Pointer(data))))
	win.SetID(id)
	if par != nil {
		par.Func94(22, uintptr(id), 0)
	}
	return win
}

func NewCheckbox(par *Window, id uint, px, py, w, h int, text string) *Window {
	draw, dfree := tempDrawData()
	defer dfree()
	*draw = *par.DrawData()

	draw.win = par.C()
	draw.style = C.int(gui.StyleCheckBox | gui.StyleMouseTrack)
	draw.SetHighlightColor(noxcolor.RGBColor(192, 128, 128))
	draw.SetTextColor(noxcolor.RGBColor(240, 180, 42))
	draw.SetText(text)
	draw.SetBackgroundImage(nox_xxx_gLoadImg("UICheckBox"))
	draw.SetSelectedImage(nox_xxx_gLoadImg("UICheckBoxLit"))
	draw.SetDisabledImage(nox_xxx_gLoadImg("UICheckBox"))
	draw.SetEnabledImage(nil)
	draw.SetHighlightImage(nil)
	status := gui.StatusEnabled | gui.StatusImage | gui.StatusNoFocus

	win := newButtonOrCheckbox(par, status, px, py, w, h, draw)
	win.SetID(id)
	if par != nil {
		par.Func94(22, uintptr(id), 0)
	}
	return win
}

func NewRadioButton(par *Window, id uint, px, py, w, h int, group int, text string) *Window {
	draw, dfree := tempDrawData()
	defer dfree()

	*draw = *par.DrawData()

	draw.win = par.C()
	draw.style = C.int(gui.StyleRadioButton | gui.StyleMouseTrack)
	draw.SetTextColor(noxcolor.RGBColor(240, 180, 42))
	draw.SetText(text)
	draw.SetBackgroundImage(nox_xxx_gLoadImg("UIRadio"))
	draw.SetSelectedImage(nox_xxx_gLoadImg("UIRadioLit"))
	draw.SetDisabledImage(nox_xxx_gLoadImg("UIRadio"))
	draw.SetEnabledImage(nil)
	draw.SetHighlightImage(nil)
	draw.SetGroup(group)
	status := gui.StatusEnabled | gui.StatusToggle | gui.StatusImage | gui.StatusSmoothText | gui.StatusNoFocus

	rdata, rfree := alloc.Calloc(1, unsafe.Sizeof(radioButtonData{}))
	defer rfree()

	win := newRadioButton(par, status, px, py, w, h, draw, (*radioButtonData)(rdata))
	win.SetID(id)
	if par != nil {
		par.Func94(22, uintptr(id), 0)
	}
	return win
}

func newButtonOrCheckbox(parent *Window, status gui.StatusFlags, px, py, w, h int, draw *WindowData) *Window {
	st := draw.StyleFlags()
	if st.IsPushButton() {
		btn := newWindowRaw(parent, status, px, py, w, h, C.nox_xxx_wndButtonProcPre_4A9250)
		if btn == nil {
			return nil
		}
		C.nox_xxx_wndButtonInit_4A8340(unsafePtrToInt(unsafe.Pointer(btn.C())))
		if draw.win == nil {
			draw.win = btn.C()
		}
		btn.CopyDrawData(draw)
		return btn
	} else if st.IsCheckBox() {
		btn := newWindowRaw(parent, status, px, py, w, h, C.nox_xxx_wndCheckboxProcMB_4A92C0)
		if btn == nil {
			return nil
		}
		C.nox_xxx_wndCheckBoxInit_4A8E60(unsafePtrToInt(unsafe.Pointer(btn.C())))
		if draw.win == nil {
			draw.win = btn.C()
		}
		btn.CopyDrawData(draw)
		return btn
	}
	return nil
}

func newRadioButton(parent *Window, status gui.StatusFlags, px, py, w, h int, draw *WindowData, data *radioButtonData) *Window {
	if !draw.StyleFlags().IsRadioButton() {
		return nil
	}
	win := newWindowRaw(parent, status, px, py, w, h, C.nox_xxx_wndRadioButtonProcPre_4A93C0)
	if win == nil {
		return nil
	}
	C.nox_xxx_wndRadioButtonSetAllFn_4A87E0(unsafePtrToInt(unsafe.Pointer(win.C())))
	if draw.win == nil {
		draw.win = win.C()
	}
	dp, _ := alloc.Calloc(1, unsafe.Sizeof(C.nox_radioButton_data{}))
	d := (*C.nox_radioButton_data)(dp)
	if data != nil {
		d.field_0 = data.field_0
	}
	win.widget_data = unsafe.Pointer(d)
	win.CopyDrawData(draw)
	return win
}
