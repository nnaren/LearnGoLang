package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
    "strconv"
)

type HeroT struct {
    Name string
    ShortBio string
    Title string
    Keywords string
    Attack string
    Defense string
    Magic string
    Difficulty string
    Hp string
    Hpperlevel string
    Mp string
    Mpperlevel string
    Movespeed string
    Armor string
    Armorperlevel string
    Spellblock string
    Spellblockperlevel string
    Attackrange string
    Hpregen string
    Hpregenperlevel string
    Mpregen string
    Mpregenperlevel string
    Crit string
    Critperlevel string
    Attackdamage string
    Attackdamageperlevel string
    Attackspeed string
    Attackspeedperlevel string
}

type SpellT struct {
    Name string
    Description string
    AbilityIconPath string
}

type Info struct{
    Hero HeroT
    Skins []interface{}
    Spells []SpellT
    Version string
    fileName string
    fileTime string
}
func main()  {
    baseUrl := "https://game.gtimg.cn/images/lol/act/img/js/hero/"
    var url string
    // ids := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,
    //     25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,48,50,51,
    //     53,54,55,56,57,58,59,60,61,62,63,64,67,68,69,72,74,75,76,77,78,79,80,81,
    //     82,83,84,85,86,89,90,91,92,96,98,99,101,102,103,104,105,106,107,110,111,
    //     112,113,114,115,117,119,120,121,122,126,127,131,133,134,136,141,142,143,
    //     145,147,150,154,157,161,163,164,166,201,202,203,222,223,234,235,236,238,
    //     240,245,246,254,266,267,268,350,360,412,420,421,427,429,432,497,498,516,
    //     517,518,523,526,555,777,875,876,887}
    ids := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24}
    num := 0
    // idValid []int
    for  _,i := range ids {
        // strconv.Itoa(int)
        url = baseUrl + strconv.Itoa(i) + ".js";
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
        var info Info
        if err := json.Unmarshal(b, &info); err != nil {
            fmt.Println(err.Error())
            // return
        } else {
            num++
            // idValid = append(idValid, i)
        }
    
        fmt.Println(i)
        fmt.Println(info.Hero.Title)
        fmt.Println("大招：《" + info.Spells[1].Name + "》")
        fmt.Println(info.Spells[1].Description)
        // FormatFloat
        // fmt.Println(strconv.ParseFloat(info.Hero.Spellblockperlevel, 32))

        
    }
    // fmt.Println(idValid)
    fmt.Println(num)
}

// 网址 https://game.gtimg.cn/images/lol/act/img/js/hero/2.js
// 分类：
// 
// 2、英雄形象

// 1、战队信息

// 排名
// 英雄
// 位置
// 胜率
// 登场率
// 被禁率
// 使用次数
// 击杀
// 死亡
// 助攻
// 造成英雄伤害
// 承受英雄伤害
// 总治疗量
// 补刀数
// 击杀敌方野怪
// 击杀己方野怪
// 金钱
