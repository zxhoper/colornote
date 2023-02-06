package colornote

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ==============================================================================================
var MyDebug = false

var testPrefix = "    ---   | "
var notePrefix = ""
var denotePrefix = "=-=-=>"

const (
	clrNon    = "0"
	clrRed    = "31"
	clrGreen  = "32"
	clrYellow = "33"
	clrBlue   = "34"
	clrPurple = "35"
	clrCyan   = "36"
	clrGray   = "37"
	clrWhite  = "97"
)

func prepareColorCode(c string) string {
	return fmt.Sprintf("\033[%sm", c)
}

// eeP  eePt wraps fmt.Println with prefix
func NoteColor(c string, a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stdout, c)
	eewP(os.Stdout, notePrefix, a...)
	_, _ = fmt.Fprintln(os.Stdout, "\033[0m")
}

func NoteGreen(a ...interface{}) {
	NoteColor(prepareColorCode(clrGreen), a...)
}
func NoteRed(a ...interface{}) {
	NoteColor(prepareColorCode(clrRed), a...)
}
func NoteYellow(a ...interface{}) {
	NoteColor(prepareColorCode(clrYellow), a...)
}
func NoteBlue(a ...interface{}) {
	NoteColor(prepareColorCode(clrBlue), a...)
}
func NotePurple(a ...interface{}) {
	NoteColor(prepareColorCode(clrPurple), a...)
}
func NoteCyan(a ...interface{}) {
	NoteColor(prepareColorCode(clrCyan), a...)
}
func NoteGray(a ...interface{}) {
	NoteColor(prepareColorCode(clrGray), a...)
}

func Note(a ...interface{}) {
	eewP(os.Stdout, notePrefix, a...)
}
func deNote(a ...interface{}) {
	if MyDebug {
		eewP(os.Stdout, denotePrefix, a...)
	}
}

// eewP  wraps fmt.Println with prefix
func eewP(w io.Writer, pfix string, a ...interface{}) {
	var prefixedAny []interface{}
	prefixedAny = append(prefixedAny, pfix)
	prefixedAny = append(prefixedAny, a...)
	_, _ = fmt.Fprintln(w, prefixedAny...)
}

// eePt wraps fmt.Println with prefix and print banner style title
func Notet(s string) {
	eewPt(os.Stdout, notePrefix, s)
}
func deNotet(s string) {
	if MyDebug {
		eewPt(os.Stdout, denotePrefix, s)
	}
}
func eewPt(w io.Writer, pfix string, s string) {
	_, _ = fmt.Fprintln(w, pfix+eeStrRepeat("=", 51))
	_, _ = fmt.Fprintf(w, pfix+"== %-45s ==\n", s)
	_, _ = fmt.Fprintln(w, pfix+eeStrRepeat("=", 51))
}

// eePf wraps fmt.Printf with prefix
func Notef(format string, a ...interface{}) {
	eewPf(os.Stdout, notePrefix, format, a...)
}
func NotefColor(c string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s", c, format, prepareColorCode(clrNon))
	eewPf(os.Stdout, notePrefix, formatColor, a...)
}

func NotefGreen(format string, a ...interface{}) {
	NotefColor(prepareColorCode(clrGreen), format, a...)
}
func NotefRed(format string, a ...interface{}) {
	NotefColor(prepareColorCode(clrRed), format, a...)
}
func NotefYellow(format string, a ...interface{}) {
	NotefColor(prepareColorCode(clrYellow), format, a...)
}
func NotefBlue(format string, a ...interface{}) {
	NotefColor(prepareColorCode(clrBlue), format, a...)
}
func NotefPurple(format string, a ...interface{}) {
	NotefColor(prepareColorCode(clrPurple), format, a...)
}
func NotefCyan(format string, a ...interface{}) {
	NotefColor(prepareColorCode(clrCyan), format, a...)
}
func NotefGray(format string, a ...interface{}) {
	NotefColor(prepareColorCode(clrGray), format, a...)
}

func deNotef(format string, a ...interface{}) {
	if MyDebug {
		eewPf(os.Stdout, denotePrefix, format, a...)
	}
}
func eewPf(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + format
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)

}

func Notefn(format string, a ...interface{}) {
	eewPfn(os.Stdout, format, notePrefix, a...)
}
func deNotefn(format string, a ...interface{}) {
	eewPfn(os.Stdout, format, denotePrefix, a...)
}
func eewPfn(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + format + "\n"
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
}

func NoteftColor(c string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s", c, format, prepareColorCode(clrNon))
	eewPft(os.Stdout, formatColor, notePrefix, a...)
}

func NoteftRed(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrRed), format, a...)
}
func NoteftGreen(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrGreen), format, a...)
}
func NoteftYellow(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrYellow), format, a...)
}
func NoteftBlue(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrBlue), format, a...)
}
func NoteftPurple(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrPurple), format, a...)
}
func NoteftCyan(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrCyan), format, a...)
}
func NoteftGray(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrGray), format, a...)
}
func NoteftWhite(format string, a ...interface{}) {
	NoteftColor(prepareColorCode(clrWhite), format, a...)
}

func Noteft(format string, a ...interface{}) {
	eewPft(os.Stdout, format, denotePrefix, a...)
}
func eewPft(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + format
	_, _ = fmt.Fprintln(w, prefixedFormat+eeStrRepeat("=", 51))
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
	_, _ = fmt.Fprintln(w, prefixedFormat+eeStrRepeat("=", 51))
}

func NoteHr() {
	eewPhr(os.Stdout, notePrefix)
}
func deNoteHr() {
	eewPhr(os.Stdout, denotePrefix)
}
func eewPhr(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+eeStrRepeat("=", 51))
}

func NoteFirst() {
	eewPFirst(os.Stdout, notePrefix)
}
func deNoteFirst() {
	eewPFirst(os.Stdout, denotePrefix)
}
func eewPFirst(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+eeStrRepeat("-", 64)+"\\\\")
}

func NoteLast() {
	eewPLast(os.Stdout, notePrefix)
}
func deNoteLast() {
	eewPLast(os.Stdout, denotePrefix)
}
func eewPLast(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+eeStrRepeat("-", 64)+"//")
}

func eeStrRepeat(s string, repeat int) string {
	var sb strings.Builder
	for i := 0; i < repeat; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

// ===============================================================================================
