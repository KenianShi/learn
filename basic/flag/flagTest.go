package main

import (
	"fmt"
	"flag"
)

func main() {
	var id int
	var name string
	var male bool

	fmt.Println("parsed? = ",flag.Parsed())

	//不能命名两个一样的参数
	flag.IntVar(&id,"id",123,"help msg for id")
	//flag.IntVar(&id,"id",123,"help msg for id")
	flag.StringVar(&name,"name","default name","help message for name")
	flag.BoolVar(&male,"male",false,"help msg for male")

	flag.Parse()
	fmt.Println("parsed = ",flag.Parsed())

	fmt.Println("-------------flag.Args return the non-flag commandLine argument ------------------")
	// return the non-flag commandLine argument that not flag in the flag.var...
	// flag.Args 返回那些未绑定的args的值
	for i,v := range flag.Args() {
		fmt.Printf("arg[%d] = (%s). \n",i,v)
	}
	fmt.Println("-------------Args End-------------------")

	fmt.Println("************* visit flag start *********************")
	// visit flag that has set
	//
	flag.Visit(func(f *flag.Flag) {
		fmt.Println(f.Name,f.Value,f.Usage,f.DefValue)
	})
	fmt.Println("************* visit flag end ***************************")

	fmt.Println("=============visitAll flag start====================")
	flag.VisitAll(func(f *flag.Flag){
		fmt.Println(f.Name,f.Value,f.Usage,f.DefValue)
	})
	fmt.Println("=============visitAll flag end======================")

	fmt.Printf("id = %d \n", id)
	fmt.Printf("name = %s \n", name)
	fmt.Printf("male = %b \n", male)

	fmt.Println("------------- printDefault start ------------------------")
	flag.PrintDefaults()
	fmt.Println("------------- PrintDefaults end  ------------------------")

	fmt.Printf("NArg = %d \n",flag.NArg())
	fmt.Printf("NFlag = %d \n",flag.NFlag())

}
