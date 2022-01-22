package process

import (
	"client/logger"
	"fmt"
	"os"
)

type UserProcess struct {

}

// 登陆成功菜单显示：
func showAfterLoginMenu() {
	logger.Info("\n----------------login succeed!----------------\n")
	logger.Info("\t\tselect what you want to do\n")
	logger.Info("\t\t1. Show all online users\n")
	logger.Info("\t\t2. Send group message\n")
	logger.Info("\t\t3. Point-to-point communication\n")
	logger.Info("\t\t4. Exit\n")
	var key int
	//var content string
	//var inputReader *bufio.Reader
	//var err error
	//inputReader = bufio.NewReader(os.Stdin)

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//messageProcess := MessageProcess{}
		//err = messageProcess.GetOnlineUerList()
		//if err != nil {
		//	logger.Error("Some error occurred when get online user list, error: %v\n", err)
		//}
		//return
	case 2:
		//logger.Notice("Say something:\n")
		//content, err = inputReader.ReadString('\n')
		//if err != nil {
		//	logger.Error("Some error occurred when you input, error: %v\n", err)
		//}
		//currentUser := model.CurrentUser
		//messageProcess := MessageProcess{}
		//err = messageProcess.SendGroupMessageToServer(0, currentUser.UserName, content)
		//if err != nil {
		//	logger.Error("Some error occurred when send data to server: %v\n", err)
		//} else {
		//	logger.Success("Send group message succeed!\n\n")
		//}
	case 3:
		//var targetUserName string
		//
		//logger.Notice("Select one friend by user name\n")
		//fmt.Scanf("%s\n", &targetUserName)
		//logger.Notice("Input message:\n")
		//content, err = inputReader.ReadString('\n')
		//if err != nil {
		//	logger.Error("Some error occurred when you input, error: %v\n", err)
		//}
		//messageProcess := MessageProcess{}
		//conn, err := messageProcess.PointToPointCommunication(targetUserName, model.CurrentUser.UserName, content)
		//if err != nil {
		//	logger.Error("Some error occurred when point to point comunication: %v\n", err)
		//	return
		//}
		//
		//errMsg := make(chan error)
		//go Response(conn, errMsg)
		//err = <-errMsg
		//
		//if err.Error() != "<nil>" {
		//	logger.Error("Send message error: %v\n", err)
		//}
	case 4:
		logger.Warn("Exit...\n")
		os.Exit(0)
	default:
		logger.Info("Selected invalid!\n")
	}
}

func (up UserProcess) Login(username string, password string) (err error){
	return
}

func (up UserProcess) Register(userName, password, passwordConfirm string) (err error){
	return
}
