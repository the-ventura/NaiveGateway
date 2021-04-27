package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// MigrationRecord representation in the database
type MigrationRecord struct {
	ID    int64
	Count int64  `pg:"default:0"`
	Name  string `pg:",unique"`
}

func (u MigrationRecord) String() string {
	return fmt.Sprintf("Migration<%d %d %s>", u.ID, u.Count, u.Name)
}

func runCustomMigrations(migrationsDirectory, instruction, migrationName string, target int) {
	files, err := ioutil.ReadDir(migrationsDirectory)
	if err != nil {
		log.Error(err)
		return
	}

	if instruction == "create" {
		createNewMigration(migrationsDirectory, migrationName)
	} else {
		migrate(files, target, migrationsDirectory)
	}
}

func enforceBounds(fileLength, offset, target int) (int, int) {
	if target > fileLength/2 {
		target = fileLength/2 - 1
	}
	if target < 0 {
		target = -1
	}
	if offset > fileLength/2 {
		offset = fileLength/2 - 1
	}
	if offset < 0 {
		offset = -1
	}
	return offset, target
}

func fileExecutionOrder(files []os.FileInfo, offset, target int) ([]string, int, int) {
	var result []string
	var relevantFiles []os.FileInfo
	offset, target = enforceBounds(len(files), offset, target)
	if offset == target {
		return result, offset, target
	}
	if offset < target {
		relevantFiles = sortMigrations(filterFilesByDirection(files, "up")[offset+1:target+1], "up")
	} else {
		relevantFiles = sortMigrations(filterFilesByDirection(files, "down")[target+1:offset+1], "down")
	}

	for idx := range relevantFiles {
		result = append(result, relevantFiles[idx].Name())
	}
	return result, offset, target
}

func runMigration(db *pg.DB, directory, fileName string) error {
	filePath := path.Join(directory, fileName)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(data))
	if err != nil {
		return err
	}
	return nil
}

func migrate(files []os.FileInfo, target int, directory string) {
	var lastFileName string
	db := NewConnection()
	lastMigrationInfo := getLastMigrationInfo(db)
	offset := int(lastMigrationInfo.Count)
	fileOrder, _, target := fileExecutionOrder(files, offset, target)
	for index, fileName := range fileOrder {
		err := runMigration(db, directory, fileName)
		if err != nil {
			log.Warning(err)
		}
		lastFileName = fileName
		log.Infof("Migrated %d %s", index+1, fileName)
	}
	updateMigrationInfo(db, int64(target), lastFileName)
}

func createMigrationInfoTable(db *pg.DB) {
	err := db.Model((*MigrationRecord)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Error(err)
	}
}

func updateMigrationInfo(db *pg.DB, count int64, name string) {
	info := &MigrationRecord{
		ID:    1,
		Count: count,
		Name:  name,
	}
	_, err := db.Model(info).OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		log.Error(err)
	}
}

func getLastMigrationInfo(db *pg.DB) *MigrationRecord {
	info := &MigrationRecord{
		ID: 1,
	}
	err := db.Model(info).WherePK().Select()
	if err != nil {
		log.Error(err)
	}
	return info
}

func filterFilesByDirection(files []os.FileInfo, requesteDirection string) []os.FileInfo {
	filteredFiles := []os.FileInfo{}
	for idx := range files {
		fileName := files[idx].Name()
		direction := filepath.Ext(strings.TrimSuffix(fileName, filepath.Ext(fileName)))[1:]
		if direction == requesteDirection {
			filteredFiles = append(filteredFiles, files[idx])
		}
	}
	return filteredFiles
}

// ByNumericalFilename sorts filenames by date
type ByNumericalFilename []os.FileInfo

func (nf ByNumericalFilename) Len() int      { return len(nf) }
func (nf ByNumericalFilename) Swap(i, j int) { nf[i], nf[j] = nf[j], nf[i] }
func (nf ByNumericalFilename) Less(i, j int) bool {

	// Use path names
	pathA := nf[i].Name()
	pathB := nf[j].Name()

	// Grab integer value of each filename by parsing the string and slicing off
	// the extension
	a, err1 := strconv.ParseInt(pathA[0:strings.Index(pathA, "_")], 10, 64)
	b, err2 := strconv.ParseInt(pathB[0:strings.Index(pathB, "_")], 10, 64)

	// If any were not numbers sort lexographically
	if err1 != nil || err2 != nil {
		return pathA < pathB
	}

	// Which integer is smaller?
	return a < b
}

func sortMigrations(files []os.FileInfo, direction string) []os.FileInfo {
	if direction == "up" {
		sort.Sort(ByNumericalFilename(files))
	} else {
		sort.Sort(sort.Reverse(ByNumericalFilename(files)))
	}
	return files
}

func createNewMigration(directory, name string) {
	dateFormat := "20060102150405"
	dateStr := time.Now().Format(dateFormat)
	baseFileName := fmt.Sprintf("%s_%s", dateStr, name)
	up := path.Join(directory, fmt.Sprintf("%s.up.sql", baseFileName))
	down := path.Join(directory, fmt.Sprintf("%s.down.sql", baseFileName))

	emptyFile, err := os.Create(up)
	if err != nil {
		log.Error(err)
	}
	emptyFile.Close()

	emptyFile, err = os.Create(down)
	if err != nil {
		log.Error(err)
	}
	emptyFile.Close()
}

func listMigrations(migrationsDirectory string) ([]string, int) {
	files, err := ioutil.ReadDir(migrationsDirectory)
	if err != nil {
		log.Error(err)
		return []string{}, 0
	}

	db := NewConnection()
	lastMigrationInfo := getLastMigrationInfo(db)
	offset := int(lastMigrationInfo.Count)
	filteredFiles := filterFilesByDirection(files, "up")
	result := []string{}
	for idx := range filteredFiles {
		result = append(result, strings.Split(filteredFiles[idx].Name(), ".up")[0])
	}
	return result, offset
}
