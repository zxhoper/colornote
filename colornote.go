package colornote

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func PrintAllColor() {
	NoteBlockFirst()

	Note("Default Note")
	Notef("Default Notef\n")

	NoteGreen("Green")
	NoteRed("Red")
	NoteYellow("Yellow")
	NoteBlue("Blue")
	NotePurple("Purple")
	NoteCyan("Cyan")
	NoteGray("Gray")
	NoteWhite("White")
	NoteBlockLast()

	NoteHr()
}
func PrintAllColorNotef() {
	NoteBlockFirst()

	Notef("Default Notef\n")

	NoteGreenf("Green")
	NoteRedf("Red")
	NoteYellowf("Yellow")
	NoteBluef("Blue")
	NotePurplef("Purple")
	NoteCyanf("Cyan")
	NoteGrayf("Gray")
	NoteWhitef("White\n")

	NoteWhitef("Print AUTO COLOR\n")
	for i := 0; i < GetColorType(); i++ {

		NoteAutof("Print auto color number: %d  ", i)
	}

	NoteBlockLast()

	NoteHr()
}
func PrintAllColorTitle() {
	NoteBlockFirst()
	NoteT("Default NoteT")
	NoteTf("Default Note Title \n")
	NoteTGreenf("Green\n")
	NoteTRedf("Red\n")
	NoteTYellowf("Yellow\n")
	NoteTBluef("Blue\n")
	NoteTPurplef("Purple\n")
	NoteTCyanf("Cyan\n")
	NoteTGrayf("Gray\n")
	NoteTWhitef("White\n")
	NoteBlockLast()
}

func PrintAllDeNote() {
	DeNoteBlockFirst()

	DeNote("Note for debug")
	DeNote("with debug prefix")
	DeNoteHr()
	DeNote("Default DeNore")

	DeNoteTf("Default DeNotef\n")
	DeNoteTGreenf("Green\n")
	DeNoteTRedf("Red\n")
	DeNoteTYellowf("Yellow\n")
	DeNoteTBluef("Blue\n")
	DeNoteTPurplef("Purple\n")
	DeNoteTCyanf("Cyan\n")
	DeNoteTGrayf("Gray\n")
	DeNoteTWhitef("White\n")

	DeNoteBlockLast()
}

func PrintAutoColorNotef() {
	NoteBlockFirst()

	NoteWhitef("AUTO COLOR\n")

	idx := 0
	for {
		idx++
		for i := 0; i < GetColorType(); i++ {
			NoteAutof("Color %d ", i)
		}
		fmt.Printf("\n")
		if idx > 10 {
			break
		}
	}

	NoteBlockLast()

	NoteHr()
}

// ==============================================================================================
var MyDebug = false

var TestPrefix = "    ---   | "
var NotePrefix = ""
var DeNotePrefix = "=-=-=>"

//
//const (
//	clrNon = "0"
//
//	clrRed    = "31"
//	clrGreen  = "32"
//	clrYellow = "33"
//	clrBlue   = "34"
//	clrPurple = "35"
//	clrCyan   = "36"
//	clrGray   = "37"
//	clrWhite  = "97"
//)

var clr = map[string]string{
	"Non":    "0",
	"Red":    "31",
	"Green":  "32",
	"Yellow": "33",
	"Blue":   "34",
	"Purple": "35",
	"Cyan":   "36",
	"Gray":   "37",
	"White":  "97",
}

func GetColorType() int {
	return len(clr)
}

var Acc = newAutoColorControl()

func getClrCode(c string) string {
	return fmt.Sprintf("\033[%sm", c)
}

func NoteColor(c string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, c)
	eewP(os.Stdout, NotePrefix, a...)
	_, _ = fmt.Fprintf(os.Stdout, "\033[0m")
}

func NoteGreen(a ...interface{}) {
	NoteColor(getClrCode(clr["Green"]), a...)
}
func NoteRed(a ...interface{}) {
	NoteColor(getClrCode(clr["Red"]), a...)
}
func NoteYellow(a ...interface{}) {
	NoteColor(getClrCode(clr["Yellow"]), a...)
}
func NoteBlue(a ...interface{}) {
	NoteColor(getClrCode(clr["Blue"]), a...)
}
func NotePurple(a ...interface{}) {
	NoteColor(getClrCode(clr["Purple"]), a...)
}
func NoteCyan(a ...interface{}) {
	NoteColor(getClrCode(clr["Cyan"]), a...)
}
func NoteGray(a ...interface{}) {
	NoteColor(getClrCode(clr["Gray"]), a...)
}
func NoteWhite(a ...interface{}) {
	NoteColor(getClrCode(clr["White"]), a...)
}

func Note(a ...interface{}) {
	eewP(os.Stdout, NotePrefix, a...)
}
func DeNote(a ...interface{}) {
	if MyDebug {
		eewP(os.Stdout, DeNotePrefix, a...)
	}
}

// eewP  wraps fmt.Println with prefix
func eewP(w io.Writer, pfix string, a ...interface{}) {
	var prefixedAny []interface{}
	if pfix != "" {
		prefixedAny = append(prefixedAny, pfix)
	}
	prefixedAny = append(prefixedAny, a...)
	_, _ = fmt.Fprintln(w, prefixedAny...)
}

func NoteT(s string) {
	eewPT(os.Stdout, NotePrefix, s)
}
func DeNoteT(s string) {
	if MyDebug {
		eewPT(os.Stdout, DeNotePrefix, s)
	}
}
func eewPT(w io.Writer, pfix string, s string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
	_, _ = fmt.Fprintf(w, pfix+"== %-45s ==\n", s)
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
}

