package service

import (
	"context"
	"fmt"
	"github.com/Abdurahmonjon/studentcrud/studentservice/genproto/gitlab.com/Abdurahmonjon/studentproto"
	pb "github.com/Abdurahmonjon/studentcrud/studentservice/genproto/gitlab.com/Abdurahmonjon/studentproto"
	"github.com/Abdurahmonjon/studentcrud/studentservice/pkg/logger"
	"github.com/Abdurahmonjon/studentcrud/studentservice/storage"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StudentService struct {
	storage storage.IStorage
	logger  logger.Logger
	studentproto.UnimplementedStudentServiceServer
}

func NewService(iStorage storage.IStorage, log logger.Logger) *StudentService {
	return &StudentService{
		storage: iStorage,
		logger:  log,
	}
}

func (s *StudentService) RegisterStudent(ctx context.Context, req *pb.RegisterStudentRequest) (*pb.RegisterStudentResponse, error) {
	id := uuid.New()
	var resp pb.RegisterStudentResponse
	resp.Id = id.String()
	fmt.Println("coming infos:", req.UserName, req.FirstName, req.LastName)
	resp.FirstName = req.FirstName
	resp.LastName = req.LastName
	resp.UserName = req.UserName
	student, err := s.storage.Student().CreateStudent(resp)
	if err != nil {
		s.logger.Error("failed to create student", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to create student")
	}
	return &student, nil
}

func (s *StudentService) GetStudent(ctx context.Context, request *pb.GetStudentRequest) (*pb.RegisterStudentResponse, error) {
	student, err := s.storage.Student().GetStudent(request.Id)
	if err != nil {
		s.logger.Error("failed to get student", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get student")
	}
	return &student, nil
}

func (s *StudentService) GetAllStudents(ctx context.Context, request *pb.GetAllStudentsRequest) (*pb.GetAllStudentsResponse, error) {
	students, err := s.storage.Student().GetAllStudents(request.GetPage(), request.GetLimit())
	if err != nil {
		s.logger.Error("failed to get all students", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get list of students")
	}

	return &pb.GetAllStudentsResponse{Students: students}, nil
}

func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.Response, error) {
	response, err := s.storage.Student().UpdateStudent(*req)
	if err != nil {
		s.logger.Error("failed to update student", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to update student")
	}
	return &response, nil
}

func (s *StudentService) DeleteStudent(ctx context.Context, request *pb.DeleteStudentRequest) (*emptypb.Empty, error) {
	_, err := s.storage.Student().DeleteStudent(*request)
	if err != nil {
		s.logger.Error("failed to delete student", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete student")
	}
	return &emptypb.Empty{}, nil
}
