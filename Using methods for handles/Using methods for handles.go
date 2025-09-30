// Method 3: Using the context Package
// For passing request-scoped data, Go’s context package offers a standardized way
// to transport data across API boundaries and between processes.
// Для передачи данных в рамках запроса, context (пакет Go) предлагает стандартизированный способ
// передачи данных через границы API и между процессами.
func yourHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value("yourKey").(YourType) // Type assertion
	fmt.Fprintf(w, "Value from context: %+v", value)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "yourKey", YourValue)
		r = r.WithContext(ctx)
		yourHandler(w, r)
	})
	http.ListenAndServe(":8080", nil)
}

// // Method 2: Structs Implementing http.Handler
// // Another approach is defining a struct that includes the dependencies as fields
// // and then implementing the ServeHTTP method on this struct.
// // Другой подход заключается в создании (=определении) структуры, которая в качестве полей будет иметь зависимости,
// // а затем реализации метода ServeHTTP для этой структуры.
// // 📌
// //
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type App struct {
// 	Logger *Logger
// }

// func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	app.Logger.Log("Received a request")
// 	fmt.Fprintln(w, "Request logged(зарегистрирован) successfully(успешно)")
// }

// func main() {
// 	logger := NewLogger()
// 	app := &App{
// 		Logger: logger,
// 	}
// 	// Как я понимаю отличие от первого метода только в возможности использования
// 	// http.Handle вместо http.HandleFunc
// 	http.Handle("/", app)
// 	http.ListenAndServe(":8080", nil)
// }

// // Method 1: Using Closure to Capture External Variables
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type Env struct {
// 	db Database
// }

// func main() {
// 	env := &Env{
// 		db: InitializeDatabase(),
// 	}
// 	http.HandleFunc("/endpoint", env.handler())
// 	http.ListenAndServe(":8080", nil)
// }

// func (e *Env) handler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// You can (вы можете) now use e.db in your handler
// 		user := e.db.GetUser() // Example function call
// 		fmt.Fprintf(w, "User Details: %+v", user)
// 	}
// }

// Each method has its use cases and choosing the right one depends on your specific requirements.
// Каждый метод имеет свои варианты применения, и выбор подходящего метода зависит от ваших конкретных требований.
// For long-lived dependencies like database connections or configuration data,
// using closures or struct methods is typically preferred.
// Для долгоживущих зависимостей, таких как подключения к базе данных или данные конфигурации,
// обычно предпочтительнее использовать замыкания или методы структур.
//
// For request-scoped data, leveraging the context package can be more appropriate.
// Для данных, относящихся в области запроса, более целесообразным может оказаться использование пакета context.