func Notef(format string, a ...interface{}) {
	eewPf(os.Stdout, NotePrefix, format, a...)
}

type AutoCololrCtl struct {
	TotalColor   int
	Colors       []string
	CurrentColor int
}

func newAutoColorControl() *AutoCololrCtl {
	var colors []string
	for _, c := range clr {
		colors = append(colors, c)
	}

	return &AutoCololrCtl{
		TotalColor:   GetColorType(),
		Colors:       colors,
		CurrentColor: 0,
	}
}

func (acc *AutoCololrCtl) popCurColor() string {
	current := acc.CurrentColor
	if acc.CurrentColor == acc.TotalColor-1 {
		acc.CurrentColor = 0
	} else {
		acc.CurrentColor = current + 1
	}
	return acc.Colors[current]
}

func NoteAutof(format string, a ...interface{}) {
	NoteColorf(getClrCode(Acc.popCurColor()), format, a...)
}

func NoteColorf(colorCode string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s", colorCode, format, getClrCode(clr["Non"]))
	eewPf(os.Stdout, NotePrefix, formatColor, a...)
}

func NoteGreenf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Green"]), format, a...)
}
func NoteRedf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Red"]), format, a...)
}
func NoteYellowf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Yellow"]), format, a...)
}
func NoteBluef(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Blue"]), format, a...)
}
func NotePurplef(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Purple"]), format, a...)
}
func NoteCyanf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Cyan"]), format, a...)
}
func NoteGrayf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Gray"]), format, a...)
}
func NoteWhitef(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["White"]), format, a...)
}

func DeNotef(format string, a ...interface{}) {
	if MyDebug {
		eewPf(os.Stdout, DeNotePrefix, format, a...)
	}
}
func eewPf(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + format
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)

}

func Notefn(format string, a ...interface{}) {
	eewPfn(os.Stdout, format, NotePrefix, a...)
}
func DeNotefn(format string, a ...interface{}) {
	if MyDebug {
		eewPfn(os.Stdout, DeNotePrefix, format, a...)
	}
}
func eewPfn(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + format + "\n"
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
}

func NoteTColorf(c string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s%s", c, "== ", format, getClrCode(clr["Non"]))
	prefixedFormatColor := NotePrefix + formatColor
	hrColor := fmt.Sprintf("%s%s%s", c, StringRepeat("=", 51), getClrCode(clr["Non"]))
	_, _ = fmt.Fprintln(os.Stdout, NotePrefix+hrColor)
	_, _ = fmt.Fprintf(os.Stdout, prefixedFormatColor, a...)
	_, _ = fmt.Fprintln(os.Stdout, NotePrefix+hrColor)
}

func NoteTRedf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Red"]), format, a...)
}
func NoteTGreenf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Green"]), format, a...)
}
func NoteTYellowf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Yellow"]), format, a...)
}
func NoteTBluef(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Blue"]), format, a...)
}
func NoteTPurplef(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Purple"]), format, a...)
}
func NoteTCyanf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Cyan"]), format, a...)
}
func NoteTGrayf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Gray"]), format, a...)
}
func NoteTWhitef(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["White"]), format, a...)
}

func NoteTf(format string, a ...interface{}) {
	eewPft(os.Stdout, NotePrefix, format, a...)
}

func DeNoteTf(format string, a ...interface{}) {
	if MyDebug {
		eewPft(os.Stdout, DeNotePrefix, format, a...)
	}
}
func eewPft(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + "== " + format
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
}

func DeNoteTColorf(c string, format string, a ...interface{}) {
	if MyDebug {
		formatColor := fmt.Sprintf("%s%s%s%s", c, "== ", format, getClrCode(clr["Non"]))
		prefixedFormatColor := DeNotePrefix + formatColor
		hrColor := fmt.Sprintf("%s%s%s", c, StringRepeat("=", 51), getClrCode(clr["Non"]))
		_, _ = fmt.Fprintln(os.Stdout, DeNotePrefix+hrColor)
		_, _ = fmt.Fprintf(os.Stdout, prefixedFormatColor, a...)
		_, _ = fmt.Fprintln(os.Stdout, DeNotePrefix+hrColor)
	}
}

func DeNoteTRedf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Red"]), format, a...)
}
func DeNoteTGreenf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Green"]), format, a...)
}
func DeNoteTYellowf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Yellow"]), format, a...)
}
func DeNoteTBluef(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Blue"]), format, a...)
}
func DeNoteTPurplef(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Purple"]), format, a...)
}
func DeNoteTCyanf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Cyan"]), format, a...)
}
func DeNoteTGrayf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Gray"]), format, a...)
}
func DeNoteTWhitef(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["White"]), format, a...)
}

func NoteHr() {
	eewPhr(os.Stdout, NotePrefix)
}
func DeNoteHr() {
	if MyDebug {
		eewPhr(os.Stdout, DeNotePrefix)
	}
}
func eewPhr(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 72))
}

func NoteBlockFirst() {
	eewPFirst(os.Stdout, NotePrefix)
}
func NoteBlockLast() {
	eewPLast(os.Stdout, NotePrefix)
}

func DeNoteBlockFirst() {
	if MyDebug {
		eewPFirst(os.Stdout, DeNotePrefix)
	}
}
func DeNoteBlockLast() {
	if MyDebug {
		eewPLast(os.Stdout, DeNotePrefix)
	}
}

func eewPFirst(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("-", 64)+"\\\\")
}
func eewPLast(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("-", 64)+"//")
}

func StringRepeat(s string, repeatTimes int) string {
	var sb strings.Builder
	for i := 0; i < repeatTimes; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

// ===============================================================================================
