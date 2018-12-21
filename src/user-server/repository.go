package user_server

import (
	pb "./proto/user"
	"database/sql"
	"log"
)

type Repository interface {
	Get(id string) (*pb.User,error)
	GetAll()([]*pb.User,error)
	Create(*pb.User,error)
	GetByEmail(email string)(*pb.User,error)
}

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) Get(id string) (*pb.User,error){
	user:=&pb.User{}
	db:=ur.db
	err:=db.QueryRow("select * from t_user where  id = ? ",id).Scan(&user.Id,&user.Name,&user.Password,&user.Company,&user.Email)
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	return user,nil
}

func (ur *UserRepository)GetAll()([]*pb.User,error){
	db:=ur.db
	users:=[]*pb.User{}
	rows,err:=db.Query("select * from t_user ")
	for rows.Next(){
		user:=pb.User{}
		err:=rows.Scan(&user.Id,&user.Name,&user.Password,&user.Company,&user.Email)
		if err != nil {
			log.Fatalln(err)
			return nil,err
		}
		users=append(users,&user)
	}
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	return users,nil
}

func (ur *UserRepository)Create(user *pb.User)error{
	db:=ur.db
	_,err:=db.Exec("insert into t_user values(?,?,?,?,?)",user.Id,user.Name,user.Password,user.Company,user.Email)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (ur *UserRepository)GetByEmail(email string)(*pb.User,error){
	user:=&pb.User{}
	db:=ur.db
	err:=db.QueryRow("select * from t_user where email = ?",email).Scan(&user.Id,&user.Name,&user.Password,&user.Company,&user.Email)
	if err!=nil{
		log.Fatalln(err)
		return nil,err
	}
	return user,nil
}
