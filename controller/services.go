package controller

import (
	"context"

	"github.com/khafidprayoga/grpc-basic/proto/pb"
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

func (s *StudentServer) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreatedStudentResponse, error) {
	return AddStudent(s.DB, req)
}
func (s *StudentServer) GetStudentById(ctx context.Context, req *pb.GetStudentByIdRequest) (*pb.StudentDetailResponse, error) {
	return GetStudentById(s.DB, req)
}
func (s *StudentServer) GetAllStudent(ctx context.Context, req *emptypb.Empty) (*pb.ListStudent, error) {
	return GetAllStudent(s.DB, req)
}
func (s *StudentServer) UpdateStudentAddress(ctx context.Context, req *pb.UpdateStudentAddressRequest) (*pb.GlobalResponse, error) {
	return UpdateStudentAddress(s.DB, req)
}

func (s *StudentServer) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.GlobalResponse, error) {
	return DeleteStudent(s.DB, req)
}
