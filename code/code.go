package code

import "fmt"
import "math"
import "math/rand"
import "time"
import "strings"
import "unicode/utf8"
import "strconv"
import "bytes"

import "github.com/Omnikron13/gomind/input"

type Code struct {
    l, d int
    c []int
}

//Creates a new Code object with a randomly generated code
//with the supplied length & depth
func New(l, d int) (c Code) {
    c = Code{l, d, make([]int, l)}
    rand.Seed(time.Now().Unix())
    for i := range c.c {
        c.c[i] = rand.Int() % d
    }
    return
}

//Creates a new Code object from a string of characters
//representing the code, and an explicit depth
func Load(s string, d int) (c Code) {
    l := utf8.RuneCountInString(s)
    c = Code{l, d, make([]int, l)}
    rn := strings.Split(s, "")
    for x := range rn {
        i, _ := strconv.ParseInt(rn[x], 16, 0)
        c.c[x] = int(i)
    }
    return
}

func (c Code) Length() int {
    return c.l
}
func (c Code) Depth() int {
    return c.d
}

func (c Code) String() string {
    var b bytes.Buffer
    for _, x := range c.c {
        b.WriteString(fmt.Sprintf("%X", x))
    }
    return b.String()
}

func (c Code) Read(p string) (rtn []int){
    io := input.GetReader()
    Start: for {
        s := io(p)
        //Check for incorrect try length...
        if utf8.RuneCountInString(s) != c.l {
            fmt.Printf("Code length is %d\n", c.l)
            continue Start
        }
        //Prep for parsing...
        rn := strings.Split(s, "")
        rtn = make([]int, c.l)
        //Attempt to parse the code string...
        for i := range rn {
            x, e := strconv.ParseInt(rn[i], 16, 0)
            if e != nil {
                fmt.Println("Code consists of (hex) numbers")
                continue Start
            }
            if int(x) >= c.d {
                fmt.Printf("Code depth is only %d\n", c.d)
                continue Start
            }
            rtn[i] = int(x)
        }
        return
    }
}

func (c Code) Try(g []int) (b, w int) {
    //Check black pegs...
    bp := make([]bool, c.l)
    for x := range c.c {
        if c.c[x] == g[x] {
            b++
            bp[x] = true
        }
    }
    //Check white pegs...
    wp := make([]bool, c.l)
    for x := range c.c {
        if !bp[x] {
            for y := range c.c {
                if !bp[y] && !wp[y] && g[x] == c.c[y] {
                    w++
                    wp[y] = true
                }
            }
        }
    }
    //fmt.Println(bp, wp) //Cheap debug...
    return
}

//Return 
func (c Code) Score(t int) (s, d float64) {
    //Calculate 'difficulty' - number of possible codes
    d = math.Pow(float64(c.Depth()), float64(c.Length()))
    //Calculate 'score' - d / 'tries'
    s = d / float64(t)
    return
}
