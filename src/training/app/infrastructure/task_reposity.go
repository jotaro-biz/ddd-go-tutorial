package infrastructure

import (
	"training/app/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

var (
	DBMS     = "mysql"
	USER     = "root"
	PASS     = "password"
	PROTOCOL = "tcp(ddd-training-mysql:3306)"
	DBNAME   = "gopher"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
)

//DB初期化
func (t *TaskRepository) DbInit() {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！（dbInit）")
	}
	db.AutoMigrate(&domain.Task{})
	defer db.Close()
}

//DB追加
func (t *TaskRepository) Create(text string, status string) {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&domain.Task{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func (t *TaskRepository) Update(id int, text string, status string) {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！（dbUpdate)")
	}
	var todo domain.Task
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func (t *TaskRepository) Delete(id int) {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！（dbDelete)")
	}
	var todo domain.Task
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func (t *TaskRepository) GetAll() []domain.Task {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var tasks []domain.Task
	db.Order("created_at desc").Find(&tasks)
	db.Close()
	return tasks
}

//DB一つ取得
func (t *TaskRepository) GetById(id int) domain.Task {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var todo domain.Task
	db.First(&todo, id)
	db.Close()
	return todo
}
