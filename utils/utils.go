package utils

import(
    "strings"
    "path/filepath"
    "fmt"
    "io/ioutil"
    "sort"
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
/*func Getchronotime()(dates []string){
    files, err := ioutil.ReadDir("./markdown/")
    if err != nil{
        fmt.Println(err)
    }
    sort.Slice(files, func(i,j int) bool{
        return files[i].ModTime().After(files[j].ModTime())
    })
    dates = make([]string,0,len(files))
    for _,file := range files{
        f := file.ModTime().Format("2006-01-02 3:4:5 pm")
        dates = append(dates, f)
    }
    return dates

}*/

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
