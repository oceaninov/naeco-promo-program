package _interface

import (
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
)

type Service interface {
	pb.ProgramServiceServer
}
