package controller

import (
	"github.com/khafidprayoga/grpc-basic/common/entities"
	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"gorm.io/gorm"
)

func GetStudentById(db *gorm.DB, req *pb.GetStudentByIdRequest) (*pb.StudentDetailResponse, error) {
	var student entities.Student
	db.First(&student).Where("id = ?", req.Id)
	var studentData = pb.StudentDetailResponse{
		Data: &pb.Student{
			Id:        int32(student.Id),
			FullName:  student.FullName,
			BirthDate: student.BirthDate.Unix() * 1000,
			Class:     student.Class,
			Address:   student.Address,
		},
	}
	return &studentData, nil
}
