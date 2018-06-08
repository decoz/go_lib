package main
import(
  	"github.com/decoz/go_lib/dotsrv"
    "github.com/decoz/go_lib/dotsrv/dmail"
    "runtime"
)

func main() {

  runtime.GOMAXPROCS(1000)

  dm := dmail.New()
  dm.Load("mail.acc")
  dotsrv.Lunch( 8515,  dm )

}
