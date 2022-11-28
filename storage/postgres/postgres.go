package postgres

import (
	"database/sql"
	"fmt"
	pb "github.com/Abdurahmonjon/studentcrud/studentservice/genproto/gitlab.com/Abdurahmonjon/studentproto"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r PostgresRepo) CreateStudent(sr pb.RegisterStudentResponse) (pb.RegisterStudentResponse, error) {
	var id string
	err := r.db.QueryRow(`insert into students (id , first_name,last_name,user_name) values ($1,$2,$3,$4) returning id`, sr.Id, sr.FirstName, sr.LastName, sr.UserName).Scan(&id)
	if err != nil {
		return pb.RegisterStudentResponse{}, err
	}
	student, err := r.GetStudent(id)
	if err != nil {
		return pb.RegisterStudentResponse{}, err
	}
	return student, nil
}

func (r PostgresRepo) GetStudent(id string) (pb.RegisterStudentResponse, error) {
	var student pb.RegisterStudentResponse
	err := r.db.QueryRow(`select * from students where id = $1`, id).Scan(
		&student.Id, &student.FirstName, &student.LastName, &student.UserName)
	if err != nil {
		return pb.RegisterStudentResponse{}, err
	}
	return student, nil
}

func (r PostgresRepo) GetAllStudents(page, limit int32) ([]*pb.RegisterStudentResponse, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(`
				SELECT id, first_name, last_name, user_name
				FROM students LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	//students := make([]pb.RegisterStudentResponse, 0)
	var students []*pb.RegisterStudentResponse
	for rows.Next() {
		var student *pb.RegisterStudentResponse
		err = rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.UserName)
		students = append(students, student)
	}

	return students, nil

}

func (r PostgresRepo) DeleteStudent(sr pb.DeleteStudentRequest) (pb.Response, error) {
	result, err := r.db.Exec(`delete from students where id = $1`, &sr.Id)
	if err != nil {
		return pb.Response{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Response{}, sql.ErrNoRows
	}
	return pb.Response{}, nil
}

func (r PostgresRepo) UpdateStudent(sr pb.UpdateStudentRequest) (pb.Response, error) {
	result, err := r.db.Exec(`update students set first_name=$1,last_name=$2,user_name=$3 where id = $4`,
		&sr.FirstName, &sr.LastName, &sr.UserName, &sr.Id)
	if err != nil {
		fmt.Println("user id:", sr.Id, sr.UserName, sr.FirstName)
		return pb.Response{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Response{}, sql.ErrNoRows
	}
	return pb.Response{}, nil
}
