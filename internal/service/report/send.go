package report

import (
	"base/internal/data/enums"
	"base/internal/handler/filter"
	"base/internal/models"
	"context"
	"fmt"
	"time"
)

func (s *reportService) GenerateMorningReport(ctx context.Context) (string, error) {
	input := filter.TaskListParams{}
	var res []models.Task
	var err = s.taskRepo.R_Search(ctx, &input, &res)
	if err != nil {
		return "", err
	}
	grouped := make(map[string][]models.Task)
	for _, task := range res {
		grouped[task.Status] = append(grouped[task.Status], task)
	}
	now := time.Now().Format("02-01-2006")

	// Thay bằng dữ liệu thực từ DB
	done := grouped[enums.TaskStatusTypeCompleted.String()]
	inProgress := grouped[enums.TaskStatusTypeInProgress.String()]
	notStarted := grouped[enums.TaskStatusTypePending.String()]

	msg := fmt.Sprintf("*Báo cáo công việc %s*", now)
	//done
	msg += "\n\nTổng quan:"
	msg += "\n- Tổng số công việc: " + fmt.Sprint(len(res))
	msg += fmt.Sprintf("\n- Đã xong: %d", len(done))
	msg += fmt.Sprintf("\n | title | description | assigned to | status | deadline | create by ")
	for i := 0; i < len(done); i++ {
		msg += fmt.Sprintf("\n | %s | %s | %s | %s | %s | %s ", done[i].Title, done[i].Description, done[i].AssignedToEmail, enums.StringToTaskStatusType(done[i].Status).Val(), done[i].Deadline.Format("02-01-2006"), done[i].CreatedByEmail)
	}
	msg += "\n- Đang làm: " + fmt.Sprint(len(inProgress))
	msg += fmt.Sprintf("\n | title | description | assigned to | status | deadline | create by ")
	for i := 0; i < len(inProgress); i++ {
		msg += fmt.Sprintf("\n | %s | %s | %s | %s | %s | %s ", inProgress[i].Title, inProgress[i].Description, inProgress[i].AssignedToEmail, enums.StringToTaskStatusType(inProgress[i].Status).Val(), inProgress[i].Deadline.Format("02-01-2006"), inProgress[i].CreatedByEmail)
	}
	msg += "\n- Chưa làm: " + fmt.Sprint(len(notStarted))
	msg += fmt.Sprintf("\n | title | description | assigned to | status | deadline | create by ")
	for i := 0; i < len(notStarted); i++ {
		msg += fmt.Sprintf("\n | %s | %s | %s | %s | %s | %s ", notStarted[i].Title, notStarted[i].Description, notStarted[i].AssignedToEmail, enums.StringToTaskStatusType(notStarted[i].Status).Val(), notStarted[i].Deadline.Format("02-01-2006"), notStarted[i].CreatedByEmail)
	}

	msg += "\n\n*Lưu ý: Đây là báo cáo tự động, vui lòng không trả lời tin nhắn này.*"
	msg += "\n\n*Chúc bạn một ngày làm việc hiệu quả!*"
	msg += "\n\n*Để biết thêm thông tin chi tiết, vui lòng truy cập vào hệ thống quản lý công việc.*"
	msg += "\n\n*Trân trọng,*"
	return msg, nil
}
