package main
import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)
var apis map[int]string
var c chan map[int]interface{}
func fetchData(API int) {
    url := apis[API]
    if resp, err := http.Get(url); err == nil {
        defer resp.Body.Close()
        if body, err := io.ReadAll(resp.Body);
            err == nil {
            var result map[string]interface{}
            json.Unmarshal([]byte(body), &result)
            switch API {
            case 1:   // for the data.fixer.io/api/API
                if result["success"] == true {
                    fmt.Println(result["rates"].(
                        map[string]interface{})["USD"])
                } else {
                    fmt.Println(result["error"].(
                        map[string]interface{})["info"])
}
            case 2:  // for the openweathermap.org API
                if result["main"] != nil {
                    fmt.Println(result["main"].(
                        map[string]interface{})["temp"])
                } else {
                    fmt.Println(result["message"])
}
}
} else {
            log.Fatal(err)
        }
    } else {
        log.Fatal(err)


		}
	

	}
func main() {
    apis = make(map[int]string)
   
apis[1] = "http://data.fixer.io/api/latest?access_key=" + "940c85855dfd5f6943db08da3c0b4d8a"
apis[2] = "http://api.openweathermap.org/data/2.5/weather?" + "q=SINGAPORE&appid=5173174765098f8144be5d150fdab03f"
go fetchData(1) 
go fetchData(2)
    fmt.Scanln()
   
    // we expect two results in the channel
    for i := 0; i < 2; i++ {
        fmt.Println(<-c)
    }
    fmt.Println("Done!")
    fmt.Scanln()
}