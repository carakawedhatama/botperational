package on_leave

import (
	"fmt"
	"time"
)

const (
	FOOTER_TEXT = "Mohon dapat menghubungi lead / head department terkait biar kamu tetep bisa koordinasi yah! 😊"
	DOC_TYPE    = "on_leave"
)

func ContentMsg() string {
	tm := time.Now()
	currDt := tm.Format("02 Jan 2006")
	return fmt.Sprintf("📣 Yang ngajukan cuti (%s) 🏖️", currDt)
}

func NoLeaveData() string {
	tm := time.Now()
	currDt := tm.Format("02 Jan 2006")
	return fmt.Sprintf("🎉 Ga ada yang ngajukan cuti (%s) 🚀", currDt)
}
