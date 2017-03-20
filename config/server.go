package config

import (
	"database/sql"
	"github.com/causticlab/filesync/api"
	"github.com/causticlab/filesync/index"
	vars "github.com/causticlab/filesync/vars"
	"github.com/howeyc/fsnotify"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func StartServer() {
	vars := vars.GetConfig()

	for _, v := range vars.Monitors {
		watcher, _ := fsnotify.NewWatcher()
		monitored, _ := v.(string)
		monitored = index.PathSafe(monitored)
		targetPath := index.SlashSuffix(monitored)
		dbPath := targetPath + ".sync/index.db"

		if !index.Exists(monitored) {
			log.Println("Path does not exist, creating: ", monitored)

			if os.MkdirAll(monitored, os.ModePerm) != nil {
				log.Fatal("Could not create directory: ", monitored)
				continue
			}
		}

		if !index.Writable(targetPath) {
			log.Fatal("Path is not writeable: ", targetPath)
		}

		db, err := sql.Open("sqlite3", dbPath)
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
