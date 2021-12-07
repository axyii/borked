package utils

import(
    "strings"
    "path/filepath"
    "io"
    "fmt"
    "bufio"
    "sort"
    "github.com/russross/blackfriday"
    "os"
    "time"
    "regexp"
)
type pageinfo struct{
    pagename string
    date time.Time
}

func Formatasname(p string)string{
    p = strings.TrimSuffix(p, filepath.Ext(p))
    return fmt.Sprintf(p)
}

func Genpages(){
    path,err := os.Getwd()
    if err !=nil{
        Logger.Println(err)
    }
    files, err := os.ReadDir(path+"/markdown/")
    if err != nil{
        Logger.Println(err)
    }
    for _,file := range files{
        mdfile, err := os.ReadFile(path +"/markdown/" + file.Name())
        if err != nil{
            Logger.Println(err)
        }
        html := blackfriday.MarkdownCommon([]byte(mdfile))
        err = os.WriteFile(path +"/articles/" + Formatasname(file.Name())+".html", html, 0644)
        if err != nil {
            Logger.Println(err)
        }
    }
}

func Formatasdate(name string)(string){
    wd := "/markdown/"
    path,err := os.Getwd()
    if err !=nil{
        Logger.Println(err)
    }
    newname := Formatasname(name)
    mainfile,err := os.Open(path+wd+newname+".md")
    defer mainfile.Close()
    if err != nil{
        Logger.Println(err)
    }
    datex,err := ReadLine(mainfile,1)
    if err != nil{
        Logger.Println(err)
    }
    actualdates := findconvert(datex)
    return actualdates.Format("02 January 2006")
}

func Getchrononames()([]string){
    wd := "/markdown/"
    path, err := os.Getwd()
    if err != nil{
        Logger.Println(err)
    }
    files, err := os.ReadDir(path + wd )
    if err != nil{
        Logger.Println(err)
    }
    nof := len(files)
    var allpageinfo = make([]pageinfo,0,nof)
    for _,file := range files{
        mainfile,err := os.Open(path + wd +file.Name())
        defer mainfile.Close()
        if err != nil{
            Logger.Println(err)
        }
        datex,err := ReadLine(mainfile,1)
        if err != nil{
            Logger.Println(err)
        }
        actualdates := findconvert(datex)
        filename := fmt.Sprint(Formatasname(file.Name())+".html")
        allpageinfo = append(allpageinfo,pageinfo{pagename:filename,date:actualdates})
    }
    sort.Slice(allpageinfo, func(i, j int) bool { return allpageinfo[i].date.After(allpageinfo[j].date) })
    names := make([]string,0,nof)
    for i := 0; i < nof; i++{
        names = append(names,allpageinfo[i].pagename)
    }
    return names

}

func findconvert(datex string)(time.Time){
    layoutISO := "2006-01-02"
    re := regexp.MustCompile(`\d\d\d\d-\d\d-\d\d`) 
    datestring := re.FindString(datex)
    actualdates,err := time.Parse(layoutISO,datestring)
    if err != nil{
        Logger.Println(err)
    }
    return actualdates

}

func ReadLine(r io.Reader, lineNum int) (line string, err error) {
    sc := bufio.NewScanner(r)
    lastLine := 0
    for sc.Scan() {
        lastLine++
        if lastLine == lineNum {
            return sc.Text(), sc.Err()
        }
    }
    return line,io.EOF
}
