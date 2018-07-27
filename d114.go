package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"time"
)



/*
import (
	"sync"
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a' +26 ;char ++ {
				fmt.Printf("%c",char)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i:= 0; i< 3;i++ {
			for char := 'A'; char < 'A' + 26;char++ {
				fmt.Printf("%c",char)
			}
		}
	}()

	wg.Wait()
}
*/

/*
import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/astaxie/beego"
	"time"
	"strings"
	"fmt"
)
func httpDo() {
	client := &http.Client{}
    url := "http://cmdb.jd.com/v1.0/device/info"
    var payload = make(map[string]interface{})
    var p = make(map[string]string)
    p["field"] = "op_group"
    p["value"] = "JC_PIDLB"
	payload["search"] = []map[string]string{p}
	req, err := http.NewRequest("POST", url, strings.NewReader(p))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpdo() {

	arkurl := "http://127.0.0.1:8090"
	beego.Debug("The arkurl here is : ", arkurl)
	request, err := http.NewRequest("GET", arkurl, nil)
	if err != nil {
		beego.Error("Failed !", err)

	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		beego.Debug("Failed !", err)

	}
	defer resp.Body.Close()
	var rawData interface{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBytes, &rawData); err != nil {
		beego.Debug(err)
	}
}

func main() {

	httpdo()
	time.Sleep(20*time.Second)
	beego.Debug("done")
}
*/


func Hget(ip string) {
	//req, err := http.NewRequest("GET", "http://scmdb.jdcloud.com/v1/arkall/getServerDictByCondition/?eth0=%2010.237.61.25", nil)
	//targetUrl := "http://127.0.0.1:8090/v1/arkall/getServerDictByCondition/?eth0=" + ip
	targetUrl := "http://127.0.0.1:8090/v1/arkall/getServerDictByCondition/?eidc=" + ip
	//targetUrl := "http://scmdb.jdcloud.com/v1/arkall/getServerDictByCondition/?eth0=" + ip
	req, err := http.NewRequest("GET",targetUrl , nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}



	fmt.Println(req.URL.String())

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}


	defer resp.Body.Close()

}
func main() {
be:
//ips := []string{"10.237.2.2","10.237.2.3","10.237.2.4","10.237.2.5","10.237.2.6","10.237.2.7","10.237.2.8","10.237.2.9","10.237.2.10"}
	ips := []string{"prod_bj02","prod_bj03"}
	for _,i := range ips {
		Hget(i)
		time.Sleep(3 * time.Second)
		log.Print(time.Now())
	}

	goto be
}
