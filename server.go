package main

import (
  "database/sql"
  "github.com/valyala/fasthttp"

  _ "github.com/mattn/go-sqlite3"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

)

type (
  Users struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Surname string `json:"surname"`
    Email string `json:"email"`
  }
  Data struct {
    Name string `json:"name"`
    Value string `json:"value"`
  }
)

func main() {
  // Echo instance
  e := echo.New()

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
  }))

  // Route => handler
  e.POST("/api/postUser", PostUser)
	e.GET("/api/getUsers", GetUsers)
	e.PUT("/api/updateUser/:id", UpdateUser)
	e.DELETE("/api/deleteUser/:id", DeleteUser)

  e.Logger.Fatal(e.Start(":4000"))
}

func GetUsers (c echo.Context) error {
  db, err := sql.Open("sqlite3", "./storage.db")
  checkErr(err)

  defer db.Close()

  res, err := db.Query("SELECT uid, name, surname, email FROM users WHERE uid > 0")
  checkErr(err)

  var(
    user Users;
    users []Users;
  )

  for res.Next(){
    err := res.Scan(&user.Id, &user.Name, &user.Surname, &user.Email)
    checkErr(err)

    users = append(users, user)
  }
  defer res.Close()

  return c.JSON(fasthttp.StatusOK, users)
}

func PostUser (c echo.Context) error {
  db, err := sql.Open("sqlite3", "./storage.db")
  checkErr(err)

  stmt, err := db.Prepare("INSERT INTO users (name, surname, email) VALUES(?, ?, ?)")
  checkErr(err)

  res, err := stmt.Exec("name", "surname", "email@mail.com")
  checkErr(err)

  _ = res

  db.Close()

  return c.JSON(fasthttp.StatusOK, nil)
}

func DeleteUser (c echo.Context) error {
  id := c.Param("id")

  db, err := sql.Open("sqlite3", "./storage.db")
  checkErr(err)

  stmt, err := db.Prepare("DELETE FROM users WHERE uid = ?")
  checkErr(err)

  res, err := stmt.Exec(id)
  checkErr(err)

  res, err = db.Exec("UPDATE SQLITE_SEQUENCE SET SEQ = 0 WHERE NAME = 'users'")
  checkErr(err)

  _ = res

  db.Close()

  return c.JSON(fasthttp.StatusOK, nil)
}

func UpdateUser (c echo.Context) error {
  id := c.Param("id")
  d := new(Data)
  c.Bind(d)
  db, err := sql.Open("sqlite3", "./storage.db")
  checkErr(err)

  stmt, err := db.Prepare("UPDATE users SET " + d.Name + "=? WHERE uid=?")
  checkErr(err)

  res, err := stmt.Exec(d.Value, id)
  checkErr(err)

  _ = res

  db.Close()

  return c.JSON(fasthttp.StatusOK, nil)
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
