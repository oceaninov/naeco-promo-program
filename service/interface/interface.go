package _interface

import (
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
)

type Service interface {
	pb.ProgramServiceServer
}
