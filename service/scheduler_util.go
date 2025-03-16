package service

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	clockwerk "github.com/nightsilvertech/clockwerk/client"
	programconst "github.com/oceaninov/naeco-promo-program/constants"
	"github.com/oceaninov/naeco-promo-program/gvars"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	"github.com/oceaninov/naeco-promo-util/er"
	"github.com/oceaninov/naeco-promo-util/lgr"
	"google.golang.org/grpc/codes"
	"time"
)

func (s service) refreshDailyQuotaScheduler(req *pb.AddProgramReq, programID string) error {
	systemUsername := s.vault.Get("/basic_auth_system:username")
	systemPassword := s.vault.Get("/basic_auth_system:password")

	if req.RefreshProgramQuotaDaily != 0 {
		_, err := s.repo.Scheduler.Add(clockwerk.SchedulerHTTP{
			Name:           fmt.Sprintf("Refresh daily quota whitelist program id %s", programID),
			URL:            fmt.Sprintf(programconst.RefreshDailyQuotaURL+"%s", programID),
			ReferenceId:    programID,
			Executor:       clockwerk.HTTP,
			Method:         clockwerk.PUT,
			Body:           "",
			Spec:           fmt.Sprintf("0 0 * * *"),
			Disabled:       false,
			Persist:        true,
			Retry:          15,
			RetryThreshold: 1,
			HTTPHeader: []clockwerk.HTTPHeader{
				{K: "Content-Type", V: "application/json"},
				{K: "Authorization", V: fmt.Sprintf("Basic %s:%s", systemUsername, systemPassword)},
			},
		})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return er.Ebl(codes.Internal, "failed to create scheduler for refresh daily quota", err)
		}
	}
	return nil
}

func (s service) statusUpdateScheduler(req *pb.AddProgramReq, programID string) error {
	systemUsername := s.vault.Get("/basic_auth_system:username")
	systemPassword := s.vault.Get("/basic_auth_system:password")

	if req.Status == 0 {
		// start at update program to active
		startAtTime := time.Unix(req.StartAt/1000, 0)
		_, err := s.repo.Scheduler.Add(clockwerk.SchedulerHTTP{
			Name:           fmt.Sprintf("Update program with id %s status to active", programID),
			URL:            fmt.Sprintf(programconst.UpdateProgramStatusURL+"%s/%d", programID, programconst.ProgramActive),
			ReferenceId:    programID,
			Executor:       clockwerk.HTTP,
			Method:         clockwerk.PUT,
			Body:           "",
			Spec:           fmt.Sprintf("0 0 %d %d *", startAtTime.Day(), int(startAtTime.Month())),
			Disabled:       false,
			Persist:        false,
			Retry:          15,
			RetryThreshold: 1,
			HTTPHeader: []clockwerk.HTTPHeader{
				{K: "Content-Type", V: "application/json"},
				{K: "Authorization", V: fmt.Sprintf("Basic %s:%s", systemUsername, systemPassword)},
			},
		})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return er.Ebl(codes.Internal, "failed to create scheduler for activate program", err)
		}

		// end at update program to terminate
		endAtTime := time.Unix(req.EndAt/1000, 0).AddDate(0, 0, 1)
		_, err = s.repo.Scheduler.Add(clockwerk.SchedulerHTTP{
			Name:           fmt.Sprintf("Update program with id %s status to termintate", programID),
			URL:            fmt.Sprintf(programconst.UpdateProgramStatusURL+"%s/%d", programID, programconst.ProgramTerminate),
			ReferenceId:    programID,
			Executor:       clockwerk.HTTP,
			Method:         clockwerk.PUT,
			Body:           "",
			Spec:           fmt.Sprintf("0 0 %d %d *", endAtTime.Day(), int(endAtTime.Month())),
			Disabled:       false,
			Persist:        false,
			Retry:          15,
			RetryThreshold: 1,
			HTTPHeader: []clockwerk.HTTPHeader{
				{K: "Content-Type", V: "application/json"},
				{K: "Authorization", V: fmt.Sprintf("Basic %s:%s", systemUsername, systemPassword)},
			},
		})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return er.Ebl(codes.Internal, "failed to create scheduler for terminate program", err)
		}
	} else {
		// end at update program to terminate
		endAtTime := time.Unix(req.EndAt/1000, 0).AddDate(0, 0, 1)
		_, err := s.repo.Scheduler.Add(clockwerk.SchedulerHTTP{
			Name:           fmt.Sprintf("Update program with id %s status to termintate", programID),
			URL:            fmt.Sprintf(programconst.UpdateProgramStatusURL+"%s/%d", programID, programconst.ProgramTerminate),
			ReferenceId:    programID,
			Executor:       clockwerk.HTTP,
			Method:         clockwerk.PUT,
			Body:           "",
			Spec:           fmt.Sprintf("0 0 %d %d *", endAtTime.Day(), int(endAtTime.Month())),
			Disabled:       false,
			Persist:        false,
			Retry:          15,
			RetryThreshold: 1,
			HTTPHeader: []clockwerk.HTTPHeader{
				{K: "Content-Type", V: "application/json"},
				{K: "Authorization", V: fmt.Sprintf("Basic %s:%s", systemUsername, systemPassword)},
			},
		})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return er.Ebl(codes.Internal, "failed to create scheduler for terminate program", err)
		}
	}
	return nil
}
