package main

import (
    "fmt"
    "os"
    "time"
    
    "golang.org/x/exp/slog"
    
    "github.com/SpaceTent/db/mysql"
)

type UpdatePerson struct {
    Id      int       `db:"column=id primarykey=yes table=Users"`
    Name    string    `db:"column=name"`
    Dtadded time.Time `db:"column=dtadded omit=yes"`
    Status  int       `db:"column=status"`
}

func main() {
    
    DSN := ""
    textHandler := slog.NewTextHandler(os.Stdout, nil)
    l := slog.New(textHandler)
    
    db, err := mysql.New(DSN, *l)
    if err != nil {
        l.Error(err.Error())
    }
    
    p := UpdatePerson{
        Id:      12,
        Name:    "Test",
        Dtadded: time.Now(),
        Status:  1,
    }
    
    sql, err := db.Update(p)
    if err != nil {
        l.Error(err.Error())
    }
    fmt.Println(sql)
    
}