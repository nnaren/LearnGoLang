package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
)
// type Head struct{

// }


type MedalsInfoT struct {
    Bronze          string
    Gold            string
    NocBronzeRank   string
    NocGoldRank     string
    NocId           string
    NocLogo         string
    NocName         string
    NocRank         string
    NocShortName    string
    NocSilverRank   string
    NocUrl          string
    Silver          string
    Total           string
}

type MedalsInfoT2 struct {
    Bronze          string
    Gold            string
    NocBronzeRank   string
    NocGoldRank     string
    NocId           string
    NocLogo         string
    NocName         string
    NocRank         string
    NocSilverRank   string
    NocUrl          string
    Silver          string
    Total           string
}
type Descrb struct{
    Desc string
    Endesc string
}
type ListT struct{
    Total []MedalsInfoT `json:"total"`
    Men []MedalsInfoT   `json:"men"`
    Women []MedalsInfoT `json:"women"`
    Mix []MedalsInfoT   `json:"mix"`
}
type Content struct{
    Head []Descrb       
    Data ListT        
    MedalsInfo MedalsInfoT2 
}
type Info struct{
    Code int                   `json:"code"`
    Data Content              `json:"data"`
    Version string            `json:"version"`

}


func main()  {
    for _, url := range os.Args[1:]{
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        b, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        
        // fmt.Printf("%s\n", b)
        // object, err := NewJsonObject(b)
        // fmt.Println(object.GetString("nocShortName"))


        var info Info
        if err := json.Unmarshal(b, &info); err != nil {
            fmt.Println(err.Error())
            return
        }
        fmt.Println("No. 1: " + info.Data.Data.Total[0].NocName + 
        " 金/银/铜/总: " + 
        info.Data.Data.Total[0].Gold + "/" + 
        info.Data.Data.Total[0].Silver + "/" +
        info.Data.Data.Total[0].Bronze + "/" +
        info.Data.Data.Total[0].Total + "." )
        fmt.Println("No. 2: " + 
        info.Data.Data.Total[1].NocName + 
        " 金/银/铜/总: " + 
        info.Data.Data.Total[1].Gold + "/" + 
        info.Data.Data.Total[1].Silver + "/" +
        info.Data.Data.Total[1].Bronze + "/" +
        info.Data.Data.Total[1].Total + "." ) 
        fmt.Println("No. 3: " + 
        info.Data.Data.Total[2].NocName + 
        " 金/银/铜/总: " + 
        info.Data.Data.Total[2].Gold + "/" + 
        info.Data.Data.Total[2].Silver + "/" +
        info.Data.Data.Total[2].Bronze + "/" +
        info.Data.Data.Total[2].Total + "." ) 
        // fmt.Println(info.Data.Head)
        // fmt.Println(info.Data.MedalsInfo)
        // fmt.Println(info.Version)

        // datamap := make(map[string]interface{})
        // //创建解码器
        // decoder := json.NewDecoder(resp.Body)
        // //解码
        // _ = decoder.Decode(&datamap)
        // fmt.Println(datamap)
    }
}

//   ./main https://app.sports.qq.com/tokyoOly/medalsList/?from%5C=h5%5C&medalsType%5C=all