package lock

//
//import (
//	"fmt"
//	"log"
//	"testing"
//)
//
//var db = map[string]string{
//	"t": "630",
//	"j": "589",
//	"s": "567",
//}
//
//func TestGet(t *testing.T) {
//	loadCnts := make(map[string]string, len(db))
//	NewGroup("score", 2<<10, GetterFunc(
//		func(key string) ([]byte, error) {
//			log.Println("[slow-db] key", key)
//			if v, ok := db[key]; ok {
//				if _, ok := loadCnts[key]; !ok {
//					loadCnts[key] = 0
//				}
//				loadCnts[key] += 1
//				return []byte(v), nil
//			}
//			return nil, fmt.Errorf("%s not exist", key)
//		},
//	))
//
//}
