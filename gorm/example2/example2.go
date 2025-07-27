package example2

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	//_ "github.com/go-sql-driver/mysql" // 替换为你的数据库驱动
	"log"
)

// Employee 员工结构体（字段名与数据库列名通过 `db` 标签映射）
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"` // 使用float64存储工资（假设数据库是DECIMAL/FLOAT类型）
}

// GetEmployeesByDepartment 查询指定部门的所有员工
func GetEmployeesByDepartment(db *sqlx.DB, dept string) ([]Employee, error) {
	var employees []Employee
	query := "SELECT id, name, department, salary FROM employees WHERE department = ?"
	err := db.Select(&employees, query, dept)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return employees, nil
}

// GetTopSalaryEmployee 查询工资最高的员工
func GetTopSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	// 方法1：使用ORDER BY + LIMIT（推荐）
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)
	if err != nil {
		return Employee{}, fmt.Errorf("查询失败: %v", err)
	}
	return employee, nil
}

// Book 结构体与 books 表字段对应（使用 `db` 标签确保正确映射）
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"` // 使用 float64 对应数据库的 DECIMAL/FLOAT 类型
}

// GetExpensiveBooks 查询价格大于指定值的书籍
func GetExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book

	query := `
		SELECT id, title, author, price 
		FROM books 
		WHERE price > ? 
		ORDER BY price DESC`

	// 执行查询并映射到结构体切片
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}

	return books, nil
}

func Run() {
	// 初始化数据库连接（替换为你的DSN）
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}
	defer db.Close()

	// 1. 查询技术部所有员工
	techEmployees, err := GetEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Println("查询技术部员工失败:", err)
	} else {
		fmt.Println("技术部员工列表:")
		for _, emp := range techEmployees {
			fmt.Printf("ID: %d, 姓名: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Salary)
		}
	}
	// 2. 查询工资最高的员工
	topEarner, err := GetTopSalaryEmployee(db)
	if err != nil {
		log.Println("查询工资最高员工失败:", err)
	} else {
		fmt.Printf("\n工资最高的员工:\nID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
			topEarner.ID, topEarner.Name, topEarner.Department, topEarner.Salary)
	}

	expensiveBooks, err := GetExpensiveBooks(db, 50.0)
	if err != nil {
		log.Fatalln("查询失败:", err)
	}

	// 打印结果
	fmt.Println("价格 > 50 元的书籍:")
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n",
			book.ID, book.Title, book.Author, book.Price)
	}
}
