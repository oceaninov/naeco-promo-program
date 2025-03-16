package service

import (
	"context"
	"github.com/go-kit/kit/log/level"
	"github.com/oceaninov/naeco-promo-program/gvars"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	pbwhitelist "github.com/oceaninov/naeco-promo-whitelist/protocs/api/v1"
	"github.com/oceaninov/naeco-promo-util/er"
	"github.com/oceaninov/naeco-promo-util/jwt"
	"github.com/oceaninov/naeco-promo-util/lgr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func CreateContext(ctx context.Context, token string) context.Context {
	md := metadata.MD{}
	md.Set("authorization", "Bearer "+token)
	return metadata.NewIncomingContext(ctx, md)
}

func (s service) bulkInsertCustCSV(ctx context.Context, req *pb.AddProgramReq, programID string) error {
	jwtToken, err := jwt.Bearer(ctx, s.vault.Get("/jwt:secret"))
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return er.Ebl(codes.Internal, "failed to read jwt", err)
	}

	_, err = s.repo.Microservice.WhitelistSvc.AddCustomerWhitelistBulk(CreateContext(ctx, jwtToken.Bearer), &pbwhitelist.AddCustomerWhitelistBulkReq{
		ProgramId: programID,
		CreatedBy: req.CreatedBy,
		Url:       req.CustomerCsvUrl,
	})
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return err
	}
	return nil
}

func (s service) bulkInsertMerchCSV(ctx context.Context, req *pb.AddProgramReq, programID string) error {
	jwtToken, err := jwt.Bearer(ctx, s.vault.Get("/jwt:secret"))
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return er.Ebl(codes.Internal, "failed to read jwt", err)
	}

	_, err = s.repo.Microservice.WhitelistSvc.AddMerchantWhitelistBulk(CreateContext(ctx, jwtToken.Bearer), &pbwhitelist.AddMerchantWhitelistBulkReq{
		ProgramId: programID,
		CreatedBy: req.CreatedBy,
		Url:       req.MerchantCsvUrl,
	})
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return err
	}
	return nil
}
