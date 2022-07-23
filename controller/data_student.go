package controller

import (
	"github.com/khafidprayoga/grpc-basic/common/entities"
	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
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

func GetAllStudent(db *gorm.DB, req *emptypb.Empty) (*pb.ListStudent, error) {
	var students []entities.Student
	var studentsPb []*pb.Student
	db.Find(&students)

	for _, val := range students {
		student := &pb.Student{
			Id:        int32(val.Id),
			FullName:  val.FullName,
			BirthDate: val.BirthDate.Unix(),
			Class:     val.Class,
			Address:   val.Address,
		}
		studentsPb = append(studentsPb, student)
	}

	return &pb.ListStudent{
		Students: studentsPb,
	}, nil
}
