package modes

import "github.com/bennicholls/delvetown/ui"

type Log struct {
	width, height int
	logs          []string
	logbox        *ui.Textbox
}

func NewLog(logbox *ui.Textbox) Log {
	w, h := logbox.Dims()
	return Log{w, h, make([]string, 0, 50), logbox}
}

//Adds a log entry, then updates the ui element. Currently just adds spaces to the end of each line and
//throws a big long ugly string into the box that simulates word wrap. Replace this someday with a
//solution that isn't embarassing why don't you.
func (l *Log) AddMessage(message string) {

	l.logs = append(l.logs, message)

	//build string for ui and ship out
	text := ""

	for i := 0; i < l.height; i++ {
		if i == len(l.logs) {
			break
		}
		text += l.logs[len(l.logs)-1-i]
		for n := len(l.logs[len(l.logs)-1-i]); n < l.width; n++ {
			text += " "
		}
	}

	l.logbox.ChangeText(text)
}
