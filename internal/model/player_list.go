package model

type PlayerList struct {
	ID                int    `json:"id"`
	UserId            int    `json:"user_id"`
	ScheduleId        int    `json:"schedule_id"`
	Attendance_Status string `json:"attendance_status"`
	Payment_Status    string `json:"payment_status"`
	Payment_Amount    int    `json:"payment_amount"`
}
