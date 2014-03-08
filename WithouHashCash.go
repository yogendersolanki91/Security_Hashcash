package main

import (
	"HashCashProject"
	"time"
	"fmt"
	"strconv"
)



var Serevre [10]node;
type node struct{
 	timeAfter int;
	mynode HashCashProject.ServerObj;
	last time.Time;
	totalsend int;
}
func main(){
	for i:=0;i<10;i++{
	Serevre[i].mynode=HashCashProject.New(i+1,"/home/blackh/IdeaProjects/untitled/src/config.json");

		Serevre[i].last=time.Now();

			Serevre[i].timeAfter=20
				go simpleCommunication(&Serevre[i])
	}
	time.Sleep(time.Hour)

}


func simpleCommunication(n *node){
	for {
		select {

		case <-time.After(100*time.Millisecond):
		go time.AfterFunc(time.Millisecond*time.Duration(n.timeAfter),func(){
		n.mynode.Outbox()<-&HashCashProject.Envelope{Pid: -1, MsgId: 2, Msg: "MEsaage"}
		if time.Now().Sub(n.last).Seconds()>5 {
			n.last=time.Now();
			fmt.Println(strconv.Itoa(n.mynode.ID)+"		"+strconv.Itoa(n.totalsend))
			n.totalsend=0;
		}else{

			n.totalsend++;
		}})
	//	select {


		case <-n.mynode.Inbox():
			//var i float64
			/*for i=0;i<23242;i++{
			 math.Sqrt(i*i*8*i*i*i*math.Sqrt(i*i*i*i*i*i*i*i*i*i*i*i*i*i)*i*i*i*i*math.Sqrt(i*i*i*i*i*i*i*i*i*i*i*i*i*i)*i*i*i*i*i*i*i*i*math.Sqrt(i*i*i*i*i*i*i*i*i*i*i*i*i*i))*/
		}
}



	}



