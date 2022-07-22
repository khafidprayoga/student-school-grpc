package controller

import (
	"context"

	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type StudentServer struct {
	pb.UnimplementedStudentServiceServer
	DB *gorm.DB
}

func NewStudentServer(db *gorm.DB) StudentServer {
	return StudentServer{
		DB: db,
	}
}

func (s *StudentServer) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (s *StudentServer) GetStudentById(ctx context.Context, req *pb.GetStudentByIdRequest) (*pb.Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentById not implemented")
}
func (s *StudentServer) GetAllStudent(ctx context.Context, req *emptypb.Empty) (*pb.ListStudent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStudent not implemented")
}
func (s *StudentServer) UpdateStudentAddress(ctx context.Context, req *pb.UpdateStudentAddressRequest) (*pb.Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudentAddress not implemented")
}
func (s *StudentServer) UpdateStudentPhone(ctx context.Context, req *pb.UpdateStudentPhoneRequest) (*pb.Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudentPhone not implemented")
}
