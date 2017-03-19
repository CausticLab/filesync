package config

import (
	"os"
	"log"
	"database/sql"
	"filesync/api"
	"filesync/index"
	vars "filesync/vars"
	"github.com/howeyc/fsnotify"
	_ "github.com/mattn/go-sqlite3"
)

func StartServer() {
	vars := vars.GetConfig();

	for _, v := range vars.Monitors {
		watcher, _ := fsnotify.NewWatcher()
		monitored, _ := v.(string)
		monitored = index.PathSafe(monitored)

		if(!index.Exists(monitored)){
			log.Println("Path does not exist, creating: ", monitored)

      if os.MkdirAll(monitored, os.ModePerm) != nil {
      	log.Fatal("Could not create directory: ", monitored)
      	continue
      }
		}

		db, err := sql.Open("sqlite3", index.SlashSuffix(monitored)+".sync/index.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.Exec("VACUUM;")

		index.InitIndex(monitored, db)
		go index.ProcessEvent(watcher, monitored)
		index.WatchRecursively(watcher, monitored, monitored)

	}

	log.Println("Serving now...")
	api.RunWeb(vars.Ip, vars.Port, vars.Monitors)
	//watcher.Close()
}
