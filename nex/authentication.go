package nex

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/resident-evil-revelations-3ds/globals"
)

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(1)
	globals.AuthenticationEndpoint.ServerAccount = globals.AuthenticationServerAccount
	globals.AuthenticationEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.AuthenticationEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername

	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)

	globals.AuthenticationServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(2, 3, 2))
	globals.AuthenticationServer.AccessKey = "f26f5737"

	globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()
		protocol := globals.GetProtocolByID(request.ProtocolID)

		fmt.Println("== Resident Evil: Revelations - Auth ==")
		fmt.Printf("User: %d\n", packet.Sender().PID())
		fmt.Printf("Protocol: %d (%s)\n", request.ProtocolID, protocol.Protocol())
		fmt.Printf("Method: %d (%s)\n", request.MethodID, protocol.GetMethodByID(request.MethodID))
		fmt.Println("===============")
	})

	globals.AuthenticationEndpoint.OnError(func(err *nex.Error) {
		globals.Logger.Errorf("Auth: %v", err)
	})

	registerCommonAuthenticationServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_REVELATIONS_AUTHENTICATION_SERVER_PORT"))

	globals.AuthenticationServer.Listen(port)
}
