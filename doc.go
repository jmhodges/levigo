/*
Package levigo provides the ability to create and access LevelDB databases.

levigo.Open() opens and creates databases.

	opts := levigo.NewOptions()
	opts.SetCache(levigo.NewLRUCache(3<<30))
	opts.SetCreateIfMissing(true)
	db, err := levigo.Open("/path/to/db", opts)

*DB.Get(), .Put() and .Delete(), respectively, get the data related to a
 single key, put data for a single key into the database, and deletes data for
 a single key. Do not worry about copying the keys and values in the arguments
 and return values of these methods. LevelDB does this for you.

	ro := levigo.NewReadOptions()
	wo := levigo.NewWriteOptions()
	// if ro and wo are not used again, be sure to Close them.

 	data, err := db.Get(ro, []byte("key"))
	...
	err = db.Put(wo, []byte("anotherkey"), data)
	...
	err = db.Delete(wo, []byte("key"))

 For bulk reads, use an Iterator. For ones that you do not want to disturb
your live traffic, be sure to call SetFillCache(false) on the ReadOptions you
use when creating the Iterator.

	ro := levigo.NewReadOptions()
	ro.SetFillCache(false)
	t := db.NewIterator(ro)
	defer it.Close()
	for it; it.Valid(); it.Next() {
		munge(it.Key(), it.Value())
	}

Batching and atomically making writes can be performed with a WriteBatch and
*DB.Write().

	wb := levigo.NewWriteBatch()
	// defer wb.Close() or use wb.Clear() and reuse.
	wb.Delete([]byte("removed"))
	wb.Put([]byte("added"), []byte("data"))
	wb.Put([]byte("anotheradded"), []byte("more"))
	err := db.Write(wo, wb)
*/
package levigo
