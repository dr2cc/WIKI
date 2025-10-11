package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// **// Method 2: (using http.Handler) Structs Implementing http.Handler
// Another approach is defining a struct that includes the dependencies as fields
// and then implementing the ServeHTTP method on this struct.
// Другой подход заключается в создании (=определении) структуры,
// которая в качестве типов полей будет иметь "зависимости"
// (другой пользовательский тип или его поля!),
// а затем реализации метода ServeHTTP для этой структуры.
// ServeHTTP это единственный метод интерфейса Handler.
// 📌

// Любой тип, реализующий метод ServeHTTP(ResponseWriter, *Request),
// считается http.Handler
type app struct {
	logger *slog.Logger
}

// Вообще (пока спорно):
// чтобы функция считалась ручкой (handler), она должна реализовывать метод ServeHTTP()
// со следующей сигнатурой:
// ServeHTTP(w http.ResponseWriter, r *http.Request)
//
// А вот это точно.
// http.Handler:
// Это интерфейс, который определяет один метод: ServeHTTP(ResponseWriter, *Request).
// Любой тип, реализующий этот (ServeHTTP) метод, считается http.Handler.
func (app *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// // Log у slog это сложный метод..
	//app.Logger.Log("Received a request")
	app.logger.Info("Method 2 processed. Received a request")
	fmt.Fprintln(w, "Method 2 processed. Request logged(зарегистрирован) successfully(успешно)")
}

// Довожу до рабочего варианта
// По шаблону статьи создаю здесь экземпляр (объект) логгера (slog)
func newLogger() *slog.Logger {
	// скопировал у Тузова, не понимаю, что где значит
	// Вроде как os.Stdout это выходной поток (даже толком не знаю, что это)
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

//*// Method 2

//**// Method 1: Using Closure
// to Capture External Variables
// Использование замыкания (а по сути ручку делаем методом структуры Env)
// для захвата внешних переменных
// (сервис Env становится частью ❗"состояния"❗ клиента (ручки myHandler) и доступен для использования!!!)

type env struct {
	db *sql.DB
}

// Переделал в ServeHTTP
func (e *env) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := e.db.Ping()
	if err != nil {
		panic(err)
		// fmt.Fprintf(w, "Ping db details: %+v", err)
		// return
	}

	fmt.Fprintf(w, "Ping db details: %+v", "Method 1 processed!")
}

// func (e *env) myHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// You can (вы можете) now use e.db in your handler
// 		err := e.db.Ping() // Example function call
// 		if err != nil {
// 			fmt.Fprintf(w, "Ping db details: %+v", err)
// 			return
// 		}

// 		fmt.Fprintf(w, "Ping db details: %+v", "OK")
// 	}
// }

// Довожу до рабочего варианта
// По шаблону статьи создаю здесь подключение к sqlite
func initializeDatabase() *sql.DB {
	// Подключаемся к БД
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		return nil
	}
	return db
}

//*// Method 1: Using Closure

func main() {
	// ❗Cуть методов 1 и 2❗
	// Благодаря реализации ServeHTTP как метода структуры
	// (можно использовать и http.Handler и http.HandlerFunc)
	// обработчик ("ручка") может работать с ее полями и производить внутри себя не только действия с http запросом,
	// но и с "зависимостями" - подключаться к db, передавать данные логгеру и т.д..
	// Все эти данные становятся частью "состояния" функции обработчика.

	//**// Method 1

	env := &env{
		db: initializeDatabase(),
	}

	// мой
	http.Handle("/endpoint", env)
	// // из примера
	// http.HandleFunc("/endpoint", env.myHandler())
	//*// Method 1

	//**// Method 2
	// Создаем объект slog
	logger := newLogger()
	app := &app{
		logger: logger,
	}

	http.Handle("/", app)
	//*// Method 2

	// Общая часть примера
	http.ListenAndServe("localhost:8080", nil)
}

// Each method has its use cases and choosing the right one depends on your specific requirements.
// Каждый метод имеет свои варианты применения, и выбор подходящего метода зависит от ваших конкретных требований.
// For long-lived dependencies like database connections or configuration data,
// using closures or struct methods is typically preferred.
// Для долгоживущих зависимостей, таких как подключения к базе данных или данные конфигурации,
// обычно предпочтительнее использовать замыкания или методы структур.

// For request-scoped data, leveraging the context package can be more appropriate.
// Для данных, относящихся в области запроса, более целесообразным может оказаться использование пакета context.
