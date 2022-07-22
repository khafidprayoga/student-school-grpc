package controller

import (
	"errors"
	"time"

	"github.com/khafidprayoga/grpc-basic/common/entities"
	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"gorm.io/gorm"
)

func AddStudent(db *gorm.DB, req *pb.CreateStudentRequest) (*pb.CreatedStudentResponse, error) {
	student := entities.Student{
		FullName:  req.GetFullName(),
		BirthDate: time.Unix(int64(req.GetBirthDate()), 0),
		Class:     req.GetClass(),
		Address:   req.GetAddress(),
	}
	if err := db.Create(&student).Error; err != nil {
		return nil, errors.New("Gagal menambahkan murid baru")
	}

	result := pb.CreatedStudentResponse{
		Id: int64(student.Id),
	}
	return &result, nil
}
