package utils

import(
    "strings"
    "path/filepath"
    "fmt"
    "io/ioutil"
    "sort"
    "github.com/russross/blackfriday"
    "os"
)

func Formatasname(p string)string{
    p = strings.TrimSuffix(p, filepath.Ext(p))
    return fmt.Sprintf(p)
}

func Getchrononames()(names []string){
    path,err := os.Getwd()
    if err !=nil{
        fmt.Println(err)
    }
    files, err := ioutil.ReadDir(path+"/markdown/")
    if err != nil{
        fmt.Println(err)
    }
    sort.Slice(files, func(i,j int) bool{
        return files[i].ModTime().After(files[j].ModTime())
    })
    names = make([]string,0,len(files))
    for _,file := range files{
        names = append(names, file.Name())
    }
    return names

}

func Genpages(){
    path,err := os.Getwd()
    if err !=nil{
        fmt.Println(err)
    }
    files, err := os.ReadDir(path+"/markdown/")
    if err != nil{
        fmt.Println(err)
    }
    for _,file := range files{
        mdfile, err := os.ReadFile(path +"/markdown/" + file.Name())
        if err != nil{
            fmt.Println(err)
        }
        html := blackfriday.MarkdownCommon([]byte(mdfile))
        err = os.WriteFile(path +"/articles/" + Formatasname(file.Name())+".html", html, 0644)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func Formatasdate(name string)(string){
    path,err := os.Getwd()
    if err !=nil{
        fmt.Println(err)
    }

    file, err := os.Stat(path+"/markdown/"+name)
    if err != nil{
        fmt.Println(err)
    }
    return file.ModTime().Format("02 January 2006")
}
