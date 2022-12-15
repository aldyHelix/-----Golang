GOGEN Framework Implement Clean Architecture

Model
	Todo
		Messsage string
		Checked bool

Usecase
 * Membuat todo item
 * Melihat semua todo item
 * Men-check-list todo item

7 langkah membuat aplikasi todo list
	1. membuat domain
		todocore
	2. membuat model
		Todo
	3. membuat usecase
		RunTodoCreate
		GetAllTodo
		RunTodoCheck
	4. membuat repository / service
		akses database via interface
	5. membuat gateway
		implementasi dari interface repository / service
	6. membuat controller
		RunTodoCreate : POST /todo
		GetAllTodo : GET /todo
		RunTodoCheck : PUT /todo/:todo_id
	7. membuat application
		usecase + gateway + controller