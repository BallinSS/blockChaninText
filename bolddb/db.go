package bolddb

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// 初始化嵌入式数据库
func OpenDb() {
	// func bolt.Open(path string, mode fs.FileMode, options *bolt.Options) (*bolt.DB, error)
	// 参数1:文件名 参数2:文件权限 参数3:该机配置 - nil为使用默认配置
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	// Update(fn func(*bolt.Tx) error) error
	err = db.Update(func(tx *bolt.Tx) error {

		// 创建BlockBucket表
		// func (tx *bolt.Tx) CreateBucket(name []byte) (*bolt.Bucket, error)
		b, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		// 往表里面存储数据
		if b != nil {
			err := b.Put([]byte("l"), []byte("Send 100 BTC To Ballin......"))
			if err != nil {
				log.Panic("数据存储失败......")
			}
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
}
