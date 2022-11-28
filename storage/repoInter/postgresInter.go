package repoInter

import (
	pb "github.com/Abdurahmonjon/studentcrud/studentservice/genproto/gitlab.com/Abdurahmonjon/studentproto"
)

type TaskStorageI interface {
	CreateStudent(request pb.RegisterStudentResponse) (pb.RegisterStudentResponse, error)
	GetStudent(request string) (pb.RegisterStudentResponse, error)
	DeleteStudent(request pb.DeleteStudentRequest) (pb.Response, error)
	UpdateStudent(request pb.UpdateStudentRequest) (pb.Response, error)
	GetAllStudents(page, limit int32) ([]*pb.RegisterStudentResponse, error)
}
