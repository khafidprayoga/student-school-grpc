package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/khafidprayoga/grpc-basic/common/entities"
	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
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

func GetStudentById(db *gorm.DB, req *pb.GetStudentByIdRequest) (*pb.StudentDetailResponse, error) {
	var student entities.Student
	res := db.First(&student).Where("id = ?", req.Id)

	if res.Error != nil {
		return nil, errors.New("Failed get student details")
	}
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
	result := db.Find(&students)
	if result.Error != nil {
		return nil, errors.New("Failed fetch all student data")
	}

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

func UpdateStudentAddress(db *gorm.DB, req *pb.UpdateStudentAddressRequest) (*pb.GlobalResponse, error) {
	var student entities.Student
	studentId := req.GetId()
	newAddress := req.GetNewAddress()

	result := db.Model(&student).Where("id = ?", studentId).Update("address", newAddress)

	if result.Error != nil {
		return nil, errors.New("Failed update student address")
	}
	return &pb.GlobalResponse{
		Code:    http.StatusOK,
		Message: "Address successfully updated",
	}, nil
}

func DeleteStudent(db *gorm.DB, req *pb.DeleteStudentRequest) (*pb.GlobalResponse, error) {
	var student entities.Student
	res := db.Delete(&student, req.GetId())
	if res.Error != nil {
		return nil, errors.New("Failed to delete student")
	}

	return &pb.GlobalResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Student deleted with Id  %v", req.GetId()),
	}, nil
}
