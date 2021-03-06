/*
* @Author: Ins
* @Date:   2018-10-30 16:21:00
* @Last Modified by:   Ins
* @Last Modified time: 2018-12-12 09:07:59
*/
package rados_io_op

import (
    "bytes"
)

func RadosListObj(cluster_name string, user_name string, conf_file string, pool_name string) (bool, []byte) {
    conn, err := NewConn(cluster_name, user_name, conf_file)
    if err != nil {
        return false, []byte(err.Error() + ",error when invoke a new connection:" + ERR_INFO[err.Error()])
    }
    defer conn.Shutdown()

    // open a pool handle
    ioctx, err := conn.OpenIOContext(pool_name)
    if err != nil {
        return false, []byte(err.Error() + ",error when openIOContext:" + ERR_INFO[err.Error()])
    }
    defer ioctx.Destroy()

    //open a iter
    iter, err := ioctx.Iter()
    if err != nil {
        return false, []byte(err.Error() + ",error when openIter:" + ERR_INFO[err.Error()])
    }
    defer iter.Close()

    iter.Next()
    oids := []byte(iter.Value())
    for iter.Next() {
        var buffer bytes.Buffer
        buffer.Write(oids)
        buffer.Write([]byte("," + iter.Value()))
        oids = buffer.Bytes()
    }
    return true, oids
}