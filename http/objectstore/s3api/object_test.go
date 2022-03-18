package s3api

import (
	"fmt"
	"github.com/filedag-project/filedag-storage/http/objectstore/utils/testsign"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestS3ApiServer_PutObjectHandler(t *testing.T) {
	u := "http://127.0.0.1:9985/test/1.txt"
	req := testsign.MustNewSignedV4Request(http.MethodPut, u, 0, nil, "s3", t)

	//req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
